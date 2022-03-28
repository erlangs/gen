package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configMigrationModelRouter(router *httprouter.Router) {
	router.GET("/migrationmodel", GetAllMigrationModel)
	router.POST("/migrationmodel", AddMigrationModel)
	router.GET("/migrationmodel/:argID", GetMigrationModel)
	router.PUT("/migrationmodel/:argID", UpdateMigrationModel)
	router.DELETE("/migrationmodel/:argID", DeleteMigrationModel)
}

func configGinMigrationModelRouter(router gin.IRoutes) {
	router.GET("/migrationmodel", ConverHttprouterToGin(GetAllMigrationModel))
	router.POST("/migrationmodel", ConverHttprouterToGin(AddMigrationModel))
	router.GET("/migrationmodel/:argID", ConverHttprouterToGin(GetMigrationModel))
	router.PUT("/migrationmodel/:argID", ConverHttprouterToGin(UpdateMigrationModel))
	router.DELETE("/migrationmodel/:argID", ConverHttprouterToGin(DeleteMigrationModel))
}

// GetAllMigrationModel is a function to get a slice of record(s) from migration_model table in the keycloak database
// @Summary Get list of MigrationModel
// @Tags MigrationModel
// @Description GetAllMigrationModel is a handler to get a slice of record(s) from migration_model table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.MigrationModel}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /migrationmodel [get]
// http "http://localhost:8080/migrationmodel?page=0&pagesize=20" X-Api-User:user123
func GetAllMigrationModel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "migration_model", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllMigrationModel(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetMigrationModel is a function to get a single record from the migration_model table in the keycloak database
// @Summary Get record from table MigrationModel by  argID
// @Tags MigrationModel
// @ID argID
// @Description GetMigrationModel is a function to get a single record from the migration_model table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.MigrationModel
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /migrationmodel/{argID} [get]
// http "http://localhost:8080/migrationmodel/hello world" X-Api-User:user123
func GetMigrationModel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "migration_model", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetMigrationModel(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddMigrationModel add to add a single record to migration_model table in the keycloak database
// @Summary Add an record to migration_model table
// @Description add to add a single record to migration_model table in the keycloak database
// @Tags MigrationModel
// @Accept  json
// @Produce  json
// @Param MigrationModel body model.MigrationModel true "Add MigrationModel"
// @Success 200 {object} model.MigrationModel
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /migrationmodel [post]
// echo '{"id": "CClsVFinNxQYwhcjMEbUZaeWF","version": "LSdYTmaoEjjrDbcAIsFBhwVoT","update_time": 49}' | http POST "http://localhost:8080/migrationmodel" X-Api-User:user123
func AddMigrationModel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	migrationmodel := &model.MigrationModel{}

	if err := readJSON(r, migrationmodel); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := migrationmodel.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	migrationmodel.Prepare()

	if err := migrationmodel.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "migration_model", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	migrationmodel, _, err = dao.AddMigrationModel(ctx, migrationmodel)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, migrationmodel)
}

// UpdateMigrationModel Update a single record from migration_model table in the keycloak database
// @Summary Update an record in table migration_model
// @Description Update a single record from migration_model table in the keycloak database
// @Tags MigrationModel
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  MigrationModel body model.MigrationModel true "Update MigrationModel record"
// @Success 200 {object} model.MigrationModel
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /migrationmodel/{argID} [put]
// echo '{"id": "CClsVFinNxQYwhcjMEbUZaeWF","version": "LSdYTmaoEjjrDbcAIsFBhwVoT","update_time": 49}' | http PUT "http://localhost:8080/migrationmodel/hello world"  X-Api-User:user123
func UpdateMigrationModel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	migrationmodel := &model.MigrationModel{}
	if err := readJSON(r, migrationmodel); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := migrationmodel.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	migrationmodel.Prepare()

	if err := migrationmodel.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "migration_model", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	migrationmodel, _, err = dao.UpdateMigrationModel(ctx,
		argID,
		migrationmodel)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, migrationmodel)
}

// DeleteMigrationModel Delete a single record from migration_model table in the keycloak database
// @Summary Delete a record from migration_model
// @Description Delete a single record from migration_model table in the keycloak database
// @Tags MigrationModel
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.MigrationModel
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /migrationmodel/{argID} [delete]
// http DELETE "http://localhost:8080/migrationmodel/hello world" X-Api-User:user123
func DeleteMigrationModel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "migration_model", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteMigrationModel(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
