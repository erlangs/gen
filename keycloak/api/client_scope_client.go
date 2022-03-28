package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configClientScopeClientRouter(router *httprouter.Router) {
	router.GET("/clientscopeclient", GetAllClientScopeClient)
	router.POST("/clientscopeclient", AddClientScopeClient)
	router.GET("/clientscopeclient/:argClientID/:argScopeID", GetClientScopeClient)
	router.PUT("/clientscopeclient/:argClientID/:argScopeID", UpdateClientScopeClient)
	router.DELETE("/clientscopeclient/:argClientID/:argScopeID", DeleteClientScopeClient)
}

func configGinClientScopeClientRouter(router gin.IRoutes) {
	router.GET("/clientscopeclient", ConverHttprouterToGin(GetAllClientScopeClient))
	router.POST("/clientscopeclient", ConverHttprouterToGin(AddClientScopeClient))
	router.GET("/clientscopeclient/:argClientID/:argScopeID", ConverHttprouterToGin(GetClientScopeClient))
	router.PUT("/clientscopeclient/:argClientID/:argScopeID", ConverHttprouterToGin(UpdateClientScopeClient))
	router.DELETE("/clientscopeclient/:argClientID/:argScopeID", ConverHttprouterToGin(DeleteClientScopeClient))
}

// GetAllClientScopeClient is a function to get a slice of record(s) from client_scope_client table in the keycloak database
// @Summary Get list of ClientScopeClient
// @Tags ClientScopeClient
// @Description GetAllClientScopeClient is a handler to get a slice of record(s) from client_scope_client table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ClientScopeClient}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientscopeclient [get]
// http "http://localhost:8080/clientscopeclient?page=0&pagesize=20" X-Api-User:user123
func GetAllClientScopeClient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "client_scope_client", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllClientScopeClient(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetClientScopeClient is a function to get a single record from the client_scope_client table in the keycloak database
// @Summary Get record from table ClientScopeClient by  argClientID  argScopeID
// @Tags ClientScopeClient
// @ID argClientID
// @ID argScopeID
// @Description GetClientScopeClient is a function to get a single record from the client_scope_client table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"
// @Param  argScopeID path string true "scope_id"
// @Success 200 {object} model.ClientScopeClient
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /clientscopeclient/{argClientID}/{argScopeID} [get]
// http "http://localhost:8080/clientscopeclient/hello world/hello world" X-Api-User:user123
func GetClientScopeClient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_scope_client", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetClientScopeClient(ctx, argClientID, argScopeID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddClientScopeClient add to add a single record to client_scope_client table in the keycloak database
// @Summary Add an record to client_scope_client table
// @Description add to add a single record to client_scope_client table in the keycloak database
// @Tags ClientScopeClient
// @Accept  json
// @Produce  json
// @Param ClientScopeClient body model.ClientScopeClient true "Add ClientScopeClient"
// @Success 200 {object} model.ClientScopeClient
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientscopeclient [post]
// echo '{"client_id": "vOtIVpcSlvcSLrEAmBWrReWYa","scope_id": "ZquYgPqitcIbsfvSXhtaKnyht","default_scope": true}' | http POST "http://localhost:8080/clientscopeclient" X-Api-User:user123
func AddClientScopeClient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	clientscopeclient := &model.ClientScopeClient{}

	if err := readJSON(r, clientscopeclient); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientscopeclient.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientscopeclient.Prepare()

	if err := clientscopeclient.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_scope_client", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	clientscopeclient, _, err = dao.AddClientScopeClient(ctx, clientscopeclient)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientscopeclient)
}

// UpdateClientScopeClient Update a single record from client_scope_client table in the keycloak database
// @Summary Update an record in table client_scope_client
// @Description Update a single record from client_scope_client table in the keycloak database
// @Tags ClientScopeClient
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"// @Param  argScopeID path string true "scope_id"
// @Param  ClientScopeClient body model.ClientScopeClient true "Update ClientScopeClient record"
// @Success 200 {object} model.ClientScopeClient
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientscopeclient/{argClientID}/{argScopeID} [put]
// echo '{"client_id": "vOtIVpcSlvcSLrEAmBWrReWYa","scope_id": "ZquYgPqitcIbsfvSXhtaKnyht","default_scope": true}' | http PUT "http://localhost:8080/clientscopeclient/hello world/hello world"  X-Api-User:user123
func UpdateClientScopeClient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientscopeclient := &model.ClientScopeClient{}
	if err := readJSON(r, clientscopeclient); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientscopeclient.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientscopeclient.Prepare()

	if err := clientscopeclient.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_scope_client", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientscopeclient, _, err = dao.UpdateClientScopeClient(ctx,
		argClientID, argScopeID,
		clientscopeclient)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientscopeclient)
}

// DeleteClientScopeClient Delete a single record from client_scope_client table in the keycloak database
// @Summary Delete a record from client_scope_client
// @Description Delete a single record from client_scope_client table in the keycloak database
// @Tags ClientScopeClient
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"// @Param  argScopeID path string true "scope_id"
// @Success 204 {object} model.ClientScopeClient
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /clientscopeclient/{argClientID}/{argScopeID} [delete]
// http DELETE "http://localhost:8080/clientscopeclient/hello world/hello world" X-Api-User:user123
func DeleteClientScopeClient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_scope_client", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteClientScopeClient(ctx, argClientID, argScopeID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
