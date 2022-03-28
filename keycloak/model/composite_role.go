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


Table: composite_role
[ 0] composite                                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] child_role                                     VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "composite": "jMYoumZBSntorKYImMBbyLhRm",    "child_role": "suFYXEHslhJkipLfQPCfsfwAk"}



*/

// CompositeRole struct is a row record of the composite_role table in the keycloak database
type CompositeRole struct {
	//[ 0] composite                                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	Composite string `gorm:"primary_key;column:composite;type:VARCHAR(36);size:36;" json:"composite"`
	//[ 1] child_role                                     VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ChildRole string `gorm:"primary_key;column:child_role;type:VARCHAR(36);size:36;" json:"child_role"`
}

var composite_roleTableInfo = &TableInfo{
	Name: "composite_role",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "composite",
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
			GoFieldName:        "Composite",
			GoFieldType:        "string",
			JSONFieldName:      "composite",
			ProtobufFieldName:  "composite",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "child_role",
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
			GoFieldName:        "ChildRole",
			GoFieldType:        "string",
			JSONFieldName:      "child_role",
			ProtobufFieldName:  "child_role",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CompositeRole) TableName() string {
	return "composite_role"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CompositeRole) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CompositeRole) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CompositeRole) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CompositeRole) TableInfo() *TableInfo {
	return composite_roleTableInfo
}
