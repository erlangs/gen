package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configProtocolMapperConfigRouter(router *httprouter.Router) {
	router.GET("/protocolmapperconfig", GetAllProtocolMapperConfig)
	router.POST("/protocolmapperconfig", AddProtocolMapperConfig)
	router.GET("/protocolmapperconfig/:argProtocolMapperID/:argName", GetProtocolMapperConfig)
	router.PUT("/protocolmapperconfig/:argProtocolMapperID/:argName", UpdateProtocolMapperConfig)
	router.DELETE("/protocolmapperconfig/:argProtocolMapperID/:argName", DeleteProtocolMapperConfig)
}

func configGinProtocolMapperConfigRouter(router gin.IRoutes) {
	router.GET("/protocolmapperconfig", ConverHttprouterToGin(GetAllProtocolMapperConfig))
	router.POST("/protocolmapperconfig", ConverHttprouterToGin(AddProtocolMapperConfig))
	router.GET("/protocolmapperconfig/:argProtocolMapperID/:argName", ConverHttprouterToGin(GetProtocolMapperConfig))
	router.PUT("/protocolmapperconfig/:argProtocolMapperID/:argName", ConverHttprouterToGin(UpdateProtocolMapperConfig))
	router.DELETE("/protocolmapperconfig/:argProtocolMapperID/:argName", ConverHttprouterToGin(DeleteProtocolMapperConfig))
}

// GetAllProtocolMapperConfig is a function to get a slice of record(s) from protocol_mapper_config table in the keycloak database
// @Summary Get list of ProtocolMapperConfig
// @Tags ProtocolMapperConfig
// @Description GetAllProtocolMapperConfig is a handler to get a slice of record(s) from protocol_mapper_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ProtocolMapperConfig}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /protocolmapperconfig [get]
// http "http://localhost:8080/protocolmapperconfig?page=0&pagesize=20" X-Api-User:user123
func GetAllProtocolMapperConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "protocol_mapper_config", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllProtocolMapperConfig(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetProtocolMapperConfig is a function to get a single record from the protocol_mapper_config table in the keycloak database
// @Summary Get record from table ProtocolMapperConfig by  argProtocolMapperID  argName
// @Tags ProtocolMapperConfig
// @ID argProtocolMapperID
// @ID argName
// @Description GetProtocolMapperConfig is a function to get a single record from the protocol_mapper_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argProtocolMapperID path string true "protocol_mapper_id"
// @Param  argName path string true "name"
// @Success 200 {object} model.ProtocolMapperConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /protocolmapperconfig/{argProtocolMapperID}/{argName} [get]
// http "http://localhost:8080/protocolmapperconfig/hello world/hello world" X-Api-User:user123
func GetProtocolMapperConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argProtocolMapperID, err := parseString(ps, "argProtocolMapperID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "protocol_mapper_config", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetProtocolMapperConfig(ctx, argProtocolMapperID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddProtocolMapperConfig add to add a single record to protocol_mapper_config table in the keycloak database
// @Summary Add an record to protocol_mapper_config table
// @Description add to add a single record to protocol_mapper_config table in the keycloak database
// @Tags ProtocolMapperConfig
// @Accept  json
// @Produce  json
// @Param ProtocolMapperConfig body model.ProtocolMapperConfig true "Add ProtocolMapperConfig"
// @Success 200 {object} model.ProtocolMapperConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /protocolmapperconfig [post]
// echo '{"protocol_mapper_id": "JjhNDTPIvjoDxZBVTqYAFrYGB","value": "crhBLEftlqWKfjAoinZOHkRmL","name": "mrBocWpHCHFCRbchspmKdLSds"}' | http POST "http://localhost:8080/protocolmapperconfig" X-Api-User:user123
func AddProtocolMapperConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	protocolmapperconfig := &model.ProtocolMapperConfig{}

	if err := readJSON(r, protocolmapperconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := protocolmapperconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	protocolmapperconfig.Prepare()

	if err := protocolmapperconfig.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "protocol_mapper_config", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	protocolmapperconfig, _, err = dao.AddProtocolMapperConfig(ctx, protocolmapperconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, protocolmapperconfig)
}

// UpdateProtocolMapperConfig Update a single record from protocol_mapper_config table in the keycloak database
// @Summary Update an record in table protocol_mapper_config
// @Description Update a single record from protocol_mapper_config table in the keycloak database
// @Tags ProtocolMapperConfig
// @Accept  json
// @Produce  json
// @Param  argProtocolMapperID path string true "protocol_mapper_id"// @Param  argName path string true "name"
// @Param  ProtocolMapperConfig body model.ProtocolMapperConfig true "Update ProtocolMapperConfig record"
// @Success 200 {object} model.ProtocolMapperConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /protocolmapperconfig/{argProtocolMapperID}/{argName} [put]
// echo '{"protocol_mapper_id": "JjhNDTPIvjoDxZBVTqYAFrYGB","value": "crhBLEftlqWKfjAoinZOHkRmL","name": "mrBocWpHCHFCRbchspmKdLSds"}' | http PUT "http://localhost:8080/protocolmapperconfig/hello world/hello world"  X-Api-User:user123
func UpdateProtocolMapperConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argProtocolMapperID, err := parseString(ps, "argProtocolMapperID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	protocolmapperconfig := &model.ProtocolMapperConfig{}
	if err := readJSON(r, protocolmapperconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := protocolmapperconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	protocolmapperconfig.Prepare()

	if err := protocolmapperconfig.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "protocol_mapper_config", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	protocolmapperconfig, _, err = dao.UpdateProtocolMapperConfig(ctx,
		argProtocolMapperID, argName,
		protocolmapperconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, protocolmapperconfig)
}

// DeleteProtocolMapperConfig Delete a single record from protocol_mapper_config table in the keycloak database
// @Summary Delete a record from protocol_mapper_config
// @Description Delete a single record from protocol_mapper_config table in the keycloak database
// @Tags ProtocolMapperConfig
// @Accept  json
// @Produce  json
// @Param  argProtocolMapperID path string true "protocol_mapper_id"// @Param  argName path string true "name"
// @Success 204 {object} model.ProtocolMapperConfig
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /protocolmapperconfig/{argProtocolMapperID}/{argName} [delete]
// http DELETE "http://localhost:8080/protocolmapperconfig/hello world/hello world" X-Api-User:user123
func DeleteProtocolMapperConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argProtocolMapperID, err := parseString(ps, "argProtocolMapperID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "protocol_mapper_config", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteProtocolMapperConfig(ctx, argProtocolMapperID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
