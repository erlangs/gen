package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configFedUserRequiredActionRouter(router *httprouter.Router) {
	router.GET("/feduserrequiredaction", GetAllFedUserRequiredAction)
	router.POST("/feduserrequiredaction", AddFedUserRequiredAction)
	router.GET("/feduserrequiredaction/:argRequiredAction/:argUserID", GetFedUserRequiredAction)
	router.PUT("/feduserrequiredaction/:argRequiredAction/:argUserID", UpdateFedUserRequiredAction)
	router.DELETE("/feduserrequiredaction/:argRequiredAction/:argUserID", DeleteFedUserRequiredAction)
}

func configGinFedUserRequiredActionRouter(router gin.IRoutes) {
	router.GET("/feduserrequiredaction", ConverHttprouterToGin(GetAllFedUserRequiredAction))
	router.POST("/feduserrequiredaction", ConverHttprouterToGin(AddFedUserRequiredAction))
	router.GET("/feduserrequiredaction/:argRequiredAction/:argUserID", ConverHttprouterToGin(GetFedUserRequiredAction))
	router.PUT("/feduserrequiredaction/:argRequiredAction/:argUserID", ConverHttprouterToGin(UpdateFedUserRequiredAction))
	router.DELETE("/feduserrequiredaction/:argRequiredAction/:argUserID", ConverHttprouterToGin(DeleteFedUserRequiredAction))
}

// GetAllFedUserRequiredAction is a function to get a slice of record(s) from fed_user_required_action table in the keycloak database
// @Summary Get list of FedUserRequiredAction
// @Tags FedUserRequiredAction
// @Description GetAllFedUserRequiredAction is a handler to get a slice of record(s) from fed_user_required_action table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.FedUserRequiredAction}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /feduserrequiredaction [get]
// http "http://localhost:8080/feduserrequiredaction?page=0&pagesize=20" X-Api-User:user123
func GetAllFedUserRequiredAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "fed_user_required_action", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllFedUserRequiredAction(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetFedUserRequiredAction is a function to get a single record from the fed_user_required_action table in the keycloak database
// @Summary Get record from table FedUserRequiredAction by  argRequiredAction  argUserID
// @Tags FedUserRequiredAction
// @ID argRequiredAction
// @ID argUserID
// @Description GetFedUserRequiredAction is a function to get a single record from the fed_user_required_action table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argRequiredAction path string true "required_action"
// @Param  argUserID path string true "user_id"
// @Success 200 {object} model.FedUserRequiredAction
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /feduserrequiredaction/{argRequiredAction}/{argUserID} [get]
// http "http://localhost:8080/feduserrequiredaction/hello world/hello world" X-Api-User:user123
func GetFedUserRequiredAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRequiredAction, err := parseString(ps, "argRequiredAction")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argUserID, err := parseString(ps, "argUserID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_required_action", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetFedUserRequiredAction(ctx, argRequiredAction, argUserID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddFedUserRequiredAction add to add a single record to fed_user_required_action table in the keycloak database
// @Summary Add an record to fed_user_required_action table
// @Description add to add a single record to fed_user_required_action table in the keycloak database
// @Tags FedUserRequiredAction
// @Accept  json
// @Produce  json
// @Param FedUserRequiredAction body model.FedUserRequiredAction true "Add FedUserRequiredAction"
// @Success 200 {object} model.FedUserRequiredAction
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /feduserrequiredaction [post]
// echo '{"required_action": "MdqgkNgeENGGdqFEIGUWpykiq","user_id": "cnyjqNAkGZGIKEvoOVGoZmYGp","realm_id": "RQeFPclUFxColZVyYpVSyGdNe","storage_provider_id": "CwEnPTCryVKrskooLbUyYPXos"}' | http POST "http://localhost:8080/feduserrequiredaction" X-Api-User:user123
func AddFedUserRequiredAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	feduserrequiredaction := &model.FedUserRequiredAction{}

	if err := readJSON(r, feduserrequiredaction); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := feduserrequiredaction.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	feduserrequiredaction.Prepare()

	if err := feduserrequiredaction.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_required_action", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	feduserrequiredaction, _, err = dao.AddFedUserRequiredAction(ctx, feduserrequiredaction)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, feduserrequiredaction)
}

// UpdateFedUserRequiredAction Update a single record from fed_user_required_action table in the keycloak database
// @Summary Update an record in table fed_user_required_action
// @Description Update a single record from fed_user_required_action table in the keycloak database
// @Tags FedUserRequiredAction
// @Accept  json
// @Produce  json
// @Param  argRequiredAction path string true "required_action"// @Param  argUserID path string true "user_id"
// @Param  FedUserRequiredAction body model.FedUserRequiredAction true "Update FedUserRequiredAction record"
// @Success 200 {object} model.FedUserRequiredAction
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /feduserrequiredaction/{argRequiredAction}/{argUserID} [put]
// echo '{"required_action": "MdqgkNgeENGGdqFEIGUWpykiq","user_id": "cnyjqNAkGZGIKEvoOVGoZmYGp","realm_id": "RQeFPclUFxColZVyYpVSyGdNe","storage_provider_id": "CwEnPTCryVKrskooLbUyYPXos"}' | http PUT "http://localhost:8080/feduserrequiredaction/hello world/hello world"  X-Api-User:user123
func UpdateFedUserRequiredAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRequiredAction, err := parseString(ps, "argRequiredAction")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argUserID, err := parseString(ps, "argUserID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	feduserrequiredaction := &model.FedUserRequiredAction{}
	if err := readJSON(r, feduserrequiredaction); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := feduserrequiredaction.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	feduserrequiredaction.Prepare()

	if err := feduserrequiredaction.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_required_action", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	feduserrequiredaction, _, err = dao.UpdateFedUserRequiredAction(ctx,
		argRequiredAction, argUserID,
		feduserrequiredaction)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, feduserrequiredaction)
}

// DeleteFedUserRequiredAction Delete a single record from fed_user_required_action table in the keycloak database
// @Summary Delete a record from fed_user_required_action
// @Description Delete a single record from fed_user_required_action table in the keycloak database
// @Tags FedUserRequiredAction
// @Accept  json
// @Produce  json
// @Param  argRequiredAction path string true "required_action"// @Param  argUserID path string true "user_id"
// @Success 204 {object} model.FedUserRequiredAction
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /feduserrequiredaction/{argRequiredAction}/{argUserID} [delete]
// http DELETE "http://localhost:8080/feduserrequiredaction/hello world/hello world" X-Api-User:user123
func DeleteFedUserRequiredAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRequiredAction, err := parseString(ps, "argRequiredAction")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argUserID, err := parseString(ps, "argUserID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_required_action", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteFedUserRequiredAction(ctx, argRequiredAction, argUserID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
