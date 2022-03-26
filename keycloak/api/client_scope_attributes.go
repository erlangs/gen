package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configClientScopeAttributesRouter(router *httprouter.Router) {
	router.GET("/clientscopeattributes", GetAllClientScopeAttributes)
	router.POST("/clientscopeattributes", AddClientScopeAttributes)
	router.GET("/clientscopeattributes/:argScopeID/:argName", GetClientScopeAttributes)
	router.PUT("/clientscopeattributes/:argScopeID/:argName", UpdateClientScopeAttributes)
	router.DELETE("/clientscopeattributes/:argScopeID/:argName", DeleteClientScopeAttributes)
}

func configGinClientScopeAttributesRouter(router gin.IRoutes) {
	router.GET("/clientscopeattributes", ConverHttprouterToGin(GetAllClientScopeAttributes))
	router.POST("/clientscopeattributes", ConverHttprouterToGin(AddClientScopeAttributes))
	router.GET("/clientscopeattributes/:argScopeID/:argName", ConverHttprouterToGin(GetClientScopeAttributes))
	router.PUT("/clientscopeattributes/:argScopeID/:argName", ConverHttprouterToGin(UpdateClientScopeAttributes))
	router.DELETE("/clientscopeattributes/:argScopeID/:argName", ConverHttprouterToGin(DeleteClientScopeAttributes))
}

// GetAllClientScopeAttributes is a function to get a slice of record(s) from client_scope_attributes table in the keycloak database
// @Summary Get list of ClientScopeAttributes
// @Tags ClientScopeAttributes
// @Description GetAllClientScopeAttributes is a handler to get a slice of record(s) from client_scope_attributes table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ClientScopeAttributes}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientscopeattributes [get]
// http "http://localhost:8080/clientscopeattributes?page=0&pagesize=20" X-Api-User:user123
func GetAllClientScopeAttributes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "client_scope_attributes", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllClientScopeAttributes(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetClientScopeAttributes is a function to get a single record from the client_scope_attributes table in the keycloak database
// @Summary Get record from table ClientScopeAttributes by  argScopeID  argName
// @Tags ClientScopeAttributes
// @ID argScopeID
// @ID argName
// @Description GetClientScopeAttributes is a function to get a single record from the client_scope_attributes table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argScopeID path string true "scope_id"
// @Param  argName path string true "name"
// @Success 200 {object} model.ClientScopeAttributes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /clientscopeattributes/{argScopeID}/{argName} [get]
// http "http://localhost:8080/clientscopeattributes/hello world/hello world" X-Api-User:user123
func GetClientScopeAttributes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_scope_attributes", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetClientScopeAttributes(ctx, argScopeID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddClientScopeAttributes add to add a single record to client_scope_attributes table in the keycloak database
// @Summary Add an record to client_scope_attributes table
// @Description add to add a single record to client_scope_attributes table in the keycloak database
// @Tags ClientScopeAttributes
// @Accept  json
// @Produce  json
// @Param ClientScopeAttributes body model.ClientScopeAttributes true "Add ClientScopeAttributes"
// @Success 200 {object} model.ClientScopeAttributes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientscopeattributes [post]
// echo '{"scope_id": "besGpPNvJtgeaYZAWsMOmHgkN","value": "FOconlfDnswvWLPKwerVVSOBG","name": "FTOyhsNNkyNiRnRPKfepvelaU"}' | http POST "http://localhost:8080/clientscopeattributes" X-Api-User:user123
func AddClientScopeAttributes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	clientscopeattributes := &model.ClientScopeAttributes{}

	if err := readJSON(r, clientscopeattributes); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientscopeattributes.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientscopeattributes.Prepare()

	if err := clientscopeattributes.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_scope_attributes", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	clientscopeattributes, _, err = dao.AddClientScopeAttributes(ctx, clientscopeattributes)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientscopeattributes)
}

// UpdateClientScopeAttributes Update a single record from client_scope_attributes table in the keycloak database
// @Summary Update an record in table client_scope_attributes
// @Description Update a single record from client_scope_attributes table in the keycloak database
// @Tags ClientScopeAttributes
// @Accept  json
// @Produce  json
// @Param  argScopeID path string true "scope_id"// @Param  argName path string true "name"
// @Param  ClientScopeAttributes body model.ClientScopeAttributes true "Update ClientScopeAttributes record"
// @Success 200 {object} model.ClientScopeAttributes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientscopeattributes/{argScopeID}/{argName} [put]
// echo '{"scope_id": "besGpPNvJtgeaYZAWsMOmHgkN","value": "FOconlfDnswvWLPKwerVVSOBG","name": "FTOyhsNNkyNiRnRPKfepvelaU"}' | http PUT "http://localhost:8080/clientscopeattributes/hello world/hello world"  X-Api-User:user123
func UpdateClientScopeAttributes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientscopeattributes := &model.ClientScopeAttributes{}
	if err := readJSON(r, clientscopeattributes); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientscopeattributes.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientscopeattributes.Prepare()

	if err := clientscopeattributes.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_scope_attributes", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientscopeattributes, _, err = dao.UpdateClientScopeAttributes(ctx,
		argScopeID, argName,
		clientscopeattributes)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientscopeattributes)
}

// DeleteClientScopeAttributes Delete a single record from client_scope_attributes table in the keycloak database
// @Summary Delete a record from client_scope_attributes
// @Description Delete a single record from client_scope_attributes table in the keycloak database
// @Tags ClientScopeAttributes
// @Accept  json
// @Produce  json
// @Param  argScopeID path string true "scope_id"// @Param  argName path string true "name"
// @Success 204 {object} model.ClientScopeAttributes
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /clientscopeattributes/{argScopeID}/{argName} [delete]
// http DELETE "http://localhost:8080/clientscopeattributes/hello world/hello world" X-Api-User:user123
func DeleteClientScopeAttributes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_scope_attributes", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteClientScopeAttributes(ctx, argScopeID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
