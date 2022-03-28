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


Table: client_node_registrations
[ 0] client_id                                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] value                                          INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "client_id": "ZGXaYyTjYPeqamgxeBRhWsoyN",    "value": 53,    "name": "FgpPXXcXkObKRItuamOmihhgG"}



*/

// ClientNodeRegistrations struct is a row record of the client_node_registrations table in the keycloak database
type ClientNodeRegistrations struct {
	//[ 0] client_id                                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ClientID string `gorm:"primary_key;column:client_id;type:VARCHAR(36);size:36;" json:"client_id"`
	//[ 1] value                                          INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Value sql.NullInt32 `gorm:"column:value;type:INT4;" json:"value"`
	//[ 2] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Name string `gorm:"primary_key;column:name;type:VARCHAR(255);size:255;" json:"name"`
}

var client_node_registrationsTableInfo = &TableInfo{
	Name: "client_node_registrations",
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
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Value",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "int32",
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
func (c *ClientNodeRegistrations) TableName() string {
	return "client_node_registrations"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ClientNodeRegistrations) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ClientNodeRegistrations) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ClientNodeRegistrations) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ClientNodeRegistrations) TableInfo() *TableInfo {
	return client_node_registrationsTableInfo
}
