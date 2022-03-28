package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configUserFederationMapperConfigRouter(router *httprouter.Router) {
	router.GET("/userfederationmapperconfig", GetAllUserFederationMapperConfig)
	router.POST("/userfederationmapperconfig", AddUserFederationMapperConfig)
	router.GET("/userfederationmapperconfig/:argUserFederationMapperID/:argName", GetUserFederationMapperConfig)
	router.PUT("/userfederationmapperconfig/:argUserFederationMapperID/:argName", UpdateUserFederationMapperConfig)
	router.DELETE("/userfederationmapperconfig/:argUserFederationMapperID/:argName", DeleteUserFederationMapperConfig)
}

func configGinUserFederationMapperConfigRouter(router gin.IRoutes) {
	router.GET("/userfederationmapperconfig", ConverHttprouterToGin(GetAllUserFederationMapperConfig))
	router.POST("/userfederationmapperconfig", ConverHttprouterToGin(AddUserFederationMapperConfig))
	router.GET("/userfederationmapperconfig/:argUserFederationMapperID/:argName", ConverHttprouterToGin(GetUserFederationMapperConfig))
	router.PUT("/userfederationmapperconfig/:argUserFederationMapperID/:argName", ConverHttprouterToGin(UpdateUserFederationMapperConfig))
	router.DELETE("/userfederationmapperconfig/:argUserFederationMapperID/:argName", ConverHttprouterToGin(DeleteUserFederationMapperConfig))
}

// GetAllUserFederationMapperConfig is a function to get a slice of record(s) from user_federation_mapper_config table in the keycloak database
// @Summary Get list of UserFederationMapperConfig
// @Tags UserFederationMapperConfig
// @Description GetAllUserFederationMapperConfig is a handler to get a slice of record(s) from user_federation_mapper_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.UserFederationMapperConfig}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userfederationmapperconfig [get]
// http "http://localhost:8080/userfederationmapperconfig?page=0&pagesize=20" X-Api-User:user123
func GetAllUserFederationMapperConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "user_federation_mapper_config", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllUserFederationMapperConfig(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetUserFederationMapperConfig is a function to get a single record from the user_federation_mapper_config table in the keycloak database
// @Summary Get record from table UserFederationMapperConfig by  argUserFederationMapperID  argName
// @Tags UserFederationMapperConfig
// @ID argUserFederationMapperID
// @ID argName
// @Description GetUserFederationMapperConfig is a function to get a single record from the user_federation_mapper_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argUserFederationMapperID path string true "user_federation_mapper_id"
// @Param  argName path string true "name"
// @Success 200 {object} model.UserFederationMapperConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /userfederationmapperconfig/{argUserFederationMapperID}/{argName} [get]
// http "http://localhost:8080/userfederationmapperconfig/hello world/hello world" X-Api-User:user123
func GetUserFederationMapperConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserFederationMapperID, err := parseString(ps, "argUserFederationMapperID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_federation_mapper_config", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetUserFederationMapperConfig(ctx, argUserFederationMapperID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddUserFederationMapperConfig add to add a single record to user_federation_mapper_config table in the keycloak database
// @Summary Add an record to user_federation_mapper_config table
// @Description add to add a single record to user_federation_mapper_config table in the keycloak database
// @Tags UserFederationMapperConfig
// @Accept  json
// @Produce  json
// @Param UserFederationMapperConfig body model.UserFederationMapperConfig true "Add UserFederationMapperConfig"
// @Success 200 {object} model.UserFederationMapperConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userfederationmapperconfig [post]
// echo '{"user_federation_mapper_id": "xFFdVTKaqsdVFTSlwjYaTeKkx","value": "XcRpdlbYoKjpTsmXotuclHgbU","name": "KArJjWQKPwPXiJOtRKrBWAmpe"}' | http POST "http://localhost:8080/userfederationmapperconfig" X-Api-User:user123
func AddUserFederationMapperConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	userfederationmapperconfig := &model.UserFederationMapperConfig{}

	if err := readJSON(r, userfederationmapperconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userfederationmapperconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userfederationmapperconfig.Prepare()

	if err := userfederationmapperconfig.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_federation_mapper_config", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	userfederationmapperconfig, _, err = dao.AddUserFederationMapperConfig(ctx, userfederationmapperconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userfederationmapperconfig)
}

// UpdateUserFederationMapperConfig Update a single record from user_federation_mapper_config table in the keycloak database
// @Summary Update an record in table user_federation_mapper_config
// @Description Update a single record from user_federation_mapper_config table in the keycloak database
// @Tags UserFederationMapperConfig
// @Accept  json
// @Produce  json
// @Param  argUserFederationMapperID path string true "user_federation_mapper_id"// @Param  argName path string true "name"
// @Param  UserFederationMapperConfig body model.UserFederationMapperConfig true "Update UserFederationMapperConfig record"
// @Success 200 {object} model.UserFederationMapperConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userfederationmapperconfig/{argUserFederationMapperID}/{argName} [put]
// echo '{"user_federation_mapper_id": "xFFdVTKaqsdVFTSlwjYaTeKkx","value": "XcRpdlbYoKjpTsmXotuclHgbU","name": "KArJjWQKPwPXiJOtRKrBWAmpe"}' | http PUT "http://localhost:8080/userfederationmapperconfig/hello world/hello world"  X-Api-User:user123
func UpdateUserFederationMapperConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserFederationMapperID, err := parseString(ps, "argUserFederationMapperID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	userfederationmapperconfig := &model.UserFederationMapperConfig{}
	if err := readJSON(r, userfederationmapperconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userfederationmapperconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userfederationmapperconfig.Prepare()

	if err := userfederationmapperconfig.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_federation_mapper_config", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	userfederationmapperconfig, _, err = dao.UpdateUserFederationMapperConfig(ctx,
		argUserFederationMapperID, argName,
		userfederationmapperconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userfederationmapperconfig)
}

// DeleteUserFederationMapperConfig Delete a single record from user_federation_mapper_config table in the keycloak database
// @Summary Delete a record from user_federation_mapper_config
// @Description Delete a single record from user_federation_mapper_config table in the keycloak database
// @Tags UserFederationMapperConfig
// @Accept  json
// @Produce  json
// @Param  argUserFederationMapperID path string true "user_federation_mapper_id"// @Param  argName path string true "name"
// @Success 204 {object} model.UserFederationMapperConfig
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /userfederationmapperconfig/{argUserFederationMapperID}/{argName} [delete]
// http DELETE "http://localhost:8080/userfederationmapperconfig/hello world/hello world" X-Api-User:user123
func DeleteUserFederationMapperConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserFederationMapperID, err := parseString(ps, "argUserFederationMapperID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_federation_mapper_config", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteUserFederationMapperConfig(ctx, argUserFederationMapperID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
