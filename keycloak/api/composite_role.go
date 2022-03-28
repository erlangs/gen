package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configCompositeRoleRouter(router *httprouter.Router) {
	router.GET("/compositerole", GetAllCompositeRole)
	router.POST("/compositerole", AddCompositeRole)
	router.GET("/compositerole/:argComposite/:argChildRole", GetCompositeRole)
	router.PUT("/compositerole/:argComposite/:argChildRole", UpdateCompositeRole)
	router.DELETE("/compositerole/:argComposite/:argChildRole", DeleteCompositeRole)
}

func configGinCompositeRoleRouter(router gin.IRoutes) {
	router.GET("/compositerole", ConverHttprouterToGin(GetAllCompositeRole))
	router.POST("/compositerole", ConverHttprouterToGin(AddCompositeRole))
	router.GET("/compositerole/:argComposite/:argChildRole", ConverHttprouterToGin(GetCompositeRole))
	router.PUT("/compositerole/:argComposite/:argChildRole", ConverHttprouterToGin(UpdateCompositeRole))
	router.DELETE("/compositerole/:argComposite/:argChildRole", ConverHttprouterToGin(DeleteCompositeRole))
}

// GetAllCompositeRole is a function to get a slice of record(s) from composite_role table in the keycloak database
// @Summary Get list of CompositeRole
// @Tags CompositeRole
// @Description GetAllCompositeRole is a handler to get a slice of record(s) from composite_role table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.CompositeRole}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /compositerole [get]
// http "http://localhost:8080/compositerole?page=0&pagesize=20" X-Api-User:user123
func GetAllCompositeRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "composite_role", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllCompositeRole(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetCompositeRole is a function to get a single record from the composite_role table in the keycloak database
// @Summary Get record from table CompositeRole by  argComposite  argChildRole
// @Tags CompositeRole
// @ID argComposite
// @ID argChildRole
// @Description GetCompositeRole is a function to get a single record from the composite_role table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argComposite path string true "composite"
// @Param  argChildRole path string true "child_role"
// @Success 200 {object} model.CompositeRole
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /compositerole/{argComposite}/{argChildRole} [get]
// http "http://localhost:8080/compositerole/hello world/hello world" X-Api-User:user123
func GetCompositeRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argComposite, err := parseString(ps, "argComposite")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argChildRole, err := parseString(ps, "argChildRole")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "composite_role", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetCompositeRole(ctx, argComposite, argChildRole)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddCompositeRole add to add a single record to composite_role table in the keycloak database
// @Summary Add an record to composite_role table
// @Description add to add a single record to composite_role table in the keycloak database
// @Tags CompositeRole
// @Accept  json
// @Produce  json
// @Param CompositeRole body model.CompositeRole true "Add CompositeRole"
// @Success 200 {object} model.CompositeRole
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /compositerole [post]
// echo '{"composite": "jMYoumZBSntorKYImMBbyLhRm","child_role": "suFYXEHslhJkipLfQPCfsfwAk"}' | http POST "http://localhost:8080/compositerole" X-Api-User:user123
func AddCompositeRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	compositerole := &model.CompositeRole{}

	if err := readJSON(r, compositerole); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := compositerole.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	compositerole.Prepare()

	if err := compositerole.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "composite_role", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	compositerole, _, err = dao.AddCompositeRole(ctx, compositerole)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, compositerole)
}

// UpdateCompositeRole Update a single record from composite_role table in the keycloak database
// @Summary Update an record in table composite_role
// @Description Update a single record from composite_role table in the keycloak database
// @Tags CompositeRole
// @Accept  json
// @Produce  json
// @Param  argComposite path string true "composite"// @Param  argChildRole path string true "child_role"
// @Param  CompositeRole body model.CompositeRole true "Update CompositeRole record"
// @Success 200 {object} model.CompositeRole
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /compositerole/{argComposite}/{argChildRole} [put]
// echo '{"composite": "jMYoumZBSntorKYImMBbyLhRm","child_role": "suFYXEHslhJkipLfQPCfsfwAk"}' | http PUT "http://localhost:8080/compositerole/hello world/hello world"  X-Api-User:user123
func UpdateCompositeRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argComposite, err := parseString(ps, "argComposite")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argChildRole, err := parseString(ps, "argChildRole")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	compositerole := &model.CompositeRole{}
	if err := readJSON(r, compositerole); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := compositerole.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	compositerole.Prepare()

	if err := compositerole.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "composite_role", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	compositerole, _, err = dao.UpdateCompositeRole(ctx,
		argComposite, argChildRole,
		compositerole)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, compositerole)
}

// DeleteCompositeRole Delete a single record from composite_role table in the keycloak database
// @Summary Delete a record from composite_role
// @Description Delete a single record from composite_role table in the keycloak database
// @Tags CompositeRole
// @Accept  json
// @Produce  json
// @Param  argComposite path string true "composite"// @Param  argChildRole path string true "child_role"
// @Success 204 {object} model.CompositeRole
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /compositerole/{argComposite}/{argChildRole} [delete]
// http DELETE "http://localhost:8080/compositerole/hello world/hello world" X-Api-User:user123
func DeleteCompositeRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argComposite, err := parseString(ps, "argComposite")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argChildRole, err := parseString(ps, "argChildRole")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "composite_role", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteCompositeRole(ctx, argComposite, argChildRole)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
