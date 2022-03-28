package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configRealmRequiredCredentialRouter(router *httprouter.Router) {
	router.GET("/realmrequiredcredential", GetAllRealmRequiredCredential)
	router.POST("/realmrequiredcredential", AddRealmRequiredCredential)
	router.GET("/realmrequiredcredential/:argType/:argRealmID", GetRealmRequiredCredential)
	router.PUT("/realmrequiredcredential/:argType/:argRealmID", UpdateRealmRequiredCredential)
	router.DELETE("/realmrequiredcredential/:argType/:argRealmID", DeleteRealmRequiredCredential)
}

func configGinRealmRequiredCredentialRouter(router gin.IRoutes) {
	router.GET("/realmrequiredcredential", ConverHttprouterToGin(GetAllRealmRequiredCredential))
	router.POST("/realmrequiredcredential", ConverHttprouterToGin(AddRealmRequiredCredential))
	router.GET("/realmrequiredcredential/:argType/:argRealmID", ConverHttprouterToGin(GetRealmRequiredCredential))
	router.PUT("/realmrequiredcredential/:argType/:argRealmID", ConverHttprouterToGin(UpdateRealmRequiredCredential))
	router.DELETE("/realmrequiredcredential/:argType/:argRealmID", ConverHttprouterToGin(DeleteRealmRequiredCredential))
}

// GetAllRealmRequiredCredential is a function to get a slice of record(s) from realm_required_credential table in the keycloak database
// @Summary Get list of RealmRequiredCredential
// @Tags RealmRequiredCredential
// @Description GetAllRealmRequiredCredential is a handler to get a slice of record(s) from realm_required_credential table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.RealmRequiredCredential}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmrequiredcredential [get]
// http "http://localhost:8080/realmrequiredcredential?page=0&pagesize=20" X-Api-User:user123
func GetAllRealmRequiredCredential(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "realm_required_credential", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllRealmRequiredCredential(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetRealmRequiredCredential is a function to get a single record from the realm_required_credential table in the keycloak database
// @Summary Get record from table RealmRequiredCredential by  argType  argRealmID
// @Tags RealmRequiredCredential
// @ID argType
// @ID argRealmID
// @Description GetRealmRequiredCredential is a function to get a single record from the realm_required_credential table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argType path string true "type"
// @Param  argRealmID path string true "realm_id"
// @Success 200 {object} model.RealmRequiredCredential
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /realmrequiredcredential/{argType}/{argRealmID} [get]
// http "http://localhost:8080/realmrequiredcredential/hello world/hello world" X-Api-User:user123
func GetRealmRequiredCredential(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argType, err := parseString(ps, "argType")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_required_credential", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetRealmRequiredCredential(ctx, argType, argRealmID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddRealmRequiredCredential add to add a single record to realm_required_credential table in the keycloak database
// @Summary Add an record to realm_required_credential table
// @Description add to add a single record to realm_required_credential table in the keycloak database
// @Tags RealmRequiredCredential
// @Accept  json
// @Produce  json
// @Param RealmRequiredCredential body model.RealmRequiredCredential true "Add RealmRequiredCredential"
// @Success 200 {object} model.RealmRequiredCredential
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmrequiredcredential [post]
// echo '{"type": "PTJIXPSCyiKLsjBcuyYqWgvno","form_label": "UmRVLAtSbaVPHYKZlNuJoiLjd","input": false,"secret": true,"realm_id": "YccILRTUcGvkoBRWUWHxiaYlm"}' | http POST "http://localhost:8080/realmrequiredcredential" X-Api-User:user123
func AddRealmRequiredCredential(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	realmrequiredcredential := &model.RealmRequiredCredential{}

	if err := readJSON(r, realmrequiredcredential); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := realmrequiredcredential.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	realmrequiredcredential.Prepare()

	if err := realmrequiredcredential.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_required_credential", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	realmrequiredcredential, _, err = dao.AddRealmRequiredCredential(ctx, realmrequiredcredential)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, realmrequiredcredential)
}

// UpdateRealmRequiredCredential Update a single record from realm_required_credential table in the keycloak database
// @Summary Update an record in table realm_required_credential
// @Description Update a single record from realm_required_credential table in the keycloak database
// @Tags RealmRequiredCredential
// @Accept  json
// @Produce  json
// @Param  argType path string true "type"// @Param  argRealmID path string true "realm_id"
// @Param  RealmRequiredCredential body model.RealmRequiredCredential true "Update RealmRequiredCredential record"
// @Success 200 {object} model.RealmRequiredCredential
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmrequiredcredential/{argType}/{argRealmID} [put]
// echo '{"type": "PTJIXPSCyiKLsjBcuyYqWgvno","form_label": "UmRVLAtSbaVPHYKZlNuJoiLjd","input": false,"secret": true,"realm_id": "YccILRTUcGvkoBRWUWHxiaYlm"}' | http PUT "http://localhost:8080/realmrequiredcredential/hello world/hello world"  X-Api-User:user123
func UpdateRealmRequiredCredential(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argType, err := parseString(ps, "argType")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	realmrequiredcredential := &model.RealmRequiredCredential{}
	if err := readJSON(r, realmrequiredcredential); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := realmrequiredcredential.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	realmrequiredcredential.Prepare()

	if err := realmrequiredcredential.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_required_credential", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	realmrequiredcredential, _, err = dao.UpdateRealmRequiredCredential(ctx,
		argType, argRealmID,
		realmrequiredcredential)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, realmrequiredcredential)
}

// DeleteRealmRequiredCredential Delete a single record from realm_required_credential table in the keycloak database
// @Summary Delete a record from realm_required_credential
// @Description Delete a single record from realm_required_credential table in the keycloak database
// @Tags RealmRequiredCredential
// @Accept  json
// @Produce  json
// @Param  argType path string true "type"// @Param  argRealmID path string true "realm_id"
// @Success 204 {object} model.RealmRequiredCredential
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /realmrequiredcredential/{argType}/{argRealmID} [delete]
// http DELETE "http://localhost:8080/realmrequiredcredential/hello world/hello world" X-Api-User:user123
func DeleteRealmRequiredCredential(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argType, err := parseString(ps, "argType")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_required_credential", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteRealmRequiredCredential(ctx, argType, argRealmID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
