package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configResourceServerResourceRouter(router *httprouter.Router) {
	router.GET("/resourceserverresource", GetAllResourceServerResource)
	router.POST("/resourceserverresource", AddResourceServerResource)
	router.GET("/resourceserverresource/:argID", GetResourceServerResource)
	router.PUT("/resourceserverresource/:argID", UpdateResourceServerResource)
	router.DELETE("/resourceserverresource/:argID", DeleteResourceServerResource)
}

func configGinResourceServerResourceRouter(router gin.IRoutes) {
	router.GET("/resourceserverresource", ConverHttprouterToGin(GetAllResourceServerResource))
	router.POST("/resourceserverresource", ConverHttprouterToGin(AddResourceServerResource))
	router.GET("/resourceserverresource/:argID", ConverHttprouterToGin(GetResourceServerResource))
	router.PUT("/resourceserverresource/:argID", ConverHttprouterToGin(UpdateResourceServerResource))
	router.DELETE("/resourceserverresource/:argID", ConverHttprouterToGin(DeleteResourceServerResource))
}

// GetAllResourceServerResource is a function to get a slice of record(s) from resource_server_resource table in the keycloak database
// @Summary Get list of ResourceServerResource
// @Tags ResourceServerResource
// @Description GetAllResourceServerResource is a handler to get a slice of record(s) from resource_server_resource table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ResourceServerResource}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceserverresource [get]
// http "http://localhost:8080/resourceserverresource?page=0&pagesize=20" X-Api-User:user123
func GetAllResourceServerResource(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "resource_server_resource", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllResourceServerResource(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetResourceServerResource is a function to get a single record from the resource_server_resource table in the keycloak database
// @Summary Get record from table ResourceServerResource by  argID
// @Tags ResourceServerResource
// @ID argID
// @Description GetResourceServerResource is a function to get a single record from the resource_server_resource table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.ResourceServerResource
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /resourceserverresource/{argID} [get]
// http "http://localhost:8080/resourceserverresource/hello world" X-Api-User:user123
func GetResourceServerResource(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server_resource", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetResourceServerResource(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddResourceServerResource add to add a single record to resource_server_resource table in the keycloak database
// @Summary Add an record to resource_server_resource table
// @Description add to add a single record to resource_server_resource table in the keycloak database
// @Tags ResourceServerResource
// @Accept  json
// @Produce  json
// @Param ResourceServerResource body model.ResourceServerResource true "Add ResourceServerResource"
// @Success 200 {object} model.ResourceServerResource
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceserverresource [post]
// echo '{"id": "WOeeJVLfRdsPUeuGWpRXSFWgn","name": "FsdsvncdNVgbxscIZNjBtyJgL","type": "lwwkrGZBiQTSgiLvMPEbHApTl","icon_uri": "jgGXyOwJntbDLtvwYIeCqWXbh","owner": "uyAQDLjiCwWpJZRTSIsrgsUgk","resource_server_id": "EqOaksmdjWjMYMcLCqIsKdiZG","owner_managed_access": false,"display_name": "mnmYOygSGOVUQtwslDaRpXRsK"}' | http POST "http://localhost:8080/resourceserverresource" X-Api-User:user123
func AddResourceServerResource(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	resourceserverresource := &model.ResourceServerResource{}

	if err := readJSON(r, resourceserverresource); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resourceserverresource.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resourceserverresource.Prepare()

	if err := resourceserverresource.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server_resource", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	resourceserverresource, _, err = dao.AddResourceServerResource(ctx, resourceserverresource)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resourceserverresource)
}

// UpdateResourceServerResource Update a single record from resource_server_resource table in the keycloak database
// @Summary Update an record in table resource_server_resource
// @Description Update a single record from resource_server_resource table in the keycloak database
// @Tags ResourceServerResource
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  ResourceServerResource body model.ResourceServerResource true "Update ResourceServerResource record"
// @Success 200 {object} model.ResourceServerResource
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceserverresource/{argID} [put]
// echo '{"id": "WOeeJVLfRdsPUeuGWpRXSFWgn","name": "FsdsvncdNVgbxscIZNjBtyJgL","type": "lwwkrGZBiQTSgiLvMPEbHApTl","icon_uri": "jgGXyOwJntbDLtvwYIeCqWXbh","owner": "uyAQDLjiCwWpJZRTSIsrgsUgk","resource_server_id": "EqOaksmdjWjMYMcLCqIsKdiZG","owner_managed_access": false,"display_name": "mnmYOygSGOVUQtwslDaRpXRsK"}' | http PUT "http://localhost:8080/resourceserverresource/hello world"  X-Api-User:user123
func UpdateResourceServerResource(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resourceserverresource := &model.ResourceServerResource{}
	if err := readJSON(r, resourceserverresource); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resourceserverresource.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resourceserverresource.Prepare()

	if err := resourceserverresource.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server_resource", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resourceserverresource, _, err = dao.UpdateResourceServerResource(ctx,
		argID,
		resourceserverresource)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resourceserverresource)
}

// DeleteResourceServerResource Delete a single record from resource_server_resource table in the keycloak database
// @Summary Delete a record from resource_server_resource
// @Description Delete a single record from resource_server_resource table in the keycloak database
// @Tags ResourceServerResource
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.ResourceServerResource
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /resourceserverresource/{argID} [delete]
// http DELETE "http://localhost:8080/resourceserverresource/hello world" X-Api-User:user123
func DeleteResourceServerResource(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server_resource", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteResourceServerResource(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
