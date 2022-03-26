package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configKeycloakRoleRouter(router *httprouter.Router) {
	router.GET("/keycloakrole", GetAllKeycloakRole)
	router.POST("/keycloakrole", AddKeycloakRole)
	router.GET("/keycloakrole/:argID", GetKeycloakRole)
	router.PUT("/keycloakrole/:argID", UpdateKeycloakRole)
	router.DELETE("/keycloakrole/:argID", DeleteKeycloakRole)
}

func configGinKeycloakRoleRouter(router gin.IRoutes) {
	router.GET("/keycloakrole", ConverHttprouterToGin(GetAllKeycloakRole))
	router.POST("/keycloakrole", ConverHttprouterToGin(AddKeycloakRole))
	router.GET("/keycloakrole/:argID", ConverHttprouterToGin(GetKeycloakRole))
	router.PUT("/keycloakrole/:argID", ConverHttprouterToGin(UpdateKeycloakRole))
	router.DELETE("/keycloakrole/:argID", ConverHttprouterToGin(DeleteKeycloakRole))
}

// GetAllKeycloakRole is a function to get a slice of record(s) from keycloak_role table in the keycloak database
// @Summary Get list of KeycloakRole
// @Tags KeycloakRole
// @Description GetAllKeycloakRole is a handler to get a slice of record(s) from keycloak_role table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.KeycloakRole}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /keycloakrole [get]
// http "http://localhost:8080/keycloakrole?page=0&pagesize=20" X-Api-User:user123
func GetAllKeycloakRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "keycloak_role", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllKeycloakRole(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetKeycloakRole is a function to get a single record from the keycloak_role table in the keycloak database
// @Summary Get record from table KeycloakRole by  argID
// @Tags KeycloakRole
// @ID argID
// @Description GetKeycloakRole is a function to get a single record from the keycloak_role table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.KeycloakRole
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /keycloakrole/{argID} [get]
// http "http://localhost:8080/keycloakrole/hello world" X-Api-User:user123
func GetKeycloakRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "keycloak_role", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetKeycloakRole(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddKeycloakRole add to add a single record to keycloak_role table in the keycloak database
// @Summary Add an record to keycloak_role table
// @Description add to add a single record to keycloak_role table in the keycloak database
// @Tags KeycloakRole
// @Accept  json
// @Produce  json
// @Param KeycloakRole body model.KeycloakRole true "Add KeycloakRole"
// @Success 200 {object} model.KeycloakRole
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /keycloakrole [post]
// echo '{"id": "biAiEChYdRQPpowvIyFAgRjDx","client_realm_constraint": "qEWDaMwkitIADFjSQvEfwxGVs","client_role": true,"description": "bfydNrobUjCUyXMFEBQQUmZYK","name": "HxBnnwTrXxZowbrdwdPDcoGQJ","realm_id": "BtPngPKabjUrPvBymHPqNvIgv","client": "jZeAEkCLjOJYrjcQTeFDLGZYv","realm": "axvdPNYeuUAScgZTGICgOhEvP"}' | http POST "http://localhost:8080/keycloakrole" X-Api-User:user123
func AddKeycloakRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	keycloakrole := &model.KeycloakRole{}

	if err := readJSON(r, keycloakrole); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := keycloakrole.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	keycloakrole.Prepare()

	if err := keycloakrole.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "keycloak_role", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	keycloakrole, _, err = dao.AddKeycloakRole(ctx, keycloakrole)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, keycloakrole)
}

// UpdateKeycloakRole Update a single record from keycloak_role table in the keycloak database
// @Summary Update an record in table keycloak_role
// @Description Update a single record from keycloak_role table in the keycloak database
// @Tags KeycloakRole
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  KeycloakRole body model.KeycloakRole true "Update KeycloakRole record"
// @Success 200 {object} model.KeycloakRole
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /keycloakrole/{argID} [put]
// echo '{"id": "biAiEChYdRQPpowvIyFAgRjDx","client_realm_constraint": "qEWDaMwkitIADFjSQvEfwxGVs","client_role": true,"description": "bfydNrobUjCUyXMFEBQQUmZYK","name": "HxBnnwTrXxZowbrdwdPDcoGQJ","realm_id": "BtPngPKabjUrPvBymHPqNvIgv","client": "jZeAEkCLjOJYrjcQTeFDLGZYv","realm": "axvdPNYeuUAScgZTGICgOhEvP"}' | http PUT "http://localhost:8080/keycloakrole/hello world"  X-Api-User:user123
func UpdateKeycloakRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	keycloakrole := &model.KeycloakRole{}
	if err := readJSON(r, keycloakrole); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := keycloakrole.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	keycloakrole.Prepare()

	if err := keycloakrole.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "keycloak_role", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	keycloakrole, _, err = dao.UpdateKeycloakRole(ctx,
		argID,
		keycloakrole)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, keycloakrole)
}

// DeleteKeycloakRole Delete a single record from keycloak_role table in the keycloak database
// @Summary Delete a record from keycloak_role
// @Description Delete a single record from keycloak_role table in the keycloak database
// @Tags KeycloakRole
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.KeycloakRole
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /keycloakrole/{argID} [delete]
// http DELETE "http://localhost:8080/keycloakrole/hello world" X-Api-User:user123
func DeleteKeycloakRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "keycloak_role", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteKeycloakRole(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
