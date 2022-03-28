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


Table: web_origins
[ 0] client_id                                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] value                                          VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "client_id": "VicweBNLWfQYQNEvtdhTvidki",    "value": "EKidYXqcXWDWYmmtQqAMkdPth"}



*/

// WebOrigins struct is a row record of the web_origins table in the keycloak database
type WebOrigins struct {
	//[ 0] client_id                                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ClientID string `gorm:"primary_key;column:client_id;type:VARCHAR(36);size:36;" json:"client_id"`
	//[ 1] value                                          VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Value string `gorm:"primary_key;column:value;type:VARCHAR(255);size:255;" json:"value"`
}

var web_originsTableInfo = &TableInfo{
	Name: "web_origins",
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
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Value",
			GoFieldType:        "string",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (w *WebOrigins) TableName() string {
	return "web_origins"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (w *WebOrigins) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WebOrigins) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WebOrigins) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *WebOrigins) TableInfo() *TableInfo {
	return web_originsTableInfo
}
