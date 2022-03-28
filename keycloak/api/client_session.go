package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configClientSessionRouter(router *httprouter.Router) {
	router.GET("/clientsession", GetAllClientSession)
	router.POST("/clientsession", AddClientSession)
	router.GET("/clientsession/:argID", GetClientSession)
	router.PUT("/clientsession/:argID", UpdateClientSession)
	router.DELETE("/clientsession/:argID", DeleteClientSession)
}

func configGinClientSessionRouter(router gin.IRoutes) {
	router.GET("/clientsession", ConverHttprouterToGin(GetAllClientSession))
	router.POST("/clientsession", ConverHttprouterToGin(AddClientSession))
	router.GET("/clientsession/:argID", ConverHttprouterToGin(GetClientSession))
	router.PUT("/clientsession/:argID", ConverHttprouterToGin(UpdateClientSession))
	router.DELETE("/clientsession/:argID", ConverHttprouterToGin(DeleteClientSession))
}

// GetAllClientSession is a function to get a slice of record(s) from client_session table in the keycloak database
// @Summary Get list of ClientSession
// @Tags ClientSession
// @Description GetAllClientSession is a handler to get a slice of record(s) from client_session table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ClientSession}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientsession [get]
// http "http://localhost:8080/clientsession?page=0&pagesize=20" X-Api-User:user123
func GetAllClientSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "client_session", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllClientSession(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetClientSession is a function to get a single record from the client_session table in the keycloak database
// @Summary Get record from table ClientSession by  argID
// @Tags ClientSession
// @ID argID
// @Description GetClientSession is a function to get a single record from the client_session table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.ClientSession
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /clientsession/{argID} [get]
// http "http://localhost:8080/clientsession/hello world" X-Api-User:user123
func GetClientSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_session", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetClientSession(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddClientSession add to add a single record to client_session table in the keycloak database
// @Summary Add an record to client_session table
// @Description add to add a single record to client_session table in the keycloak database
// @Tags ClientSession
// @Accept  json
// @Produce  json
// @Param ClientSession body model.ClientSession true "Add ClientSession"
// @Success 200 {object} model.ClientSession
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientsession [post]
// echo '{"id": "aUPyCHbcQNyqoUhBomVogEFnu","client_id": "KipPBZHggtiKEVynnhAdmAxrc","redirect_uri": "VoGFgjWAEiyYIBHskewmmDsal","state": "hJSMXsTUaHZNsGyWQvwlLtYQJ","timestamp": 75,"session_id": "ahrBsNKPNYQNxgYRDxwjLrVxd","auth_method": "eAYgUKJxJNbqLjyJAKErQsVvb","realm_id": "XEBTrxoIanpkEsuSQCrAeCfBu","auth_user_id": "xadmvSkbXUbwqBEswxvrVhqpP","current_action": "VfPYMuoavgBBHiypnhhFbRcdY"}' | http POST "http://localhost:8080/clientsession" X-Api-User:user123
func AddClientSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	clientsession := &model.ClientSession{}

	if err := readJSON(r, clientsession); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientsession.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientsession.Prepare()

	if err := clientsession.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_session", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	clientsession, _, err = dao.AddClientSession(ctx, clientsession)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientsession)
}

// UpdateClientSession Update a single record from client_session table in the keycloak database
// @Summary Update an record in table client_session
// @Description Update a single record from client_session table in the keycloak database
// @Tags ClientSession
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  ClientSession body model.ClientSession true "Update ClientSession record"
// @Success 200 {object} model.ClientSession
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientsession/{argID} [put]
// echo '{"id": "aUPyCHbcQNyqoUhBomVogEFnu","client_id": "KipPBZHggtiKEVynnhAdmAxrc","redirect_uri": "VoGFgjWAEiyYIBHskewmmDsal","state": "hJSMXsTUaHZNsGyWQvwlLtYQJ","timestamp": 75,"session_id": "ahrBsNKPNYQNxgYRDxwjLrVxd","auth_method": "eAYgUKJxJNbqLjyJAKErQsVvb","realm_id": "XEBTrxoIanpkEsuSQCrAeCfBu","auth_user_id": "xadmvSkbXUbwqBEswxvrVhqpP","current_action": "VfPYMuoavgBBHiypnhhFbRcdY"}' | http PUT "http://localhost:8080/clientsession/hello world"  X-Api-User:user123
func UpdateClientSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientsession := &model.ClientSession{}
	if err := readJSON(r, clientsession); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientsession.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientsession.Prepare()

	if err := clientsession.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_session", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientsession, _, err = dao.UpdateClientSession(ctx,
		argID,
		clientsession)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientsession)
}

// DeleteClientSession Delete a single record from client_session table in the keycloak database
// @Summary Delete a record from client_session
// @Description Delete a single record from client_session table in the keycloak database
// @Tags ClientSession
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.ClientSession
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /clientsession/{argID} [delete]
// http DELETE "http://localhost:8080/clientsession/hello world" X-Api-User:user123
func DeleteClientSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_session", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteClientSession(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
