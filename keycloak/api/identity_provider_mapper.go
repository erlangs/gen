package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configIdentityProviderMapperRouter(router *httprouter.Router) {
	router.GET("/identityprovidermapper", GetAllIdentityProviderMapper)
	router.POST("/identityprovidermapper", AddIdentityProviderMapper)
	router.GET("/identityprovidermapper/:argID", GetIdentityProviderMapper)
	router.PUT("/identityprovidermapper/:argID", UpdateIdentityProviderMapper)
	router.DELETE("/identityprovidermapper/:argID", DeleteIdentityProviderMapper)
}

func configGinIdentityProviderMapperRouter(router gin.IRoutes) {
	router.GET("/identityprovidermapper", ConverHttprouterToGin(GetAllIdentityProviderMapper))
	router.POST("/identityprovidermapper", ConverHttprouterToGin(AddIdentityProviderMapper))
	router.GET("/identityprovidermapper/:argID", ConverHttprouterToGin(GetIdentityProviderMapper))
	router.PUT("/identityprovidermapper/:argID", ConverHttprouterToGin(UpdateIdentityProviderMapper))
	router.DELETE("/identityprovidermapper/:argID", ConverHttprouterToGin(DeleteIdentityProviderMapper))
}

// GetAllIdentityProviderMapper is a function to get a slice of record(s) from identity_provider_mapper table in the keycloak database
// @Summary Get list of IdentityProviderMapper
// @Tags IdentityProviderMapper
// @Description GetAllIdentityProviderMapper is a handler to get a slice of record(s) from identity_provider_mapper table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.IdentityProviderMapper}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /identityprovidermapper [get]
// http "http://localhost:8080/identityprovidermapper?page=0&pagesize=20" X-Api-User:user123
func GetAllIdentityProviderMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "identity_provider_mapper", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllIdentityProviderMapper(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetIdentityProviderMapper is a function to get a single record from the identity_provider_mapper table in the keycloak database
// @Summary Get record from table IdentityProviderMapper by  argID
// @Tags IdentityProviderMapper
// @ID argID
// @Description GetIdentityProviderMapper is a function to get a single record from the identity_provider_mapper table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.IdentityProviderMapper
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /identityprovidermapper/{argID} [get]
// http "http://localhost:8080/identityprovidermapper/hello world" X-Api-User:user123
func GetIdentityProviderMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "identity_provider_mapper", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetIdentityProviderMapper(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddIdentityProviderMapper add to add a single record to identity_provider_mapper table in the keycloak database
// @Summary Add an record to identity_provider_mapper table
// @Description add to add a single record to identity_provider_mapper table in the keycloak database
// @Tags IdentityProviderMapper
// @Accept  json
// @Produce  json
// @Param IdentityProviderMapper body model.IdentityProviderMapper true "Add IdentityProviderMapper"
// @Success 200 {object} model.IdentityProviderMapper
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /identityprovidermapper [post]
// echo '{"id": "ZYvVUOmgbGOfkvNjrPStpGaYk","name": "dFGdxReXQRYGxhyANdQABbsTl","idp_alias": "tArQdTQZbtKavNymBCWnVCixB","idp_mapper_name": "hSpWYptSYaPtApRGigrwuMUbt","realm_id": "VNGOZHsjBpRxjDlyYwgysQXbE"}' | http POST "http://localhost:8080/identityprovidermapper" X-Api-User:user123
func AddIdentityProviderMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	identityprovidermapper := &model.IdentityProviderMapper{}

	if err := readJSON(r, identityprovidermapper); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := identityprovidermapper.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	identityprovidermapper.Prepare()

	if err := identityprovidermapper.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "identity_provider_mapper", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	identityprovidermapper, _, err = dao.AddIdentityProviderMapper(ctx, identityprovidermapper)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, identityprovidermapper)
}

// UpdateIdentityProviderMapper Update a single record from identity_provider_mapper table in the keycloak database
// @Summary Update an record in table identity_provider_mapper
// @Description Update a single record from identity_provider_mapper table in the keycloak database
// @Tags IdentityProviderMapper
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  IdentityProviderMapper body model.IdentityProviderMapper true "Update IdentityProviderMapper record"
// @Success 200 {object} model.IdentityProviderMapper
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /identityprovidermapper/{argID} [put]
// echo '{"id": "ZYvVUOmgbGOfkvNjrPStpGaYk","name": "dFGdxReXQRYGxhyANdQABbsTl","idp_alias": "tArQdTQZbtKavNymBCWnVCixB","idp_mapper_name": "hSpWYptSYaPtApRGigrwuMUbt","realm_id": "VNGOZHsjBpRxjDlyYwgysQXbE"}' | http PUT "http://localhost:8080/identityprovidermapper/hello world"  X-Api-User:user123
func UpdateIdentityProviderMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	identityprovidermapper := &model.IdentityProviderMapper{}
	if err := readJSON(r, identityprovidermapper); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := identityprovidermapper.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	identityprovidermapper.Prepare()

	if err := identityprovidermapper.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "identity_provider_mapper", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	identityprovidermapper, _, err = dao.UpdateIdentityProviderMapper(ctx,
		argID,
		identityprovidermapper)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, identityprovidermapper)
}

// DeleteIdentityProviderMapper Delete a single record from identity_provider_mapper table in the keycloak database
// @Summary Delete a record from identity_provider_mapper
// @Description Delete a single record from identity_provider_mapper table in the keycloak database
// @Tags IdentityProviderMapper
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.IdentityProviderMapper
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /identityprovidermapper/{argID} [delete]
// http DELETE "http://localhost:8080/identityprovidermapper/hello world" X-Api-User:user123
func DeleteIdentityProviderMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "identity_provider_mapper", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteIdentityProviderMapper(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
