package api

import (
	"net/http"

	"example.com/rest/example/dao"
	"example.com/rest/example/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configFailedJobsRouter(router *httprouter.Router) {
	router.GET("/failedjobs", GetAllFailedJobs)
	router.POST("/failedjobs", AddFailedJobs)
	router.GET("/failedjobs/:argID", GetFailedJobs)
	router.PUT("/failedjobs/:argID", UpdateFailedJobs)
	router.DELETE("/failedjobs/:argID", DeleteFailedJobs)
}

func configGinFailedJobsRouter(router gin.IRoutes) {
	router.GET("/failedjobs", ConverHttprouterToGin(GetAllFailedJobs))
	router.POST("/failedjobs", ConverHttprouterToGin(AddFailedJobs))
	router.GET("/failedjobs/:argID", ConverHttprouterToGin(GetFailedJobs))
	router.PUT("/failedjobs/:argID", ConverHttprouterToGin(UpdateFailedJobs))
	router.DELETE("/failedjobs/:argID", ConverHttprouterToGin(DeleteFailedJobs))
}

// GetAllFailedJobs is a function to get a slice of record(s) from failed_jobs table in the test1 database
// @Summary Get list of FailedJobs
// @Tags FailedJobs
// @Description GetAllFailedJobs is a handler to get a slice of record(s) from failed_jobs table in the test1 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.FailedJobs}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /failedjobs [get]
// http "http://localhost:8080/failedjobs?page=0&pagesize=20" X-Api-User:user123
func GetAllFailedJobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "failed_jobs", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllFailedJobs(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetFailedJobs is a function to get a single record from the failed_jobs table in the test1 database
// @Summary Get record from table FailedJobs by  argID
// @Tags FailedJobs
// @ID argID
// @Description GetFailedJobs is a function to get a single record from the failed_jobs table in the test1 database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.FailedJobs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /failedjobs/{argID} [get]
// http "http://localhost:8080/failedjobs/1" X-Api-User:user123
func GetFailedJobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseUint64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "failed_jobs", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetFailedJobs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddFailedJobs add to add a single record to failed_jobs table in the test1 database
// @Summary Add an record to failed_jobs table
// @Description add to add a single record to failed_jobs table in the test1 database
// @Tags FailedJobs
// @Accept  json
// @Produce  json
// @Param FailedJobs body model.FailedJobs true "Add FailedJobs"
// @Success 200 {object} model.FailedJobs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /failedjobs [post]
// echo '{"id": 10,"uuid": "MYdrQNGchVxCfcqeHaXErhABB","connection": "uDoNQMCYdjyNiLXwkvaGltftm","queue": "ZrVGOnjjDowIZdvUtxQEJyqug","payload": "joHQdNlKjIVnWdbTobRnJthCF","exception": "evlNvHKwiKyAqvlxccZVWyZib","failed_at": "2175-07-21T22:32:28.124048268+08:00"}' | http POST "http://localhost:8080/failedjobs" X-Api-User:user123
func AddFailedJobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	failedjobs := &model.FailedJobs{}

	if err := readJSON(r, failedjobs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := failedjobs.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	failedjobs.Prepare()

	if err := failedjobs.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "failed_jobs", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	failedjobs, _, err = dao.AddFailedJobs(ctx, failedjobs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, failedjobs)
}

// UpdateFailedJobs Update a single record from failed_jobs table in the test1 database
// @Summary Update an record in table failed_jobs
// @Description Update a single record from failed_jobs table in the test1 database
// @Tags FailedJobs
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  FailedJobs body model.FailedJobs true "Update FailedJobs record"
// @Success 200 {object} model.FailedJobs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /failedjobs/{argID} [put]
// echo '{"id": 10,"uuid": "MYdrQNGchVxCfcqeHaXErhABB","connection": "uDoNQMCYdjyNiLXwkvaGltftm","queue": "ZrVGOnjjDowIZdvUtxQEJyqug","payload": "joHQdNlKjIVnWdbTobRnJthCF","exception": "evlNvHKwiKyAqvlxccZVWyZib","failed_at": "2175-07-21T22:32:28.124048268+08:00"}' | http PUT "http://localhost:8080/failedjobs/1"  X-Api-User:user123
func UpdateFailedJobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseUint64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	failedjobs := &model.FailedJobs{}
	if err := readJSON(r, failedjobs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := failedjobs.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	failedjobs.Prepare()

	if err := failedjobs.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "failed_jobs", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	failedjobs, _, err = dao.UpdateFailedJobs(ctx,
		argID,
		failedjobs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, failedjobs)
}

// DeleteFailedJobs Delete a single record from failed_jobs table in the test1 database
// @Summary Delete a record from failed_jobs
// @Description Delete a single record from failed_jobs table in the test1 database
// @Tags FailedJobs
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.FailedJobs
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /failedjobs/{argID} [delete]
// http DELETE "http://localhost:8080/failedjobs/1" X-Api-User:user123
func DeleteFailedJobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseUint64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "failed_jobs", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteFailedJobs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
