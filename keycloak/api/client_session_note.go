package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configClientSessionNoteRouter(router *httprouter.Router) {
	router.GET("/clientsessionnote", GetAllClientSessionNote)
	router.POST("/clientsessionnote", AddClientSessionNote)
	router.GET("/clientsessionnote/:argName/:argClientSession", GetClientSessionNote)
	router.PUT("/clientsessionnote/:argName/:argClientSession", UpdateClientSessionNote)
	router.DELETE("/clientsessionnote/:argName/:argClientSession", DeleteClientSessionNote)
}

func configGinClientSessionNoteRouter(router gin.IRoutes) {
	router.GET("/clientsessionnote", ConverHttprouterToGin(GetAllClientSessionNote))
	router.POST("/clientsessionnote", ConverHttprouterToGin(AddClientSessionNote))
	router.GET("/clientsessionnote/:argName/:argClientSession", ConverHttprouterToGin(GetClientSessionNote))
	router.PUT("/clientsessionnote/:argName/:argClientSession", ConverHttprouterToGin(UpdateClientSessionNote))
	router.DELETE("/clientsessionnote/:argName/:argClientSession", ConverHttprouterToGin(DeleteClientSessionNote))
}

// GetAllClientSessionNote is a function to get a slice of record(s) from client_session_note table in the keycloak database
// @Summary Get list of ClientSessionNote
// @Tags ClientSessionNote
// @Description GetAllClientSessionNote is a handler to get a slice of record(s) from client_session_note table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ClientSessionNote}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientsessionnote [get]
// http "http://localhost:8080/clientsessionnote?page=0&pagesize=20" X-Api-User:user123
func GetAllClientSessionNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "client_session_note", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllClientSessionNote(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetClientSessionNote is a function to get a single record from the client_session_note table in the keycloak database
// @Summary Get record from table ClientSessionNote by  argName  argClientSession
// @Tags ClientSessionNote
// @ID argName
// @ID argClientSession
// @Description GetClientSessionNote is a function to get a single record from the client_session_note table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argName path string true "name"
// @Param  argClientSession path string true "client_session"
// @Success 200 {object} model.ClientSessionNote
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /clientsessionnote/{argName}/{argClientSession} [get]
// http "http://localhost:8080/clientsessionnote/hello world/hello world" X-Api-User:user123
func GetClientSessionNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "client_session_note", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetClientSessionNote(ctx, argName, argClientSession)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddClientSessionNote add to add a single record to client_session_note table in the keycloak database
// @Summary Add an record to client_session_note table
// @Description add to add a single record to client_session_note table in the keycloak database
// @Tags ClientSessionNote
// @Accept  json
// @Produce  json
// @Param ClientSessionNote body model.ClientSessionNote true "Add ClientSessionNote"
// @Success 200 {object} model.ClientSessionNote
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientsessionnote [post]
// echo '{"name": "ujQACRLvWWfMKnESlshrOnhFu","value": "qxpZqhibyLxSGrdIskkYLnEWk","client_session": "IjSJfuaFUGqJdiSQfxqqovWlX"}' | http POST "http://localhost:8080/clientsessionnote" X-Api-User:user123
func AddClientSessionNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	clientsessionnote := &model.ClientSessionNote{}

	if err := readJSON(r, clientsessionnote); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientsessionnote.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientsessionnote.Prepare()

	if err := clientsessionnote.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_session_note", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	clientsessionnote, _, err = dao.AddClientSessionNote(ctx, clientsessionnote)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientsessionnote)
}

// UpdateClientSessionNote Update a single record from client_session_note table in the keycloak database
// @Summary Update an record in table client_session_note
// @Description Update a single record from client_session_note table in the keycloak database
// @Tags ClientSessionNote
// @Accept  json
// @Produce  json
// @Param  argName path string true "name"// @Param  argClientSession path string true "client_session"
// @Param  ClientSessionNote body model.ClientSessionNote true "Update ClientSessionNote record"
// @Success 200 {object} model.ClientSessionNote
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientsessionnote/{argName}/{argClientSession} [put]
// echo '{"name": "ujQACRLvWWfMKnESlshrOnhFu","value": "qxpZqhibyLxSGrdIskkYLnEWk","client_session": "IjSJfuaFUGqJdiSQfxqqovWlX"}' | http PUT "http://localhost:8080/clientsessionnote/hello world/hello world"  X-Api-User:user123
func UpdateClientSessionNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	clientsessionnote := &model.ClientSessionNote{}
	if err := readJSON(r, clientsessionnote); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientsessionnote.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientsessionnote.Prepare()

	if err := clientsessionnote.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_session_note", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientsessionnote, _, err = dao.UpdateClientSessionNote(ctx,
		argName, argClientSession,
		clientsessionnote)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientsessionnote)
}

// DeleteClientSessionNote Delete a single record from client_session_note table in the keycloak database
// @Summary Delete a record from client_session_note
// @Description Delete a single record from client_session_note table in the keycloak database
// @Tags ClientSessionNote
// @Accept  json
// @Produce  json
// @Param  argName path string true "name"// @Param  argClientSession path string true "client_session"
// @Success 204 {object} model.ClientSessionNote
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /clientsessionnote/{argName}/{argClientSession} [delete]
// http DELETE "http://localhost:8080/clientsessionnote/hello world/hello world" X-Api-User:user123
func DeleteClientSessionNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "client_session_note", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteClientSessionNote(ctx, argName, argClientSession)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
