package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configGroupRoleMappingRouter(router *httprouter.Router) {
	router.GET("/grouprolemapping", GetAllGroupRoleMapping)
	router.POST("/grouprolemapping", AddGroupRoleMapping)
	router.GET("/grouprolemapping/:argRoleID/:argGroupID", GetGroupRoleMapping)
	router.PUT("/grouprolemapping/:argRoleID/:argGroupID", UpdateGroupRoleMapping)
	router.DELETE("/grouprolemapping/:argRoleID/:argGroupID", DeleteGroupRoleMapping)
}

func configGinGroupRoleMappingRouter(router gin.IRoutes) {
	router.GET("/grouprolemapping", ConverHttprouterToGin(GetAllGroupRoleMapping))
	router.POST("/grouprolemapping", ConverHttprouterToGin(AddGroupRoleMapping))
	router.GET("/grouprolemapping/:argRoleID/:argGroupID", ConverHttprouterToGin(GetGroupRoleMapping))
	router.PUT("/grouprolemapping/:argRoleID/:argGroupID", ConverHttprouterToGin(UpdateGroupRoleMapping))
	router.DELETE("/grouprolemapping/:argRoleID/:argGroupID", ConverHttprouterToGin(DeleteGroupRoleMapping))
}

// GetAllGroupRoleMapping is a function to get a slice of record(s) from group_role_mapping table in the keycloak database
// @Summary Get list of GroupRoleMapping
// @Tags GroupRoleMapping
// @Description GetAllGroupRoleMapping is a handler to get a slice of record(s) from group_role_mapping table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.GroupRoleMapping}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /grouprolemapping [get]
// http "http://localhost:8080/grouprolemapping?page=0&pagesize=20" X-Api-User:user123
func GetAllGroupRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "group_role_mapping", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllGroupRoleMapping(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetGroupRoleMapping is a function to get a single record from the group_role_mapping table in the keycloak database
// @Summary Get record from table GroupRoleMapping by  argRoleID  argGroupID
// @Tags GroupRoleMapping
// @ID argRoleID
// @ID argGroupID
// @Description GetGroupRoleMapping is a function to get a single record from the group_role_mapping table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argRoleID path string true "role_id"
// @Param  argGroupID path string true "group_id"
// @Success 200 {object} model.GroupRoleMapping
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /grouprolemapping/{argRoleID}/{argGroupID} [get]
// http "http://localhost:8080/grouprolemapping/hello world/hello world" X-Api-User:user123
func GetGroupRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRoleID, err := parseString(ps, "argRoleID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argGroupID, err := parseString(ps, "argGroupID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "group_role_mapping", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetGroupRoleMapping(ctx, argRoleID, argGroupID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddGroupRoleMapping add to add a single record to group_role_mapping table in the keycloak database
// @Summary Add an record to group_role_mapping table
// @Description add to add a single record to group_role_mapping table in the keycloak database
// @Tags GroupRoleMapping
// @Accept  json
// @Produce  json
// @Param GroupRoleMapping body model.GroupRoleMapping true "Add GroupRoleMapping"
// @Success 200 {object} model.GroupRoleMapping
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /grouprolemapping [post]
// echo '{"role_id": "RwaJLUWmRDtyoDFOPIUpNQbLJ","group_id": "SoCixMFWfDyRKxmqXhxxGjxih"}' | http POST "http://localhost:8080/grouprolemapping" X-Api-User:user123
func AddGroupRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	grouprolemapping := &model.GroupRoleMapping{}

	if err := readJSON(r, grouprolemapping); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := grouprolemapping.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	grouprolemapping.Prepare()

	if err := grouprolemapping.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "group_role_mapping", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	grouprolemapping, _, err = dao.AddGroupRoleMapping(ctx, grouprolemapping)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, grouprolemapping)
}

// UpdateGroupRoleMapping Update a single record from group_role_mapping table in the keycloak database
// @Summary Update an record in table group_role_mapping
// @Description Update a single record from group_role_mapping table in the keycloak database
// @Tags GroupRoleMapping
// @Accept  json
// @Produce  json
// @Param  argRoleID path string true "role_id"// @Param  argGroupID path string true "group_id"
// @Param  GroupRoleMapping body model.GroupRoleMapping true "Update GroupRoleMapping record"
// @Success 200 {object} model.GroupRoleMapping
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /grouprolemapping/{argRoleID}/{argGroupID} [put]
// echo '{"role_id": "RwaJLUWmRDtyoDFOPIUpNQbLJ","group_id": "SoCixMFWfDyRKxmqXhxxGjxih"}' | http PUT "http://localhost:8080/grouprolemapping/hello world/hello world"  X-Api-User:user123
func UpdateGroupRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRoleID, err := parseString(ps, "argRoleID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argGroupID, err := parseString(ps, "argGroupID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	grouprolemapping := &model.GroupRoleMapping{}
	if err := readJSON(r, grouprolemapping); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := grouprolemapping.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	grouprolemapping.Prepare()

	if err := grouprolemapping.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "group_role_mapping", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	grouprolemapping, _, err = dao.UpdateGroupRoleMapping(ctx,
		argRoleID, argGroupID,
		grouprolemapping)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, grouprolemapping)
}

// DeleteGroupRoleMapping Delete a single record from group_role_mapping table in the keycloak database
// @Summary Delete a record from group_role_mapping
// @Description Delete a single record from group_role_mapping table in the keycloak database
// @Tags GroupRoleMapping
// @Accept  json
// @Produce  json
// @Param  argRoleID path string true "role_id"// @Param  argGroupID path string true "group_id"
// @Success 204 {object} model.GroupRoleMapping
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /grouprolemapping/{argRoleID}/{argGroupID} [delete]
// http DELETE "http://localhost:8080/grouprolemapping/hello world/hello world" X-Api-User:user123
func DeleteGroupRoleMapping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argRoleID, err := parseString(ps, "argRoleID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argGroupID, err := parseString(ps, "argGroupID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "group_role_mapping", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteGroupRoleMapping(ctx, argRoleID, argGroupID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
