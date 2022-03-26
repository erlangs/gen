package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configUserGroupMembershipRouter(router *httprouter.Router) {
	router.GET("/usergroupmembership", GetAllUserGroupMembership)
	router.POST("/usergroupmembership", AddUserGroupMembership)
	router.GET("/usergroupmembership/:argGroupID/:argUserID", GetUserGroupMembership)
	router.PUT("/usergroupmembership/:argGroupID/:argUserID", UpdateUserGroupMembership)
	router.DELETE("/usergroupmembership/:argGroupID/:argUserID", DeleteUserGroupMembership)
}

func configGinUserGroupMembershipRouter(router gin.IRoutes) {
	router.GET("/usergroupmembership", ConverHttprouterToGin(GetAllUserGroupMembership))
	router.POST("/usergroupmembership", ConverHttprouterToGin(AddUserGroupMembership))
	router.GET("/usergroupmembership/:argGroupID/:argUserID", ConverHttprouterToGin(GetUserGroupMembership))
	router.PUT("/usergroupmembership/:argGroupID/:argUserID", ConverHttprouterToGin(UpdateUserGroupMembership))
	router.DELETE("/usergroupmembership/:argGroupID/:argUserID", ConverHttprouterToGin(DeleteUserGroupMembership))
}

// GetAllUserGroupMembership is a function to get a slice of record(s) from user_group_membership table in the keycloak database
// @Summary Get list of UserGroupMembership
// @Tags UserGroupMembership
// @Description GetAllUserGroupMembership is a handler to get a slice of record(s) from user_group_membership table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.UserGroupMembership}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /usergroupmembership [get]
// http "http://localhost:8080/usergroupmembership?page=0&pagesize=20" X-Api-User:user123
func GetAllUserGroupMembership(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "user_group_membership", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllUserGroupMembership(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetUserGroupMembership is a function to get a single record from the user_group_membership table in the keycloak database
// @Summary Get record from table UserGroupMembership by  argGroupID  argUserID
// @Tags UserGroupMembership
// @ID argGroupID
// @ID argUserID
// @Description GetUserGroupMembership is a function to get a single record from the user_group_membership table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argGroupID path string true "group_id"
// @Param  argUserID path string true "user_id"
// @Success 200 {object} model.UserGroupMembership
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /usergroupmembership/{argGroupID}/{argUserID} [get]
// http "http://localhost:8080/usergroupmembership/hello world/hello world" X-Api-User:user123
func GetUserGroupMembership(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argGroupID, err := parseString(ps, "argGroupID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argUserID, err := parseString(ps, "argUserID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_group_membership", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetUserGroupMembership(ctx, argGroupID, argUserID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddUserGroupMembership add to add a single record to user_group_membership table in the keycloak database
// @Summary Add an record to user_group_membership table
// @Description add to add a single record to user_group_membership table in the keycloak database
// @Tags UserGroupMembership
// @Accept  json
// @Produce  json
// @Param UserGroupMembership body model.UserGroupMembership true "Add UserGroupMembership"
// @Success 200 {object} model.UserGroupMembership
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /usergroupmembership [post]
// echo '{"group_id": "pGvigLhoWrdOTgmKribxPUWdT","user_id": "pTSINLeQyOKFhAidPEeISXNHf"}' | http POST "http://localhost:8080/usergroupmembership" X-Api-User:user123
func AddUserGroupMembership(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	usergroupmembership := &model.UserGroupMembership{}

	if err := readJSON(r, usergroupmembership); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := usergroupmembership.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	usergroupmembership.Prepare()

	if err := usergroupmembership.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_group_membership", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	usergroupmembership, _, err = dao.AddUserGroupMembership(ctx, usergroupmembership)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, usergroupmembership)
}

// UpdateUserGroupMembership Update a single record from user_group_membership table in the keycloak database
// @Summary Update an record in table user_group_membership
// @Description Update a single record from user_group_membership table in the keycloak database
// @Tags UserGroupMembership
// @Accept  json
// @Produce  json
// @Param  argGroupID path string true "group_id"// @Param  argUserID path string true "user_id"
// @Param  UserGroupMembership body model.UserGroupMembership true "Update UserGroupMembership record"
// @Success 200 {object} model.UserGroupMembership
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /usergroupmembership/{argGroupID}/{argUserID} [put]
// echo '{"group_id": "pGvigLhoWrdOTgmKribxPUWdT","user_id": "pTSINLeQyOKFhAidPEeISXNHf"}' | http PUT "http://localhost:8080/usergroupmembership/hello world/hello world"  X-Api-User:user123
func UpdateUserGroupMembership(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argGroupID, err := parseString(ps, "argGroupID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argUserID, err := parseString(ps, "argUserID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	usergroupmembership := &model.UserGroupMembership{}
	if err := readJSON(r, usergroupmembership); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := usergroupmembership.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	usergroupmembership.Prepare()

	if err := usergroupmembership.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "user_group_membership", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	usergroupmembership, _, err = dao.UpdateUserGroupMembership(ctx,
		argGroupID, argUserID,
		usergroupmembership)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, usergroupmembership)
}

// DeleteUserGroupMembership Delete a single record from user_group_membership table in the keycloak database
// @Summary Delete a record from user_group_membership
// @Description Delete a single record from user_group_membership table in the keycloak database
// @Tags UserGroupMembership
// @Accept  json
// @Produce  json
// @Param  argGroupID path string true "group_id"// @Param  argUserID path string true "user_id"
// @Success 204 {object} model.UserGroupMembership
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /usergroupmembership/{argGroupID}/{argUserID} [delete]
// http DELETE "http://localhost:8080/usergroupmembership/hello world/hello world" X-Api-User:user123
func DeleteUserGroupMembership(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argGroupID, err := parseString(ps, "argGroupID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argUserID, err := parseString(ps, "argUserID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "user_group_membership", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteUserGroupMembership(ctx, argGroupID, argUserID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
