package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configBrokerLinkRouter(router *httprouter.Router) {
	router.GET("/brokerlink", GetAllBrokerLink)
	router.POST("/brokerlink", AddBrokerLink)
	router.GET("/brokerlink/:argIdentityProvider/:argUserID", GetBrokerLink)
	router.PUT("/brokerlink/:argIdentityProvider/:argUserID", UpdateBrokerLink)
	router.DELETE("/brokerlink/:argIdentityProvider/:argUserID", DeleteBrokerLink)
}

func configGinBrokerLinkRouter(router gin.IRoutes) {
	router.GET("/brokerlink", ConverHttprouterToGin(GetAllBrokerLink))
	router.POST("/brokerlink", ConverHttprouterToGin(AddBrokerLink))
	router.GET("/brokerlink/:argIdentityProvider/:argUserID", ConverHttprouterToGin(GetBrokerLink))
	router.PUT("/brokerlink/:argIdentityProvider/:argUserID", ConverHttprouterToGin(UpdateBrokerLink))
	router.DELETE("/brokerlink/:argIdentityProvider/:argUserID", ConverHttprouterToGin(DeleteBrokerLink))
}

// GetAllBrokerLink is a function to get a slice of record(s) from broker_link table in the keycloak database
// @Summary Get list of BrokerLink
// @Tags BrokerLink
// @Description GetAllBrokerLink is a handler to get a slice of record(s) from broker_link table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.BrokerLink}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /brokerlink [get]
// http "http://localhost:8080/brokerlink?page=0&pagesize=20" X-Api-User:user123
func GetAllBrokerLink(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "broker_link", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBrokerLink(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetBrokerLink is a function to get a single record from the broker_link table in the keycloak database
// @Summary Get record from table BrokerLink by  argIdentityProvider  argUserID
// @Tags BrokerLink
// @ID argIdentityProvider
// @ID argUserID
// @Description GetBrokerLink is a function to get a single record from the broker_link table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argIdentityProvider path string true "identity_provider"
// @Param  argUserID path string true "user_id"
// @Success 200 {object} model.BrokerLink
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /brokerlink/{argIdentityProvider}/{argUserID} [get]
// http "http://localhost:8080/brokerlink/hello world/hello world" X-Api-User:user123
func GetBrokerLink(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argIdentityProvider, err := parseString(ps, "argIdentityProvider")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argUserID, err := parseString(ps, "argUserID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "broker_link", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBrokerLink(ctx, argIdentityProvider, argUserID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBrokerLink add to add a single record to broker_link table in the keycloak database
// @Summary Add an record to broker_link table
// @Description add to add a single record to broker_link table in the keycloak database
// @Tags BrokerLink
// @Accept  json
// @Produce  json
// @Param BrokerLink body model.BrokerLink true "Add BrokerLink"
// @Success 200 {object} model.BrokerLink
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /brokerlink [post]
// echo '{"identity_provider": "QEurbkYQdyPMYqjOaqyaygvYB","storage_provider_id": "VxncJhRcRbLNyyricLWbbTCXE","realm_id": "nCjjpkVMBmLfDptfjNRYKpUmq","broker_user_id": "KCBuppUsptYPfvFUhXPnTBdhh","broker_username": "kACOSGEJkiTJdaYGuNssIvtyJ","token": "RuQkoSeviRaGkEPCwWZkBHWln","user_id": "JqFQICbDdRyWJQaIRxNKTVPXI"}' | http POST "http://localhost:8080/brokerlink" X-Api-User:user123
func AddBrokerLink(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	brokerlink := &model.BrokerLink{}

	if err := readJSON(r, brokerlink); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := brokerlink.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	brokerlink.Prepare()

	if err := brokerlink.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "broker_link", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	brokerlink, _, err = dao.AddBrokerLink(ctx, brokerlink)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, brokerlink)
}

// UpdateBrokerLink Update a single record from broker_link table in the keycloak database
// @Summary Update an record in table broker_link
// @Description Update a single record from broker_link table in the keycloak database
// @Tags BrokerLink
// @Accept  json
// @Produce  json
// @Param  argIdentityProvider path string true "identity_provider"// @Param  argUserID path string true "user_id"
// @Param  BrokerLink body model.BrokerLink true "Update BrokerLink record"
// @Success 200 {object} model.BrokerLink
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /brokerlink/{argIdentityProvider}/{argUserID} [put]
// echo '{"identity_provider": "QEurbkYQdyPMYqjOaqyaygvYB","storage_provider_id": "VxncJhRcRbLNyyricLWbbTCXE","realm_id": "nCjjpkVMBmLfDptfjNRYKpUmq","broker_user_id": "KCBuppUsptYPfvFUhXPnTBdhh","broker_username": "kACOSGEJkiTJdaYGuNssIvtyJ","token": "RuQkoSeviRaGkEPCwWZkBHWln","user_id": "JqFQICbDdRyWJQaIRxNKTVPXI"}' | http PUT "http://localhost:8080/brokerlink/hello world/hello world"  X-Api-User:user123
func UpdateBrokerLink(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argIdentityProvider, err := parseString(ps, "argIdentityProvider")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argUserID, err := parseString(ps, "argUserID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	brokerlink := &model.BrokerLink{}
	if err := readJSON(r, brokerlink); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := brokerlink.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	brokerlink.Prepare()

	if err := brokerlink.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "broker_link", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	brokerlink, _, err = dao.UpdateBrokerLink(ctx,
		argIdentityProvider, argUserID,
		brokerlink)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, brokerlink)
}

// DeleteBrokerLink Delete a single record from broker_link table in the keycloak database
// @Summary Delete a record from broker_link
// @Description Delete a single record from broker_link table in the keycloak database
// @Tags BrokerLink
// @Accept  json
// @Produce  json
// @Param  argIdentityProvider path string true "identity_provider"// @Param  argUserID path string true "user_id"
// @Success 204 {object} model.BrokerLink
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /brokerlink/{argIdentityProvider}/{argUserID} [delete]
// http DELETE "http://localhost:8080/brokerlink/hello world/hello world" X-Api-User:user123
func DeleteBrokerLink(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argIdentityProvider, err := parseString(ps, "argIdentityProvider")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argUserID, err := parseString(ps, "argUserID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "broker_link", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBrokerLink(ctx, argIdentityProvider, argUserID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
