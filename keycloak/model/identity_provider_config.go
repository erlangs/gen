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


Table: identity_provider_config
[ 0] identity_provider_id                           VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] value                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "identity_provider_id": "kxuaBWvDbonXiNdsIiaZomQDQ",    "value": "YBynjFCTWYFxOGVAfXsFHbNkr",    "name": "hxlWqbtvoyLkMNKvjrOvuGcAR"}



*/

// IdentityProviderConfig struct is a row record of the identity_provider_config table in the keycloak database
type IdentityProviderConfig struct {
	//[ 0] identity_provider_id                           VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	IdentityProviderID string `gorm:"primary_key;column:identity_provider_id;type:VARCHAR;size:36;" json:"identity_provider_id"`
	//[ 1] value                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Value sql.NullString `gorm:"column:value;type:TEXT;" json:"value"`
	//[ 2] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Name string `gorm:"primary_key;column:name;type:VARCHAR;size:255;" json:"name"`
}

var identity_provider_configTableInfo = &TableInfo{
	Name: "identity_provider_config",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "identity_provider_id",
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
			GoFieldName:        "IdentityProviderID",
			GoFieldType:        "string",
			JSONFieldName:      "identity_provider_id",
			ProtobufFieldName:  "identity_provider_id",
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
func (i *IdentityProviderConfig) TableName() string {
	return "identity_provider_config"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *IdentityProviderConfig) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *IdentityProviderConfig) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *IdentityProviderConfig) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (i *IdentityProviderConfig) TableInfo() *TableInfo {
	return identity_provider_configTableInfo
}
