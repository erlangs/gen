package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configRequiredActionProviderRouter(router *httprouter.Router) {
	router.GET("/requiredactionprovider", GetAllRequiredActionProvider)
	router.POST("/requiredactionprovider", AddRequiredActionProvider)
	router.GET("/requiredactionprovider/:argID", GetRequiredActionProvider)
	router.PUT("/requiredactionprovider/:argID", UpdateRequiredActionProvider)
	router.DELETE("/requiredactionprovider/:argID", DeleteRequiredActionProvider)
}

func configGinRequiredActionProviderRouter(router gin.IRoutes) {
	router.GET("/requiredactionprovider", ConverHttprouterToGin(GetAllRequiredActionProvider))
	router.POST("/requiredactionprovider", ConverHttprouterToGin(AddRequiredActionProvider))
	router.GET("/requiredactionprovider/:argID", ConverHttprouterToGin(GetRequiredActionProvider))
	router.PUT("/requiredactionprovider/:argID", ConverHttprouterToGin(UpdateRequiredActionProvider))
	router.DELETE("/requiredactionprovider/:argID", ConverHttprouterToGin(DeleteRequiredActionProvider))
}

// GetAllRequiredActionProvider is a function to get a slice of record(s) from required_action_provider table in the keycloak database
// @Summary Get list of RequiredActionProvider
// @Tags RequiredActionProvider
// @Description GetAllRequiredActionProvider is a handler to get a slice of record(s) from required_action_provider table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.RequiredActionProvider}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /requiredactionprovider [get]
// http "http://localhost:8080/requiredactionprovider?page=0&pagesize=20" X-Api-User:user123
func GetAllRequiredActionProvider(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "required_action_provider", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllRequiredActionProvider(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetRequiredActionProvider is a function to get a single record from the required_action_provider table in the keycloak database
// @Summary Get record from table RequiredActionProvider by  argID
// @Tags RequiredActionProvider
// @ID argID
// @Description GetRequiredActionProvider is a function to get a single record from the required_action_provider table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.RequiredActionProvider
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /requiredactionprovider/{argID} [get]
// http "http://localhost:8080/requiredactionprovider/hello world" X-Api-User:user123
func GetRequiredActionProvider(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "required_action_provider", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetRequiredActionProvider(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddRequiredActionProvider add to add a single record to required_action_provider table in the keycloak database
// @Summary Add an record to required_action_provider table
// @Description add to add a single record to required_action_provider table in the keycloak database
// @Tags RequiredActionProvider
// @Accept  json
// @Produce  json
// @Param RequiredActionProvider body model.RequiredActionProvider true "Add RequiredActionProvider"
// @Success 200 {object} model.RequiredActionProvider
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /requiredactionprovider [post]
// echo '{"id": "QPDVPUZySXcRRePKJNxySZggo","alias": "JBiUViIYurGhQpPxtYFaESwBS","name": "ADZFDDVZTcgxJycRQoxBOLkTc","realm_id": "UEDYlroONgqoKKbdjibyepggL","enabled": false,"default_action": true,"provider_id": "GBiKicikEyMhjWqQWvdVHxoJF","priority": 59}' | http POST "http://localhost:8080/requiredactionprovider" X-Api-User:user123
func AddRequiredActionProvider(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	requiredactionprovider := &model.RequiredActionProvider{}

	if err := readJSON(r, requiredactionprovider); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := requiredactionprovider.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	requiredactionprovider.Prepare()

	if err := requiredactionprovider.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "required_action_provider", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	requiredactionprovider, _, err = dao.AddRequiredActionProvider(ctx, requiredactionprovider)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, requiredactionprovider)
}

// UpdateRequiredActionProvider Update a single record from required_action_provider table in the keycloak database
// @Summary Update an record in table required_action_provider
// @Description Update a single record from required_action_provider table in the keycloak database
// @Tags RequiredActionProvider
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  RequiredActionProvider body model.RequiredActionProvider true "Update RequiredActionProvider record"
// @Success 200 {object} model.RequiredActionProvider
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /requiredactionprovider/{argID} [put]
// echo '{"id": "QPDVPUZySXcRRePKJNxySZggo","alias": "JBiUViIYurGhQpPxtYFaESwBS","name": "ADZFDDVZTcgxJycRQoxBOLkTc","realm_id": "UEDYlroONgqoKKbdjibyepggL","enabled": false,"default_action": true,"provider_id": "GBiKicikEyMhjWqQWvdVHxoJF","priority": 59}' | http PUT "http://localhost:8080/requiredactionprovider/hello world"  X-Api-User:user123
func UpdateRequiredActionProvider(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	requiredactionprovider := &model.RequiredActionProvider{}
	if err := readJSON(r, requiredactionprovider); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := requiredactionprovider.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	requiredactionprovider.Prepare()

	if err := requiredactionprovider.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "required_action_provider", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	requiredactionprovider, _, err = dao.UpdateRequiredActionProvider(ctx,
		argID,
		requiredactionprovider)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, requiredactionprovider)
}

// DeleteRequiredActionProvider Delete a single record from required_action_provider table in the keycloak database
// @Summary Delete a record from required_action_provider
// @Description Delete a single record from required_action_provider table in the keycloak database
// @Tags RequiredActionProvider
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.RequiredActionProvider
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /requiredactionprovider/{argID} [delete]
// http DELETE "http://localhost:8080/requiredactionprovider/hello world" X-Api-User:user123
func DeleteRequiredActionProvider(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "required_action_provider", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteRequiredActionProvider(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
