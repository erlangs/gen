package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configEventEntityRouter(router *httprouter.Router) {
	router.GET("/evententity", GetAllEventEntity)
	router.POST("/evententity", AddEventEntity)
	router.GET("/evententity/:argID", GetEventEntity)
	router.PUT("/evententity/:argID", UpdateEventEntity)
	router.DELETE("/evententity/:argID", DeleteEventEntity)
}

func configGinEventEntityRouter(router gin.IRoutes) {
	router.GET("/evententity", ConverHttprouterToGin(GetAllEventEntity))
	router.POST("/evententity", ConverHttprouterToGin(AddEventEntity))
	router.GET("/evententity/:argID", ConverHttprouterToGin(GetEventEntity))
	router.PUT("/evententity/:argID", ConverHttprouterToGin(UpdateEventEntity))
	router.DELETE("/evententity/:argID", ConverHttprouterToGin(DeleteEventEntity))
}

// GetAllEventEntity is a function to get a slice of record(s) from event_entity table in the keycloak database
// @Summary Get list of EventEntity
// @Tags EventEntity
// @Description GetAllEventEntity is a handler to get a slice of record(s) from event_entity table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.EventEntity}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /evententity [get]
// http "http://localhost:8080/evententity?page=0&pagesize=20" X-Api-User:user123
func GetAllEventEntity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "event_entity", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllEventEntity(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetEventEntity is a function to get a single record from the event_entity table in the keycloak database
// @Summary Get record from table EventEntity by  argID
// @Tags EventEntity
// @ID argID
// @Description GetEventEntity is a function to get a single record from the event_entity table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.EventEntity
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /evententity/{argID} [get]
// http "http://localhost:8080/evententity/hello world" X-Api-User:user123
func GetEventEntity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "event_entity", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetEventEntity(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddEventEntity add to add a single record to event_entity table in the keycloak database
// @Summary Add an record to event_entity table
// @Description add to add a single record to event_entity table in the keycloak database
// @Tags EventEntity
// @Accept  json
// @Produce  json
// @Param EventEntity body model.EventEntity true "Add EventEntity"
// @Success 200 {object} model.EventEntity
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /evententity [post]
// echo '{"id": "IvLpjvvpUJnycgJcrblZiZGZw","client_id": "rCgjdCbhvBidEOAGyIeGQosGW","details_json": "CRHrilyyfspQOFrZGZTrvmZCv","error": "EkLdDsIhiMMGZCvOSEJbnmTPH","ip_address": "LjrpHPfZovGjUidGdIKaLCyke","realm_id": "LpvnClEbUPVokedoHjMYgveRB","session_id": "tRrDPFkQPGjfLHaXeDRpUSiZL","event_time": 8,"type": "clZfcixSOdYQgeGBrmnHcqOEL","user_id": "nUcoAeFTKhdXweJyQFiNWLtYr"}' | http POST "http://localhost:8080/evententity" X-Api-User:user123
func AddEventEntity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	evententity := &model.EventEntity{}

	if err := readJSON(r, evententity); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := evententity.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	evententity.Prepare()

	if err := evententity.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "event_entity", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	evententity, _, err = dao.AddEventEntity(ctx, evententity)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, evententity)
}

// UpdateEventEntity Update a single record from event_entity table in the keycloak database
// @Summary Update an record in table event_entity
// @Description Update a single record from event_entity table in the keycloak database
// @Tags EventEntity
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  EventEntity body model.EventEntity true "Update EventEntity record"
// @Success 200 {object} model.EventEntity
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /evententity/{argID} [put]
// echo '{"id": "IvLpjvvpUJnycgJcrblZiZGZw","client_id": "rCgjdCbhvBidEOAGyIeGQosGW","details_json": "CRHrilyyfspQOFrZGZTrvmZCv","error": "EkLdDsIhiMMGZCvOSEJbnmTPH","ip_address": "LjrpHPfZovGjUidGdIKaLCyke","realm_id": "LpvnClEbUPVokedoHjMYgveRB","session_id": "tRrDPFkQPGjfLHaXeDRpUSiZL","event_time": 8,"type": "clZfcixSOdYQgeGBrmnHcqOEL","user_id": "nUcoAeFTKhdXweJyQFiNWLtYr"}' | http PUT "http://localhost:8080/evententity/hello world"  X-Api-User:user123
func UpdateEventEntity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	evententity := &model.EventEntity{}
	if err := readJSON(r, evententity); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := evententity.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	evententity.Prepare()

	if err := evententity.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "event_entity", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	evententity, _, err = dao.UpdateEventEntity(ctx,
		argID,
		evententity)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, evententity)
}

// DeleteEventEntity Delete a single record from event_entity table in the keycloak database
// @Summary Delete a record from event_entity
// @Description Delete a single record from event_entity table in the keycloak database
// @Tags EventEntity
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.EventEntity
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /evententity/{argID} [delete]
// http DELETE "http://localhost:8080/evententity/hello world" X-Api-User:user123
func DeleteEventEntity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "event_entity", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteEventEntity(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
