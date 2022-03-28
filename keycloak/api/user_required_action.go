package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configUserRequiredActionRouter(router *httprouter.Router) {
	router.GET("/userrequiredaction", GetAllUserRequiredAction)
	router.POST("/userrequiredaction", AddUserRequiredAction)
	router.GET("/userrequiredaction/:argUserID/:argRequiredAction", GetUserRequiredAction)
	router.PUT("/userrequiredaction/:argUserID/:argRequiredAction", UpdateUserRequiredAction)
	router.DELETE("/userrequiredaction/:argUserID/:argRequiredAction", DeleteUserRequiredAction)
}

func configGinUserRequiredActionRouter(router gin.IRoutes) {
	router.GET("/userrequiredaction", ConverHttprouterToGin(GetAllUserRequiredAction))
	router.POST("/userrequiredaction", ConverHttprouterToGin(AddUserRequiredAction))
	router.GET("/userrequiredaction/:argUserID/:argRequiredAction", ConverHttprouterToGin(GetUserRequiredAction))
	router.PUT("/userrequiredaction/:argUserID/:argRequiredAction", ConverHttprouterToGin(UpdateUserRequiredAction))
	router.DELETE("/userrequiredaction/:argUserID/:argRequiredAction", ConverHttprouterToGin(DeleteUserRequiredAction))
}

// GetAllUserRequiredAction is a function to get a slice of record(s) from user_required_action table in the keycloak database
// @Summary Get list of UserRequiredAction
// @Tags UserRequiredAction
// @Description GetAllUserRequiredAction is a handler to get a slice of record(s) from user_required_action table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.UserRequiredAction}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userrequiredaction [get]
// http "http://localhost:8080/userrequiredaction?page=0&pagesize=20" X-Api-User:user123
func GetAllUserRequiredAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "user_required_action", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllUserRequiredAction(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetUserRequiredAction is a function to get a single record from the user_required_action table in the keycloak database
// @Summary Get record from table UserRequiredAction by  argUserID  argRequiredAction
// @Tags UserRequiredAction
// @ID argUserID
// @ID argRequiredAction
// @Description GetUserRequiredAction is a function to get a single record from the user_required_action table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argUserID path string true "user_id"
// @Param  argRequiredAction path string true "required_action"
// @Success 200 {object} model.UserRequiredAction
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /userrequiredaction/{argUserID}/{argRequiredAction} [get]
// http "http://localhost:8080/userrequiredaction/hello world/hello world" X-Api-User:user123
func GetUserRequiredAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserID, err := parseString(ps, "argUserID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argRequiredAction, err := parseString(ps, "argRequiredAction")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_required_action", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetUserRequiredAction(ctx, argUserID, argRequiredAction)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddUserRequiredAction add to add a single record to user_required_action table in the keycloak database
// @Summary Add an record to user_required_action table
// @Description add to add a single record to user_required_action table in the keycloak database
// @Tags UserRequiredAction
// @Accept  json
// @Produce  json
// @Param UserRequiredAction body model.UserRequiredAction true "Add UserRequiredAction"
// @Success 200 {object} model.UserRequiredAction
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userrequiredaction [post]
// echo '{"user_id": "JRZvhHGkBxibPLTPQAXBtogau","required_action": "RPSoynwJRFXZqDBLeaNXTuGvn"}' | http POST "http://localhost:8080/userrequiredaction" X-Api-User:user123
func AddUserRequiredAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	userrequiredaction := &model.UserRequiredAction{}

	if err := readJSON(r, userrequiredaction); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userrequiredaction.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userrequiredaction.Prepare()

	if err := userrequiredaction.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_required_action", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	userrequiredaction, _, err = dao.AddUserRequiredAction(ctx, userrequiredaction)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userrequiredaction)
}

// UpdateUserRequiredAction Update a single record from user_required_action table in the keycloak database
// @Summary Update an record in table user_required_action
// @Description Update a single record from user_required_action table in the keycloak database
// @Tags UserRequiredAction
// @Accept  json
// @Produce  json
// @Param  argUserID path string true "user_id"// @Param  argRequiredAction path string true "required_action"
// @Param  UserRequiredAction body model.UserRequiredAction true "Update UserRequiredAction record"
// @Success 200 {object} model.UserRequiredAction
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userrequiredaction/{argUserID}/{argRequiredAction} [put]
// echo '{"user_id": "JRZvhHGkBxibPLTPQAXBtogau","required_action": "RPSoynwJRFXZqDBLeaNXTuGvn"}' | http PUT "http://localhost:8080/userrequiredaction/hello world/hello world"  X-Api-User:user123
func UpdateUserRequiredAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserID, err := parseString(ps, "argUserID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argRequiredAction, err := parseString(ps, "argRequiredAction")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	userrequiredaction := &model.UserRequiredAction{}
	if err := readJSON(r, userrequiredaction); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userrequiredaction.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userrequiredaction.Prepare()

	if err := userrequiredaction.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_required_action", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	userrequiredaction, _, err = dao.UpdateUserRequiredAction(ctx,
		argUserID, argRequiredAction,
		userrequiredaction)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userrequiredaction)
}

// DeleteUserRequiredAction Delete a single record from user_required_action table in the keycloak database
// @Summary Delete a record from user_required_action
// @Description Delete a single record from user_required_action table in the keycloak database
// @Tags UserRequiredAction
// @Accept  json
// @Produce  json
// @Param  argUserID path string true "user_id"// @Param  argRequiredAction path string true "required_action"
// @Success 204 {object} model.UserRequiredAction
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /userrequiredaction/{argUserID}/{argRequiredAction} [delete]
// http DELETE "http://localhost:8080/userrequiredaction/hello world/hello world" X-Api-User:user123
func DeleteUserRequiredAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserID, err := parseString(ps, "argUserID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argRequiredAction, err := parseString(ps, "argRequiredAction")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_required_action", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteUserRequiredAction(ctx, argUserID, argRequiredAction)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
