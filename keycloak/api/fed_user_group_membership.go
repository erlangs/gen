package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configFedUserGroupMembershipRouter(router *httprouter.Router) {
	router.GET("/fedusergroupmembership", GetAllFedUserGroupMembership)
	router.POST("/fedusergroupmembership", AddFedUserGroupMembership)
	router.GET("/fedusergroupmembership/:argGroupID/:argUserID", GetFedUserGroupMembership)
	router.PUT("/fedusergroupmembership/:argGroupID/:argUserID", UpdateFedUserGroupMembership)
	router.DELETE("/fedusergroupmembership/:argGroupID/:argUserID", DeleteFedUserGroupMembership)
}

func configGinFedUserGroupMembershipRouter(router gin.IRoutes) {
	router.GET("/fedusergroupmembership", ConverHttprouterToGin(GetAllFedUserGroupMembership))
	router.POST("/fedusergroupmembership", ConverHttprouterToGin(AddFedUserGroupMembership))
	router.GET("/fedusergroupmembership/:argGroupID/:argUserID", ConverHttprouterToGin(GetFedUserGroupMembership))
	router.PUT("/fedusergroupmembership/:argGroupID/:argUserID", ConverHttprouterToGin(UpdateFedUserGroupMembership))
	router.DELETE("/fedusergroupmembership/:argGroupID/:argUserID", ConverHttprouterToGin(DeleteFedUserGroupMembership))
}

// GetAllFedUserGroupMembership is a function to get a slice of record(s) from fed_user_group_membership table in the keycloak database
// @Summary Get list of FedUserGroupMembership
// @Tags FedUserGroupMembership
// @Description GetAllFedUserGroupMembership is a handler to get a slice of record(s) from fed_user_group_membership table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.FedUserGroupMembership}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /fedusergroupmembership [get]
// http "http://localhost:8080/fedusergroupmembership?page=0&pagesize=20" X-Api-User:user123
func GetAllFedUserGroupMembership(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "fed_user_group_membership", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllFedUserGroupMembership(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetFedUserGroupMembership is a function to get a single record from the fed_user_group_membership table in the keycloak database
// @Summary Get record from table FedUserGroupMembership by  argGroupID  argUserID
// @Tags FedUserGroupMembership
// @ID argGroupID
// @ID argUserID
// @Description GetFedUserGroupMembership is a function to get a single record from the fed_user_group_membership table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argGroupID path string true "group_id"
// @Param  argUserID path string true "user_id"
// @Success 200 {object} model.FedUserGroupMembership
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /fedusergroupmembership/{argGroupID}/{argUserID} [get]
// http "http://localhost:8080/fedusergroupmembership/hello world/hello world" X-Api-User:user123
func GetFedUserGroupMembership(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "fed_user_group_membership", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetFedUserGroupMembership(ctx, argGroupID, argUserID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddFedUserGroupMembership add to add a single record to fed_user_group_membership table in the keycloak database
// @Summary Add an record to fed_user_group_membership table
// @Description add to add a single record to fed_user_group_membership table in the keycloak database
// @Tags FedUserGroupMembership
// @Accept  json
// @Produce  json
// @Param FedUserGroupMembership body model.FedUserGroupMembership true "Add FedUserGroupMembership"
// @Success 200 {object} model.FedUserGroupMembership
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /fedusergroupmembership [post]
// echo '{"group_id": "gwNhohsWlgIdSwSrYagWmlnSk","user_id": "GxkmXJKUVTeSmqMXNwiQIWiQT","realm_id": "YNHhlVsfRQwTAoGVJFStKTHom","storage_provider_id": "UdhtgURKCwyaKpuSOfdIXIfGR"}' | http POST "http://localhost:8080/fedusergroupmembership" X-Api-User:user123
func AddFedUserGroupMembership(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	fedusergroupmembership := &model.FedUserGroupMembership{}

	if err := readJSON(r, fedusergroupmembership); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := fedusergroupmembership.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	fedusergroupmembership.Prepare()

	if err := fedusergroupmembership.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_group_membership", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	fedusergroupmembership, _, err = dao.AddFedUserGroupMembership(ctx, fedusergroupmembership)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, fedusergroupmembership)
}

// UpdateFedUserGroupMembership Update a single record from fed_user_group_membership table in the keycloak database
// @Summary Update an record in table fed_user_group_membership
// @Description Update a single record from fed_user_group_membership table in the keycloak database
// @Tags FedUserGroupMembership
// @Accept  json
// @Produce  json
// @Param  argGroupID path string true "group_id"// @Param  argUserID path string true "user_id"
// @Param  FedUserGroupMembership body model.FedUserGroupMembership true "Update FedUserGroupMembership record"
// @Success 200 {object} model.FedUserGroupMembership
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /fedusergroupmembership/{argGroupID}/{argUserID} [put]
// echo '{"group_id": "gwNhohsWlgIdSwSrYagWmlnSk","user_id": "GxkmXJKUVTeSmqMXNwiQIWiQT","realm_id": "YNHhlVsfRQwTAoGVJFStKTHom","storage_provider_id": "UdhtgURKCwyaKpuSOfdIXIfGR"}' | http PUT "http://localhost:8080/fedusergroupmembership/hello world/hello world"  X-Api-User:user123
func UpdateFedUserGroupMembership(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	fedusergroupmembership := &model.FedUserGroupMembership{}
	if err := readJSON(r, fedusergroupmembership); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := fedusergroupmembership.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	fedusergroupmembership.Prepare()

	if err := fedusergroupmembership.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "fed_user_group_membership", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	fedusergroupmembership, _, err = dao.UpdateFedUserGroupMembership(ctx,
		argGroupID, argUserID,
		fedusergroupmembership)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, fedusergroupmembership)
}

// DeleteFedUserGroupMembership Delete a single record from fed_user_group_membership table in the keycloak database
// @Summary Delete a record from fed_user_group_membership
// @Description Delete a single record from fed_user_group_membership table in the keycloak database
// @Tags FedUserGroupMembership
// @Accept  json
// @Produce  json
// @Param  argGroupID path string true "group_id"// @Param  argUserID path string true "user_id"
// @Success 204 {object} model.FedUserGroupMembership
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /fedusergroupmembership/{argGroupID}/{argUserID} [delete]
// http DELETE "http://localhost:8080/fedusergroupmembership/hello world/hello world" X-Api-User:user123
func DeleteFedUserGroupMembership(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "fed_user_group_membership", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteFedUserGroupMembership(ctx, argGroupID, argUserID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
