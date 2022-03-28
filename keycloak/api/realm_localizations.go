package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configRealmLocalizationsRouter(router *httprouter.Router) {
	router.GET("/realmlocalizations", GetAllRealmLocalizations)
	router.POST("/realmlocalizations", AddRealmLocalizations)
	router.GET("/realmlocalizations/:argRealmID/:argLocale", GetRealmLocalizations)
	router.PUT("/realmlocalizations/:argRealmID/:argLocale", UpdateRealmLocalizations)
	router.DELETE("/realmlocalizations/:argRealmID/:argLocale", DeleteRealmLocalizations)
}

func configGinRealmLocalizationsRouter(router gin.IRoutes) {
	router.GET("/realmlocalizations", ConverHttprouterToGin(GetAllRealmLocalizations))
	router.POST("/realmlocalizations", ConverHttprouterToGin(AddRealmLocalizations))
	router.GET("/realmlocalizations/:argRealmID/:argLocale", ConverHttprouterToGin(GetRealmLocalizations))
	router.PUT("/realmlocalizations/:argRealmID/:argLocale", ConverHttprouterToGin(UpdateRealmLocalizations))
	router.DELETE("/realmlocalizations/:argRealmID/:argLocale", ConverHttprouterToGin(DeleteRealmLocalizations))
}

// GetAllRealmLocalizations is a function to get a slice of record(s) from realm_localizations table in the keycloak database
// @Summary Get list of RealmLocalizations
// @Tags RealmLocalizations
// @Description GetAllRealmLocalizations is a handler to get a slice of record(s) from realm_localizations table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.RealmLocalizations}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmlocalizations [get]
// http "http://localhost:8080/realmlocalizations?page=0&pagesize=20" X-Api-User:user123
func GetAllRealmLocalizations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "realm_localizations", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllRealmLocalizations(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetRealmLocalizations is a function to get a single record from the realm_localizations table in the keycloak database
// @Summary Get record from table RealmLocalizations by  argRealmID  argLocale
// @Tags RealmLocalizations
// @ID argRealmID
// @ID argLocale
// @Description GetRealmLocalizations is a function to get a single record from the realm_localizations table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"
// @Param  argLocale path string true "locale"
// @Success 200 {object} model.RealmLocalizations
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /realmlocalizations/{argRealmID}/{argLocale} [get]
// http "http://localhost:8080/realmlocalizations/hello world/hello world" X-Api-User:user123
func GetRealmLocalizations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argLocale, err := parseString(ps, "argLocale")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_localizations", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetRealmLocalizations(ctx, argRealmID, argLocale)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddRealmLocalizations add to add a single record to realm_localizations table in the keycloak database
// @Summary Add an record to realm_localizations table
// @Description add to add a single record to realm_localizations table in the keycloak database
// @Tags RealmLocalizations
// @Accept  json
// @Produce  json
// @Param RealmLocalizations body model.RealmLocalizations true "Add RealmLocalizations"
// @Success 200 {object} model.RealmLocalizations
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmlocalizations [post]
// echo '{"realm_id": "jCOXUkADnUMfFLdJxgBSvsUmN","locale": "FHLfdKtPYvhRhEtAZPoxOXpcZ","texts": "blTqIiVBtCHtyVxJbVgZxgqJk"}' | http POST "http://localhost:8080/realmlocalizations" X-Api-User:user123
func AddRealmLocalizations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	realmlocalizations := &model.RealmLocalizations{}

	if err := readJSON(r, realmlocalizations); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := realmlocalizations.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	realmlocalizations.Prepare()

	if err := realmlocalizations.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_localizations", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	realmlocalizations, _, err = dao.AddRealmLocalizations(ctx, realmlocalizations)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, realmlocalizations)
}

// UpdateRealmLocalizations Update a single record from realm_localizations table in the keycloak database
// @Summary Update an record in table realm_localizations
// @Description Update a single record from realm_localizations table in the keycloak database
// @Tags RealmLocalizations
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"// @Param  argLocale path string true "locale"
// @Param  RealmLocalizations body model.RealmLocalizations true "Update RealmLocalizations record"
// @Success 200 {object} model.RealmLocalizations
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmlocalizations/{argRealmID}/{argLocale} [put]
// echo '{"realm_id": "jCOXUkADnUMfFLdJxgBSvsUmN","locale": "FHLfdKtPYvhRhEtAZPoxOXpcZ","texts": "blTqIiVBtCHtyVxJbVgZxgqJk"}' | http PUT "http://localhost:8080/realmlocalizations/hello world/hello world"  X-Api-User:user123
func UpdateRealmLocalizations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argLocale, err := parseString(ps, "argLocale")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	realmlocalizations := &model.RealmLocalizations{}
	if err := readJSON(r, realmlocalizations); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := realmlocalizations.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	realmlocalizations.Prepare()

	if err := realmlocalizations.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_localizations", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	realmlocalizations, _, err = dao.UpdateRealmLocalizations(ctx,
		argRealmID, argLocale,
		realmlocalizations)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, realmlocalizations)
}

// DeleteRealmLocalizations Delete a single record from realm_localizations table in the keycloak database
// @Summary Delete a record from realm_localizations
// @Description Delete a single record from realm_localizations table in the keycloak database
// @Tags RealmLocalizations
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"// @Param  argLocale path string true "locale"
// @Success 204 {object} model.RealmLocalizations
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /realmlocalizations/{argRealmID}/{argLocale} [delete]
// http DELETE "http://localhost:8080/realmlocalizations/hello world/hello world" X-Api-User:user123
func DeleteRealmLocalizations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argLocale, err := parseString(ps, "argLocale")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_localizations", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteRealmLocalizations(ctx, argRealmID, argLocale)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
