package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configFedUserCredentialRouter(router *httprouter.Router) {
	router.GET("/fedusercredential", GetAllFedUserCredential)
	router.POST("/fedusercredential", AddFedUserCredential)
	router.GET("/fedusercredential/:argID", GetFedUserCredential)
	router.PUT("/fedusercredential/:argID", UpdateFedUserCredential)
	router.DELETE("/fedusercredential/:argID", DeleteFedUserCredential)
}

func configGinFedUserCredentialRouter(router gin.IRoutes) {
	router.GET("/fedusercredential", ConverHttprouterToGin(GetAllFedUserCredential))
	router.POST("/fedusercredential", ConverHttprouterToGin(AddFedUserCredential))
	router.GET("/fedusercredential/:argID", ConverHttprouterToGin(GetFedUserCredential))
	router.PUT("/fedusercredential/:argID", ConverHttprouterToGin(UpdateFedUserCredential))
	router.DELETE("/fedusercredential/:argID", ConverHttprouterToGin(DeleteFedUserCredential))
}

// GetAllFedUserCredential is a function to get a slice of record(s) from fed_user_credential table in the keycloak database
// @Summary Get list of FedUserCredential
// @Tags FedUserCredential
// @Description GetAllFedUserCredential is a handler to get a slice of record(s) from fed_user_credential table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.FedUserCredential}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /fedusercredential [get]
// http "http://localhost:8080/fedusercredential?page=0&pagesize=20" X-Api-User:user123
func GetAllFedUserCredential(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "fed_user_credential", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllFedUserCredential(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetFedUserCredential is a function to get a single record from the fed_user_credential table in the keycloak database
// @Summary Get record from table FedUserCredential by  argID
// @Tags FedUserCredential
// @ID argID
// @Description GetFedUserCredential is a function to get a single record from the fed_user_credential table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.FedUserCredential
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /fedusercredential/{argID} [get]
// http "http://localhost:8080/fedusercredential/hello world" X-Api-User:user123
func GetFedUserCredential(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_credential", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetFedUserCredential(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddFedUserCredential add to add a single record to fed_user_credential table in the keycloak database
// @Summary Add an record to fed_user_credential table
// @Description add to add a single record to fed_user_credential table in the keycloak database
// @Tags FedUserCredential
// @Accept  json
// @Produce  json
// @Param FedUserCredential body model.FedUserCredential true "Add FedUserCredential"
// @Success 200 {object} model.FedUserCredential
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /fedusercredential [post]
// echo '{"id": "FANEcoNcZIMViHrdQDcHnTZUt","salt": "aIIHWOHKVbgVpSoqxWOMlxiWV","type": "wGjYrofvqWqtYnAorFWoqsWIn","created_date": 18,"user_id": "RJgdPQwGRqRvywswNuIIihLrk","realm_id": "OVhQFYMSXKjmMfkdrQXUyZfCY","storage_provider_id": "jxaAFvIwQRqHIYMhCCfIWLbiD","user_label": "jSJAUhchbZKmNogKWUaXWpmMv","secret_data": "xOIWNLvGkTmuFAbIKZfnuVeKU","credential_data": "EXetyFubTdDRhmNTvmnufJxih","priority": 46}' | http POST "http://localhost:8080/fedusercredential" X-Api-User:user123
func AddFedUserCredential(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	fedusercredential := &model.FedUserCredential{}

	if err := readJSON(r, fedusercredential); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := fedusercredential.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	fedusercredential.Prepare()

	if err := fedusercredential.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_credential", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	fedusercredential, _, err = dao.AddFedUserCredential(ctx, fedusercredential)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, fedusercredential)
}

// UpdateFedUserCredential Update a single record from fed_user_credential table in the keycloak database
// @Summary Update an record in table fed_user_credential
// @Description Update a single record from fed_user_credential table in the keycloak database
// @Tags FedUserCredential
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  FedUserCredential body model.FedUserCredential true "Update FedUserCredential record"
// @Success 200 {object} model.FedUserCredential
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /fedusercredential/{argID} [put]
// echo '{"id": "FANEcoNcZIMViHrdQDcHnTZUt","salt": "aIIHWOHKVbgVpSoqxWOMlxiWV","type": "wGjYrofvqWqtYnAorFWoqsWIn","created_date": 18,"user_id": "RJgdPQwGRqRvywswNuIIihLrk","realm_id": "OVhQFYMSXKjmMfkdrQXUyZfCY","storage_provider_id": "jxaAFvIwQRqHIYMhCCfIWLbiD","user_label": "jSJAUhchbZKmNogKWUaXWpmMv","secret_data": "xOIWNLvGkTmuFAbIKZfnuVeKU","credential_data": "EXetyFubTdDRhmNTvmnufJxih","priority": 46}' | http PUT "http://localhost:8080/fedusercredential/hello world"  X-Api-User:user123
func UpdateFedUserCredential(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	fedusercredential := &model.FedUserCredential{}
	if err := readJSON(r, fedusercredential); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := fedusercredential.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	fedusercredential.Prepare()

	if err := fedusercredential.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_credential", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	fedusercredential, _, err = dao.UpdateFedUserCredential(ctx,
		argID,
		fedusercredential)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, fedusercredential)
}

// DeleteFedUserCredential Delete a single record from fed_user_credential table in the keycloak database
// @Summary Delete a record from fed_user_credential
// @Description Delete a single record from fed_user_credential table in the keycloak database
// @Tags FedUserCredential
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.FedUserCredential
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /fedusercredential/{argID} [delete]
// http DELETE "http://localhost:8080/fedusercredential/hello world" X-Api-User:user123
func DeleteFedUserCredential(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_credential", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteFedUserCredential(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
