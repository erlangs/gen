package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configRoleAttributeRouter(router *httprouter.Router) {
	router.GET("/roleattribute", GetAllRoleAttribute)
	router.POST("/roleattribute", AddRoleAttribute)
	router.GET("/roleattribute/:argID", GetRoleAttribute)
	router.PUT("/roleattribute/:argID", UpdateRoleAttribute)
	router.DELETE("/roleattribute/:argID", DeleteRoleAttribute)
}

func configGinRoleAttributeRouter(router gin.IRoutes) {
	router.GET("/roleattribute", ConverHttprouterToGin(GetAllRoleAttribute))
	router.POST("/roleattribute", ConverHttprouterToGin(AddRoleAttribute))
	router.GET("/roleattribute/:argID", ConverHttprouterToGin(GetRoleAttribute))
	router.PUT("/roleattribute/:argID", ConverHttprouterToGin(UpdateRoleAttribute))
	router.DELETE("/roleattribute/:argID", ConverHttprouterToGin(DeleteRoleAttribute))
}

// GetAllRoleAttribute is a function to get a slice of record(s) from role_attribute table in the keycloak database
// @Summary Get list of RoleAttribute
// @Tags RoleAttribute
// @Description GetAllRoleAttribute is a handler to get a slice of record(s) from role_attribute table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.RoleAttribute}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /roleattribute [get]
// http "http://localhost:8080/roleattribute?page=0&pagesize=20" X-Api-User:user123
func GetAllRoleAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "role_attribute", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllRoleAttribute(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetRoleAttribute is a function to get a single record from the role_attribute table in the keycloak database
// @Summary Get record from table RoleAttribute by  argID
// @Tags RoleAttribute
// @ID argID
// @Description GetRoleAttribute is a function to get a single record from the role_attribute table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.RoleAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /roleattribute/{argID} [get]
// http "http://localhost:8080/roleattribute/hello world" X-Api-User:user123
func GetRoleAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "role_attribute", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetRoleAttribute(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddRoleAttribute add to add a single record to role_attribute table in the keycloak database
// @Summary Add an record to role_attribute table
// @Description add to add a single record to role_attribute table in the keycloak database
// @Tags RoleAttribute
// @Accept  json
// @Produce  json
// @Param RoleAttribute body model.RoleAttribute true "Add RoleAttribute"
// @Success 200 {object} model.RoleAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /roleattribute [post]
// echo '{"id": "WXoDaUnFLmEtQwEEMaDPvctYq","role_id": "vSOfpbwmsxvNFdvSXNRiyaZmO","name": "kPGDxhftBFldxTZexaoYtXdoo","value": "forgBtIfQnMsGaOvdtYqpYDqG"}' | http POST "http://localhost:8080/roleattribute" X-Api-User:user123
func AddRoleAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	roleattribute := &model.RoleAttribute{}

	if err := readJSON(r, roleattribute); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := roleattribute.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	roleattribute.Prepare()

	if err := roleattribute.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "role_attribute", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	roleattribute, _, err = dao.AddRoleAttribute(ctx, roleattribute)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, roleattribute)
}

// UpdateRoleAttribute Update a single record from role_attribute table in the keycloak database
// @Summary Update an record in table role_attribute
// @Description Update a single record from role_attribute table in the keycloak database
// @Tags RoleAttribute
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  RoleAttribute body model.RoleAttribute true "Update RoleAttribute record"
// @Success 200 {object} model.RoleAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /roleattribute/{argID} [put]
// echo '{"id": "WXoDaUnFLmEtQwEEMaDPvctYq","role_id": "vSOfpbwmsxvNFdvSXNRiyaZmO","name": "kPGDxhftBFldxTZexaoYtXdoo","value": "forgBtIfQnMsGaOvdtYqpYDqG"}' | http PUT "http://localhost:8080/roleattribute/hello world"  X-Api-User:user123
func UpdateRoleAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	roleattribute := &model.RoleAttribute{}
	if err := readJSON(r, roleattribute); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := roleattribute.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	roleattribute.Prepare()

	if err := roleattribute.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "role_attribute", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	roleattribute, _, err = dao.UpdateRoleAttribute(ctx,
		argID,
		roleattribute)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, roleattribute)
}

// DeleteRoleAttribute Delete a single record from role_attribute table in the keycloak database
// @Summary Delete a record from role_attribute
// @Description Delete a single record from role_attribute table in the keycloak database
// @Tags RoleAttribute
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.RoleAttribute
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /roleattribute/{argID} [delete]
// http DELETE "http://localhost:8080/roleattribute/hello world" X-Api-User:user123
func DeleteRoleAttribute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "role_attribute", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteRoleAttribute(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
