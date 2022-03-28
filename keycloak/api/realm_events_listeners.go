package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configRealmEventsListenersRouter(router *httprouter.Router) {
	router.GET("/realmeventslisteners", GetAllRealmEventsListeners)
	router.POST("/realmeventslisteners", AddRealmEventsListeners)
	router.GET("/realmeventslisteners/:argRealmID/:argValue", GetRealmEventsListeners)
	router.PUT("/realmeventslisteners/:argRealmID/:argValue", UpdateRealmEventsListeners)
	router.DELETE("/realmeventslisteners/:argRealmID/:argValue", DeleteRealmEventsListeners)
}

func configGinRealmEventsListenersRouter(router gin.IRoutes) {
	router.GET("/realmeventslisteners", ConverHttprouterToGin(GetAllRealmEventsListeners))
	router.POST("/realmeventslisteners", ConverHttprouterToGin(AddRealmEventsListeners))
	router.GET("/realmeventslisteners/:argRealmID/:argValue", ConverHttprouterToGin(GetRealmEventsListeners))
	router.PUT("/realmeventslisteners/:argRealmID/:argValue", ConverHttprouterToGin(UpdateRealmEventsListeners))
	router.DELETE("/realmeventslisteners/:argRealmID/:argValue", ConverHttprouterToGin(DeleteRealmEventsListeners))
}

// GetAllRealmEventsListeners is a function to get a slice of record(s) from realm_events_listeners table in the keycloak database
// @Summary Get list of RealmEventsListeners
// @Tags RealmEventsListeners
// @Description GetAllRealmEventsListeners is a handler to get a slice of record(s) from realm_events_listeners table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.RealmEventsListeners}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmeventslisteners [get]
// http "http://localhost:8080/realmeventslisteners?page=0&pagesize=20" X-Api-User:user123
func GetAllRealmEventsListeners(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "realm_events_listeners", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllRealmEventsListeners(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetRealmEventsListeners is a function to get a single record from the realm_events_listeners table in the keycloak database
// @Summary Get record from table RealmEventsListeners by  argRealmID  argValue
// @Tags RealmEventsListeners
// @ID argRealmID
// @ID argValue
// @Description GetRealmEventsListeners is a function to get a single record from the realm_events_listeners table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"
// @Param  argValue path string true "value"
// @Success 200 {object} model.RealmEventsListeners
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /realmeventslisteners/{argRealmID}/{argValue} [get]
// http "http://localhost:8080/realmeventslisteners/hello world/hello world" X-Api-User:user123
func GetRealmEventsListeners(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "realm_events_listeners", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetRealmEventsListeners(ctx, argRealmID, argValue)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddRealmEventsListeners add to add a single record to realm_events_listeners table in the keycloak database
// @Summary Add an record to realm_events_listeners table
// @Description add to add a single record to realm_events_listeners table in the keycloak database
// @Tags RealmEventsListeners
// @Accept  json
// @Produce  json
// @Param RealmEventsListeners body model.RealmEventsListeners true "Add RealmEventsListeners"
// @Success 200 {object} model.RealmEventsListeners
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmeventslisteners [post]
// echo '{"realm_id": "ZRdepJcqOAbTFoYfLSJDgAOxN","value": "aoUQXtjNgbnwTLJLIlIZaSGLg"}' | http POST "http://localhost:8080/realmeventslisteners" X-Api-User:user123
func AddRealmEventsListeners(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	realmeventslisteners := &model.RealmEventsListeners{}

	if err := readJSON(r, realmeventslisteners); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := realmeventslisteners.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	realmeventslisteners.Prepare()

	if err := realmeventslisteners.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_events_listeners", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	realmeventslisteners, _, err = dao.AddRealmEventsListeners(ctx, realmeventslisteners)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, realmeventslisteners)
}

// UpdateRealmEventsListeners Update a single record from realm_events_listeners table in the keycloak database
// @Summary Update an record in table realm_events_listeners
// @Description Update a single record from realm_events_listeners table in the keycloak database
// @Tags RealmEventsListeners
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"// @Param  argValue path string true "value"
// @Param  RealmEventsListeners body model.RealmEventsListeners true "Update RealmEventsListeners record"
// @Success 200 {object} model.RealmEventsListeners
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmeventslisteners/{argRealmID}/{argValue} [put]
// echo '{"realm_id": "ZRdepJcqOAbTFoYfLSJDgAOxN","value": "aoUQXtjNgbnwTLJLIlIZaSGLg"}' | http PUT "http://localhost:8080/realmeventslisteners/hello world/hello world"  X-Api-User:user123
func UpdateRealmEventsListeners(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	realmeventslisteners := &model.RealmEventsListeners{}
	if err := readJSON(r, realmeventslisteners); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := realmeventslisteners.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	realmeventslisteners.Prepare()

	if err := realmeventslisteners.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_events_listeners", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	realmeventslisteners, _, err = dao.UpdateRealmEventsListeners(ctx,
		argRealmID, argValue,
		realmeventslisteners)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, realmeventslisteners)
}

// DeleteRealmEventsListeners Delete a single record from realm_events_listeners table in the keycloak database
// @Summary Delete a record from realm_events_listeners
// @Description Delete a single record from realm_events_listeners table in the keycloak database
// @Tags RealmEventsListeners
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"// @Param  argValue path string true "value"
// @Success 204 {object} model.RealmEventsListeners
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /realmeventslisteners/{argRealmID}/{argValue} [delete]
// http DELETE "http://localhost:8080/realmeventslisteners/hello world/hello world" X-Api-User:user123
func DeleteRealmEventsListeners(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "realm_events_listeners", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteRealmEventsListeners(ctx, argRealmID, argValue)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
