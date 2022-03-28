package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configRealmDefaultGroupsRouter(router *httprouter.Router) {
	router.GET("/realmdefaultgroups", GetAllRealmDefaultGroups)
	router.POST("/realmdefaultgroups", AddRealmDefaultGroups)
	router.GET("/realmdefaultgroups/:argRealmID/:argGroupID", GetRealmDefaultGroups)
	router.PUT("/realmdefaultgroups/:argRealmID/:argGroupID", UpdateRealmDefaultGroups)
	router.DELETE("/realmdefaultgroups/:argRealmID/:argGroupID", DeleteRealmDefaultGroups)
}

func configGinRealmDefaultGroupsRouter(router gin.IRoutes) {
	router.GET("/realmdefaultgroups", ConverHttprouterToGin(GetAllRealmDefaultGroups))
	router.POST("/realmdefaultgroups", ConverHttprouterToGin(AddRealmDefaultGroups))
	router.GET("/realmdefaultgroups/:argRealmID/:argGroupID", ConverHttprouterToGin(GetRealmDefaultGroups))
	router.PUT("/realmdefaultgroups/:argRealmID/:argGroupID", ConverHttprouterToGin(UpdateRealmDefaultGroups))
	router.DELETE("/realmdefaultgroups/:argRealmID/:argGroupID", ConverHttprouterToGin(DeleteRealmDefaultGroups))
}

// GetAllRealmDefaultGroups is a function to get a slice of record(s) from realm_default_groups table in the keycloak database
// @Summary Get list of RealmDefaultGroups
// @Tags RealmDefaultGroups
// @Description GetAllRealmDefaultGroups is a handler to get a slice of record(s) from realm_default_groups table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.RealmDefaultGroups}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmdefaultgroups [get]
// http "http://localhost:8080/realmdefaultgroups?page=0&pagesize=20" X-Api-User:user123
func GetAllRealmDefaultGroups(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "realm_default_groups", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllRealmDefaultGroups(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetRealmDefaultGroups is a function to get a single record from the realm_default_groups table in the keycloak database
// @Summary Get record from table RealmDefaultGroups by  argRealmID  argGroupID
// @Tags RealmDefaultGroups
// @ID argRealmID
// @ID argGroupID
// @Description GetRealmDefaultGroups is a function to get a single record from the realm_default_groups table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"
// @Param  argGroupID path string true "group_id"
// @Success 200 {object} model.RealmDefaultGroups
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /realmdefaultgroups/{argRealmID}/{argGroupID} [get]
// http "http://localhost:8080/realmdefaultgroups/hello world/hello world" X-Api-User:user123
func GetRealmDefaultGroups(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argGroupID, err := parseString(ps, "argGroupID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_default_groups", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetRealmDefaultGroups(ctx, argRealmID, argGroupID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddRealmDefaultGroups add to add a single record to realm_default_groups table in the keycloak database
// @Summary Add an record to realm_default_groups table
// @Description add to add a single record to realm_default_groups table in the keycloak database
// @Tags RealmDefaultGroups
// @Accept  json
// @Produce  json
// @Param RealmDefaultGroups body model.RealmDefaultGroups true "Add RealmDefaultGroups"
// @Success 200 {object} model.RealmDefaultGroups
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmdefaultgroups [post]
// echo '{"realm_id": "ttDuRQhnCWCDHRhNZIOvZmOvT","group_id": "WfcBYmAvtxRBsyEWNjowAMoLC"}' | http POST "http://localhost:8080/realmdefaultgroups" X-Api-User:user123
func AddRealmDefaultGroups(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	realmdefaultgroups := &model.RealmDefaultGroups{}

	if err := readJSON(r, realmdefaultgroups); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := realmdefaultgroups.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	realmdefaultgroups.Prepare()

	if err := realmdefaultgroups.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_default_groups", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	realmdefaultgroups, _, err = dao.AddRealmDefaultGroups(ctx, realmdefaultgroups)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, realmdefaultgroups)
}

// UpdateRealmDefaultGroups Update a single record from realm_default_groups table in the keycloak database
// @Summary Update an record in table realm_default_groups
// @Description Update a single record from realm_default_groups table in the keycloak database
// @Tags RealmDefaultGroups
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"// @Param  argGroupID path string true "group_id"
// @Param  RealmDefaultGroups body model.RealmDefaultGroups true "Update RealmDefaultGroups record"
// @Success 200 {object} model.RealmDefaultGroups
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmdefaultgroups/{argRealmID}/{argGroupID} [put]
// echo '{"realm_id": "ttDuRQhnCWCDHRhNZIOvZmOvT","group_id": "WfcBYmAvtxRBsyEWNjowAMoLC"}' | http PUT "http://localhost:8080/realmdefaultgroups/hello world/hello world"  X-Api-User:user123
func UpdateRealmDefaultGroups(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argGroupID, err := parseString(ps, "argGroupID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	realmdefaultgroups := &model.RealmDefaultGroups{}
	if err := readJSON(r, realmdefaultgroups); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := realmdefaultgroups.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	realmdefaultgroups.Prepare()

	if err := realmdefaultgroups.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_default_groups", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	realmdefaultgroups, _, err = dao.UpdateRealmDefaultGroups(ctx,
		argRealmID, argGroupID,
		realmdefaultgroups)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, realmdefaultgroups)
}

// DeleteRealmDefaultGroups Delete a single record from realm_default_groups table in the keycloak database
// @Summary Delete a record from realm_default_groups
// @Description Delete a single record from realm_default_groups table in the keycloak database
// @Tags RealmDefaultGroups
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"// @Param  argGroupID path string true "group_id"
// @Success 204 {object} model.RealmDefaultGroups
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /realmdefaultgroups/{argRealmID}/{argGroupID} [delete]
// http DELETE "http://localhost:8080/realmdefaultgroups/hello world/hello world" X-Api-User:user123
func DeleteRealmDefaultGroups(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argGroupID, err := parseString(ps, "argGroupID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_default_groups", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteRealmDefaultGroups(ctx, argRealmID, argGroupID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
