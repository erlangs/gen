package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configDefaultClientScopeRouter(router *httprouter.Router) {
	router.GET("/defaultclientscope", GetAllDefaultClientScope)
	router.POST("/defaultclientscope", AddDefaultClientScope)
	router.GET("/defaultclientscope/:argRealmID/:argScopeID", GetDefaultClientScope)
	router.PUT("/defaultclientscope/:argRealmID/:argScopeID", UpdateDefaultClientScope)
	router.DELETE("/defaultclientscope/:argRealmID/:argScopeID", DeleteDefaultClientScope)
}

func configGinDefaultClientScopeRouter(router gin.IRoutes) {
	router.GET("/defaultclientscope", ConverHttprouterToGin(GetAllDefaultClientScope))
	router.POST("/defaultclientscope", ConverHttprouterToGin(AddDefaultClientScope))
	router.GET("/defaultclientscope/:argRealmID/:argScopeID", ConverHttprouterToGin(GetDefaultClientScope))
	router.PUT("/defaultclientscope/:argRealmID/:argScopeID", ConverHttprouterToGin(UpdateDefaultClientScope))
	router.DELETE("/defaultclientscope/:argRealmID/:argScopeID", ConverHttprouterToGin(DeleteDefaultClientScope))
}

// GetAllDefaultClientScope is a function to get a slice of record(s) from default_client_scope table in the keycloak database
// @Summary Get list of DefaultClientScope
// @Tags DefaultClientScope
// @Description GetAllDefaultClientScope is a handler to get a slice of record(s) from default_client_scope table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.DefaultClientScope}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /defaultclientscope [get]
// http "http://localhost:8080/defaultclientscope?page=0&pagesize=20" X-Api-User:user123
func GetAllDefaultClientScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "default_client_scope", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllDefaultClientScope(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetDefaultClientScope is a function to get a single record from the default_client_scope table in the keycloak database
// @Summary Get record from table DefaultClientScope by  argRealmID  argScopeID
// @Tags DefaultClientScope
// @ID argRealmID
// @ID argScopeID
// @Description GetDefaultClientScope is a function to get a single record from the default_client_scope table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"
// @Param  argScopeID path string true "scope_id"
// @Success 200 {object} model.DefaultClientScope
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /defaultclientscope/{argRealmID}/{argScopeID} [get]
// http "http://localhost:8080/defaultclientscope/hello world/hello world" X-Api-User:user123
func GetDefaultClientScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "default_client_scope", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetDefaultClientScope(ctx, argRealmID, argScopeID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddDefaultClientScope add to add a single record to default_client_scope table in the keycloak database
// @Summary Add an record to default_client_scope table
// @Description add to add a single record to default_client_scope table in the keycloak database
// @Tags DefaultClientScope
// @Accept  json
// @Produce  json
// @Param DefaultClientScope body model.DefaultClientScope true "Add DefaultClientScope"
// @Success 200 {object} model.DefaultClientScope
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /defaultclientscope [post]
// echo '{"realm_id": "mTdXGBqPOaWFEcrLvIKXGpYnB","scope_id": "UChEmFykmLUsoUeNuvCYstJhW","default_scope": true}' | http POST "http://localhost:8080/defaultclientscope" X-Api-User:user123
func AddDefaultClientScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	defaultclientscope := &model.DefaultClientScope{}

	if err := readJSON(r, defaultclientscope); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := defaultclientscope.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	defaultclientscope.Prepare()

	if err := defaultclientscope.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "default_client_scope", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	defaultclientscope, _, err = dao.AddDefaultClientScope(ctx, defaultclientscope)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, defaultclientscope)
}

// UpdateDefaultClientScope Update a single record from default_client_scope table in the keycloak database
// @Summary Update an record in table default_client_scope
// @Description Update a single record from default_client_scope table in the keycloak database
// @Tags DefaultClientScope
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"// @Param  argScopeID path string true "scope_id"
// @Param  DefaultClientScope body model.DefaultClientScope true "Update DefaultClientScope record"
// @Success 200 {object} model.DefaultClientScope
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /defaultclientscope/{argRealmID}/{argScopeID} [put]
// echo '{"realm_id": "mTdXGBqPOaWFEcrLvIKXGpYnB","scope_id": "UChEmFykmLUsoUeNuvCYstJhW","default_scope": true}' | http PUT "http://localhost:8080/defaultclientscope/hello world/hello world"  X-Api-User:user123
func UpdateDefaultClientScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	defaultclientscope := &model.DefaultClientScope{}
	if err := readJSON(r, defaultclientscope); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := defaultclientscope.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	defaultclientscope.Prepare()

	if err := defaultclientscope.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "default_client_scope", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	defaultclientscope, _, err = dao.UpdateDefaultClientScope(ctx,
		argRealmID, argScopeID,
		defaultclientscope)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, defaultclientscope)
}

// DeleteDefaultClientScope Delete a single record from default_client_scope table in the keycloak database
// @Summary Delete a record from default_client_scope
// @Description Delete a single record from default_client_scope table in the keycloak database
// @Tags DefaultClientScope
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"// @Param  argScopeID path string true "scope_id"
// @Success 204 {object} model.DefaultClientScope
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /defaultclientscope/{argRealmID}/{argScopeID} [delete]
// http DELETE "http://localhost:8080/defaultclientscope/hello world/hello world" X-Api-User:user123
func DeleteDefaultClientScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "default_client_scope", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteDefaultClientScope(ctx, argRealmID, argScopeID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
