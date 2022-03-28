package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configUserEntityRouter(router *httprouter.Router) {
	router.GET("/userentity", GetAllUserEntity)
	router.POST("/userentity", AddUserEntity)
	router.GET("/userentity/:argID", GetUserEntity)
	router.PUT("/userentity/:argID", UpdateUserEntity)
	router.DELETE("/userentity/:argID", DeleteUserEntity)
}

func configGinUserEntityRouter(router gin.IRoutes) {
	router.GET("/userentity", ConverHttprouterToGin(GetAllUserEntity))
	router.POST("/userentity", ConverHttprouterToGin(AddUserEntity))
	router.GET("/userentity/:argID", ConverHttprouterToGin(GetUserEntity))
	router.PUT("/userentity/:argID", ConverHttprouterToGin(UpdateUserEntity))
	router.DELETE("/userentity/:argID", ConverHttprouterToGin(DeleteUserEntity))
}

// GetAllUserEntity is a function to get a slice of record(s) from user_entity table in the keycloak database
// @Summary Get list of UserEntity
// @Tags UserEntity
// @Description GetAllUserEntity is a handler to get a slice of record(s) from user_entity table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.UserEntity}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userentity [get]
// http "http://localhost:8080/userentity?page=0&pagesize=20" X-Api-User:user123
func GetAllUserEntity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "user_entity", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllUserEntity(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetUserEntity is a function to get a single record from the user_entity table in the keycloak database
// @Summary Get record from table UserEntity by  argID
// @Tags UserEntity
// @ID argID
// @Description GetUserEntity is a function to get a single record from the user_entity table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.UserEntity
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /userentity/{argID} [get]
// http "http://localhost:8080/userentity/hello world" X-Api-User:user123
func GetUserEntity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_entity", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetUserEntity(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddUserEntity add to add a single record to user_entity table in the keycloak database
// @Summary Add an record to user_entity table
// @Description add to add a single record to user_entity table in the keycloak database
// @Tags UserEntity
// @Accept  json
// @Produce  json
// @Param UserEntity body model.UserEntity true "Add UserEntity"
// @Success 200 {object} model.UserEntity
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userentity [post]
// echo '{"id": "hXaAuMRkfoSTwdZtPXFdtbETT","email": "WcMwEckuMbOYSBdLnxVBNUQgh","email_constraint": "tKrRGKhlfskFYWExyMTAMARxe","email_verified": true,"enabled": true,"federation_link": "MQZBhyMauTUmepXeYVOevFdwg","first_name": "mkYNQWlWBixchMOWekuhslAZo","last_name": "oAWxleewsPewsaJtqcQTXWmvt","realm_id": "YpvLaSYsBkLRahZsLxNIGBmKO","username": "JAQDEsHCdHXGnYDQbOwnsDhbp","created_timestamp": 82,"service_account_client_link": "hcjNVwOJkKVLLcOAxJUrfassh","not_before": 57}' | http POST "http://localhost:8080/userentity" X-Api-User:user123
func AddUserEntity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	userentity := &model.UserEntity{}

	if err := readJSON(r, userentity); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userentity.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userentity.Prepare()

	if err := userentity.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_entity", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	userentity, _, err = dao.AddUserEntity(ctx, userentity)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userentity)
}

// UpdateUserEntity Update a single record from user_entity table in the keycloak database
// @Summary Update an record in table user_entity
// @Description Update a single record from user_entity table in the keycloak database
// @Tags UserEntity
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  UserEntity body model.UserEntity true "Update UserEntity record"
// @Success 200 {object} model.UserEntity
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userentity/{argID} [put]
// echo '{"id": "hXaAuMRkfoSTwdZtPXFdtbETT","email": "WcMwEckuMbOYSBdLnxVBNUQgh","email_constraint": "tKrRGKhlfskFYWExyMTAMARxe","email_verified": true,"enabled": true,"federation_link": "MQZBhyMauTUmepXeYVOevFdwg","first_name": "mkYNQWlWBixchMOWekuhslAZo","last_name": "oAWxleewsPewsaJtqcQTXWmvt","realm_id": "YpvLaSYsBkLRahZsLxNIGBmKO","username": "JAQDEsHCdHXGnYDQbOwnsDhbp","created_timestamp": 82,"service_account_client_link": "hcjNVwOJkKVLLcOAxJUrfassh","not_before": 57}' | http PUT "http://localhost:8080/userentity/hello world"  X-Api-User:user123
func UpdateUserEntity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	userentity := &model.UserEntity{}
	if err := readJSON(r, userentity); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userentity.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userentity.Prepare()

	if err := userentity.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_entity", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	userentity, _, err = dao.UpdateUserEntity(ctx,
		argID,
		userentity)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userentity)
}

// DeleteUserEntity Delete a single record from user_entity table in the keycloak database
// @Summary Delete a record from user_entity
// @Description Delete a single record from user_entity table in the keycloak database
// @Tags UserEntity
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.UserEntity
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /userentity/{argID} [delete]
// http DELETE "http://localhost:8080/userentity/hello world" X-Api-User:user123
func DeleteUserEntity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_entity", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteUserEntity(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
