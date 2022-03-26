package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configRealmAttributeRouter(router *httprouter.Router) {
	router.GET("/realmattribute", GetAllRealmAttribute)
	router.POST("/realmattribute", AddRealmAttribute)
	router.GET("/realmattribute/:argName/:argRealmID", GetRealmAttribute)
	router.PUT("/realmattribute/:argName/:argRealmID", UpdateRealmAttribute)
	router.DELETE("/realmattribute/:argName/:argRealmID", DeleteRealmAttribute)
}

func configGinRealmAttributeRouter(router gin.IRoutes) {
	router.GET("/realmattribute", ConverHttprouterToGin(GetAllRealmAttribute))
	router.POST("/realmattribute", ConverHttprouterToGin(AddRealmAttribute))
	router.GET("/realmattribute/:argName/:argRealmID", ConverHttprouterToGin(GetRealmAttribute))
	router.PUT("/realmattribute/:argName/:argRealmID", ConverHttprouterToGin(UpdateRealmAttribute))
	router.DELETE("/realmattribute/:argName/:argRealmID", ConverHttprouterToGin(DeleteRealmAttribute))
}

// GetAllRealmAttribute is a function to get a slice of record(s) from realm_attribute table in the keycloak database
// @Summary Get list of RealmAttribute
// @Tags RealmAttribute
// @Description GetAllRealmAttribute is a handler to get a slice of record(s) from realm_attribute table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.RealmAttribute}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmattribute [get]
// http "http://localhost:8080/realmattribute?page=0&pagesize=20" X-Api-User:user123
func GetAllRealmAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "realm_attribute", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllRealmAttribute(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetRealmAttribute is a function to get a single record from the realm_attribute table in the keycloak database
// @Summary Get record from table RealmAttribute by  argName  argRealmID
// @Tags RealmAttribute
// @ID argName
// @ID argRealmID
// @Description GetRealmAttribute is a function to get a single record from the realm_attribute table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argName path string true "name"
// @Param  argRealmID path string true "realm_id"
// @Success 200 {object} model.RealmAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /realmattribute/{argName}/{argRealmID} [get]
// http "http://localhost:8080/realmattribute/hello world/hello world" X-Api-User:user123
func GetRealmAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_attribute", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetRealmAttribute(ctx, argName, argRealmID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddRealmAttribute add to add a single record to realm_attribute table in the keycloak database
// @Summary Add an record to realm_attribute table
// @Description add to add a single record to realm_attribute table in the keycloak database
// @Tags RealmAttribute
// @Accept  json
// @Produce  json
// @Param RealmAttribute body model.RealmAttribute true "Add RealmAttribute"
// @Success 200 {object} model.RealmAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmattribute [post]
// echo '{"name": "nVrZJdSgvQltNDTuOdapOPcoH","realm_id": "pWOaRGSgTkaWeBwGAsAYItuys","value": "FlYYBFcYgcthBKqSEGkoClSab"}' | http POST "http://localhost:8080/realmattribute" X-Api-User:user123
func AddRealmAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	realmattribute := &model.RealmAttribute{}

	if err := readJSON(r, realmattribute); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := realmattribute.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	realmattribute.Prepare()

	if err := realmattribute.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_attribute", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	realmattribute, _, err = dao.AddRealmAttribute(ctx, realmattribute)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, realmattribute)
}

// UpdateRealmAttribute Update a single record from realm_attribute table in the keycloak database
// @Summary Update an record in table realm_attribute
// @Description Update a single record from realm_attribute table in the keycloak database
// @Tags RealmAttribute
// @Accept  json
// @Produce  json
// @Param  argName path string true "name"// @Param  argRealmID path string true "realm_id"
// @Param  RealmAttribute body model.RealmAttribute true "Update RealmAttribute record"
// @Success 200 {object} model.RealmAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmattribute/{argName}/{argRealmID} [put]
// echo '{"name": "nVrZJdSgvQltNDTuOdapOPcoH","realm_id": "pWOaRGSgTkaWeBwGAsAYItuys","value": "FlYYBFcYgcthBKqSEGkoClSab"}' | http PUT "http://localhost:8080/realmattribute/hello world/hello world"  X-Api-User:user123
func UpdateRealmAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	realmattribute := &model.RealmAttribute{}
	if err := readJSON(r, realmattribute); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := realmattribute.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	realmattribute.Prepare()

	if err := realmattribute.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_attribute", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	realmattribute, _, err = dao.UpdateRealmAttribute(ctx,
		argName, argRealmID,
		realmattribute)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, realmattribute)
}

// DeleteRealmAttribute Delete a single record from realm_attribute table in the keycloak database
// @Summary Delete a record from realm_attribute
// @Description Delete a single record from realm_attribute table in the keycloak database
// @Tags RealmAttribute
// @Accept  json
// @Produce  json
// @Param  argName path string true "name"// @Param  argRealmID path string true "realm_id"
// @Success 204 {object} model.RealmAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /realmattribute/{argName}/{argRealmID} [delete]
// http DELETE "http://localhost:8080/realmattribute/hello world/hello world" X-Api-User:user123
func DeleteRealmAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_attribute", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteRealmAttribute(ctx, argName, argRealmID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
