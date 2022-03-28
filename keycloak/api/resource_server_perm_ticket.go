package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configResourceServerPermTicketRouter(router *httprouter.Router) {
	router.GET("/resourceserverpermticket", GetAllResourceServerPermTicket)
	router.POST("/resourceserverpermticket", AddResourceServerPermTicket)
	router.GET("/resourceserverpermticket/:argID", GetResourceServerPermTicket)
	router.PUT("/resourceserverpermticket/:argID", UpdateResourceServerPermTicket)
	router.DELETE("/resourceserverpermticket/:argID", DeleteResourceServerPermTicket)
}

func configGinResourceServerPermTicketRouter(router gin.IRoutes) {
	router.GET("/resourceserverpermticket", ConverHttprouterToGin(GetAllResourceServerPermTicket))
	router.POST("/resourceserverpermticket", ConverHttprouterToGin(AddResourceServerPermTicket))
	router.GET("/resourceserverpermticket/:argID", ConverHttprouterToGin(GetResourceServerPermTicket))
	router.PUT("/resourceserverpermticket/:argID", ConverHttprouterToGin(UpdateResourceServerPermTicket))
	router.DELETE("/resourceserverpermticket/:argID", ConverHttprouterToGin(DeleteResourceServerPermTicket))
}

// GetAllResourceServerPermTicket is a function to get a slice of record(s) from resource_server_perm_ticket table in the keycloak database
// @Summary Get list of ResourceServerPermTicket
// @Tags ResourceServerPermTicket
// @Description GetAllResourceServerPermTicket is a handler to get a slice of record(s) from resource_server_perm_ticket table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ResourceServerPermTicket}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceserverpermticket [get]
// http "http://localhost:8080/resourceserverpermticket?page=0&pagesize=20" X-Api-User:user123
func GetAllResourceServerPermTicket(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "resource_server_perm_ticket", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllResourceServerPermTicket(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetResourceServerPermTicket is a function to get a single record from the resource_server_perm_ticket table in the keycloak database
// @Summary Get record from table ResourceServerPermTicket by  argID
// @Tags ResourceServerPermTicket
// @ID argID
// @Description GetResourceServerPermTicket is a function to get a single record from the resource_server_perm_ticket table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.ResourceServerPermTicket
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /resourceserverpermticket/{argID} [get]
// http "http://localhost:8080/resourceserverpermticket/hello world" X-Api-User:user123
func GetResourceServerPermTicket(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server_perm_ticket", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetResourceServerPermTicket(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddResourceServerPermTicket add to add a single record to resource_server_perm_ticket table in the keycloak database
// @Summary Add an record to resource_server_perm_ticket table
// @Description add to add a single record to resource_server_perm_ticket table in the keycloak database
// @Tags ResourceServerPermTicket
// @Accept  json
// @Produce  json
// @Param ResourceServerPermTicket body model.ResourceServerPermTicket true "Add ResourceServerPermTicket"
// @Success 200 {object} model.ResourceServerPermTicket
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceserverpermticket [post]
// echo '{"id": "bxuMYqMHAtgHwRjoMcjTrrGMI","owner": "JsbsCtAOhMNRenuQgDcKdTpHZ","requester": "yNNvZcjRAuUpAHGAdepPpaMpo","created_timestamp": 7,"granted_timestamp": 94,"resource_id": "uggISDNfestDaNEPByVlDqXkq","scope_id": "IXkelesOkRncHMIrewlVGVGvr","resource_server_id": "lITxJPovTwIswDeLUgQbHtIyv","policy_id": "KSmtYUQnbnAcHtEwJibgHDwJA"}' | http POST "http://localhost:8080/resourceserverpermticket" X-Api-User:user123
func AddResourceServerPermTicket(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	resourceserverpermticket := &model.ResourceServerPermTicket{}

	if err := readJSON(r, resourceserverpermticket); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resourceserverpermticket.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resourceserverpermticket.Prepare()

	if err := resourceserverpermticket.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server_perm_ticket", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	resourceserverpermticket, _, err = dao.AddResourceServerPermTicket(ctx, resourceserverpermticket)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resourceserverpermticket)
}

// UpdateResourceServerPermTicket Update a single record from resource_server_perm_ticket table in the keycloak database
// @Summary Update an record in table resource_server_perm_ticket
// @Description Update a single record from resource_server_perm_ticket table in the keycloak database
// @Tags ResourceServerPermTicket
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  ResourceServerPermTicket body model.ResourceServerPermTicket true "Update ResourceServerPermTicket record"
// @Success 200 {object} model.ResourceServerPermTicket
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resourceserverpermticket/{argID} [put]
// echo '{"id": "bxuMYqMHAtgHwRjoMcjTrrGMI","owner": "JsbsCtAOhMNRenuQgDcKdTpHZ","requester": "yNNvZcjRAuUpAHGAdepPpaMpo","created_timestamp": 7,"granted_timestamp": 94,"resource_id": "uggISDNfestDaNEPByVlDqXkq","scope_id": "IXkelesOkRncHMIrewlVGVGvr","resource_server_id": "lITxJPovTwIswDeLUgQbHtIyv","policy_id": "KSmtYUQnbnAcHtEwJibgHDwJA"}' | http PUT "http://localhost:8080/resourceserverpermticket/hello world"  X-Api-User:user123
func UpdateResourceServerPermTicket(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resourceserverpermticket := &model.ResourceServerPermTicket{}
	if err := readJSON(r, resourceserverpermticket); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resourceserverpermticket.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resourceserverpermticket.Prepare()

	if err := resourceserverpermticket.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server_perm_ticket", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resourceserverpermticket, _, err = dao.UpdateResourceServerPermTicket(ctx,
		argID,
		resourceserverpermticket)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resourceserverpermticket)
}

// DeleteResourceServerPermTicket Delete a single record from resource_server_perm_ticket table in the keycloak database
// @Summary Delete a record from resource_server_perm_ticket
// @Description Delete a single record from resource_server_perm_ticket table in the keycloak database
// @Tags ResourceServerPermTicket
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.ResourceServerPermTicket
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /resourceserverpermticket/{argID} [delete]
// http DELETE "http://localhost:8080/resourceserverpermticket/hello world" X-Api-User:user123
func DeleteResourceServerPermTicket(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resource_server_perm_ticket", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteResourceServerPermTicket(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
