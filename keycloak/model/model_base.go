package model

import "fmt"

// Action CRUD actions
type Action int32

var (
	// Create action when record is created
	Create = Action(0)

	// RetrieveOne action when a record is retrieved from db
	RetrieveOne = Action(1)

	// RetrieveMany action when record(s) are retrieved from db
	RetrieveMany = Action(2)

	// Update action when record is updated in db
	Update = Action(3)

	// Delete action when record is deleted in db
	Delete = Action(4)

	// FetchDDL action when fetching ddl info from db
	FetchDDL = Action(5)

	tables map[string]*TableInfo
)

func init() {
	tables = make(map[string]*TableInfo)

	tables["admin_event_entity"] = admin_event_entityTableInfo
	tables["associated_policy"] = associated_policyTableInfo
	tables["authentication_execution"] = authentication_executionTableInfo
	tables["authentication_flow"] = authentication_flowTableInfo
	tables["authenticator_config"] = authenticator_configTableInfo
	tables["authenticator_config_entry"] = authenticator_config_entryTableInfo
	tables["broker_link"] = broker_linkTableInfo
	tables["client"] = clientTableInfo
	tables["client_attributes"] = client_attributesTableInfo
	tables["client_auth_flow_bindings"] = client_auth_flow_bindingsTableInfo
	tables["client_initial_access"] = client_initial_accessTableInfo
	tables["client_node_registrations"] = client_node_registrationsTableInfo
	tables["client_scope"] = client_scopeTableInfo
	tables["client_scope_attributes"] = client_scope_attributesTableInfo
	tables["client_scope_client"] = client_scope_clientTableInfo
	tables["client_scope_role_mapping"] = client_scope_role_mappingTableInfo
	tables["client_session"] = client_sessionTableInfo
	tables["client_session_auth_status"] = client_session_auth_statusTableInfo
	tables["client_session_note"] = client_session_noteTableInfo
	tables["client_session_prot_mapper"] = client_session_prot_mapperTableInfo
	tables["client_session_role"] = client_session_roleTableInfo
	tables["client_user_session_note"] = client_user_session_noteTableInfo
	tables["component"] = componentTableInfo
	tables["component_config"] = component_configTableInfo
	tables["composite_role"] = composite_roleTableInfo
	tables["credential"] = credentialTableInfo
	tables["databasechangelog"] = databasechangelogTableInfo
	tables["databasechangeloglock"] = databasechangeloglockTableInfo
	tables["default_client_scope"] = default_client_scopeTableInfo
	tables["event_entity"] = event_entityTableInfo
	tables["fed_user_attribute"] = fed_user_attributeTableInfo
	tables["fed_user_consent"] = fed_user_consentTableInfo
	tables["fed_user_consent_cl_scope"] = fed_user_consent_cl_scopeTableInfo
	tables["fed_user_credential"] = fed_user_credentialTableInfo
	tables["fed_user_group_membership"] = fed_user_group_membershipTableInfo
	tables["fed_user_required_action"] = fed_user_required_actionTableInfo
	tables["fed_user_role_mapping"] = fed_user_role_mappingTableInfo
	tables["federated_identity"] = federated_identityTableInfo
	tables["federated_user"] = federated_userTableInfo
	tables["group_attribute"] = group_attributeTableInfo
	tables["group_role_mapping"] = group_role_mappingTableInfo
	tables["identity_provider"] = identity_providerTableInfo
	tables["identity_provider_config"] = identity_provider_configTableInfo
	tables["identity_provider_mapper"] = identity_provider_mapperTableInfo
	tables["idp_mapper_config"] = idp_mapper_configTableInfo
	tables["keycloak_group"] = keycloak_groupTableInfo
	tables["keycloak_role"] = keycloak_roleTableInfo
	tables["migration_model"] = migration_modelTableInfo
	tables["offline_client_session"] = offline_client_sessionTableInfo
	tables["offline_user_session"] = offline_user_sessionTableInfo
	tables["policy_config"] = policy_configTableInfo
	tables["protocol_mapper"] = protocol_mapperTableInfo
	tables["protocol_mapper_config"] = protocol_mapper_configTableInfo
	tables["realm"] = realmTableInfo
	tables["realm_attribute"] = realm_attributeTableInfo
	tables["realm_default_groups"] = realm_default_groupsTableInfo
	tables["realm_enabled_event_types"] = realm_enabled_event_typesTableInfo
	tables["realm_events_listeners"] = realm_events_listenersTableInfo
	tables["realm_localizations"] = realm_localizationsTableInfo
	tables["realm_required_credential"] = realm_required_credentialTableInfo
	tables["realm_smtp_config"] = realm_smtp_configTableInfo
	tables["realm_supported_locales"] = realm_supported_localesTableInfo
	tables["redirect_uris"] = redirect_urisTableInfo
	tables["required_action_config"] = required_action_configTableInfo
	tables["required_action_provider"] = required_action_providerTableInfo
	tables["resource_attribute"] = resource_attributeTableInfo
	tables["resource_policy"] = resource_policyTableInfo
	tables["resource_scope"] = resource_scopeTableInfo
	tables["resource_server"] = resource_serverTableInfo
	tables["resource_server_perm_ticket"] = resource_server_perm_ticketTableInfo
	tables["resource_server_policy"] = resource_server_policyTableInfo
	tables["resource_server_resource"] = resource_server_resourceTableInfo
	tables["resource_server_scope"] = resource_server_scopeTableInfo
	tables["resource_uris"] = resource_urisTableInfo
	tables["role_attribute"] = role_attributeTableInfo
	tables["scope_mapping"] = scope_mappingTableInfo
	tables["scope_policy"] = scope_policyTableInfo
	tables["user_attribute"] = user_attributeTableInfo
	tables["user_consent"] = user_consentTableInfo
	tables["user_consent_client_scope"] = user_consent_client_scopeTableInfo
	tables["user_entity"] = user_entityTableInfo
	tables["user_federation_config"] = user_federation_configTableInfo
	tables["user_federation_mapper"] = user_federation_mapperTableInfo
	tables["user_federation_mapper_config"] = user_federation_mapper_configTableInfo
	tables["user_federation_provider"] = user_federation_providerTableInfo
	tables["user_group_membership"] = user_group_membershipTableInfo
	tables["user_required_action"] = user_required_actionTableInfo
	tables["user_role_mapping"] = user_role_mappingTableInfo
	tables["user_session"] = user_sessionTableInfo
	tables["user_session_note"] = user_session_noteTableInfo
	tables["username_login_failure"] = username_login_failureTableInfo
	tables["web_origins"] = web_originsTableInfo
}

// String describe the action
func (i Action) String() string {
	switch i {
	case Create:
		return "Create"
	case RetrieveOne:
		return "RetrieveOne"
	case RetrieveMany:
		return "RetrieveMany"
	case Update:
		return "Update"
	case Delete:
		return "Delete"
	case FetchDDL:
		return "FetchDDL"
	default:
		return fmt.Sprintf("unknown action: %d", int(i))
	}
}

// Model interface methods for database structs generated
type Model interface {
	TableName() string
	BeforeSave() error
	Prepare()
	Validate(action Action) error
	TableInfo() *TableInfo
}

// TableInfo describes a table in the database
type TableInfo struct {
	Name    string        `json:"name"`
	Columns []*ColumnInfo `json:"columns"`
}

// ColumnInfo describes a column in the database table
type ColumnInfo struct {
	Index              int    `json:"index"`
	GoFieldName        string `json:"go_field_name"`
	GoFieldType        string `json:"go_field_type"`
	JSONFieldName      string `json:"json_field_name"`
	ProtobufFieldName  string `json:"protobuf_field_name"`
	ProtobufType       string `json:"protobuf_field_type"`
	ProtobufPos        int    `json:"protobuf_field_pos"`
	Comment            string `json:"comment"`
	Notes              string `json:"notes"`
	Name               string `json:"name"`
	Nullable           bool   `json:"is_nullable"`
	DatabaseTypeName   string `json:"database_type_name"`
	DatabaseTypePretty string `json:"database_type_pretty"`
	IsPrimaryKey       bool   `json:"is_primary_key"`
	IsAutoIncrement    bool   `json:"is_auto_increment"`
	IsArray            bool   `json:"is_array"`
	ColumnType         string `json:"column_type"`
	ColumnLength       int64  `json:"column_length"`
	DefaultValue       string `json:"default_value"`
}

// GetTableInfo retrieve TableInfo for a table
func GetTableInfo(name string) (*TableInfo, bool) {
	val, ok := tables[name]
	return val, ok
}
