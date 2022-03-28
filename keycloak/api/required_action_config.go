package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configRequiredActionConfigRouter(router *httprouter.Router) {
	router.GET("/requiredactionconfig", GetAllRequiredActionConfig)
	router.POST("/requiredactionconfig", AddRequiredActionConfig)
	router.GET("/requiredactionconfig/:argRequiredActionID/:argName", GetRequiredActionConfig)
	router.PUT("/requiredactionconfig/:argRequiredActionID/:argName", UpdateRequiredActionConfig)
	router.DELETE("/requiredactionconfig/:argRequiredActionID/:argName", DeleteRequiredActionConfig)
}

func configGinRequiredActionConfigRouter(router gin.IRoutes) {
	router.GET("/requiredactionconfig", ConverHttprouterToGin(GetAllRequiredActionConfig))
	router.POST("/requiredactionconfig", ConverHttprouterToGin(AddRequiredActionConfig))
	router.GET("/requiredactionconfig/:argRequiredActionID/:argName", ConverHttprouterToGin(GetRequiredActionConfig))
	router.PUT("/requiredactionconfig/:argRequiredActionID/:argName", ConverHttprouterToGin(UpdateRequiredActionConfig))
	router.DELETE("/requiredactionconfig/:argRequiredActionID/:argName", ConverHttprouterToGin(DeleteRequiredActionConfig))
}

// GetAllRequiredActionConfig is a function to get a slice of record(s) from required_action_config table in the keycloak database
// @Summary Get list of RequiredActionConfig
// @Tags RequiredActionConfig
// @Description GetAllRequiredActionConfig is a handler to get a slice of record(s) from required_action_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.RequiredActionConfig}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /requiredactionconfig [get]
// http "http://localhost:8080/requiredactionconfig?page=0&pagesize=20" X-Api-User:user123
func GetAllRequiredActionConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "required_action_config", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllRequiredActionConfig(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetRequiredActionConfig is a function to get a single record from the required_action_config table in the keycloak database
// @Summary Get record from table RequiredActionConfig by  argRequiredActionID  argName
// @Tags RequiredActionConfig
// @ID argRequiredActionID
// @ID argName
// @Description GetRequiredActionConfig is a function to get a single record from the required_action_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argRequiredActionID path string true "required_action_id"
// @Param  argName path string true "name"
// @Success 200 {object} model.RequiredActionConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /requiredactionconfig/{argRequiredActionID}/{argName} [get]
// http "http://localhost:8080/requiredactionconfig/hello world/hello world" X-Api-User:user123
func GetRequiredActionConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRequiredActionID, err := parseString(ps, "argRequiredActionID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "required_action_config", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetRequiredActionConfig(ctx, argRequiredActionID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddRequiredActionConfig add to add a single record to required_action_config table in the keycloak database
// @Summary Add an record to required_action_config table
// @Description add to add a single record to required_action_config table in the keycloak database
// @Tags RequiredActionConfig
// @Accept  json
// @Produce  json
// @Param RequiredActionConfig body model.RequiredActionConfig true "Add RequiredActionConfig"
// @Success 200 {object} model.RequiredActionConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /requiredactionconfig [post]
// echo '{"required_action_id": "iDhjmTZtdoliEeaZBODsMmJym","value": "tneAoposWKktZcRvQijnxYrET","name": "KnFujbbuMUxemQWxwgrthlYyP"}' | http POST "http://localhost:8080/requiredactionconfig" X-Api-User:user123
func AddRequiredActionConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	requiredactionconfig := &model.RequiredActionConfig{}

	if err := readJSON(r, requiredactionconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := requiredactionconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	requiredactionconfig.Prepare()

	if err := requiredactionconfig.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "required_action_config", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	requiredactionconfig, _, err = dao.AddRequiredActionConfig(ctx, requiredactionconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, requiredactionconfig)
}

// UpdateRequiredActionConfig Update a single record from required_action_config table in the keycloak database
// @Summary Update an record in table required_action_config
// @Description Update a single record from required_action_config table in the keycloak database
// @Tags RequiredActionConfig
// @Accept  json
// @Produce  json
// @Param  argRequiredActionID path string true "required_action_id"// @Param  argName path string true "name"
// @Param  RequiredActionConfig body model.RequiredActionConfig true "Update RequiredActionConfig record"
// @Success 200 {object} model.RequiredActionConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /requiredactionconfig/{argRequiredActionID}/{argName} [put]
// echo '{"required_action_id": "iDhjmTZtdoliEeaZBODsMmJym","value": "tneAoposWKktZcRvQijnxYrET","name": "KnFujbbuMUxemQWxwgrthlYyP"}' | http PUT "http://localhost:8080/requiredactionconfig/hello world/hello world"  X-Api-User:user123
func UpdateRequiredActionConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRequiredActionID, err := parseString(ps, "argRequiredActionID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	requiredactionconfig := &model.RequiredActionConfig{}
	if err := readJSON(r, requiredactionconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := requiredactionconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	requiredactionconfig.Prepare()

	if err := requiredactionconfig.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "required_action_config", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	requiredactionconfig, _, err = dao.UpdateRequiredActionConfig(ctx,
		argRequiredActionID, argName,
		requiredactionconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, requiredactionconfig)
}

// DeleteRequiredActionConfig Delete a single record from required_action_config table in the keycloak database
// @Summary Delete a record from required_action_config
// @Description Delete a single record from required_action_config table in the keycloak database
// @Tags RequiredActionConfig
// @Accept  json
// @Produce  json
// @Param  argRequiredActionID path string true "required_action_id"// @Param  argName path string true "name"
// @Success 204 {object} model.RequiredActionConfig
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /requiredactionconfig/{argRequiredActionID}/{argName} [delete]
// http DELETE "http://localhost:8080/requiredactionconfig/hello world/hello world" X-Api-User:user123
func DeleteRequiredActionConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRequiredActionID, err := parseString(ps, "argRequiredActionID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "required_action_config", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteRequiredActionConfig(ctx, argRequiredActionID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
