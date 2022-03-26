package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configIdentityProviderRouter(router *httprouter.Router) {
	router.GET("/identityprovider", GetAllIdentityProvider)
	router.POST("/identityprovider", AddIdentityProvider)
	router.GET("/identityprovider/:argInternalID", GetIdentityProvider)
	router.PUT("/identityprovider/:argInternalID", UpdateIdentityProvider)
	router.DELETE("/identityprovider/:argInternalID", DeleteIdentityProvider)
}

func configGinIdentityProviderRouter(router gin.IRoutes) {
	router.GET("/identityprovider", ConverHttprouterToGin(GetAllIdentityProvider))
	router.POST("/identityprovider", ConverHttprouterToGin(AddIdentityProvider))
	router.GET("/identityprovider/:argInternalID", ConverHttprouterToGin(GetIdentityProvider))
	router.PUT("/identityprovider/:argInternalID", ConverHttprouterToGin(UpdateIdentityProvider))
	router.DELETE("/identityprovider/:argInternalID", ConverHttprouterToGin(DeleteIdentityProvider))
}

// GetAllIdentityProvider is a function to get a slice of record(s) from identity_provider table in the keycloak database
// @Summary Get list of IdentityProvider
// @Tags IdentityProvider
// @Description GetAllIdentityProvider is a handler to get a slice of record(s) from identity_provider table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.IdentityProvider}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /identityprovider [get]
// http "http://localhost:8080/identityprovider?page=0&pagesize=20" X-Api-User:user123
func GetAllIdentityProvider(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "identity_provider", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllIdentityProvider(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetIdentityProvider is a function to get a single record from the identity_provider table in the keycloak database
// @Summary Get record from table IdentityProvider by  argInternalID
// @Tags IdentityProvider
// @ID argInternalID
// @Description GetIdentityProvider is a function to get a single record from the identity_provider table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argInternalID path string true "internal_id"
// @Success 200 {object} model.IdentityProvider
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /identityprovider/{argInternalID} [get]
// http "http://localhost:8080/identityprovider/hello world" X-Api-User:user123
func GetIdentityProvider(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argInternalID, err := parseString(ps, "argInternalID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "identity_provider", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetIdentityProvider(ctx, argInternalID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddIdentityProvider add to add a single record to identity_provider table in the keycloak database
// @Summary Add an record to identity_provider table
// @Description add to add a single record to identity_provider table in the keycloak database
// @Tags IdentityProvider
// @Accept  json
// @Produce  json
// @Param IdentityProvider body model.IdentityProvider true "Add IdentityProvider"
// @Success 200 {object} model.IdentityProvider
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /identityprovider [post]
// echo '{"internal_id": "bufABomFBuTnIVLjfCyQeOPaQ","enabled": false,"provider_alias": "GeFfRfltlnXRPPndWMPoMnkCU","provider_id": "lqruxhlrpPiOxrXkqgeinGkoU","store_token": false,"authenticate_by_default": false,"realm_id": "hyOvhMvKsSrbcOxOMNKharPvN","add_token_role": false,"trust_email": true,"first_broker_login_flow_id": "YqiVMXrxhGfkHfoWTnYoXHASh","post_broker_login_flow_id": "ZFGQMmWjfXoElCrWjuyLcbuxE","provider_display_name": "QWrZZnSKiaDCcAjNcbHhobdId","link_only": false}' | http POST "http://localhost:8080/identityprovider" X-Api-User:user123
func AddIdentityProvider(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	identityprovider := &model.IdentityProvider{}

	if err := readJSON(r, identityprovider); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := identityprovider.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	identityprovider.Prepare()

	if err := identityprovider.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "identity_provider", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	identityprovider, _, err = dao.AddIdentityProvider(ctx, identityprovider)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, identityprovider)
}

// UpdateIdentityProvider Update a single record from identity_provider table in the keycloak database
// @Summary Update an record in table identity_provider
// @Description Update a single record from identity_provider table in the keycloak database
// @Tags IdentityProvider
// @Accept  json
// @Produce  json
// @Param  argInternalID path string true "internal_id"
// @Param  IdentityProvider body model.IdentityProvider true "Update IdentityProvider record"
// @Success 200 {object} model.IdentityProvider
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /identityprovider/{argInternalID} [put]
// echo '{"internal_id": "bufABomFBuTnIVLjfCyQeOPaQ","enabled": false,"provider_alias": "GeFfRfltlnXRPPndWMPoMnkCU","provider_id": "lqruxhlrpPiOxrXkqgeinGkoU","store_token": false,"authenticate_by_default": false,"realm_id": "hyOvhMvKsSrbcOxOMNKharPvN","add_token_role": false,"trust_email": true,"first_broker_login_flow_id": "YqiVMXrxhGfkHfoWTnYoXHASh","post_broker_login_flow_id": "ZFGQMmWjfXoElCrWjuyLcbuxE","provider_display_name": "QWrZZnSKiaDCcAjNcbHhobdId","link_only": false}' | http PUT "http://localhost:8080/identityprovider/hello world"  X-Api-User:user123
func UpdateIdentityProvider(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argInternalID, err := parseString(ps, "argInternalID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	identityprovider := &model.IdentityProvider{}
	if err := readJSON(r, identityprovider); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := identityprovider.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	identityprovider.Prepare()

	if err := identityprovider.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "identity_provider", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	identityprovider, _, err = dao.UpdateIdentityProvider(ctx,
		argInternalID,
		identityprovider)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, identityprovider)
}

// DeleteIdentityProvider Delete a single record from identity_provider table in the keycloak database
// @Summary Delete a record from identity_provider
// @Description Delete a single record from identity_provider table in the keycloak database
// @Tags IdentityProvider
// @Accept  json
// @Produce  json
// @Param  argInternalID path string true "internal_id"
// @Success 204 {object} model.IdentityProvider
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /identityprovider/{argInternalID} [delete]
// http DELETE "http://localhost:8080/identityprovider/hello world" X-Api-User:user123
func DeleteIdentityProvider(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argInternalID, err := parseString(ps, "argInternalID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "identity_provider", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteIdentityProvider(ctx, argInternalID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
