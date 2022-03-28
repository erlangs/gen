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


Table: user_consent_client_scope
[ 0] user_consent_id                                VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] scope_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "user_consent_id": "NDmTGUHGWPXDRhkrvqycvCSyy",    "scope_id": "hojjxUVWyZStIOxXqThwGLVIY"}



*/

// UserConsentClientScope struct is a row record of the user_consent_client_scope table in the keycloak database
type UserConsentClientScope struct {
	//[ 0] user_consent_id                                VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	UserConsentID string `gorm:"primary_key;column:user_consent_id;type:VARCHAR(36);size:36;" json:"user_consent_id"`
	//[ 1] scope_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ScopeID string `gorm:"primary_key;column:scope_id;type:VARCHAR(36);size:36;" json:"scope_id"`
}

var user_consent_client_scopeTableInfo = &TableInfo{
	Name: "user_consent_client_scope",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "user_consent_id",
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
			GoFieldName:        "UserConsentID",
			GoFieldType:        "string",
			JSONFieldName:      "user_consent_id",
			ProtobufFieldName:  "user_consent_id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "scope_id",
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
			GoFieldName:        "ScopeID",
			GoFieldType:        "string",
			JSONFieldName:      "scope_id",
			ProtobufFieldName:  "scope_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *UserConsentClientScope) TableName() string {
	return "user_consent_client_scope"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *UserConsentClientScope) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *UserConsentClientScope) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *UserConsentClientScope) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *UserConsentClientScope) TableInfo() *TableInfo {
	return user_consent_client_scopeTableInfo
}
