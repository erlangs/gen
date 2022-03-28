package model

import (
	"database/sql"
	"time"

	//"github.com/satori/go.uuid"

	"gorm.io/gorm"
)

/*
DB Table Details
-------------------------------------


Table: authentication_execution
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] alias                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] authenticator                                  VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 3] realm_id                                       VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 4] flow_id                                        VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 5] requirement                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 6] priority                                       INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 7] authenticator_flow                             BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[ 8] auth_flow_id                                   VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 9] auth_config                                    VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "id": "kHjcZESkGJcnKMDlLyGJObfFe",    "alias": "CmMuHaSWyGkWiOMkILPeyBSMU",    "authenticator": "CaYxPSvWPvdKFwTHPPMPDZAiC",    "realm_id": "cWPRVmKTKZCXdAIOokPdbENLk",    "flow_id": "epifXlgjJobfAonQrZYknTNJG",    "requirement": 11,    "priority": 45,    "authenticator_flow": false,    "auth_flow_id": "OgRaatruOvfCjGnnrFYtsaODy",    "auth_config": "KiOjMwBwAsKhrZNyQmULBIuFO"}



*/

// AuthenticationExecution struct is a row record of the authentication_execution table in the keycloak database
type AuthenticationExecution struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR(36);size:36;" json:"id"`
	//[ 1] alias                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Alias sql.NullString `gorm:"column:alias;type:VARCHAR(255);size:255;" json:"alias"`
	//[ 2] authenticator                                  VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	Authenticator sql.NullString `gorm:"column:authenticator;type:VARCHAR(36);size:36;" json:"authenticator"`
	//[ 3] realm_id                                       VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID sql.NullString `gorm:"column:realm_id;type:VARCHAR(36);size:36;" json:"realm_id"`
	//[ 4] flow_id                                        VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	FlowID sql.NullString `gorm:"column:flow_id;type:VARCHAR(36);size:36;" json:"flow_id"`
	//[ 5] requirement                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Requirement sql.NullInt32 `gorm:"column:requirement;type:INT4;" json:"requirement"`
	//[ 6] priority                                       INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Priority sql.NullInt32 `gorm:"column:priority;type:INT4;" json:"priority"`
	//[ 7] authenticator_flow                             BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	AuthenticatorFlow bool `gorm:"column:authenticator_flow;type:BOOL;default:false;" json:"authenticator_flow"`
	//[ 8] auth_flow_id                                   VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	AuthFlowID sql.NullString `gorm:"column:auth_flow_id;type:VARCHAR(36);size:36;" json:"auth_flow_id"`
	//[ 9] auth_config                                    VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	AuthConfig sql.NullString `gorm:"column:auth_config;type:VARCHAR(36);size:36;" json:"auth_config"`
}

var authentication_executionTableInfo = &TableInfo{
	Name: "authentication_execution",
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
			Name:               "alias",
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
			GoFieldName:        "Alias",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "alias",
			ProtobufFieldName:  "alias",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "authenticator",
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
			GoFieldName:        "Authenticator",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "authenticator",
			ProtobufFieldName:  "authenticator",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "flow_id",
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
			GoFieldName:        "FlowID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "flow_id",
			ProtobufFieldName:  "flow_id",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "requirement",
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
			GoFieldName:        "Requirement",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "requirement",
			ProtobufFieldName:  "requirement",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "priority",
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
			GoFieldName:        "Priority",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "priority",
			ProtobufFieldName:  "priority",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "authenticator_flow",
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
			GoFieldName:        "AuthenticatorFlow",
			GoFieldType:        "bool",
			JSONFieldName:      "authenticator_flow",
			ProtobufFieldName:  "authenticator_flow",
			ProtobufType:       "bool",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "auth_flow_id",
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
			GoFieldName:        "AuthFlowID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "auth_flow_id",
			ProtobufFieldName:  "auth_flow_id",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "auth_config",
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
			GoFieldName:        "AuthConfig",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "auth_config",
			ProtobufFieldName:  "auth_config",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AuthenticationExecution) TableName() string {
	return "authentication_execution"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AuthenticationExecution) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AuthenticationExecution) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AuthenticationExecution) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AuthenticationExecution) TableInfo() *TableInfo {
	return authentication_executionTableInfo
}
