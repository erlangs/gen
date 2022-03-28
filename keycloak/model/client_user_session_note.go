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


Table: client_user_session_note
[ 0] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] value                                          VARCHAR(2048)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 2048    default: []
[ 2] client_session                                 VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "name": "xHtJatQlmsIgcIrZqqWOUOvfO",    "value": "rhSndQMrbsSwpSUBgCCaZLmRO",    "client_session": "NiPuDHtRJDnserlSAjOUQqnjv"}



*/

// ClientUserSessionNote struct is a row record of the client_user_session_note table in the keycloak database
type ClientUserSessionNote struct {
	//[ 0] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Name string `gorm:"primary_key;column:name;type:VARCHAR(255);size:255;" json:"name"`
	//[ 1] value                                          VARCHAR(2048)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 2048    default: []
	Value sql.NullString `gorm:"column:value;type:VARCHAR(2048);size:2048;" json:"value"`
	//[ 2] client_session                                 VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ClientSession string `gorm:"primary_key;column:client_session;type:VARCHAR(36);size:36;" json:"client_session"`
}

var client_user_session_noteTableInfo = &TableInfo{
	Name: "client_user_session_note",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
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
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "value",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(2048)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       2048,
			GoFieldName:        "Value",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "client_session",
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
			GoFieldName:        "ClientSession",
			GoFieldType:        "string",
			JSONFieldName:      "client_session",
			ProtobufFieldName:  "client_session",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *ClientUserSessionNote) TableName() string {
	return "client_user_session_note"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ClientUserSessionNote) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ClientUserSessionNote) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ClientUserSessionNote) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ClientUserSessionNote) TableInfo() *TableInfo {
	return client_user_session_noteTableInfo
}
