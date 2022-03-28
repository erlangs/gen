package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configComponentConfigRouter(router *httprouter.Router) {
	router.GET("/componentconfig", GetAllComponentConfig)
	router.POST("/componentconfig", AddComponentConfig)
	router.GET("/componentconfig/:argID", GetComponentConfig)
	router.PUT("/componentconfig/:argID", UpdateComponentConfig)
	router.DELETE("/componentconfig/:argID", DeleteComponentConfig)
}

func configGinComponentConfigRouter(router gin.IRoutes) {
	router.GET("/componentconfig", ConverHttprouterToGin(GetAllComponentConfig))
	router.POST("/componentconfig", ConverHttprouterToGin(AddComponentConfig))
	router.GET("/componentconfig/:argID", ConverHttprouterToGin(GetComponentConfig))
	router.PUT("/componentconfig/:argID", ConverHttprouterToGin(UpdateComponentConfig))
	router.DELETE("/componentconfig/:argID", ConverHttprouterToGin(DeleteComponentConfig))
}

// GetAllComponentConfig is a function to get a slice of record(s) from component_config table in the keycloak database
// @Summary Get list of ComponentConfig
// @Tags ComponentConfig
// @Description GetAllComponentConfig is a handler to get a slice of record(s) from component_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ComponentConfig}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /componentconfig [get]
// http "http://localhost:8080/componentconfig?page=0&pagesize=20" X-Api-User:user123
func GetAllComponentConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "component_config", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllComponentConfig(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetComponentConfig is a function to get a single record from the component_config table in the keycloak database
// @Summary Get record from table ComponentConfig by  argID
// @Tags ComponentConfig
// @ID argID
// @Description GetComponentConfig is a function to get a single record from the component_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.ComponentConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /componentconfig/{argID} [get]
// http "http://localhost:8080/componentconfig/hello world" X-Api-User:user123
func GetComponentConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "component_config", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetComponentConfig(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddComponentConfig add to add a single record to component_config table in the keycloak database
// @Summary Add an record to component_config table
// @Description add to add a single record to component_config table in the keycloak database
// @Tags ComponentConfig
// @Accept  json
// @Produce  json
// @Param ComponentConfig body model.ComponentConfig true "Add ComponentConfig"
// @Success 200 {object} model.ComponentConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /componentconfig [post]
// echo '{"id": "KVeKvUqcnVlrlGGYfEnoVgFRD","component_id": "JNTUCEphapWEiPaCTmWpEUjQY","name": "UlBKqMHfGpLyKSfQVgToOPwix","value": "aNYTCyTRAZhUSMnUtlpMACNeV"}' | http POST "http://localhost:8080/componentconfig" X-Api-User:user123
func AddComponentConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	componentconfig := &model.ComponentConfig{}

	if err := readJSON(r, componentconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := componentconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	componentconfig.Prepare()

	if err := componentconfig.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "component_config", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	componentconfig, _, err = dao.AddComponentConfig(ctx, componentconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, componentconfig)
}

// UpdateComponentConfig Update a single record from component_config table in the keycloak database
// @Summary Update an record in table component_config
// @Description Update a single record from component_config table in the keycloak database
// @Tags ComponentConfig
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  ComponentConfig body model.ComponentConfig true "Update ComponentConfig record"
// @Success 200 {object} model.ComponentConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /componentconfig/{argID} [put]
// echo '{"id": "KVeKvUqcnVlrlGGYfEnoVgFRD","component_id": "JNTUCEphapWEiPaCTmWpEUjQY","name": "UlBKqMHfGpLyKSfQVgToOPwix","value": "aNYTCyTRAZhUSMnUtlpMACNeV"}' | http PUT "http://localhost:8080/componentconfig/hello world"  X-Api-User:user123
func UpdateComponentConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	componentconfig := &model.ComponentConfig{}
	if err := readJSON(r, componentconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := componentconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	componentconfig.Prepare()

	if err := componentconfig.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "component_config", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	componentconfig, _, err = dao.UpdateComponentConfig(ctx,
		argID,
		componentconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, componentconfig)
}

// DeleteComponentConfig Delete a single record from component_config table in the keycloak database
// @Summary Delete a record from component_config
// @Description Delete a single record from component_config table in the keycloak database
// @Tags ComponentConfig
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.ComponentConfig
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /componentconfig/{argID} [delete]
// http DELETE "http://localhost:8080/componentconfig/hello world" X-Api-User:user123
func DeleteComponentConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "component_config", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteComponentConfig(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
