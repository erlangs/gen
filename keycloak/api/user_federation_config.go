package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configUserFederationConfigRouter(router *httprouter.Router) {
	router.GET("/userfederationconfig", GetAllUserFederationConfig)
	router.POST("/userfederationconfig", AddUserFederationConfig)
	router.GET("/userfederationconfig/:argUserFederationProviderID/:argName", GetUserFederationConfig)
	router.PUT("/userfederationconfig/:argUserFederationProviderID/:argName", UpdateUserFederationConfig)
	router.DELETE("/userfederationconfig/:argUserFederationProviderID/:argName", DeleteUserFederationConfig)
}

func configGinUserFederationConfigRouter(router gin.IRoutes) {
	router.GET("/userfederationconfig", ConverHttprouterToGin(GetAllUserFederationConfig))
	router.POST("/userfederationconfig", ConverHttprouterToGin(AddUserFederationConfig))
	router.GET("/userfederationconfig/:argUserFederationProviderID/:argName", ConverHttprouterToGin(GetUserFederationConfig))
	router.PUT("/userfederationconfig/:argUserFederationProviderID/:argName", ConverHttprouterToGin(UpdateUserFederationConfig))
	router.DELETE("/userfederationconfig/:argUserFederationProviderID/:argName", ConverHttprouterToGin(DeleteUserFederationConfig))
}

// GetAllUserFederationConfig is a function to get a slice of record(s) from user_federation_config table in the keycloak database
// @Summary Get list of UserFederationConfig
// @Tags UserFederationConfig
// @Description GetAllUserFederationConfig is a handler to get a slice of record(s) from user_federation_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.UserFederationConfig}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userfederationconfig [get]
// http "http://localhost:8080/userfederationconfig?page=0&pagesize=20" X-Api-User:user123
func GetAllUserFederationConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "user_federation_config", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllUserFederationConfig(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetUserFederationConfig is a function to get a single record from the user_federation_config table in the keycloak database
// @Summary Get record from table UserFederationConfig by  argUserFederationProviderID  argName
// @Tags UserFederationConfig
// @ID argUserFederationProviderID
// @ID argName
// @Description GetUserFederationConfig is a function to get a single record from the user_federation_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argUserFederationProviderID path string true "user_federation_provider_id"
// @Param  argName path string true "name"
// @Success 200 {object} model.UserFederationConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /userfederationconfig/{argUserFederationProviderID}/{argName} [get]
// http "http://localhost:8080/userfederationconfig/hello world/hello world" X-Api-User:user123
func GetUserFederationConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserFederationProviderID, err := parseString(ps, "argUserFederationProviderID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_federation_config", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetUserFederationConfig(ctx, argUserFederationProviderID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddUserFederationConfig add to add a single record to user_federation_config table in the keycloak database
// @Summary Add an record to user_federation_config table
// @Description add to add a single record to user_federation_config table in the keycloak database
// @Tags UserFederationConfig
// @Accept  json
// @Produce  json
// @Param UserFederationConfig body model.UserFederationConfig true "Add UserFederationConfig"
// @Success 200 {object} model.UserFederationConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userfederationconfig [post]
// echo '{"user_federation_provider_id": "sUknukExVlChDmLPaiNRhpeFG","value": "fOQoVdWyUkphsOHkXHmkYZmMA","name": "QHWqPWeTHxuUfPLFWSyDKrmeX"}' | http POST "http://localhost:8080/userfederationconfig" X-Api-User:user123
func AddUserFederationConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	userfederationconfig := &model.UserFederationConfig{}

	if err := readJSON(r, userfederationconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userfederationconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userfederationconfig.Prepare()

	if err := userfederationconfig.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_federation_config", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	userfederationconfig, _, err = dao.AddUserFederationConfig(ctx, userfederationconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userfederationconfig)
}

// UpdateUserFederationConfig Update a single record from user_federation_config table in the keycloak database
// @Summary Update an record in table user_federation_config
// @Description Update a single record from user_federation_config table in the keycloak database
// @Tags UserFederationConfig
// @Accept  json
// @Produce  json
// @Param  argUserFederationProviderID path string true "user_federation_provider_id"// @Param  argName path string true "name"
// @Param  UserFederationConfig body model.UserFederationConfig true "Update UserFederationConfig record"
// @Success 200 {object} model.UserFederationConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userfederationconfig/{argUserFederationProviderID}/{argName} [put]
// echo '{"user_federation_provider_id": "sUknukExVlChDmLPaiNRhpeFG","value": "fOQoVdWyUkphsOHkXHmkYZmMA","name": "QHWqPWeTHxuUfPLFWSyDKrmeX"}' | http PUT "http://localhost:8080/userfederationconfig/hello world/hello world"  X-Api-User:user123
func UpdateUserFederationConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserFederationProviderID, err := parseString(ps, "argUserFederationProviderID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	userfederationconfig := &model.UserFederationConfig{}
	if err := readJSON(r, userfederationconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userfederationconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userfederationconfig.Prepare()

	if err := userfederationconfig.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_federation_config", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	userfederationconfig, _, err = dao.UpdateUserFederationConfig(ctx,
		argUserFederationProviderID, argName,
		userfederationconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userfederationconfig)
}

// DeleteUserFederationConfig Delete a single record from user_federation_config table in the keycloak database
// @Summary Delete a record from user_federation_config
// @Description Delete a single record from user_federation_config table in the keycloak database
// @Tags UserFederationConfig
// @Accept  json
// @Produce  json
// @Param  argUserFederationProviderID path string true "user_federation_provider_id"// @Param  argName path string true "name"
// @Success 204 {object} model.UserFederationConfig
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /userfederationconfig/{argUserFederationProviderID}/{argName} [delete]
// http DELETE "http://localhost:8080/userfederationconfig/hello world/hello world" X-Api-User:user123
func DeleteUserFederationConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserFederationProviderID, err := parseString(ps, "argUserFederationProviderID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_federation_config", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteUserFederationConfig(ctx, argUserFederationProviderID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
