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


Table: client_scope_attributes
[ 0] scope_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] value                                          VARCHAR(2048)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 2048    default: []
[ 2] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "scope_id": "sZLGcCskroVJIsQsVCAcAuSPq",    "value": "dkgciIqSMCpTCFCnkhPQxPOuy",    "name": "DBQKYfkGOApscfjBGiLsFNGpc"}



*/

// ClientScopeAttributes struct is a row record of the client_scope_attributes table in the keycloak database
type ClientScopeAttributes struct {
	//[ 0] scope_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ScopeID string `gorm:"primary_key;column:scope_id;type:VARCHAR(36);size:36;" json:"scope_id"`
	//[ 1] value                                          VARCHAR(2048)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 2048    default: []
	Value sql.NullString `gorm:"column:value;type:VARCHAR(2048);size:2048;" json:"value"`
	//[ 2] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Name string `gorm:"primary_key;column:name;type:VARCHAR(255);size:255;" json:"name"`
}

var client_scope_attributesTableInfo = &TableInfo{
	Name: "client_scope_attributes",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
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
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "value",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(2048)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       2048,
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
func (c *ClientScopeAttributes) TableName() string {
	return "client_scope_attributes"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ClientScopeAttributes) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ClientScopeAttributes) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ClientScopeAttributes) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ClientScopeAttributes) TableInfo() *TableInfo {
	return client_scope_attributesTableInfo
}
