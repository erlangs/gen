package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configDatabasechangeloglockRouter(router *httprouter.Router) {
	router.GET("/databasechangeloglock", GetAllDatabasechangeloglock)
	router.POST("/databasechangeloglock", AddDatabasechangeloglock)
	router.GET("/databasechangeloglock/:argID", GetDatabasechangeloglock)
	router.PUT("/databasechangeloglock/:argID", UpdateDatabasechangeloglock)
	router.DELETE("/databasechangeloglock/:argID", DeleteDatabasechangeloglock)
}

func configGinDatabasechangeloglockRouter(router gin.IRoutes) {
	router.GET("/databasechangeloglock", ConverHttprouterToGin(GetAllDatabasechangeloglock))
	router.POST("/databasechangeloglock", ConverHttprouterToGin(AddDatabasechangeloglock))
	router.GET("/databasechangeloglock/:argID", ConverHttprouterToGin(GetDatabasechangeloglock))
	router.PUT("/databasechangeloglock/:argID", ConverHttprouterToGin(UpdateDatabasechangeloglock))
	router.DELETE("/databasechangeloglock/:argID", ConverHttprouterToGin(DeleteDatabasechangeloglock))
}

// GetAllDatabasechangeloglock is a function to get a slice of record(s) from databasechangeloglock table in the keycloak database
// @Summary Get list of Databasechangeloglock
// @Tags Databasechangeloglock
// @Description GetAllDatabasechangeloglock is a handler to get a slice of record(s) from databasechangeloglock table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Databasechangeloglock}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /databasechangeloglock [get]
// http "http://localhost:8080/databasechangeloglock?page=0&pagesize=20" X-Api-User:user123
func GetAllDatabasechangeloglock(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "databasechangeloglock", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllDatabasechangeloglock(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetDatabasechangeloglock is a function to get a single record from the databasechangeloglock table in the keycloak database
// @Summary Get record from table Databasechangeloglock by  argID
// @Tags Databasechangeloglock
// @ID argID
// @Description GetDatabasechangeloglock is a function to get a single record from the databasechangeloglock table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 200 {object} model.Databasechangeloglock
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /databasechangeloglock/{argID} [get]
// http "http://localhost:8080/databasechangeloglock/1" X-Api-User:user123
func GetDatabasechangeloglock(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "databasechangeloglock", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetDatabasechangeloglock(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddDatabasechangeloglock add to add a single record to databasechangeloglock table in the keycloak database
// @Summary Add an record to databasechangeloglock table
// @Description add to add a single record to databasechangeloglock table in the keycloak database
// @Tags Databasechangeloglock
// @Accept  json
// @Produce  json
// @Param Databasechangeloglock body model.Databasechangeloglock true "Add Databasechangeloglock"
// @Success 200 {object} model.Databasechangeloglock
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /databasechangeloglock [post]
// echo '{"id": 93,"locked": false,"lockgranted": "2273-01-15T15:18:15.108612853+08:00","lockedby": "fjLtQSKplJNgtQiIVkCnIfTtG"}' | http POST "http://localhost:8080/databasechangeloglock" X-Api-User:user123
func AddDatabasechangeloglock(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	databasechangeloglock := &model.Databasechangeloglock{}

	if err := readJSON(r, databasechangeloglock); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := databasechangeloglock.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	databasechangeloglock.Prepare()

	if err := databasechangeloglock.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "databasechangeloglock", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	databasechangeloglock, _, err = dao.AddDatabasechangeloglock(ctx, databasechangeloglock)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, databasechangeloglock)
}

// UpdateDatabasechangeloglock Update a single record from databasechangeloglock table in the keycloak database
// @Summary Update an record in table databasechangeloglock
// @Description Update a single record from databasechangeloglock table in the keycloak database
// @Tags Databasechangeloglock
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Param  Databasechangeloglock body model.Databasechangeloglock true "Update Databasechangeloglock record"
// @Success 200 {object} model.Databasechangeloglock
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /databasechangeloglock/{argID} [put]
// echo '{"id": 93,"locked": false,"lockgranted": "2273-01-15T15:18:15.108612853+08:00","lockedby": "fjLtQSKplJNgtQiIVkCnIfTtG"}' | http PUT "http://localhost:8080/databasechangeloglock/1"  X-Api-User:user123
func UpdateDatabasechangeloglock(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	databasechangeloglock := &model.Databasechangeloglock{}
	if err := readJSON(r, databasechangeloglock); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := databasechangeloglock.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	databasechangeloglock.Prepare()

	if err := databasechangeloglock.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "databasechangeloglock", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	databasechangeloglock, _, err = dao.UpdateDatabasechangeloglock(ctx,
		argID,
		databasechangeloglock)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, databasechangeloglock)
}

// DeleteDatabasechangeloglock Delete a single record from databasechangeloglock table in the keycloak database
// @Summary Delete a record from databasechangeloglock
// @Description Delete a single record from databasechangeloglock table in the keycloak database
// @Tags Databasechangeloglock
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 204 {object} model.Databasechangeloglock
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /databasechangeloglock/{argID} [delete]
// http DELETE "http://localhost:8080/databasechangeloglock/1" X-Api-User:user123
func DeleteDatabasechangeloglock(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "databasechangeloglock", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteDatabasechangeloglock(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
