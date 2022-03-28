package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configClientScopeRoleMappingRouter(router *httprouter.Router) {
	router.GET("/clientscoperolemapping", GetAllClientScopeRoleMapping)
	router.POST("/clientscoperolemapping", AddClientScopeRoleMapping)
	router.GET("/clientscoperolemapping/:argScopeID/:argRoleID", GetClientScopeRoleMapping)
	router.PUT("/clientscoperolemapping/:argScopeID/:argRoleID", UpdateClientScopeRoleMapping)
	router.DELETE("/clientscoperolemapping/:argScopeID/:argRoleID", DeleteClientScopeRoleMapping)
}

func configGinClientScopeRoleMappingRouter(router gin.IRoutes) {
	router.GET("/clientscoperolemapping", ConverHttprouterToGin(GetAllClientScopeRoleMapping))
	router.POST("/clientscoperolemapping", ConverHttprouterToGin(AddClientScopeRoleMapping))
	router.GET("/clientscoperolemapping/:argScopeID/:argRoleID", ConverHttprouterToGin(GetClientScopeRoleMapping))
	router.PUT("/clientscoperolemapping/:argScopeID/:argRoleID", ConverHttprouterToGin(UpdateClientScopeRoleMapping))
	router.DELETE("/clientscoperolemapping/:argScopeID/:argRoleID", ConverHttprouterToGin(DeleteClientScopeRoleMapping))
}

// GetAllClientScopeRoleMapping is a function to get a slice of record(s) from client_scope_role_mapping table in the keycloak database
// @Summary Get list of ClientScopeRoleMapping
// @Tags ClientScopeRoleMapping
// @Description GetAllClientScopeRoleMapping is a handler to get a slice of record(s) from client_scope_role_mapping table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ClientScopeRoleMapping}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientscoperolemapping [get]
// http "http://localhost:8080/clientscoperolemapping?page=0&pagesize=20" X-Api-User:user123
func GetAllClientScopeRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "client_scope_role_mapping", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllClientScopeRoleMapping(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetClientScopeRoleMapping is a function to get a single record from the client_scope_role_mapping table in the keycloak database
// @Summary Get record from table ClientScopeRoleMapping by  argScopeID  argRoleID
// @Tags ClientScopeRoleMapping
// @ID argScopeID
// @ID argRoleID
// @Description GetClientScopeRoleMapping is a function to get a single record from the client_scope_role_mapping table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argScopeID path string true "scope_id"
// @Param  argRoleID path string true "role_id"
// @Success 200 {object} model.ClientScopeRoleMapping
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /clientscoperolemapping/{argScopeID}/{argRoleID} [get]
// http "http://localhost:8080/clientscoperolemapping/hello world/hello world" X-Api-User:user123
func GetClientScopeRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argRoleID, err := parseString(ps, "argRoleID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_scope_role_mapping", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetClientScopeRoleMapping(ctx, argScopeID, argRoleID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddClientScopeRoleMapping add to add a single record to client_scope_role_mapping table in the keycloak database
// @Summary Add an record to client_scope_role_mapping table
// @Description add to add a single record to client_scope_role_mapping table in the keycloak database
// @Tags ClientScopeRoleMapping
// @Accept  json
// @Produce  json
// @Param ClientScopeRoleMapping body model.ClientScopeRoleMapping true "Add ClientScopeRoleMapping"
// @Success 200 {object} model.ClientScopeRoleMapping
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientscoperolemapping [post]
// echo '{"scope_id": "pCQELmcAMtBKQSIHUxMqmMaxo","role_id": "yDaoleUCHPNUrSnbYMHjDyNmL"}' | http POST "http://localhost:8080/clientscoperolemapping" X-Api-User:user123
func AddClientScopeRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	clientscoperolemapping := &model.ClientScopeRoleMapping{}

	if err := readJSON(r, clientscoperolemapping); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientscoperolemapping.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientscoperolemapping.Prepare()

	if err := clientscoperolemapping.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_scope_role_mapping", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	clientscoperolemapping, _, err = dao.AddClientScopeRoleMapping(ctx, clientscoperolemapping)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientscoperolemapping)
}

// UpdateClientScopeRoleMapping Update a single record from client_scope_role_mapping table in the keycloak database
// @Summary Update an record in table client_scope_role_mapping
// @Description Update a single record from client_scope_role_mapping table in the keycloak database
// @Tags ClientScopeRoleMapping
// @Accept  json
// @Produce  json
// @Param  argScopeID path string true "scope_id"// @Param  argRoleID path string true "role_id"
// @Param  ClientScopeRoleMapping body model.ClientScopeRoleMapping true "Update ClientScopeRoleMapping record"
// @Success 200 {object} model.ClientScopeRoleMapping
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /clientscoperolemapping/{argScopeID}/{argRoleID} [put]
// echo '{"scope_id": "pCQELmcAMtBKQSIHUxMqmMaxo","role_id": "yDaoleUCHPNUrSnbYMHjDyNmL"}' | http PUT "http://localhost:8080/clientscoperolemapping/hello world/hello world"  X-Api-User:user123
func UpdateClientScopeRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argRoleID, err := parseString(ps, "argRoleID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientscoperolemapping := &model.ClientScopeRoleMapping{}
	if err := readJSON(r, clientscoperolemapping); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := clientscoperolemapping.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	clientscoperolemapping.Prepare()

	if err := clientscoperolemapping.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "client_scope_role_mapping", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	clientscoperolemapping, _, err = dao.UpdateClientScopeRoleMapping(ctx,
		argScopeID, argRoleID,
		clientscoperolemapping)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, clientscoperolemapping)
}

// DeleteClientScopeRoleMapping Delete a single record from client_scope_role_mapping table in the keycloak database
// @Summary Delete a record from client_scope_role_mapping
// @Description Delete a single record from client_scope_role_mapping table in the keycloak database
// @Tags ClientScopeRoleMapping
// @Accept  json
// @Produce  json
// @Param  argScopeID path string true "scope_id"// @Param  argRoleID path string true "role_id"
// @Success 204 {object} model.ClientScopeRoleMapping
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /clientscoperolemapping/{argScopeID}/{argRoleID} [delete]
// http DELETE "http://localhost:8080/clientscoperolemapping/hello world/hello world" X-Api-User:user123
func DeleteClientScopeRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argScopeID, err := parseString(ps, "argScopeID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argRoleID, err := parseString(ps, "argRoleID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "client_scope_role_mapping", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteClientScopeRoleMapping(ctx, argScopeID, argRoleID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
