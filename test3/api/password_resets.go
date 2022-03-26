package api

import (
	"net/http"

	"example.com/rest/example/dao"
	"example.com/rest/example/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configPasswordResetsRouter(router *httprouter.Router) {
	router.GET("/passwordresets", GetAllPasswordResets)
	router.POST("/passwordresets", AddPasswordResets)
	router.GET("/passwordresets/:argEmail", GetPasswordResets)
	router.PUT("/passwordresets/:argEmail", UpdatePasswordResets)
	router.DELETE("/passwordresets/:argEmail", DeletePasswordResets)
}

func configGinPasswordResetsRouter(router gin.IRoutes) {
	router.GET("/passwordresets", ConverHttprouterToGin(GetAllPasswordResets))
	router.POST("/passwordresets", ConverHttprouterToGin(AddPasswordResets))
	router.GET("/passwordresets/:argEmail", ConverHttprouterToGin(GetPasswordResets))
	router.PUT("/passwordresets/:argEmail", ConverHttprouterToGin(UpdatePasswordResets))
	router.DELETE("/passwordresets/:argEmail", ConverHttprouterToGin(DeletePasswordResets))
}

// GetAllPasswordResets is a function to get a slice of record(s) from password_resets table in the test1 database
// @Summary Get list of PasswordResets
// @Tags PasswordResets
// @Description GetAllPasswordResets is a handler to get a slice of record(s) from password_resets table in the test1 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.PasswordResets}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /passwordresets [get]
// http "http://localhost:8080/passwordresets?page=0&pagesize=20" X-Api-User:user123
func GetAllPasswordResets(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	pagesize, err := readInt(r, "pagesize", 20)
	if err != nil || pagesize <= 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	order := r.FormValue("order")

	if err := ValidateRequest(ctx, r, "password_resets", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllPasswordResets(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetPasswordResets is a function to get a single record from the password_resets table in the test1 database
// @Summary Get record from table PasswordResets by  argEmail
// @Tags PasswordResets
// @ID argEmail
// @Description GetPasswordResets is a function to get a single record from the password_resets table in the test1 database
// @Accept  json
// @Produce  json
// @Param  argEmail path string true "email"
// @Success 200 {object} model.PasswordResets
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /passwordresets/{argEmail} [get]
// http "http://localhost:8080/passwordresets/hello world" X-Api-User:user123
func GetPasswordResets(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argEmail, err := parseString(ps, "argEmail")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "password_resets", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetPasswordResets(ctx, argEmail)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddPasswordResets add to add a single record to password_resets table in the test1 database
// @Summary Add an record to password_resets table
// @Description add to add a single record to password_resets table in the test1 database
// @Tags PasswordResets
// @Accept  json
// @Produce  json
// @Param PasswordResets body model.PasswordResets true "Add PasswordResets"
// @Success 200 {object} model.PasswordResets
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /passwordresets [post]
// echo '{"email": "CUEMFCPHxqFlrrvfSrDueqxJI","token": "qyOBhtHkVCfkXbOLBKwCvNALE","created_at": "2074-09-23T07:53:04.07671926+08:00"}' | http POST "http://localhost:8080/passwordresets" X-Api-User:user123
func AddPasswordResets(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	passwordresets := &model.PasswordResets{}

	if err := readJSON(r, passwordresets); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := passwordresets.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	passwordresets.Prepare()

	if err := passwordresets.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "password_resets", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	passwordresets, _, err = dao.AddPasswordResets(ctx, passwordresets)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, passwordresets)
}

// UpdatePasswordResets Update a single record from password_resets table in the test1 database
// @Summary Update an record in table password_resets
// @Description Update a single record from password_resets table in the test1 database
// @Tags PasswordResets
// @Accept  json
// @Produce  json
// @Param  argEmail path string true "email"
// @Param  PasswordResets body model.PasswordResets true "Update PasswordResets record"
// @Success 200 {object} model.PasswordResets
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /passwordresets/{argEmail} [put]
// echo '{"email": "CUEMFCPHxqFlrrvfSrDueqxJI","token": "qyOBhtHkVCfkXbOLBKwCvNALE","created_at": "2074-09-23T07:53:04.07671926+08:00"}' | http PUT "http://localhost:8080/passwordresets/hello world"  X-Api-User:user123
func UpdatePasswordResets(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argEmail, err := parseString(ps, "argEmail")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	passwordresets := &model.PasswordResets{}
	if err := readJSON(r, passwordresets); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := passwordresets.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	passwordresets.Prepare()

	if err := passwordresets.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "password_resets", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	passwordresets, _, err = dao.UpdatePasswordResets(ctx,
		argEmail,
		passwordresets)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, passwordresets)
}

// DeletePasswordResets Delete a single record from password_resets table in the test1 database
// @Summary Delete a record from password_resets
// @Description Delete a single record from password_resets table in the test1 database
// @Tags PasswordResets
// @Accept  json
// @Produce  json
// @Param  argEmail path string true "email"
// @Success 204 {object} model.PasswordResets
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /passwordresets/{argEmail} [delete]
// http DELETE "http://localhost:8080/passwordresets/hello world" X-Api-User:user123
func DeletePasswordResets(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argEmail, err := parseString(ps, "argEmail")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "password_resets", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeletePasswordResets(ctx, argEmail)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
