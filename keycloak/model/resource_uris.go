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


Table: resource_uris
[ 0] resource_id                                    VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] value                                          VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "resource_id": "oMhQPFbIQqIlrANNKidRjlJlr",    "value": "pUhIvyLoFTMcLMAGPvQiDAots"}



*/

// ResourceUris struct is a row record of the resource_uris table in the keycloak database
type ResourceUris struct {
	//[ 0] resource_id                                    VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ResourceID string `gorm:"primary_key;column:resource_id;type:VARCHAR(36);size:36;" json:"resource_id"`
	//[ 1] value                                          VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Value string `gorm:"primary_key;column:value;type:VARCHAR(255);size:255;" json:"value"`
}

var resource_urisTableInfo = &TableInfo{
	Name: "resource_uris",
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
			Name:               "value",
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
			GoFieldName:        "Value",
			GoFieldType:        "string",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *ResourceUris) TableName() string {
	return "resource_uris"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *ResourceUris) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *ResourceUris) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *ResourceUris) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *ResourceUris) TableInfo() *TableInfo {
	return resource_urisTableInfo
}
