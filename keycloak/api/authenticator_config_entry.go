package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configAuthenticatorConfigEntryRouter(router *httprouter.Router) {
	router.GET("/authenticatorconfigentry", GetAllAuthenticatorConfigEntry)
	router.POST("/authenticatorconfigentry", AddAuthenticatorConfigEntry)
	router.GET("/authenticatorconfigentry/:argAuthenticatorID/:argName", GetAuthenticatorConfigEntry)
	router.PUT("/authenticatorconfigentry/:argAuthenticatorID/:argName", UpdateAuthenticatorConfigEntry)
	router.DELETE("/authenticatorconfigentry/:argAuthenticatorID/:argName", DeleteAuthenticatorConfigEntry)
}

func configGinAuthenticatorConfigEntryRouter(router gin.IRoutes) {
	router.GET("/authenticatorconfigentry", ConverHttprouterToGin(GetAllAuthenticatorConfigEntry))
	router.POST("/authenticatorconfigentry", ConverHttprouterToGin(AddAuthenticatorConfigEntry))
	router.GET("/authenticatorconfigentry/:argAuthenticatorID/:argName", ConverHttprouterToGin(GetAuthenticatorConfigEntry))
	router.PUT("/authenticatorconfigentry/:argAuthenticatorID/:argName", ConverHttprouterToGin(UpdateAuthenticatorConfigEntry))
	router.DELETE("/authenticatorconfigentry/:argAuthenticatorID/:argName", ConverHttprouterToGin(DeleteAuthenticatorConfigEntry))
}

// GetAllAuthenticatorConfigEntry is a function to get a slice of record(s) from authenticator_config_entry table in the keycloak database
// @Summary Get list of AuthenticatorConfigEntry
// @Tags AuthenticatorConfigEntry
// @Description GetAllAuthenticatorConfigEntry is a handler to get a slice of record(s) from authenticator_config_entry table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.AuthenticatorConfigEntry}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /authenticatorconfigentry [get]
// http "http://localhost:8080/authenticatorconfigentry?page=0&pagesize=20" X-Api-User:user123
func GetAllAuthenticatorConfigEntry(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "authenticator_config_entry", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllAuthenticatorConfigEntry(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetAuthenticatorConfigEntry is a function to get a single record from the authenticator_config_entry table in the keycloak database
// @Summary Get record from table AuthenticatorConfigEntry by  argAuthenticatorID  argName
// @Tags AuthenticatorConfigEntry
// @ID argAuthenticatorID
// @ID argName
// @Description GetAuthenticatorConfigEntry is a function to get a single record from the authenticator_config_entry table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argAuthenticatorID path string true "authenticator_id"
// @Param  argName path string true "name"
// @Success 200 {object} model.AuthenticatorConfigEntry
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /authenticatorconfigentry/{argAuthenticatorID}/{argName} [get]
// http "http://localhost:8080/authenticatorconfigentry/hello world/hello world" X-Api-User:user123
func GetAuthenticatorConfigEntry(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argAuthenticatorID, err := parseString(ps, "argAuthenticatorID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "authenticator_config_entry", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetAuthenticatorConfigEntry(ctx, argAuthenticatorID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddAuthenticatorConfigEntry add to add a single record to authenticator_config_entry table in the keycloak database
// @Summary Add an record to authenticator_config_entry table
// @Description add to add a single record to authenticator_config_entry table in the keycloak database
// @Tags AuthenticatorConfigEntry
// @Accept  json
// @Produce  json
// @Param AuthenticatorConfigEntry body model.AuthenticatorConfigEntry true "Add AuthenticatorConfigEntry"
// @Success 200 {object} model.AuthenticatorConfigEntry
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /authenticatorconfigentry [post]
// echo '{"authenticator_id": "sIQLihIRWGCDhURmKwXpseisN","value": "XVJYqwLtXevqCVJLONkZTMEjy","name": "hVsLuTemYWVdtRGnDIqbFshkE"}' | http POST "http://localhost:8080/authenticatorconfigentry" X-Api-User:user123
func AddAuthenticatorConfigEntry(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	authenticatorconfigentry := &model.AuthenticatorConfigEntry{}

	if err := readJSON(r, authenticatorconfigentry); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := authenticatorconfigentry.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	authenticatorconfigentry.Prepare()

	if err := authenticatorconfigentry.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "authenticator_config_entry", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	authenticatorconfigentry, _, err = dao.AddAuthenticatorConfigEntry(ctx, authenticatorconfigentry)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, authenticatorconfigentry)
}

// UpdateAuthenticatorConfigEntry Update a single record from authenticator_config_entry table in the keycloak database
// @Summary Update an record in table authenticator_config_entry
// @Description Update a single record from authenticator_config_entry table in the keycloak database
// @Tags AuthenticatorConfigEntry
// @Accept  json
// @Produce  json
// @Param  argAuthenticatorID path string true "authenticator_id"// @Param  argName path string true "name"
// @Param  AuthenticatorConfigEntry body model.AuthenticatorConfigEntry true "Update AuthenticatorConfigEntry record"
// @Success 200 {object} model.AuthenticatorConfigEntry
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /authenticatorconfigentry/{argAuthenticatorID}/{argName} [put]
// echo '{"authenticator_id": "sIQLihIRWGCDhURmKwXpseisN","value": "XVJYqwLtXevqCVJLONkZTMEjy","name": "hVsLuTemYWVdtRGnDIqbFshkE"}' | http PUT "http://localhost:8080/authenticatorconfigentry/hello world/hello world"  X-Api-User:user123
func UpdateAuthenticatorConfigEntry(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argAuthenticatorID, err := parseString(ps, "argAuthenticatorID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	authenticatorconfigentry := &model.AuthenticatorConfigEntry{}
	if err := readJSON(r, authenticatorconfigentry); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := authenticatorconfigentry.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	authenticatorconfigentry.Prepare()

	if err := authenticatorconfigentry.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "authenticator_config_entry", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	authenticatorconfigentry, _, err = dao.UpdateAuthenticatorConfigEntry(ctx,
		argAuthenticatorID, argName,
		authenticatorconfigentry)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, authenticatorconfigentry)
}

// DeleteAuthenticatorConfigEntry Delete a single record from authenticator_config_entry table in the keycloak database
// @Summary Delete a record from authenticator_config_entry
// @Description Delete a single record from authenticator_config_entry table in the keycloak database
// @Tags AuthenticatorConfigEntry
// @Accept  json
// @Produce  json
// @Param  argAuthenticatorID path string true "authenticator_id"// @Param  argName path string true "name"
// @Success 204 {object} model.AuthenticatorConfigEntry
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /authenticatorconfigentry/{argAuthenticatorID}/{argName} [delete]
// http DELETE "http://localhost:8080/authenticatorconfigentry/hello world/hello world" X-Api-User:user123
func DeleteAuthenticatorConfigEntry(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argAuthenticatorID, err := parseString(ps, "argAuthenticatorID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "authenticator_config_entry", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteAuthenticatorConfigEntry(ctx, argAuthenticatorID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
