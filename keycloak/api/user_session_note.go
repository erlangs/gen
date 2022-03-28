package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configUserSessionNoteRouter(router *httprouter.Router) {
	router.GET("/usersessionnote", GetAllUserSessionNote)
	router.POST("/usersessionnote", AddUserSessionNote)
	router.GET("/usersessionnote/:argUserSession/:argName", GetUserSessionNote)
	router.PUT("/usersessionnote/:argUserSession/:argName", UpdateUserSessionNote)
	router.DELETE("/usersessionnote/:argUserSession/:argName", DeleteUserSessionNote)
}

func configGinUserSessionNoteRouter(router gin.IRoutes) {
	router.GET("/usersessionnote", ConverHttprouterToGin(GetAllUserSessionNote))
	router.POST("/usersessionnote", ConverHttprouterToGin(AddUserSessionNote))
	router.GET("/usersessionnote/:argUserSession/:argName", ConverHttprouterToGin(GetUserSessionNote))
	router.PUT("/usersessionnote/:argUserSession/:argName", ConverHttprouterToGin(UpdateUserSessionNote))
	router.DELETE("/usersessionnote/:argUserSession/:argName", ConverHttprouterToGin(DeleteUserSessionNote))
}

// GetAllUserSessionNote is a function to get a slice of record(s) from user_session_note table in the keycloak database
// @Summary Get list of UserSessionNote
// @Tags UserSessionNote
// @Description GetAllUserSessionNote is a handler to get a slice of record(s) from user_session_note table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.UserSessionNote}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /usersessionnote [get]
// http "http://localhost:8080/usersessionnote?page=0&pagesize=20" X-Api-User:user123
func GetAllUserSessionNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "user_session_note", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllUserSessionNote(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetUserSessionNote is a function to get a single record from the user_session_note table in the keycloak database
// @Summary Get record from table UserSessionNote by  argUserSession  argName
// @Tags UserSessionNote
// @ID argUserSession
// @ID argName
// @Description GetUserSessionNote is a function to get a single record from the user_session_note table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argUserSession path string true "user_session"
// @Param  argName path string true "name"
// @Success 200 {object} model.UserSessionNote
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /usersessionnote/{argUserSession}/{argName} [get]
// http "http://localhost:8080/usersessionnote/hello world/hello world" X-Api-User:user123
func GetUserSessionNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserSession, err := parseString(ps, "argUserSession")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_session_note", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetUserSessionNote(ctx, argUserSession, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddUserSessionNote add to add a single record to user_session_note table in the keycloak database
// @Summary Add an record to user_session_note table
// @Description add to add a single record to user_session_note table in the keycloak database
// @Tags UserSessionNote
// @Accept  json
// @Produce  json
// @Param UserSessionNote body model.UserSessionNote true "Add UserSessionNote"
// @Success 200 {object} model.UserSessionNote
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /usersessionnote [post]
// echo '{"user_session": "nTxVYlqtKXhimdjsDSMeBMGhU","name": "wFWtRGHomOnepPNeljOAVSevv","value": "ptkGrmdUPjQSRAOCORlKPKxEE"}' | http POST "http://localhost:8080/usersessionnote" X-Api-User:user123
func AddUserSessionNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	usersessionnote := &model.UserSessionNote{}

	if err := readJSON(r, usersessionnote); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := usersessionnote.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	usersessionnote.Prepare()

	if err := usersessionnote.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_session_note", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	usersessionnote, _, err = dao.AddUserSessionNote(ctx, usersessionnote)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, usersessionnote)
}

// UpdateUserSessionNote Update a single record from user_session_note table in the keycloak database
// @Summary Update an record in table user_session_note
// @Description Update a single record from user_session_note table in the keycloak database
// @Tags UserSessionNote
// @Accept  json
// @Produce  json
// @Param  argUserSession path string true "user_session"// @Param  argName path string true "name"
// @Param  UserSessionNote body model.UserSessionNote true "Update UserSessionNote record"
// @Success 200 {object} model.UserSessionNote
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /usersessionnote/{argUserSession}/{argName} [put]
// echo '{"user_session": "nTxVYlqtKXhimdjsDSMeBMGhU","name": "wFWtRGHomOnepPNeljOAVSevv","value": "ptkGrmdUPjQSRAOCORlKPKxEE"}' | http PUT "http://localhost:8080/usersessionnote/hello world/hello world"  X-Api-User:user123
func UpdateUserSessionNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserSession, err := parseString(ps, "argUserSession")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	usersessionnote := &model.UserSessionNote{}
	if err := readJSON(r, usersessionnote); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := usersessionnote.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	usersessionnote.Prepare()

	if err := usersessionnote.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_session_note", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	usersessionnote, _, err = dao.UpdateUserSessionNote(ctx,
		argUserSession, argName,
		usersessionnote)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, usersessionnote)
}

// DeleteUserSessionNote Delete a single record from user_session_note table in the keycloak database
// @Summary Delete a record from user_session_note
// @Description Delete a single record from user_session_note table in the keycloak database
// @Tags UserSessionNote
// @Accept  json
// @Produce  json
// @Param  argUserSession path string true "user_session"// @Param  argName path string true "name"
// @Success 204 {object} model.UserSessionNote
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /usersessionnote/{argUserSession}/{argName} [delete]
// http DELETE "http://localhost:8080/usersessionnote/hello world/hello world" X-Api-User:user123
func DeleteUserSessionNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserSession, err := parseString(ps, "argUserSession")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_session_note", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteUserSessionNote(ctx, argUserSession, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
