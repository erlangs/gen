package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configAdminEventEntityRouter(router *httprouter.Router) {
	router.GET("/adminevententity", GetAllAdminEventEntity)
	router.POST("/adminevententity", AddAdminEventEntity)
	router.GET("/adminevententity/:argID", GetAdminEventEntity)
	router.PUT("/adminevententity/:argID", UpdateAdminEventEntity)
	router.DELETE("/adminevententity/:argID", DeleteAdminEventEntity)
}

func configGinAdminEventEntityRouter(router gin.IRoutes) {
	router.GET("/adminevententity", ConverHttprouterToGin(GetAllAdminEventEntity))
	router.POST("/adminevententity", ConverHttprouterToGin(AddAdminEventEntity))
	router.GET("/adminevententity/:argID", ConverHttprouterToGin(GetAdminEventEntity))
	router.PUT("/adminevententity/:argID", ConverHttprouterToGin(UpdateAdminEventEntity))
	router.DELETE("/adminevententity/:argID", ConverHttprouterToGin(DeleteAdminEventEntity))
}

// GetAllAdminEventEntity is a function to get a slice of record(s) from admin_event_entity table in the keycloak database
// @Summary Get list of AdminEventEntity
// @Tags AdminEventEntity
// @Description GetAllAdminEventEntity is a handler to get a slice of record(s) from admin_event_entity table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.AdminEventEntity}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /adminevententity [get]
// http "http://localhost:8080/adminevententity?page=0&pagesize=20" X-Api-User:user123
func GetAllAdminEventEntity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "admin_event_entity", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllAdminEventEntity(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetAdminEventEntity is a function to get a single record from the admin_event_entity table in the keycloak database
// @Summary Get record from table AdminEventEntity by  argID
// @Tags AdminEventEntity
// @ID argID
// @Description GetAdminEventEntity is a function to get a single record from the admin_event_entity table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.AdminEventEntity
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /adminevententity/{argID} [get]
// http "http://localhost:8080/adminevententity/hello world" X-Api-User:user123
func GetAdminEventEntity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "admin_event_entity", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetAdminEventEntity(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddAdminEventEntity add to add a single record to admin_event_entity table in the keycloak database
// @Summary Add an record to admin_event_entity table
// @Description add to add a single record to admin_event_entity table in the keycloak database
// @Tags AdminEventEntity
// @Accept  json
// @Produce  json
// @Param AdminEventEntity body model.AdminEventEntity true "Add AdminEventEntity"
// @Success 200 {object} model.AdminEventEntity
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /adminevententity [post]
// echo '{"id": "yFnCrnQFhBBplCAFgBmMXgGwi","admin_event_time": 73,"realm_id": "sDDKctFOTpCxovXNhmlFQquCe","operation_type": "BvWRjvNGPoaShTlssqbwnJPOW","auth_realm_id": "HbiFvrQErBbLOyHENCkGDMMJe","auth_client_id": "lMnIeIDrfpJvfZGGQAxWlyepV","auth_user_id": "ZAcktnRBmAngeQqwGkisSnwsg","ip_address": "kbOAQoGpIpxwOjvTfGqGupySi","resource_path": "oBNjoSOyQgAYrqjkDyeRduxHy","representation": "vvbUKFvDNQnWobsTfbpiyUepN","error": "ZKROLXsRWUhjiHdGCxlRUkVkd","resource_type": "NSlfZZNHohHXxPWqsCGuFYoro"}' | http POST "http://localhost:8080/adminevententity" X-Api-User:user123
func AddAdminEventEntity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	adminevententity := &model.AdminEventEntity{}

	if err := readJSON(r, adminevententity); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := adminevententity.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	adminevententity.Prepare()

	if err := adminevententity.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "admin_event_entity", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	adminevententity, _, err = dao.AddAdminEventEntity(ctx, adminevententity)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, adminevententity)
}

// UpdateAdminEventEntity Update a single record from admin_event_entity table in the keycloak database
// @Summary Update an record in table admin_event_entity
// @Description Update a single record from admin_event_entity table in the keycloak database
// @Tags AdminEventEntity
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  AdminEventEntity body model.AdminEventEntity true "Update AdminEventEntity record"
// @Success 200 {object} model.AdminEventEntity
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /adminevententity/{argID} [put]
// echo '{"id": "yFnCrnQFhBBplCAFgBmMXgGwi","admin_event_time": 73,"realm_id": "sDDKctFOTpCxovXNhmlFQquCe","operation_type": "BvWRjvNGPoaShTlssqbwnJPOW","auth_realm_id": "HbiFvrQErBbLOyHENCkGDMMJe","auth_client_id": "lMnIeIDrfpJvfZGGQAxWlyepV","auth_user_id": "ZAcktnRBmAngeQqwGkisSnwsg","ip_address": "kbOAQoGpIpxwOjvTfGqGupySi","resource_path": "oBNjoSOyQgAYrqjkDyeRduxHy","representation": "vvbUKFvDNQnWobsTfbpiyUepN","error": "ZKROLXsRWUhjiHdGCxlRUkVkd","resource_type": "NSlfZZNHohHXxPWqsCGuFYoro"}' | http PUT "http://localhost:8080/adminevententity/hello world"  X-Api-User:user123
func UpdateAdminEventEntity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	adminevententity := &model.AdminEventEntity{}
	if err := readJSON(r, adminevententity); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := adminevententity.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	adminevententity.Prepare()

	if err := adminevententity.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "admin_event_entity", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	adminevententity, _, err = dao.UpdateAdminEventEntity(ctx,
		argID,
		adminevententity)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, adminevententity)
}

// DeleteAdminEventEntity Delete a single record from admin_event_entity table in the keycloak database
// @Summary Delete a record from admin_event_entity
// @Description Delete a single record from admin_event_entity table in the keycloak database
// @Tags AdminEventEntity
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.AdminEventEntity
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /adminevententity/{argID} [delete]
// http DELETE "http://localhost:8080/adminevententity/hello world" X-Api-User:user123
func DeleteAdminEventEntity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "admin_event_entity", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteAdminEventEntity(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
