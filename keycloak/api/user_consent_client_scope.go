package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configUserConsentClientScopeRouter(router *httprouter.Router) {
	router.GET("/userconsentclientscope", GetAllUserConsentClientScope)
	router.POST("/userconsentclientscope", AddUserConsentClientScope)
	router.GET("/userconsentclientscope/:argUserConsentID/:argScopeID", GetUserConsentClientScope)
	router.PUT("/userconsentclientscope/:argUserConsentID/:argScopeID", UpdateUserConsentClientScope)
	router.DELETE("/userconsentclientscope/:argUserConsentID/:argScopeID", DeleteUserConsentClientScope)
}

func configGinUserConsentClientScopeRouter(router gin.IRoutes) {
	router.GET("/userconsentclientscope", ConverHttprouterToGin(GetAllUserConsentClientScope))
	router.POST("/userconsentclientscope", ConverHttprouterToGin(AddUserConsentClientScope))
	router.GET("/userconsentclientscope/:argUserConsentID/:argScopeID", ConverHttprouterToGin(GetUserConsentClientScope))
	router.PUT("/userconsentclientscope/:argUserConsentID/:argScopeID", ConverHttprouterToGin(UpdateUserConsentClientScope))
	router.DELETE("/userconsentclientscope/:argUserConsentID/:argScopeID", ConverHttprouterToGin(DeleteUserConsentClientScope))
}

// GetAllUserConsentClientScope is a function to get a slice of record(s) from user_consent_client_scope table in the keycloak database
// @Summary Get list of UserConsentClientScope
// @Tags UserConsentClientScope
// @Description GetAllUserConsentClientScope is a handler to get a slice of record(s) from user_consent_client_scope table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.UserConsentClientScope}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userconsentclientscope [get]
// http "http://localhost:8080/userconsentclientscope?page=0&pagesize=20" X-Api-User:user123
func GetAllUserConsentClientScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "user_consent_client_scope", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllUserConsentClientScope(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetUserConsentClientScope is a function to get a single record from the user_consent_client_scope table in the keycloak database
// @Summary Get record from table UserConsentClientScope by  argUserConsentID  argScopeID
// @Tags UserConsentClientScope
// @ID argUserConsentID
// @ID argScopeID
// @Description GetUserConsentClientScope is a function to get a single record from the user_consent_client_scope table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argUserConsentID path string true "user_consent_id"
// @Param  argScopeID path string true "scope_id"
// @Success 200 {object} model.UserConsentClientScope
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /userconsentclientscope/{argUserConsentID}/{argScopeID} [get]
// http "http://localhost:8080/userconsentclientscope/hello world/hello world" X-Api-User:user123
func GetUserConsentClientScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserConsentID, err := parseString(ps, "argUserConsentID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_consent_client_scope", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetUserConsentClientScope(ctx, argUserConsentID, argScopeID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddUserConsentClientScope add to add a single record to user_consent_client_scope table in the keycloak database
// @Summary Add an record to user_consent_client_scope table
// @Description add to add a single record to user_consent_client_scope table in the keycloak database
// @Tags UserConsentClientScope
// @Accept  json
// @Produce  json
// @Param UserConsentClientScope body model.UserConsentClientScope true "Add UserConsentClientScope"
// @Success 200 {object} model.UserConsentClientScope
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userconsentclientscope [post]
// echo '{"user_consent_id": "NDmTGUHGWPXDRhkrvqycvCSyy","scope_id": "hojjxUVWyZStIOxXqThwGLVIY"}' | http POST "http://localhost:8080/userconsentclientscope" X-Api-User:user123
func AddUserConsentClientScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	userconsentclientscope := &model.UserConsentClientScope{}

	if err := readJSON(r, userconsentclientscope); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userconsentclientscope.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userconsentclientscope.Prepare()

	if err := userconsentclientscope.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_consent_client_scope", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	userconsentclientscope, _, err = dao.AddUserConsentClientScope(ctx, userconsentclientscope)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userconsentclientscope)
}

// UpdateUserConsentClientScope Update a single record from user_consent_client_scope table in the keycloak database
// @Summary Update an record in table user_consent_client_scope
// @Description Update a single record from user_consent_client_scope table in the keycloak database
// @Tags UserConsentClientScope
// @Accept  json
// @Produce  json
// @Param  argUserConsentID path string true "user_consent_id"// @Param  argScopeID path string true "scope_id"
// @Param  UserConsentClientScope body model.UserConsentClientScope true "Update UserConsentClientScope record"
// @Success 200 {object} model.UserConsentClientScope
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userconsentclientscope/{argUserConsentID}/{argScopeID} [put]
// echo '{"user_consent_id": "NDmTGUHGWPXDRhkrvqycvCSyy","scope_id": "hojjxUVWyZStIOxXqThwGLVIY"}' | http PUT "http://localhost:8080/userconsentclientscope/hello world/hello world"  X-Api-User:user123
func UpdateUserConsentClientScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserConsentID, err := parseString(ps, "argUserConsentID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	userconsentclientscope := &model.UserConsentClientScope{}
	if err := readJSON(r, userconsentclientscope); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userconsentclientscope.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userconsentclientscope.Prepare()

	if err := userconsentclientscope.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_consent_client_scope", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	userconsentclientscope, _, err = dao.UpdateUserConsentClientScope(ctx,
		argUserConsentID, argScopeID,
		userconsentclientscope)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userconsentclientscope)
}

// DeleteUserConsentClientScope Delete a single record from user_consent_client_scope table in the keycloak database
// @Summary Delete a record from user_consent_client_scope
// @Description Delete a single record from user_consent_client_scope table in the keycloak database
// @Tags UserConsentClientScope
// @Accept  json
// @Produce  json
// @Param  argUserConsentID path string true "user_consent_id"// @Param  argScopeID path string true "scope_id"
// @Success 204 {object} model.UserConsentClientScope
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /userconsentclientscope/{argUserConsentID}/{argScopeID} [delete]
// http DELETE "http://localhost:8080/userconsentclientscope/hello world/hello world" X-Api-User:user123
func DeleteUserConsentClientScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserConsentID, err := parseString(ps, "argUserConsentID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_consent_client_scope", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteUserConsentClientScope(ctx, argUserConsentID, argScopeID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
