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


Table: client_auth_flow_bindings
[ 0] client_id                                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] flow_id                                        VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 2] binding_name                                   VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "client_id": "NyyVthyeGjRuEBfUwVUrZVNSf",    "flow_id": "dYFwETVYwnHxLmSFfwmLqOnbR",    "binding_name": "XcEBRgyuYqCRBHMOapQAvVMJT"}



*/

// ClientAuthFlowBindings struct is a row record of the client_auth_flow_bindings table in the keycloak database
type ClientAuthFlowBindings struct {
	//[ 0] client_id                                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ClientID string `gorm:"primary_key;column:client_id;type:VARCHAR(36);size:36;" json:"client_id"`
	//[ 1] flow_id                                        VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	FlowID sql.NullString `gorm:"column:flow_id;type:VARCHAR(36);size:36;" json:"flow_id"`
	//[ 2] binding_name                                   VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	BindingName string `gorm:"primary_key;column:binding_name;type:VARCHAR(255);size:255;" json:"binding_name"`
}

var client_auth_flow_bindingsTableInfo = &TableInfo{
	Name: "client_auth_flow_bindings",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "client_id",
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
			GoFieldName:        "ClientID",
			GoFieldType:        "string",
			JSONFieldName:      "client_id",
			ProtobufFieldName:  "client_id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "flow_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(36)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       36,
			GoFieldName:        "FlowID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "flow_id",
			ProtobufFieldName:  "flow_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "binding_name",
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
			GoFieldName:        "BindingName",
			GoFieldType:        "string",
			JSONFieldName:      "binding_name",
			ProtobufFieldName:  "binding_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *ClientAuthFlowBindings) TableName() string {
	return "client_auth_flow_bindings"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ClientAuthFlowBindings) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ClientAuthFlowBindings) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ClientAuthFlowBindings) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ClientAuthFlowBindings) TableInfo() *TableInfo {
	return client_auth_flow_bindingsTableInfo
}
