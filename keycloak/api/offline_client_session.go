package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configOfflineClientSessionRouter(router *httprouter.Router) {
	router.GET("/offlineclientsession", GetAllOfflineClientSession)
	router.POST("/offlineclientsession", AddOfflineClientSession)
	router.GET("/offlineclientsession/:argUserSessionID/:argClientID/:argOfflineFlag/:argClientStorageProvider/:argExternalClientID", GetOfflineClientSession)
	router.PUT("/offlineclientsession/:argUserSessionID/:argClientID/:argOfflineFlag/:argClientStorageProvider/:argExternalClientID", UpdateOfflineClientSession)
	router.DELETE("/offlineclientsession/:argUserSessionID/:argClientID/:argOfflineFlag/:argClientStorageProvider/:argExternalClientID", DeleteOfflineClientSession)
}

func configGinOfflineClientSessionRouter(router gin.IRoutes) {
	router.GET("/offlineclientsession", ConverHttprouterToGin(GetAllOfflineClientSession))
	router.POST("/offlineclientsession", ConverHttprouterToGin(AddOfflineClientSession))
	router.GET("/offlineclientsession/:argUserSessionID/:argClientID/:argOfflineFlag/:argClientStorageProvider/:argExternalClientID", ConverHttprouterToGin(GetOfflineClientSession))
	router.PUT("/offlineclientsession/:argUserSessionID/:argClientID/:argOfflineFlag/:argClientStorageProvider/:argExternalClientID", ConverHttprouterToGin(UpdateOfflineClientSession))
	router.DELETE("/offlineclientsession/:argUserSessionID/:argClientID/:argOfflineFlag/:argClientStorageProvider/:argExternalClientID", ConverHttprouterToGin(DeleteOfflineClientSession))
}

// GetAllOfflineClientSession is a function to get a slice of record(s) from offline_client_session table in the keycloak database
// @Summary Get list of OfflineClientSession
// @Tags OfflineClientSession
// @Description GetAllOfflineClientSession is a handler to get a slice of record(s) from offline_client_session table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.OfflineClientSession}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /offlineclientsession [get]
// http "http://localhost:8080/offlineclientsession?page=0&pagesize=20" X-Api-User:user123
func GetAllOfflineClientSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "offline_client_session", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllOfflineClientSession(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetOfflineClientSession is a function to get a single record from the offline_client_session table in the keycloak database
// @Summary Get record from table OfflineClientSession by  argUserSessionID  argClientID  argOfflineFlag  argClientStorageProvider  argExternalClientID
// @Tags OfflineClientSession
// @ID argUserSessionID
// @ID argClientID
// @ID argOfflineFlag
// @ID argClientStorageProvider
// @ID argExternalClientID
// @Description GetOfflineClientSession is a function to get a single record from the offline_client_session table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argUserSessionID path string true "user_session_id"
// @Param  argClientID path string true "client_id"
// @Param  argOfflineFlag path string true "offline_flag"
// @Param  argClientStorageProvider path string true "client_storage_provider"
// @Param  argExternalClientID path string true "external_client_id"
// @Success 200 {object} model.OfflineClientSession
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /offlineclientsession/{argUserSessionID}/{argClientID}/{argOfflineFlag}/{argClientStorageProvider}/{argExternalClientID} [get]
// http "http://localhost:8080/offlineclientsession/hello world/hello world/hello world/hello world/hello world" X-Api-User:user123
func GetOfflineClientSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserSessionID, err := parseString(ps, "argUserSessionID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argOfflineFlag, err := parseString(ps, "argOfflineFlag")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argClientStorageProvider, err := parseString(ps, "argClientStorageProvider")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argExternalClientID, err := parseString(ps, "argExternalClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "offline_client_session", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetOfflineClientSession(ctx, argUserSessionID, argClientID, argOfflineFlag, argClientStorageProvider, argExternalClientID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddOfflineClientSession add to add a single record to offline_client_session table in the keycloak database
// @Summary Add an record to offline_client_session table
// @Description add to add a single record to offline_client_session table in the keycloak database
// @Tags OfflineClientSession
// @Accept  json
// @Produce  json
// @Param OfflineClientSession body model.OfflineClientSession true "Add OfflineClientSession"
// @Success 200 {object} model.OfflineClientSession
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /offlineclientsession [post]
// echo '{"user_session_id": "dGLdvddddUCWDrcjAxAwQWNFv","client_id": "vynOselRaFpFOuNadpdqGIFgQ","offline_flag": "hJJoRvGZLAOOlMdhOGliVMNOu","timestamp": 22,"data": "kXAZcBFqOcGtDcDMCLRCqocdO","client_storage_provider": "bPKkktGhvVdULDGjWyXfhwLQs","external_client_id": "VfMSHGoOfKKYkNmFMAVXbWSMV"}' | http POST "http://localhost:8080/offlineclientsession" X-Api-User:user123
func AddOfflineClientSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	offlineclientsession := &model.OfflineClientSession{}

	if err := readJSON(r, offlineclientsession); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := offlineclientsession.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	offlineclientsession.Prepare()

	if err := offlineclientsession.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "offline_client_session", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	offlineclientsession, _, err = dao.AddOfflineClientSession(ctx, offlineclientsession)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, offlineclientsession)
}

// UpdateOfflineClientSession Update a single record from offline_client_session table in the keycloak database
// @Summary Update an record in table offline_client_session
// @Description Update a single record from offline_client_session table in the keycloak database
// @Tags OfflineClientSession
// @Accept  json
// @Produce  json
// @Param  argUserSessionID path string true "user_session_id"// @Param  argClientID path string true "client_id"// @Param  argOfflineFlag path string true "offline_flag"// @Param  argClientStorageProvider path string true "client_storage_provider"// @Param  argExternalClientID path string true "external_client_id"
// @Param  OfflineClientSession body model.OfflineClientSession true "Update OfflineClientSession record"
// @Success 200 {object} model.OfflineClientSession
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /offlineclientsession/{argUserSessionID}/{argClientID}/{argOfflineFlag}/{argClientStorageProvider}/{argExternalClientID} [put]
// echo '{"user_session_id": "dGLdvddddUCWDrcjAxAwQWNFv","client_id": "vynOselRaFpFOuNadpdqGIFgQ","offline_flag": "hJJoRvGZLAOOlMdhOGliVMNOu","timestamp": 22,"data": "kXAZcBFqOcGtDcDMCLRCqocdO","client_storage_provider": "bPKkktGhvVdULDGjWyXfhwLQs","external_client_id": "VfMSHGoOfKKYkNmFMAVXbWSMV"}' | http PUT "http://localhost:8080/offlineclientsession/hello world/hello world/hello world/hello world/hello world"  X-Api-User:user123
func UpdateOfflineClientSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserSessionID, err := parseString(ps, "argUserSessionID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argOfflineFlag, err := parseString(ps, "argOfflineFlag")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argClientStorageProvider, err := parseString(ps, "argClientStorageProvider")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argExternalClientID, err := parseString(ps, "argExternalClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	offlineclientsession := &model.OfflineClientSession{}
	if err := readJSON(r, offlineclientsession); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := offlineclientsession.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	offlineclientsession.Prepare()

	if err := offlineclientsession.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "offline_client_session", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	offlineclientsession, _, err = dao.UpdateOfflineClientSession(ctx,
		argUserSessionID, argClientID, argOfflineFlag, argClientStorageProvider, argExternalClientID,
		offlineclientsession)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, offlineclientsession)
}

// DeleteOfflineClientSession Delete a single record from offline_client_session table in the keycloak database
// @Summary Delete a record from offline_client_session
// @Description Delete a single record from offline_client_session table in the keycloak database
// @Tags OfflineClientSession
// @Accept  json
// @Produce  json
// @Param  argUserSessionID path string true "user_session_id"// @Param  argClientID path string true "client_id"// @Param  argOfflineFlag path string true "offline_flag"// @Param  argClientStorageProvider path string true "client_storage_provider"// @Param  argExternalClientID path string true "external_client_id"
// @Success 204 {object} model.OfflineClientSession
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /offlineclientsession/{argUserSessionID}/{argClientID}/{argOfflineFlag}/{argClientStorageProvider}/{argExternalClientID} [delete]
// http DELETE "http://localhost:8080/offlineclientsession/hello world/hello world/hello world/hello world/hello world" X-Api-User:user123
func DeleteOfflineClientSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argUserSessionID, err := parseString(ps, "argUserSessionID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argOfflineFlag, err := parseString(ps, "argOfflineFlag")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argClientStorageProvider, err := parseString(ps, "argClientStorageProvider")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argExternalClientID, err := parseString(ps, "argExternalClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "offline_client_session", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteOfflineClientSession(ctx, argUserSessionID, argClientID, argOfflineFlag, argClientStorageProvider, argExternalClientID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
