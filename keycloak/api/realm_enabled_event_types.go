package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configRealmEnabledEventTypesRouter(router *httprouter.Router) {
	router.GET("/realmenabledeventtypes", GetAllRealmEnabledEventTypes)
	router.POST("/realmenabledeventtypes", AddRealmEnabledEventTypes)
	router.GET("/realmenabledeventtypes/:argRealmID/:argValue", GetRealmEnabledEventTypes)
	router.PUT("/realmenabledeventtypes/:argRealmID/:argValue", UpdateRealmEnabledEventTypes)
	router.DELETE("/realmenabledeventtypes/:argRealmID/:argValue", DeleteRealmEnabledEventTypes)
}

func configGinRealmEnabledEventTypesRouter(router gin.IRoutes) {
	router.GET("/realmenabledeventtypes", ConverHttprouterToGin(GetAllRealmEnabledEventTypes))
	router.POST("/realmenabledeventtypes", ConverHttprouterToGin(AddRealmEnabledEventTypes))
	router.GET("/realmenabledeventtypes/:argRealmID/:argValue", ConverHttprouterToGin(GetRealmEnabledEventTypes))
	router.PUT("/realmenabledeventtypes/:argRealmID/:argValue", ConverHttprouterToGin(UpdateRealmEnabledEventTypes))
	router.DELETE("/realmenabledeventtypes/:argRealmID/:argValue", ConverHttprouterToGin(DeleteRealmEnabledEventTypes))
}

// GetAllRealmEnabledEventTypes is a function to get a slice of record(s) from realm_enabled_event_types table in the keycloak database
// @Summary Get list of RealmEnabledEventTypes
// @Tags RealmEnabledEventTypes
// @Description GetAllRealmEnabledEventTypes is a handler to get a slice of record(s) from realm_enabled_event_types table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.RealmEnabledEventTypes}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmenabledeventtypes [get]
// http "http://localhost:8080/realmenabledeventtypes?page=0&pagesize=20" X-Api-User:user123
func GetAllRealmEnabledEventTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "realm_enabled_event_types", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllRealmEnabledEventTypes(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetRealmEnabledEventTypes is a function to get a single record from the realm_enabled_event_types table in the keycloak database
// @Summary Get record from table RealmEnabledEventTypes by  argRealmID  argValue
// @Tags RealmEnabledEventTypes
// @ID argRealmID
// @ID argValue
// @Description GetRealmEnabledEventTypes is a function to get a single record from the realm_enabled_event_types table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"
// @Param  argValue path string true "value"
// @Success 200 {object} model.RealmEnabledEventTypes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /realmenabledeventtypes/{argRealmID}/{argValue} [get]
// http "http://localhost:8080/realmenabledeventtypes/hello world/hello world" X-Api-User:user123
func GetRealmEnabledEventTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argValue, err := parseString(ps, "argValue")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_enabled_event_types", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetRealmEnabledEventTypes(ctx, argRealmID, argValue)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddRealmEnabledEventTypes add to add a single record to realm_enabled_event_types table in the keycloak database
// @Summary Add an record to realm_enabled_event_types table
// @Description add to add a single record to realm_enabled_event_types table in the keycloak database
// @Tags RealmEnabledEventTypes
// @Accept  json
// @Produce  json
// @Param RealmEnabledEventTypes body model.RealmEnabledEventTypes true "Add RealmEnabledEventTypes"
// @Success 200 {object} model.RealmEnabledEventTypes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmenabledeventtypes [post]
// echo '{"realm_id": "ipmBUFwJavFMfNdBeqYBIVUAR","value": "VdhPULVCCvHoShsiUhOmFqLEU"}' | http POST "http://localhost:8080/realmenabledeventtypes" X-Api-User:user123
func AddRealmEnabledEventTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	realmenabledeventtypes := &model.RealmEnabledEventTypes{}

	if err := readJSON(r, realmenabledeventtypes); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := realmenabledeventtypes.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	realmenabledeventtypes.Prepare()

	if err := realmenabledeventtypes.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_enabled_event_types", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	realmenabledeventtypes, _, err = dao.AddRealmEnabledEventTypes(ctx, realmenabledeventtypes)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, realmenabledeventtypes)
}

// UpdateRealmEnabledEventTypes Update a single record from realm_enabled_event_types table in the keycloak database
// @Summary Update an record in table realm_enabled_event_types
// @Description Update a single record from realm_enabled_event_types table in the keycloak database
// @Tags RealmEnabledEventTypes
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"// @Param  argValue path string true "value"
// @Param  RealmEnabledEventTypes body model.RealmEnabledEventTypes true "Update RealmEnabledEventTypes record"
// @Success 200 {object} model.RealmEnabledEventTypes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmenabledeventtypes/{argRealmID}/{argValue} [put]
// echo '{"realm_id": "ipmBUFwJavFMfNdBeqYBIVUAR","value": "VdhPULVCCvHoShsiUhOmFqLEU"}' | http PUT "http://localhost:8080/realmenabledeventtypes/hello world/hello world"  X-Api-User:user123
func UpdateRealmEnabledEventTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argValue, err := parseString(ps, "argValue")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	realmenabledeventtypes := &model.RealmEnabledEventTypes{}
	if err := readJSON(r, realmenabledeventtypes); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := realmenabledeventtypes.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	realmenabledeventtypes.Prepare()

	if err := realmenabledeventtypes.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_enabled_event_types", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	realmenabledeventtypes, _, err = dao.UpdateRealmEnabledEventTypes(ctx,
		argRealmID, argValue,
		realmenabledeventtypes)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, realmenabledeventtypes)
}

// DeleteRealmEnabledEventTypes Delete a single record from realm_enabled_event_types table in the keycloak database
// @Summary Delete a record from realm_enabled_event_types
// @Description Delete a single record from realm_enabled_event_types table in the keycloak database
// @Tags RealmEnabledEventTypes
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"// @Param  argValue path string true "value"
// @Success 204 {object} model.RealmEnabledEventTypes
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /realmenabledeventtypes/{argRealmID}/{argValue} [delete]
// http DELETE "http://localhost:8080/realmenabledeventtypes/hello world/hello world" X-Api-User:user123
func DeleteRealmEnabledEventTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argValue, err := parseString(ps, "argValue")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_enabled_event_types", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteRealmEnabledEventTypes(ctx, argRealmID, argValue)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
