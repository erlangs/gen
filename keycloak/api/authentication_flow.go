package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configAuthenticationFlowRouter(router *httprouter.Router) {
	router.GET("/authenticationflow", GetAllAuthenticationFlow)
	router.POST("/authenticationflow", AddAuthenticationFlow)
	router.GET("/authenticationflow/:argID", GetAuthenticationFlow)
	router.PUT("/authenticationflow/:argID", UpdateAuthenticationFlow)
	router.DELETE("/authenticationflow/:argID", DeleteAuthenticationFlow)
}

func configGinAuthenticationFlowRouter(router gin.IRoutes) {
	router.GET("/authenticationflow", ConverHttprouterToGin(GetAllAuthenticationFlow))
	router.POST("/authenticationflow", ConverHttprouterToGin(AddAuthenticationFlow))
	router.GET("/authenticationflow/:argID", ConverHttprouterToGin(GetAuthenticationFlow))
	router.PUT("/authenticationflow/:argID", ConverHttprouterToGin(UpdateAuthenticationFlow))
	router.DELETE("/authenticationflow/:argID", ConverHttprouterToGin(DeleteAuthenticationFlow))
}

// GetAllAuthenticationFlow is a function to get a slice of record(s) from authentication_flow table in the keycloak database
// @Summary Get list of AuthenticationFlow
// @Tags AuthenticationFlow
// @Description GetAllAuthenticationFlow is a handler to get a slice of record(s) from authentication_flow table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.AuthenticationFlow}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /authenticationflow [get]
// http "http://localhost:8080/authenticationflow?page=0&pagesize=20" X-Api-User:user123
func GetAllAuthenticationFlow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "authentication_flow", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllAuthenticationFlow(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetAuthenticationFlow is a function to get a single record from the authentication_flow table in the keycloak database
// @Summary Get record from table AuthenticationFlow by  argID
// @Tags AuthenticationFlow
// @ID argID
// @Description GetAuthenticationFlow is a function to get a single record from the authentication_flow table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.AuthenticationFlow
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /authenticationflow/{argID} [get]
// http "http://localhost:8080/authenticationflow/hello world" X-Api-User:user123
func GetAuthenticationFlow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "authentication_flow", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetAuthenticationFlow(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddAuthenticationFlow add to add a single record to authentication_flow table in the keycloak database
// @Summary Add an record to authentication_flow table
// @Description add to add a single record to authentication_flow table in the keycloak database
// @Tags AuthenticationFlow
// @Accept  json
// @Produce  json
// @Param AuthenticationFlow body model.AuthenticationFlow true "Add AuthenticationFlow"
// @Success 200 {object} model.AuthenticationFlow
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /authenticationflow [post]
// echo '{"id": "JRBISDiAIeFXZDYppSuaSMQLI","alias": "BghfHSoCiQnIKapipeJciqypc","description": "XRByivSyHhhJsUXeSQgXafhDH","realm_id": "nAvcgQKqgeESZaTexDJuErHHl","provider_id": "TVnIDdhtFQPJrViCHqayUsKdo","top_level": false,"built_in": true}' | http POST "http://localhost:8080/authenticationflow" X-Api-User:user123
func AddAuthenticationFlow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	authenticationflow := &model.AuthenticationFlow{}

	if err := readJSON(r, authenticationflow); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := authenticationflow.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	authenticationflow.Prepare()

	if err := authenticationflow.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "authentication_flow", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	authenticationflow, _, err = dao.AddAuthenticationFlow(ctx, authenticationflow)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, authenticationflow)
}

// UpdateAuthenticationFlow Update a single record from authentication_flow table in the keycloak database
// @Summary Update an record in table authentication_flow
// @Description Update a single record from authentication_flow table in the keycloak database
// @Tags AuthenticationFlow
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  AuthenticationFlow body model.AuthenticationFlow true "Update AuthenticationFlow record"
// @Success 200 {object} model.AuthenticationFlow
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /authenticationflow/{argID} [put]
// echo '{"id": "JRBISDiAIeFXZDYppSuaSMQLI","alias": "BghfHSoCiQnIKapipeJciqypc","description": "XRByivSyHhhJsUXeSQgXafhDH","realm_id": "nAvcgQKqgeESZaTexDJuErHHl","provider_id": "TVnIDdhtFQPJrViCHqayUsKdo","top_level": false,"built_in": true}' | http PUT "http://localhost:8080/authenticationflow/hello world"  X-Api-User:user123
func UpdateAuthenticationFlow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	authenticationflow := &model.AuthenticationFlow{}
	if err := readJSON(r, authenticationflow); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := authenticationflow.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	authenticationflow.Prepare()

	if err := authenticationflow.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "authentication_flow", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	authenticationflow, _, err = dao.UpdateAuthenticationFlow(ctx,
		argID,
		authenticationflow)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, authenticationflow)
}

// DeleteAuthenticationFlow Delete a single record from authentication_flow table in the keycloak database
// @Summary Delete a record from authentication_flow
// @Description Delete a single record from authentication_flow table in the keycloak database
// @Tags AuthenticationFlow
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.AuthenticationFlow
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /authenticationflow/{argID} [delete]
// http DELETE "http://localhost:8080/authenticationflow/hello world" X-Api-User:user123
func DeleteAuthenticationFlow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "authentication_flow", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteAuthenticationFlow(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
