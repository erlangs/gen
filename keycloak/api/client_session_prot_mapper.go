package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configClientSessionProtMapperRouter(router *httprouter.Router) {
	router.GET("/clientsessionprotmapper", GetAllClientSessionProtMapper)
	router.POST("/clientsessionprotmapper", AddClientSessionProtMapper)
	router.GET("/clientsessionprotmapper/:argProtocolMapperID/:argClientSession", GetClientSessionProtMapper)
	router.PUT("/clientsessionprotmapper/:argProtocolMapperID/:argClientSession", UpdateClientSessionProtMapper)
	router.DELETE("/clientsessionprotmapper/:argProtocolMapperID/:argClientSession", DeleteClientSessionProtMapper)
}

func configGinClientSessionProtMapperRouter(router gin.IRoutes) {
	router.GET("/clientsessionprotmapper", ConverHttprouterToGin(GetAllClientSessionProtMapper))
	router.POST("/clientsessionprotmapper", ConverHttprouterToGin(AddClientSessionProtMapper))
	router.GET("/clientsessionprotmapper/:argProtocolMapperID/:argClientSession", ConverHttprouterToGin(GetClientSessionProtMapper))
	router.PUT("/clientsessionprotmapper/:argProtocolMapperID/:argClientSession", ConverHttprouterToGin(UpdateClientSessionProtMapper))
	router.DELETE("/clientsessionprotmapper/:argProtocolMapperID/:argClientSession", ConverHttprouterToGin(DeleteClientSessionProtMapper))
}

// GetAllClientSessionProtMapper is a function to get a slice of record(s) from client_session_prot_mapper table in the keycloak database
// @Summary Get list of ClientSessionProtMapper
// @Tags ClientSessionProtMapper
// @Description GetAllClientSessionProtMapper is a handler to get a slice of record(s) from client_session_prot_mapper table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ClientSessionProtMapper}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientsessionprotmapper [get]
// http "http://localhost:8080/clientsessionprotmapper?page=0&pagesize=20" X-Api-User:user123
func GetAllClientSessionProtMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "client_session_prot_mapper", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllClientSessionProtMapper(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetClientSessionProtMapper is a function to get a single record from the client_session_prot_mapper table in the keycloak database
// @Summary Get record from table ClientSessionProtMapper by  argProtocolMapperID  argClientSession
// @Tags ClientSessionProtMapper
// @ID argProtocolMapperID
// @ID argClientSession
// @Description GetClientSessionProtMapper is a function to get a single record from the client_session_prot_mapper table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argProtocolMapperID path string true "protocol_mapper_id"
// @Param  argClientSession path string true "client_session"
// @Success 200 {object} model.ClientSessionProtMapper
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /clientsessionprotmapper/{argProtocolMapperID}/{argClientSession} [get]
// http "http://localhost:8080/clientsessionprotmapper/hello world/hello world" X-Api-User:user123
func GetClientSessionProtMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argProtocolMapperID, err := parseString(ps, "argProtocolMapperID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argClientSession, err := parseString(ps, "argClientSession")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_session_prot_mapper", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetClientSessionProtMapper(ctx, argProtocolMapperID, argClientSession)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddClientSessionProtMapper add to add a single record to client_session_prot_mapper table in the keycloak database
// @Summary Add an record to client_session_prot_mapper table
// @Description add to add a single record to client_session_prot_mapper table in the keycloak database
// @Tags ClientSessionProtMapper
// @Accept  json
// @Produce  json
// @Param ClientSessionProtMapper body model.ClientSessionProtMapper true "Add ClientSessionProtMapper"
// @Success 200 {object} model.ClientSessionProtMapper
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientsessionprotmapper [post]
// echo '{"protocol_mapper_id": "GYIsBvaYKXNiSdrcQOxwadfiK","client_session": "LhYoGBJMsowcGOZlYEOABQwto"}' | http POST "http://localhost:8080/clientsessionprotmapper" X-Api-User:user123
func AddClientSessionProtMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	clientsessionprotmapper := &model.ClientSessionProtMapper{}

	if err := readJSON(r, clientsessionprotmapper); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientsessionprotmapper.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientsessionprotmapper.Prepare()

	if err := clientsessionprotmapper.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_session_prot_mapper", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	clientsessionprotmapper, _, err = dao.AddClientSessionProtMapper(ctx, clientsessionprotmapper)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientsessionprotmapper)
}

// UpdateClientSessionProtMapper Update a single record from client_session_prot_mapper table in the keycloak database
// @Summary Update an record in table client_session_prot_mapper
// @Description Update a single record from client_session_prot_mapper table in the keycloak database
// @Tags ClientSessionProtMapper
// @Accept  json
// @Produce  json
// @Param  argProtocolMapperID path string true "protocol_mapper_id"// @Param  argClientSession path string true "client_session"
// @Param  ClientSessionProtMapper body model.ClientSessionProtMapper true "Update ClientSessionProtMapper record"
// @Success 200 {object} model.ClientSessionProtMapper
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientsessionprotmapper/{argProtocolMapperID}/{argClientSession} [put]
// echo '{"protocol_mapper_id": "GYIsBvaYKXNiSdrcQOxwadfiK","client_session": "LhYoGBJMsowcGOZlYEOABQwto"}' | http PUT "http://localhost:8080/clientsessionprotmapper/hello world/hello world"  X-Api-User:user123
func UpdateClientSessionProtMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argProtocolMapperID, err := parseString(ps, "argProtocolMapperID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argClientSession, err := parseString(ps, "argClientSession")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientsessionprotmapper := &model.ClientSessionProtMapper{}
	if err := readJSON(r, clientsessionprotmapper); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientsessionprotmapper.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientsessionprotmapper.Prepare()

	if err := clientsessionprotmapper.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_session_prot_mapper", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientsessionprotmapper, _, err = dao.UpdateClientSessionProtMapper(ctx,
		argProtocolMapperID, argClientSession,
		clientsessionprotmapper)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientsessionprotmapper)
}

// DeleteClientSessionProtMapper Delete a single record from client_session_prot_mapper table in the keycloak database
// @Summary Delete a record from client_session_prot_mapper
// @Description Delete a single record from client_session_prot_mapper table in the keycloak database
// @Tags ClientSessionProtMapper
// @Accept  json
// @Produce  json
// @Param  argProtocolMapperID path string true "protocol_mapper_id"// @Param  argClientSession path string true "client_session"
// @Success 204 {object} model.ClientSessionProtMapper
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /clientsessionprotmapper/{argProtocolMapperID}/{argClientSession} [delete]
// http DELETE "http://localhost:8080/clientsessionprotmapper/hello world/hello world" X-Api-User:user123
func DeleteClientSessionProtMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argProtocolMapperID, err := parseString(ps, "argProtocolMapperID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argClientSession, err := parseString(ps, "argClientSession")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_session_prot_mapper", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteClientSessionProtMapper(ctx, argProtocolMapperID, argClientSession)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
