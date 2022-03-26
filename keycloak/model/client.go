package model

import (
	"database/sql"
	//"time"

	//"github.com/satori/go.uuid"

	"gorm.io/gorm"
)

/*
DB Table Details
-------------------------------------


Table: client
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] enabled                                        BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[ 2] full_scope_allowed                             BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[ 3] client_id                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] not_before                                     INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 5] public_client                                  BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[ 6] secret                                         VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 7] base_url                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] bearer_only                                    BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[ 9] management_url                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[10] surrogate_auth_required                        BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[11] realm_id                                       VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[12] protocol                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[13] node_rereg_timeout                             INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[14] frontchannel_logout                            BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[15] consent_required                               BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[16] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[17] service_accounts_enabled                       BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[18] client_authenticator_type                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[19] root_url                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[20] description                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[21] registration_token                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[22] standard_flow_enabled                          BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [true]
[23] implicit_flow_enabled                          BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[24] direct_access_grants_enabled                   BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[25] always_display_in_console                      BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]


JSON Sample
-------------------------------------
{    "id": "RUxjlQfcoClRhwbkXSKtvWkWI",    "enabled": false,    "full_scope_allowed": false,    "client_id": "jAvRuieDwpArhZNtdDURaEaEV",    "not_before": 66,    "public_client": false,    "secret": "KKrByHxiKkyxnMoidAkTlkMQw",    "base_url": "lSrOboSkyHCNNxwadOopZETOv",    "bearer_only": false,    "management_url": "eEgbrZPHTaFqIXIuOxohNAclj",    "surrogate_auth_required": true,    "realm_id": "AHqtlbIjnYLXVYTtAOmMDOqhZ",    "protocol": "jSeAoJFsSFVWLcZTsUSXUTLYn",    "node_rereg_timeout": 34,    "frontchannel_logout": true,    "consent_required": true,    "name": "YacSGFapuyCZgdlREIMEYcDJT",    "service_accounts_enabled": true,    "client_authenticator_type": "fkcflHvQSGLFSPTqBWvmpPDWy",    "root_url": "mUCqSTPACCuuxCadjtaKjJonj",    "description": "nGXqTbXOtqdoBLSEgmVIsyepQ",    "registration_token": "mECHvYkoBgYZhTHWrBaKrYurW",    "standard_flow_enabled": true,    "implicit_flow_enabled": false,    "direct_access_grants_enabled": false,    "always_display_in_console": true}



*/

// Client struct is a row record of the client table in the keycloak database
type Client struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR;size:36;" json:"id"`
	//[ 1] enabled                                        BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	Enabled bool `gorm:"column:enabled;type:BOOL;default:false;" json:"enabled"`
	//[ 2] full_scope_allowed                             BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	FullScopeAllowed bool `gorm:"column:full_scope_allowed;type:BOOL;default:false;" json:"full_scope_allowed"`
	//[ 3] client_id                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ClientID sql.NullString `gorm:"column:client_id;type:VARCHAR;size:255;" json:"client_id"`
	//[ 4] not_before                                     INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	NotBefore sql.NullInt32 `gorm:"column:not_before;type:INT4;" json:"not_before"`
	//[ 5] public_client                                  BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	PublicClient bool `gorm:"column:public_client;type:BOOL;default:false;" json:"public_client"`
	//[ 6] secret                                         VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Secret sql.NullString `gorm:"column:secret;type:VARCHAR;size:255;" json:"secret"`
	//[ 7] base_url                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	BaseURL sql.NullString `gorm:"column:base_url;type:VARCHAR;size:255;" json:"base_url"`
	//[ 8] bearer_only                                    BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	BearerOnly bool `gorm:"column:bearer_only;type:BOOL;default:false;" json:"bearer_only"`
	//[ 9] management_url                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ManagementURL sql.NullString `gorm:"column:management_url;type:VARCHAR;size:255;" json:"management_url"`
	//[10] surrogate_auth_required                        BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	SurrogateAuthRequired bool `gorm:"column:surrogate_auth_required;type:BOOL;default:false;" json:"surrogate_auth_required"`
	//[11] realm_id                                       VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID sql.NullString `gorm:"column:realm_id;type:VARCHAR;size:36;" json:"realm_id"`
	//[12] protocol                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Protocol sql.NullString `gorm:"column:protocol;type:VARCHAR;size:255;" json:"protocol"`
	//[13] node_rereg_timeout                             INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
	NodeReregTimeout sql.NullInt32 `gorm:"column:node_rereg_timeout;type:INT4;default:0;" json:"node_rereg_timeout"`
	//[14] frontchannel_logout                            BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	FrontchannelLogout bool `gorm:"column:frontchannel_logout;type:BOOL;default:false;" json:"frontchannel_logout"`
	//[15] consent_required                               BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	ConsentRequired bool `gorm:"column:consent_required;type:BOOL;default:false;" json:"consent_required"`
	//[16] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Name sql.NullString `gorm:"column:name;type:VARCHAR;size:255;" json:"name"`
	//[17] service_accounts_enabled                       BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	ServiceAccountsEnabled bool `gorm:"column:service_accounts_enabled;type:BOOL;default:false;" json:"service_accounts_enabled"`
	//[18] client_authenticator_type                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ClientAuthenticatorType sql.NullString `gorm:"column:client_authenticator_type;type:VARCHAR;size:255;" json:"client_authenticator_type"`
	//[19] root_url                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	RootURL sql.NullString `gorm:"column:root_url;type:VARCHAR;size:255;" json:"root_url"`
	//[20] description                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Description sql.NullString `gorm:"column:description;type:VARCHAR;size:255;" json:"description"`
	//[21] registration_token                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	RegistrationToken sql.NullString `gorm:"column:registration_token;type:VARCHAR;size:255;" json:"registration_token"`
	//[22] standard_flow_enabled                          BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [true]
	StandardFlowEnabled bool `gorm:"column:standard_flow_enabled;type:BOOL;default:true;" json:"standard_flow_enabled"`
	//[23] implicit_flow_enabled                          BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	ImplicitFlowEnabled bool `gorm:"column:implicit_flow_enabled;type:BOOL;default:false;" json:"implicit_flow_enabled"`
	//[24] direct_access_grants_enabled                   BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	DirectAccessGrantsEnabled bool `gorm:"column:direct_access_grants_enabled;type:BOOL;default:false;" json:"direct_access_grants_enabled"`
	//[25] always_display_in_console                      BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	AlwaysDisplayInConsole bool `gorm:"column:always_display_in_console;type:BOOL;default:false;" json:"always_display_in_console"`
}

var clientTableInfo = &TableInfo{
	Name: "client",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(36)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       36,
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "enabled",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "Enabled",
			GoFieldType:        "bool",
			JSONFieldName:      "enabled",
			ProtobufFieldName:  "enabled",
			ProtobufType:       "bool",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "full_scope_allowed",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "FullScopeAllowed",
			GoFieldType:        "bool",
			JSONFieldName:      "full_scope_allowed",
			ProtobufFieldName:  "full_scope_allowed",
			ProtobufType:       "bool",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "client_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ClientID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "client_id",
			ProtobufFieldName:  "client_id",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "not_before",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "NotBefore",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "not_before",
			ProtobufFieldName:  "not_before",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "public_client",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "PublicClient",
			GoFieldType:        "bool",
			JSONFieldName:      "public_client",
			ProtobufFieldName:  "public_client",
			ProtobufType:       "bool",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "secret",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Secret",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "secret",
			ProtobufFieldName:  "secret",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "base_url",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "BaseURL",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "base_url",
			ProtobufFieldName:  "base_url",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "bearer_only",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "BearerOnly",
			GoFieldType:        "bool",
			JSONFieldName:      "bearer_only",
			ProtobufFieldName:  "bearer_only",
			ProtobufType:       "bool",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "management_url",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ManagementURL",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "management_url",
			ProtobufFieldName:  "management_url",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "surrogate_auth_required",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "SurrogateAuthRequired",
			GoFieldType:        "bool",
			JSONFieldName:      "surrogate_auth_required",
			ProtobufFieldName:  "surrogate_auth_required",
			ProtobufType:       "bool",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "realm_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(36)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       36,
			GoFieldName:        "RealmID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "realm_id",
			ProtobufFieldName:  "realm_id",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "protocol",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Protocol",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "protocol",
			ProtobufFieldName:  "protocol",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "node_rereg_timeout",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "NodeReregTimeout",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "node_rereg_timeout",
			ProtobufFieldName:  "node_rereg_timeout",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "frontchannel_logout",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "FrontchannelLogout",
			GoFieldType:        "bool",
			JSONFieldName:      "frontchannel_logout",
			ProtobufFieldName:  "frontchannel_logout",
			ProtobufType:       "bool",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "consent_required",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "ConsentRequired",
			GoFieldType:        "bool",
			JSONFieldName:      "consent_required",
			ProtobufFieldName:  "consent_required",
			ProtobufType:       "bool",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "service_accounts_enabled",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "ServiceAccountsEnabled",
			GoFieldType:        "bool",
			JSONFieldName:      "service_accounts_enabled",
			ProtobufFieldName:  "service_accounts_enabled",
			ProtobufType:       "bool",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "client_authenticator_type",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ClientAuthenticatorType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "client_authenticator_type",
			ProtobufFieldName:  "client_authenticator_type",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "root_url",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "RootURL",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "root_url",
			ProtobufFieldName:  "root_url",
			ProtobufType:       "string",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "description",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Description",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "description",
			ProtobufFieldName:  "description",
			ProtobufType:       "string",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "registration_token",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "RegistrationToken",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "registration_token",
			ProtobufFieldName:  "registration_token",
			ProtobufType:       "string",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "standard_flow_enabled",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "StandardFlowEnabled",
			GoFieldType:        "bool",
			JSONFieldName:      "standard_flow_enabled",
			ProtobufFieldName:  "standard_flow_enabled",
			ProtobufType:       "bool",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "implicit_flow_enabled",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "ImplicitFlowEnabled",
			GoFieldType:        "bool",
			JSONFieldName:      "implicit_flow_enabled",
			ProtobufFieldName:  "implicit_flow_enabled",
			ProtobufType:       "bool",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "direct_access_grants_enabled",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "DirectAccessGrantsEnabled",
			GoFieldType:        "bool",
			JSONFieldName:      "direct_access_grants_enabled",
			ProtobufFieldName:  "direct_access_grants_enabled",
			ProtobufType:       "bool",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "always_display_in_console",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "AlwaysDisplayInConsole",
			GoFieldType:        "bool",
			JSONFieldName:      "always_display_in_console",
			ProtobufFieldName:  "always_display_in_console",
			ProtobufType:       "bool",
			ProtobufPos:        26,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *Client) TableName() string {
	return "client"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *Client) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *Client) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *Client) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *Client) TableInfo() *TableInfo {
	return clientTableInfo
}
