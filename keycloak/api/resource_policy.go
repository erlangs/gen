package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configResourcePolicyRouter(router *httprouter.Router) {
	router.GET("/resourcepolicy", GetAllResourcePolicy)
	router.POST("/resourcepolicy", AddResourcePolicy)
	router.GET("/resourcepolicy/:argResourceID/:argPolicyID", GetResourcePolicy)
	router.PUT("/resourcepolicy/:argResourceID/:argPolicyID", UpdateResourcePolicy)
	router.DELETE("/resourcepolicy/:argResourceID/:argPolicyID", DeleteResourcePolicy)
}

func configGinResourcePolicyRouter(router gin.IRoutes) {
	router.GET("/resourcepolicy", ConverHttprouterToGin(GetAllResourcePolicy))
	router.POST("/resourcepolicy", ConverHttprouterToGin(AddResourcePolicy))
	router.GET("/resourcepolicy/:argResourceID/:argPolicyID", ConverHttprouterToGin(GetResourcePolicy))
	router.PUT("/resourcepolicy/:argResourceID/:argPolicyID", ConverHttprouterToGin(UpdateResourcePolicy))
	router.DELETE("/resourcepolicy/:argResourceID/:argPolicyID", ConverHttprouterToGin(DeleteResourcePolicy))
}

// GetAllResourcePolicy is a function to get a slice of record(s) from resource_policy table in the keycloak database
// @Summary Get list of ResourcePolicy
// @Tags ResourcePolicy
// @Description GetAllResourcePolicy is a handler to get a slice of record(s) from resource_policy table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ResourcePolicy}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourcepolicy [get]
// http "http://localhost:8080/resourcepolicy?page=0&pagesize=20" X-Api-User:user123
func GetAllResourcePolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "resource_policy", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllResourcePolicy(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetResourcePolicy is a function to get a single record from the resource_policy table in the keycloak database
// @Summary Get record from table ResourcePolicy by  argResourceID  argPolicyID
// @Tags ResourcePolicy
// @ID argResourceID
// @ID argPolicyID
// @Description GetResourcePolicy is a function to get a single record from the resource_policy table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argResourceID path string true "resource_id"
// @Param  argPolicyID path string true "policy_id"
// @Success 200 {object} model.ResourcePolicy
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /resourcepolicy/{argResourceID}/{argPolicyID} [get]
// http "http://localhost:8080/resourcepolicy/hello world/hello world" X-Api-User:user123
func GetResourcePolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argResourceID, err := parseString(ps, "argResourceID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argPolicyID, err := parseString(ps, "argPolicyID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_policy", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetResourcePolicy(ctx, argResourceID, argPolicyID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddResourcePolicy add to add a single record to resource_policy table in the keycloak database
// @Summary Add an record to resource_policy table
// @Description add to add a single record to resource_policy table in the keycloak database
// @Tags ResourcePolicy
// @Accept  json
// @Produce  json
// @Param ResourcePolicy body model.ResourcePolicy true "Add ResourcePolicy"
// @Success 200 {object} model.ResourcePolicy
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourcepolicy [post]
// echo '{"resource_id": "GwAXGoJByvpycHOkSTNlBjtTr","policy_id": "QRwgJlTFtOPRYEDeOESpXfcZN"}' | http POST "http://localhost:8080/resourcepolicy" X-Api-User:user123
func AddResourcePolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	resourcepolicy := &model.ResourcePolicy{}

	if err := readJSON(r, resourcepolicy); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resourcepolicy.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resourcepolicy.Prepare()

	if err := resourcepolicy.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_policy", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	resourcepolicy, _, err = dao.AddResourcePolicy(ctx, resourcepolicy)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resourcepolicy)
}

// UpdateResourcePolicy Update a single record from resource_policy table in the keycloak database
// @Summary Update an record in table resource_policy
// @Description Update a single record from resource_policy table in the keycloak database
// @Tags ResourcePolicy
// @Accept  json
// @Produce  json
// @Param  argResourceID path string true "resource_id"// @Param  argPolicyID path string true "policy_id"
// @Param  ResourcePolicy body model.ResourcePolicy true "Update ResourcePolicy record"
// @Success 200 {object} model.ResourcePolicy
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourcepolicy/{argResourceID}/{argPolicyID} [put]
// echo '{"resource_id": "GwAXGoJByvpycHOkSTNlBjtTr","policy_id": "QRwgJlTFtOPRYEDeOESpXfcZN"}' | http PUT "http://localhost:8080/resourcepolicy/hello world/hello world"  X-Api-User:user123
func UpdateResourcePolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argResourceID, err := parseString(ps, "argResourceID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argPolicyID, err := parseString(ps, "argPolicyID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resourcepolicy := &model.ResourcePolicy{}
	if err := readJSON(r, resourcepolicy); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resourcepolicy.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resourcepolicy.Prepare()

	if err := resourcepolicy.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_policy", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resourcepolicy, _, err = dao.UpdateResourcePolicy(ctx,
		argResourceID, argPolicyID,
		resourcepolicy)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resourcepolicy)
}

// DeleteResourcePolicy Delete a single record from resource_policy table in the keycloak database
// @Summary Delete a record from resource_policy
// @Description Delete a single record from resource_policy table in the keycloak database
// @Tags ResourcePolicy
// @Accept  json
// @Produce  json
// @Param  argResourceID path string true "resource_id"// @Param  argPolicyID path string true "policy_id"
// @Success 204 {object} model.ResourcePolicy
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /resourcepolicy/{argResourceID}/{argPolicyID} [delete]
// http DELETE "http://localhost:8080/resourcepolicy/hello world/hello world" X-Api-User:user123
func DeleteResourcePolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argResourceID, err := parseString(ps, "argResourceID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argPolicyID, err := parseString(ps, "argPolicyID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_policy", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteResourcePolicy(ctx, argResourceID, argPolicyID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
