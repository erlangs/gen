package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configUserAttributeRouter(router *httprouter.Router) {
	router.GET("/userattribute", GetAllUserAttribute)
	router.POST("/userattribute", AddUserAttribute)
	router.GET("/userattribute/:argID", GetUserAttribute)
	router.PUT("/userattribute/:argID", UpdateUserAttribute)
	router.DELETE("/userattribute/:argID", DeleteUserAttribute)
}

func configGinUserAttributeRouter(router gin.IRoutes) {
	router.GET("/userattribute", ConverHttprouterToGin(GetAllUserAttribute))
	router.POST("/userattribute", ConverHttprouterToGin(AddUserAttribute))
	router.GET("/userattribute/:argID", ConverHttprouterToGin(GetUserAttribute))
	router.PUT("/userattribute/:argID", ConverHttprouterToGin(UpdateUserAttribute))
	router.DELETE("/userattribute/:argID", ConverHttprouterToGin(DeleteUserAttribute))
}

// GetAllUserAttribute is a function to get a slice of record(s) from user_attribute table in the keycloak database
// @Summary Get list of UserAttribute
// @Tags UserAttribute
// @Description GetAllUserAttribute is a handler to get a slice of record(s) from user_attribute table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.UserAttribute}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userattribute [get]
// http "http://localhost:8080/userattribute?page=0&pagesize=20" X-Api-User:user123
func GetAllUserAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "user_attribute", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllUserAttribute(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetUserAttribute is a function to get a single record from the user_attribute table in the keycloak database
// @Summary Get record from table UserAttribute by  argID
// @Tags UserAttribute
// @ID argID
// @Description GetUserAttribute is a function to get a single record from the user_attribute table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.UserAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /userattribute/{argID} [get]
// http "http://localhost:8080/userattribute/hello world" X-Api-User:user123
func GetUserAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_attribute", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetUserAttribute(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddUserAttribute add to add a single record to user_attribute table in the keycloak database
// @Summary Add an record to user_attribute table
// @Description add to add a single record to user_attribute table in the keycloak database
// @Tags UserAttribute
// @Accept  json
// @Produce  json
// @Param UserAttribute body model.UserAttribute true "Add UserAttribute"
// @Success 200 {object} model.UserAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userattribute [post]
// echo '{"name": "CHtgpiLLHAqsPWgZBcfWQGqpJ","value": "ViAWITKwTMEKsHFJgvPbPosbg","user_id": "aAiVeRwJkuHqVcXtbCMmpCiKE","id": "bGlwSCMOxlkyUiywSTxGnOCJf"}' | http POST "http://localhost:8080/userattribute" X-Api-User:user123
func AddUserAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	userattribute := &model.UserAttribute{}

	if err := readJSON(r, userattribute); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userattribute.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userattribute.Prepare()

	if err := userattribute.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_attribute", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	userattribute, _, err = dao.AddUserAttribute(ctx, userattribute)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userattribute)
}

// UpdateUserAttribute Update a single record from user_attribute table in the keycloak database
// @Summary Update an record in table user_attribute
// @Description Update a single record from user_attribute table in the keycloak database
// @Tags UserAttribute
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  UserAttribute body model.UserAttribute true "Update UserAttribute record"
// @Success 200 {object} model.UserAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userattribute/{argID} [put]
// echo '{"name": "CHtgpiLLHAqsPWgZBcfWQGqpJ","value": "ViAWITKwTMEKsHFJgvPbPosbg","user_id": "aAiVeRwJkuHqVcXtbCMmpCiKE","id": "bGlwSCMOxlkyUiywSTxGnOCJf"}' | http PUT "http://localhost:8080/userattribute/hello world"  X-Api-User:user123
func UpdateUserAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	userattribute := &model.UserAttribute{}
	if err := readJSON(r, userattribute); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userattribute.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userattribute.Prepare()

	if err := userattribute.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_attribute", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	userattribute, _, err = dao.UpdateUserAttribute(ctx,
		argID,
		userattribute)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userattribute)
}

// DeleteUserAttribute Delete a single record from user_attribute table in the keycloak database
// @Summary Delete a record from user_attribute
// @Description Delete a single record from user_attribute table in the keycloak database
// @Tags UserAttribute
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.UserAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /userattribute/{argID} [delete]
// http DELETE "http://localhost:8080/userattribute/hello world" X-Api-User:user123
func DeleteUserAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_attribute", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteUserAttribute(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
