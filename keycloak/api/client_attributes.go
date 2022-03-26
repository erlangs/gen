package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configClientAttributesRouter(router *httprouter.Router) {
	router.GET("/clientattributes", GetAllClientAttributes)
	router.POST("/clientattributes", AddClientAttributes)
	router.GET("/clientattributes/:argClientID/:argName", GetClientAttributes)
	router.PUT("/clientattributes/:argClientID/:argName", UpdateClientAttributes)
	router.DELETE("/clientattributes/:argClientID/:argName", DeleteClientAttributes)
}

func configGinClientAttributesRouter(router gin.IRoutes) {
	router.GET("/clientattributes", ConverHttprouterToGin(GetAllClientAttributes))
	router.POST("/clientattributes", ConverHttprouterToGin(AddClientAttributes))
	router.GET("/clientattributes/:argClientID/:argName", ConverHttprouterToGin(GetClientAttributes))
	router.PUT("/clientattributes/:argClientID/:argName", ConverHttprouterToGin(UpdateClientAttributes))
	router.DELETE("/clientattributes/:argClientID/:argName", ConverHttprouterToGin(DeleteClientAttributes))
}

// GetAllClientAttributes is a function to get a slice of record(s) from client_attributes table in the keycloak database
// @Summary Get list of ClientAttributes
// @Tags ClientAttributes
// @Description GetAllClientAttributes is a handler to get a slice of record(s) from client_attributes table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ClientAttributes}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientattributes [get]
// http "http://localhost:8080/clientattributes?page=0&pagesize=20" X-Api-User:user123
func GetAllClientAttributes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "client_attributes", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllClientAttributes(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetClientAttributes is a function to get a single record from the client_attributes table in the keycloak database
// @Summary Get record from table ClientAttributes by  argClientID  argName
// @Tags ClientAttributes
// @ID argClientID
// @ID argName
// @Description GetClientAttributes is a function to get a single record from the client_attributes table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"
// @Param  argName path string true "name"
// @Success 200 {object} model.ClientAttributes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /clientattributes/{argClientID}/{argName} [get]
// http "http://localhost:8080/clientattributes/hello world/hello world" X-Api-User:user123
func GetClientAttributes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_attributes", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetClientAttributes(ctx, argClientID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddClientAttributes add to add a single record to client_attributes table in the keycloak database
// @Summary Add an record to client_attributes table
// @Description add to add a single record to client_attributes table in the keycloak database
// @Tags ClientAttributes
// @Accept  json
// @Produce  json
// @Param ClientAttributes body model.ClientAttributes true "Add ClientAttributes"
// @Success 200 {object} model.ClientAttributes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientattributes [post]
// echo '{"client_id": "PanbPUtJqpMNNmaXlnVfWSGPR","value": "LCfnLZflMGQoeHQcZMsDkehji","name": "cxKnfPqjdfHdaTJhWmZjEqwJa"}' | http POST "http://localhost:8080/clientattributes" X-Api-User:user123
func AddClientAttributes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	clientattributes := &model.ClientAttributes{}

	if err := readJSON(r, clientattributes); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientattributes.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientattributes.Prepare()

	if err := clientattributes.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_attributes", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	clientattributes, _, err = dao.AddClientAttributes(ctx, clientattributes)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientattributes)
}

// UpdateClientAttributes Update a single record from client_attributes table in the keycloak database
// @Summary Update an record in table client_attributes
// @Description Update a single record from client_attributes table in the keycloak database
// @Tags ClientAttributes
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"// @Param  argName path string true "name"
// @Param  ClientAttributes body model.ClientAttributes true "Update ClientAttributes record"
// @Success 200 {object} model.ClientAttributes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientattributes/{argClientID}/{argName} [put]
// echo '{"client_id": "PanbPUtJqpMNNmaXlnVfWSGPR","value": "LCfnLZflMGQoeHQcZMsDkehji","name": "cxKnfPqjdfHdaTJhWmZjEqwJa"}' | http PUT "http://localhost:8080/clientattributes/hello world/hello world"  X-Api-User:user123
func UpdateClientAttributes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientattributes := &model.ClientAttributes{}
	if err := readJSON(r, clientattributes); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientattributes.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientattributes.Prepare()

	if err := clientattributes.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_attributes", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientattributes, _, err = dao.UpdateClientAttributes(ctx,
		argClientID, argName,
		clientattributes)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientattributes)
}

// DeleteClientAttributes Delete a single record from client_attributes table in the keycloak database
// @Summary Delete a record from client_attributes
// @Description Delete a single record from client_attributes table in the keycloak database
// @Tags ClientAttributes
// @Accept  json
// @Produce  json
// @Param  argClientID path string true "client_id"// @Param  argName path string true "name"
// @Success 204 {object} model.ClientAttributes
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /clientattributes/{argClientID}/{argName} [delete]
// http DELETE "http://localhost:8080/clientattributes/hello world/hello world" X-Api-User:user123
func DeleteClientAttributes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argClientID, err := parseString(ps, "argClientID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argName, err := parseString(ps, "argName")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_attributes", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteClientAttributes(ctx, argClientID, argName)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
