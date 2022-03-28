package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configScopeMappingRouter(router *httprouter.Router) {
	router.GET("/scopemapping", GetAllScopeMapping)
	router.POST("/scopemapping", AddScopeMapping)
	router.GET("/scopemapping/:argClientID/:argRoleID", GetScopeMapping)
	router.PUT("/scopemapping/:argClientID/:argRoleID", UpdateScopeMapping)
	router.DELETE("/scopemapping/:argClientID/:argRoleID", DeleteScopeMapping)
}

func configGinScopeMappingRouter(router gin.IRoutes) {
	router.GET("/scopemapping", ConverHttprouterToGin(GetAllScopeMapping))
	router.POST("/scopemapping", ConverHttprouterToGin(AddScopeMapping))
	router.GET("/scopemapping/:argClientID/:argRoleID", ConverHttprouterToGin(GetScopeMapping))
	router.PUT("/scopemapping/:argClientID/:argRoleID", ConverHttprouterToGin(UpdateScopeMapping))
	router.DELETE("/scopemapping/:argClientID/:argRoleID", ConverHttprouterToGin(DeleteScopeMapping))
}

// GetAllScopeMapping is a function to get a slice of record(s) from scope_mapping table in the keycloak database
// @Summary Get list of ScopeMapping
// @Tags ScopeMapping
// @Description GetAllScopeMapping is a handler to get a slice of record(s) from scope_mapping table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ScopeMapping}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /scopemapping [get]
// http "http://localhost:8080/scopemapping?page=0&pagesize=20" X-Api-User:user123
func GetAllScopeMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "scope_mapping", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllScopeMapping(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetScopeMapping is a function to get a single record from the scope_mapping table in the keycloak database
// @Summary Get record from table ScopeMapping by  argClientID  argRoleID
// @Tags ScopeMapping
// @ID argClientID
// @ID argRoleID
// @Description GetScopeMapping is a function to get a single record from the scope_mapping table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"
// @Param  argRoleID path string true "role_id"
// @Success 200 {object} model.ScopeMapping
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /scopemapping/{argClientID}/{argRoleID} [get]
// http "http://localhost:8080/scopemapping/hello world/hello world" X-Api-User:user123
func GetScopeMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argRoleID, err := parseString(ps, "argRoleID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "scope_mapping", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetScopeMapping(ctx, argClientID, argRoleID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddScopeMapping add to add a single record to scope_mapping table in the keycloak database
// @Summary Add an record to scope_mapping table
// @Description add to add a single record to scope_mapping table in the keycloak database
// @Tags ScopeMapping
// @Accept  json
// @Produce  json
// @Param ScopeMapping body model.ScopeMapping true "Add ScopeMapping"
// @Success 200 {object} model.ScopeMapping
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /scopemapping [post]
// echo '{"client_id": "wcLvjJTMbMsEmyoOPjTJXOnWJ","role_id": "HabWxXfiUkCNfyhBEDlHCyRmI"}' | http POST "http://localhost:8080/scopemapping" X-Api-User:user123
func AddScopeMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	scopemapping := &model.ScopeMapping{}

	if err := readJSON(r, scopemapping); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := scopemapping.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	scopemapping.Prepare()

	if err := scopemapping.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "scope_mapping", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	scopemapping, _, err = dao.AddScopeMapping(ctx, scopemapping)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, scopemapping)
}

// UpdateScopeMapping Update a single record from scope_mapping table in the keycloak database
// @Summary Update an record in table scope_mapping
// @Description Update a single record from scope_mapping table in the keycloak database
// @Tags ScopeMapping
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"// @Param  argRoleID path string true "role_id"
// @Param  ScopeMapping body model.ScopeMapping true "Update ScopeMapping record"
// @Success 200 {object} model.ScopeMapping
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /scopemapping/{argClientID}/{argRoleID} [put]
// echo '{"client_id": "wcLvjJTMbMsEmyoOPjTJXOnWJ","role_id": "HabWxXfiUkCNfyhBEDlHCyRmI"}' | http PUT "http://localhost:8080/scopemapping/hello world/hello world"  X-Api-User:user123
func UpdateScopeMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argRoleID, err := parseString(ps, "argRoleID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	scopemapping := &model.ScopeMapping{}
	if err := readJSON(r, scopemapping); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := scopemapping.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	scopemapping.Prepare()

	if err := scopemapping.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "scope_mapping", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	scopemapping, _, err = dao.UpdateScopeMapping(ctx,
		argClientID, argRoleID,
		scopemapping)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, scopemapping)
}

// DeleteScopeMapping Delete a single record from scope_mapping table in the keycloak database
// @Summary Delete a record from scope_mapping
// @Description Delete a single record from scope_mapping table in the keycloak database
// @Tags ScopeMapping
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"// @Param  argRoleID path string true "role_id"
// @Success 204 {object} model.ScopeMapping
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /scopemapping/{argClientID}/{argRoleID} [delete]
// http DELETE "http://localhost:8080/scopemapping/hello world/hello world" X-Api-User:user123
func DeleteScopeMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argRoleID, err := parseString(ps, "argRoleID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "scope_mapping", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteScopeMapping(ctx, argClientID, argRoleID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
