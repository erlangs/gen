package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configClientAuthFlowBindingsRouter(router *httprouter.Router) {
	router.GET("/clientauthflowbindings", GetAllClientAuthFlowBindings)
	router.POST("/clientauthflowbindings", AddClientAuthFlowBindings)
	router.GET("/clientauthflowbindings/:argClientID/:argBindingName", GetClientAuthFlowBindings)
	router.PUT("/clientauthflowbindings/:argClientID/:argBindingName", UpdateClientAuthFlowBindings)
	router.DELETE("/clientauthflowbindings/:argClientID/:argBindingName", DeleteClientAuthFlowBindings)
}

func configGinClientAuthFlowBindingsRouter(router gin.IRoutes) {
	router.GET("/clientauthflowbindings", ConverHttprouterToGin(GetAllClientAuthFlowBindings))
	router.POST("/clientauthflowbindings", ConverHttprouterToGin(AddClientAuthFlowBindings))
	router.GET("/clientauthflowbindings/:argClientID/:argBindingName", ConverHttprouterToGin(GetClientAuthFlowBindings))
	router.PUT("/clientauthflowbindings/:argClientID/:argBindingName", ConverHttprouterToGin(UpdateClientAuthFlowBindings))
	router.DELETE("/clientauthflowbindings/:argClientID/:argBindingName", ConverHttprouterToGin(DeleteClientAuthFlowBindings))
}

// GetAllClientAuthFlowBindings is a function to get a slice of record(s) from client_auth_flow_bindings table in the keycloak database
// @Summary Get list of ClientAuthFlowBindings
// @Tags ClientAuthFlowBindings
// @Description GetAllClientAuthFlowBindings is a handler to get a slice of record(s) from client_auth_flow_bindings table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ClientAuthFlowBindings}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientauthflowbindings [get]
// http "http://localhost:8080/clientauthflowbindings?page=0&pagesize=20" X-Api-User:user123
func GetAllClientAuthFlowBindings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "client_auth_flow_bindings", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllClientAuthFlowBindings(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetClientAuthFlowBindings is a function to get a single record from the client_auth_flow_bindings table in the keycloak database
// @Summary Get record from table ClientAuthFlowBindings by  argClientID  argBindingName
// @Tags ClientAuthFlowBindings
// @ID argClientID
// @ID argBindingName
// @Description GetClientAuthFlowBindings is a function to get a single record from the client_auth_flow_bindings table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"
// @Param  argBindingName path string true "binding_name"
// @Success 200 {object} model.ClientAuthFlowBindings
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /clientauthflowbindings/{argClientID}/{argBindingName} [get]
// http "http://localhost:8080/clientauthflowbindings/hello world/hello world" X-Api-User:user123
func GetClientAuthFlowBindings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argBindingName, err := parseString(ps, "argBindingName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_auth_flow_bindings", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetClientAuthFlowBindings(ctx, argClientID, argBindingName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddClientAuthFlowBindings add to add a single record to client_auth_flow_bindings table in the keycloak database
// @Summary Add an record to client_auth_flow_bindings table
// @Description add to add a single record to client_auth_flow_bindings table in the keycloak database
// @Tags ClientAuthFlowBindings
// @Accept  json
// @Produce  json
// @Param ClientAuthFlowBindings body model.ClientAuthFlowBindings true "Add ClientAuthFlowBindings"
// @Success 200 {object} model.ClientAuthFlowBindings
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientauthflowbindings [post]
// echo '{"client_id": "KtfNwnLMSuHonDTHkDgGxFQPC","flow_id": "VhEwKYNsToVhgFMbIWJdiJPaH","binding_name": "potSwJnPWISnjBroDOxtceYED"}' | http POST "http://localhost:8080/clientauthflowbindings" X-Api-User:user123
func AddClientAuthFlowBindings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	clientauthflowbindings := &model.ClientAuthFlowBindings{}

	if err := readJSON(r, clientauthflowbindings); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientauthflowbindings.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientauthflowbindings.Prepare()

	if err := clientauthflowbindings.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_auth_flow_bindings", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	clientauthflowbindings, _, err = dao.AddClientAuthFlowBindings(ctx, clientauthflowbindings)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientauthflowbindings)
}

// UpdateClientAuthFlowBindings Update a single record from client_auth_flow_bindings table in the keycloak database
// @Summary Update an record in table client_auth_flow_bindings
// @Description Update a single record from client_auth_flow_bindings table in the keycloak database
// @Tags ClientAuthFlowBindings
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"// @Param  argBindingName path string true "binding_name"
// @Param  ClientAuthFlowBindings body model.ClientAuthFlowBindings true "Update ClientAuthFlowBindings record"
// @Success 200 {object} model.ClientAuthFlowBindings
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientauthflowbindings/{argClientID}/{argBindingName} [put]
// echo '{"client_id": "KtfNwnLMSuHonDTHkDgGxFQPC","flow_id": "VhEwKYNsToVhgFMbIWJdiJPaH","binding_name": "potSwJnPWISnjBroDOxtceYED"}' | http PUT "http://localhost:8080/clientauthflowbindings/hello world/hello world"  X-Api-User:user123
func UpdateClientAuthFlowBindings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argBindingName, err := parseString(ps, "argBindingName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientauthflowbindings := &model.ClientAuthFlowBindings{}
	if err := readJSON(r, clientauthflowbindings); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientauthflowbindings.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientauthflowbindings.Prepare()

	if err := clientauthflowbindings.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_auth_flow_bindings", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientauthflowbindings, _, err = dao.UpdateClientAuthFlowBindings(ctx,
		argClientID, argBindingName,
		clientauthflowbindings)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientauthflowbindings)
}

// DeleteClientAuthFlowBindings Delete a single record from client_auth_flow_bindings table in the keycloak database
// @Summary Delete a record from client_auth_flow_bindings
// @Description Delete a single record from client_auth_flow_bindings table in the keycloak database
// @Tags ClientAuthFlowBindings
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"// @Param  argBindingName path string true "binding_name"
// @Success 204 {object} model.ClientAuthFlowBindings
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /clientauthflowbindings/{argClientID}/{argBindingName} [delete]
// http DELETE "http://localhost:8080/clientauthflowbindings/hello world/hello world" X-Api-User:user123
func DeleteClientAuthFlowBindings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argBindingName, err := parseString(ps, "argBindingName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_auth_flow_bindings", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteClientAuthFlowBindings(ctx, argClientID, argBindingName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
