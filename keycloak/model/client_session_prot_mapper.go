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


Table: client_session_prot_mapper
[ 0] protocol_mapper_id                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] client_session                                 VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "protocol_mapper_id": "GYIsBvaYKXNiSdrcQOxwadfiK",    "client_session": "LhYoGBJMsowcGOZlYEOABQwto"}



*/

// ClientSessionProtMapper struct is a row record of the client_session_prot_mapper table in the keycloak database
type ClientSessionProtMapper struct {
	//[ 0] protocol_mapper_id                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ProtocolMapperID string `gorm:"primary_key;column:protocol_mapper_id;type:VARCHAR;size:36;" json:"protocol_mapper_id"`
	//[ 1] client_session                                 VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ClientSession string `gorm:"primary_key;column:client_session;type:VARCHAR;size:36;" json:"client_session"`
}

var client_session_prot_mapperTableInfo = &TableInfo{
	Name: "client_session_prot_mapper",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "protocol_mapper_id",
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
			GoFieldName:        "ProtocolMapperID",
			GoFieldType:        "string",
			JSONFieldName:      "protocol_mapper_id",
			ProtobufFieldName:  "protocol_mapper_id",
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
func (c *ClientSessionProtMapper) TableName() string {
	return "client_session_prot_mapper"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ClientSessionProtMapper) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ClientSessionProtMapper) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ClientSessionProtMapper) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ClientSessionProtMapper) TableInfo() *TableInfo {
	return client_session_prot_mapperTableInfo
}
