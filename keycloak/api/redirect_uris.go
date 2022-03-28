package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configRedirectUrisRouter(router *httprouter.Router) {
	router.GET("/redirecturis", GetAllRedirectUris)
	router.POST("/redirecturis", AddRedirectUris)
	router.GET("/redirecturis/:argClientID/:argValue", GetRedirectUris)
	router.PUT("/redirecturis/:argClientID/:argValue", UpdateRedirectUris)
	router.DELETE("/redirecturis/:argClientID/:argValue", DeleteRedirectUris)
}

func configGinRedirectUrisRouter(router gin.IRoutes) {
	router.GET("/redirecturis", ConverHttprouterToGin(GetAllRedirectUris))
	router.POST("/redirecturis", ConverHttprouterToGin(AddRedirectUris))
	router.GET("/redirecturis/:argClientID/:argValue", ConverHttprouterToGin(GetRedirectUris))
	router.PUT("/redirecturis/:argClientID/:argValue", ConverHttprouterToGin(UpdateRedirectUris))
	router.DELETE("/redirecturis/:argClientID/:argValue", ConverHttprouterToGin(DeleteRedirectUris))
}

// GetAllRedirectUris is a function to get a slice of record(s) from redirect_uris table in the keycloak database
// @Summary Get list of RedirectUris
// @Tags RedirectUris
// @Description GetAllRedirectUris is a handler to get a slice of record(s) from redirect_uris table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.RedirectUris}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /redirecturis [get]
// http "http://localhost:8080/redirecturis?page=0&pagesize=20" X-Api-User:user123
func GetAllRedirectUris(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "redirect_uris", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllRedirectUris(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetRedirectUris is a function to get a single record from the redirect_uris table in the keycloak database
// @Summary Get record from table RedirectUris by  argClientID  argValue
// @Tags RedirectUris
// @ID argClientID
// @ID argValue
// @Description GetRedirectUris is a function to get a single record from the redirect_uris table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"
// @Param  argValue path string true "value"
// @Success 200 {object} model.RedirectUris
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /redirecturis/{argClientID}/{argValue} [get]
// http "http://localhost:8080/redirecturis/hello world/hello world" X-Api-User:user123
func GetRedirectUris(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argValue, err := parseString(ps, "argValue")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "redirect_uris", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetRedirectUris(ctx, argClientID, argValue)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddRedirectUris add to add a single record to redirect_uris table in the keycloak database
// @Summary Add an record to redirect_uris table
// @Description add to add a single record to redirect_uris table in the keycloak database
// @Tags RedirectUris
// @Accept  json
// @Produce  json
// @Param RedirectUris body model.RedirectUris true "Add RedirectUris"
// @Success 200 {object} model.RedirectUris
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /redirecturis [post]
// echo '{"client_id": "CZmKZvSCXoUCxaeaEEuxAiqgs","value": "lPCBkVNNGuseXtpYJtUXMJTkL"}' | http POST "http://localhost:8080/redirecturis" X-Api-User:user123
func AddRedirectUris(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	redirecturis := &model.RedirectUris{}

	if err := readJSON(r, redirecturis); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := redirecturis.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	redirecturis.Prepare()

	if err := redirecturis.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "redirect_uris", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	redirecturis, _, err = dao.AddRedirectUris(ctx, redirecturis)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, redirecturis)
}

// UpdateRedirectUris Update a single record from redirect_uris table in the keycloak database
// @Summary Update an record in table redirect_uris
// @Description Update a single record from redirect_uris table in the keycloak database
// @Tags RedirectUris
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"// @Param  argValue path string true "value"
// @Param  RedirectUris body model.RedirectUris true "Update RedirectUris record"
// @Success 200 {object} model.RedirectUris
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /redirecturis/{argClientID}/{argValue} [put]
// echo '{"client_id": "CZmKZvSCXoUCxaeaEEuxAiqgs","value": "lPCBkVNNGuseXtpYJtUXMJTkL"}' | http PUT "http://localhost:8080/redirecturis/hello world/hello world"  X-Api-User:user123
func UpdateRedirectUris(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argValue, err := parseString(ps, "argValue")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	redirecturis := &model.RedirectUris{}
	if err := readJSON(r, redirecturis); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := redirecturis.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	redirecturis.Prepare()

	if err := redirecturis.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "redirect_uris", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	redirecturis, _, err = dao.UpdateRedirectUris(ctx,
		argClientID, argValue,
		redirecturis)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, redirecturis)
}

// DeleteRedirectUris Delete a single record from redirect_uris table in the keycloak database
// @Summary Delete a record from redirect_uris
// @Description Delete a single record from redirect_uris table in the keycloak database
// @Tags RedirectUris
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"// @Param  argValue path string true "value"
// @Success 204 {object} model.RedirectUris
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /redirecturis/{argClientID}/{argValue} [delete]
// http DELETE "http://localhost:8080/redirecturis/hello world/hello world" X-Api-User:user123
func DeleteRedirectUris(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argValue, err := parseString(ps, "argValue")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "redirect_uris", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteRedirectUris(ctx, argClientID, argValue)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
