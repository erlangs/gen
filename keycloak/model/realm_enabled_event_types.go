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


Table: realm_enabled_event_types
[ 0] realm_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] value                                          VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "realm_id": "ipmBUFwJavFMfNdBeqYBIVUAR",    "value": "VdhPULVCCvHoShsiUhOmFqLEU"}



*/

// RealmEnabledEventTypes struct is a row record of the realm_enabled_event_types table in the keycloak database
type RealmEnabledEventTypes struct {
	//[ 0] realm_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID string `gorm:"primary_key;column:realm_id;type:VARCHAR(36);size:36;" json:"realm_id"`
	//[ 1] value                                          VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Value string `gorm:"primary_key;column:value;type:VARCHAR(255);size:255;" json:"value"`
}

var realm_enabled_event_typesTableInfo = &TableInfo{
	Name: "realm_enabled_event_types",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "realm_id",
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
			GoFieldName:        "RealmID",
			GoFieldType:        "string",
			JSONFieldName:      "realm_id",
			ProtobufFieldName:  "realm_id",
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
func (r *RealmEnabledEventTypes) TableName() string {
	return "realm_enabled_event_types"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RealmEnabledEventTypes) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RealmEnabledEventTypes) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RealmEnabledEventTypes) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *RealmEnabledEventTypes) TableInfo() *TableInfo {
	return realm_enabled_event_typesTableInfo
}
