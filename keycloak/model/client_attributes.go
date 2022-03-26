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


Table: client_attributes
[ 0] client_id                                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] value                                          VARCHAR(4000)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 4000    default: []
[ 2] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "client_id": "PanbPUtJqpMNNmaXlnVfWSGPR",    "value": "LCfnLZflMGQoeHQcZMsDkehji",    "name": "cxKnfPqjdfHdaTJhWmZjEqwJa"}



*/

// ClientAttributes struct is a row record of the client_attributes table in the keycloak database
type ClientAttributes struct {
	//[ 0] client_id                                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ClientID string `gorm:"primary_key;column:client_id;type:VARCHAR;size:36;" json:"client_id"`
	//[ 1] value                                          VARCHAR(4000)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 4000    default: []
	Value sql.NullString `gorm:"column:value;type:VARCHAR;size:4000;" json:"value"`
	//[ 2] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Name string `gorm:"primary_key;column:name;type:VARCHAR;size:255;" json:"name"`
}

var client_attributesTableInfo = &TableInfo{
	Name: "client_attributes",
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
			Name:               "value",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(4000)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       4000,
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
func (c *ClientAttributes) TableName() string {
	return "client_attributes"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ClientAttributes) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ClientAttributes) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ClientAttributes) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ClientAttributes) TableInfo() *TableInfo {
	return client_attributesTableInfo
}
