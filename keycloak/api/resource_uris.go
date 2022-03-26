package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configResourceUrisRouter(router *httprouter.Router) {
	router.GET("/resourceuris", GetAllResourceUris)
	router.POST("/resourceuris", AddResourceUris)
	router.GET("/resourceuris/:argResourceID/:argValue", GetResourceUris)
	router.PUT("/resourceuris/:argResourceID/:argValue", UpdateResourceUris)
	router.DELETE("/resourceuris/:argResourceID/:argValue", DeleteResourceUris)
}

func configGinResourceUrisRouter(router gin.IRoutes) {
	router.GET("/resourceuris", ConverHttprouterToGin(GetAllResourceUris))
	router.POST("/resourceuris", ConverHttprouterToGin(AddResourceUris))
	router.GET("/resourceuris/:argResourceID/:argValue", ConverHttprouterToGin(GetResourceUris))
	router.PUT("/resourceuris/:argResourceID/:argValue", ConverHttprouterToGin(UpdateResourceUris))
	router.DELETE("/resourceuris/:argResourceID/:argValue", ConverHttprouterToGin(DeleteResourceUris))
}

// GetAllResourceUris is a function to get a slice of record(s) from resource_uris table in the keycloak database
// @Summary Get list of ResourceUris
// @Tags ResourceUris
// @Description GetAllResourceUris is a handler to get a slice of record(s) from resource_uris table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ResourceUris}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceuris [get]
// http "http://localhost:8080/resourceuris?page=0&pagesize=20" X-Api-User:user123
func GetAllResourceUris(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "resource_uris", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllResourceUris(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetResourceUris is a function to get a single record from the resource_uris table in the keycloak database
// @Summary Get record from table ResourceUris by  argResourceID  argValue
// @Tags ResourceUris
// @ID argResourceID
// @ID argValue
// @Description GetResourceUris is a function to get a single record from the resource_uris table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argResourceID path string true "resource_id"
// @Param  argValue path string true "value"
// @Success 200 {object} model.ResourceUris
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /resourceuris/{argResourceID}/{argValue} [get]
// http "http://localhost:8080/resourceuris/hello world/hello world" X-Api-User:user123
func GetResourceUris(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argResourceID, err := parseString(ps, "argResourceID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argValue, err := parseString(ps, "argValue")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_uris", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetResourceUris(ctx, argResourceID, argValue)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddResourceUris add to add a single record to resource_uris table in the keycloak database
// @Summary Add an record to resource_uris table
// @Description add to add a single record to resource_uris table in the keycloak database
// @Tags ResourceUris
// @Accept  json
// @Produce  json
// @Param ResourceUris body model.ResourceUris true "Add ResourceUris"
// @Success 200 {object} model.ResourceUris
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceuris [post]
// echo '{"resource_id": "cBCgvJgSLKThsZagOwlldZmqu","value": "YfBZxBLUdxOinYCNFxUmyEVSa"}' | http POST "http://localhost:8080/resourceuris" X-Api-User:user123
func AddResourceUris(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	resourceuris := &model.ResourceUris{}

	if err := readJSON(r, resourceuris); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resourceuris.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resourceuris.Prepare()

	if err := resourceuris.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_uris", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	resourceuris, _, err = dao.AddResourceUris(ctx, resourceuris)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resourceuris)
}

// UpdateResourceUris Update a single record from resource_uris table in the keycloak database
// @Summary Update an record in table resource_uris
// @Description Update a single record from resource_uris table in the keycloak database
// @Tags ResourceUris
// @Accept  json
// @Produce  json
// @Param  argResourceID path string true "resource_id"// @Param  argValue path string true "value"
// @Param  ResourceUris body model.ResourceUris true "Update ResourceUris record"
// @Success 200 {object} model.ResourceUris
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceuris/{argResourceID}/{argValue} [put]
// echo '{"resource_id": "cBCgvJgSLKThsZagOwlldZmqu","value": "YfBZxBLUdxOinYCNFxUmyEVSa"}' | http PUT "http://localhost:8080/resourceuris/hello world/hello world"  X-Api-User:user123
func UpdateResourceUris(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argResourceID, err := parseString(ps, "argResourceID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argValue, err := parseString(ps, "argValue")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resourceuris := &model.ResourceUris{}
	if err := readJSON(r, resourceuris); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resourceuris.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resourceuris.Prepare()

	if err := resourceuris.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_uris", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resourceuris, _, err = dao.UpdateResourceUris(ctx,
		argResourceID, argValue,
		resourceuris)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resourceuris)
}

// DeleteResourceUris Delete a single record from resource_uris table in the keycloak database
// @Summary Delete a record from resource_uris
// @Description Delete a single record from resource_uris table in the keycloak database
// @Tags ResourceUris
// @Accept  json
// @Produce  json
// @Param  argResourceID path string true "resource_id"// @Param  argValue path string true "value"
// @Success 204 {object} model.ResourceUris
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /resourceuris/{argResourceID}/{argValue} [delete]
// http DELETE "http://localhost:8080/resourceuris/hello world/hello world" X-Api-User:user123
func DeleteResourceUris(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argResourceID, err := parseString(ps, "argResourceID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argValue, err := parseString(ps, "argValue")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_uris", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteResourceUris(ctx, argResourceID, argValue)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
