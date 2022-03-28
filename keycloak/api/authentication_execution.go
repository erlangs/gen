package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configAuthenticationExecutionRouter(router *httprouter.Router) {
	router.GET("/authenticationexecution", GetAllAuthenticationExecution)
	router.POST("/authenticationexecution", AddAuthenticationExecution)
	router.GET("/authenticationexecution/:argID", GetAuthenticationExecution)
	router.PUT("/authenticationexecution/:argID", UpdateAuthenticationExecution)
	router.DELETE("/authenticationexecution/:argID", DeleteAuthenticationExecution)
}

func configGinAuthenticationExecutionRouter(router gin.IRoutes) {
	router.GET("/authenticationexecution", ConverHttprouterToGin(GetAllAuthenticationExecution))
	router.POST("/authenticationexecution", ConverHttprouterToGin(AddAuthenticationExecution))
	router.GET("/authenticationexecution/:argID", ConverHttprouterToGin(GetAuthenticationExecution))
	router.PUT("/authenticationexecution/:argID", ConverHttprouterToGin(UpdateAuthenticationExecution))
	router.DELETE("/authenticationexecution/:argID", ConverHttprouterToGin(DeleteAuthenticationExecution))
}

// GetAllAuthenticationExecution is a function to get a slice of record(s) from authentication_execution table in the keycloak database
// @Summary Get list of AuthenticationExecution
// @Tags AuthenticationExecution
// @Description GetAllAuthenticationExecution is a handler to get a slice of record(s) from authentication_execution table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.AuthenticationExecution}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /authenticationexecution [get]
// http "http://localhost:8080/authenticationexecution?page=0&pagesize=20" X-Api-User:user123
func GetAllAuthenticationExecution(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "authentication_execution", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllAuthenticationExecution(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetAuthenticationExecution is a function to get a single record from the authentication_execution table in the keycloak database
// @Summary Get record from table AuthenticationExecution by  argID
// @Tags AuthenticationExecution
// @ID argID
// @Description GetAuthenticationExecution is a function to get a single record from the authentication_execution table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.AuthenticationExecution
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /authenticationexecution/{argID} [get]
// http "http://localhost:8080/authenticationexecution/hello world" X-Api-User:user123
func GetAuthenticationExecution(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "authentication_execution", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetAuthenticationExecution(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddAuthenticationExecution add to add a single record to authentication_execution table in the keycloak database
// @Summary Add an record to authentication_execution table
// @Description add to add a single record to authentication_execution table in the keycloak database
// @Tags AuthenticationExecution
// @Accept  json
// @Produce  json
// @Param AuthenticationExecution body model.AuthenticationExecution true "Add AuthenticationExecution"
// @Success 200 {object} model.AuthenticationExecution
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /authenticationexecution [post]
// echo '{"id": "kHjcZESkGJcnKMDlLyGJObfFe","alias": "CmMuHaSWyGkWiOMkILPeyBSMU","authenticator": "CaYxPSvWPvdKFwTHPPMPDZAiC","realm_id": "cWPRVmKTKZCXdAIOokPdbENLk","flow_id": "epifXlgjJobfAonQrZYknTNJG","requirement": 11,"priority": 45,"authenticator_flow": false,"auth_flow_id": "OgRaatruOvfCjGnnrFYtsaODy","auth_config": "KiOjMwBwAsKhrZNyQmULBIuFO"}' | http POST "http://localhost:8080/authenticationexecution" X-Api-User:user123
func AddAuthenticationExecution(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	authenticationexecution := &model.AuthenticationExecution{}

	if err := readJSON(r, authenticationexecution); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := authenticationexecution.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	authenticationexecution.Prepare()

	if err := authenticationexecution.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "authentication_execution", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	authenticationexecution, _, err = dao.AddAuthenticationExecution(ctx, authenticationexecution)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, authenticationexecution)
}

// UpdateAuthenticationExecution Update a single record from authentication_execution table in the keycloak database
// @Summary Update an record in table authentication_execution
// @Description Update a single record from authentication_execution table in the keycloak database
// @Tags AuthenticationExecution
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  AuthenticationExecution body model.AuthenticationExecution true "Update AuthenticationExecution record"
// @Success 200 {object} model.AuthenticationExecution
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /authenticationexecution/{argID} [put]
// echo '{"id": "kHjcZESkGJcnKMDlLyGJObfFe","alias": "CmMuHaSWyGkWiOMkILPeyBSMU","authenticator": "CaYxPSvWPvdKFwTHPPMPDZAiC","realm_id": "cWPRVmKTKZCXdAIOokPdbENLk","flow_id": "epifXlgjJobfAonQrZYknTNJG","requirement": 11,"priority": 45,"authenticator_flow": false,"auth_flow_id": "OgRaatruOvfCjGnnrFYtsaODy","auth_config": "KiOjMwBwAsKhrZNyQmULBIuFO"}' | http PUT "http://localhost:8080/authenticationexecution/hello world"  X-Api-User:user123
func UpdateAuthenticationExecution(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	authenticationexecution := &model.AuthenticationExecution{}
	if err := readJSON(r, authenticationexecution); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := authenticationexecution.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	authenticationexecution.Prepare()

	if err := authenticationexecution.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "authentication_execution", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	authenticationexecution, _, err = dao.UpdateAuthenticationExecution(ctx,
		argID,
		authenticationexecution)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, authenticationexecution)
}

// DeleteAuthenticationExecution Delete a single record from authentication_execution table in the keycloak database
// @Summary Delete a record from authentication_execution
// @Description Delete a single record from authentication_execution table in the keycloak database
// @Tags AuthenticationExecution
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.AuthenticationExecution
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /authenticationexecution/{argID} [delete]
// http DELETE "http://localhost:8080/authenticationexecution/hello world" X-Api-User:user123
func DeleteAuthenticationExecution(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "authentication_execution", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteAuthenticationExecution(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
