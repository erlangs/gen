package api

import (
	"net/http"

	"example.com/rest/example/dao"
	"example.com/rest/example/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configUsers_Router(router *httprouter.Router) {
	router.GET("/users_", GetAllUsers_)
	router.POST("/users_", AddUsers_)
	router.GET("/users_/:argID", GetUsers_)
	router.PUT("/users_/:argID", UpdateUsers_)
	router.DELETE("/users_/:argID", DeleteUsers_)
}

func configGinUsers_Router(router gin.IRoutes) {
	router.GET("/users_", ConverHttprouterToGin(GetAllUsers_))
	router.POST("/users_", ConverHttprouterToGin(AddUsers_))
	router.GET("/users_/:argID", ConverHttprouterToGin(GetUsers_))
	router.PUT("/users_/:argID", ConverHttprouterToGin(UpdateUsers_))
	router.DELETE("/users_/:argID", ConverHttprouterToGin(DeleteUsers_))
}

// GetAllUsers_ is a function to get a slice of record(s) from users table in the test1 database
// @Summary Get list of Users_
// @Tags Users_
// @Description GetAllUsers_ is a handler to get a slice of record(s) from users table in the test1 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Users_}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /users_ [get]
// http "http://localhost:8080/users_?page=0&pagesize=20" X-Api-User:user123
func GetAllUsers_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "users", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllUsers_(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetUsers_ is a function to get a single record from the users table in the test1 database
// @Summary Get record from table Users_ by  argID
// @Tags Users_
// @ID argID
// @Description GetUsers_ is a function to get a single record from the users table in the test1 database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Users_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /users_/{argID} [get]
// http "http://localhost:8080/users_/1" X-Api-User:user123
func GetUsers_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseUint64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "users", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetUsers_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddUsers_ add to add a single record to users table in the test1 database
// @Summary Add an record to users table
// @Description add to add a single record to users table in the test1 database
// @Tags Users_
// @Accept  json
// @Produce  json
// @Param Users_ body model.Users_ true "Add Users_"
// @Success 200 {object} model.Users_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /users_ [post]
// echo '{"id": 85,"name": "YxmeqFOGChasUxlbKxTBEZdpe","email": "UPyrEaUpwmqtypxpscriCeKgT","email_verified_at": "2041-08-17T18:12:15.208363605+08:00","password": "nZHGxbXrQKCNxUHCyUvmDWXon","remember_token": "TXxRWcQSSQggQiebFKbXLfUVX","created_at": "2277-02-10T03:06:55.986321258+08:00","updated_at": "2282-07-19T23:33:11.225075029+08:00"}' | http POST "http://localhost:8080/users_" X-Api-User:user123
func AddUsers_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	users_ := &model.Users_{}

	if err := readJSON(r, users_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := users_.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	users_.Prepare()

	if err := users_.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "users", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	users_, _, err = dao.AddUsers_(ctx, users_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, users_)
}

// UpdateUsers_ Update a single record from users table in the test1 database
// @Summary Update an record in table users
// @Description Update a single record from users table in the test1 database
// @Tags Users_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Users_ body model.Users_ true "Update Users_ record"
// @Success 200 {object} model.Users_
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /users_/{argID} [put]
// echo '{"id": 85,"name": "YxmeqFOGChasUxlbKxTBEZdpe","email": "UPyrEaUpwmqtypxpscriCeKgT","email_verified_at": "2041-08-17T18:12:15.208363605+08:00","password": "nZHGxbXrQKCNxUHCyUvmDWXon","remember_token": "TXxRWcQSSQggQiebFKbXLfUVX","created_at": "2277-02-10T03:06:55.986321258+08:00","updated_at": "2282-07-19T23:33:11.225075029+08:00"}' | http PUT "http://localhost:8080/users_/1"  X-Api-User:user123
func UpdateUsers_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseUint64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	users_ := &model.Users_{}
	if err := readJSON(r, users_); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := users_.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	users_.Prepare()

	if err := users_.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "users", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	users_, _, err = dao.UpdateUsers_(ctx,
		argID,
		users_)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, users_)
}

// DeleteUsers_ Delete a single record from users table in the test1 database
// @Summary Delete a record from users
// @Description Delete a single record from users table in the test1 database
// @Tags Users_
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Users_
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /users_/{argID} [delete]
// http DELETE "http://localhost:8080/users_/1" X-Api-User:user123
func DeleteUsers_(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseUint64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "users", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteUsers_(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
