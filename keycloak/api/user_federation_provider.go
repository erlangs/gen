package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configUserFederationProviderRouter(router *httprouter.Router) {
	router.GET("/userfederationprovider", GetAllUserFederationProvider)
	router.POST("/userfederationprovider", AddUserFederationProvider)
	router.GET("/userfederationprovider/:argID", GetUserFederationProvider)
	router.PUT("/userfederationprovider/:argID", UpdateUserFederationProvider)
	router.DELETE("/userfederationprovider/:argID", DeleteUserFederationProvider)
}

func configGinUserFederationProviderRouter(router gin.IRoutes) {
	router.GET("/userfederationprovider", ConverHttprouterToGin(GetAllUserFederationProvider))
	router.POST("/userfederationprovider", ConverHttprouterToGin(AddUserFederationProvider))
	router.GET("/userfederationprovider/:argID", ConverHttprouterToGin(GetUserFederationProvider))
	router.PUT("/userfederationprovider/:argID", ConverHttprouterToGin(UpdateUserFederationProvider))
	router.DELETE("/userfederationprovider/:argID", ConverHttprouterToGin(DeleteUserFederationProvider))
}

// GetAllUserFederationProvider is a function to get a slice of record(s) from user_federation_provider table in the keycloak database
// @Summary Get list of UserFederationProvider
// @Tags UserFederationProvider
// @Description GetAllUserFederationProvider is a handler to get a slice of record(s) from user_federation_provider table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.UserFederationProvider}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userfederationprovider [get]
// http "http://localhost:8080/userfederationprovider?page=0&pagesize=20" X-Api-User:user123
func GetAllUserFederationProvider(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "user_federation_provider", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllUserFederationProvider(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetUserFederationProvider is a function to get a single record from the user_federation_provider table in the keycloak database
// @Summary Get record from table UserFederationProvider by  argID
// @Tags UserFederationProvider
// @ID argID
// @Description GetUserFederationProvider is a function to get a single record from the user_federation_provider table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.UserFederationProvider
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /userfederationprovider/{argID} [get]
// http "http://localhost:8080/userfederationprovider/hello world" X-Api-User:user123
func GetUserFederationProvider(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_federation_provider", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetUserFederationProvider(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddUserFederationProvider add to add a single record to user_federation_provider table in the keycloak database
// @Summary Add an record to user_federation_provider table
// @Description add to add a single record to user_federation_provider table in the keycloak database
// @Tags UserFederationProvider
// @Accept  json
// @Produce  json
// @Param UserFederationProvider body model.UserFederationProvider true "Add UserFederationProvider"
// @Success 200 {object} model.UserFederationProvider
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userfederationprovider [post]
// echo '{"id": "AbYAXDvRukjuTLlbJmQMllvKX","changed_sync_period": 72,"display_name": "BidAXjtZxIPDCoQceeyfyRHBy","full_sync_period": 17,"last_sync": 75,"priority": 87,"provider_name": "CdTcQQCwkKPLiyOFDpvDmnXNa","realm_id": "EnVbBSyXHeHTRHbkGjTvVbRHd"}' | http POST "http://localhost:8080/userfederationprovider" X-Api-User:user123
func AddUserFederationProvider(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	userfederationprovider := &model.UserFederationProvider{}

	if err := readJSON(r, userfederationprovider); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userfederationprovider.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userfederationprovider.Prepare()

	if err := userfederationprovider.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_federation_provider", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	userfederationprovider, _, err = dao.AddUserFederationProvider(ctx, userfederationprovider)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userfederationprovider)
}

// UpdateUserFederationProvider Update a single record from user_federation_provider table in the keycloak database
// @Summary Update an record in table user_federation_provider
// @Description Update a single record from user_federation_provider table in the keycloak database
// @Tags UserFederationProvider
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  UserFederationProvider body model.UserFederationProvider true "Update UserFederationProvider record"
// @Success 200 {object} model.UserFederationProvider
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userfederationprovider/{argID} [put]
// echo '{"id": "AbYAXDvRukjuTLlbJmQMllvKX","changed_sync_period": 72,"display_name": "BidAXjtZxIPDCoQceeyfyRHBy","full_sync_period": 17,"last_sync": 75,"priority": 87,"provider_name": "CdTcQQCwkKPLiyOFDpvDmnXNa","realm_id": "EnVbBSyXHeHTRHbkGjTvVbRHd"}' | http PUT "http://localhost:8080/userfederationprovider/hello world"  X-Api-User:user123
func UpdateUserFederationProvider(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	userfederationprovider := &model.UserFederationProvider{}
	if err := readJSON(r, userfederationprovider); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userfederationprovider.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userfederationprovider.Prepare()

	if err := userfederationprovider.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_federation_provider", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	userfederationprovider, _, err = dao.UpdateUserFederationProvider(ctx,
		argID,
		userfederationprovider)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userfederationprovider)
}

// DeleteUserFederationProvider Delete a single record from user_federation_provider table in the keycloak database
// @Summary Delete a record from user_federation_provider
// @Description Delete a single record from user_federation_provider table in the keycloak database
// @Tags UserFederationProvider
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.UserFederationProvider
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /userfederationprovider/{argID} [delete]
// http DELETE "http://localhost:8080/userfederationprovider/hello world" X-Api-User:user123
func DeleteUserFederationProvider(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_federation_provider", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteUserFederationProvider(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
