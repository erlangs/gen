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


Table: resource_scope
[ 0] resource_id                                    VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] scope_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "resource_id": "QmyHVXwTxKWLhtAIBYdkANvms",    "scope_id": "jfnNNJSHpYVJlgpIsWfaEcgdR"}



*/

// ResourceScope struct is a row record of the resource_scope table in the keycloak database
type ResourceScope struct {
	//[ 0] resource_id                                    VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ResourceID string `gorm:"primary_key;column:resource_id;type:VARCHAR(36);size:36;" json:"resource_id"`
	//[ 1] scope_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ScopeID string `gorm:"primary_key;column:scope_id;type:VARCHAR(36);size:36;" json:"scope_id"`
}

var resource_scopeTableInfo = &TableInfo{
	Name: "resource_scope",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "resource_id",
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
			GoFieldName:        "ResourceID",
			GoFieldType:        "string",
			JSONFieldName:      "resource_id",
			ProtobufFieldName:  "resource_id",
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
func (r *ResourceScope) TableName() string {
	return "resource_scope"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *ResourceScope) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *ResourceScope) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *ResourceScope) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *ResourceScope) TableInfo() *TableInfo {
	return resource_scopeTableInfo
}
