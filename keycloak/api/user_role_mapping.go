package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configUserRoleMappingRouter(router *httprouter.Router) {
	router.GET("/userrolemapping", GetAllUserRoleMapping)
	router.POST("/userrolemapping", AddUserRoleMapping)
	router.GET("/userrolemapping/:argRoleID/:argUserID", GetUserRoleMapping)
	router.PUT("/userrolemapping/:argRoleID/:argUserID", UpdateUserRoleMapping)
	router.DELETE("/userrolemapping/:argRoleID/:argUserID", DeleteUserRoleMapping)
}

func configGinUserRoleMappingRouter(router gin.IRoutes) {
	router.GET("/userrolemapping", ConverHttprouterToGin(GetAllUserRoleMapping))
	router.POST("/userrolemapping", ConverHttprouterToGin(AddUserRoleMapping))
	router.GET("/userrolemapping/:argRoleID/:argUserID", ConverHttprouterToGin(GetUserRoleMapping))
	router.PUT("/userrolemapping/:argRoleID/:argUserID", ConverHttprouterToGin(UpdateUserRoleMapping))
	router.DELETE("/userrolemapping/:argRoleID/:argUserID", ConverHttprouterToGin(DeleteUserRoleMapping))
}

// GetAllUserRoleMapping is a function to get a slice of record(s) from user_role_mapping table in the keycloak database
// @Summary Get list of UserRoleMapping
// @Tags UserRoleMapping
// @Description GetAllUserRoleMapping is a handler to get a slice of record(s) from user_role_mapping table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.UserRoleMapping}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userrolemapping [get]
// http "http://localhost:8080/userrolemapping?page=0&pagesize=20" X-Api-User:user123
func GetAllUserRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "user_role_mapping", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllUserRoleMapping(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetUserRoleMapping is a function to get a single record from the user_role_mapping table in the keycloak database
// @Summary Get record from table UserRoleMapping by  argRoleID  argUserID
// @Tags UserRoleMapping
// @ID argRoleID
// @ID argUserID
// @Description GetUserRoleMapping is a function to get a single record from the user_role_mapping table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argRoleID path string true "role_id"
// @Param  argUserID path string true "user_id"
// @Success 200 {object} model.UserRoleMapping
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /userrolemapping/{argRoleID}/{argUserID} [get]
// http "http://localhost:8080/userrolemapping/hello world/hello world" X-Api-User:user123
func GetUserRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "user_role_mapping", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetUserRoleMapping(ctx, argRoleID, argUserID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddUserRoleMapping add to add a single record to user_role_mapping table in the keycloak database
// @Summary Add an record to user_role_mapping table
// @Description add to add a single record to user_role_mapping table in the keycloak database
// @Tags UserRoleMapping
// @Accept  json
// @Produce  json
// @Param UserRoleMapping body model.UserRoleMapping true "Add UserRoleMapping"
// @Success 200 {object} model.UserRoleMapping
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userrolemapping [post]
// echo '{"role_id": "QHAelmhZFaaHBXfUDwfYgNHNA","user_id": "fIFkfLqgMgVSMZfsOlyCZaGZy"}' | http POST "http://localhost:8080/userrolemapping" X-Api-User:user123
func AddUserRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	userrolemapping := &model.UserRoleMapping{}

	if err := readJSON(r, userrolemapping); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userrolemapping.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userrolemapping.Prepare()

	if err := userrolemapping.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_role_mapping", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	userrolemapping, _, err = dao.AddUserRoleMapping(ctx, userrolemapping)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userrolemapping)
}

// UpdateUserRoleMapping Update a single record from user_role_mapping table in the keycloak database
// @Summary Update an record in table user_role_mapping
// @Description Update a single record from user_role_mapping table in the keycloak database
// @Tags UserRoleMapping
// @Accept  json
// @Produce  json
// @Param  argRoleID path string true "role_id"// @Param  argUserID path string true "user_id"
// @Param  UserRoleMapping body model.UserRoleMapping true "Update UserRoleMapping record"
// @Success 200 {object} model.UserRoleMapping
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userrolemapping/{argRoleID}/{argUserID} [put]
// echo '{"role_id": "QHAelmhZFaaHBXfUDwfYgNHNA","user_id": "fIFkfLqgMgVSMZfsOlyCZaGZy"}' | http PUT "http://localhost:8080/userrolemapping/hello world/hello world"  X-Api-User:user123
func UpdateUserRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	userrolemapping := &model.UserRoleMapping{}
	if err := readJSON(r, userrolemapping); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userrolemapping.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userrolemapping.Prepare()

	if err := userrolemapping.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_role_mapping", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	userrolemapping, _, err = dao.UpdateUserRoleMapping(ctx,
		argRoleID, argUserID,
		userrolemapping)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userrolemapping)
}

// DeleteUserRoleMapping Delete a single record from user_role_mapping table in the keycloak database
// @Summary Delete a record from user_role_mapping
// @Description Delete a single record from user_role_mapping table in the keycloak database
// @Tags UserRoleMapping
// @Accept  json
// @Produce  json
// @Param  argRoleID path string true "role_id"// @Param  argUserID path string true "user_id"
// @Success 204 {object} model.UserRoleMapping
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /userrolemapping/{argRoleID}/{argUserID} [delete]
// http DELETE "http://localhost:8080/userrolemapping/hello world/hello world" X-Api-User:user123
func DeleteUserRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "user_role_mapping", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteUserRoleMapping(ctx, argRoleID, argUserID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
