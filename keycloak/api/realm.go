package api

import (
	"net/http"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"

	"github.com/julienschmidt/httprouter"
)

func configRealmRouter(router *httprouter.Router) {
	router.GET("/realm", GetAllRealm)
	router.POST("/realm", AddRealm)
	router.GET("/realm/:argID", GetRealm)
	router.PUT("/realm/:argID", UpdateRealm)
	router.DELETE("/realm/:argID", DeleteRealm)
}

func configGinRealmRouter(router gin.IRoutes) {
	router.GET("/realm", ConverHttprouterToGin(GetAllRealm))
	router.POST("/realm", ConverHttprouterToGin(AddRealm))
	router.GET("/realm/:argID", ConverHttprouterToGin(GetRealm))
	router.PUT("/realm/:argID", ConverHttprouterToGin(UpdateRealm))
	router.DELETE("/realm/:argID", ConverHttprouterToGin(DeleteRealm))
}

// GetAllRealm is a function to get a slice of record(s) from realm table in the keycloak database
// @Summary Get list of Realm
// @Tags Realm
// @Description GetAllRealm is a handler to get a slice of record(s) from realm table in the keycloak database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Realm}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realm [get]
// http "http://localhost:8080/realm?page=0&pagesize=20" X-Api-User:user123
func GetAllRealm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "realm", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllRealm(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetRealm is a function to get a single record from the realm table in the keycloak database
// @Summary Get record from table Realm by  argID
// @Tags Realm
// @ID argID
// @Description GetRealm is a function to get a single record from the realm table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 200 {object} model.Realm
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /realm/{argID} [get]
// http "http://localhost:8080/realm/hello world" X-Api-User:user123
func GetRealm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "realm", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetRealm(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddRealm add to add a single record to realm table in the keycloak database
// @Summary Add an record to realm table
// @Description add to add a single record to realm table in the keycloak database
// @Tags Realm
// @Accept  json
// @Produce  json
// @Param Realm body model.Realm true "Add Realm"
// @Success 200 {object} model.Realm
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realm [post]
// echo '{"id": "qlYOcmbcmSURJxCLVmDlSMvwA","access_code_lifespan": 69,"user_action_lifespan": 17,"access_token_lifespan": 50,"account_theme": "JuOqefMnyvqPBiuQiFPMNiAJW","admin_theme": "pYgdrQZqqxYexwBPyaISedvao","email_theme": "hnuwOVdWGCtwwymFwSUsCcTNI","enabled": true,"events_enabled": false,"events_expiration": 34,"login_theme": "sJSxJFvXDInPqpWcbMUdRYUar","name": "gHBDSBtZrvHKxHsidFIasjcpA","not_before": 26,"password_policy": "TswUpYfBiBifxUHEKkEgFYUlm","registration_allowed": false,"remember_me": false,"reset_password_allowed": false,"social": true,"ssl_required": "aHeRTfmtqgdKKLPEHDrYMZjoD","sso_idle_timeout": 70,"sso_max_lifespan": 40,"update_profile_on_soc_login": true,"verify_email": false,"master_admin_client": "cFStJCWZQlyDdJdmUmAJuJYIr","login_lifespan": 30,"internationalization_enabled": true,"default_locale": "foZtFGTahZfsCrNFcQisrkfnN","reg_email_as_username": false,"admin_events_enabled": true,"admin_events_details_enabled": true,"edit_username_allowed": true,"otp_policy_counter": 17,"otp_policy_window": 81,"otp_policy_period": 16,"otp_policy_digits": 97,"otp_policy_alg": "ZDjIAsJUBCIxeIyWFMGZfkBHF","otp_policy_type": "KoceZgpiBwbHnchlQrMKgcTpO","browser_flow": "nEHMdOoabBXnYOmMhMViHmHjO","registration_flow": "OLDpdsOtgQfqexUkIlkSjWDng","direct_grant_flow": "YQXNOMiFRbPadjDPXdedYhZge","reset_credentials_flow": "QCRdFcklYuoeBsrBqdnRYZJIc","client_auth_flow": "odvQbPvPOdZysAscRMIiXFmKr","offline_session_idle_timeout": 70,"revoke_refresh_token": false,"access_token_life_implicit": 95,"login_with_email_allowed": false,"duplicate_emails_allowed": true,"docker_auth_flow": "avBMHYgSFSlmBQebmwhwnmnDB","refresh_token_max_reuse": 23,"allow_user_managed_access": false,"sso_max_lifespan_remember_me": 8,"sso_idle_timeout_remember_me": 71,"default_role": "JdLqjRVmkbefNFAUWEumEoCXE"}' | http POST "http://localhost:8080/realm" X-Api-User:user123
func AddRealm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	realm := &model.Realm{}

	if err := readJSON(r, realm); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := realm.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	realm.Prepare()

	if err := realm.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "realm", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	realm, _, err = dao.AddRealm(ctx, realm)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, realm)
}

// UpdateRealm Update a single record from realm table in the keycloak database
// @Summary Update an record in table realm
// @Description Update a single record from realm table in the keycloak database
// @Tags Realm
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Param  Realm body model.Realm true "Update Realm record"
// @Success 200 {object} model.Realm
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /realm/{argID} [put]
// echo '{"id": "qlYOcmbcmSURJxCLVmDlSMvwA","access_code_lifespan": 69,"user_action_lifespan": 17,"access_token_lifespan": 50,"account_theme": "JuOqefMnyvqPBiuQiFPMNiAJW","admin_theme": "pYgdrQZqqxYexwBPyaISedvao","email_theme": "hnuwOVdWGCtwwymFwSUsCcTNI","enabled": true,"events_enabled": false,"events_expiration": 34,"login_theme": "sJSxJFvXDInPqpWcbMUdRYUar","name": "gHBDSBtZrvHKxHsidFIasjcpA","not_before": 26,"password_policy": "TswUpYfBiBifxUHEKkEgFYUlm","registration_allowed": false,"remember_me": false,"reset_password_allowed": false,"social": true,"ssl_required": "aHeRTfmtqgdKKLPEHDrYMZjoD","sso_idle_timeout": 70,"sso_max_lifespan": 40,"update_profile_on_soc_login": true,"verify_email": false,"master_admin_client": "cFStJCWZQlyDdJdmUmAJuJYIr","login_lifespan": 30,"internationalization_enabled": true,"default_locale": "foZtFGTahZfsCrNFcQisrkfnN","reg_email_as_username": false,"admin_events_enabled": true,"admin_events_details_enabled": true,"edit_username_allowed": true,"otp_policy_counter": 17,"otp_policy_window": 81,"otp_policy_period": 16,"otp_policy_digits": 97,"otp_policy_alg": "ZDjIAsJUBCIxeIyWFMGZfkBHF","otp_policy_type": "KoceZgpiBwbHnchlQrMKgcTpO","browser_flow": "nEHMdOoabBXnYOmMhMViHmHjO","registration_flow": "OLDpdsOtgQfqexUkIlkSjWDng","direct_grant_flow": "YQXNOMiFRbPadjDPXdedYhZge","reset_credentials_flow": "QCRdFcklYuoeBsrBqdnRYZJIc","client_auth_flow": "odvQbPvPOdZysAscRMIiXFmKr","offline_session_idle_timeout": 70,"revoke_refresh_token": false,"access_token_life_implicit": 95,"login_with_email_allowed": false,"duplicate_emails_allowed": true,"docker_auth_flow": "avBMHYgSFSlmBQebmwhwnmnDB","refresh_token_max_reuse": 23,"allow_user_managed_access": false,"sso_max_lifespan_remember_me": 8,"sso_idle_timeout_remember_me": 71,"default_role": "JdLqjRVmkbefNFAUWEumEoCXE"}' | http PUT "http://localhost:8080/realm/hello world"  X-Api-User:user123
func UpdateRealm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	realm := &model.Realm{}
	if err := readJSON(r, realm); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := realm.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	realm.Prepare()

	if err := realm.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "realm", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	realm, _, err = dao.UpdateRealm(ctx,
		argID,
		realm)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, realm)
}

// DeleteRealm Delete a single record from realm table in the keycloak database
// @Summary Delete a record from realm
// @Description Delete a single record from realm table in the keycloak database
// @Tags Realm
// @Accept  json
// @Produce  json
// @Param  argID path string true "id"
// @Success 204 {object} model.Realm
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /realm/{argID} [delete]
// http DELETE "http://localhost:8080/realm/hello world" X-Api-User:user123
func DeleteRealm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseString(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "realm", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteRealm(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
