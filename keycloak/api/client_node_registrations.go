package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configClientNodeRegistrationsRouter(router *httprouter.Router) {
	router.GET("/clientnoderegistrations", GetAllClientNodeRegistrations)
	router.POST("/clientnoderegistrations", AddClientNodeRegistrations)
	router.GET("/clientnoderegistrations/:argClientID/:argName", GetClientNodeRegistrations)
	router.PUT("/clientnoderegistrations/:argClientID/:argName", UpdateClientNodeRegistrations)
	router.DELETE("/clientnoderegistrations/:argClientID/:argName", DeleteClientNodeRegistrations)
}

func configGinClientNodeRegistrationsRouter(router gin.IRoutes) {
	router.GET("/clientnoderegistrations", ConverHttprouterToGin(GetAllClientNodeRegistrations))
	router.POST("/clientnoderegistrations", ConverHttprouterToGin(AddClientNodeRegistrations))
	router.GET("/clientnoderegistrations/:argClientID/:argName", ConverHttprouterToGin(GetClientNodeRegistrations))
	router.PUT("/clientnoderegistrations/:argClientID/:argName", ConverHttprouterToGin(UpdateClientNodeRegistrations))
	router.DELETE("/clientnoderegistrations/:argClientID/:argName", ConverHttprouterToGin(DeleteClientNodeRegistrations))
}

// GetAllClientNodeRegistrations is a function to get a slice of record(s) from client_node_registrations table in the keycloak database
// @Summary Get list of ClientNodeRegistrations
// @Tags ClientNodeRegistrations
// @Description GetAllClientNodeRegistrations is a handler to get a slice of record(s) from client_node_registrations table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ClientNodeRegistrations}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientnoderegistrations [get]
// http "http://localhost:8080/clientnoderegistrations?page=0&pagesize=20" X-Api-User:user123
func GetAllClientNodeRegistrations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "client_node_registrations", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllClientNodeRegistrations(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetClientNodeRegistrations is a function to get a single record from the client_node_registrations table in the keycloak database
// @Summary Get record from table ClientNodeRegistrations by  argClientID  argName
// @Tags ClientNodeRegistrations
// @ID argClientID
// @ID argName
// @Description GetClientNodeRegistrations is a function to get a single record from the client_node_registrations table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"
// @Param  argName path string true "name"
// @Success 200 {object} model.ClientNodeRegistrations
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /clientnoderegistrations/{argClientID}/{argName} [get]
// http "http://localhost:8080/clientnoderegistrations/hello world/hello world" X-Api-User:user123
func GetClientNodeRegistrations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_node_registrations", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetClientNodeRegistrations(ctx, argClientID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddClientNodeRegistrations add to add a single record to client_node_registrations table in the keycloak database
// @Summary Add an record to client_node_registrations table
// @Description add to add a single record to client_node_registrations table in the keycloak database
// @Tags ClientNodeRegistrations
// @Accept  json
// @Produce  json
// @Param ClientNodeRegistrations body model.ClientNodeRegistrations true "Add ClientNodeRegistrations"
// @Success 200 {object} model.ClientNodeRegistrations
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientnoderegistrations [post]
// echo '{"client_id": "mJHkvfEhooGuQuUmLrhHRcNWU","value": 9,"name": "QuJIcPcsGMYDqKbaYJnBDIJUd"}' | http POST "http://localhost:8080/clientnoderegistrations" X-Api-User:user123
func AddClientNodeRegistrations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	clientnoderegistrations := &model.ClientNodeRegistrations{}

	if err := readJSON(r, clientnoderegistrations); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientnoderegistrations.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientnoderegistrations.Prepare()

	if err := clientnoderegistrations.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_node_registrations", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	clientnoderegistrations, _, err = dao.AddClientNodeRegistrations(ctx, clientnoderegistrations)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientnoderegistrations)
}

// UpdateClientNodeRegistrations Update a single record from client_node_registrations table in the keycloak database
// @Summary Update an record in table client_node_registrations
// @Description Update a single record from client_node_registrations table in the keycloak database
// @Tags ClientNodeRegistrations
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"// @Param  argName path string true "name"
// @Param  ClientNodeRegistrations body model.ClientNodeRegistrations true "Update ClientNodeRegistrations record"
// @Success 200 {object} model.ClientNodeRegistrations
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientnoderegistrations/{argClientID}/{argName} [put]
// echo '{"client_id": "mJHkvfEhooGuQuUmLrhHRcNWU","value": 9,"name": "QuJIcPcsGMYDqKbaYJnBDIJUd"}' | http PUT "http://localhost:8080/clientnoderegistrations/hello world/hello world"  X-Api-User:user123
func UpdateClientNodeRegistrations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientnoderegistrations := &model.ClientNodeRegistrations{}
	if err := readJSON(r, clientnoderegistrations); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientnoderegistrations.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientnoderegistrations.Prepare()

	if err := clientnoderegistrations.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_node_registrations", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientnoderegistrations, _, err = dao.UpdateClientNodeRegistrations(ctx,
		argClientID, argName,
		clientnoderegistrations)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientnoderegistrations)
}

// DeleteClientNodeRegistrations Delete a single record from client_node_registrations table in the keycloak database
// @Summary Delete a record from client_node_registrations
// @Description Delete a single record from client_node_registrations table in the keycloak database
// @Tags ClientNodeRegistrations
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"// @Param  argName path string true "name"
// @Success 204 {object} model.ClientNodeRegistrations
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /clientnoderegistrations/{argClientID}/{argName} [delete]
// http DELETE "http://localhost:8080/clientnoderegistrations/hello world/hello world" X-Api-User:user123
func DeleteClientNodeRegistrations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_node_registrations", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteClientNodeRegistrations(ctx, argClientID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
