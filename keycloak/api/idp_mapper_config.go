package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configIdpMapperConfigRouter(router *httprouter.Router) {
	router.GET("/idpmapperconfig", GetAllIdpMapperConfig)
	router.POST("/idpmapperconfig", AddIdpMapperConfig)
	router.GET("/idpmapperconfig/:argIdpMapperID/:argName", GetIdpMapperConfig)
	router.PUT("/idpmapperconfig/:argIdpMapperID/:argName", UpdateIdpMapperConfig)
	router.DELETE("/idpmapperconfig/:argIdpMapperID/:argName", DeleteIdpMapperConfig)
}

func configGinIdpMapperConfigRouter(router gin.IRoutes) {
	router.GET("/idpmapperconfig", ConverHttprouterToGin(GetAllIdpMapperConfig))
	router.POST("/idpmapperconfig", ConverHttprouterToGin(AddIdpMapperConfig))
	router.GET("/idpmapperconfig/:argIdpMapperID/:argName", ConverHttprouterToGin(GetIdpMapperConfig))
	router.PUT("/idpmapperconfig/:argIdpMapperID/:argName", ConverHttprouterToGin(UpdateIdpMapperConfig))
	router.DELETE("/idpmapperconfig/:argIdpMapperID/:argName", ConverHttprouterToGin(DeleteIdpMapperConfig))
}

// GetAllIdpMapperConfig is a function to get a slice of record(s) from idp_mapper_config table in the keycloak database
// @Summary Get list of IdpMapperConfig
// @Tags IdpMapperConfig
// @Description GetAllIdpMapperConfig is a handler to get a slice of record(s) from idp_mapper_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.IdpMapperConfig}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /idpmapperconfig [get]
// http "http://localhost:8080/idpmapperconfig?page=0&pagesize=20" X-Api-User:user123
func GetAllIdpMapperConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "idp_mapper_config", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllIdpMapperConfig(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetIdpMapperConfig is a function to get a single record from the idp_mapper_config table in the keycloak database
// @Summary Get record from table IdpMapperConfig by  argIdpMapperID  argName
// @Tags IdpMapperConfig
// @ID argIdpMapperID
// @ID argName
// @Description GetIdpMapperConfig is a function to get a single record from the idp_mapper_config table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argIdpMapperID path string true "idp_mapper_id"
// @Param  argName path string true "name"
// @Success 200 {object} model.IdpMapperConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /idpmapperconfig/{argIdpMapperID}/{argName} [get]
// http "http://localhost:8080/idpmapperconfig/hello world/hello world" X-Api-User:user123
func GetIdpMapperConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argIdpMapperID, err := parseString(ps, "argIdpMapperID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "idp_mapper_config", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetIdpMapperConfig(ctx, argIdpMapperID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddIdpMapperConfig add to add a single record to idp_mapper_config table in the keycloak database
// @Summary Add an record to idp_mapper_config table
// @Description add to add a single record to idp_mapper_config table in the keycloak database
// @Tags IdpMapperConfig
// @Accept  json
// @Produce  json
// @Param IdpMapperConfig body model.IdpMapperConfig true "Add IdpMapperConfig"
// @Success 200 {object} model.IdpMapperConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /idpmapperconfig [post]
// echo '{"idp_mapper_id": "RKvRKbFnKiWlZqBgLwdwNJxSt","value": "lQndviXROsHdsJERyvCHIvewo","name": "fVCeHMbTpgsmDKfGmDjflNbeN"}' | http POST "http://localhost:8080/idpmapperconfig" X-Api-User:user123
func AddIdpMapperConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	idpmapperconfig := &model.IdpMapperConfig{}

	if err := readJSON(r, idpmapperconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := idpmapperconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	idpmapperconfig.Prepare()

	if err := idpmapperconfig.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "idp_mapper_config", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	idpmapperconfig, _, err = dao.AddIdpMapperConfig(ctx, idpmapperconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, idpmapperconfig)
}

// UpdateIdpMapperConfig Update a single record from idp_mapper_config table in the keycloak database
// @Summary Update an record in table idp_mapper_config
// @Description Update a single record from idp_mapper_config table in the keycloak database
// @Tags IdpMapperConfig
// @Accept  json
// @Produce  json
// @Param  argIdpMapperID path string true "idp_mapper_id"// @Param  argName path string true "name"
// @Param  IdpMapperConfig body model.IdpMapperConfig true "Update IdpMapperConfig record"
// @Success 200 {object} model.IdpMapperConfig
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /idpmapperconfig/{argIdpMapperID}/{argName} [put]
// echo '{"idp_mapper_id": "RKvRKbFnKiWlZqBgLwdwNJxSt","value": "lQndviXROsHdsJERyvCHIvewo","name": "fVCeHMbTpgsmDKfGmDjflNbeN"}' | http PUT "http://localhost:8080/idpmapperconfig/hello world/hello world"  X-Api-User:user123
func UpdateIdpMapperConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argIdpMapperID, err := parseString(ps, "argIdpMapperID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	idpmapperconfig := &model.IdpMapperConfig{}
	if err := readJSON(r, idpmapperconfig); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := idpmapperconfig.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	idpmapperconfig.Prepare()

	if err := idpmapperconfig.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "idp_mapper_config", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	idpmapperconfig, _, err = dao.UpdateIdpMapperConfig(ctx,
		argIdpMapperID, argName,
		idpmapperconfig)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, idpmapperconfig)
}

// DeleteIdpMapperConfig Delete a single record from idp_mapper_config table in the keycloak database
// @Summary Delete a record from idp_mapper_config
// @Description Delete a single record from idp_mapper_config table in the keycloak database
// @Tags IdpMapperConfig
// @Accept  json
// @Produce  json
// @Param  argIdpMapperID path string true "idp_mapper_id"// @Param  argName path string true "name"
// @Success 204 {object} model.IdpMapperConfig
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /idpmapperconfig/{argIdpMapperID}/{argName} [delete]
// http DELETE "http://localhost:8080/idpmapperconfig/hello world/hello world" X-Api-User:user123
func DeleteIdpMapperConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argIdpMapperID, err := parseString(ps, "argIdpMapperID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "idp_mapper_config", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteIdpMapperConfig(ctx, argIdpMapperID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
