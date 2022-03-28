package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configUserRouter(router *httprouter.Router) {
	router.GET("/user", GetAllUser)
	router.POST("/user", AddUser)
	router.GET("/user/:argID", GetUser)
	router.PUT("/user/:argID", UpdateUser)
	router.DELETE("/user/:argID", DeleteUser)
}

func configGinUserRouter(router gin.IRoutes) {
	router.GET("/user", ConverHttprouterToGin(GetAllUser))
	router.POST("/user", ConverHttprouterToGin(AddUser))
	router.GET("/user/:argID", ConverHttprouterToGin(GetUser))
	router.PUT("/user/:argID", ConverHttprouterToGin(UpdateUser))
	router.DELETE("/user/:argID", ConverHttprouterToGin(DeleteUser))
}

// GetAllUser is a function to get a slice of record(s) from user table in the keycloak database
// @Summary Get list of User
// @Tags User
// @Description GetAllUser is a handler to get a slice of record(s) from user table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.User}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /user [get]
// http "http://localhost:8080/user?page=0&pagesize=20" X-Api-User:user123
func GetAllUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "user", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllUser(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetUser is a function to get a single record from the user table in the keycloak database
// @Summary Get record from table User by  argID
// @Tags User
// @ID argID
// @Description GetUser is a function to get a single record from the user table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.User
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /user/{argID} [get]
// http "http://localhost:8080/user/1" X-Api-User:user123
func GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetUser(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddUser add to add a single record to user table in the keycloak database
// @Summary Add an record to user table
// @Description add to add a single record to user table in the keycloak database
// @Tags User
// @Accept  json
// @Produce  json
// @Param User body model.User true "Add User"
// @Success 200 {object} model.User
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /user [post]
// echo '{"id": 2,"user_name": "qjkYehGDIOpGfaAnQOHSXOtNn","account": "GSRErVYejmGapZkgoQwGaIYdo","nickname": "bhUPjoKdKLhMmaBERrviaebNr","password": "xNFCyaCCTmLBbksOFZQwDGCED","salt": "XuKLMqPNBnTGcPKPxyImWMoWh","email": "IdcGXpXvaeRwOhaQyNtbtVoXD","state": 42,"registration_date": 82}' | http POST "http://localhost:8080/user" X-Api-User:user123
func AddUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	user := &model.User{}

	if err := readJSON(r, user); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := user.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	user.Prepare()

	if err := user.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	user, _, err = dao.AddUser(ctx, user)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, user)
}

// UpdateUser Update a single record from user table in the keycloak database
// @Summary Update an record in table user
// @Description Update a single record from user table in the keycloak database
// @Tags User
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  User body model.User true "Update User record"
// @Success 200 {object} model.User
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /user/{argID} [put]
// echo '{"id": 2,"user_name": "qjkYehGDIOpGfaAnQOHSXOtNn","account": "GSRErVYejmGapZkgoQwGaIYdo","nickname": "bhUPjoKdKLhMmaBERrviaebNr","password": "xNFCyaCCTmLBbksOFZQwDGCED","salt": "XuKLMqPNBnTGcPKPxyImWMoWh","email": "IdcGXpXvaeRwOhaQyNtbtVoXD","state": 42,"registration_date": 82}' | http PUT "http://localhost:8080/user/1"  X-Api-User:user123
func UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	user := &model.User{}
	if err := readJSON(r, user); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := user.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	user.Prepare()

	if err := user.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	user, _, err = dao.UpdateUser(ctx,
		argID,
		user)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, user)
}

// DeleteUser Delete a single record from user table in the keycloak database
// @Summary Delete a record from user
// @Description Delete a single record from user table in the keycloak database
// @Tags User
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.User
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /user/{argID} [delete]
// http DELETE "http://localhost:8080/user/1" X-Api-User:user123
func DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteUser(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
