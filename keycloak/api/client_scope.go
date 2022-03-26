package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configClientScopeRouter(router *httprouter.Router) {
	router.GET("/clientscope", GetAllClientScope)
	router.POST("/clientscope", AddClientScope)
	router.GET("/clientscope/:argID", GetClientScope)
	router.PUT("/clientscope/:argID", UpdateClientScope)
	router.DELETE("/clientscope/:argID", DeleteClientScope)
}

func configGinClientScopeRouter(router gin.IRoutes) {
	router.GET("/clientscope", ConverHttprouterToGin(GetAllClientScope))
	router.POST("/clientscope", ConverHttprouterToGin(AddClientScope))
	router.GET("/clientscope/:argID", ConverHttprouterToGin(GetClientScope))
	router.PUT("/clientscope/:argID", ConverHttprouterToGin(UpdateClientScope))
	router.DELETE("/clientscope/:argID", ConverHttprouterToGin(DeleteClientScope))
}

// GetAllClientScope is a function to get a slice of record(s) from client_scope table in the keycloak database
// @Summary Get list of ClientScope
// @Tags ClientScope
// @Description GetAllClientScope is a handler to get a slice of record(s) from client_scope table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ClientScope}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientscope [get]
// http "http://localhost:8080/clientscope?page=0&pagesize=20" X-Api-User:user123
func GetAllClientScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "client_scope", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllClientScope(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetClientScope is a function to get a single record from the client_scope table in the keycloak database
// @Summary Get record from table ClientScope by  argID
// @Tags ClientScope
// @ID argID
// @Description GetClientScope is a function to get a single record from the client_scope table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.ClientScope
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /clientscope/{argID} [get]
// http "http://localhost:8080/clientscope/hello world" X-Api-User:user123
func GetClientScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_scope", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetClientScope(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddClientScope add to add a single record to client_scope table in the keycloak database
// @Summary Add an record to client_scope table
// @Description add to add a single record to client_scope table in the keycloak database
// @Tags ClientScope
// @Accept  json
// @Produce  json
// @Param ClientScope body model.ClientScope true "Add ClientScope"
// @Success 200 {object} model.ClientScope
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientscope [post]
// echo '{"id": "JgrAOOoKudgCFpUxcovfjLRSF","name": "dTWbIjOoOVMMPPUvNGonYNYPJ","realm_id": "pcfFqygTRTdUhWtvhnbwjfclB","description": "FCbmNUkkfAscBsHKwbuhEGbCH","protocol": "QnmFNsBTTabPlLOhyexhjBEHr"}' | http POST "http://localhost:8080/clientscope" X-Api-User:user123
func AddClientScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	clientscope := &model.ClientScope{}

	if err := readJSON(r, clientscope); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientscope.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientscope.Prepare()

	if err := clientscope.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_scope", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	clientscope, _, err = dao.AddClientScope(ctx, clientscope)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientscope)
}

// UpdateClientScope Update a single record from client_scope table in the keycloak database
// @Summary Update an record in table client_scope
// @Description Update a single record from client_scope table in the keycloak database
// @Tags ClientScope
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  ClientScope body model.ClientScope true "Update ClientScope record"
// @Success 200 {object} model.ClientScope
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientscope/{argID} [put]
// echo '{"id": "JgrAOOoKudgCFpUxcovfjLRSF","name": "dTWbIjOoOVMMPPUvNGonYNYPJ","realm_id": "pcfFqygTRTdUhWtvhnbwjfclB","description": "FCbmNUkkfAscBsHKwbuhEGbCH","protocol": "QnmFNsBTTabPlLOhyexhjBEHr"}' | http PUT "http://localhost:8080/clientscope/hello world"  X-Api-User:user123
func UpdateClientScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientscope := &model.ClientScope{}
	if err := readJSON(r, clientscope); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientscope.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientscope.Prepare()

	if err := clientscope.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_scope", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientscope, _, err = dao.UpdateClientScope(ctx,
		argID,
		clientscope)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientscope)
}

// DeleteClientScope Delete a single record from client_scope table in the keycloak database
// @Summary Delete a record from client_scope
// @Description Delete a single record from client_scope table in the keycloak database
// @Tags ClientScope
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.ClientScope
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /clientscope/{argID} [delete]
// http DELETE "http://localhost:8080/clientscope/hello world" X-Api-User:user123
func DeleteClientScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_scope", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteClientScope(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
