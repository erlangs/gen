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


Table: identity_provider
[ 0] internal_id                                    VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] enabled                                        BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[ 2] provider_alias                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] provider_id                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] store_token                                    BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[ 5] authenticate_by_default                        BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[ 6] realm_id                                       VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 7] add_token_role                                 BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [true]
[ 8] trust_email                                    BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[ 9] first_broker_login_flow_id                     VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[10] post_broker_login_flow_id                      VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[11] provider_display_name                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[12] link_only                                      BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]


JSON Sample
-------------------------------------
{    "internal_id": "bufABomFBuTnIVLjfCyQeOPaQ",    "enabled": false,    "provider_alias": "GeFfRfltlnXRPPndWMPoMnkCU",    "provider_id": "lqruxhlrpPiOxrXkqgeinGkoU",    "store_token": false,    "authenticate_by_default": false,    "realm_id": "hyOvhMvKsSrbcOxOMNKharPvN",    "add_token_role": false,    "trust_email": true,    "first_broker_login_flow_id": "YqiVMXrxhGfkHfoWTnYoXHASh",    "post_broker_login_flow_id": "ZFGQMmWjfXoElCrWjuyLcbuxE",    "provider_display_name": "QWrZZnSKiaDCcAjNcbHhobdId",    "link_only": false}



*/

// IdentityProvider struct is a row record of the identity_provider table in the keycloak database
type IdentityProvider struct {
	//[ 0] internal_id                                    VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	InternalID string `gorm:"primary_key;column:internal_id;type:VARCHAR;size:36;" json:"internal_id"`
	//[ 1] enabled                                        BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	Enabled bool `gorm:"column:enabled;type:BOOL;default:false;" json:"enabled"`
	//[ 2] provider_alias                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ProviderAlias sql.NullString `gorm:"column:provider_alias;type:VARCHAR;size:255;" json:"provider_alias"`
	//[ 3] provider_id                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ProviderID sql.NullString `gorm:"column:provider_id;type:VARCHAR;size:255;" json:"provider_id"`
	//[ 4] store_token                                    BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	StoreToken bool `gorm:"column:store_token;type:BOOL;default:false;" json:"store_token"`
	//[ 5] authenticate_by_default                        BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	AuthenticateByDefault bool `gorm:"column:authenticate_by_default;type:BOOL;default:false;" json:"authenticate_by_default"`
	//[ 6] realm_id                                       VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID sql.NullString `gorm:"column:realm_id;type:VARCHAR;size:36;" json:"realm_id"`
	//[ 7] add_token_role                                 BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [true]
	AddTokenRole bool `gorm:"column:add_token_role;type:BOOL;default:true;" json:"add_token_role"`
	//[ 8] trust_email                                    BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	TrustEmail bool `gorm:"column:trust_email;type:BOOL;default:false;" json:"trust_email"`
	//[ 9] first_broker_login_flow_id                     VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	FirstBrokerLoginFlowID sql.NullString `gorm:"column:first_broker_login_flow_id;type:VARCHAR;size:36;" json:"first_broker_login_flow_id"`
	//[10] post_broker_login_flow_id                      VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	PostBrokerLoginFlowID sql.NullString `gorm:"column:post_broker_login_flow_id;type:VARCHAR;size:36;" json:"post_broker_login_flow_id"`
	//[11] provider_display_name                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ProviderDisplayName sql.NullString `gorm:"column:provider_display_name;type:VARCHAR;size:255;" json:"provider_display_name"`
	//[12] link_only                                      BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	LinkOnly bool `gorm:"column:link_only;type:BOOL;default:false;" json:"link_only"`
}

var identity_providerTableInfo = &TableInfo{
	Name: "identity_provider",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "internal_id",
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
			GoFieldName:        "InternalID",
			GoFieldType:        "string",
			JSONFieldName:      "internal_id",
			ProtobufFieldName:  "internal_id",
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
			Name:               "provider_alias",
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
			GoFieldName:        "ProviderAlias",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "provider_alias",
			ProtobufFieldName:  "provider_alias",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "provider_id",
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
			GoFieldName:        "ProviderID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "provider_id",
			ProtobufFieldName:  "provider_id",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "store_token",
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
			GoFieldName:        "StoreToken",
			GoFieldType:        "bool",
			JSONFieldName:      "store_token",
			ProtobufFieldName:  "store_token",
			ProtobufType:       "bool",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "authenticate_by_default",
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
			GoFieldName:        "AuthenticateByDefault",
			GoFieldType:        "bool",
			JSONFieldName:      "authenticate_by_default",
			ProtobufFieldName:  "authenticate_by_default",
			ProtobufType:       "bool",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "add_token_role",
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
			GoFieldName:        "AddTokenRole",
			GoFieldType:        "bool",
			JSONFieldName:      "add_token_role",
			ProtobufFieldName:  "add_token_role",
			ProtobufType:       "bool",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "trust_email",
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
			GoFieldName:        "TrustEmail",
			GoFieldType:        "bool",
			JSONFieldName:      "trust_email",
			ProtobufFieldName:  "trust_email",
			ProtobufType:       "bool",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "first_broker_login_flow_id",
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
			GoFieldName:        "FirstBrokerLoginFlowID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "first_broker_login_flow_id",
			ProtobufFieldName:  "first_broker_login_flow_id",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "post_broker_login_flow_id",
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
			GoFieldName:        "PostBrokerLoginFlowID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "post_broker_login_flow_id",
			ProtobufFieldName:  "post_broker_login_flow_id",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "provider_display_name",
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
			GoFieldName:        "ProviderDisplayName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "provider_display_name",
			ProtobufFieldName:  "provider_display_name",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "link_only",
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
			GoFieldName:        "LinkOnly",
			GoFieldType:        "bool",
			JSONFieldName:      "link_only",
			ProtobufFieldName:  "link_only",
			ProtobufType:       "bool",
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (i *IdentityProvider) TableName() string {
	return "identity_provider"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *IdentityProvider) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *IdentityProvider) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *IdentityProvider) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (i *IdentityProvider) TableInfo() *TableInfo {
	return identity_providerTableInfo
}
