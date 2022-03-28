package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configAssociatedPolicyRouter(router *httprouter.Router) {
	router.GET("/associatedpolicy", GetAllAssociatedPolicy)
	router.POST("/associatedpolicy", AddAssociatedPolicy)
	router.GET("/associatedpolicy/:argPolicyID/:argAssociatedPolicyID", GetAssociatedPolicy)
	router.PUT("/associatedpolicy/:argPolicyID/:argAssociatedPolicyID", UpdateAssociatedPolicy)
	router.DELETE("/associatedpolicy/:argPolicyID/:argAssociatedPolicyID", DeleteAssociatedPolicy)
}

func configGinAssociatedPolicyRouter(router gin.IRoutes) {
	router.GET("/associatedpolicy", ConverHttprouterToGin(GetAllAssociatedPolicy))
	router.POST("/associatedpolicy", ConverHttprouterToGin(AddAssociatedPolicy))
	router.GET("/associatedpolicy/:argPolicyID/:argAssociatedPolicyID", ConverHttprouterToGin(GetAssociatedPolicy))
	router.PUT("/associatedpolicy/:argPolicyID/:argAssociatedPolicyID", ConverHttprouterToGin(UpdateAssociatedPolicy))
	router.DELETE("/associatedpolicy/:argPolicyID/:argAssociatedPolicyID", ConverHttprouterToGin(DeleteAssociatedPolicy))
}

// GetAllAssociatedPolicy is a function to get a slice of record(s) from associated_policy table in the keycloak database
// @Summary Get list of AssociatedPolicy
// @Tags AssociatedPolicy
// @Description GetAllAssociatedPolicy is a handler to get a slice of record(s) from associated_policy table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.AssociatedPolicy}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /associatedpolicy [get]
// http "http://localhost:8080/associatedpolicy?page=0&pagesize=20" X-Api-User:user123
func GetAllAssociatedPolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "associated_policy", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllAssociatedPolicy(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetAssociatedPolicy is a function to get a single record from the associated_policy table in the keycloak database
// @Summary Get record from table AssociatedPolicy by  argPolicyID  argAssociatedPolicyID
// @Tags AssociatedPolicy
// @ID argPolicyID
// @ID argAssociatedPolicyID
// @Description GetAssociatedPolicy is a function to get a single record from the associated_policy table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argPolicyID path string true "policy_id"
// @Param  argAssociatedPolicyID path string true "associated_policy_id"
// @Success 200 {object} model.AssociatedPolicy
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /associatedpolicy/{argPolicyID}/{argAssociatedPolicyID} [get]
// http "http://localhost:8080/associatedpolicy/hello world/hello world" X-Api-User:user123
func GetAssociatedPolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argPolicyID, err := parseString(ps, "argPolicyID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argAssociatedPolicyID, err := parseString(ps, "argAssociatedPolicyID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "associated_policy", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetAssociatedPolicy(ctx, argPolicyID, argAssociatedPolicyID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddAssociatedPolicy add to add a single record to associated_policy table in the keycloak database
// @Summary Add an record to associated_policy table
// @Description add to add a single record to associated_policy table in the keycloak database
// @Tags AssociatedPolicy
// @Accept  json
// @Produce  json
// @Param AssociatedPolicy body model.AssociatedPolicy true "Add AssociatedPolicy"
// @Success 200 {object} model.AssociatedPolicy
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /associatedpolicy [post]
// echo '{"policy_id": "wVtWmTHnfrwNhhtKAovYkAnrZ","associated_policy_id": "PwtrxgYUlGUomgHfLwArBWaYm"}' | http POST "http://localhost:8080/associatedpolicy" X-Api-User:user123
func AddAssociatedPolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	associatedpolicy := &model.AssociatedPolicy{}

	if err := readJSON(r, associatedpolicy); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := associatedpolicy.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	associatedpolicy.Prepare()

	if err := associatedpolicy.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "associated_policy", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	associatedpolicy, _, err = dao.AddAssociatedPolicy(ctx, associatedpolicy)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, associatedpolicy)
}

// UpdateAssociatedPolicy Update a single record from associated_policy table in the keycloak database
// @Summary Update an record in table associated_policy
// @Description Update a single record from associated_policy table in the keycloak database
// @Tags AssociatedPolicy
// @Accept  json
// @Produce  json
// @Param  argPolicyID path string true "policy_id"// @Param  argAssociatedPolicyID path string true "associated_policy_id"
// @Param  AssociatedPolicy body model.AssociatedPolicy true "Update AssociatedPolicy record"
// @Success 200 {object} model.AssociatedPolicy
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /associatedpolicy/{argPolicyID}/{argAssociatedPolicyID} [put]
// echo '{"policy_id": "wVtWmTHnfrwNhhtKAovYkAnrZ","associated_policy_id": "PwtrxgYUlGUomgHfLwArBWaYm"}' | http PUT "http://localhost:8080/associatedpolicy/hello world/hello world"  X-Api-User:user123
func UpdateAssociatedPolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argPolicyID, err := parseString(ps, "argPolicyID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argAssociatedPolicyID, err := parseString(ps, "argAssociatedPolicyID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	associatedpolicy := &model.AssociatedPolicy{}
	if err := readJSON(r, associatedpolicy); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := associatedpolicy.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	associatedpolicy.Prepare()

	if err := associatedpolicy.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "associated_policy", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	associatedpolicy, _, err = dao.UpdateAssociatedPolicy(ctx,
		argPolicyID, argAssociatedPolicyID,
		associatedpolicy)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, associatedpolicy)
}

// DeleteAssociatedPolicy Delete a single record from associated_policy table in the keycloak database
// @Summary Delete a record from associated_policy
// @Description Delete a single record from associated_policy table in the keycloak database
// @Tags AssociatedPolicy
// @Accept  json
// @Produce  json
// @Param  argPolicyID path string true "policy_id"// @Param  argAssociatedPolicyID path string true "associated_policy_id"
// @Success 204 {object} model.AssociatedPolicy
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /associatedpolicy/{argPolicyID}/{argAssociatedPolicyID} [delete]
// http DELETE "http://localhost:8080/associatedpolicy/hello world/hello world" X-Api-User:user123
func DeleteAssociatedPolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argPolicyID, err := parseString(ps, "argPolicyID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argAssociatedPolicyID, err := parseString(ps, "argAssociatedPolicyID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "associated_policy", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteAssociatedPolicy(ctx, argPolicyID, argAssociatedPolicyID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
