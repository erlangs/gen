package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configComponentRouter(router *httprouter.Router) {
	router.GET("/component", GetAllComponent)
	router.POST("/component", AddComponent)
	router.GET("/component/:argID", GetComponent)
	router.PUT("/component/:argID", UpdateComponent)
	router.DELETE("/component/:argID", DeleteComponent)
}

func configGinComponentRouter(router gin.IRoutes) {
	router.GET("/component", ConverHttprouterToGin(GetAllComponent))
	router.POST("/component", ConverHttprouterToGin(AddComponent))
	router.GET("/component/:argID", ConverHttprouterToGin(GetComponent))
	router.PUT("/component/:argID", ConverHttprouterToGin(UpdateComponent))
	router.DELETE("/component/:argID", ConverHttprouterToGin(DeleteComponent))
}

// GetAllComponent is a function to get a slice of record(s) from component table in the keycloak database
// @Summary Get list of Component
// @Tags Component
// @Description GetAllComponent is a handler to get a slice of record(s) from component table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Component}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /component [get]
// http "http://localhost:8080/component?page=0&pagesize=20" X-Api-User:user123
func GetAllComponent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "component", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllComponent(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetComponent is a function to get a single record from the component table in the keycloak database
// @Summary Get record from table Component by  argID
// @Tags Component
// @ID argID
// @Description GetComponent is a function to get a single record from the component table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.Component
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /component/{argID} [get]
// http "http://localhost:8080/component/hello world" X-Api-User:user123
func GetComponent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "component", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetComponent(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddComponent add to add a single record to component table in the keycloak database
// @Summary Add an record to component table
// @Description add to add a single record to component table in the keycloak database
// @Tags Component
// @Accept  json
// @Produce  json
// @Param Component body model.Component true "Add Component"
// @Success 200 {object} model.Component
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /component [post]
// echo '{"id": "DZaaAZuaACrMEsAbDsBmWYpMX","name": "CjQEENFOanBTrRZmLaQkyFmTq","parent_id": "THtxXXULFEKEHYOExeIJjiqku","provider_id": "JIZZLGRwhlittnUvhckwdwkiO","provider_type": "UMHUdfbqQjnQaChfJCrcncZPc","realm_id": "NnCFZTDecqccpaiPdlLLZufqa","sub_type": "YpFslimHGAtrHaZXTRXwKmkSs"}' | http POST "http://localhost:8080/component" X-Api-User:user123
func AddComponent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	component := &model.Component{}

	if err := readJSON(r, component); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := component.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	component.Prepare()

	if err := component.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "component", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	component, _, err = dao.AddComponent(ctx, component)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, component)
}

// UpdateComponent Update a single record from component table in the keycloak database
// @Summary Update an record in table component
// @Description Update a single record from component table in the keycloak database
// @Tags Component
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  Component body model.Component true "Update Component record"
// @Success 200 {object} model.Component
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /component/{argID} [put]
// echo '{"id": "DZaaAZuaACrMEsAbDsBmWYpMX","name": "CjQEENFOanBTrRZmLaQkyFmTq","parent_id": "THtxXXULFEKEHYOExeIJjiqku","provider_id": "JIZZLGRwhlittnUvhckwdwkiO","provider_type": "UMHUdfbqQjnQaChfJCrcncZPc","realm_id": "NnCFZTDecqccpaiPdlLLZufqa","sub_type": "YpFslimHGAtrHaZXTRXwKmkSs"}' | http PUT "http://localhost:8080/component/hello world"  X-Api-User:user123
func UpdateComponent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	component := &model.Component{}
	if err := readJSON(r, component); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := component.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	component.Prepare()

	if err := component.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "component", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	component, _, err = dao.UpdateComponent(ctx,
		argID,
		component)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, component)
}

// DeleteComponent Delete a single record from component table in the keycloak database
// @Summary Delete a record from component
// @Description Delete a single record from component table in the keycloak database
// @Tags Component
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.Component
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /component/{argID} [delete]
// http DELETE "http://localhost:8080/component/hello world" X-Api-User:user123
func DeleteComponent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "component", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteComponent(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
