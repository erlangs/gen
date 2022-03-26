package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configRealmSMTPConfigRouter(router *httprouter.Router) {
	router.GET("/realmsmtpconfig", GetAllRealmSMTPConfig)
	router.POST("/realmsmtpconfig", AddRealmSMTPConfig)
	router.GET("/realmsmtpconfig/:argRealmID/:argName", GetRealmSMTPConfig)
	router.PUT("/realmsmtpconfig/:argRealmID/:argName", UpdateRealmSMTPConfig)
	router.DELETE("/realmsmtpconfig/:argRealmID/:argName", DeleteRealmSMTPConfig)
}

func configGinRealmSMTPConfigRouter(router gin.IRoutes) {
	router.GET("/realmsmtpconfig", ConverHttprouterToGin(GetAllRealmSMTPConfig))
	router.POST("/realmsmtpconfig", ConverHttprouterToGin(AddRealmSMTPConfig))
	router.GET("/realmsmtpconfig/:argRealmID/:argName", ConverHttprouterToGin(GetRealmSMTPConfig))
	router.PUT("/realmsmtpconfig/:argRealmID/:argName", ConverHttprouterToGin(UpdateRealmSMTPConfig))
	router.DELETE("/realmsmtpconfig/:argRealmID/:argName", ConverHttprouterToGin(DeleteRealmSMTPConfig))
}

// GetAllRealmSMTPConfig is a function to get a slice of record(s) from realm_smtp_config table in the keycloak database
// @Summary Get list of RealmSMTPConfig
// @Tags RealmSMTPConfig
// @Description GetAllRealmSMTPConfig is a handler to get a slice of record(s) from realm_smtp_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.RealmSMTPConfig}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmsmtpconfig [get]
// http "http://localhost:8080/realmsmtpconfig?page=0&pagesize=20" X-Api-User:user123
func GetAllRealmSMTPConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "realm_smtp_config", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllRealmSMTPConfig(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetRealmSMTPConfig is a function to get a single record from the realm_smtp_config table in the keycloak database
// @Summary Get record from table RealmSMTPConfig by  argRealmID  argName
// @Tags RealmSMTPConfig
// @ID argRealmID
// @ID argName
// @Description GetRealmSMTPConfig is a function to get a single record from the realm_smtp_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"
// @Param  argName path string true "name"
// @Success 200 {object} model.RealmSMTPConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /realmsmtpconfig/{argRealmID}/{argName} [get]
// http "http://localhost:8080/realmsmtpconfig/hello world/hello world" X-Api-User:user123
func GetRealmSMTPConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_smtp_config", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetRealmSMTPConfig(ctx, argRealmID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddRealmSMTPConfig add to add a single record to realm_smtp_config table in the keycloak database
// @Summary Add an record to realm_smtp_config table
// @Description add to add a single record to realm_smtp_config table in the keycloak database
// @Tags RealmSMTPConfig
// @Accept  json
// @Produce  json
// @Param RealmSMTPConfig body model.RealmSMTPConfig true "Add RealmSMTPConfig"
// @Success 200 {object} model.RealmSMTPConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmsmtpconfig [post]
// echo '{"realm_id": "WTseRGcelrsUvkDriuEDKWCoC","value": "gJZGAYAjooSURMscSmMwVFIUA","name": "WiYYqHMlKoIJOXKyIXrKxkKoI"}' | http POST "http://localhost:8080/realmsmtpconfig" X-Api-User:user123
func AddRealmSMTPConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	realmsmtpconfig := &model.RealmSMTPConfig{}

	if err := readJSON(r, realmsmtpconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := realmsmtpconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	realmsmtpconfig.Prepare()

	if err := realmsmtpconfig.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_smtp_config", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	realmsmtpconfig, _, err = dao.AddRealmSMTPConfig(ctx, realmsmtpconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, realmsmtpconfig)
}

// UpdateRealmSMTPConfig Update a single record from realm_smtp_config table in the keycloak database
// @Summary Update an record in table realm_smtp_config
// @Description Update a single record from realm_smtp_config table in the keycloak database
// @Tags RealmSMTPConfig
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"// @Param  argName path string true "name"
// @Param  RealmSMTPConfig body model.RealmSMTPConfig true "Update RealmSMTPConfig record"
// @Success 200 {object} model.RealmSMTPConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmsmtpconfig/{argRealmID}/{argName} [put]
// echo '{"realm_id": "WTseRGcelrsUvkDriuEDKWCoC","value": "gJZGAYAjooSURMscSmMwVFIUA","name": "WiYYqHMlKoIJOXKyIXrKxkKoI"}' | http PUT "http://localhost:8080/realmsmtpconfig/hello world/hello world"  X-Api-User:user123
func UpdateRealmSMTPConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	realmsmtpconfig := &model.RealmSMTPConfig{}
	if err := readJSON(r, realmsmtpconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := realmsmtpconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	realmsmtpconfig.Prepare()

	if err := realmsmtpconfig.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_smtp_config", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	realmsmtpconfig, _, err = dao.UpdateRealmSMTPConfig(ctx,
		argRealmID, argName,
		realmsmtpconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, realmsmtpconfig)
}

// DeleteRealmSMTPConfig Delete a single record from realm_smtp_config table in the keycloak database
// @Summary Delete a record from realm_smtp_config
// @Description Delete a single record from realm_smtp_config table in the keycloak database
// @Tags RealmSMTPConfig
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"// @Param  argName path string true "name"
// @Success 204 {object} model.RealmSMTPConfig
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /realmsmtpconfig/{argRealmID}/{argName} [delete]
// http DELETE "http://localhost:8080/realmsmtpconfig/hello world/hello world" X-Api-User:user123
func DeleteRealmSMTPConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_smtp_config", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteRealmSMTPConfig(ctx, argRealmID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
