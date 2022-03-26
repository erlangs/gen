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


Table: resource_policy
[ 0] resource_id                                    VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] policy_id                                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "resource_id": "UDnWkmUmyCJDqgfcnlIDbQFgq",    "policy_id": "WEUQqgCCUdMGiIoUixkfXMNSs"}



*/

// ResourcePolicy struct is a row record of the resource_policy table in the keycloak database
type ResourcePolicy struct {
	//[ 0] resource_id                                    VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ResourceID string `gorm:"primary_key;column:resource_id;type:VARCHAR;size:36;" json:"resource_id"`
	//[ 1] policy_id                                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	PolicyID string `gorm:"primary_key;column:policy_id;type:VARCHAR;size:36;" json:"policy_id"`
}

var resource_policyTableInfo = &TableInfo{
	Name: "resource_policy",
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
			Name:               "policy_id",
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
			GoFieldName:        "PolicyID",
			GoFieldType:        "string",
			JSONFieldName:      "policy_id",
			ProtobufFieldName:  "policy_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *ResourcePolicy) TableName() string {
	return "resource_policy"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *ResourcePolicy) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *ResourcePolicy) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *ResourcePolicy) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *ResourcePolicy) TableInfo() *TableInfo {
	return resource_policyTableInfo
}
