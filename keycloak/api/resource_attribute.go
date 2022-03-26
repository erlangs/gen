package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configResourceAttributeRouter(router *httprouter.Router) {
	router.GET("/resourceattribute", GetAllResourceAttribute)
	router.POST("/resourceattribute", AddResourceAttribute)
	router.GET("/resourceattribute/:argID", GetResourceAttribute)
	router.PUT("/resourceattribute/:argID", UpdateResourceAttribute)
	router.DELETE("/resourceattribute/:argID", DeleteResourceAttribute)
}

func configGinResourceAttributeRouter(router gin.IRoutes) {
	router.GET("/resourceattribute", ConverHttprouterToGin(GetAllResourceAttribute))
	router.POST("/resourceattribute", ConverHttprouterToGin(AddResourceAttribute))
	router.GET("/resourceattribute/:argID", ConverHttprouterToGin(GetResourceAttribute))
	router.PUT("/resourceattribute/:argID", ConverHttprouterToGin(UpdateResourceAttribute))
	router.DELETE("/resourceattribute/:argID", ConverHttprouterToGin(DeleteResourceAttribute))
}

// GetAllResourceAttribute is a function to get a slice of record(s) from resource_attribute table in the keycloak database
// @Summary Get list of ResourceAttribute
// @Tags ResourceAttribute
// @Description GetAllResourceAttribute is a handler to get a slice of record(s) from resource_attribute table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ResourceAttribute}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceattribute [get]
// http "http://localhost:8080/resourceattribute?page=0&pagesize=20" X-Api-User:user123
func GetAllResourceAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "resource_attribute", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllResourceAttribute(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetResourceAttribute is a function to get a single record from the resource_attribute table in the keycloak database
// @Summary Get record from table ResourceAttribute by  argID
// @Tags ResourceAttribute
// @ID argID
// @Description GetResourceAttribute is a function to get a single record from the resource_attribute table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.ResourceAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /resourceattribute/{argID} [get]
// http "http://localhost:8080/resourceattribute/hello world" X-Api-User:user123
func GetResourceAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_attribute", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetResourceAttribute(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddResourceAttribute add to add a single record to resource_attribute table in the keycloak database
// @Summary Add an record to resource_attribute table
// @Description add to add a single record to resource_attribute table in the keycloak database
// @Tags ResourceAttribute
// @Accept  json
// @Produce  json
// @Param ResourceAttribute body model.ResourceAttribute true "Add ResourceAttribute"
// @Success 200 {object} model.ResourceAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceattribute [post]
// echo '{"id": "hAcWmPqiQZWfgqxqMqyucrClS","name": "LbAPCImeIfGeaAbaEpNHLAYiV","value": "ZNQrnaSmFanxUogseansuaPjm","resource_id": "pyxsUbXSxsROdLICxApcDSWRO"}' | http POST "http://localhost:8080/resourceattribute" X-Api-User:user123
func AddResourceAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	resourceattribute := &model.ResourceAttribute{}

	if err := readJSON(r, resourceattribute); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resourceattribute.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resourceattribute.Prepare()

	if err := resourceattribute.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_attribute", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	resourceattribute, _, err = dao.AddResourceAttribute(ctx, resourceattribute)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resourceattribute)
}

// UpdateResourceAttribute Update a single record from resource_attribute table in the keycloak database
// @Summary Update an record in table resource_attribute
// @Description Update a single record from resource_attribute table in the keycloak database
// @Tags ResourceAttribute
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  ResourceAttribute body model.ResourceAttribute true "Update ResourceAttribute record"
// @Success 200 {object} model.ResourceAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceattribute/{argID} [put]
// echo '{"id": "hAcWmPqiQZWfgqxqMqyucrClS","name": "LbAPCImeIfGeaAbaEpNHLAYiV","value": "ZNQrnaSmFanxUogseansuaPjm","resource_id": "pyxsUbXSxsROdLICxApcDSWRO"}' | http PUT "http://localhost:8080/resourceattribute/hello world"  X-Api-User:user123
func UpdateResourceAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resourceattribute := &model.ResourceAttribute{}
	if err := readJSON(r, resourceattribute); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resourceattribute.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resourceattribute.Prepare()

	if err := resourceattribute.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_attribute", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resourceattribute, _, err = dao.UpdateResourceAttribute(ctx,
		argID,
		resourceattribute)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resourceattribute)
}

// DeleteResourceAttribute Delete a single record from resource_attribute table in the keycloak database
// @Summary Delete a record from resource_attribute
// @Description Delete a single record from resource_attribute table in the keycloak database
// @Tags ResourceAttribute
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.ResourceAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /resourceattribute/{argID} [delete]
// http DELETE "http://localhost:8080/resourceattribute/hello world" X-Api-User:user123
func DeleteResourceAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_attribute", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteResourceAttribute(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
