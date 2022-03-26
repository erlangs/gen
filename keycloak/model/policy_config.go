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


Table: policy_config
[ 0] policy_id                                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] value                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "policy_id": "dLojwwrQgibTfFvZHhbprxnXb",    "name": "NQXvfTwWwCqrqXBhqGcZVqGwG",    "value": "xUilURaCbImOFILuJrbqyBEyS"}



*/

// PolicyConfig struct is a row record of the policy_config table in the keycloak database
type PolicyConfig struct {
	//[ 0] policy_id                                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	PolicyID string `gorm:"primary_key;column:policy_id;type:VARCHAR;size:36;" json:"policy_id"`
	//[ 1] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Name string `gorm:"primary_key;column:name;type:VARCHAR;size:255;" json:"name"`
	//[ 2] value                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Value sql.NullString `gorm:"column:value;type:TEXT;" json:"value"`
}

var policy_configTableInfo = &TableInfo{
	Name: "policy_config",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
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
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *PolicyConfig) TableName() string {
	return "policy_config"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PolicyConfig) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PolicyConfig) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PolicyConfig) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PolicyConfig) TableInfo() *TableInfo {
	return policy_configTableInfo
}
