package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configFederatedUserRouter(router *httprouter.Router) {
	router.GET("/federateduser", GetAllFederatedUser)
	router.POST("/federateduser", AddFederatedUser)
	router.GET("/federateduser/:argID", GetFederatedUser)
	router.PUT("/federateduser/:argID", UpdateFederatedUser)
	router.DELETE("/federateduser/:argID", DeleteFederatedUser)
}

func configGinFederatedUserRouter(router gin.IRoutes) {
	router.GET("/federateduser", ConverHttprouterToGin(GetAllFederatedUser))
	router.POST("/federateduser", ConverHttprouterToGin(AddFederatedUser))
	router.GET("/federateduser/:argID", ConverHttprouterToGin(GetFederatedUser))
	router.PUT("/federateduser/:argID", ConverHttprouterToGin(UpdateFederatedUser))
	router.DELETE("/federateduser/:argID", ConverHttprouterToGin(DeleteFederatedUser))
}

// GetAllFederatedUser is a function to get a slice of record(s) from federated_user table in the keycloak database
// @Summary Get list of FederatedUser
// @Tags FederatedUser
// @Description GetAllFederatedUser is a handler to get a slice of record(s) from federated_user table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.FederatedUser}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /federateduser [get]
// http "http://localhost:8080/federateduser?page=0&pagesize=20" X-Api-User:user123
func GetAllFederatedUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "federated_user", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllFederatedUser(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetFederatedUser is a function to get a single record from the federated_user table in the keycloak database
// @Summary Get record from table FederatedUser by  argID
// @Tags FederatedUser
// @ID argID
// @Description GetFederatedUser is a function to get a single record from the federated_user table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.FederatedUser
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /federateduser/{argID} [get]
// http "http://localhost:8080/federateduser/hello world" X-Api-User:user123
func GetFederatedUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "federated_user", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetFederatedUser(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddFederatedUser add to add a single record to federated_user table in the keycloak database
// @Summary Add an record to federated_user table
// @Description add to add a single record to federated_user table in the keycloak database
// @Tags FederatedUser
// @Accept  json
// @Produce  json
// @Param FederatedUser body model.FederatedUser true "Add FederatedUser"
// @Success 200 {object} model.FederatedUser
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /federateduser [post]
// echo '{"id": "NNboLSEGIKCjZECqmiXOiXHQU","storage_provider_id": "XksvGMqDMiZxJawynsfRfCPkb","realm_id": "ovqVrytYMGQZslJRIXUUNtPuD"}' | http POST "http://localhost:8080/federateduser" X-Api-User:user123
func AddFederatedUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	federateduser := &model.FederatedUser{}

	if err := readJSON(r, federateduser); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := federateduser.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	federateduser.Prepare()

	if err := federateduser.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "federated_user", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	federateduser, _, err = dao.AddFederatedUser(ctx, federateduser)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, federateduser)
}

// UpdateFederatedUser Update a single record from federated_user table in the keycloak database
// @Summary Update an record in table federated_user
// @Description Update a single record from federated_user table in the keycloak database
// @Tags FederatedUser
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  FederatedUser body model.FederatedUser true "Update FederatedUser record"
// @Success 200 {object} model.FederatedUser
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /federateduser/{argID} [put]
// echo '{"id": "NNboLSEGIKCjZECqmiXOiXHQU","storage_provider_id": "XksvGMqDMiZxJawynsfRfCPkb","realm_id": "ovqVrytYMGQZslJRIXUUNtPuD"}' | http PUT "http://localhost:8080/federateduser/hello world"  X-Api-User:user123
func UpdateFederatedUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	federateduser := &model.FederatedUser{}
	if err := readJSON(r, federateduser); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := federateduser.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	federateduser.Prepare()

	if err := federateduser.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "federated_user", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	federateduser, _, err = dao.UpdateFederatedUser(ctx,
		argID,
		federateduser)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, federateduser)
}

// DeleteFederatedUser Delete a single record from federated_user table in the keycloak database
// @Summary Delete a record from federated_user
// @Description Delete a single record from federated_user table in the keycloak database
// @Tags FederatedUser
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.FederatedUser
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /federateduser/{argID} [delete]
// http DELETE "http://localhost:8080/federateduser/hello world" X-Api-User:user123
func DeleteFederatedUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "federated_user", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteFederatedUser(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
