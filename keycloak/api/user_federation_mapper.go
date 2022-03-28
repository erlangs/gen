package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configUserFederationMapperRouter(router *httprouter.Router) {
	router.GET("/userfederationmapper", GetAllUserFederationMapper)
	router.POST("/userfederationmapper", AddUserFederationMapper)
	router.GET("/userfederationmapper/:argID", GetUserFederationMapper)
	router.PUT("/userfederationmapper/:argID", UpdateUserFederationMapper)
	router.DELETE("/userfederationmapper/:argID", DeleteUserFederationMapper)
}

func configGinUserFederationMapperRouter(router gin.IRoutes) {
	router.GET("/userfederationmapper", ConverHttprouterToGin(GetAllUserFederationMapper))
	router.POST("/userfederationmapper", ConverHttprouterToGin(AddUserFederationMapper))
	router.GET("/userfederationmapper/:argID", ConverHttprouterToGin(GetUserFederationMapper))
	router.PUT("/userfederationmapper/:argID", ConverHttprouterToGin(UpdateUserFederationMapper))
	router.DELETE("/userfederationmapper/:argID", ConverHttprouterToGin(DeleteUserFederationMapper))
}

// GetAllUserFederationMapper is a function to get a slice of record(s) from user_federation_mapper table in the keycloak database
// @Summary Get list of UserFederationMapper
// @Tags UserFederationMapper
// @Description GetAllUserFederationMapper is a handler to get a slice of record(s) from user_federation_mapper table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.UserFederationMapper}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userfederationmapper [get]
// http "http://localhost:8080/userfederationmapper?page=0&pagesize=20" X-Api-User:user123
func GetAllUserFederationMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "user_federation_mapper", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllUserFederationMapper(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetUserFederationMapper is a function to get a single record from the user_federation_mapper table in the keycloak database
// @Summary Get record from table UserFederationMapper by  argID
// @Tags UserFederationMapper
// @ID argID
// @Description GetUserFederationMapper is a function to get a single record from the user_federation_mapper table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.UserFederationMapper
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /userfederationmapper/{argID} [get]
// http "http://localhost:8080/userfederationmapper/hello world" X-Api-User:user123
func GetUserFederationMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_federation_mapper", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetUserFederationMapper(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddUserFederationMapper add to add a single record to user_federation_mapper table in the keycloak database
// @Summary Add an record to user_federation_mapper table
// @Description add to add a single record to user_federation_mapper table in the keycloak database
// @Tags UserFederationMapper
// @Accept  json
// @Produce  json
// @Param UserFederationMapper body model.UserFederationMapper true "Add UserFederationMapper"
// @Success 200 {object} model.UserFederationMapper
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userfederationmapper [post]
// echo '{"id": "lCcCQgErtaJRvxHIeVLomqAyX","name": "enHBOIeXgTRgvdNeVLmmUFunK","federation_provider_id": "LIicMIOpltOiDdXnRrMChthiJ","federation_mapper_type": "DsgMnVrcbFctwetIUWmDGqBsg","realm_id": "sytsLbZDSqEaAMddoKmJvhkcP"}' | http POST "http://localhost:8080/userfederationmapper" X-Api-User:user123
func AddUserFederationMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	userfederationmapper := &model.UserFederationMapper{}

	if err := readJSON(r, userfederationmapper); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userfederationmapper.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userfederationmapper.Prepare()

	if err := userfederationmapper.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_federation_mapper", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	userfederationmapper, _, err = dao.AddUserFederationMapper(ctx, userfederationmapper)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userfederationmapper)
}

// UpdateUserFederationMapper Update a single record from user_federation_mapper table in the keycloak database
// @Summary Update an record in table user_federation_mapper
// @Description Update a single record from user_federation_mapper table in the keycloak database
// @Tags UserFederationMapper
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  UserFederationMapper body model.UserFederationMapper true "Update UserFederationMapper record"
// @Success 200 {object} model.UserFederationMapper
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userfederationmapper/{argID} [put]
// echo '{"id": "lCcCQgErtaJRvxHIeVLomqAyX","name": "enHBOIeXgTRgvdNeVLmmUFunK","federation_provider_id": "LIicMIOpltOiDdXnRrMChthiJ","federation_mapper_type": "DsgMnVrcbFctwetIUWmDGqBsg","realm_id": "sytsLbZDSqEaAMddoKmJvhkcP"}' | http PUT "http://localhost:8080/userfederationmapper/hello world"  X-Api-User:user123
func UpdateUserFederationMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	userfederationmapper := &model.UserFederationMapper{}
	if err := readJSON(r, userfederationmapper); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userfederationmapper.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userfederationmapper.Prepare()

	if err := userfederationmapper.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_federation_mapper", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	userfederationmapper, _, err = dao.UpdateUserFederationMapper(ctx,
		argID,
		userfederationmapper)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userfederationmapper)
}

// DeleteUserFederationMapper Delete a single record from user_federation_mapper table in the keycloak database
// @Summary Delete a record from user_federation_mapper
// @Description Delete a single record from user_federation_mapper table in the keycloak database
// @Tags UserFederationMapper
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.UserFederationMapper
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /userfederationmapper/{argID} [delete]
// http DELETE "http://localhost:8080/userfederationmapper/hello world" X-Api-User:user123
func DeleteUserFederationMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_federation_mapper", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteUserFederationMapper(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
