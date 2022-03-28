package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configCredentialRouter(router *httprouter.Router) {
	router.GET("/credential", GetAllCredential)
	router.POST("/credential", AddCredential)
	router.GET("/credential/:argID", GetCredential)
	router.PUT("/credential/:argID", UpdateCredential)
	router.DELETE("/credential/:argID", DeleteCredential)
}

func configGinCredentialRouter(router gin.IRoutes) {
	router.GET("/credential", ConverHttprouterToGin(GetAllCredential))
	router.POST("/credential", ConverHttprouterToGin(AddCredential))
	router.GET("/credential/:argID", ConverHttprouterToGin(GetCredential))
	router.PUT("/credential/:argID", ConverHttprouterToGin(UpdateCredential))
	router.DELETE("/credential/:argID", ConverHttprouterToGin(DeleteCredential))
}

// GetAllCredential is a function to get a slice of record(s) from credential table in the keycloak database
// @Summary Get list of Credential
// @Tags Credential
// @Description GetAllCredential is a handler to get a slice of record(s) from credential table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Credential}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /credential [get]
// http "http://localhost:8080/credential?page=0&pagesize=20" X-Api-User:user123
func GetAllCredential(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "credential", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllCredential(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetCredential is a function to get a single record from the credential table in the keycloak database
// @Summary Get record from table Credential by  argID
// @Tags Credential
// @ID argID
// @Description GetCredential is a function to get a single record from the credential table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.Credential
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /credential/{argID} [get]
// http "http://localhost:8080/credential/hello world" X-Api-User:user123
func GetCredential(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "credential", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetCredential(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddCredential add to add a single record to credential table in the keycloak database
// @Summary Add an record to credential table
// @Description add to add a single record to credential table in the keycloak database
// @Tags Credential
// @Accept  json
// @Produce  json
// @Param Credential body model.Credential true "Add Credential"
// @Success 200 {object} model.Credential
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /credential [post]
// echo '{"id": "GisBamUHwPBCMcXRvJLoTrMFy","salt": "cXCtperDeHigirgMYryMdcwEw","type": "hiJXfcRXcOVCXTsViWMvPscxe","user_id": "RRwkouTMiXQNPtAijaZgiWFAZ","created_date": 33,"user_label": "wxgvZKODKIHQRQcoGepvtXQHH","secret_data": "atyMFyiPiUsgIyGyKkQSqRLtY","credential_data": "RaYNcZDeGAqvKuWLQqRMSrngF","priority": 49}' | http POST "http://localhost:8080/credential" X-Api-User:user123
func AddCredential(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	credential := &model.Credential{}

	if err := readJSON(r, credential); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := credential.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	credential.Prepare()

	if err := credential.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "credential", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	credential, _, err = dao.AddCredential(ctx, credential)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, credential)
}

// UpdateCredential Update a single record from credential table in the keycloak database
// @Summary Update an record in table credential
// @Description Update a single record from credential table in the keycloak database
// @Tags Credential
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  Credential body model.Credential true "Update Credential record"
// @Success 200 {object} model.Credential
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /credential/{argID} [put]
// echo '{"id": "GisBamUHwPBCMcXRvJLoTrMFy","salt": "cXCtperDeHigirgMYryMdcwEw","type": "hiJXfcRXcOVCXTsViWMvPscxe","user_id": "RRwkouTMiXQNPtAijaZgiWFAZ","created_date": 33,"user_label": "wxgvZKODKIHQRQcoGepvtXQHH","secret_data": "atyMFyiPiUsgIyGyKkQSqRLtY","credential_data": "RaYNcZDeGAqvKuWLQqRMSrngF","priority": 49}' | http PUT "http://localhost:8080/credential/hello world"  X-Api-User:user123
func UpdateCredential(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	credential := &model.Credential{}
	if err := readJSON(r, credential); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := credential.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	credential.Prepare()

	if err := credential.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "credential", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	credential, _, err = dao.UpdateCredential(ctx,
		argID,
		credential)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, credential)
}

// DeleteCredential Delete a single record from credential table in the keycloak database
// @Summary Delete a record from credential
// @Description Delete a single record from credential table in the keycloak database
// @Tags Credential
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.Credential
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /credential/{argID} [delete]
// http DELETE "http://localhost:8080/credential/hello world" X-Api-User:user123
func DeleteCredential(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "credential", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteCredential(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
