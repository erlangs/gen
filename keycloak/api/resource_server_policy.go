package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configResourceServerPolicyRouter(router *httprouter.Router) {
	router.GET("/resourceserverpolicy", GetAllResourceServerPolicy)
	router.POST("/resourceserverpolicy", AddResourceServerPolicy)
	router.GET("/resourceserverpolicy/:argID", GetResourceServerPolicy)
	router.PUT("/resourceserverpolicy/:argID", UpdateResourceServerPolicy)
	router.DELETE("/resourceserverpolicy/:argID", DeleteResourceServerPolicy)
}

func configGinResourceServerPolicyRouter(router gin.IRoutes) {
	router.GET("/resourceserverpolicy", ConverHttprouterToGin(GetAllResourceServerPolicy))
	router.POST("/resourceserverpolicy", ConverHttprouterToGin(AddResourceServerPolicy))
	router.GET("/resourceserverpolicy/:argID", ConverHttprouterToGin(GetResourceServerPolicy))
	router.PUT("/resourceserverpolicy/:argID", ConverHttprouterToGin(UpdateResourceServerPolicy))
	router.DELETE("/resourceserverpolicy/:argID", ConverHttprouterToGin(DeleteResourceServerPolicy))
}

// GetAllResourceServerPolicy is a function to get a slice of record(s) from resource_server_policy table in the keycloak database
// @Summary Get list of ResourceServerPolicy
// @Tags ResourceServerPolicy
// @Description GetAllResourceServerPolicy is a handler to get a slice of record(s) from resource_server_policy table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ResourceServerPolicy}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceserverpolicy [get]
// http "http://localhost:8080/resourceserverpolicy?page=0&pagesize=20" X-Api-User:user123
func GetAllResourceServerPolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "resource_server_policy", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllResourceServerPolicy(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetResourceServerPolicy is a function to get a single record from the resource_server_policy table in the keycloak database
// @Summary Get record from table ResourceServerPolicy by  argID
// @Tags ResourceServerPolicy
// @ID argID
// @Description GetResourceServerPolicy is a function to get a single record from the resource_server_policy table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.ResourceServerPolicy
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /resourceserverpolicy/{argID} [get]
// http "http://localhost:8080/resourceserverpolicy/hello world" X-Api-User:user123
func GetResourceServerPolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server_policy", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetResourceServerPolicy(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddResourceServerPolicy add to add a single record to resource_server_policy table in the keycloak database
// @Summary Add an record to resource_server_policy table
// @Description add to add a single record to resource_server_policy table in the keycloak database
// @Tags ResourceServerPolicy
// @Accept  json
// @Produce  json
// @Param ResourceServerPolicy body model.ResourceServerPolicy true "Add ResourceServerPolicy"
// @Success 200 {object} model.ResourceServerPolicy
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceserverpolicy [post]
// echo '{"id": "kygdDtwOPGRZRfkpuJWccuTPy","name": "RhdPTmdCrycnolrYsjsIZlEma","description": "wdUrFfxXTOWmxEUOUnIJdKnGa","type": "UioqtLpQHTvFTRZxTZtCVECHm","decision_strategy": "xTFrbwjrwiHdUAllaiayEghnT","logic": "XdfUCTnMEDIAegqgDttnjgmKy","resource_server_id": "tQreAOQvONEroAJGJljpEqmRw","owner": "TcyhbLvqmVkWBtXyEuTvFPvee"}' | http POST "http://localhost:8080/resourceserverpolicy" X-Api-User:user123
func AddResourceServerPolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	resourceserverpolicy := &model.ResourceServerPolicy{}

	if err := readJSON(r, resourceserverpolicy); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resourceserverpolicy.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resourceserverpolicy.Prepare()

	if err := resourceserverpolicy.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server_policy", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	resourceserverpolicy, _, err = dao.AddResourceServerPolicy(ctx, resourceserverpolicy)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resourceserverpolicy)
}

// UpdateResourceServerPolicy Update a single record from resource_server_policy table in the keycloak database
// @Summary Update an record in table resource_server_policy
// @Description Update a single record from resource_server_policy table in the keycloak database
// @Tags ResourceServerPolicy
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  ResourceServerPolicy body model.ResourceServerPolicy true "Update ResourceServerPolicy record"
// @Success 200 {object} model.ResourceServerPolicy
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceserverpolicy/{argID} [put]
// echo '{"id": "kygdDtwOPGRZRfkpuJWccuTPy","name": "RhdPTmdCrycnolrYsjsIZlEma","description": "wdUrFfxXTOWmxEUOUnIJdKnGa","type": "UioqtLpQHTvFTRZxTZtCVECHm","decision_strategy": "xTFrbwjrwiHdUAllaiayEghnT","logic": "XdfUCTnMEDIAegqgDttnjgmKy","resource_server_id": "tQreAOQvONEroAJGJljpEqmRw","owner": "TcyhbLvqmVkWBtXyEuTvFPvee"}' | http PUT "http://localhost:8080/resourceserverpolicy/hello world"  X-Api-User:user123
func UpdateResourceServerPolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resourceserverpolicy := &model.ResourceServerPolicy{}
	if err := readJSON(r, resourceserverpolicy); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resourceserverpolicy.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resourceserverpolicy.Prepare()

	if err := resourceserverpolicy.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server_policy", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resourceserverpolicy, _, err = dao.UpdateResourceServerPolicy(ctx,
		argID,
		resourceserverpolicy)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resourceserverpolicy)
}

// DeleteResourceServerPolicy Delete a single record from resource_server_policy table in the keycloak database
// @Summary Delete a record from resource_server_policy
// @Description Delete a single record from resource_server_policy table in the keycloak database
// @Tags ResourceServerPolicy
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.ResourceServerPolicy
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /resourceserverpolicy/{argID} [delete]
// http DELETE "http://localhost:8080/resourceserverpolicy/hello world" X-Api-User:user123
func DeleteResourceServerPolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server_policy", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteResourceServerPolicy(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
