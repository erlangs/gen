package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configFederatedIdentityRouter(router *httprouter.Router) {
	router.GET("/federatedidentity", GetAllFederatedIdentity)
	router.POST("/federatedidentity", AddFederatedIdentity)
	router.GET("/federatedidentity/:argIdentityProvider/:argUserID", GetFederatedIdentity)
	router.PUT("/federatedidentity/:argIdentityProvider/:argUserID", UpdateFederatedIdentity)
	router.DELETE("/federatedidentity/:argIdentityProvider/:argUserID", DeleteFederatedIdentity)
}

func configGinFederatedIdentityRouter(router gin.IRoutes) {
	router.GET("/federatedidentity", ConverHttprouterToGin(GetAllFederatedIdentity))
	router.POST("/federatedidentity", ConverHttprouterToGin(AddFederatedIdentity))
	router.GET("/federatedidentity/:argIdentityProvider/:argUserID", ConverHttprouterToGin(GetFederatedIdentity))
	router.PUT("/federatedidentity/:argIdentityProvider/:argUserID", ConverHttprouterToGin(UpdateFederatedIdentity))
	router.DELETE("/federatedidentity/:argIdentityProvider/:argUserID", ConverHttprouterToGin(DeleteFederatedIdentity))
}

// GetAllFederatedIdentity is a function to get a slice of record(s) from federated_identity table in the keycloak database
// @Summary Get list of FederatedIdentity
// @Tags FederatedIdentity
// @Description GetAllFederatedIdentity is a handler to get a slice of record(s) from federated_identity table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.FederatedIdentity}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /federatedidentity [get]
// http "http://localhost:8080/federatedidentity?page=0&pagesize=20" X-Api-User:user123
func GetAllFederatedIdentity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "federated_identity", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllFederatedIdentity(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetFederatedIdentity is a function to get a single record from the federated_identity table in the keycloak database
// @Summary Get record from table FederatedIdentity by  argIdentityProvider  argUserID
// @Tags FederatedIdentity
// @ID argIdentityProvider
// @ID argUserID
// @Description GetFederatedIdentity is a function to get a single record from the federated_identity table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argIdentityProvider path string true "identity_provider"
// @Param  argUserID path string true "user_id"
// @Success 200 {object} model.FederatedIdentity
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /federatedidentity/{argIdentityProvider}/{argUserID} [get]
// http "http://localhost:8080/federatedidentity/hello world/hello world" X-Api-User:user123
func GetFederatedIdentity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argIdentityProvider, err := parseString(ps, "argIdentityProvider")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argUserID, err := parseString(ps, "argUserID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "federated_identity", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetFederatedIdentity(ctx, argIdentityProvider, argUserID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddFederatedIdentity add to add a single record to federated_identity table in the keycloak database
// @Summary Add an record to federated_identity table
// @Description add to add a single record to federated_identity table in the keycloak database
// @Tags FederatedIdentity
// @Accept  json
// @Produce  json
// @Param FederatedIdentity body model.FederatedIdentity true "Add FederatedIdentity"
// @Success 200 {object} model.FederatedIdentity
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /federatedidentity [post]
// echo '{"identity_provider": "SdPxPjwhfdwnGkFHSLwHRVnJe","realm_id": "oyWVkqaHYJplKCmTLTUdJLqxB","federated_user_id": "ZxTkgnKhnbIhSTkgxdBTadRaW","federated_username": "fTsAvLvgWkwNEAehGSkDuRWkA","token": "CJuOMYQtCOhnxibEPHMqpNxKc","user_id": "TeojvKKpqZmUFAVfXmUASvKcw"}' | http POST "http://localhost:8080/federatedidentity" X-Api-User:user123
func AddFederatedIdentity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	federatedidentity := &model.FederatedIdentity{}

	if err := readJSON(r, federatedidentity); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := federatedidentity.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	federatedidentity.Prepare()

	if err := federatedidentity.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "federated_identity", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	federatedidentity, _, err = dao.AddFederatedIdentity(ctx, federatedidentity)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, federatedidentity)
}

// UpdateFederatedIdentity Update a single record from federated_identity table in the keycloak database
// @Summary Update an record in table federated_identity
// @Description Update a single record from federated_identity table in the keycloak database
// @Tags FederatedIdentity
// @Accept  json
// @Produce  json
// @Param  argIdentityProvider path string true "identity_provider"// @Param  argUserID path string true "user_id"
// @Param  FederatedIdentity body model.FederatedIdentity true "Update FederatedIdentity record"
// @Success 200 {object} model.FederatedIdentity
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /federatedidentity/{argIdentityProvider}/{argUserID} [put]
// echo '{"identity_provider": "SdPxPjwhfdwnGkFHSLwHRVnJe","realm_id": "oyWVkqaHYJplKCmTLTUdJLqxB","federated_user_id": "ZxTkgnKhnbIhSTkgxdBTadRaW","federated_username": "fTsAvLvgWkwNEAehGSkDuRWkA","token": "CJuOMYQtCOhnxibEPHMqpNxKc","user_id": "TeojvKKpqZmUFAVfXmUASvKcw"}' | http PUT "http://localhost:8080/federatedidentity/hello world/hello world"  X-Api-User:user123
func UpdateFederatedIdentity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argIdentityProvider, err := parseString(ps, "argIdentityProvider")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argUserID, err := parseString(ps, "argUserID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	federatedidentity := &model.FederatedIdentity{}
	if err := readJSON(r, federatedidentity); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := federatedidentity.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	federatedidentity.Prepare()

	if err := federatedidentity.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "federated_identity", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	federatedidentity, _, err = dao.UpdateFederatedIdentity(ctx,
		argIdentityProvider, argUserID,
		federatedidentity)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, federatedidentity)
}

// DeleteFederatedIdentity Delete a single record from federated_identity table in the keycloak database
// @Summary Delete a record from federated_identity
// @Description Delete a single record from federated_identity table in the keycloak database
// @Tags FederatedIdentity
// @Accept  json
// @Produce  json
// @Param  argIdentityProvider path string true "identity_provider"// @Param  argUserID path string true "user_id"
// @Success 204 {object} model.FederatedIdentity
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /federatedidentity/{argIdentityProvider}/{argUserID} [delete]
// http DELETE "http://localhost:8080/federatedidentity/hello world/hello world" X-Api-User:user123
func DeleteFederatedIdentity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argIdentityProvider, err := parseString(ps, "argIdentityProvider")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argUserID, err := parseString(ps, "argUserID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "federated_identity", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteFederatedIdentity(ctx, argIdentityProvider, argUserID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
