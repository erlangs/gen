package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configAuthenticatorConfigRouter(router *httprouter.Router) {
	router.GET("/authenticatorconfig", GetAllAuthenticatorConfig)
	router.POST("/authenticatorconfig", AddAuthenticatorConfig)
	router.GET("/authenticatorconfig/:argID", GetAuthenticatorConfig)
	router.PUT("/authenticatorconfig/:argID", UpdateAuthenticatorConfig)
	router.DELETE("/authenticatorconfig/:argID", DeleteAuthenticatorConfig)
}

func configGinAuthenticatorConfigRouter(router gin.IRoutes) {
	router.GET("/authenticatorconfig", ConverHttprouterToGin(GetAllAuthenticatorConfig))
	router.POST("/authenticatorconfig", ConverHttprouterToGin(AddAuthenticatorConfig))
	router.GET("/authenticatorconfig/:argID", ConverHttprouterToGin(GetAuthenticatorConfig))
	router.PUT("/authenticatorconfig/:argID", ConverHttprouterToGin(UpdateAuthenticatorConfig))
	router.DELETE("/authenticatorconfig/:argID", ConverHttprouterToGin(DeleteAuthenticatorConfig))
}

// GetAllAuthenticatorConfig is a function to get a slice of record(s) from authenticator_config table in the keycloak database
// @Summary Get list of AuthenticatorConfig
// @Tags AuthenticatorConfig
// @Description GetAllAuthenticatorConfig is a handler to get a slice of record(s) from authenticator_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.AuthenticatorConfig}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /authenticatorconfig [get]
// http "http://localhost:8080/authenticatorconfig?page=0&pagesize=20" X-Api-User:user123
func GetAllAuthenticatorConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "authenticator_config", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllAuthenticatorConfig(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetAuthenticatorConfig is a function to get a single record from the authenticator_config table in the keycloak database
// @Summary Get record from table AuthenticatorConfig by  argID
// @Tags AuthenticatorConfig
// @ID argID
// @Description GetAuthenticatorConfig is a function to get a single record from the authenticator_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.AuthenticatorConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /authenticatorconfig/{argID} [get]
// http "http://localhost:8080/authenticatorconfig/hello world" X-Api-User:user123
func GetAuthenticatorConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "authenticator_config", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetAuthenticatorConfig(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddAuthenticatorConfig add to add a single record to authenticator_config table in the keycloak database
// @Summary Add an record to authenticator_config table
// @Description add to add a single record to authenticator_config table in the keycloak database
// @Tags AuthenticatorConfig
// @Accept  json
// @Produce  json
// @Param AuthenticatorConfig body model.AuthenticatorConfig true "Add AuthenticatorConfig"
// @Success 200 {object} model.AuthenticatorConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /authenticatorconfig [post]
// echo '{"id": "GuOQdaKwlZCboeQMHVZSPTrNS","alias": "HjotNcivigqiUbinPxtPFEVvN","realm_id": "PYFeWBermXtJhQHXiLQkmxGiA"}' | http POST "http://localhost:8080/authenticatorconfig" X-Api-User:user123
func AddAuthenticatorConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	authenticatorconfig := &model.AuthenticatorConfig{}

	if err := readJSON(r, authenticatorconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := authenticatorconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	authenticatorconfig.Prepare()

	if err := authenticatorconfig.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "authenticator_config", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	authenticatorconfig, _, err = dao.AddAuthenticatorConfig(ctx, authenticatorconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, authenticatorconfig)
}

// UpdateAuthenticatorConfig Update a single record from authenticator_config table in the keycloak database
// @Summary Update an record in table authenticator_config
// @Description Update a single record from authenticator_config table in the keycloak database
// @Tags AuthenticatorConfig
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  AuthenticatorConfig body model.AuthenticatorConfig true "Update AuthenticatorConfig record"
// @Success 200 {object} model.AuthenticatorConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /authenticatorconfig/{argID} [put]
// echo '{"id": "GuOQdaKwlZCboeQMHVZSPTrNS","alias": "HjotNcivigqiUbinPxtPFEVvN","realm_id": "PYFeWBermXtJhQHXiLQkmxGiA"}' | http PUT "http://localhost:8080/authenticatorconfig/hello world"  X-Api-User:user123
func UpdateAuthenticatorConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	authenticatorconfig := &model.AuthenticatorConfig{}
	if err := readJSON(r, authenticatorconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := authenticatorconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	authenticatorconfig.Prepare()

	if err := authenticatorconfig.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "authenticator_config", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	authenticatorconfig, _, err = dao.UpdateAuthenticatorConfig(ctx,
		argID,
		authenticatorconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, authenticatorconfig)
}

// DeleteAuthenticatorConfig Delete a single record from authenticator_config table in the keycloak database
// @Summary Delete a record from authenticator_config
// @Description Delete a single record from authenticator_config table in the keycloak database
// @Tags AuthenticatorConfig
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.AuthenticatorConfig
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /authenticatorconfig/{argID} [delete]
// http DELETE "http://localhost:8080/authenticatorconfig/hello world" X-Api-User:user123
func DeleteAuthenticatorConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "authenticator_config", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteAuthenticatorConfig(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
