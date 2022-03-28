package api

import (
	"net/http"

	"example.com/rest/example/dao"
	"example.com/rest/example/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configPersonalAccessTokensRouter(router *httprouter.Router) {
	router.GET("/personalaccesstokens", GetAllPersonalAccessTokens)
	router.POST("/personalaccesstokens", AddPersonalAccessTokens)
	router.GET("/personalaccesstokens/:argID", GetPersonalAccessTokens)
	router.PUT("/personalaccesstokens/:argID", UpdatePersonalAccessTokens)
	router.DELETE("/personalaccesstokens/:argID", DeletePersonalAccessTokens)
}

func configGinPersonalAccessTokensRouter(router gin.IRoutes) {
	router.GET("/personalaccesstokens", ConverHttprouterToGin(GetAllPersonalAccessTokens))
	router.POST("/personalaccesstokens", ConverHttprouterToGin(AddPersonalAccessTokens))
	router.GET("/personalaccesstokens/:argID", ConverHttprouterToGin(GetPersonalAccessTokens))
	router.PUT("/personalaccesstokens/:argID", ConverHttprouterToGin(UpdatePersonalAccessTokens))
	router.DELETE("/personalaccesstokens/:argID", ConverHttprouterToGin(DeletePersonalAccessTokens))
}

// GetAllPersonalAccessTokens is a function to get a slice of record(s) from personal_access_tokens table in the test1 database
// @Summary Get list of PersonalAccessTokens
// @Tags PersonalAccessTokens
// @Description GetAllPersonalAccessTokens is a handler to get a slice of record(s) from personal_access_tokens table in the test1 database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.PersonalAccessTokens}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /personalaccesstokens [get]
// http "http://localhost:8080/personalaccesstokens?page=0&pagesize=20" X-Api-User:user123
func GetAllPersonalAccessTokens(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "personal_access_tokens", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllPersonalAccessTokens(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetPersonalAccessTokens is a function to get a single record from the personal_access_tokens table in the test1 database
// @Summary Get record from table PersonalAccessTokens by  argID
// @Tags PersonalAccessTokens
// @ID argID
// @Description GetPersonalAccessTokens is a function to get a single record from the personal_access_tokens table in the test1 database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.PersonalAccessTokens
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /personalaccesstokens/{argID} [get]
// http "http://localhost:8080/personalaccesstokens/1" X-Api-User:user123
func GetPersonalAccessTokens(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseUint64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "personal_access_tokens", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetPersonalAccessTokens(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddPersonalAccessTokens add to add a single record to personal_access_tokens table in the test1 database
// @Summary Add an record to personal_access_tokens table
// @Description add to add a single record to personal_access_tokens table in the test1 database
// @Tags PersonalAccessTokens
// @Accept  json
// @Produce  json
// @Param PersonalAccessTokens body model.PersonalAccessTokens true "Add PersonalAccessTokens"
// @Success 200 {object} model.PersonalAccessTokens
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /personalaccesstokens [post]
// echo '{"id": 20,"tokenable_type": "gJTXZbympgguDdRSbKRrZUHEe","tokenable_id": 2,"name": "OtrHhjbxlbkpxldCThEnZrFie","token": "IGobwpJTgmEjdjKqVwRGPTLdQ","abilities": "wtBxQYSMMSOwhtrRtbmwhScdI","last_used_at": "2273-07-03T09:53:20.870639989+08:00","created_at": "2279-07-03T20:55:45.958757079+08:00","updated_at": "2261-02-22T12:33:14.767856874+08:00"}' | http POST "http://localhost:8080/personalaccesstokens" X-Api-User:user123
func AddPersonalAccessTokens(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	personalaccesstokens := &model.PersonalAccessTokens{}

	if err := readJSON(r, personalaccesstokens); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := personalaccesstokens.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	personalaccesstokens.Prepare()

	if err := personalaccesstokens.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "personal_access_tokens", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	personalaccesstokens, _, err = dao.AddPersonalAccessTokens(ctx, personalaccesstokens)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, personalaccesstokens)
}

// UpdatePersonalAccessTokens Update a single record from personal_access_tokens table in the test1 database
// @Summary Update an record in table personal_access_tokens
// @Description Update a single record from personal_access_tokens table in the test1 database
// @Tags PersonalAccessTokens
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  PersonalAccessTokens body model.PersonalAccessTokens true "Update PersonalAccessTokens record"
// @Success 200 {object} model.PersonalAccessTokens
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /personalaccesstokens/{argID} [put]
// echo '{"id": 20,"tokenable_type": "gJTXZbympgguDdRSbKRrZUHEe","tokenable_id": 2,"name": "OtrHhjbxlbkpxldCThEnZrFie","token": "IGobwpJTgmEjdjKqVwRGPTLdQ","abilities": "wtBxQYSMMSOwhtrRtbmwhScdI","last_used_at": "2273-07-03T09:53:20.870639989+08:00","created_at": "2279-07-03T20:55:45.958757079+08:00","updated_at": "2261-02-22T12:33:14.767856874+08:00"}' | http PUT "http://localhost:8080/personalaccesstokens/1"  X-Api-User:user123
func UpdatePersonalAccessTokens(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseUint64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	personalaccesstokens := &model.PersonalAccessTokens{}
	if err := readJSON(r, personalaccesstokens); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := personalaccesstokens.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	personalaccesstokens.Prepare()

	if err := personalaccesstokens.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "personal_access_tokens", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	personalaccesstokens, _, err = dao.UpdatePersonalAccessTokens(ctx,
		argID,
		personalaccesstokens)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, personalaccesstokens)
}

// DeletePersonalAccessTokens Delete a single record from personal_access_tokens table in the test1 database
// @Summary Delete a record from personal_access_tokens
// @Description Delete a single record from personal_access_tokens table in the test1 database
// @Tags PersonalAccessTokens
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.PersonalAccessTokens
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /personalaccesstokens/{argID} [delete]
// http DELETE "http://localhost:8080/personalaccesstokens/1" X-Api-User:user123
func DeletePersonalAccessTokens(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseUint64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "personal_access_tokens", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeletePersonalAccessTokens(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
