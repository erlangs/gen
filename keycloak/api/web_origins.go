package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configWebOriginsRouter(router *httprouter.Router) {
	router.GET("/weborigins", GetAllWebOrigins)
	router.POST("/weborigins", AddWebOrigins)
	router.GET("/weborigins/:argClientID/:argValue", GetWebOrigins)
	router.PUT("/weborigins/:argClientID/:argValue", UpdateWebOrigins)
	router.DELETE("/weborigins/:argClientID/:argValue", DeleteWebOrigins)
}

func configGinWebOriginsRouter(router gin.IRoutes) {
	router.GET("/weborigins", ConverHttprouterToGin(GetAllWebOrigins))
	router.POST("/weborigins", ConverHttprouterToGin(AddWebOrigins))
	router.GET("/weborigins/:argClientID/:argValue", ConverHttprouterToGin(GetWebOrigins))
	router.PUT("/weborigins/:argClientID/:argValue", ConverHttprouterToGin(UpdateWebOrigins))
	router.DELETE("/weborigins/:argClientID/:argValue", ConverHttprouterToGin(DeleteWebOrigins))
}

// GetAllWebOrigins is a function to get a slice of record(s) from web_origins table in the keycloak database
// @Summary Get list of WebOrigins
// @Tags WebOrigins
// @Description GetAllWebOrigins is a handler to get a slice of record(s) from web_origins table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.WebOrigins}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /weborigins [get]
// http "http://localhost:8080/weborigins?page=0&pagesize=20" X-Api-User:user123
func GetAllWebOrigins(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "web_origins", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllWebOrigins(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetWebOrigins is a function to get a single record from the web_origins table in the keycloak database
// @Summary Get record from table WebOrigins by  argClientID  argValue
// @Tags WebOrigins
// @ID argClientID
// @ID argValue
// @Description GetWebOrigins is a function to get a single record from the web_origins table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"
// @Param  argValue path string true "value"
// @Success 200 {object} model.WebOrigins
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /weborigins/{argClientID}/{argValue} [get]
// http "http://localhost:8080/weborigins/hello world/hello world" X-Api-User:user123
func GetWebOrigins(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "web_origins", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetWebOrigins(ctx, argClientID, argValue)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddWebOrigins add to add a single record to web_origins table in the keycloak database
// @Summary Add an record to web_origins table
// @Description add to add a single record to web_origins table in the keycloak database
// @Tags WebOrigins
// @Accept  json
// @Produce  json
// @Param WebOrigins body model.WebOrigins true "Add WebOrigins"
// @Success 200 {object} model.WebOrigins
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /weborigins [post]
// echo '{"client_id": "VicweBNLWfQYQNEvtdhTvidki","value": "EKidYXqcXWDWYmmtQqAMkdPth"}' | http POST "http://localhost:8080/weborigins" X-Api-User:user123
func AddWebOrigins(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	weborigins := &model.WebOrigins{}

	if err := readJSON(r, weborigins); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := weborigins.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	weborigins.Prepare()

	if err := weborigins.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "web_origins", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	weborigins, _, err = dao.AddWebOrigins(ctx, weborigins)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, weborigins)
}

// UpdateWebOrigins Update a single record from web_origins table in the keycloak database
// @Summary Update an record in table web_origins
// @Description Update a single record from web_origins table in the keycloak database
// @Tags WebOrigins
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"// @Param  argValue path string true "value"
// @Param  WebOrigins body model.WebOrigins true "Update WebOrigins record"
// @Success 200 {object} model.WebOrigins
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /weborigins/{argClientID}/{argValue} [put]
// echo '{"client_id": "VicweBNLWfQYQNEvtdhTvidki","value": "EKidYXqcXWDWYmmtQqAMkdPth"}' | http PUT "http://localhost:8080/weborigins/hello world/hello world"  X-Api-User:user123
func UpdateWebOrigins(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	weborigins := &model.WebOrigins{}
	if err := readJSON(r, weborigins); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := weborigins.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	weborigins.Prepare()

	if err := weborigins.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "web_origins", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	weborigins, _, err = dao.UpdateWebOrigins(ctx,
		argClientID, argValue,
		weborigins)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, weborigins)
}

// DeleteWebOrigins Delete a single record from web_origins table in the keycloak database
// @Summary Delete a record from web_origins
// @Description Delete a single record from web_origins table in the keycloak database
// @Tags WebOrigins
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"// @Param  argValue path string true "value"
// @Success 204 {object} model.WebOrigins
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /weborigins/{argClientID}/{argValue} [delete]
// http DELETE "http://localhost:8080/weborigins/hello world/hello world" X-Api-User:user123
func DeleteWebOrigins(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "web_origins", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteWebOrigins(ctx, argClientID, argValue)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
