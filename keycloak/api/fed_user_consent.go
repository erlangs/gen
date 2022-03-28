package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configFedUserConsentRouter(router *httprouter.Router) {
	router.GET("/feduserconsent", GetAllFedUserConsent)
	router.POST("/feduserconsent", AddFedUserConsent)
	router.GET("/feduserconsent/:argID", GetFedUserConsent)
	router.PUT("/feduserconsent/:argID", UpdateFedUserConsent)
	router.DELETE("/feduserconsent/:argID", DeleteFedUserConsent)
}

func configGinFedUserConsentRouter(router gin.IRoutes) {
	router.GET("/feduserconsent", ConverHttprouterToGin(GetAllFedUserConsent))
	router.POST("/feduserconsent", ConverHttprouterToGin(AddFedUserConsent))
	router.GET("/feduserconsent/:argID", ConverHttprouterToGin(GetFedUserConsent))
	router.PUT("/feduserconsent/:argID", ConverHttprouterToGin(UpdateFedUserConsent))
	router.DELETE("/feduserconsent/:argID", ConverHttprouterToGin(DeleteFedUserConsent))
}

// GetAllFedUserConsent is a function to get a slice of record(s) from fed_user_consent table in the keycloak database
// @Summary Get list of FedUserConsent
// @Tags FedUserConsent
// @Description GetAllFedUserConsent is a handler to get a slice of record(s) from fed_user_consent table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.FedUserConsent}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /feduserconsent [get]
// http "http://localhost:8080/feduserconsent?page=0&pagesize=20" X-Api-User:user123
func GetAllFedUserConsent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "fed_user_consent", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllFedUserConsent(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetFedUserConsent is a function to get a single record from the fed_user_consent table in the keycloak database
// @Summary Get record from table FedUserConsent by  argID
// @Tags FedUserConsent
// @ID argID
// @Description GetFedUserConsent is a function to get a single record from the fed_user_consent table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.FedUserConsent
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /feduserconsent/{argID} [get]
// http "http://localhost:8080/feduserconsent/hello world" X-Api-User:user123
func GetFedUserConsent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_consent", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetFedUserConsent(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddFedUserConsent add to add a single record to fed_user_consent table in the keycloak database
// @Summary Add an record to fed_user_consent table
// @Description add to add a single record to fed_user_consent table in the keycloak database
// @Tags FedUserConsent
// @Accept  json
// @Produce  json
// @Param FedUserConsent body model.FedUserConsent true "Add FedUserConsent"
// @Success 200 {object} model.FedUserConsent
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /feduserconsent [post]
// echo '{"id": "pFMhaLmPNxCOrVOscovXvEmjJ","client_id": "jkYoprGPrslAsUGQGnNZopuLh","user_id": "YJhpLBAVwbqZbsZSWICYarPyC","realm_id": "TYapqgRAEUGZtXTvFSZeWBHAD","storage_provider_id": "eivbwbQaRHGnjTPCfDKngnBGk","created_date": 88,"last_updated_date": 44,"client_storage_provider": "SkwmExEyAnyAJiVuBFtOHRwnY","external_client_id": "bTEtTBiBlkmCgkHIWgNwWQmxn"}' | http POST "http://localhost:8080/feduserconsent" X-Api-User:user123
func AddFedUserConsent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	feduserconsent := &model.FedUserConsent{}

	if err := readJSON(r, feduserconsent); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := feduserconsent.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	feduserconsent.Prepare()

	if err := feduserconsent.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_consent", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	feduserconsent, _, err = dao.AddFedUserConsent(ctx, feduserconsent)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, feduserconsent)
}

// UpdateFedUserConsent Update a single record from fed_user_consent table in the keycloak database
// @Summary Update an record in table fed_user_consent
// @Description Update a single record from fed_user_consent table in the keycloak database
// @Tags FedUserConsent
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  FedUserConsent body model.FedUserConsent true "Update FedUserConsent record"
// @Success 200 {object} model.FedUserConsent
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /feduserconsent/{argID} [put]
// echo '{"id": "pFMhaLmPNxCOrVOscovXvEmjJ","client_id": "jkYoprGPrslAsUGQGnNZopuLh","user_id": "YJhpLBAVwbqZbsZSWICYarPyC","realm_id": "TYapqgRAEUGZtXTvFSZeWBHAD","storage_provider_id": "eivbwbQaRHGnjTPCfDKngnBGk","created_date": 88,"last_updated_date": 44,"client_storage_provider": "SkwmExEyAnyAJiVuBFtOHRwnY","external_client_id": "bTEtTBiBlkmCgkHIWgNwWQmxn"}' | http PUT "http://localhost:8080/feduserconsent/hello world"  X-Api-User:user123
func UpdateFedUserConsent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	feduserconsent := &model.FedUserConsent{}
	if err := readJSON(r, feduserconsent); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := feduserconsent.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	feduserconsent.Prepare()

	if err := feduserconsent.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_consent", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	feduserconsent, _, err = dao.UpdateFedUserConsent(ctx,
		argID,
		feduserconsent)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, feduserconsent)
}

// DeleteFedUserConsent Delete a single record from fed_user_consent table in the keycloak database
// @Summary Delete a record from fed_user_consent
// @Description Delete a single record from fed_user_consent table in the keycloak database
// @Tags FedUserConsent
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.FedUserConsent
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /feduserconsent/{argID} [delete]
// http DELETE "http://localhost:8080/feduserconsent/hello world" X-Api-User:user123
func DeleteFedUserConsent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_consent", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteFedUserConsent(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
