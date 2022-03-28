package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configResourceServerRouter(router *httprouter.Router) {
	router.GET("/resourceserver", GetAllResourceServer)
	router.POST("/resourceserver", AddResourceServer)
	router.GET("/resourceserver/:argID", GetResourceServer)
	router.PUT("/resourceserver/:argID", UpdateResourceServer)
	router.DELETE("/resourceserver/:argID", DeleteResourceServer)
}

func configGinResourceServerRouter(router gin.IRoutes) {
	router.GET("/resourceserver", ConverHttprouterToGin(GetAllResourceServer))
	router.POST("/resourceserver", ConverHttprouterToGin(AddResourceServer))
	router.GET("/resourceserver/:argID", ConverHttprouterToGin(GetResourceServer))
	router.PUT("/resourceserver/:argID", ConverHttprouterToGin(UpdateResourceServer))
	router.DELETE("/resourceserver/:argID", ConverHttprouterToGin(DeleteResourceServer))
}

// GetAllResourceServer is a function to get a slice of record(s) from resource_server table in the keycloak database
// @Summary Get list of ResourceServer
// @Tags ResourceServer
// @Description GetAllResourceServer is a handler to get a slice of record(s) from resource_server table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ResourceServer}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceserver [get]
// http "http://localhost:8080/resourceserver?page=0&pagesize=20" X-Api-User:user123
func GetAllResourceServer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "resource_server", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllResourceServer(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetResourceServer is a function to get a single record from the resource_server table in the keycloak database
// @Summary Get record from table ResourceServer by  argID
// @Tags ResourceServer
// @ID argID
// @Description GetResourceServer is a function to get a single record from the resource_server table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.ResourceServer
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /resourceserver/{argID} [get]
// http "http://localhost:8080/resourceserver/hello world" X-Api-User:user123
func GetResourceServer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetResourceServer(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddResourceServer add to add a single record to resource_server table in the keycloak database
// @Summary Add an record to resource_server table
// @Description add to add a single record to resource_server table in the keycloak database
// @Tags ResourceServer
// @Accept  json
// @Produce  json
// @Param ResourceServer body model.ResourceServer true "Add ResourceServer"
// @Success 200 {object} model.ResourceServer
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceserver [post]
// echo '{"id": "UimvTunbmGsItxTFKPwHAylQa","allow_rs_remote_mgmt": true,"policy_enforce_mode": "lHQqeDwZLehPZMDvlirLfkHxB","decision_strategy": 63}' | http POST "http://localhost:8080/resourceserver" X-Api-User:user123
func AddResourceServer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	resourceserver := &model.ResourceServer{}

	if err := readJSON(r, resourceserver); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resourceserver.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resourceserver.Prepare()

	if err := resourceserver.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	resourceserver, _, err = dao.AddResourceServer(ctx, resourceserver)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resourceserver)
}

// UpdateResourceServer Update a single record from resource_server table in the keycloak database
// @Summary Update an record in table resource_server
// @Description Update a single record from resource_server table in the keycloak database
// @Tags ResourceServer
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  ResourceServer body model.ResourceServer true "Update ResourceServer record"
// @Success 200 {object} model.ResourceServer
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceserver/{argID} [put]
// echo '{"id": "UimvTunbmGsItxTFKPwHAylQa","allow_rs_remote_mgmt": true,"policy_enforce_mode": "lHQqeDwZLehPZMDvlirLfkHxB","decision_strategy": 63}' | http PUT "http://localhost:8080/resourceserver/hello world"  X-Api-User:user123
func UpdateResourceServer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resourceserver := &model.ResourceServer{}
	if err := readJSON(r, resourceserver); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resourceserver.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resourceserver.Prepare()

	if err := resourceserver.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resourceserver, _, err = dao.UpdateResourceServer(ctx,
		argID,
		resourceserver)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resourceserver)
}

// DeleteResourceServer Delete a single record from resource_server table in the keycloak database
// @Summary Delete a record from resource_server
// @Description Delete a single record from resource_server table in the keycloak database
// @Tags ResourceServer
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.ResourceServer
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /resourceserver/{argID} [delete]
// http DELETE "http://localhost:8080/resourceserver/hello world" X-Api-User:user123
func DeleteResourceServer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteResourceServer(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
