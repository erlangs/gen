package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configClientUserSessionNoteRouter(router *httprouter.Router) {
	router.GET("/clientusersessionnote", GetAllClientUserSessionNote)
	router.POST("/clientusersessionnote", AddClientUserSessionNote)
	router.GET("/clientusersessionnote/:argName/:argClientSession", GetClientUserSessionNote)
	router.PUT("/clientusersessionnote/:argName/:argClientSession", UpdateClientUserSessionNote)
	router.DELETE("/clientusersessionnote/:argName/:argClientSession", DeleteClientUserSessionNote)
}

func configGinClientUserSessionNoteRouter(router gin.IRoutes) {
	router.GET("/clientusersessionnote", ConverHttprouterToGin(GetAllClientUserSessionNote))
	router.POST("/clientusersessionnote", ConverHttprouterToGin(AddClientUserSessionNote))
	router.GET("/clientusersessionnote/:argName/:argClientSession", ConverHttprouterToGin(GetClientUserSessionNote))
	router.PUT("/clientusersessionnote/:argName/:argClientSession", ConverHttprouterToGin(UpdateClientUserSessionNote))
	router.DELETE("/clientusersessionnote/:argName/:argClientSession", ConverHttprouterToGin(DeleteClientUserSessionNote))
}

// GetAllClientUserSessionNote is a function to get a slice of record(s) from client_user_session_note table in the keycloak database
// @Summary Get list of ClientUserSessionNote
// @Tags ClientUserSessionNote
// @Description GetAllClientUserSessionNote is a handler to get a slice of record(s) from client_user_session_note table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ClientUserSessionNote}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientusersessionnote [get]
// http "http://localhost:8080/clientusersessionnote?page=0&pagesize=20" X-Api-User:user123
func GetAllClientUserSessionNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "client_user_session_note", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllClientUserSessionNote(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetClientUserSessionNote is a function to get a single record from the client_user_session_note table in the keycloak database
// @Summary Get record from table ClientUserSessionNote by  argName  argClientSession
// @Tags ClientUserSessionNote
// @ID argName
// @ID argClientSession
// @Description GetClientUserSessionNote is a function to get a single record from the client_user_session_note table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argName path string true "name"
// @Param  argClientSession path string true "client_session"
// @Success 200 {object} model.ClientUserSessionNote
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /clientusersessionnote/{argName}/{argClientSession} [get]
// http "http://localhost:8080/clientusersessionnote/hello world/hello world" X-Api-User:user123
func GetClientUserSessionNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argClientSession, err := parseString(ps, "argClientSession")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_user_session_note", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetClientUserSessionNote(ctx, argName, argClientSession)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddClientUserSessionNote add to add a single record to client_user_session_note table in the keycloak database
// @Summary Add an record to client_user_session_note table
// @Description add to add a single record to client_user_session_note table in the keycloak database
// @Tags ClientUserSessionNote
// @Accept  json
// @Produce  json
// @Param ClientUserSessionNote body model.ClientUserSessionNote true "Add ClientUserSessionNote"
// @Success 200 {object} model.ClientUserSessionNote
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientusersessionnote [post]
// echo '{"name": "oujWrfnwCjCMfktbNlUEClVPx","value": "xlyrNviXkLlXhJpYsMhBYcFFU","client_session": "tlUQBEsrWgJugQsfpGgwvDASO"}' | http POST "http://localhost:8080/clientusersessionnote" X-Api-User:user123
func AddClientUserSessionNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	clientusersessionnote := &model.ClientUserSessionNote{}

	if err := readJSON(r, clientusersessionnote); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientusersessionnote.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientusersessionnote.Prepare()

	if err := clientusersessionnote.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_user_session_note", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	clientusersessionnote, _, err = dao.AddClientUserSessionNote(ctx, clientusersessionnote)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientusersessionnote)
}

// UpdateClientUserSessionNote Update a single record from client_user_session_note table in the keycloak database
// @Summary Update an record in table client_user_session_note
// @Description Update a single record from client_user_session_note table in the keycloak database
// @Tags ClientUserSessionNote
// @Accept  json
// @Produce  json
// @Param  argName path string true "name"// @Param  argClientSession path string true "client_session"
// @Param  ClientUserSessionNote body model.ClientUserSessionNote true "Update ClientUserSessionNote record"
// @Success 200 {object} model.ClientUserSessionNote
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientusersessionnote/{argName}/{argClientSession} [put]
// echo '{"name": "oujWrfnwCjCMfktbNlUEClVPx","value": "xlyrNviXkLlXhJpYsMhBYcFFU","client_session": "tlUQBEsrWgJugQsfpGgwvDASO"}' | http PUT "http://localhost:8080/clientusersessionnote/hello world/hello world"  X-Api-User:user123
func UpdateClientUserSessionNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argClientSession, err := parseString(ps, "argClientSession")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientusersessionnote := &model.ClientUserSessionNote{}
	if err := readJSON(r, clientusersessionnote); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientusersessionnote.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientusersessionnote.Prepare()

	if err := clientusersessionnote.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_user_session_note", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientusersessionnote, _, err = dao.UpdateClientUserSessionNote(ctx,
		argName, argClientSession,
		clientusersessionnote)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientusersessionnote)
}

// DeleteClientUserSessionNote Delete a single record from client_user_session_note table in the keycloak database
// @Summary Delete a record from client_user_session_note
// @Description Delete a single record from client_user_session_note table in the keycloak database
// @Tags ClientUserSessionNote
// @Accept  json
// @Produce  json
// @Param  argName path string true "name"// @Param  argClientSession path string true "client_session"
// @Success 204 {object} model.ClientUserSessionNote
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /clientusersessionnote/{argName}/{argClientSession} [delete]
// http DELETE "http://localhost:8080/clientusersessionnote/hello world/hello world" X-Api-User:user123
func DeleteClientUserSessionNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argClientSession, err := parseString(ps, "argClientSession")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_user_session_note", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteClientUserSessionNote(ctx, argName, argClientSession)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
