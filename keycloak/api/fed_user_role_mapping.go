package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configFedUserRoleMappingRouter(router *httprouter.Router) {
	router.GET("/feduserrolemapping", GetAllFedUserRoleMapping)
	router.POST("/feduserrolemapping", AddFedUserRoleMapping)
	router.GET("/feduserrolemapping/:argRoleID/:argUserID", GetFedUserRoleMapping)
	router.PUT("/feduserrolemapping/:argRoleID/:argUserID", UpdateFedUserRoleMapping)
	router.DELETE("/feduserrolemapping/:argRoleID/:argUserID", DeleteFedUserRoleMapping)
}

func configGinFedUserRoleMappingRouter(router gin.IRoutes) {
	router.GET("/feduserrolemapping", ConverHttprouterToGin(GetAllFedUserRoleMapping))
	router.POST("/feduserrolemapping", ConverHttprouterToGin(AddFedUserRoleMapping))
	router.GET("/feduserrolemapping/:argRoleID/:argUserID", ConverHttprouterToGin(GetFedUserRoleMapping))
	router.PUT("/feduserrolemapping/:argRoleID/:argUserID", ConverHttprouterToGin(UpdateFedUserRoleMapping))
	router.DELETE("/feduserrolemapping/:argRoleID/:argUserID", ConverHttprouterToGin(DeleteFedUserRoleMapping))
}

// GetAllFedUserRoleMapping is a function to get a slice of record(s) from fed_user_role_mapping table in the keycloak database
// @Summary Get list of FedUserRoleMapping
// @Tags FedUserRoleMapping
// @Description GetAllFedUserRoleMapping is a handler to get a slice of record(s) from fed_user_role_mapping table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.FedUserRoleMapping}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /feduserrolemapping [get]
// http "http://localhost:8080/feduserrolemapping?page=0&pagesize=20" X-Api-User:user123
func GetAllFedUserRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "fed_user_role_mapping", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllFedUserRoleMapping(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetFedUserRoleMapping is a function to get a single record from the fed_user_role_mapping table in the keycloak database
// @Summary Get record from table FedUserRoleMapping by  argRoleID  argUserID
// @Tags FedUserRoleMapping
// @ID argRoleID
// @ID argUserID
// @Description GetFedUserRoleMapping is a function to get a single record from the fed_user_role_mapping table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argRoleID path string true "role_id"
// @Param  argUserID path string true "user_id"
// @Success 200 {object} model.FedUserRoleMapping
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /feduserrolemapping/{argRoleID}/{argUserID} [get]
// http "http://localhost:8080/feduserrolemapping/hello world/hello world" X-Api-User:user123
func GetFedUserRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRoleID, err := parseString(ps, "argRoleID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argUserID, err := parseString(ps, "argUserID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_role_mapping", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetFedUserRoleMapping(ctx, argRoleID, argUserID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddFedUserRoleMapping add to add a single record to fed_user_role_mapping table in the keycloak database
// @Summary Add an record to fed_user_role_mapping table
// @Description add to add a single record to fed_user_role_mapping table in the keycloak database
// @Tags FedUserRoleMapping
// @Accept  json
// @Produce  json
// @Param FedUserRoleMapping body model.FedUserRoleMapping true "Add FedUserRoleMapping"
// @Success 200 {object} model.FedUserRoleMapping
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /feduserrolemapping [post]
// echo '{"role_id": "aAZBoFRSRosYdsEuWgWLOqXEI","user_id": "QEJjxqeXMiRNpalYcWcHARqXa","realm_id": "vHioeFiowmNxTerZQIAbmjNry","storage_provider_id": "IUQfGYyQCfpAyoyKvCdspdIbb"}' | http POST "http://localhost:8080/feduserrolemapping" X-Api-User:user123
func AddFedUserRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	feduserrolemapping := &model.FedUserRoleMapping{}

	if err := readJSON(r, feduserrolemapping); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := feduserrolemapping.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	feduserrolemapping.Prepare()

	if err := feduserrolemapping.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_role_mapping", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	feduserrolemapping, _, err = dao.AddFedUserRoleMapping(ctx, feduserrolemapping)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, feduserrolemapping)
}

// UpdateFedUserRoleMapping Update a single record from fed_user_role_mapping table in the keycloak database
// @Summary Update an record in table fed_user_role_mapping
// @Description Update a single record from fed_user_role_mapping table in the keycloak database
// @Tags FedUserRoleMapping
// @Accept  json
// @Produce  json
// @Param  argRoleID path string true "role_id"// @Param  argUserID path string true "user_id"
// @Param  FedUserRoleMapping body model.FedUserRoleMapping true "Update FedUserRoleMapping record"
// @Success 200 {object} model.FedUserRoleMapping
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /feduserrolemapping/{argRoleID}/{argUserID} [put]
// echo '{"role_id": "aAZBoFRSRosYdsEuWgWLOqXEI","user_id": "QEJjxqeXMiRNpalYcWcHARqXa","realm_id": "vHioeFiowmNxTerZQIAbmjNry","storage_provider_id": "IUQfGYyQCfpAyoyKvCdspdIbb"}' | http PUT "http://localhost:8080/feduserrolemapping/hello world/hello world"  X-Api-User:user123
func UpdateFedUserRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRoleID, err := parseString(ps, "argRoleID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argUserID, err := parseString(ps, "argUserID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	feduserrolemapping := &model.FedUserRoleMapping{}
	if err := readJSON(r, feduserrolemapping); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := feduserrolemapping.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	feduserrolemapping.Prepare()

	if err := feduserrolemapping.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_role_mapping", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	feduserrolemapping, _, err = dao.UpdateFedUserRoleMapping(ctx,
		argRoleID, argUserID,
		feduserrolemapping)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, feduserrolemapping)
}

// DeleteFedUserRoleMapping Delete a single record from fed_user_role_mapping table in the keycloak database
// @Summary Delete a record from fed_user_role_mapping
// @Description Delete a single record from fed_user_role_mapping table in the keycloak database
// @Tags FedUserRoleMapping
// @Accept  json
// @Produce  json
// @Param  argRoleID path string true "role_id"// @Param  argUserID path string true "user_id"
// @Success 204 {object} model.FedUserRoleMapping
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /feduserrolemapping/{argRoleID}/{argUserID} [delete]
// http DELETE "http://localhost:8080/feduserrolemapping/hello world/hello world" X-Api-User:user123
func DeleteFedUserRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRoleID, err := parseString(ps, "argRoleID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argUserID, err := parseString(ps, "argUserID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_role_mapping", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteFedUserRoleMapping(ctx, argRoleID, argUserID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
