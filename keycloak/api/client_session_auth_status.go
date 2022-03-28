package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configClientSessionAuthStatusRouter(router *httprouter.Router) {
	router.GET("/clientsessionauthstatus", GetAllClientSessionAuthStatus)
	router.POST("/clientsessionauthstatus", AddClientSessionAuthStatus)
	router.GET("/clientsessionauthstatus/:argAuthenticator/:argClientSession", GetClientSessionAuthStatus)
	router.PUT("/clientsessionauthstatus/:argAuthenticator/:argClientSession", UpdateClientSessionAuthStatus)
	router.DELETE("/clientsessionauthstatus/:argAuthenticator/:argClientSession", DeleteClientSessionAuthStatus)
}

func configGinClientSessionAuthStatusRouter(router gin.IRoutes) {
	router.GET("/clientsessionauthstatus", ConverHttprouterToGin(GetAllClientSessionAuthStatus))
	router.POST("/clientsessionauthstatus", ConverHttprouterToGin(AddClientSessionAuthStatus))
	router.GET("/clientsessionauthstatus/:argAuthenticator/:argClientSession", ConverHttprouterToGin(GetClientSessionAuthStatus))
	router.PUT("/clientsessionauthstatus/:argAuthenticator/:argClientSession", ConverHttprouterToGin(UpdateClientSessionAuthStatus))
	router.DELETE("/clientsessionauthstatus/:argAuthenticator/:argClientSession", ConverHttprouterToGin(DeleteClientSessionAuthStatus))
}

// GetAllClientSessionAuthStatus is a function to get a slice of record(s) from client_session_auth_status table in the keycloak database
// @Summary Get list of ClientSessionAuthStatus
// @Tags ClientSessionAuthStatus
// @Description GetAllClientSessionAuthStatus is a handler to get a slice of record(s) from client_session_auth_status table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ClientSessionAuthStatus}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientsessionauthstatus [get]
// http "http://localhost:8080/clientsessionauthstatus?page=0&pagesize=20" X-Api-User:user123
func GetAllClientSessionAuthStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "client_session_auth_status", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllClientSessionAuthStatus(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetClientSessionAuthStatus is a function to get a single record from the client_session_auth_status table in the keycloak database
// @Summary Get record from table ClientSessionAuthStatus by  argAuthenticator  argClientSession
// @Tags ClientSessionAuthStatus
// @ID argAuthenticator
// @ID argClientSession
// @Description GetClientSessionAuthStatus is a function to get a single record from the client_session_auth_status table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argAuthenticator path string true "authenticator"
// @Param  argClientSession path string true "client_session"
// @Success 200 {object} model.ClientSessionAuthStatus
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /clientsessionauthstatus/{argAuthenticator}/{argClientSession} [get]
// http "http://localhost:8080/clientsessionauthstatus/hello world/hello world" X-Api-User:user123
func GetClientSessionAuthStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argAuthenticator, err := parseString(ps, "argAuthenticator")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argClientSession, err := parseString(ps, "argClientSession")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_session_auth_status", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetClientSessionAuthStatus(ctx, argAuthenticator, argClientSession)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddClientSessionAuthStatus add to add a single record to client_session_auth_status table in the keycloak database
// @Summary Add an record to client_session_auth_status table
// @Description add to add a single record to client_session_auth_status table in the keycloak database
// @Tags ClientSessionAuthStatus
// @Accept  json
// @Produce  json
// @Param ClientSessionAuthStatus body model.ClientSessionAuthStatus true "Add ClientSessionAuthStatus"
// @Success 200 {object} model.ClientSessionAuthStatus
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientsessionauthstatus [post]
// echo '{"authenticator": "irheVBRjHBDgIaNhIBWcZKayV","status": 45,"client_session": "VbJrPFjUwaVxhyRkKKFENRBdm"}' | http POST "http://localhost:8080/clientsessionauthstatus" X-Api-User:user123
func AddClientSessionAuthStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	clientsessionauthstatus := &model.ClientSessionAuthStatus{}

	if err := readJSON(r, clientsessionauthstatus); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientsessionauthstatus.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientsessionauthstatus.Prepare()

	if err := clientsessionauthstatus.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_session_auth_status", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	clientsessionauthstatus, _, err = dao.AddClientSessionAuthStatus(ctx, clientsessionauthstatus)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientsessionauthstatus)
}

// UpdateClientSessionAuthStatus Update a single record from client_session_auth_status table in the keycloak database
// @Summary Update an record in table client_session_auth_status
// @Description Update a single record from client_session_auth_status table in the keycloak database
// @Tags ClientSessionAuthStatus
// @Accept  json
// @Produce  json
// @Param  argAuthenticator path string true "authenticator"// @Param  argClientSession path string true "client_session"
// @Param  ClientSessionAuthStatus body model.ClientSessionAuthStatus true "Update ClientSessionAuthStatus record"
// @Success 200 {object} model.ClientSessionAuthStatus
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientsessionauthstatus/{argAuthenticator}/{argClientSession} [put]
// echo '{"authenticator": "irheVBRjHBDgIaNhIBWcZKayV","status": 45,"client_session": "VbJrPFjUwaVxhyRkKKFENRBdm"}' | http PUT "http://localhost:8080/clientsessionauthstatus/hello world/hello world"  X-Api-User:user123
func UpdateClientSessionAuthStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argAuthenticator, err := parseString(ps, "argAuthenticator")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argClientSession, err := parseString(ps, "argClientSession")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientsessionauthstatus := &model.ClientSessionAuthStatus{}
	if err := readJSON(r, clientsessionauthstatus); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientsessionauthstatus.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientsessionauthstatus.Prepare()

	if err := clientsessionauthstatus.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_session_auth_status", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientsessionauthstatus, _, err = dao.UpdateClientSessionAuthStatus(ctx,
		argAuthenticator, argClientSession,
		clientsessionauthstatus)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientsessionauthstatus)
}

// DeleteClientSessionAuthStatus Delete a single record from client_session_auth_status table in the keycloak database
// @Summary Delete a record from client_session_auth_status
// @Description Delete a single record from client_session_auth_status table in the keycloak database
// @Tags ClientSessionAuthStatus
// @Accept  json
// @Produce  json
// @Param  argAuthenticator path string true "authenticator"// @Param  argClientSession path string true "client_session"
// @Success 204 {object} model.ClientSessionAuthStatus
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /clientsessionauthstatus/{argAuthenticator}/{argClientSession} [delete]
// http DELETE "http://localhost:8080/clientsessionauthstatus/hello world/hello world" X-Api-User:user123
func DeleteClientSessionAuthStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argAuthenticator, err := parseString(ps, "argAuthenticator")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argClientSession, err := parseString(ps, "argClientSession")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_session_auth_status", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteClientSessionAuthStatus(ctx, argAuthenticator, argClientSession)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
