package model

import (
	//"database/sql"
	//"time"

	//"github.com/satori/go.uuid"

	"gorm.io/gorm"
)

/*
DB Table Details
-------------------------------------


Table: client_scope_role_mapping
[ 0] scope_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] role_id                                        VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "scope_id": "dGNfMPLdeodLifuoSQODVmNIh",    "role_id": "QhwNnjxaKSLMfKGFihWdrctRN"}



*/

// ClientScopeRoleMapping struct is a row record of the client_scope_role_mapping table in the keycloak database
type ClientScopeRoleMapping struct {
	//[ 0] scope_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ScopeID string `gorm:"primary_key;column:scope_id;type:VARCHAR(36);size:36;" json:"scope_id"`
	//[ 1] role_id                                        VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RoleID string `gorm:"primary_key;column:role_id;type:VARCHAR(36);size:36;" json:"role_id"`
}

var client_scope_role_mappingTableInfo = &TableInfo{
	Name: "client_scope_role_mapping",
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
			Name:               "role_id",
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
			GoFieldName:        "RoleID",
			GoFieldType:        "string",
			JSONFieldName:      "role_id",
			ProtobufFieldName:  "role_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *ClientScopeRoleMapping) TableName() string {
	return "client_scope_role_mapping"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ClientScopeRoleMapping) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ClientScopeRoleMapping) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ClientScopeRoleMapping) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ClientScopeRoleMapping) TableInfo() *TableInfo {
	return client_scope_role_mappingTableInfo
}
