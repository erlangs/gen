package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configDatabasechangelogRouter(router *httprouter.Router) {
	router.GET("/databasechangelog", GetAllDatabasechangelog)
	router.POST("/databasechangelog", AddDatabasechangelog)
	router.GET("/databasechangelog/:argID", GetDatabasechangelog)
	router.PUT("/databasechangelog/:argID", UpdateDatabasechangelog)
	router.DELETE("/databasechangelog/:argID", DeleteDatabasechangelog)
}

func configGinDatabasechangelogRouter(router gin.IRoutes) {
	router.GET("/databasechangelog", ConverHttprouterToGin(GetAllDatabasechangelog))
	router.POST("/databasechangelog", ConverHttprouterToGin(AddDatabasechangelog))
	router.GET("/databasechangelog/:argID", ConverHttprouterToGin(GetDatabasechangelog))
	router.PUT("/databasechangelog/:argID", ConverHttprouterToGin(UpdateDatabasechangelog))
	router.DELETE("/databasechangelog/:argID", ConverHttprouterToGin(DeleteDatabasechangelog))
}

// GetAllDatabasechangelog is a function to get a slice of record(s) from databasechangelog table in the keycloak database
// @Summary Get list of Databasechangelog
// @Tags Databasechangelog
// @Description GetAllDatabasechangelog is a handler to get a slice of record(s) from databasechangelog table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Databasechangelog}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /databasechangelog [get]
// http "http://localhost:8080/databasechangelog?page=0&pagesize=20" X-Api-User:user123
func GetAllDatabasechangelog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "databasechangelog", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllDatabasechangelog(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetDatabasechangelog is a function to get a single record from the databasechangelog table in the keycloak database
// @Summary Get record from table Databasechangelog by  argID
// @Tags Databasechangelog
// @ID argID
// @Description GetDatabasechangelog is a function to get a single record from the databasechangelog table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.Databasechangelog
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /databasechangelog/{argID} [get]
// http "http://localhost:8080/databasechangelog/hello world" X-Api-User:user123
func GetDatabasechangelog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "databasechangelog", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetDatabasechangelog(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddDatabasechangelog add to add a single record to databasechangelog table in the keycloak database
// @Summary Add an record to databasechangelog table
// @Description add to add a single record to databasechangelog table in the keycloak database
// @Tags Databasechangelog
// @Accept  json
// @Produce  json
// @Param Databasechangelog body model.Databasechangelog true "Add Databasechangelog"
// @Success 200 {object} model.Databasechangelog
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /databasechangelog [post]
// echo '{"id": "UyjiSgAFCLeIiQyFPSqSMgBgB","author": "mHFVFMHviNUUJhOncxXCCdvLU","filename": "qTYYEVIBwmoXSZZmdLoUIwxrq","dateexecuted": "2276-02-06T13:39:13.307404738+08:00","orderexecuted": 76,"exectype": "YZikIFBfNljssxLMOfVGPBhfm","md_5_sum": "UyMbYuwZDQWJpMfySaAUnBAmt","description": "xJDMytHOLFTgtMHhyZPBTZhEY","comments": "fVJViHkoZQqZgnfuukyjPncJo","tag": "HtZgpYfehLSpXTOBytHnAAZbL","liquibase": "yvtofkYJaSMsXZLeECryrkgGU","contexts": "UOKyEqjFfcRFceFELgkvZCnLc","labels": "UaGclABowpcsJfgGtxVTWprDE","deployment_id": "eOwBnxpXVZCXjCYiEkdFnPoQi"}' | http POST "http://localhost:8080/databasechangelog" X-Api-User:user123
func AddDatabasechangelog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	databasechangelog := &model.Databasechangelog{}

	if err := readJSON(r, databasechangelog); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := databasechangelog.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	databasechangelog.Prepare()

	if err := databasechangelog.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "databasechangelog", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	databasechangelog, _, err = dao.AddDatabasechangelog(ctx, databasechangelog)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, databasechangelog)
}

// UpdateDatabasechangelog Update a single record from databasechangelog table in the keycloak database
// @Summary Update an record in table databasechangelog
// @Description Update a single record from databasechangelog table in the keycloak database
// @Tags Databasechangelog
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  Databasechangelog body model.Databasechangelog true "Update Databasechangelog record"
// @Success 200 {object} model.Databasechangelog
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /databasechangelog/{argID} [put]
// echo '{"id": "UyjiSgAFCLeIiQyFPSqSMgBgB","author": "mHFVFMHviNUUJhOncxXCCdvLU","filename": "qTYYEVIBwmoXSZZmdLoUIwxrq","dateexecuted": "2276-02-06T13:39:13.307404738+08:00","orderexecuted": 76,"exectype": "YZikIFBfNljssxLMOfVGPBhfm","md_5_sum": "UyMbYuwZDQWJpMfySaAUnBAmt","description": "xJDMytHOLFTgtMHhyZPBTZhEY","comments": "fVJViHkoZQqZgnfuukyjPncJo","tag": "HtZgpYfehLSpXTOBytHnAAZbL","liquibase": "yvtofkYJaSMsXZLeECryrkgGU","contexts": "UOKyEqjFfcRFceFELgkvZCnLc","labels": "UaGclABowpcsJfgGtxVTWprDE","deployment_id": "eOwBnxpXVZCXjCYiEkdFnPoQi"}' | http PUT "http://localhost:8080/databasechangelog/hello world"  X-Api-User:user123
func UpdateDatabasechangelog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	databasechangelog := &model.Databasechangelog{}
	if err := readJSON(r, databasechangelog); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := databasechangelog.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	databasechangelog.Prepare()

	if err := databasechangelog.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "databasechangelog", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	databasechangelog, _, err = dao.UpdateDatabasechangelog(ctx,
		argID,
		databasechangelog)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, databasechangelog)
}

// DeleteDatabasechangelog Delete a single record from databasechangelog table in the keycloak database
// @Summary Delete a record from databasechangelog
// @Description Delete a single record from databasechangelog table in the keycloak database
// @Tags Databasechangelog
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.Databasechangelog
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /databasechangelog/{argID} [delete]
// http DELETE "http://localhost:8080/databasechangelog/hello world" X-Api-User:user123
func DeleteDatabasechangelog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "databasechangelog", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteDatabasechangelog(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
