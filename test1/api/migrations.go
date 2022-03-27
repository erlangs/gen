package api

import (
	"net/http"

	"example.com/rest/example/dao"
	"example.com/rest/example/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configMigrationsRouter(router *httprouter.Router) {
	router.GET("/migrations", GetAllMigrations)
	router.POST("/migrations", AddMigrations)
	router.GET("/migrations/:argID", GetMigrations)
	router.PUT("/migrations/:argID", UpdateMigrations)
	router.DELETE("/migrations/:argID", DeleteMigrations)
}

func configGinMigrationsRouter(router gin.IRoutes) {
	router.GET("/migrations", ConverHttprouterToGin(GetAllMigrations))
	router.POST("/migrations", ConverHttprouterToGin(AddMigrations))
	router.GET("/migrations/:argID", ConverHttprouterToGin(GetMigrations))
	router.PUT("/migrations/:argID", ConverHttprouterToGin(UpdateMigrations))
	router.DELETE("/migrations/:argID", ConverHttprouterToGin(DeleteMigrations))
}

// GetAllMigrations is a function to get a slice of record(s) from migrations table in the test1 database
// @Summary Get list of Migrations
// @Tags Migrations
// @Description GetAllMigrations is a handler to get a slice of record(s) from migrations table in the test1 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Migrations}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /migrations [get]
// http "http://localhost:8080/migrations?page=0&pagesize=20" X-Api-User:user123
func GetAllMigrations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "migrations", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllMigrations(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetMigrations is a function to get a single record from the migrations table in the test1 database
// @Summary Get record from table Migrations by  argID
// @Tags Migrations
// @ID argID
// @Description GetMigrations is a function to get a single record from the migrations table in the test1 database
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 200 {object} model.Migrations
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /migrations/{argID} [get]
// http "http://localhost:8080/migrations/1" X-Api-User:user123
func GetMigrations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseUint32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "migrations", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetMigrations(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddMigrations add to add a single record to migrations table in the test1 database
// @Summary Add an record to migrations table
// @Description add to add a single record to migrations table in the test1 database
// @Tags Migrations
// @Accept  json
// @Produce  json
// @Param Migrations body model.Migrations true "Add Migrations"
// @Success 200 {object} model.Migrations
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /migrations [post]
// echo '{"id": 95,"migration": "BJGcgVgOFMWccdVNuwhZlkeTs","batch": 82}' | http POST "http://localhost:8080/migrations" X-Api-User:user123
func AddMigrations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	migrations := &model.Migrations{}

	if err := readJSON(r, migrations); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := migrations.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	migrations.Prepare()

	if err := migrations.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "migrations", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	migrations, _, err = dao.AddMigrations(ctx, migrations)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, migrations)
}

// UpdateMigrations Update a single record from migrations table in the test1 database
// @Summary Update an record in table migrations
// @Description Update a single record from migrations table in the test1 database
// @Tags Migrations
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Param  Migrations body model.Migrations true "Update Migrations record"
// @Success 200 {object} model.Migrations
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /migrations/{argID} [put]
// echo '{"id": 95,"migration": "BJGcgVgOFMWccdVNuwhZlkeTs","batch": 82}' | http PUT "http://localhost:8080/migrations/1"  X-Api-User:user123
func UpdateMigrations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseUint32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	migrations := &model.Migrations{}
	if err := readJSON(r, migrations); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := migrations.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	migrations.Prepare()

	if err := migrations.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "migrations", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	migrations, _, err = dao.UpdateMigrations(ctx,
		argID,
		migrations)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, migrations)
}

// DeleteMigrations Delete a single record from migrations table in the test1 database
// @Summary Delete a record from migrations
// @Description Delete a single record from migrations table in the test1 database
// @Tags Migrations
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 204 {object} model.Migrations
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /migrations/{argID} [delete]
// http DELETE "http://localhost:8080/migrations/1" X-Api-User:user123
func DeleteMigrations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseUint32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "migrations", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteMigrations(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
