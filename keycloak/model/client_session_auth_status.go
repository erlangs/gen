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


Table: client_session_auth_status
[ 0] authenticator                                  VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] client_session                                 VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "authenticator": "giAqhgYBQydghwUwTUTpgtaFQ",    "status": 58,    "client_session": "GRCXOVleodNJFtTJpNprSTbVn"}



*/

// ClientSessionAuthStatus struct is a row record of the client_session_auth_status table in the keycloak database
type ClientSessionAuthStatus struct {
	//[ 0] authenticator                                  VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	Authenticator string `gorm:"primary_key;column:authenticator;type:VARCHAR;size:36;" json:"authenticator"`
	//[ 1] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Status sql.NullInt32 `gorm:"column:status;type:INT4;" json:"status"`
	//[ 2] client_session                                 VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ClientSession string `gorm:"primary_key;column:client_session;type:VARCHAR;size:36;" json:"client_session"`
}

var client_session_auth_statusTableInfo = &TableInfo{
	Name: "client_session_auth_status",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "authenticator",
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
			GoFieldName:        "Authenticator",
			GoFieldType:        "string",
			JSONFieldName:      "authenticator",
			ProtobufFieldName:  "authenticator",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "status",
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
			GoFieldName:        "Status",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
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
func (c *ClientSessionAuthStatus) TableName() string {
	return "client_session_auth_status"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ClientSessionAuthStatus) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ClientSessionAuthStatus) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ClientSessionAuthStatus) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ClientSessionAuthStatus) TableInfo() *TableInfo {
	return client_session_auth_statusTableInfo
}
