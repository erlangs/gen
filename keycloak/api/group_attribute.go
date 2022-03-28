package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configGroupAttributeRouter(router *httprouter.Router) {
	router.GET("/groupattribute", GetAllGroupAttribute)
	router.POST("/groupattribute", AddGroupAttribute)
	router.GET("/groupattribute/:argID", GetGroupAttribute)
	router.PUT("/groupattribute/:argID", UpdateGroupAttribute)
	router.DELETE("/groupattribute/:argID", DeleteGroupAttribute)
}

func configGinGroupAttributeRouter(router gin.IRoutes) {
	router.GET("/groupattribute", ConverHttprouterToGin(GetAllGroupAttribute))
	router.POST("/groupattribute", ConverHttprouterToGin(AddGroupAttribute))
	router.GET("/groupattribute/:argID", ConverHttprouterToGin(GetGroupAttribute))
	router.PUT("/groupattribute/:argID", ConverHttprouterToGin(UpdateGroupAttribute))
	router.DELETE("/groupattribute/:argID", ConverHttprouterToGin(DeleteGroupAttribute))
}

// GetAllGroupAttribute is a function to get a slice of record(s) from group_attribute table in the keycloak database
// @Summary Get list of GroupAttribute
// @Tags GroupAttribute
// @Description GetAllGroupAttribute is a handler to get a slice of record(s) from group_attribute table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.GroupAttribute}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /groupattribute [get]
// http "http://localhost:8080/groupattribute?page=0&pagesize=20" X-Api-User:user123
func GetAllGroupAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "group_attribute", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllGroupAttribute(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetGroupAttribute is a function to get a single record from the group_attribute table in the keycloak database
// @Summary Get record from table GroupAttribute by  argID
// @Tags GroupAttribute
// @ID argID
// @Description GetGroupAttribute is a function to get a single record from the group_attribute table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.GroupAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /groupattribute/{argID} [get]
// http "http://localhost:8080/groupattribute/hello world" X-Api-User:user123
func GetGroupAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "group_attribute", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetGroupAttribute(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddGroupAttribute add to add a single record to group_attribute table in the keycloak database
// @Summary Add an record to group_attribute table
// @Description add to add a single record to group_attribute table in the keycloak database
// @Tags GroupAttribute
// @Accept  json
// @Produce  json
// @Param GroupAttribute body model.GroupAttribute true "Add GroupAttribute"
// @Success 200 {object} model.GroupAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /groupattribute [post]
// echo '{"id": "VQYncgPagZqqEGeaoeCHBpyrT","name": "VrKYaATvgCbMwUAieQrtBgkTN","value": "BbQVpRWVUOyNPOFxlEXvQodcF","group_id": "YXmSRXTOIcdkVlTEVEMsZDcZy"}' | http POST "http://localhost:8080/groupattribute" X-Api-User:user123
func AddGroupAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	groupattribute := &model.GroupAttribute{}

	if err := readJSON(r, groupattribute); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := groupattribute.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	groupattribute.Prepare()

	if err := groupattribute.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "group_attribute", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	groupattribute, _, err = dao.AddGroupAttribute(ctx, groupattribute)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, groupattribute)
}

// UpdateGroupAttribute Update a single record from group_attribute table in the keycloak database
// @Summary Update an record in table group_attribute
// @Description Update a single record from group_attribute table in the keycloak database
// @Tags GroupAttribute
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  GroupAttribute body model.GroupAttribute true "Update GroupAttribute record"
// @Success 200 {object} model.GroupAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /groupattribute/{argID} [put]
// echo '{"id": "VQYncgPagZqqEGeaoeCHBpyrT","name": "VrKYaATvgCbMwUAieQrtBgkTN","value": "BbQVpRWVUOyNPOFxlEXvQodcF","group_id": "YXmSRXTOIcdkVlTEVEMsZDcZy"}' | http PUT "http://localhost:8080/groupattribute/hello world"  X-Api-User:user123
func UpdateGroupAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	groupattribute := &model.GroupAttribute{}
	if err := readJSON(r, groupattribute); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := groupattribute.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	groupattribute.Prepare()

	if err := groupattribute.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "group_attribute", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	groupattribute, _, err = dao.UpdateGroupAttribute(ctx,
		argID,
		groupattribute)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, groupattribute)
}

// DeleteGroupAttribute Delete a single record from group_attribute table in the keycloak database
// @Summary Delete a record from group_attribute
// @Description Delete a single record from group_attribute table in the keycloak database
// @Tags GroupAttribute
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.GroupAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /groupattribute/{argID} [delete]
// http DELETE "http://localhost:8080/groupattribute/hello world" X-Api-User:user123
func DeleteGroupAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "group_attribute", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteGroupAttribute(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
