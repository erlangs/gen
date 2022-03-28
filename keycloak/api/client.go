package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configClientRouter(router *httprouter.Router) {
	router.GET("/client", GetAllClient)
	router.POST("/client", AddClient)
	router.GET("/client/:argID", GetClient)
	router.PUT("/client/:argID", UpdateClient)
	router.DELETE("/client/:argID", DeleteClient)
}

func configGinClientRouter(router gin.IRoutes) {
	router.GET("/client", ConverHttprouterToGin(GetAllClient))
	router.POST("/client", ConverHttprouterToGin(AddClient))
	router.GET("/client/:argID", ConverHttprouterToGin(GetClient))
	router.PUT("/client/:argID", ConverHttprouterToGin(UpdateClient))
	router.DELETE("/client/:argID", ConverHttprouterToGin(DeleteClient))
}

// GetAllClient is a function to get a slice of record(s) from client table in the keycloak database
// @Summary Get list of Client
// @Tags Client
// @Description GetAllClient is a handler to get a slice of record(s) from client table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Client}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /client [get]
// http "http://localhost:8080/client?page=0&pagesize=20" X-Api-User:user123
func GetAllClient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "client", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllClient(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetClient is a function to get a single record from the client table in the keycloak database
// @Summary Get record from table Client by  argID
// @Tags Client
// @ID argID
// @Description GetClient is a function to get a single record from the client table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.Client
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /client/{argID} [get]
// http "http://localhost:8080/client/hello world" X-Api-User:user123
func GetClient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetClient(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddClient add to add a single record to client table in the keycloak database
// @Summary Add an record to client table
// @Description add to add a single record to client table in the keycloak database
// @Tags Client
// @Accept  json
// @Produce  json
// @Param Client body model.Client true "Add Client"
// @Success 200 {object} model.Client
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /client [post]
// echo '{"id": "bFZVTdODDpaofiKGSVhYHCNKh","enabled": false,"full_scope_allowed": true,"client_id": "wMPBREIMWQeaHvvEvCjimcYMX","not_before": 75,"public_client": false,"secret": "XrRGwJxfrfuMyxrgbHxXiYxVt","base_url": "sFVLASnWUdhvvbUYLXXlfvCjU","bearer_only": false,"management_url": "LfurIbwXMqtfxKSgRIaeBvEMa","surrogate_auth_required": true,"realm_id": "QvQiAbfInsunFPNyoSvShWNOt","protocol": "uhSupLDqunxuVAAFmHsDFGTRZ","node_rereg_timeout": 75,"frontchannel_logout": false,"consent_required": true,"name": "KWqghhmPdGjmrxmTZItUQOMeI","service_accounts_enabled": false,"client_authenticator_type": "nExWNXoqpAnkmNZaaqwrMcEpq","root_url": "eArkskOBLPtygCsOpBPXjqoNu","description": "EhbwlTtUXYsbNZcrXUpJLenCY","registration_token": "DfjJeNvUSRbSvqJJPpsZVCcDj","standard_flow_enabled": false,"implicit_flow_enabled": false,"direct_access_grants_enabled": true,"always_display_in_console": false}' | http POST "http://localhost:8080/client" X-Api-User:user123
func AddClient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	client := &model.Client{}

	if err := readJSON(r, client); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := client.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	client.Prepare()

	if err := client.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	client, _, err = dao.AddClient(ctx, client)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, client)
}

// UpdateClient Update a single record from client table in the keycloak database
// @Summary Update an record in table client
// @Description Update a single record from client table in the keycloak database
// @Tags Client
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  Client body model.Client true "Update Client record"
// @Success 200 {object} model.Client
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /client/{argID} [put]
// echo '{"id": "bFZVTdODDpaofiKGSVhYHCNKh","enabled": false,"full_scope_allowed": true,"client_id": "wMPBREIMWQeaHvvEvCjimcYMX","not_before": 75,"public_client": false,"secret": "XrRGwJxfrfuMyxrgbHxXiYxVt","base_url": "sFVLASnWUdhvvbUYLXXlfvCjU","bearer_only": false,"management_url": "LfurIbwXMqtfxKSgRIaeBvEMa","surrogate_auth_required": true,"realm_id": "QvQiAbfInsunFPNyoSvShWNOt","protocol": "uhSupLDqunxuVAAFmHsDFGTRZ","node_rereg_timeout": 75,"frontchannel_logout": false,"consent_required": true,"name": "KWqghhmPdGjmrxmTZItUQOMeI","service_accounts_enabled": false,"client_authenticator_type": "nExWNXoqpAnkmNZaaqwrMcEpq","root_url": "eArkskOBLPtygCsOpBPXjqoNu","description": "EhbwlTtUXYsbNZcrXUpJLenCY","registration_token": "DfjJeNvUSRbSvqJJPpsZVCcDj","standard_flow_enabled": false,"implicit_flow_enabled": false,"direct_access_grants_enabled": true,"always_display_in_console": false}' | http PUT "http://localhost:8080/client/hello world"  X-Api-User:user123
func UpdateClient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	client := &model.Client{}
	if err := readJSON(r, client); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := client.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	client.Prepare()

	if err := client.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	client, _, err = dao.UpdateClient(ctx,
		argID,
		client)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, client)
}

// DeleteClient Delete a single record from client table in the keycloak database
// @Summary Delete a record from client
// @Description Delete a single record from client table in the keycloak database
// @Tags Client
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.Client
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /client/{argID} [delete]
// http DELETE "http://localhost:8080/client/hello world" X-Api-User:user123
func DeleteClient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteClient(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
