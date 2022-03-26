package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configResourceServerScopeRouter(router *httprouter.Router) {
	router.GET("/resourceserverscope", GetAllResourceServerScope)
	router.POST("/resourceserverscope", AddResourceServerScope)
	router.GET("/resourceserverscope/:argID", GetResourceServerScope)
	router.PUT("/resourceserverscope/:argID", UpdateResourceServerScope)
	router.DELETE("/resourceserverscope/:argID", DeleteResourceServerScope)
}

func configGinResourceServerScopeRouter(router gin.IRoutes) {
	router.GET("/resourceserverscope", ConverHttprouterToGin(GetAllResourceServerScope))
	router.POST("/resourceserverscope", ConverHttprouterToGin(AddResourceServerScope))
	router.GET("/resourceserverscope/:argID", ConverHttprouterToGin(GetResourceServerScope))
	router.PUT("/resourceserverscope/:argID", ConverHttprouterToGin(UpdateResourceServerScope))
	router.DELETE("/resourceserverscope/:argID", ConverHttprouterToGin(DeleteResourceServerScope))
}

// GetAllResourceServerScope is a function to get a slice of record(s) from resource_server_scope table in the keycloak database
// @Summary Get list of ResourceServerScope
// @Tags ResourceServerScope
// @Description GetAllResourceServerScope is a handler to get a slice of record(s) from resource_server_scope table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ResourceServerScope}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceserverscope [get]
// http "http://localhost:8080/resourceserverscope?page=0&pagesize=20" X-Api-User:user123
func GetAllResourceServerScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "resource_server_scope", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllResourceServerScope(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetResourceServerScope is a function to get a single record from the resource_server_scope table in the keycloak database
// @Summary Get record from table ResourceServerScope by  argID
// @Tags ResourceServerScope
// @ID argID
// @Description GetResourceServerScope is a function to get a single record from the resource_server_scope table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.ResourceServerScope
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /resourceserverscope/{argID} [get]
// http "http://localhost:8080/resourceserverscope/hello world" X-Api-User:user123
func GetResourceServerScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server_scope", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetResourceServerScope(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddResourceServerScope add to add a single record to resource_server_scope table in the keycloak database
// @Summary Add an record to resource_server_scope table
// @Description add to add a single record to resource_server_scope table in the keycloak database
// @Tags ResourceServerScope
// @Accept  json
// @Produce  json
// @Param ResourceServerScope body model.ResourceServerScope true "Add ResourceServerScope"
// @Success 200 {object} model.ResourceServerScope
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceserverscope [post]
// echo '{"id": "FhgfxtmBmVyqaIyjPlpNGxsWB","name": "ToySwgIoUMBqaulfdVLTvZhRr","icon_uri": "ssKFUZLoVEuxBwOmjqGOSeLbS","resource_server_id": "kvmRyEYcnquhdflcqdndrLZYL","display_name": "JutibSaOqpuWtqvmEuMFYlkTx"}' | http POST "http://localhost:8080/resourceserverscope" X-Api-User:user123
func AddResourceServerScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	resourceserverscope := &model.ResourceServerScope{}

	if err := readJSON(r, resourceserverscope); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resourceserverscope.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resourceserverscope.Prepare()

	if err := resourceserverscope.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server_scope", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	resourceserverscope, _, err = dao.AddResourceServerScope(ctx, resourceserverscope)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resourceserverscope)
}

// UpdateResourceServerScope Update a single record from resource_server_scope table in the keycloak database
// @Summary Update an record in table resource_server_scope
// @Description Update a single record from resource_server_scope table in the keycloak database
// @Tags ResourceServerScope
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  ResourceServerScope body model.ResourceServerScope true "Update ResourceServerScope record"
// @Success 200 {object} model.ResourceServerScope
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceserverscope/{argID} [put]
// echo '{"id": "FhgfxtmBmVyqaIyjPlpNGxsWB","name": "ToySwgIoUMBqaulfdVLTvZhRr","icon_uri": "ssKFUZLoVEuxBwOmjqGOSeLbS","resource_server_id": "kvmRyEYcnquhdflcqdndrLZYL","display_name": "JutibSaOqpuWtqvmEuMFYlkTx"}' | http PUT "http://localhost:8080/resourceserverscope/hello world"  X-Api-User:user123
func UpdateResourceServerScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resourceserverscope := &model.ResourceServerScope{}
	if err := readJSON(r, resourceserverscope); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resourceserverscope.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resourceserverscope.Prepare()

	if err := resourceserverscope.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server_scope", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resourceserverscope, _, err = dao.UpdateResourceServerScope(ctx,
		argID,
		resourceserverscope)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resourceserverscope)
}

// DeleteResourceServerScope Delete a single record from resource_server_scope table in the keycloak database
// @Summary Delete a record from resource_server_scope
// @Description Delete a single record from resource_server_scope table in the keycloak database
// @Tags ResourceServerScope
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.ResourceServerScope
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /resourceserverscope/{argID} [delete]
// http DELETE "http://localhost:8080/resourceserverscope/hello world" X-Api-User:user123
func DeleteResourceServerScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server_scope", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteResourceServerScope(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
