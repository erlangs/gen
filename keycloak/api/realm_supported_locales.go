package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configRealmSupportedLocalesRouter(router *httprouter.Router) {
	router.GET("/realmsupportedlocales", GetAllRealmSupportedLocales)
	router.POST("/realmsupportedlocales", AddRealmSupportedLocales)
	router.GET("/realmsupportedlocales/:argRealmID/:argValue", GetRealmSupportedLocales)
	router.PUT("/realmsupportedlocales/:argRealmID/:argValue", UpdateRealmSupportedLocales)
	router.DELETE("/realmsupportedlocales/:argRealmID/:argValue", DeleteRealmSupportedLocales)
}

func configGinRealmSupportedLocalesRouter(router gin.IRoutes) {
	router.GET("/realmsupportedlocales", ConverHttprouterToGin(GetAllRealmSupportedLocales))
	router.POST("/realmsupportedlocales", ConverHttprouterToGin(AddRealmSupportedLocales))
	router.GET("/realmsupportedlocales/:argRealmID/:argValue", ConverHttprouterToGin(GetRealmSupportedLocales))
	router.PUT("/realmsupportedlocales/:argRealmID/:argValue", ConverHttprouterToGin(UpdateRealmSupportedLocales))
	router.DELETE("/realmsupportedlocales/:argRealmID/:argValue", ConverHttprouterToGin(DeleteRealmSupportedLocales))
}

// GetAllRealmSupportedLocales is a function to get a slice of record(s) from realm_supported_locales table in the keycloak database
// @Summary Get list of RealmSupportedLocales
// @Tags RealmSupportedLocales
// @Description GetAllRealmSupportedLocales is a handler to get a slice of record(s) from realm_supported_locales table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.RealmSupportedLocales}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmsupportedlocales [get]
// http "http://localhost:8080/realmsupportedlocales?page=0&pagesize=20" X-Api-User:user123
func GetAllRealmSupportedLocales(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "realm_supported_locales", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllRealmSupportedLocales(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetRealmSupportedLocales is a function to get a single record from the realm_supported_locales table in the keycloak database
// @Summary Get record from table RealmSupportedLocales by  argRealmID  argValue
// @Tags RealmSupportedLocales
// @ID argRealmID
// @ID argValue
// @Description GetRealmSupportedLocales is a function to get a single record from the realm_supported_locales table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"
// @Param  argValue path string true "value"
// @Success 200 {object} model.RealmSupportedLocales
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /realmsupportedlocales/{argRealmID}/{argValue} [get]
// http "http://localhost:8080/realmsupportedlocales/hello world/hello world" X-Api-User:user123
func GetRealmSupportedLocales(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argValue, err := parseString(ps, "argValue")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_supported_locales", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetRealmSupportedLocales(ctx, argRealmID, argValue)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddRealmSupportedLocales add to add a single record to realm_supported_locales table in the keycloak database
// @Summary Add an record to realm_supported_locales table
// @Description add to add a single record to realm_supported_locales table in the keycloak database
// @Tags RealmSupportedLocales
// @Accept  json
// @Produce  json
// @Param RealmSupportedLocales body model.RealmSupportedLocales true "Add RealmSupportedLocales"
// @Success 200 {object} model.RealmSupportedLocales
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmsupportedlocales [post]
// echo '{"realm_id": "GjCeSTlTfVQIuektMNPSTDbdy","value": "eRIuTirtkHRKIOioRjvWJdmyB"}' | http POST "http://localhost:8080/realmsupportedlocales" X-Api-User:user123
func AddRealmSupportedLocales(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	realmsupportedlocales := &model.RealmSupportedLocales{}

	if err := readJSON(r, realmsupportedlocales); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := realmsupportedlocales.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	realmsupportedlocales.Prepare()

	if err := realmsupportedlocales.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_supported_locales", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	realmsupportedlocales, _, err = dao.AddRealmSupportedLocales(ctx, realmsupportedlocales)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, realmsupportedlocales)
}

// UpdateRealmSupportedLocales Update a single record from realm_supported_locales table in the keycloak database
// @Summary Update an record in table realm_supported_locales
// @Description Update a single record from realm_supported_locales table in the keycloak database
// @Tags RealmSupportedLocales
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"// @Param  argValue path string true "value"
// @Param  RealmSupportedLocales body model.RealmSupportedLocales true "Update RealmSupportedLocales record"
// @Success 200 {object} model.RealmSupportedLocales
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realmsupportedlocales/{argRealmID}/{argValue} [put]
// echo '{"realm_id": "GjCeSTlTfVQIuektMNPSTDbdy","value": "eRIuTirtkHRKIOioRjvWJdmyB"}' | http PUT "http://localhost:8080/realmsupportedlocales/hello world/hello world"  X-Api-User:user123
func UpdateRealmSupportedLocales(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argValue, err := parseString(ps, "argValue")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	realmsupportedlocales := &model.RealmSupportedLocales{}
	if err := readJSON(r, realmsupportedlocales); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := realmsupportedlocales.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	realmsupportedlocales.Prepare()

	if err := realmsupportedlocales.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_supported_locales", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	realmsupportedlocales, _, err = dao.UpdateRealmSupportedLocales(ctx,
		argRealmID, argValue,
		realmsupportedlocales)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, realmsupportedlocales)
}

// DeleteRealmSupportedLocales Delete a single record from realm_supported_locales table in the keycloak database
// @Summary Delete a record from realm_supported_locales
// @Description Delete a single record from realm_supported_locales table in the keycloak database
// @Tags RealmSupportedLocales
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"// @Param  argValue path string true "value"
// @Success 204 {object} model.RealmSupportedLocales
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /realmsupportedlocales/{argRealmID}/{argValue} [delete]
// http DELETE "http://localhost:8080/realmsupportedlocales/hello world/hello world" X-Api-User:user123
func DeleteRealmSupportedLocales(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argValue, err := parseString(ps, "argValue")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "realm_supported_locales", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteRealmSupportedLocales(ctx, argRealmID, argValue)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
