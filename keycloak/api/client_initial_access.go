package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configClientInitialAccessRouter(router *httprouter.Router) {
	router.GET("/clientinitialaccess", GetAllClientInitialAccess)
	router.POST("/clientinitialaccess", AddClientInitialAccess)
	router.GET("/clientinitialaccess/:argID", GetClientInitialAccess)
	router.PUT("/clientinitialaccess/:argID", UpdateClientInitialAccess)
	router.DELETE("/clientinitialaccess/:argID", DeleteClientInitialAccess)
}

func configGinClientInitialAccessRouter(router gin.IRoutes) {
	router.GET("/clientinitialaccess", ConverHttprouterToGin(GetAllClientInitialAccess))
	router.POST("/clientinitialaccess", ConverHttprouterToGin(AddClientInitialAccess))
	router.GET("/clientinitialaccess/:argID", ConverHttprouterToGin(GetClientInitialAccess))
	router.PUT("/clientinitialaccess/:argID", ConverHttprouterToGin(UpdateClientInitialAccess))
	router.DELETE("/clientinitialaccess/:argID", ConverHttprouterToGin(DeleteClientInitialAccess))
}

// GetAllClientInitialAccess is a function to get a slice of record(s) from client_initial_access table in the keycloak database
// @Summary Get list of ClientInitialAccess
// @Tags ClientInitialAccess
// @Description GetAllClientInitialAccess is a handler to get a slice of record(s) from client_initial_access table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ClientInitialAccess}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientinitialaccess [get]
// http "http://localhost:8080/clientinitialaccess?page=0&pagesize=20" X-Api-User:user123
func GetAllClientInitialAccess(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "client_initial_access", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllClientInitialAccess(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetClientInitialAccess is a function to get a single record from the client_initial_access table in the keycloak database
// @Summary Get record from table ClientInitialAccess by  argID
// @Tags ClientInitialAccess
// @ID argID
// @Description GetClientInitialAccess is a function to get a single record from the client_initial_access table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.ClientInitialAccess
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /clientinitialaccess/{argID} [get]
// http "http://localhost:8080/clientinitialaccess/hello world" X-Api-User:user123
func GetClientInitialAccess(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_initial_access", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetClientInitialAccess(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddClientInitialAccess add to add a single record to client_initial_access table in the keycloak database
// @Summary Add an record to client_initial_access table
// @Description add to add a single record to client_initial_access table in the keycloak database
// @Tags ClientInitialAccess
// @Accept  json
// @Produce  json
// @Param ClientInitialAccess body model.ClientInitialAccess true "Add ClientInitialAccess"
// @Success 200 {object} model.ClientInitialAccess
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientinitialaccess [post]
// echo '{"id": "yqRZvriWlmsRgscQODAueWfte","realm_id": "MNycTiqwJDSRXbhdUlOYdGwLn","timestamp": 83,"expiration": 76,"count": 21,"remaining_count": 15}' | http POST "http://localhost:8080/clientinitialaccess" X-Api-User:user123
func AddClientInitialAccess(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	clientinitialaccess := &model.ClientInitialAccess{}

	if err := readJSON(r, clientinitialaccess); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientinitialaccess.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientinitialaccess.Prepare()

	if err := clientinitialaccess.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_initial_access", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	clientinitialaccess, _, err = dao.AddClientInitialAccess(ctx, clientinitialaccess)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientinitialaccess)
}

// UpdateClientInitialAccess Update a single record from client_initial_access table in the keycloak database
// @Summary Update an record in table client_initial_access
// @Description Update a single record from client_initial_access table in the keycloak database
// @Tags ClientInitialAccess
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  ClientInitialAccess body model.ClientInitialAccess true "Update ClientInitialAccess record"
// @Success 200 {object} model.ClientInitialAccess
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientinitialaccess/{argID} [put]
// echo '{"id": "yqRZvriWlmsRgscQODAueWfte","realm_id": "MNycTiqwJDSRXbhdUlOYdGwLn","timestamp": 83,"expiration": 76,"count": 21,"remaining_count": 15}' | http PUT "http://localhost:8080/clientinitialaccess/hello world"  X-Api-User:user123
func UpdateClientInitialAccess(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientinitialaccess := &model.ClientInitialAccess{}
	if err := readJSON(r, clientinitialaccess); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientinitialaccess.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientinitialaccess.Prepare()

	if err := clientinitialaccess.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_initial_access", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientinitialaccess, _, err = dao.UpdateClientInitialAccess(ctx,
		argID,
		clientinitialaccess)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientinitialaccess)
}

// DeleteClientInitialAccess Delete a single record from client_initial_access table in the keycloak database
// @Summary Delete a record from client_initial_access
// @Description Delete a single record from client_initial_access table in the keycloak database
// @Tags ClientInitialAccess
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.ClientInitialAccess
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /clientinitialaccess/{argID} [delete]
// http DELETE "http://localhost:8080/clientinitialaccess/hello world" X-Api-User:user123
func DeleteClientInitialAccess(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_initial_access", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteClientInitialAccess(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
