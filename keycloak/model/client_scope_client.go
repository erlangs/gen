package model

import (
	//"database/sql"
	//"time"

	//"github.com/satori/go.uuid"

	"gorm.io/gorm"
)

/*
DB Table Details
-------------------------------------


Table: client_scope_client
[ 0] client_id                                      VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] scope_id                                       VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] default_scope                                  BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]


JSON Sample
-------------------------------------
{    "client_id": "quHAJYEuritytwWkPGxQdmFMX",    "scope_id": "sTdMvLSlUWdeAkndufgjiZakO",    "default_scope": false}



*/

// ClientScopeClient struct is a row record of the client_scope_client table in the keycloak database
type ClientScopeClient struct {
	//[ 0] client_id                                      VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ClientID string `gorm:"primary_key;column:client_id;type:VARCHAR;size:255;" json:"client_id"`
	//[ 1] scope_id                                       VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ScopeID string `gorm:"primary_key;column:scope_id;type:VARCHAR;size:255;" json:"scope_id"`
	//[ 2] default_scope                                  BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	DefaultScope bool `gorm:"column:default_scope;type:BOOL;default:false;" json:"default_scope"`
}

var client_scope_clientTableInfo = &TableInfo{
	Name: "client_scope_client",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "client_id",
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
			GoFieldName:        "ClientID",
			GoFieldType:        "string",
			JSONFieldName:      "client_id",
			ProtobufFieldName:  "client_id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "scope_id",
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
			GoFieldName:        "ScopeID",
			GoFieldType:        "string",
			JSONFieldName:      "scope_id",
			ProtobufFieldName:  "scope_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "default_scope",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "DefaultScope",
			GoFieldType:        "bool",
			JSONFieldName:      "default_scope",
			ProtobufFieldName:  "default_scope",
			ProtobufType:       "bool",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *ClientScopeClient) TableName() string {
	return "client_scope_client"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ClientScopeClient) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ClientScopeClient) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ClientScopeClient) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ClientScopeClient) TableInfo() *TableInfo {
	return client_scope_clientTableInfo
}
