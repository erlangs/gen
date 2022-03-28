package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configOfflineUserSessionRouter(router *httprouter.Router) {
	router.GET("/offlineusersession", GetAllOfflineUserSession)
	router.POST("/offlineusersession", AddOfflineUserSession)
	router.GET("/offlineusersession/:argUserSessionID/:argOfflineFlag", GetOfflineUserSession)
	router.PUT("/offlineusersession/:argUserSessionID/:argOfflineFlag", UpdateOfflineUserSession)
	router.DELETE("/offlineusersession/:argUserSessionID/:argOfflineFlag", DeleteOfflineUserSession)
}

func configGinOfflineUserSessionRouter(router gin.IRoutes) {
	router.GET("/offlineusersession", ConverHttprouterToGin(GetAllOfflineUserSession))
	router.POST("/offlineusersession", ConverHttprouterToGin(AddOfflineUserSession))
	router.GET("/offlineusersession/:argUserSessionID/:argOfflineFlag", ConverHttprouterToGin(GetOfflineUserSession))
	router.PUT("/offlineusersession/:argUserSessionID/:argOfflineFlag", ConverHttprouterToGin(UpdateOfflineUserSession))
	router.DELETE("/offlineusersession/:argUserSessionID/:argOfflineFlag", ConverHttprouterToGin(DeleteOfflineUserSession))
}

// GetAllOfflineUserSession is a function to get a slice of record(s) from offline_user_session table in the keycloak database
// @Summary Get list of OfflineUserSession
// @Tags OfflineUserSession
// @Description GetAllOfflineUserSession is a handler to get a slice of record(s) from offline_user_session table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.OfflineUserSession}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /offlineusersession [get]
// http "http://localhost:8080/offlineusersession?page=0&pagesize=20" X-Api-User:user123
func GetAllOfflineUserSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "offline_user_session", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllOfflineUserSession(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetOfflineUserSession is a function to get a single record from the offline_user_session table in the keycloak database
// @Summary Get record from table OfflineUserSession by  argUserSessionID  argOfflineFlag
// @Tags OfflineUserSession
// @ID argUserSessionID
// @ID argOfflineFlag
// @Description GetOfflineUserSession is a function to get a single record from the offline_user_session table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argUserSessionID path string true "user_session_id"
// @Param  argOfflineFlag path string true "offline_flag"
// @Success 200 {object} model.OfflineUserSession
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /offlineusersession/{argUserSessionID}/{argOfflineFlag} [get]
// http "http://localhost:8080/offlineusersession/hello world/hello world" X-Api-User:user123
func GetOfflineUserSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserSessionID, err := parseString(ps, "argUserSessionID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argOfflineFlag, err := parseString(ps, "argOfflineFlag")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "offline_user_session", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetOfflineUserSession(ctx, argUserSessionID, argOfflineFlag)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddOfflineUserSession add to add a single record to offline_user_session table in the keycloak database
// @Summary Add an record to offline_user_session table
// @Description add to add a single record to offline_user_session table in the keycloak database
// @Tags OfflineUserSession
// @Accept  json
// @Produce  json
// @Param OfflineUserSession body model.OfflineUserSession true "Add OfflineUserSession"
// @Success 200 {object} model.OfflineUserSession
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /offlineusersession [post]
// echo '{"user_session_id": "nRvsZPuGfgMFYHgUETstMfwoM","user_id": "wFMVOGoJPcwJZcHiNeuhlolFn","realm_id": "ogsNNBTjwTMjCAHBgGrMVLaBZ","created_on": 32,"offline_flag": "MmoitLDISGFoaCgsJUPkejDdk","data": "asfBRfHyxPkrKxgbYcCZWuQPG","last_session_refresh": 9}' | http POST "http://localhost:8080/offlineusersession" X-Api-User:user123
func AddOfflineUserSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	offlineusersession := &model.OfflineUserSession{}

	if err := readJSON(r, offlineusersession); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := offlineusersession.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	offlineusersession.Prepare()

	if err := offlineusersession.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "offline_user_session", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	offlineusersession, _, err = dao.AddOfflineUserSession(ctx, offlineusersession)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, offlineusersession)
}

// UpdateOfflineUserSession Update a single record from offline_user_session table in the keycloak database
// @Summary Update an record in table offline_user_session
// @Description Update a single record from offline_user_session table in the keycloak database
// @Tags OfflineUserSession
// @Accept  json
// @Produce  json
// @Param  argUserSessionID path string true "user_session_id"// @Param  argOfflineFlag path string true "offline_flag"
// @Param  OfflineUserSession body model.OfflineUserSession true "Update OfflineUserSession record"
// @Success 200 {object} model.OfflineUserSession
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /offlineusersession/{argUserSessionID}/{argOfflineFlag} [put]
// echo '{"user_session_id": "nRvsZPuGfgMFYHgUETstMfwoM","user_id": "wFMVOGoJPcwJZcHiNeuhlolFn","realm_id": "ogsNNBTjwTMjCAHBgGrMVLaBZ","created_on": 32,"offline_flag": "MmoitLDISGFoaCgsJUPkejDdk","data": "asfBRfHyxPkrKxgbYcCZWuQPG","last_session_refresh": 9}' | http PUT "http://localhost:8080/offlineusersession/hello world/hello world"  X-Api-User:user123
func UpdateOfflineUserSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserSessionID, err := parseString(ps, "argUserSessionID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argOfflineFlag, err := parseString(ps, "argOfflineFlag")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	offlineusersession := &model.OfflineUserSession{}
	if err := readJSON(r, offlineusersession); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := offlineusersession.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	offlineusersession.Prepare()

	if err := offlineusersession.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "offline_user_session", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	offlineusersession, _, err = dao.UpdateOfflineUserSession(ctx,
		argUserSessionID, argOfflineFlag,
		offlineusersession)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, offlineusersession)
}

// DeleteOfflineUserSession Delete a single record from offline_user_session table in the keycloak database
// @Summary Delete a record from offline_user_session
// @Description Delete a single record from offline_user_session table in the keycloak database
// @Tags OfflineUserSession
// @Accept  json
// @Produce  json
// @Param  argUserSessionID path string true "user_session_id"// @Param  argOfflineFlag path string true "offline_flag"
// @Success 204 {object} model.OfflineUserSession
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /offlineusersession/{argUserSessionID}/{argOfflineFlag} [delete]
// http DELETE "http://localhost:8080/offlineusersession/hello world/hello world" X-Api-User:user123
func DeleteOfflineUserSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserSessionID, err := parseString(ps, "argUserSessionID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argOfflineFlag, err := parseString(ps, "argOfflineFlag")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "offline_user_session", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteOfflineUserSession(ctx, argUserSessionID, argOfflineFlag)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
