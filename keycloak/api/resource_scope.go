package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configResourceScopeRouter(router *httprouter.Router) {
	router.GET("/resourcescope", GetAllResourceScope)
	router.POST("/resourcescope", AddResourceScope)
	router.GET("/resourcescope/:argResourceID/:argScopeID", GetResourceScope)
	router.PUT("/resourcescope/:argResourceID/:argScopeID", UpdateResourceScope)
	router.DELETE("/resourcescope/:argResourceID/:argScopeID", DeleteResourceScope)
}

func configGinResourceScopeRouter(router gin.IRoutes) {
	router.GET("/resourcescope", ConverHttprouterToGin(GetAllResourceScope))
	router.POST("/resourcescope", ConverHttprouterToGin(AddResourceScope))
	router.GET("/resourcescope/:argResourceID/:argScopeID", ConverHttprouterToGin(GetResourceScope))
	router.PUT("/resourcescope/:argResourceID/:argScopeID", ConverHttprouterToGin(UpdateResourceScope))
	router.DELETE("/resourcescope/:argResourceID/:argScopeID", ConverHttprouterToGin(DeleteResourceScope))
}

// GetAllResourceScope is a function to get a slice of record(s) from resource_scope table in the keycloak database
// @Summary Get list of ResourceScope
// @Tags ResourceScope
// @Description GetAllResourceScope is a handler to get a slice of record(s) from resource_scope table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ResourceScope}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourcescope [get]
// http "http://localhost:8080/resourcescope?page=0&pagesize=20" X-Api-User:user123
func GetAllResourceScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "resource_scope", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllResourceScope(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetResourceScope is a function to get a single record from the resource_scope table in the keycloak database
// @Summary Get record from table ResourceScope by  argResourceID  argScopeID
// @Tags ResourceScope
// @ID argResourceID
// @ID argScopeID
// @Description GetResourceScope is a function to get a single record from the resource_scope table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argResourceID path string true "resource_id"
// @Param  argScopeID path string true "scope_id"
// @Success 200 {object} model.ResourceScope
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /resourcescope/{argResourceID}/{argScopeID} [get]
// http "http://localhost:8080/resourcescope/hello world/hello world" X-Api-User:user123
func GetResourceScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argResourceID, err := parseString(ps, "argResourceID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_scope", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetResourceScope(ctx, argResourceID, argScopeID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddResourceScope add to add a single record to resource_scope table in the keycloak database
// @Summary Add an record to resource_scope table
// @Description add to add a single record to resource_scope table in the keycloak database
// @Tags ResourceScope
// @Accept  json
// @Produce  json
// @Param ResourceScope body model.ResourceScope true "Add ResourceScope"
// @Success 200 {object} model.ResourceScope
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourcescope [post]
// echo '{"resource_id": "aGpGQOWQHlqSxTyRjNMbPTwsH","scope_id": "hUoIgyqHcAAOtvlwKvlmbqqGM"}' | http POST "http://localhost:8080/resourcescope" X-Api-User:user123
func AddResourceScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	resourcescope := &model.ResourceScope{}

	if err := readJSON(r, resourcescope); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resourcescope.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resourcescope.Prepare()

	if err := resourcescope.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_scope", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	resourcescope, _, err = dao.AddResourceScope(ctx, resourcescope)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resourcescope)
}

// UpdateResourceScope Update a single record from resource_scope table in the keycloak database
// @Summary Update an record in table resource_scope
// @Description Update a single record from resource_scope table in the keycloak database
// @Tags ResourceScope
// @Accept  json
// @Produce  json
// @Param  argResourceID path string true "resource_id"// @Param  argScopeID path string true "scope_id"
// @Param  ResourceScope body model.ResourceScope true "Update ResourceScope record"
// @Success 200 {object} model.ResourceScope
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourcescope/{argResourceID}/{argScopeID} [put]
// echo '{"resource_id": "aGpGQOWQHlqSxTyRjNMbPTwsH","scope_id": "hUoIgyqHcAAOtvlwKvlmbqqGM"}' | http PUT "http://localhost:8080/resourcescope/hello world/hello world"  X-Api-User:user123
func UpdateResourceScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argResourceID, err := parseString(ps, "argResourceID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resourcescope := &model.ResourceScope{}
	if err := readJSON(r, resourcescope); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resourcescope.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resourcescope.Prepare()

	if err := resourcescope.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_scope", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resourcescope, _, err = dao.UpdateResourceScope(ctx,
		argResourceID, argScopeID,
		resourcescope)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resourcescope)
}

// DeleteResourceScope Delete a single record from resource_scope table in the keycloak database
// @Summary Delete a record from resource_scope
// @Description Delete a single record from resource_scope table in the keycloak database
// @Tags ResourceScope
// @Accept  json
// @Produce  json
// @Param  argResourceID path string true "resource_id"// @Param  argScopeID path string true "scope_id"
// @Success 204 {object} model.ResourceScope
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /resourcescope/{argResourceID}/{argScopeID} [delete]
// http DELETE "http://localhost:8080/resourcescope/hello world/hello world" X-Api-User:user123
func DeleteResourceScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argResourceID, err := parseString(ps, "argResourceID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_scope", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteResourceScope(ctx, argResourceID, argScopeID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
