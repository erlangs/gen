package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configUserSessionRouter(router *httprouter.Router) {
	router.GET("/usersession", GetAllUserSession)
	router.POST("/usersession", AddUserSession)
	router.GET("/usersession/:argID", GetUserSession)
	router.PUT("/usersession/:argID", UpdateUserSession)
	router.DELETE("/usersession/:argID", DeleteUserSession)
}

func configGinUserSessionRouter(router gin.IRoutes) {
	router.GET("/usersession", ConverHttprouterToGin(GetAllUserSession))
	router.POST("/usersession", ConverHttprouterToGin(AddUserSession))
	router.GET("/usersession/:argID", ConverHttprouterToGin(GetUserSession))
	router.PUT("/usersession/:argID", ConverHttprouterToGin(UpdateUserSession))
	router.DELETE("/usersession/:argID", ConverHttprouterToGin(DeleteUserSession))
}

// GetAllUserSession is a function to get a slice of record(s) from user_session table in the keycloak database
// @Summary Get list of UserSession
// @Tags UserSession
// @Description GetAllUserSession is a handler to get a slice of record(s) from user_session table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.UserSession}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /usersession [get]
// http "http://localhost:8080/usersession?page=0&pagesize=20" X-Api-User:user123
func GetAllUserSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "user_session", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllUserSession(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetUserSession is a function to get a single record from the user_session table in the keycloak database
// @Summary Get record from table UserSession by  argID
// @Tags UserSession
// @ID argID
// @Description GetUserSession is a function to get a single record from the user_session table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.UserSession
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /usersession/{argID} [get]
// http "http://localhost:8080/usersession/hello world" X-Api-User:user123
func GetUserSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_session", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetUserSession(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddUserSession add to add a single record to user_session table in the keycloak database
// @Summary Add an record to user_session table
// @Description add to add a single record to user_session table in the keycloak database
// @Tags UserSession
// @Accept  json
// @Produce  json
// @Param UserSession body model.UserSession true "Add UserSession"
// @Success 200 {object} model.UserSession
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /usersession [post]
// echo '{"id": "EEIBpsZOLPfbyrCssonwtDNGH","auth_method": "oZPrVqkTxOrIAlbgOUNEGivvm","ip_address": "QxaSZWdGiltLGVloFgqOFrTBr","last_session_refresh": 67,"login_username": "vdGAKcPZMakqspJVpEfCrwGLe","realm_id": "XmkBreXduOBIcFDaXlEtUYGZX","remember_me": false,"started": 93,"user_id": "oWZLLVPlXoYGRMrSqwJVyvYRX","user_session_state": 95,"broker_session_id": "nBjbsiiQmHLsnkpvRbMrkWlRZ","broker_user_id": "oSNkJGjbAcosWmhBinCQZtpeE"}' | http POST "http://localhost:8080/usersession" X-Api-User:user123
func AddUserSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	usersession := &model.UserSession{}

	if err := readJSON(r, usersession); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := usersession.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	usersession.Prepare()

	if err := usersession.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_session", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	usersession, _, err = dao.AddUserSession(ctx, usersession)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, usersession)
}

// UpdateUserSession Update a single record from user_session table in the keycloak database
// @Summary Update an record in table user_session
// @Description Update a single record from user_session table in the keycloak database
// @Tags UserSession
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  UserSession body model.UserSession true "Update UserSession record"
// @Success 200 {object} model.UserSession
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /usersession/{argID} [put]
// echo '{"id": "EEIBpsZOLPfbyrCssonwtDNGH","auth_method": "oZPrVqkTxOrIAlbgOUNEGivvm","ip_address": "QxaSZWdGiltLGVloFgqOFrTBr","last_session_refresh": 67,"login_username": "vdGAKcPZMakqspJVpEfCrwGLe","realm_id": "XmkBreXduOBIcFDaXlEtUYGZX","remember_me": false,"started": 93,"user_id": "oWZLLVPlXoYGRMrSqwJVyvYRX","user_session_state": 95,"broker_session_id": "nBjbsiiQmHLsnkpvRbMrkWlRZ","broker_user_id": "oSNkJGjbAcosWmhBinCQZtpeE"}' | http PUT "http://localhost:8080/usersession/hello world"  X-Api-User:user123
func UpdateUserSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	usersession := &model.UserSession{}
	if err := readJSON(r, usersession); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := usersession.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	usersession.Prepare()

	if err := usersession.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_session", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	usersession, _, err = dao.UpdateUserSession(ctx,
		argID,
		usersession)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, usersession)
}

// DeleteUserSession Delete a single record from user_session table in the keycloak database
// @Summary Delete a record from user_session
// @Description Delete a single record from user_session table in the keycloak database
// @Tags UserSession
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.UserSession
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /usersession/{argID} [delete]
// http DELETE "http://localhost:8080/usersession/hello world" X-Api-User:user123
func DeleteUserSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_session", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteUserSession(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
