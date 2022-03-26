package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configKeycloakGroupRouter(router *httprouter.Router) {
	router.GET("/keycloakgroup", GetAllKeycloakGroup)
	router.POST("/keycloakgroup", AddKeycloakGroup)
	router.GET("/keycloakgroup/:argID", GetKeycloakGroup)
	router.PUT("/keycloakgroup/:argID", UpdateKeycloakGroup)
	router.DELETE("/keycloakgroup/:argID", DeleteKeycloakGroup)
}

func configGinKeycloakGroupRouter(router gin.IRoutes) {
	router.GET("/keycloakgroup", ConverHttprouterToGin(GetAllKeycloakGroup))
	router.POST("/keycloakgroup", ConverHttprouterToGin(AddKeycloakGroup))
	router.GET("/keycloakgroup/:argID", ConverHttprouterToGin(GetKeycloakGroup))
	router.PUT("/keycloakgroup/:argID", ConverHttprouterToGin(UpdateKeycloakGroup))
	router.DELETE("/keycloakgroup/:argID", ConverHttprouterToGin(DeleteKeycloakGroup))
}

// GetAllKeycloakGroup is a function to get a slice of record(s) from keycloak_group table in the keycloak database
// @Summary Get list of KeycloakGroup
// @Tags KeycloakGroup
// @Description GetAllKeycloakGroup is a handler to get a slice of record(s) from keycloak_group table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.KeycloakGroup}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /keycloakgroup [get]
// http "http://localhost:8080/keycloakgroup?page=0&pagesize=20" X-Api-User:user123
func GetAllKeycloakGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "keycloak_group", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllKeycloakGroup(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetKeycloakGroup is a function to get a single record from the keycloak_group table in the keycloak database
// @Summary Get record from table KeycloakGroup by  argID
// @Tags KeycloakGroup
// @ID argID
// @Description GetKeycloakGroup is a function to get a single record from the keycloak_group table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.KeycloakGroup
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /keycloakgroup/{argID} [get]
// http "http://localhost:8080/keycloakgroup/hello world" X-Api-User:user123
func GetKeycloakGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "keycloak_group", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetKeycloakGroup(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddKeycloakGroup add to add a single record to keycloak_group table in the keycloak database
// @Summary Add an record to keycloak_group table
// @Description add to add a single record to keycloak_group table in the keycloak database
// @Tags KeycloakGroup
// @Accept  json
// @Produce  json
// @Param KeycloakGroup body model.KeycloakGroup true "Add KeycloakGroup"
// @Success 200 {object} model.KeycloakGroup
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /keycloakgroup [post]
// echo '{"id": "FbtuYDcmjogEnVLkeZlOcrwpw","name": "BSarhqRBvfZChfWSHbSeoIBok","parent_group": "lmCvEVkbUPpiOgiIXrqjciXlf","realm_id": "OoCpBBcuoeSpgqlShTiUQLWOA"}' | http POST "http://localhost:8080/keycloakgroup" X-Api-User:user123
func AddKeycloakGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	keycloakgroup := &model.KeycloakGroup{}

	if err := readJSON(r, keycloakgroup); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := keycloakgroup.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	keycloakgroup.Prepare()

	if err := keycloakgroup.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "keycloak_group", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	keycloakgroup, _, err = dao.AddKeycloakGroup(ctx, keycloakgroup)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, keycloakgroup)
}

// UpdateKeycloakGroup Update a single record from keycloak_group table in the keycloak database
// @Summary Update an record in table keycloak_group
// @Description Update a single record from keycloak_group table in the keycloak database
// @Tags KeycloakGroup
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  KeycloakGroup body model.KeycloakGroup true "Update KeycloakGroup record"
// @Success 200 {object} model.KeycloakGroup
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /keycloakgroup/{argID} [put]
// echo '{"id": "FbtuYDcmjogEnVLkeZlOcrwpw","name": "BSarhqRBvfZChfWSHbSeoIBok","parent_group": "lmCvEVkbUPpiOgiIXrqjciXlf","realm_id": "OoCpBBcuoeSpgqlShTiUQLWOA"}' | http PUT "http://localhost:8080/keycloakgroup/hello world"  X-Api-User:user123
func UpdateKeycloakGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	keycloakgroup := &model.KeycloakGroup{}
	if err := readJSON(r, keycloakgroup); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := keycloakgroup.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	keycloakgroup.Prepare()

	if err := keycloakgroup.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "keycloak_group", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	keycloakgroup, _, err = dao.UpdateKeycloakGroup(ctx,
		argID,
		keycloakgroup)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, keycloakgroup)
}

// DeleteKeycloakGroup Delete a single record from keycloak_group table in the keycloak database
// @Summary Delete a record from keycloak_group
// @Description Delete a single record from keycloak_group table in the keycloak database
// @Tags KeycloakGroup
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.KeycloakGroup
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /keycloakgroup/{argID} [delete]
// http DELETE "http://localhost:8080/keycloakgroup/hello world" X-Api-User:user123
func DeleteKeycloakGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "keycloak_group", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteKeycloakGroup(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
