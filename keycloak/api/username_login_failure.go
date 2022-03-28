package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configUsernameLoginFailureRouter(router *httprouter.Router) {
	router.GET("/usernameloginfailure", GetAllUsernameLoginFailure)
	router.POST("/usernameloginfailure", AddUsernameLoginFailure)
	router.GET("/usernameloginfailure/:argRealmID/:argUsername", GetUsernameLoginFailure)
	router.PUT("/usernameloginfailure/:argRealmID/:argUsername", UpdateUsernameLoginFailure)
	router.DELETE("/usernameloginfailure/:argRealmID/:argUsername", DeleteUsernameLoginFailure)
}

func configGinUsernameLoginFailureRouter(router gin.IRoutes) {
	router.GET("/usernameloginfailure", ConverHttprouterToGin(GetAllUsernameLoginFailure))
	router.POST("/usernameloginfailure", ConverHttprouterToGin(AddUsernameLoginFailure))
	router.GET("/usernameloginfailure/:argRealmID/:argUsername", ConverHttprouterToGin(GetUsernameLoginFailure))
	router.PUT("/usernameloginfailure/:argRealmID/:argUsername", ConverHttprouterToGin(UpdateUsernameLoginFailure))
	router.DELETE("/usernameloginfailure/:argRealmID/:argUsername", ConverHttprouterToGin(DeleteUsernameLoginFailure))
}

// GetAllUsernameLoginFailure is a function to get a slice of record(s) from username_login_failure table in the keycloak database
// @Summary Get list of UsernameLoginFailure
// @Tags UsernameLoginFailure
// @Description GetAllUsernameLoginFailure is a handler to get a slice of record(s) from username_login_failure table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.UsernameLoginFailure}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /usernameloginfailure [get]
// http "http://localhost:8080/usernameloginfailure?page=0&pagesize=20" X-Api-User:user123
func GetAllUsernameLoginFailure(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "username_login_failure", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllUsernameLoginFailure(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetUsernameLoginFailure is a function to get a single record from the username_login_failure table in the keycloak database
// @Summary Get record from table UsernameLoginFailure by  argRealmID  argUsername
// @Tags UsernameLoginFailure
// @ID argRealmID
// @ID argUsername
// @Description GetUsernameLoginFailure is a function to get a single record from the username_login_failure table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"
// @Param  argUsername path string true "username"
// @Success 200 {object} model.UsernameLoginFailure
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /usernameloginfailure/{argRealmID}/{argUsername} [get]
// http "http://localhost:8080/usernameloginfailure/hello world/hello world" X-Api-User:user123
func GetUsernameLoginFailure(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argUsername, err := parseString(ps, "argUsername")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "username_login_failure", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetUsernameLoginFailure(ctx, argRealmID, argUsername)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddUsernameLoginFailure add to add a single record to username_login_failure table in the keycloak database
// @Summary Add an record to username_login_failure table
// @Description add to add a single record to username_login_failure table in the keycloak database
// @Tags UsernameLoginFailure
// @Accept  json
// @Produce  json
// @Param UsernameLoginFailure body model.UsernameLoginFailure true "Add UsernameLoginFailure"
// @Success 200 {object} model.UsernameLoginFailure
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /usernameloginfailure [post]
// echo '{"realm_id": "ppyakgjtMRLTcTYSSgIOFlifd","username": "CYmmLgNffNdMTfImAJvbrfrsL","failed_login_not_before": 76,"last_failure": 1,"last_ip_failure": "LwoWfRACcOdQLZpkYGMQAVCoE","num_failures": 40}' | http POST "http://localhost:8080/usernameloginfailure" X-Api-User:user123
func AddUsernameLoginFailure(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	usernameloginfailure := &model.UsernameLoginFailure{}

	if err := readJSON(r, usernameloginfailure); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := usernameloginfailure.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	usernameloginfailure.Prepare()

	if err := usernameloginfailure.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "username_login_failure", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	usernameloginfailure, _, err = dao.AddUsernameLoginFailure(ctx, usernameloginfailure)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, usernameloginfailure)
}

// UpdateUsernameLoginFailure Update a single record from username_login_failure table in the keycloak database
// @Summary Update an record in table username_login_failure
// @Description Update a single record from username_login_failure table in the keycloak database
// @Tags UsernameLoginFailure
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"// @Param  argUsername path string true "username"
// @Param  UsernameLoginFailure body model.UsernameLoginFailure true "Update UsernameLoginFailure record"
// @Success 200 {object} model.UsernameLoginFailure
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /usernameloginfailure/{argRealmID}/{argUsername} [put]
// echo '{"realm_id": "ppyakgjtMRLTcTYSSgIOFlifd","username": "CYmmLgNffNdMTfImAJvbrfrsL","failed_login_not_before": 76,"last_failure": 1,"last_ip_failure": "LwoWfRACcOdQLZpkYGMQAVCoE","num_failures": 40}' | http PUT "http://localhost:8080/usernameloginfailure/hello world/hello world"  X-Api-User:user123
func UpdateUsernameLoginFailure(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argUsername, err := parseString(ps, "argUsername")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	usernameloginfailure := &model.UsernameLoginFailure{}
	if err := readJSON(r, usernameloginfailure); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := usernameloginfailure.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	usernameloginfailure.Prepare()

	if err := usernameloginfailure.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "username_login_failure", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	usernameloginfailure, _, err = dao.UpdateUsernameLoginFailure(ctx,
		argRealmID, argUsername,
		usernameloginfailure)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, usernameloginfailure)
}

// DeleteUsernameLoginFailure Delete a single record from username_login_failure table in the keycloak database
// @Summary Delete a record from username_login_failure
// @Description Delete a single record from username_login_failure table in the keycloak database
// @Tags UsernameLoginFailure
// @Accept  json
// @Produce  json
// @Param  argRealmID path string true "realm_id"// @Param  argUsername path string true "username"
// @Success 204 {object} model.UsernameLoginFailure
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /usernameloginfailure/{argRealmID}/{argUsername} [delete]
// http DELETE "http://localhost:8080/usernameloginfailure/hello world/hello world" X-Api-User:user123
func DeleteUsernameLoginFailure(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRealmID, err := parseString(ps, "argRealmID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argUsername, err := parseString(ps, "argUsername")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "username_login_failure", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteUsernameLoginFailure(ctx, argRealmID, argUsername)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
