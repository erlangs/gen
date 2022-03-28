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


Table: client_session_role
[ 0] role_id                                        VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] client_session                                 VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "role_id": "RKWIQEKubvRlaBIOByTEWWMeJ",    "client_session": "jrkibwgrXDVGDKxROWAGwGGWw"}



*/

// ClientSessionRole struct is a row record of the client_session_role table in the keycloak database
type ClientSessionRole struct {
	//[ 0] role_id                                        VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	RoleID string `gorm:"primary_key;column:role_id;type:VARCHAR(255);size:255;" json:"role_id"`
	//[ 1] client_session                                 VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ClientSession string `gorm:"primary_key;column:client_session;type:VARCHAR(36);size:36;" json:"client_session"`
}

var client_session_roleTableInfo = &TableInfo{
	Name: "client_session_role",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "role_id",
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
			GoFieldName:        "RoleID",
			GoFieldType:        "string",
			JSONFieldName:      "role_id",
			ProtobufFieldName:  "role_id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *ClientSessionRole) TableName() string {
	return "client_session_role"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ClientSessionRole) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ClientSessionRole) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ClientSessionRole) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ClientSessionRole) TableInfo() *TableInfo {
	return client_session_roleTableInfo
}
