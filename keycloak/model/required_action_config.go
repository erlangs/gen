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


Table: required_action_config
[ 0] required_action_id                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] value                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "required_action_id": "aswSkwkLcwSXDadsLBQkqeqQM",    "value": "tJQeNAEiUaPhPNuEkcTcMpedT",    "name": "uKWbKRKQChwUeAnCXXnlUGaCB"}



*/

// RequiredActionConfig struct is a row record of the required_action_config table in the keycloak database
type RequiredActionConfig struct {
	//[ 0] required_action_id                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RequiredActionID string `gorm:"primary_key;column:required_action_id;type:VARCHAR(36);size:36;" json:"required_action_id"`
	//[ 1] value                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Value sql.NullString `gorm:"column:value;type:TEXT;" json:"value"`
	//[ 2] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Name string `gorm:"primary_key;column:name;type:VARCHAR(255);size:255;" json:"name"`
}

var required_action_configTableInfo = &TableInfo{
	Name: "required_action_config",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "required_action_id",
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
			GoFieldName:        "RequiredActionID",
			GoFieldType:        "string",
			JSONFieldName:      "required_action_id",
			ProtobufFieldName:  "required_action_id",
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
func (r *RequiredActionConfig) TableName() string {
	return "required_action_config"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RequiredActionConfig) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RequiredActionConfig) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RequiredActionConfig) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *RequiredActionConfig) TableInfo() *TableInfo {
	return required_action_configTableInfo
}
