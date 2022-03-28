package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configFedUserConsentClScopeRouter(router *httprouter.Router) {
	router.GET("/feduserconsentclscope", GetAllFedUserConsentClScope)
	router.POST("/feduserconsentclscope", AddFedUserConsentClScope)
	router.GET("/feduserconsentclscope/:argUserConsentID/:argScopeID", GetFedUserConsentClScope)
	router.PUT("/feduserconsentclscope/:argUserConsentID/:argScopeID", UpdateFedUserConsentClScope)
	router.DELETE("/feduserconsentclscope/:argUserConsentID/:argScopeID", DeleteFedUserConsentClScope)
}

func configGinFedUserConsentClScopeRouter(router gin.IRoutes) {
	router.GET("/feduserconsentclscope", ConverHttprouterToGin(GetAllFedUserConsentClScope))
	router.POST("/feduserconsentclscope", ConverHttprouterToGin(AddFedUserConsentClScope))
	router.GET("/feduserconsentclscope/:argUserConsentID/:argScopeID", ConverHttprouterToGin(GetFedUserConsentClScope))
	router.PUT("/feduserconsentclscope/:argUserConsentID/:argScopeID", ConverHttprouterToGin(UpdateFedUserConsentClScope))
	router.DELETE("/feduserconsentclscope/:argUserConsentID/:argScopeID", ConverHttprouterToGin(DeleteFedUserConsentClScope))
}

// GetAllFedUserConsentClScope is a function to get a slice of record(s) from fed_user_consent_cl_scope table in the keycloak database
// @Summary Get list of FedUserConsentClScope
// @Tags FedUserConsentClScope
// @Description GetAllFedUserConsentClScope is a handler to get a slice of record(s) from fed_user_consent_cl_scope table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.FedUserConsentClScope}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /feduserconsentclscope [get]
// http "http://localhost:8080/feduserconsentclscope?page=0&pagesize=20" X-Api-User:user123
func GetAllFedUserConsentClScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "fed_user_consent_cl_scope", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllFedUserConsentClScope(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetFedUserConsentClScope is a function to get a single record from the fed_user_consent_cl_scope table in the keycloak database
// @Summary Get record from table FedUserConsentClScope by  argUserConsentID  argScopeID
// @Tags FedUserConsentClScope
// @ID argUserConsentID
// @ID argScopeID
// @Description GetFedUserConsentClScope is a function to get a single record from the fed_user_consent_cl_scope table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argUserConsentID path string true "user_consent_id"
// @Param  argScopeID path string true "scope_id"
// @Success 200 {object} model.FedUserConsentClScope
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /feduserconsentclscope/{argUserConsentID}/{argScopeID} [get]
// http "http://localhost:8080/feduserconsentclscope/hello world/hello world" X-Api-User:user123
func GetFedUserConsentClScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "fed_user_consent_cl_scope", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetFedUserConsentClScope(ctx, argUserConsentID, argScopeID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddFedUserConsentClScope add to add a single record to fed_user_consent_cl_scope table in the keycloak database
// @Summary Add an record to fed_user_consent_cl_scope table
// @Description add to add a single record to fed_user_consent_cl_scope table in the keycloak database
// @Tags FedUserConsentClScope
// @Accept  json
// @Produce  json
// @Param FedUserConsentClScope body model.FedUserConsentClScope true "Add FedUserConsentClScope"
// @Success 200 {object} model.FedUserConsentClScope
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /feduserconsentclscope [post]
// echo '{"user_consent_id": "dkZDOStxsEhjFevBCjqEReHnX","scope_id": "qvgcgYtmmxJvKwelRBjxmUlwb"}' | http POST "http://localhost:8080/feduserconsentclscope" X-Api-User:user123
func AddFedUserConsentClScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	feduserconsentclscope := &model.FedUserConsentClScope{}

	if err := readJSON(r, feduserconsentclscope); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := feduserconsentclscope.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	feduserconsentclscope.Prepare()

	if err := feduserconsentclscope.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_consent_cl_scope", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	feduserconsentclscope, _, err = dao.AddFedUserConsentClScope(ctx, feduserconsentclscope)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, feduserconsentclscope)
}

// UpdateFedUserConsentClScope Update a single record from fed_user_consent_cl_scope table in the keycloak database
// @Summary Update an record in table fed_user_consent_cl_scope
// @Description Update a single record from fed_user_consent_cl_scope table in the keycloak database
// @Tags FedUserConsentClScope
// @Accept  json
// @Produce  json
// @Param  argUserConsentID path string true "user_consent_id"// @Param  argScopeID path string true "scope_id"
// @Param  FedUserConsentClScope body model.FedUserConsentClScope true "Update FedUserConsentClScope record"
// @Success 200 {object} model.FedUserConsentClScope
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /feduserconsentclscope/{argUserConsentID}/{argScopeID} [put]
// echo '{"user_consent_id": "dkZDOStxsEhjFevBCjqEReHnX","scope_id": "qvgcgYtmmxJvKwelRBjxmUlwb"}' | http PUT "http://localhost:8080/feduserconsentclscope/hello world/hello world"  X-Api-User:user123
func UpdateFedUserConsentClScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	feduserconsentclscope := &model.FedUserConsentClScope{}
	if err := readJSON(r, feduserconsentclscope); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := feduserconsentclscope.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	feduserconsentclscope.Prepare()

	if err := feduserconsentclscope.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_consent_cl_scope", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	feduserconsentclscope, _, err = dao.UpdateFedUserConsentClScope(ctx,
		argUserConsentID, argScopeID,
		feduserconsentclscope)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, feduserconsentclscope)
}

// DeleteFedUserConsentClScope Delete a single record from fed_user_consent_cl_scope table in the keycloak database
// @Summary Delete a record from fed_user_consent_cl_scope
// @Description Delete a single record from fed_user_consent_cl_scope table in the keycloak database
// @Tags FedUserConsentClScope
// @Accept  json
// @Produce  json
// @Param  argUserConsentID path string true "user_consent_id"// @Param  argScopeID path string true "scope_id"
// @Success 204 {object} model.FedUserConsentClScope
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /feduserconsentclscope/{argUserConsentID}/{argScopeID} [delete]
// http DELETE "http://localhost:8080/feduserconsentclscope/hello world/hello world" X-Api-User:user123
func DeleteFedUserConsentClScope(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "fed_user_consent_cl_scope", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteFedUserConsentClScope(ctx, argUserConsentID, argScopeID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
