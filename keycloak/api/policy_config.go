package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configPolicyConfigRouter(router *httprouter.Router) {
	router.GET("/policyconfig", GetAllPolicyConfig)
	router.POST("/policyconfig", AddPolicyConfig)
	router.GET("/policyconfig/:argPolicyID/:argName", GetPolicyConfig)
	router.PUT("/policyconfig/:argPolicyID/:argName", UpdatePolicyConfig)
	router.DELETE("/policyconfig/:argPolicyID/:argName", DeletePolicyConfig)
}

func configGinPolicyConfigRouter(router gin.IRoutes) {
	router.GET("/policyconfig", ConverHttprouterToGin(GetAllPolicyConfig))
	router.POST("/policyconfig", ConverHttprouterToGin(AddPolicyConfig))
	router.GET("/policyconfig/:argPolicyID/:argName", ConverHttprouterToGin(GetPolicyConfig))
	router.PUT("/policyconfig/:argPolicyID/:argName", ConverHttprouterToGin(UpdatePolicyConfig))
	router.DELETE("/policyconfig/:argPolicyID/:argName", ConverHttprouterToGin(DeletePolicyConfig))
}

// GetAllPolicyConfig is a function to get a slice of record(s) from policy_config table in the keycloak database
// @Summary Get list of PolicyConfig
// @Tags PolicyConfig
// @Description GetAllPolicyConfig is a handler to get a slice of record(s) from policy_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.PolicyConfig}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /policyconfig [get]
// http "http://localhost:8080/policyconfig?page=0&pagesize=20" X-Api-User:user123
func GetAllPolicyConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "policy_config", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllPolicyConfig(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetPolicyConfig is a function to get a single record from the policy_config table in the keycloak database
// @Summary Get record from table PolicyConfig by  argPolicyID  argName
// @Tags PolicyConfig
// @ID argPolicyID
// @ID argName
// @Description GetPolicyConfig is a function to get a single record from the policy_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argPolicyID path string true "policy_id"
// @Param  argName path string true "name"
// @Success 200 {object} model.PolicyConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /policyconfig/{argPolicyID}/{argName} [get]
// http "http://localhost:8080/policyconfig/hello world/hello world" X-Api-User:user123
func GetPolicyConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argPolicyID, err := parseString(ps, "argPolicyID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "policy_config", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetPolicyConfig(ctx, argPolicyID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddPolicyConfig add to add a single record to policy_config table in the keycloak database
// @Summary Add an record to policy_config table
// @Description add to add a single record to policy_config table in the keycloak database
// @Tags PolicyConfig
// @Accept  json
// @Produce  json
// @Param PolicyConfig body model.PolicyConfig true "Add PolicyConfig"
// @Success 200 {object} model.PolicyConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /policyconfig [post]
// echo '{"policy_id": "mOuVXtorrjiuVBxnpamRSubKU","name": "tgIKcBZZHDpFOUcGIECtYPviu","value": "DZGgIJZCNKaCnukHLxAQpcjNk"}' | http POST "http://localhost:8080/policyconfig" X-Api-User:user123
func AddPolicyConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	policyconfig := &model.PolicyConfig{}

	if err := readJSON(r, policyconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := policyconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	policyconfig.Prepare()

	if err := policyconfig.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "policy_config", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	policyconfig, _, err = dao.AddPolicyConfig(ctx, policyconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, policyconfig)
}

// UpdatePolicyConfig Update a single record from policy_config table in the keycloak database
// @Summary Update an record in table policy_config
// @Description Update a single record from policy_config table in the keycloak database
// @Tags PolicyConfig
// @Accept  json
// @Produce  json
// @Param  argPolicyID path string true "policy_id"// @Param  argName path string true "name"
// @Param  PolicyConfig body model.PolicyConfig true "Update PolicyConfig record"
// @Success 200 {object} model.PolicyConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /policyconfig/{argPolicyID}/{argName} [put]
// echo '{"policy_id": "mOuVXtorrjiuVBxnpamRSubKU","name": "tgIKcBZZHDpFOUcGIECtYPviu","value": "DZGgIJZCNKaCnukHLxAQpcjNk"}' | http PUT "http://localhost:8080/policyconfig/hello world/hello world"  X-Api-User:user123
func UpdatePolicyConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argPolicyID, err := parseString(ps, "argPolicyID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	policyconfig := &model.PolicyConfig{}
	if err := readJSON(r, policyconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := policyconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	policyconfig.Prepare()

	if err := policyconfig.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "policy_config", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	policyconfig, _, err = dao.UpdatePolicyConfig(ctx,
		argPolicyID, argName,
		policyconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, policyconfig)
}

// DeletePolicyConfig Delete a single record from policy_config table in the keycloak database
// @Summary Delete a record from policy_config
// @Description Delete a single record from policy_config table in the keycloak database
// @Tags PolicyConfig
// @Accept  json
// @Produce  json
// @Param  argPolicyID path string true "policy_id"// @Param  argName path string true "name"
// @Success 204 {object} model.PolicyConfig
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /policyconfig/{argPolicyID}/{argName} [delete]
// http DELETE "http://localhost:8080/policyconfig/hello world/hello world" X-Api-User:user123
func DeletePolicyConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argPolicyID, err := parseString(ps, "argPolicyID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "policy_config", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeletePolicyConfig(ctx, argPolicyID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
