package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configScopePolicyRouter(router *httprouter.Router) {
	router.GET("/scopepolicy", GetAllScopePolicy)
	router.POST("/scopepolicy", AddScopePolicy)
	router.GET("/scopepolicy/:argScopeID/:argPolicyID", GetScopePolicy)
	router.PUT("/scopepolicy/:argScopeID/:argPolicyID", UpdateScopePolicy)
	router.DELETE("/scopepolicy/:argScopeID/:argPolicyID", DeleteScopePolicy)
}

func configGinScopePolicyRouter(router gin.IRoutes) {
	router.GET("/scopepolicy", ConverHttprouterToGin(GetAllScopePolicy))
	router.POST("/scopepolicy", ConverHttprouterToGin(AddScopePolicy))
	router.GET("/scopepolicy/:argScopeID/:argPolicyID", ConverHttprouterToGin(GetScopePolicy))
	router.PUT("/scopepolicy/:argScopeID/:argPolicyID", ConverHttprouterToGin(UpdateScopePolicy))
	router.DELETE("/scopepolicy/:argScopeID/:argPolicyID", ConverHttprouterToGin(DeleteScopePolicy))
}

// GetAllScopePolicy is a function to get a slice of record(s) from scope_policy table in the keycloak database
// @Summary Get list of ScopePolicy
// @Tags ScopePolicy
// @Description GetAllScopePolicy is a handler to get a slice of record(s) from scope_policy table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ScopePolicy}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /scopepolicy [get]
// http "http://localhost:8080/scopepolicy?page=0&pagesize=20" X-Api-User:user123
func GetAllScopePolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "scope_policy", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllScopePolicy(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetScopePolicy is a function to get a single record from the scope_policy table in the keycloak database
// @Summary Get record from table ScopePolicy by  argScopeID  argPolicyID
// @Tags ScopePolicy
// @ID argScopeID
// @ID argPolicyID
// @Description GetScopePolicy is a function to get a single record from the scope_policy table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argScopeID path string true "scope_id"
// @Param  argPolicyID path string true "policy_id"
// @Success 200 {object} model.ScopePolicy
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /scopepolicy/{argScopeID}/{argPolicyID} [get]
// http "http://localhost:8080/scopepolicy/hello world/hello world" X-Api-User:user123
func GetScopePolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argPolicyID, err := parseString(ps, "argPolicyID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "scope_policy", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetScopePolicy(ctx, argScopeID, argPolicyID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddScopePolicy add to add a single record to scope_policy table in the keycloak database
// @Summary Add an record to scope_policy table
// @Description add to add a single record to scope_policy table in the keycloak database
// @Tags ScopePolicy
// @Accept  json
// @Produce  json
// @Param ScopePolicy body model.ScopePolicy true "Add ScopePolicy"
// @Success 200 {object} model.ScopePolicy
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /scopepolicy [post]
// echo '{"scope_id": "gOKRofCjENFFqnFNNPltITWLX","policy_id": "SQCViBdQAMdpAostFoVCdDxjO"}' | http POST "http://localhost:8080/scopepolicy" X-Api-User:user123
func AddScopePolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	scopepolicy := &model.ScopePolicy{}

	if err := readJSON(r, scopepolicy); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := scopepolicy.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	scopepolicy.Prepare()

	if err := scopepolicy.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "scope_policy", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	scopepolicy, _, err = dao.AddScopePolicy(ctx, scopepolicy)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, scopepolicy)
}

// UpdateScopePolicy Update a single record from scope_policy table in the keycloak database
// @Summary Update an record in table scope_policy
// @Description Update a single record from scope_policy table in the keycloak database
// @Tags ScopePolicy
// @Accept  json
// @Produce  json
// @Param  argScopeID path string true "scope_id"// @Param  argPolicyID path string true "policy_id"
// @Param  ScopePolicy body model.ScopePolicy true "Update ScopePolicy record"
// @Success 200 {object} model.ScopePolicy
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /scopepolicy/{argScopeID}/{argPolicyID} [put]
// echo '{"scope_id": "gOKRofCjENFFqnFNNPltITWLX","policy_id": "SQCViBdQAMdpAostFoVCdDxjO"}' | http PUT "http://localhost:8080/scopepolicy/hello world/hello world"  X-Api-User:user123
func UpdateScopePolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argPolicyID, err := parseString(ps, "argPolicyID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	scopepolicy := &model.ScopePolicy{}
	if err := readJSON(r, scopepolicy); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := scopepolicy.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	scopepolicy.Prepare()

	if err := scopepolicy.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "scope_policy", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	scopepolicy, _, err = dao.UpdateScopePolicy(ctx,
		argScopeID, argPolicyID,
		scopepolicy)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, scopepolicy)
}

// DeleteScopePolicy Delete a single record from scope_policy table in the keycloak database
// @Summary Delete a record from scope_policy
// @Description Delete a single record from scope_policy table in the keycloak database
// @Tags ScopePolicy
// @Accept  json
// @Produce  json
// @Param  argScopeID path string true "scope_id"// @Param  argPolicyID path string true "policy_id"
// @Success 204 {object} model.ScopePolicy
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /scopepolicy/{argScopeID}/{argPolicyID} [delete]
// http DELETE "http://localhost:8080/scopepolicy/hello world/hello world" X-Api-User:user123
func DeleteScopePolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argPolicyID, err := parseString(ps, "argPolicyID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "scope_policy", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteScopePolicy(ctx, argScopeID, argPolicyID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
