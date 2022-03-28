package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configFedUserAttributeRouter(router *httprouter.Router) {
	router.GET("/feduserattribute", GetAllFedUserAttribute)
	router.POST("/feduserattribute", AddFedUserAttribute)
	router.GET("/feduserattribute/:argID", GetFedUserAttribute)
	router.PUT("/feduserattribute/:argID", UpdateFedUserAttribute)
	router.DELETE("/feduserattribute/:argID", DeleteFedUserAttribute)
}

func configGinFedUserAttributeRouter(router gin.IRoutes) {
	router.GET("/feduserattribute", ConverHttprouterToGin(GetAllFedUserAttribute))
	router.POST("/feduserattribute", ConverHttprouterToGin(AddFedUserAttribute))
	router.GET("/feduserattribute/:argID", ConverHttprouterToGin(GetFedUserAttribute))
	router.PUT("/feduserattribute/:argID", ConverHttprouterToGin(UpdateFedUserAttribute))
	router.DELETE("/feduserattribute/:argID", ConverHttprouterToGin(DeleteFedUserAttribute))
}

// GetAllFedUserAttribute is a function to get a slice of record(s) from fed_user_attribute table in the keycloak database
// @Summary Get list of FedUserAttribute
// @Tags FedUserAttribute
// @Description GetAllFedUserAttribute is a handler to get a slice of record(s) from fed_user_attribute table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.FedUserAttribute}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /feduserattribute [get]
// http "http://localhost:8080/feduserattribute?page=0&pagesize=20" X-Api-User:user123
func GetAllFedUserAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "fed_user_attribute", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllFedUserAttribute(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetFedUserAttribute is a function to get a single record from the fed_user_attribute table in the keycloak database
// @Summary Get record from table FedUserAttribute by  argID
// @Tags FedUserAttribute
// @ID argID
// @Description GetFedUserAttribute is a function to get a single record from the fed_user_attribute table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.FedUserAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /feduserattribute/{argID} [get]
// http "http://localhost:8080/feduserattribute/hello world" X-Api-User:user123
func GetFedUserAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_attribute", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetFedUserAttribute(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddFedUserAttribute add to add a single record to fed_user_attribute table in the keycloak database
// @Summary Add an record to fed_user_attribute table
// @Description add to add a single record to fed_user_attribute table in the keycloak database
// @Tags FedUserAttribute
// @Accept  json
// @Produce  json
// @Param FedUserAttribute body model.FedUserAttribute true "Add FedUserAttribute"
// @Success 200 {object} model.FedUserAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /feduserattribute [post]
// echo '{"id": "eAxbJPAvZVtKbchviamjsbJYr","name": "cJWVZnfZtuCjWCvjZtDrdTymF","user_id": "deIeFqHCikEvdsHLpBmtMpESW","realm_id": "LyDpyjRahkjZBsquOojSoUlNb","storage_provider_id": "VxCJWkyLkpGRBVrFUCxJFqcaY","value": "jOwmBmkJjbTCIqwUiKELLMhDa"}' | http POST "http://localhost:8080/feduserattribute" X-Api-User:user123
func AddFedUserAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	feduserattribute := &model.FedUserAttribute{}

	if err := readJSON(r, feduserattribute); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := feduserattribute.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	feduserattribute.Prepare()

	if err := feduserattribute.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_attribute", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	feduserattribute, _, err = dao.AddFedUserAttribute(ctx, feduserattribute)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, feduserattribute)
}

// UpdateFedUserAttribute Update a single record from fed_user_attribute table in the keycloak database
// @Summary Update an record in table fed_user_attribute
// @Description Update a single record from fed_user_attribute table in the keycloak database
// @Tags FedUserAttribute
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  FedUserAttribute body model.FedUserAttribute true "Update FedUserAttribute record"
// @Success 200 {object} model.FedUserAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /feduserattribute/{argID} [put]
// echo '{"id": "eAxbJPAvZVtKbchviamjsbJYr","name": "cJWVZnfZtuCjWCvjZtDrdTymF","user_id": "deIeFqHCikEvdsHLpBmtMpESW","realm_id": "LyDpyjRahkjZBsquOojSoUlNb","storage_provider_id": "VxCJWkyLkpGRBVrFUCxJFqcaY","value": "jOwmBmkJjbTCIqwUiKELLMhDa"}' | http PUT "http://localhost:8080/feduserattribute/hello world"  X-Api-User:user123
func UpdateFedUserAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	feduserattribute := &model.FedUserAttribute{}
	if err := readJSON(r, feduserattribute); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := feduserattribute.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	feduserattribute.Prepare()

	if err := feduserattribute.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_attribute", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	feduserattribute, _, err = dao.UpdateFedUserAttribute(ctx,
		argID,
		feduserattribute)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, feduserattribute)
}

// DeleteFedUserAttribute Delete a single record from fed_user_attribute table in the keycloak database
// @Summary Delete a record from fed_user_attribute
// @Description Delete a single record from fed_user_attribute table in the keycloak database
// @Tags FedUserAttribute
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.FedUserAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /feduserattribute/{argID} [delete]
// http DELETE "http://localhost:8080/feduserattribute/hello world" X-Api-User:user123
func DeleteFedUserAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_attribute", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteFedUserAttribute(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
