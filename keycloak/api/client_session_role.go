package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configClientSessionRoleRouter(router *httprouter.Router) {
	router.GET("/clientsessionrole", GetAllClientSessionRole)
	router.POST("/clientsessionrole", AddClientSessionRole)
	router.GET("/clientsessionrole/:argRoleID/:argClientSession", GetClientSessionRole)
	router.PUT("/clientsessionrole/:argRoleID/:argClientSession", UpdateClientSessionRole)
	router.DELETE("/clientsessionrole/:argRoleID/:argClientSession", DeleteClientSessionRole)
}

func configGinClientSessionRoleRouter(router gin.IRoutes) {
	router.GET("/clientsessionrole", ConverHttprouterToGin(GetAllClientSessionRole))
	router.POST("/clientsessionrole", ConverHttprouterToGin(AddClientSessionRole))
	router.GET("/clientsessionrole/:argRoleID/:argClientSession", ConverHttprouterToGin(GetClientSessionRole))
	router.PUT("/clientsessionrole/:argRoleID/:argClientSession", ConverHttprouterToGin(UpdateClientSessionRole))
	router.DELETE("/clientsessionrole/:argRoleID/:argClientSession", ConverHttprouterToGin(DeleteClientSessionRole))
}

// GetAllClientSessionRole is a function to get a slice of record(s) from client_session_role table in the keycloak database
// @Summary Get list of ClientSessionRole
// @Tags ClientSessionRole
// @Description GetAllClientSessionRole is a handler to get a slice of record(s) from client_session_role table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ClientSessionRole}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientsessionrole [get]
// http "http://localhost:8080/clientsessionrole?page=0&pagesize=20" X-Api-User:user123
func GetAllClientSessionRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "client_session_role", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllClientSessionRole(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetClientSessionRole is a function to get a single record from the client_session_role table in the keycloak database
// @Summary Get record from table ClientSessionRole by  argRoleID  argClientSession
// @Tags ClientSessionRole
// @ID argRoleID
// @ID argClientSession
// @Description GetClientSessionRole is a function to get a single record from the client_session_role table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argRoleID path string true "role_id"
// @Param  argClientSession path string true "client_session"
// @Success 200 {object} model.ClientSessionRole
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /clientsessionrole/{argRoleID}/{argClientSession} [get]
// http "http://localhost:8080/clientsessionrole/hello world/hello world" X-Api-User:user123
func GetClientSessionRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRoleID, err := parseString(ps, "argRoleID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argClientSession, err := parseString(ps, "argClientSession")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_session_role", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetClientSessionRole(ctx, argRoleID, argClientSession)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddClientSessionRole add to add a single record to client_session_role table in the keycloak database
// @Summary Add an record to client_session_role table
// @Description add to add a single record to client_session_role table in the keycloak database
// @Tags ClientSessionRole
// @Accept  json
// @Produce  json
// @Param ClientSessionRole body model.ClientSessionRole true "Add ClientSessionRole"
// @Success 200 {object} model.ClientSessionRole
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientsessionrole [post]
// echo '{"role_id": "RKWIQEKubvRlaBIOByTEWWMeJ","client_session": "jrkibwgrXDVGDKxROWAGwGGWw"}' | http POST "http://localhost:8080/clientsessionrole" X-Api-User:user123
func AddClientSessionRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	clientsessionrole := &model.ClientSessionRole{}

	if err := readJSON(r, clientsessionrole); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientsessionrole.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientsessionrole.Prepare()

	if err := clientsessionrole.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_session_role", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	clientsessionrole, _, err = dao.AddClientSessionRole(ctx, clientsessionrole)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientsessionrole)
}

// UpdateClientSessionRole Update a single record from client_session_role table in the keycloak database
// @Summary Update an record in table client_session_role
// @Description Update a single record from client_session_role table in the keycloak database
// @Tags ClientSessionRole
// @Accept  json
// @Produce  json
// @Param  argRoleID path string true "role_id"// @Param  argClientSession path string true "client_session"
// @Param  ClientSessionRole body model.ClientSessionRole true "Update ClientSessionRole record"
// @Success 200 {object} model.ClientSessionRole
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientsessionrole/{argRoleID}/{argClientSession} [put]
// echo '{"role_id": "RKWIQEKubvRlaBIOByTEWWMeJ","client_session": "jrkibwgrXDVGDKxROWAGwGGWw"}' | http PUT "http://localhost:8080/clientsessionrole/hello world/hello world"  X-Api-User:user123
func UpdateClientSessionRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRoleID, err := parseString(ps, "argRoleID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argClientSession, err := parseString(ps, "argClientSession")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientsessionrole := &model.ClientSessionRole{}
	if err := readJSON(r, clientsessionrole); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientsessionrole.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientsessionrole.Prepare()

	if err := clientsessionrole.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_session_role", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientsessionrole, _, err = dao.UpdateClientSessionRole(ctx,
		argRoleID, argClientSession,
		clientsessionrole)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientsessionrole)
}

// DeleteClientSessionRole Delete a single record from client_session_role table in the keycloak database
// @Summary Delete a record from client_session_role
// @Description Delete a single record from client_session_role table in the keycloak database
// @Tags ClientSessionRole
// @Accept  json
// @Produce  json
// @Param  argRoleID path string true "role_id"// @Param  argClientSession path string true "client_session"
// @Success 204 {object} model.ClientSessionRole
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /clientsessionrole/{argRoleID}/{argClientSession} [delete]
// http DELETE "http://localhost:8080/clientsessionrole/hello world/hello world" X-Api-User:user123
func DeleteClientSessionRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRoleID, err := parseString(ps, "argRoleID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argClientSession, err := parseString(ps, "argClientSession")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_session_role", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteClientSessionRole(ctx, argRoleID, argClientSession)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
