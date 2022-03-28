package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configIdentityProviderConfigRouter(router *httprouter.Router) {
	router.GET("/identityproviderconfig", GetAllIdentityProviderConfig)
	router.POST("/identityproviderconfig", AddIdentityProviderConfig)
	router.GET("/identityproviderconfig/:argIdentityProviderID/:argName", GetIdentityProviderConfig)
	router.PUT("/identityproviderconfig/:argIdentityProviderID/:argName", UpdateIdentityProviderConfig)
	router.DELETE("/identityproviderconfig/:argIdentityProviderID/:argName", DeleteIdentityProviderConfig)
}

func configGinIdentityProviderConfigRouter(router gin.IRoutes) {
	router.GET("/identityproviderconfig", ConverHttprouterToGin(GetAllIdentityProviderConfig))
	router.POST("/identityproviderconfig", ConverHttprouterToGin(AddIdentityProviderConfig))
	router.GET("/identityproviderconfig/:argIdentityProviderID/:argName", ConverHttprouterToGin(GetIdentityProviderConfig))
	router.PUT("/identityproviderconfig/:argIdentityProviderID/:argName", ConverHttprouterToGin(UpdateIdentityProviderConfig))
	router.DELETE("/identityproviderconfig/:argIdentityProviderID/:argName", ConverHttprouterToGin(DeleteIdentityProviderConfig))
}

// GetAllIdentityProviderConfig is a function to get a slice of record(s) from identity_provider_config table in the keycloak database
// @Summary Get list of IdentityProviderConfig
// @Tags IdentityProviderConfig
// @Description GetAllIdentityProviderConfig is a handler to get a slice of record(s) from identity_provider_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.IdentityProviderConfig}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /identityproviderconfig [get]
// http "http://localhost:8080/identityproviderconfig?page=0&pagesize=20" X-Api-User:user123
func GetAllIdentityProviderConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "identity_provider_config", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllIdentityProviderConfig(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetIdentityProviderConfig is a function to get a single record from the identity_provider_config table in the keycloak database
// @Summary Get record from table IdentityProviderConfig by  argIdentityProviderID  argName
// @Tags IdentityProviderConfig
// @ID argIdentityProviderID
// @ID argName
// @Description GetIdentityProviderConfig is a function to get a single record from the identity_provider_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argIdentityProviderID path string true "identity_provider_id"
// @Param  argName path string true "name"
// @Success 200 {object} model.IdentityProviderConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /identityproviderconfig/{argIdentityProviderID}/{argName} [get]
// http "http://localhost:8080/identityproviderconfig/hello world/hello world" X-Api-User:user123
func GetIdentityProviderConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argIdentityProviderID, err := parseString(ps, "argIdentityProviderID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "identity_provider_config", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetIdentityProviderConfig(ctx, argIdentityProviderID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddIdentityProviderConfig add to add a single record to identity_provider_config table in the keycloak database
// @Summary Add an record to identity_provider_config table
// @Description add to add a single record to identity_provider_config table in the keycloak database
// @Tags IdentityProviderConfig
// @Accept  json
// @Produce  json
// @Param IdentityProviderConfig body model.IdentityProviderConfig true "Add IdentityProviderConfig"
// @Success 200 {object} model.IdentityProviderConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /identityproviderconfig [post]
// echo '{"identity_provider_id": "nnaZoAVgGymMayPnobXMcytQI","value": "oNVHTkBLbnBoSNEdbIUAJTCAQ","name": "uQDEadRioPsZdYxySVkZowhSu"}' | http POST "http://localhost:8080/identityproviderconfig" X-Api-User:user123
func AddIdentityProviderConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	identityproviderconfig := &model.IdentityProviderConfig{}

	if err := readJSON(r, identityproviderconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := identityproviderconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	identityproviderconfig.Prepare()

	if err := identityproviderconfig.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "identity_provider_config", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	identityproviderconfig, _, err = dao.AddIdentityProviderConfig(ctx, identityproviderconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, identityproviderconfig)
}

// UpdateIdentityProviderConfig Update a single record from identity_provider_config table in the keycloak database
// @Summary Update an record in table identity_provider_config
// @Description Update a single record from identity_provider_config table in the keycloak database
// @Tags IdentityProviderConfig
// @Accept  json
// @Produce  json
// @Param  argIdentityProviderID path string true "identity_provider_id"// @Param  argName path string true "name"
// @Param  IdentityProviderConfig body model.IdentityProviderConfig true "Update IdentityProviderConfig record"
// @Success 200 {object} model.IdentityProviderConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /identityproviderconfig/{argIdentityProviderID}/{argName} [put]
// echo '{"identity_provider_id": "nnaZoAVgGymMayPnobXMcytQI","value": "oNVHTkBLbnBoSNEdbIUAJTCAQ","name": "uQDEadRioPsZdYxySVkZowhSu"}' | http PUT "http://localhost:8080/identityproviderconfig/hello world/hello world"  X-Api-User:user123
func UpdateIdentityProviderConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argIdentityProviderID, err := parseString(ps, "argIdentityProviderID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	identityproviderconfig := &model.IdentityProviderConfig{}
	if err := readJSON(r, identityproviderconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := identityproviderconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	identityproviderconfig.Prepare()

	if err := identityproviderconfig.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "identity_provider_config", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	identityproviderconfig, _, err = dao.UpdateIdentityProviderConfig(ctx,
		argIdentityProviderID, argName,
		identityproviderconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, identityproviderconfig)
}

// DeleteIdentityProviderConfig Delete a single record from identity_provider_config table in the keycloak database
// @Summary Delete a record from identity_provider_config
// @Description Delete a single record from identity_provider_config table in the keycloak database
// @Tags IdentityProviderConfig
// @Accept  json
// @Produce  json
// @Param  argIdentityProviderID path string true "identity_provider_id"// @Param  argName path string true "name"
// @Success 204 {object} model.IdentityProviderConfig
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /identityproviderconfig/{argIdentityProviderID}/{argName} [delete]
// http DELETE "http://localhost:8080/identityproviderconfig/hello world/hello world" X-Api-User:user123
func DeleteIdentityProviderConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argIdentityProviderID, err := parseString(ps, "argIdentityProviderID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "identity_provider_config", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteIdentityProviderConfig(ctx, argIdentityProviderID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
