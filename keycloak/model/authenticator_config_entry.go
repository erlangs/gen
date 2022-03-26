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


Table: authenticator_config_entry
[ 0] authenticator_id                               VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] value                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "authenticator_id": "vVfrCOkMSIHEnVgKGvgXqrIuS",    "value": "KfTWkQHrhKijNTuKtpqajbkjw",    "name": "gxFIIHlunlSvivQDgBJHmQojH"}



*/

// AuthenticatorConfigEntry struct is a row record of the authenticator_config_entry table in the keycloak database
type AuthenticatorConfigEntry struct {
	//[ 0] authenticator_id                               VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	AuthenticatorID string `gorm:"primary_key;column:authenticator_id;type:VARCHAR;size:36;" json:"authenticator_id"`
	//[ 1] value                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Value sql.NullString `gorm:"column:value;type:TEXT;" json:"value"`
	//[ 2] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Name string `gorm:"primary_key;column:name;type:VARCHAR;size:255;" json:"name"`
}

var authenticator_config_entryTableInfo = &TableInfo{
	Name: "authenticator_config_entry",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "authenticator_id",
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
			GoFieldName:        "AuthenticatorID",
			GoFieldType:        "string",
			JSONFieldName:      "authenticator_id",
			ProtobufFieldName:  "authenticator_id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "value",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Value",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AuthenticatorConfigEntry) TableName() string {
	return "authenticator_config_entry"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AuthenticatorConfigEntry) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AuthenticatorConfigEntry) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AuthenticatorConfigEntry) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AuthenticatorConfigEntry) TableInfo() *TableInfo {
	return authenticator_config_entryTableInfo
}
