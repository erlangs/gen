package api

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"unsafe"
	//_ "github.com/satori/go.uuid"

	"keycloak/rest/api/dao"
	"keycloak/rest/api/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
)

var crudEndpoints map[string]*CrudAPI

// CrudAPI describes requests available for tables in the database
type CrudAPI struct {
	Name            string           `json:"name"`
	CreateURL       string           `json:"create_url"`
	RetrieveOneURL  string           `json:"retrieve_one_url"`
	RetrieveManyURL string           `json:"retrieve_many_url"`
	UpdateURL       string           `json:"update_url"`
	DeleteURL       string           `json:"delete_url"`
	FetchDDLURL     string           `json:"fetch_ddl_url"`
	TableInfo       *model.TableInfo `json:"table_info"`
}

// PagedResults results for pages GetAll results.
type PagedResults struct {
	Page         int64       `json:"page"`
	PageSize     int64       `json:"page_size"`
	Data         interface{} `json:"data"`
	TotalRecords int         `json:"total_records"`
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// ConfigRouter configure http.Handler router
func ConfigRouter() http.Handler {
	router := httprouter.New()
	configAdminEventEntityRouter(router)
	configAssociatedPolicyRouter(router)
	configAuthenticationExecutionRouter(router)
	configAuthenticationFlowRouter(router)
	configAuthenticatorConfigRouter(router)
	configAuthenticatorConfigEntryRouter(router)
	configBrokerLinkRouter(router)
	configClientRouter(router)
	configClientAttributesRouter(router)
	configClientAuthFlowBindingsRouter(router)
	configClientInitialAccessRouter(router)
	configClientNodeRegistrationsRouter(router)
	configClientScopeRouter(router)
	configClientScopeAttributesRouter(router)
	configClientScopeClientRouter(router)
	configClientScopeRoleMappingRouter(router)
	configClientSessionRouter(router)
	configClientSessionAuthStatusRouter(router)
	configClientSessionNoteRouter(router)
	configClientSessionProtMapperRouter(router)
	configClientSessionRoleRouter(router)
	configClientUserSessionNoteRouter(router)
	configComponentRouter(router)
	configComponentConfigRouter(router)
	configCompositeRoleRouter(router)
	configCredentialRouter(router)
	configDatabasechangelogRouter(router)
	configDatabasechangeloglockRouter(router)
	configDefaultClientScopeRouter(router)
	configEventEntityRouter(router)
	configFedUserAttributeRouter(router)
	configFedUserConsentRouter(router)
	configFedUserConsentClScopeRouter(router)
	configFedUserCredentialRouter(router)
	configFedUserGroupMembershipRouter(router)
	configFedUserRequiredActionRouter(router)
	configFedUserRoleMappingRouter(router)
	configFederatedIdentityRouter(router)
	configFederatedUserRouter(router)
	configGroupAttributeRouter(router)
	configGroupRoleMappingRouter(router)
	configIdentityProviderRouter(router)
	configIdentityProviderConfigRouter(router)
	configIdentityProviderMapperRouter(router)
	configIdpMapperConfigRouter(router)
	configKeycloakGroupRouter(router)
	configKeycloakRoleRouter(router)
	configMigrationModelRouter(router)
	configOfflineClientSessionRouter(router)
	configOfflineUserSessionRouter(router)
	configPolicyConfigRouter(router)
	configProtocolMapperRouter(router)
	configProtocolMapperConfigRouter(router)
	configRealmRouter(router)
	configRealmAttributeRouter(router)
	configRealmDefaultGroupsRouter(router)
	configRealmEnabledEventTypesRouter(router)
	configRealmEventsListenersRouter(router)
	configRealmLocalizationsRouter(router)
	configRealmRequiredCredentialRouter(router)
	configRealmSMTPConfigRouter(router)
	configRealmSupportedLocalesRouter(router)
	configRedirectUrisRouter(router)
	configRequiredActionConfigRouter(router)
	configRequiredActionProviderRouter(router)
	configResourceAttributeRouter(router)
	configResourcePolicyRouter(router)
	configResourceScopeRouter(router)
	configResourceServerRouter(router)
	configResourceServerPermTicketRouter(router)
	configResourceServerPolicyRouter(router)
	configResourceServerResourceRouter(router)
	configResourceServerScopeRouter(router)
	configResourceUrisRouter(router)
	configRoleAttributeRouter(router)
	configScopeMappingRouter(router)
	configScopePolicyRouter(router)
	configUserRouter(router)
	configUserAttributeRouter(router)
	configUserConsentRouter(router)
	configUserConsentClientScopeRouter(router)
	configUserEntityRouter(router)
	configUserFederationConfigRouter(router)
	configUserFederationMapperRouter(router)
	configUserFederationMapperConfigRouter(router)
	configUserFederationProviderRouter(router)
	configUserGroupMembershipRouter(router)
	configUserRequiredActionRouter(router)
	configUserRoleMappingRouter(router)
	configUserSessionRouter(router)
	configUserSessionNoteRouter(router)
	configUsernameLoginFailureRouter(router)
	configWebOriginsRouter(router)

	router.GET("/ddl/:argID", GetDdl)
	router.GET("/ddl", GetDdlEndpoints)
	return router
}

// ConfigGinRouter configure gin router
func ConfigGinRouter(router gin.IRoutes) {
	configGinAdminEventEntityRouter(router)
	configGinAssociatedPolicyRouter(router)
	configGinAuthenticationExecutionRouter(router)
	configGinAuthenticationFlowRouter(router)
	configGinAuthenticatorConfigRouter(router)
	configGinAuthenticatorConfigEntryRouter(router)
	configGinBrokerLinkRouter(router)
	configGinClientRouter(router)
	configGinClientAttributesRouter(router)
	configGinClientAuthFlowBindingsRouter(router)
	configGinClientInitialAccessRouter(router)
	configGinClientNodeRegistrationsRouter(router)
	configGinClientScopeRouter(router)
	configGinClientScopeAttributesRouter(router)
	configGinClientScopeClientRouter(router)
	configGinClientScopeRoleMappingRouter(router)
	configGinClientSessionRouter(router)
	configGinClientSessionAuthStatusRouter(router)
	configGinClientSessionNoteRouter(router)
	configGinClientSessionProtMapperRouter(router)
	configGinClientSessionRoleRouter(router)
	configGinClientUserSessionNoteRouter(router)
	configGinComponentRouter(router)
	configGinComponentConfigRouter(router)
	configGinCompositeRoleRouter(router)
	configGinCredentialRouter(router)
	configGinDatabasechangelogRouter(router)
	configGinDatabasechangeloglockRouter(router)
	configGinDefaultClientScopeRouter(router)
	configGinEventEntityRouter(router)
	configGinFedUserAttributeRouter(router)
	configGinFedUserConsentRouter(router)
	configGinFedUserConsentClScopeRouter(router)
	configGinFedUserCredentialRouter(router)
	configGinFedUserGroupMembershipRouter(router)
	configGinFedUserRequiredActionRouter(router)
	configGinFedUserRoleMappingRouter(router)
	configGinFederatedIdentityRouter(router)
	configGinFederatedUserRouter(router)
	configGinGroupAttributeRouter(router)
	configGinGroupRoleMappingRouter(router)
	configGinIdentityProviderRouter(router)
	configGinIdentityProviderConfigRouter(router)
	configGinIdentityProviderMapperRouter(router)
	configGinIdpMapperConfigRouter(router)
	configGinKeycloakGroupRouter(router)
	configGinKeycloakRoleRouter(router)
	configGinMigrationModelRouter(router)
	configGinOfflineClientSessionRouter(router)
	configGinOfflineUserSessionRouter(router)
	configGinPolicyConfigRouter(router)
	configGinProtocolMapperRouter(router)
	configGinProtocolMapperConfigRouter(router)
	configGinRealmRouter(router)
	configGinRealmAttributeRouter(router)
	configGinRealmDefaultGroupsRouter(router)
	configGinRealmEnabledEventTypesRouter(router)
	configGinRealmEventsListenersRouter(router)
	configGinRealmLocalizationsRouter(router)
	configGinRealmRequiredCredentialRouter(router)
	configGinRealmSMTPConfigRouter(router)
	configGinRealmSupportedLocalesRouter(router)
	configGinRedirectUrisRouter(router)
	configGinRequiredActionConfigRouter(router)
	configGinRequiredActionProviderRouter(router)
	configGinResourceAttributeRouter(router)
	configGinResourcePolicyRouter(router)
	configGinResourceScopeRouter(router)
	configGinResourceServerRouter(router)
	configGinResourceServerPermTicketRouter(router)
	configGinResourceServerPolicyRouter(router)
	configGinResourceServerResourceRouter(router)
	configGinResourceServerScopeRouter(router)
	configGinResourceUrisRouter(router)
	configGinRoleAttributeRouter(router)
	configGinScopeMappingRouter(router)
	configGinScopePolicyRouter(router)
	configGinUserRouter(router)
	configGinUserAttributeRouter(router)
	configGinUserConsentRouter(router)
	configGinUserConsentClientScopeRouter(router)
	configGinUserEntityRouter(router)
	configGinUserFederationConfigRouter(router)
	configGinUserFederationMapperRouter(router)
	configGinUserFederationMapperConfigRouter(router)
	configGinUserFederationProviderRouter(router)
	configGinUserGroupMembershipRouter(router)
	configGinUserRequiredActionRouter(router)
	configGinUserRoleMappingRouter(router)
	configGinUserSessionRouter(router)
	configGinUserSessionNoteRouter(router)
	configGinUsernameLoginFailureRouter(router)
	configGinWebOriginsRouter(router)

	router.GET("/ddl/:argID", ConverHttprouterToGin(GetDdl))
	router.GET("/ddl", ConverHttprouterToGin(GetDdlEndpoints))
	return
}

// ConverHttprouterToGin wrap httprouter.Handle to gin.HandlerFunc
func ConverHttprouterToGin(f httprouter.Handle) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params httprouter.Params
		_len := len(c.Params)
		if _len == 0 {
			params = nil
		} else {
			params = ((*[1 << 10]httprouter.Param)(unsafe.Pointer(&c.Params[0])))[:_len]
		}

		f(c.Writer, c.Request, params)
	}
}

func initializeContext(r *http.Request) (ctx context.Context) {
	if ContextInitializer != nil {
		ctx = ContextInitializer(r)
	} else {
		ctx = r.Context()
	}
	return ctx
}

func ValidateRequest(ctx context.Context, r *http.Request, table string, action model.Action) error {
	if RequestValidator != nil {
		return RequestValidator(ctx, r, table, action)
	}

	return nil
}

type RequestValidatorFunc func(ctx context.Context, r *http.Request, table string, action model.Action) error

var RequestValidator RequestValidatorFunc

type ContextInitializerFunc func(r *http.Request) (ctx context.Context)

var ContextInitializer ContextInitializerFunc

func readInt(r *http.Request, param string, v int64) (int64, error) {
	p := r.FormValue(param)
	if p == "" {
		return v, nil
	}

	return strconv.ParseInt(p, 10, 64)
}

func writeJSON(ctx context.Context, w http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func writeRowsAffected(w http.ResponseWriter, rowsAffected int64) {
	data, _ := json.Marshal(rowsAffected)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func readJSON(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, v)
}

func returnError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	status := 0
	switch err {
	case dao.ErrNotFound:
		status = http.StatusBadRequest
	case dao.ErrUnableToMarshalJSON:
		status = http.StatusBadRequest
	case dao.ErrUpdateFailed:
		status = http.StatusBadRequest
	case dao.ErrInsertFailed:
		status = http.StatusBadRequest
	case dao.ErrDeleteFailed:
		status = http.StatusBadRequest
	case dao.ErrBadParams:
		status = http.StatusBadRequest
	default:
		status = http.StatusBadRequest
	}
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}

	SendJSON(w, r, er.Code, er)
}

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

func parseUint8(ps httprouter.Params, key string) (uint8, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 8)
	if err != nil {
		return uint8(id), err
	}
	return uint8(id), err
}
func parseUint16(ps httprouter.Params, key string) (uint16, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil {
		return uint16(id), err
	}
	return uint16(id), err
}
func parseUint32(ps httprouter.Params, key string) (uint32, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return uint32(id), err
	}
	return uint32(id), err
}
func parseUint64(ps httprouter.Params, key string) (uint64, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return uint64(id), err
	}
	return uint64(id), err
}
func parseInt(ps httprouter.Params, key string) (int, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return -1, err
	}
	return int(id), err
}
func parseInt8(ps httprouter.Params, key string) (int8, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 8)
	if err != nil {
		return -1, err
	}
	return int8(id), err
}
func parseInt16(ps httprouter.Params, key string) (int16, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil {
		return -1, err
	}
	return int16(id), err
}
func parseInt32(ps httprouter.Params, key string) (int32, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return -1, err
	}
	return int32(id), err
}
func parseInt64(ps httprouter.Params, key string) (int64, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 54)
	if err != nil {
		return -1, err
	}
	return id, err
}
func parseString(ps httprouter.Params, key string) (string, error) {
	idStr := ps.ByName(key)
	return idStr, nil
}
func parseUUID(ps httprouter.Params, key string) (string, error) {
	idStr := ps.ByName(key)
	return idStr, nil
}

func parseBytes(ps httprouter.Params, key string) (string, error) {
	idStr := ps.ByName(key)
	decByte, _ := hex.DecodeString(idStr)
	return string(decByte), nil
	//return hex.DecodeString(idStr)
}

// GetDdl is a function to get table info for a table in the keycloak database
// @Summary Get table info for a table in the keycloak database by argID
// @Tags TableInfo
// @ID argID
// @Description GetDdl is a function to get table info for a table in the keycloak database
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 200 {object} api.CrudAPI
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /ddl/{argID} [get]
// http "http://localhost:8080/ddl/xyz" X-Api-User:user123
func GetDdl(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID := ps.ByName("argID")

	if err := ValidateRequest(ctx, r, "ddl", model.FetchDDL); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, ok := crudEndpoints[argID]
	if !ok {
		returnError(ctx, w, r, fmt.Errorf("unable to find table: %s", argID))
		return
	}

	writeJSON(ctx, w, record)
}

// GetDdlEndpoints is a function to get a list of ddl endpoints available for tables in the keycloak database
// @Summary Gets a list of ddl endpoints available for tables in the keycloak database
// @Tags TableInfo
// @Description GetDdlEndpoints is a function to get a list of ddl endpoints available for tables in the keycloak database
// @Accept  json
// @Produce  json
// @Success 200 {object} api.CrudAPI
// @Router /ddl [get]
// http "http://localhost:8080/ddl" X-Api-User:user123
func GetDdlEndpoints(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	if err := ValidateRequest(ctx, r, "ddl", model.FetchDDL); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, crudEndpoints)
}

func init() {
	crudEndpoints = make(map[string]*CrudAPI)

	var tmp *CrudAPI

	tmp = &CrudAPI{
		Name:            "admin_event_entity",
		CreateURL:       "/adminevententity",
		RetrieveOneURL:  "/adminevententity",
		RetrieveManyURL: "/adminevententity",
		UpdateURL:       "/adminevententity",
		DeleteURL:       "/adminevententity",
		FetchDDLURL:     "/ddl/admin_event_entity",
	}

	tmp.TableInfo, _ = model.GetTableInfo("admin_event_entity")
	crudEndpoints["admin_event_entity"] = tmp

	tmp = &CrudAPI{
		Name:            "associated_policy",
		CreateURL:       "/associatedpolicy",
		RetrieveOneURL:  "/associatedpolicy",
		RetrieveManyURL: "/associatedpolicy",
		UpdateURL:       "/associatedpolicy",
		DeleteURL:       "/associatedpolicy",
		FetchDDLURL:     "/ddl/associated_policy",
	}

	tmp.TableInfo, _ = model.GetTableInfo("associated_policy")
	crudEndpoints["associated_policy"] = tmp

	tmp = &CrudAPI{
		Name:            "authentication_execution",
		CreateURL:       "/authenticationexecution",
		RetrieveOneURL:  "/authenticationexecution",
		RetrieveManyURL: "/authenticationexecution",
		UpdateURL:       "/authenticationexecution",
		DeleteURL:       "/authenticationexecution",
		FetchDDLURL:     "/ddl/authentication_execution",
	}

	tmp.TableInfo, _ = model.GetTableInfo("authentication_execution")
	crudEndpoints["authentication_execution"] = tmp

	tmp = &CrudAPI{
		Name:            "authentication_flow",
		CreateURL:       "/authenticationflow",
		RetrieveOneURL:  "/authenticationflow",
		RetrieveManyURL: "/authenticationflow",
		UpdateURL:       "/authenticationflow",
		DeleteURL:       "/authenticationflow",
		FetchDDLURL:     "/ddl/authentication_flow",
	}

	tmp.TableInfo, _ = model.GetTableInfo("authentication_flow")
	crudEndpoints["authentication_flow"] = tmp

	tmp = &CrudAPI{
		Name:            "authenticator_config",
		CreateURL:       "/authenticatorconfig",
		RetrieveOneURL:  "/authenticatorconfig",
		RetrieveManyURL: "/authenticatorconfig",
		UpdateURL:       "/authenticatorconfig",
		DeleteURL:       "/authenticatorconfig",
		FetchDDLURL:     "/ddl/authenticator_config",
	}

	tmp.TableInfo, _ = model.GetTableInfo("authenticator_config")
	crudEndpoints["authenticator_config"] = tmp

	tmp = &CrudAPI{
		Name:            "authenticator_config_entry",
		CreateURL:       "/authenticatorconfigentry",
		RetrieveOneURL:  "/authenticatorconfigentry",
		RetrieveManyURL: "/authenticatorconfigentry",
		UpdateURL:       "/authenticatorconfigentry",
		DeleteURL:       "/authenticatorconfigentry",
		FetchDDLURL:     "/ddl/authenticator_config_entry",
	}

	tmp.TableInfo, _ = model.GetTableInfo("authenticator_config_entry")
	crudEndpoints["authenticator_config_entry"] = tmp

	tmp = &CrudAPI{
		Name:            "broker_link",
		CreateURL:       "/brokerlink",
		RetrieveOneURL:  "/brokerlink",
		RetrieveManyURL: "/brokerlink",
		UpdateURL:       "/brokerlink",
		DeleteURL:       "/brokerlink",
		FetchDDLURL:     "/ddl/broker_link",
	}

	tmp.TableInfo, _ = model.GetTableInfo("broker_link")
	crudEndpoints["broker_link"] = tmp

	tmp = &CrudAPI{
		Name:            "client",
		CreateURL:       "/client",
		RetrieveOneURL:  "/client",
		RetrieveManyURL: "/client",
		UpdateURL:       "/client",
		DeleteURL:       "/client",
		FetchDDLURL:     "/ddl/client",
	}

	tmp.TableInfo, _ = model.GetTableInfo("client")
	crudEndpoints["client"] = tmp

	tmp = &CrudAPI{
		Name:            "client_attributes",
		CreateURL:       "/clientattributes",
		RetrieveOneURL:  "/clientattributes",
		RetrieveManyURL: "/clientattributes",
		UpdateURL:       "/clientattributes",
		DeleteURL:       "/clientattributes",
		FetchDDLURL:     "/ddl/client_attributes",
	}

	tmp.TableInfo, _ = model.GetTableInfo("client_attributes")
	crudEndpoints["client_attributes"] = tmp

	tmp = &CrudAPI{
		Name:            "client_auth_flow_bindings",
		CreateURL:       "/clientauthflowbindings",
		RetrieveOneURL:  "/clientauthflowbindings",
		RetrieveManyURL: "/clientauthflowbindings",
		UpdateURL:       "/clientauthflowbindings",
		DeleteURL:       "/clientauthflowbindings",
		FetchDDLURL:     "/ddl/client_auth_flow_bindings",
	}

	tmp.TableInfo, _ = model.GetTableInfo("client_auth_flow_bindings")
	crudEndpoints["client_auth_flow_bindings"] = tmp

	tmp = &CrudAPI{
		Name:            "client_initial_access",
		CreateURL:       "/clientinitialaccess",
		RetrieveOneURL:  "/clientinitialaccess",
		RetrieveManyURL: "/clientinitialaccess",
		UpdateURL:       "/clientinitialaccess",
		DeleteURL:       "/clientinitialaccess",
		FetchDDLURL:     "/ddl/client_initial_access",
	}

	tmp.TableInfo, _ = model.GetTableInfo("client_initial_access")
	crudEndpoints["client_initial_access"] = tmp

	tmp = &CrudAPI{
		Name:            "client_node_registrations",
		CreateURL:       "/clientnoderegistrations",
		RetrieveOneURL:  "/clientnoderegistrations",
		RetrieveManyURL: "/clientnoderegistrations",
		UpdateURL:       "/clientnoderegistrations",
		DeleteURL:       "/clientnoderegistrations",
		FetchDDLURL:     "/ddl/client_node_registrations",
	}

	tmp.TableInfo, _ = model.GetTableInfo("client_node_registrations")
	crudEndpoints["client_node_registrations"] = tmp

	tmp = &CrudAPI{
		Name:            "client_scope",
		CreateURL:       "/clientscope",
		RetrieveOneURL:  "/clientscope",
		RetrieveManyURL: "/clientscope",
		UpdateURL:       "/clientscope",
		DeleteURL:       "/clientscope",
		FetchDDLURL:     "/ddl/client_scope",
	}

	tmp.TableInfo, _ = model.GetTableInfo("client_scope")
	crudEndpoints["client_scope"] = tmp

	tmp = &CrudAPI{
		Name:            "client_scope_attributes",
		CreateURL:       "/clientscopeattributes",
		RetrieveOneURL:  "/clientscopeattributes",
		RetrieveManyURL: "/clientscopeattributes",
		UpdateURL:       "/clientscopeattributes",
		DeleteURL:       "/clientscopeattributes",
		FetchDDLURL:     "/ddl/client_scope_attributes",
	}

	tmp.TableInfo, _ = model.GetTableInfo("client_scope_attributes")
	crudEndpoints["client_scope_attributes"] = tmp

	tmp = &CrudAPI{
		Name:            "client_scope_client",
		CreateURL:       "/clientscopeclient",
		RetrieveOneURL:  "/clientscopeclient",
		RetrieveManyURL: "/clientscopeclient",
		UpdateURL:       "/clientscopeclient",
		DeleteURL:       "/clientscopeclient",
		FetchDDLURL:     "/ddl/client_scope_client",
	}

	tmp.TableInfo, _ = model.GetTableInfo("client_scope_client")
	crudEndpoints["client_scope_client"] = tmp

	tmp = &CrudAPI{
		Name:            "client_scope_role_mapping",
		CreateURL:       "/clientscoperolemapping",
		RetrieveOneURL:  "/clientscoperolemapping",
		RetrieveManyURL: "/clientscoperolemapping",
		UpdateURL:       "/clientscoperolemapping",
		DeleteURL:       "/clientscoperolemapping",
		FetchDDLURL:     "/ddl/client_scope_role_mapping",
	}

	tmp.TableInfo, _ = model.GetTableInfo("client_scope_role_mapping")
	crudEndpoints["client_scope_role_mapping"] = tmp

	tmp = &CrudAPI{
		Name:            "client_session",
		CreateURL:       "/clientsession",
		RetrieveOneURL:  "/clientsession",
		RetrieveManyURL: "/clientsession",
		UpdateURL:       "/clientsession",
		DeleteURL:       "/clientsession",
		FetchDDLURL:     "/ddl/client_session",
	}

	tmp.TableInfo, _ = model.GetTableInfo("client_session")
	crudEndpoints["client_session"] = tmp

	tmp = &CrudAPI{
		Name:            "client_session_auth_status",
		CreateURL:       "/clientsessionauthstatus",
		RetrieveOneURL:  "/clientsessionauthstatus",
		RetrieveManyURL: "/clientsessionauthstatus",
		UpdateURL:       "/clientsessionauthstatus",
		DeleteURL:       "/clientsessionauthstatus",
		FetchDDLURL:     "/ddl/client_session_auth_status",
	}

	tmp.TableInfo, _ = model.GetTableInfo("client_session_auth_status")
	crudEndpoints["client_session_auth_status"] = tmp

	tmp = &CrudAPI{
		Name:            "client_session_note",
		CreateURL:       "/clientsessionnote",
		RetrieveOneURL:  "/clientsessionnote",
		RetrieveManyURL: "/clientsessionnote",
		UpdateURL:       "/clientsessionnote",
		DeleteURL:       "/clientsessionnote",
		FetchDDLURL:     "/ddl/client_session_note",
	}

	tmp.TableInfo, _ = model.GetTableInfo("client_session_note")
	crudEndpoints["client_session_note"] = tmp

	tmp = &CrudAPI{
		Name:            "client_session_prot_mapper",
		CreateURL:       "/clientsessionprotmapper",
		RetrieveOneURL:  "/clientsessionprotmapper",
		RetrieveManyURL: "/clientsessionprotmapper",
		UpdateURL:       "/clientsessionprotmapper",
		DeleteURL:       "/clientsessionprotmapper",
		FetchDDLURL:     "/ddl/client_session_prot_mapper",
	}

	tmp.TableInfo, _ = model.GetTableInfo("client_session_prot_mapper")
	crudEndpoints["client_session_prot_mapper"] = tmp

	tmp = &CrudAPI{
		Name:            "client_session_role",
		CreateURL:       "/clientsessionrole",
		RetrieveOneURL:  "/clientsessionrole",
		RetrieveManyURL: "/clientsessionrole",
		UpdateURL:       "/clientsessionrole",
		DeleteURL:       "/clientsessionrole",
		FetchDDLURL:     "/ddl/client_session_role",
	}

	tmp.TableInfo, _ = model.GetTableInfo("client_session_role")
	crudEndpoints["client_session_role"] = tmp

	tmp = &CrudAPI{
		Name:            "client_user_session_note",
		CreateURL:       "/clientusersessionnote",
		RetrieveOneURL:  "/clientusersessionnote",
		RetrieveManyURL: "/clientusersessionnote",
		UpdateURL:       "/clientusersessionnote",
		DeleteURL:       "/clientusersessionnote",
		FetchDDLURL:     "/ddl/client_user_session_note",
	}

	tmp.TableInfo, _ = model.GetTableInfo("client_user_session_note")
	crudEndpoints["client_user_session_note"] = tmp

	tmp = &CrudAPI{
		Name:            "component",
		CreateURL:       "/component",
		RetrieveOneURL:  "/component",
		RetrieveManyURL: "/component",
		UpdateURL:       "/component",
		DeleteURL:       "/component",
		FetchDDLURL:     "/ddl/component",
	}

	tmp.TableInfo, _ = model.GetTableInfo("component")
	crudEndpoints["component"] = tmp

	tmp = &CrudAPI{
		Name:            "component_config",
		CreateURL:       "/componentconfig",
		RetrieveOneURL:  "/componentconfig",
		RetrieveManyURL: "/componentconfig",
		UpdateURL:       "/componentconfig",
		DeleteURL:       "/componentconfig",
		FetchDDLURL:     "/ddl/component_config",
	}

	tmp.TableInfo, _ = model.GetTableInfo("component_config")
	crudEndpoints["component_config"] = tmp

	tmp = &CrudAPI{
		Name:            "composite_role",
		CreateURL:       "/compositerole",
		RetrieveOneURL:  "/compositerole",
		RetrieveManyURL: "/compositerole",
		UpdateURL:       "/compositerole",
		DeleteURL:       "/compositerole",
		FetchDDLURL:     "/ddl/composite_role",
	}

	tmp.TableInfo, _ = model.GetTableInfo("composite_role")
	crudEndpoints["composite_role"] = tmp

	tmp = &CrudAPI{
		Name:            "credential",
		CreateURL:       "/credential",
		RetrieveOneURL:  "/credential",
		RetrieveManyURL: "/credential",
		UpdateURL:       "/credential",
		DeleteURL:       "/credential",
		FetchDDLURL:     "/ddl/credential",
	}

	tmp.TableInfo, _ = model.GetTableInfo("credential")
	crudEndpoints["credential"] = tmp

	tmp = &CrudAPI{
		Name:            "databasechangelog",
		CreateURL:       "/databasechangelog",
		RetrieveOneURL:  "/databasechangelog",
		RetrieveManyURL: "/databasechangelog",
		UpdateURL:       "/databasechangelog",
		DeleteURL:       "/databasechangelog",
		FetchDDLURL:     "/ddl/databasechangelog",
	}

	tmp.TableInfo, _ = model.GetTableInfo("databasechangelog")
	crudEndpoints["databasechangelog"] = tmp

	tmp = &CrudAPI{
		Name:            "databasechangeloglock",
		CreateURL:       "/databasechangeloglock",
		RetrieveOneURL:  "/databasechangeloglock",
		RetrieveManyURL: "/databasechangeloglock",
		UpdateURL:       "/databasechangeloglock",
		DeleteURL:       "/databasechangeloglock",
		FetchDDLURL:     "/ddl/databasechangeloglock",
	}

	tmp.TableInfo, _ = model.GetTableInfo("databasechangeloglock")
	crudEndpoints["databasechangeloglock"] = tmp

	tmp = &CrudAPI{
		Name:            "default_client_scope",
		CreateURL:       "/defaultclientscope",
		RetrieveOneURL:  "/defaultclientscope",
		RetrieveManyURL: "/defaultclientscope",
		UpdateURL:       "/defaultclientscope",
		DeleteURL:       "/defaultclientscope",
		FetchDDLURL:     "/ddl/default_client_scope",
	}

	tmp.TableInfo, _ = model.GetTableInfo("default_client_scope")
	crudEndpoints["default_client_scope"] = tmp

	tmp = &CrudAPI{
		Name:            "event_entity",
		CreateURL:       "/evententity",
		RetrieveOneURL:  "/evententity",
		RetrieveManyURL: "/evententity",
		UpdateURL:       "/evententity",
		DeleteURL:       "/evententity",
		FetchDDLURL:     "/ddl/event_entity",
	}

	tmp.TableInfo, _ = model.GetTableInfo("event_entity")
	crudEndpoints["event_entity"] = tmp

	tmp = &CrudAPI{
		Name:            "fed_user_attribute",
		CreateURL:       "/feduserattribute",
		RetrieveOneURL:  "/feduserattribute",
		RetrieveManyURL: "/feduserattribute",
		UpdateURL:       "/feduserattribute",
		DeleteURL:       "/feduserattribute",
		FetchDDLURL:     "/ddl/fed_user_attribute",
	}

	tmp.TableInfo, _ = model.GetTableInfo("fed_user_attribute")
	crudEndpoints["fed_user_attribute"] = tmp

	tmp = &CrudAPI{
		Name:            "fed_user_consent",
		CreateURL:       "/feduserconsent",
		RetrieveOneURL:  "/feduserconsent",
		RetrieveManyURL: "/feduserconsent",
		UpdateURL:       "/feduserconsent",
		DeleteURL:       "/feduserconsent",
		FetchDDLURL:     "/ddl/fed_user_consent",
	}

	tmp.TableInfo, _ = model.GetTableInfo("fed_user_consent")
	crudEndpoints["fed_user_consent"] = tmp

	tmp = &CrudAPI{
		Name:            "fed_user_consent_cl_scope",
		CreateURL:       "/feduserconsentclscope",
		RetrieveOneURL:  "/feduserconsentclscope",
		RetrieveManyURL: "/feduserconsentclscope",
		UpdateURL:       "/feduserconsentclscope",
		DeleteURL:       "/feduserconsentclscope",
		FetchDDLURL:     "/ddl/fed_user_consent_cl_scope",
	}

	tmp.TableInfo, _ = model.GetTableInfo("fed_user_consent_cl_scope")
	crudEndpoints["fed_user_consent_cl_scope"] = tmp

	tmp = &CrudAPI{
		Name:            "fed_user_credential",
		CreateURL:       "/fedusercredential",
		RetrieveOneURL:  "/fedusercredential",
		RetrieveManyURL: "/fedusercredential",
		UpdateURL:       "/fedusercredential",
		DeleteURL:       "/fedusercredential",
		FetchDDLURL:     "/ddl/fed_user_credential",
	}

	tmp.TableInfo, _ = model.GetTableInfo("fed_user_credential")
	crudEndpoints["fed_user_credential"] = tmp

	tmp = &CrudAPI{
		Name:            "fed_user_group_membership",
		CreateURL:       "/fedusergroupmembership",
		RetrieveOneURL:  "/fedusergroupmembership",
		RetrieveManyURL: "/fedusergroupmembership",
		UpdateURL:       "/fedusergroupmembership",
		DeleteURL:       "/fedusergroupmembership",
		FetchDDLURL:     "/ddl/fed_user_group_membership",
	}

	tmp.TableInfo, _ = model.GetTableInfo("fed_user_group_membership")
	crudEndpoints["fed_user_group_membership"] = tmp

	tmp = &CrudAPI{
		Name:            "fed_user_required_action",
		CreateURL:       "/feduserrequiredaction",
		RetrieveOneURL:  "/feduserrequiredaction",
		RetrieveManyURL: "/feduserrequiredaction",
		UpdateURL:       "/feduserrequiredaction",
		DeleteURL:       "/feduserrequiredaction",
		FetchDDLURL:     "/ddl/fed_user_required_action",
	}

	tmp.TableInfo, _ = model.GetTableInfo("fed_user_required_action")
	crudEndpoints["fed_user_required_action"] = tmp

	tmp = &CrudAPI{
		Name:            "fed_user_role_mapping",
		CreateURL:       "/feduserrolemapping",
		RetrieveOneURL:  "/feduserrolemapping",
		RetrieveManyURL: "/feduserrolemapping",
		UpdateURL:       "/feduserrolemapping",
		DeleteURL:       "/feduserrolemapping",
		FetchDDLURL:     "/ddl/fed_user_role_mapping",
	}

	tmp.TableInfo, _ = model.GetTableInfo("fed_user_role_mapping")
	crudEndpoints["fed_user_role_mapping"] = tmp

	tmp = &CrudAPI{
		Name:            "federated_identity",
		CreateURL:       "/federatedidentity",
		RetrieveOneURL:  "/federatedidentity",
		RetrieveManyURL: "/federatedidentity",
		UpdateURL:       "/federatedidentity",
		DeleteURL:       "/federatedidentity",
		FetchDDLURL:     "/ddl/federated_identity",
	}

	tmp.TableInfo, _ = model.GetTableInfo("federated_identity")
	crudEndpoints["federated_identity"] = tmp

	tmp = &CrudAPI{
		Name:            "federated_user",
		CreateURL:       "/federateduser",
		RetrieveOneURL:  "/federateduser",
		RetrieveManyURL: "/federateduser",
		UpdateURL:       "/federateduser",
		DeleteURL:       "/federateduser",
		FetchDDLURL:     "/ddl/federated_user",
	}

	tmp.TableInfo, _ = model.GetTableInfo("federated_user")
	crudEndpoints["federated_user"] = tmp

	tmp = &CrudAPI{
		Name:            "group_attribute",
		CreateURL:       "/groupattribute",
		RetrieveOneURL:  "/groupattribute",
		RetrieveManyURL: "/groupattribute",
		UpdateURL:       "/groupattribute",
		DeleteURL:       "/groupattribute",
		FetchDDLURL:     "/ddl/group_attribute",
	}

	tmp.TableInfo, _ = model.GetTableInfo("group_attribute")
	crudEndpoints["group_attribute"] = tmp

	tmp = &CrudAPI{
		Name:            "group_role_mapping",
		CreateURL:       "/grouprolemapping",
		RetrieveOneURL:  "/grouprolemapping",
		RetrieveManyURL: "/grouprolemapping",
		UpdateURL:       "/grouprolemapping",
		DeleteURL:       "/grouprolemapping",
		FetchDDLURL:     "/ddl/group_role_mapping",
	}

	tmp.TableInfo, _ = model.GetTableInfo("group_role_mapping")
	crudEndpoints["group_role_mapping"] = tmp

	tmp = &CrudAPI{
		Name:            "identity_provider",
		CreateURL:       "/identityprovider",
		RetrieveOneURL:  "/identityprovider",
		RetrieveManyURL: "/identityprovider",
		UpdateURL:       "/identityprovider",
		DeleteURL:       "/identityprovider",
		FetchDDLURL:     "/ddl/identity_provider",
	}

	tmp.TableInfo, _ = model.GetTableInfo("identity_provider")
	crudEndpoints["identity_provider"] = tmp

	tmp = &CrudAPI{
		Name:            "identity_provider_config",
		CreateURL:       "/identityproviderconfig",
		RetrieveOneURL:  "/identityproviderconfig",
		RetrieveManyURL: "/identityproviderconfig",
		UpdateURL:       "/identityproviderconfig",
		DeleteURL:       "/identityproviderconfig",
		FetchDDLURL:     "/ddl/identity_provider_config",
	}

	tmp.TableInfo, _ = model.GetTableInfo("identity_provider_config")
	crudEndpoints["identity_provider_config"] = tmp

	tmp = &CrudAPI{
		Name:            "identity_provider_mapper",
		CreateURL:       "/identityprovidermapper",
		RetrieveOneURL:  "/identityprovidermapper",
		RetrieveManyURL: "/identityprovidermapper",
		UpdateURL:       "/identityprovidermapper",
		DeleteURL:       "/identityprovidermapper",
		FetchDDLURL:     "/ddl/identity_provider_mapper",
	}

	tmp.TableInfo, _ = model.GetTableInfo("identity_provider_mapper")
	crudEndpoints["identity_provider_mapper"] = tmp

	tmp = &CrudAPI{
		Name:            "idp_mapper_config",
		CreateURL:       "/idpmapperconfig",
		RetrieveOneURL:  "/idpmapperconfig",
		RetrieveManyURL: "/idpmapperconfig",
		UpdateURL:       "/idpmapperconfig",
		DeleteURL:       "/idpmapperconfig",
		FetchDDLURL:     "/ddl/idp_mapper_config",
	}

	tmp.TableInfo, _ = model.GetTableInfo("idp_mapper_config")
	crudEndpoints["idp_mapper_config"] = tmp

	tmp = &CrudAPI{
		Name:            "keycloak_group",
		CreateURL:       "/keycloakgroup",
		RetrieveOneURL:  "/keycloakgroup",
		RetrieveManyURL: "/keycloakgroup",
		UpdateURL:       "/keycloakgroup",
		DeleteURL:       "/keycloakgroup",
		FetchDDLURL:     "/ddl/keycloak_group",
	}

	tmp.TableInfo, _ = model.GetTableInfo("keycloak_group")
	crudEndpoints["keycloak_group"] = tmp

	tmp = &CrudAPI{
		Name:            "keycloak_role",
		CreateURL:       "/keycloakrole",
		RetrieveOneURL:  "/keycloakrole",
		RetrieveManyURL: "/keycloakrole",
		UpdateURL:       "/keycloakrole",
		DeleteURL:       "/keycloakrole",
		FetchDDLURL:     "/ddl/keycloak_role",
	}

	tmp.TableInfo, _ = model.GetTableInfo("keycloak_role")
	crudEndpoints["keycloak_role"] = tmp

	tmp = &CrudAPI{
		Name:            "migration_model",
		CreateURL:       "/migrationmodel",
		RetrieveOneURL:  "/migrationmodel",
		RetrieveManyURL: "/migrationmodel",
		UpdateURL:       "/migrationmodel",
		DeleteURL:       "/migrationmodel",
		FetchDDLURL:     "/ddl/migration_model",
	}

	tmp.TableInfo, _ = model.GetTableInfo("migration_model")
	crudEndpoints["migration_model"] = tmp

	tmp = &CrudAPI{
		Name:            "offline_client_session",
		CreateURL:       "/offlineclientsession",
		RetrieveOneURL:  "/offlineclientsession",
		RetrieveManyURL: "/offlineclientsession",
		UpdateURL:       "/offlineclientsession",
		DeleteURL:       "/offlineclientsession",
		FetchDDLURL:     "/ddl/offline_client_session",
	}

	tmp.TableInfo, _ = model.GetTableInfo("offline_client_session")
	crudEndpoints["offline_client_session"] = tmp

	tmp = &CrudAPI{
		Name:            "offline_user_session",
		CreateURL:       "/offlineusersession",
		RetrieveOneURL:  "/offlineusersession",
		RetrieveManyURL: "/offlineusersession",
		UpdateURL:       "/offlineusersession",
		DeleteURL:       "/offlineusersession",
		FetchDDLURL:     "/ddl/offline_user_session",
	}

	tmp.TableInfo, _ = model.GetTableInfo("offline_user_session")
	crudEndpoints["offline_user_session"] = tmp

	tmp = &CrudAPI{
		Name:            "policy_config",
		CreateURL:       "/policyconfig",
		RetrieveOneURL:  "/policyconfig",
		RetrieveManyURL: "/policyconfig",
		UpdateURL:       "/policyconfig",
		DeleteURL:       "/policyconfig",
		FetchDDLURL:     "/ddl/policy_config",
	}

	tmp.TableInfo, _ = model.GetTableInfo("policy_config")
	crudEndpoints["policy_config"] = tmp

	tmp = &CrudAPI{
		Name:            "protocol_mapper",
		CreateURL:       "/protocolmapper",
		RetrieveOneURL:  "/protocolmapper",
		RetrieveManyURL: "/protocolmapper",
		UpdateURL:       "/protocolmapper",
		DeleteURL:       "/protocolmapper",
		FetchDDLURL:     "/ddl/protocol_mapper",
	}

	tmp.TableInfo, _ = model.GetTableInfo("protocol_mapper")
	crudEndpoints["protocol_mapper"] = tmp

	tmp = &CrudAPI{
		Name:            "protocol_mapper_config",
		CreateURL:       "/protocolmapperconfig",
		RetrieveOneURL:  "/protocolmapperconfig",
		RetrieveManyURL: "/protocolmapperconfig",
		UpdateURL:       "/protocolmapperconfig",
		DeleteURL:       "/protocolmapperconfig",
		FetchDDLURL:     "/ddl/protocol_mapper_config",
	}

	tmp.TableInfo, _ = model.GetTableInfo("protocol_mapper_config")
	crudEndpoints["protocol_mapper_config"] = tmp

	tmp = &CrudAPI{
		Name:            "realm",
		CreateURL:       "/realm",
		RetrieveOneURL:  "/realm",
		RetrieveManyURL: "/realm",
		UpdateURL:       "/realm",
		DeleteURL:       "/realm",
		FetchDDLURL:     "/ddl/realm",
	}

	tmp.TableInfo, _ = model.GetTableInfo("realm")
	crudEndpoints["realm"] = tmp

	tmp = &CrudAPI{
		Name:            "realm_attribute",
		CreateURL:       "/realmattribute",
		RetrieveOneURL:  "/realmattribute",
		RetrieveManyURL: "/realmattribute",
		UpdateURL:       "/realmattribute",
		DeleteURL:       "/realmattribute",
		FetchDDLURL:     "/ddl/realm_attribute",
	}

	tmp.TableInfo, _ = model.GetTableInfo("realm_attribute")
	crudEndpoints["realm_attribute"] = tmp

	tmp = &CrudAPI{
		Name:            "realm_default_groups",
		CreateURL:       "/realmdefaultgroups",
		RetrieveOneURL:  "/realmdefaultgroups",
		RetrieveManyURL: "/realmdefaultgroups",
		UpdateURL:       "/realmdefaultgroups",
		DeleteURL:       "/realmdefaultgroups",
		FetchDDLURL:     "/ddl/realm_default_groups",
	}

	tmp.TableInfo, _ = model.GetTableInfo("realm_default_groups")
	crudEndpoints["realm_default_groups"] = tmp

	tmp = &CrudAPI{
		Name:            "realm_enabled_event_types",
		CreateURL:       "/realmenabledeventtypes",
		RetrieveOneURL:  "/realmenabledeventtypes",
		RetrieveManyURL: "/realmenabledeventtypes",
		UpdateURL:       "/realmenabledeventtypes",
		DeleteURL:       "/realmenabledeventtypes",
		FetchDDLURL:     "/ddl/realm_enabled_event_types",
	}

	tmp.TableInfo, _ = model.GetTableInfo("realm_enabled_event_types")
	crudEndpoints["realm_enabled_event_types"] = tmp

	tmp = &CrudAPI{
		Name:            "realm_events_listeners",
		CreateURL:       "/realmeventslisteners",
		RetrieveOneURL:  "/realmeventslisteners",
		RetrieveManyURL: "/realmeventslisteners",
		UpdateURL:       "/realmeventslisteners",
		DeleteURL:       "/realmeventslisteners",
		FetchDDLURL:     "/ddl/realm_events_listeners",
	}

	tmp.TableInfo, _ = model.GetTableInfo("realm_events_listeners")
	crudEndpoints["realm_events_listeners"] = tmp

	tmp = &CrudAPI{
		Name:            "realm_localizations",
		CreateURL:       "/realmlocalizations",
		RetrieveOneURL:  "/realmlocalizations",
		RetrieveManyURL: "/realmlocalizations",
		UpdateURL:       "/realmlocalizations",
		DeleteURL:       "/realmlocalizations",
		FetchDDLURL:     "/ddl/realm_localizations",
	}

	tmp.TableInfo, _ = model.GetTableInfo("realm_localizations")
	crudEndpoints["realm_localizations"] = tmp

	tmp = &CrudAPI{
		Name:            "realm_required_credential",
		CreateURL:       "/realmrequiredcredential",
		RetrieveOneURL:  "/realmrequiredcredential",
		RetrieveManyURL: "/realmrequiredcredential",
		UpdateURL:       "/realmrequiredcredential",
		DeleteURL:       "/realmrequiredcredential",
		FetchDDLURL:     "/ddl/realm_required_credential",
	}

	tmp.TableInfo, _ = model.GetTableInfo("realm_required_credential")
	crudEndpoints["realm_required_credential"] = tmp

	tmp = &CrudAPI{
		Name:            "realm_smtp_config",
		CreateURL:       "/realmsmtpconfig",
		RetrieveOneURL:  "/realmsmtpconfig",
		RetrieveManyURL: "/realmsmtpconfig",
		UpdateURL:       "/realmsmtpconfig",
		DeleteURL:       "/realmsmtpconfig",
		FetchDDLURL:     "/ddl/realm_smtp_config",
	}

	tmp.TableInfo, _ = model.GetTableInfo("realm_smtp_config")
	crudEndpoints["realm_smtp_config"] = tmp

	tmp = &CrudAPI{
		Name:            "realm_supported_locales",
		CreateURL:       "/realmsupportedlocales",
		RetrieveOneURL:  "/realmsupportedlocales",
		RetrieveManyURL: "/realmsupportedlocales",
		UpdateURL:       "/realmsupportedlocales",
		DeleteURL:       "/realmsupportedlocales",
		FetchDDLURL:     "/ddl/realm_supported_locales",
	}

	tmp.TableInfo, _ = model.GetTableInfo("realm_supported_locales")
	crudEndpoints["realm_supported_locales"] = tmp

	tmp = &CrudAPI{
		Name:            "redirect_uris",
		CreateURL:       "/redirecturis",
		RetrieveOneURL:  "/redirecturis",
		RetrieveManyURL: "/redirecturis",
		UpdateURL:       "/redirecturis",
		DeleteURL:       "/redirecturis",
		FetchDDLURL:     "/ddl/redirect_uris",
	}

	tmp.TableInfo, _ = model.GetTableInfo("redirect_uris")
	crudEndpoints["redirect_uris"] = tmp

	tmp = &CrudAPI{
		Name:            "required_action_config",
		CreateURL:       "/requiredactionconfig",
		RetrieveOneURL:  "/requiredactionconfig",
		RetrieveManyURL: "/requiredactionconfig",
		UpdateURL:       "/requiredactionconfig",
		DeleteURL:       "/requiredactionconfig",
		FetchDDLURL:     "/ddl/required_action_config",
	}

	tmp.TableInfo, _ = model.GetTableInfo("required_action_config")
	crudEndpoints["required_action_config"] = tmp

	tmp = &CrudAPI{
		Name:            "required_action_provider",
		CreateURL:       "/requiredactionprovider",
		RetrieveOneURL:  "/requiredactionprovider",
		RetrieveManyURL: "/requiredactionprovider",
		UpdateURL:       "/requiredactionprovider",
		DeleteURL:       "/requiredactionprovider",
		FetchDDLURL:     "/ddl/required_action_provider",
	}

	tmp.TableInfo, _ = model.GetTableInfo("required_action_provider")
	crudEndpoints["required_action_provider"] = tmp

	tmp = &CrudAPI{
		Name:            "resource_attribute",
		CreateURL:       "/resourceattribute",
		RetrieveOneURL:  "/resourceattribute",
		RetrieveManyURL: "/resourceattribute",
		UpdateURL:       "/resourceattribute",
		DeleteURL:       "/resourceattribute",
		FetchDDLURL:     "/ddl/resource_attribute",
	}

	tmp.TableInfo, _ = model.GetTableInfo("resource_attribute")
	crudEndpoints["resource_attribute"] = tmp

	tmp = &CrudAPI{
		Name:            "resource_policy",
		CreateURL:       "/resourcepolicy",
		RetrieveOneURL:  "/resourcepolicy",
		RetrieveManyURL: "/resourcepolicy",
		UpdateURL:       "/resourcepolicy",
		DeleteURL:       "/resourcepolicy",
		FetchDDLURL:     "/ddl/resource_policy",
	}

	tmp.TableInfo, _ = model.GetTableInfo("resource_policy")
	crudEndpoints["resource_policy"] = tmp

	tmp = &CrudAPI{
		Name:            "resource_scope",
		CreateURL:       "/resourcescope",
		RetrieveOneURL:  "/resourcescope",
		RetrieveManyURL: "/resourcescope",
		UpdateURL:       "/resourcescope",
		DeleteURL:       "/resourcescope",
		FetchDDLURL:     "/ddl/resource_scope",
	}

	tmp.TableInfo, _ = model.GetTableInfo("resource_scope")
	crudEndpoints["resource_scope"] = tmp

	tmp = &CrudAPI{
		Name:            "resource_server",
		CreateURL:       "/resourceserver",
		RetrieveOneURL:  "/resourceserver",
		RetrieveManyURL: "/resourceserver",
		UpdateURL:       "/resourceserver",
		DeleteURL:       "/resourceserver",
		FetchDDLURL:     "/ddl/resource_server",
	}

	tmp.TableInfo, _ = model.GetTableInfo("resource_server")
	crudEndpoints["resource_server"] = tmp

	tmp = &CrudAPI{
		Name:            "resource_server_perm_ticket",
		CreateURL:       "/resourceserverpermticket",
		RetrieveOneURL:  "/resourceserverpermticket",
		RetrieveManyURL: "/resourceserverpermticket",
		UpdateURL:       "/resourceserverpermticket",
		DeleteURL:       "/resourceserverpermticket",
		FetchDDLURL:     "/ddl/resource_server_perm_ticket",
	}

	tmp.TableInfo, _ = model.GetTableInfo("resource_server_perm_ticket")
	crudEndpoints["resource_server_perm_ticket"] = tmp

	tmp = &CrudAPI{
		Name:            "resource_server_policy",
		CreateURL:       "/resourceserverpolicy",
		RetrieveOneURL:  "/resourceserverpolicy",
		RetrieveManyURL: "/resourceserverpolicy",
		UpdateURL:       "/resourceserverpolicy",
		DeleteURL:       "/resourceserverpolicy",
		FetchDDLURL:     "/ddl/resource_server_policy",
	}

	tmp.TableInfo, _ = model.GetTableInfo("resource_server_policy")
	crudEndpoints["resource_server_policy"] = tmp

	tmp = &CrudAPI{
		Name:            "resource_server_resource",
		CreateURL:       "/resourceserverresource",
		RetrieveOneURL:  "/resourceserverresource",
		RetrieveManyURL: "/resourceserverresource",
		UpdateURL:       "/resourceserverresource",
		DeleteURL:       "/resourceserverresource",
		FetchDDLURL:     "/ddl/resource_server_resource",
	}

	tmp.TableInfo, _ = model.GetTableInfo("resource_server_resource")
	crudEndpoints["resource_server_resource"] = tmp

	tmp = &CrudAPI{
		Name:            "resource_server_scope",
		CreateURL:       "/resourceserverscope",
		RetrieveOneURL:  "/resourceserverscope",
		RetrieveManyURL: "/resourceserverscope",
		UpdateURL:       "/resourceserverscope",
		DeleteURL:       "/resourceserverscope",
		FetchDDLURL:     "/ddl/resource_server_scope",
	}

	tmp.TableInfo, _ = model.GetTableInfo("resource_server_scope")
	crudEndpoints["resource_server_scope"] = tmp

	tmp = &CrudAPI{
		Name:            "resource_uris",
		CreateURL:       "/resourceuris",
		RetrieveOneURL:  "/resourceuris",
		RetrieveManyURL: "/resourceuris",
		UpdateURL:       "/resourceuris",
		DeleteURL:       "/resourceuris",
		FetchDDLURL:     "/ddl/resource_uris",
	}

	tmp.TableInfo, _ = model.GetTableInfo("resource_uris")
	crudEndpoints["resource_uris"] = tmp

	tmp = &CrudAPI{
		Name:            "role_attribute",
		CreateURL:       "/roleattribute",
		RetrieveOneURL:  "/roleattribute",
		RetrieveManyURL: "/roleattribute",
		UpdateURL:       "/roleattribute",
		DeleteURL:       "/roleattribute",
		FetchDDLURL:     "/ddl/role_attribute",
	}

	tmp.TableInfo, _ = model.GetTableInfo("role_attribute")
	crudEndpoints["role_attribute"] = tmp

	tmp = &CrudAPI{
		Name:            "scope_mapping",
		CreateURL:       "/scopemapping",
		RetrieveOneURL:  "/scopemapping",
		RetrieveManyURL: "/scopemapping",
		UpdateURL:       "/scopemapping",
		DeleteURL:       "/scopemapping",
		FetchDDLURL:     "/ddl/scope_mapping",
	}

	tmp.TableInfo, _ = model.GetTableInfo("scope_mapping")
	crudEndpoints["scope_mapping"] = tmp

	tmp = &CrudAPI{
		Name:            "scope_policy",
		CreateURL:       "/scopepolicy",
		RetrieveOneURL:  "/scopepolicy",
		RetrieveManyURL: "/scopepolicy",
		UpdateURL:       "/scopepolicy",
		DeleteURL:       "/scopepolicy",
		FetchDDLURL:     "/ddl/scope_policy",
	}

	tmp.TableInfo, _ = model.GetTableInfo("scope_policy")
	crudEndpoints["scope_policy"] = tmp

	tmp = &CrudAPI{
		Name:            "user",
		CreateURL:       "/user",
		RetrieveOneURL:  "/user",
		RetrieveManyURL: "/user",
		UpdateURL:       "/user",
		DeleteURL:       "/user",
		FetchDDLURL:     "/ddl/user",
	}

	tmp.TableInfo, _ = model.GetTableInfo("user")
	crudEndpoints["user"] = tmp

	tmp = &CrudAPI{
		Name:            "user_attribute",
		CreateURL:       "/userattribute",
		RetrieveOneURL:  "/userattribute",
		RetrieveManyURL: "/userattribute",
		UpdateURL:       "/userattribute",
		DeleteURL:       "/userattribute",
		FetchDDLURL:     "/ddl/user_attribute",
	}

	tmp.TableInfo, _ = model.GetTableInfo("user_attribute")
	crudEndpoints["user_attribute"] = tmp

	tmp = &CrudAPI{
		Name:            "user_consent",
		CreateURL:       "/userconsent",
		RetrieveOneURL:  "/userconsent",
		RetrieveManyURL: "/userconsent",
		UpdateURL:       "/userconsent",
		DeleteURL:       "/userconsent",
		FetchDDLURL:     "/ddl/user_consent",
	}

	tmp.TableInfo, _ = model.GetTableInfo("user_consent")
	crudEndpoints["user_consent"] = tmp

	tmp = &CrudAPI{
		Name:            "user_consent_client_scope",
		CreateURL:       "/userconsentclientscope",
		RetrieveOneURL:  "/userconsentclientscope",
		RetrieveManyURL: "/userconsentclientscope",
		UpdateURL:       "/userconsentclientscope",
		DeleteURL:       "/userconsentclientscope",
		FetchDDLURL:     "/ddl/user_consent_client_scope",
	}

	tmp.TableInfo, _ = model.GetTableInfo("user_consent_client_scope")
	crudEndpoints["user_consent_client_scope"] = tmp

	tmp = &CrudAPI{
		Name:            "user_entity",
		CreateURL:       "/userentity",
		RetrieveOneURL:  "/userentity",
		RetrieveManyURL: "/userentity",
		UpdateURL:       "/userentity",
		DeleteURL:       "/userentity",
		FetchDDLURL:     "/ddl/user_entity",
	}

	tmp.TableInfo, _ = model.GetTableInfo("user_entity")
	crudEndpoints["user_entity"] = tmp

	tmp = &CrudAPI{
		Name:            "user_federation_config",
		CreateURL:       "/userfederationconfig",
		RetrieveOneURL:  "/userfederationconfig",
		RetrieveManyURL: "/userfederationconfig",
		UpdateURL:       "/userfederationconfig",
		DeleteURL:       "/userfederationconfig",
		FetchDDLURL:     "/ddl/user_federation_config",
	}

	tmp.TableInfo, _ = model.GetTableInfo("user_federation_config")
	crudEndpoints["user_federation_config"] = tmp

	tmp = &CrudAPI{
		Name:            "user_federation_mapper",
		CreateURL:       "/userfederationmapper",
		RetrieveOneURL:  "/userfederationmapper",
		RetrieveManyURL: "/userfederationmapper",
		UpdateURL:       "/userfederationmapper",
		DeleteURL:       "/userfederationmapper",
		FetchDDLURL:     "/ddl/user_federation_mapper",
	}

	tmp.TableInfo, _ = model.GetTableInfo("user_federation_mapper")
	crudEndpoints["user_federation_mapper"] = tmp

	tmp = &CrudAPI{
		Name:            "user_federation_mapper_config",
		CreateURL:       "/userfederationmapperconfig",
		RetrieveOneURL:  "/userfederationmapperconfig",
		RetrieveManyURL: "/userfederationmapperconfig",
		UpdateURL:       "/userfederationmapperconfig",
		DeleteURL:       "/userfederationmapperconfig",
		FetchDDLURL:     "/ddl/user_federation_mapper_config",
	}

	tmp.TableInfo, _ = model.GetTableInfo("user_federation_mapper_config")
	crudEndpoints["user_federation_mapper_config"] = tmp

	tmp = &CrudAPI{
		Name:            "user_federation_provider",
		CreateURL:       "/userfederationprovider",
		RetrieveOneURL:  "/userfederationprovider",
		RetrieveManyURL: "/userfederationprovider",
		UpdateURL:       "/userfederationprovider",
		DeleteURL:       "/userfederationprovider",
		FetchDDLURL:     "/ddl/user_federation_provider",
	}

	tmp.TableInfo, _ = model.GetTableInfo("user_federation_provider")
	crudEndpoints["user_federation_provider"] = tmp

	tmp = &CrudAPI{
		Name:            "user_group_membership",
		CreateURL:       "/usergroupmembership",
		RetrieveOneURL:  "/usergroupmembership",
		RetrieveManyURL: "/usergroupmembership",
		UpdateURL:       "/usergroupmembership",
		DeleteURL:       "/usergroupmembership",
		FetchDDLURL:     "/ddl/user_group_membership",
	}

	tmp.TableInfo, _ = model.GetTableInfo("user_group_membership")
	crudEndpoints["user_group_membership"] = tmp

	tmp = &CrudAPI{
		Name:            "user_required_action",
		CreateURL:       "/userrequiredaction",
		RetrieveOneURL:  "/userrequiredaction",
		RetrieveManyURL: "/userrequiredaction",
		UpdateURL:       "/userrequiredaction",
		DeleteURL:       "/userrequiredaction",
		FetchDDLURL:     "/ddl/user_required_action",
	}

	tmp.TableInfo, _ = model.GetTableInfo("user_required_action")
	crudEndpoints["user_required_action"] = tmp

	tmp = &CrudAPI{
		Name:            "user_role_mapping",
		CreateURL:       "/userrolemapping",
		RetrieveOneURL:  "/userrolemapping",
		RetrieveManyURL: "/userrolemapping",
		UpdateURL:       "/userrolemapping",
		DeleteURL:       "/userrolemapping",
		FetchDDLURL:     "/ddl/user_role_mapping",
	}

	tmp.TableInfo, _ = model.GetTableInfo("user_role_mapping")
	crudEndpoints["user_role_mapping"] = tmp

	tmp = &CrudAPI{
		Name:            "user_session",
		CreateURL:       "/usersession",
		RetrieveOneURL:  "/usersession",
		RetrieveManyURL: "/usersession",
		UpdateURL:       "/usersession",
		DeleteURL:       "/usersession",
		FetchDDLURL:     "/ddl/user_session",
	}

	tmp.TableInfo, _ = model.GetTableInfo("user_session")
	crudEndpoints["user_session"] = tmp

	tmp = &CrudAPI{
		Name:            "user_session_note",
		CreateURL:       "/usersessionnote",
		RetrieveOneURL:  "/usersessionnote",
		RetrieveManyURL: "/usersessionnote",
		UpdateURL:       "/usersessionnote",
		DeleteURL:       "/usersessionnote",
		FetchDDLURL:     "/ddl/user_session_note",
	}

	tmp.TableInfo, _ = model.GetTableInfo("user_session_note")
	crudEndpoints["user_session_note"] = tmp

	tmp = &CrudAPI{
		Name:            "username_login_failure",
		CreateURL:       "/usernameloginfailure",
		RetrieveOneURL:  "/usernameloginfailure",
		RetrieveManyURL: "/usernameloginfailure",
		UpdateURL:       "/usernameloginfailure",
		DeleteURL:       "/usernameloginfailure",
		FetchDDLURL:     "/ddl/username_login_failure",
	}

	tmp.TableInfo, _ = model.GetTableInfo("username_login_failure")
	crudEndpoints["username_login_failure"] = tmp

	tmp = &CrudAPI{
		Name:            "web_origins",
		CreateURL:       "/weborigins",
		RetrieveOneURL:  "/weborigins",
		RetrieveManyURL: "/weborigins",
		UpdateURL:       "/weborigins",
		DeleteURL:       "/weborigins",
		FetchDDLURL:     "/ddl/web_origins",
	}

	tmp.TableInfo, _ = model.GetTableInfo("web_origins")
	crudEndpoints["web_origins"] = tmp

}
