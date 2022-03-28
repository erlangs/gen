package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configUserConsentRouter(router *httprouter.Router) {
	router.GET("/userconsent", GetAllUserConsent)
	router.POST("/userconsent", AddUserConsent)
	router.GET("/userconsent/:argID", GetUserConsent)
	router.PUT("/userconsent/:argID", UpdateUserConsent)
	router.DELETE("/userconsent/:argID", DeleteUserConsent)
}

func configGinUserConsentRouter(router gin.IRoutes) {
	router.GET("/userconsent", ConverHttprouterToGin(GetAllUserConsent))
	router.POST("/userconsent", ConverHttprouterToGin(AddUserConsent))
	router.GET("/userconsent/:argID", ConverHttprouterToGin(GetUserConsent))
	router.PUT("/userconsent/:argID", ConverHttprouterToGin(UpdateUserConsent))
	router.DELETE("/userconsent/:argID", ConverHttprouterToGin(DeleteUserConsent))
}

// GetAllUserConsent is a function to get a slice of record(s) from user_consent table in the keycloak database
// @Summary Get list of UserConsent
// @Tags UserConsent
// @Description GetAllUserConsent is a handler to get a slice of record(s) from user_consent table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.UserConsent}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userconsent [get]
// http "http://localhost:8080/userconsent?page=0&pagesize=20" X-Api-User:user123
func GetAllUserConsent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "user_consent", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllUserConsent(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetUserConsent is a function to get a single record from the user_consent table in the keycloak database
// @Summary Get record from table UserConsent by  argID
// @Tags UserConsent
// @ID argID
// @Description GetUserConsent is a function to get a single record from the user_consent table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.UserConsent
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /userconsent/{argID} [get]
// http "http://localhost:8080/userconsent/hello world" X-Api-User:user123
func GetUserConsent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_consent", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetUserConsent(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddUserConsent add to add a single record to user_consent table in the keycloak database
// @Summary Add an record to user_consent table
// @Description add to add a single record to user_consent table in the keycloak database
// @Tags UserConsent
// @Accept  json
// @Produce  json
// @Param UserConsent body model.UserConsent true "Add UserConsent"
// @Success 200 {object} model.UserConsent
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userconsent [post]
// echo '{"id": "IZyviYHJTBUQTwyrDfxmQxZTM","client_id": "MxkxyahjivaSuyiKrxAVQUDTJ","user_id": "VpTapSpTfdIUdgeHEupUHgPoM","created_date": 16,"last_updated_date": 22,"client_storage_provider": "sniheGVErcEFSwMLVYHDdNNux","external_client_id": "KbnvKSRZIVLakMsYRwnrmVqPy"}' | http POST "http://localhost:8080/userconsent" X-Api-User:user123
func AddUserConsent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	userconsent := &model.UserConsent{}

	if err := readJSON(r, userconsent); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userconsent.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userconsent.Prepare()

	if err := userconsent.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_consent", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	userconsent, _, err = dao.AddUserConsent(ctx, userconsent)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userconsent)
}

// UpdateUserConsent Update a single record from user_consent table in the keycloak database
// @Summary Update an record in table user_consent
// @Description Update a single record from user_consent table in the keycloak database
// @Tags UserConsent
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  UserConsent body model.UserConsent true "Update UserConsent record"
// @Success 200 {object} model.UserConsent
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /userconsent/{argID} [put]
// echo '{"id": "IZyviYHJTBUQTwyrDfxmQxZTM","client_id": "MxkxyahjivaSuyiKrxAVQUDTJ","user_id": "VpTapSpTfdIUdgeHEupUHgPoM","created_date": 16,"last_updated_date": 22,"client_storage_provider": "sniheGVErcEFSwMLVYHDdNNux","external_client_id": "KbnvKSRZIVLakMsYRwnrmVqPy"}' | http PUT "http://localhost:8080/userconsent/hello world"  X-Api-User:user123
func UpdateUserConsent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	userconsent := &model.UserConsent{}
	if err := readJSON(r, userconsent); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := userconsent.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	userconsent.Prepare()

	if err := userconsent.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_consent", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	userconsent, _, err = dao.UpdateUserConsent(ctx,
		argID,
		userconsent)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, userconsent)
}

// DeleteUserConsent Delete a single record from user_consent table in the keycloak database
// @Summary Delete a record from user_consent
// @Description Delete a single record from user_consent table in the keycloak database
// @Tags UserConsent
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.UserConsent
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /userconsent/{argID} [delete]
// http DELETE "http://localhost:8080/userconsent/hello world" X-Api-User:user123
func DeleteUserConsent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_consent", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteUserConsent(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
