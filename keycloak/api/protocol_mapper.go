package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configProtocolMapperRouter(router *httprouter.Router) {
	router.GET("/protocolmapper", GetAllProtocolMapper)
	router.POST("/protocolmapper", AddProtocolMapper)
	router.GET("/protocolmapper/:argID", GetProtocolMapper)
	router.PUT("/protocolmapper/:argID", UpdateProtocolMapper)
	router.DELETE("/protocolmapper/:argID", DeleteProtocolMapper)
}

func configGinProtocolMapperRouter(router gin.IRoutes) {
	router.GET("/protocolmapper", ConverHttprouterToGin(GetAllProtocolMapper))
	router.POST("/protocolmapper", ConverHttprouterToGin(AddProtocolMapper))
	router.GET("/protocolmapper/:argID", ConverHttprouterToGin(GetProtocolMapper))
	router.PUT("/protocolmapper/:argID", ConverHttprouterToGin(UpdateProtocolMapper))
	router.DELETE("/protocolmapper/:argID", ConverHttprouterToGin(DeleteProtocolMapper))
}

// GetAllProtocolMapper is a function to get a slice of record(s) from protocol_mapper table in the keycloak database
// @Summary Get list of ProtocolMapper
// @Tags ProtocolMapper
// @Description GetAllProtocolMapper is a handler to get a slice of record(s) from protocol_mapper table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ProtocolMapper}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /protocolmapper [get]
// http "http://localhost:8080/protocolmapper?page=0&pagesize=20" X-Api-User:user123
func GetAllProtocolMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "protocol_mapper", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllProtocolMapper(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetProtocolMapper is a function to get a single record from the protocol_mapper table in the keycloak database
// @Summary Get record from table ProtocolMapper by  argID
// @Tags ProtocolMapper
// @ID argID
// @Description GetProtocolMapper is a function to get a single record from the protocol_mapper table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.ProtocolMapper
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /protocolmapper/{argID} [get]
// http "http://localhost:8080/protocolmapper/hello world" X-Api-User:user123
func GetProtocolMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "protocol_mapper", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetProtocolMapper(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddProtocolMapper add to add a single record to protocol_mapper table in the keycloak database
// @Summary Add an record to protocol_mapper table
// @Description add to add a single record to protocol_mapper table in the keycloak database
// @Tags ProtocolMapper
// @Accept  json
// @Produce  json
// @Param ProtocolMapper body model.ProtocolMapper true "Add ProtocolMapper"
// @Success 200 {object} model.ProtocolMapper
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /protocolmapper [post]
// echo '{"id": "ybyCOswbKmTGWRsKlOqjkQGtE","name": "IKLFVrTeHuOVfuThJfpMgoDqD","protocol": "DcfnDOYerQWlFwkhLwjCumYLK","protocol_mapper_name": "NZHfHQUQHeMUSLbEylmjOdrrs","client_id": "trtJoWPvbaNiINWVKEMeatnik","client_scope_id": "mPvqMLWZHGmiPVqHckdpUoTTU"}' | http POST "http://localhost:8080/protocolmapper" X-Api-User:user123
func AddProtocolMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	protocolmapper := &model.ProtocolMapper{}

	if err := readJSON(r, protocolmapper); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := protocolmapper.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	protocolmapper.Prepare()

	if err := protocolmapper.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "protocol_mapper", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	protocolmapper, _, err = dao.AddProtocolMapper(ctx, protocolmapper)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, protocolmapper)
}

// UpdateProtocolMapper Update a single record from protocol_mapper table in the keycloak database
// @Summary Update an record in table protocol_mapper
// @Description Update a single record from protocol_mapper table in the keycloak database
// @Tags ProtocolMapper
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  ProtocolMapper body model.ProtocolMapper true "Update ProtocolMapper record"
// @Success 200 {object} model.ProtocolMapper
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /protocolmapper/{argID} [put]
// echo '{"id": "ybyCOswbKmTGWRsKlOqjkQGtE","name": "IKLFVrTeHuOVfuThJfpMgoDqD","protocol": "DcfnDOYerQWlFwkhLwjCumYLK","protocol_mapper_name": "NZHfHQUQHeMUSLbEylmjOdrrs","client_id": "trtJoWPvbaNiINWVKEMeatnik","client_scope_id": "mPvqMLWZHGmiPVqHckdpUoTTU"}' | http PUT "http://localhost:8080/protocolmapper/hello world"  X-Api-User:user123
func UpdateProtocolMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	protocolmapper := &model.ProtocolMapper{}
	if err := readJSON(r, protocolmapper); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := protocolmapper.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	protocolmapper.Prepare()

	if err := protocolmapper.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "protocol_mapper", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	protocolmapper, _, err = dao.UpdateProtocolMapper(ctx,
		argID,
		protocolmapper)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, protocolmapper)
}

// DeleteProtocolMapper Delete a single record from protocol_mapper table in the keycloak database
// @Summary Delete a record from protocol_mapper
// @Description Delete a single record from protocol_mapper table in the keycloak database
// @Tags ProtocolMapper
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.ProtocolMapper
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /protocolmapper/{argID} [delete]
// http DELETE "http://localhost:8080/protocolmapper/hello world" X-Api-User:user123
func DeleteProtocolMapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "protocol_mapper", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteProtocolMapper(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
