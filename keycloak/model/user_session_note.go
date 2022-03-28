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


Table: user_session_note
[ 0] user_session                                   VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] value                                          VARCHAR(2048)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 2048    default: []


JSON Sample
-------------------------------------
{    "user_session": "lpaKxqSkXJfnFbiTXOyxBamVn",    "name": "ULxKOQNdVvRPiEJYtBooYuWuO",    "value": "MtRPipVxesCQMTqbhTyHtSvKW"}



*/

// UserSessionNote struct is a row record of the user_session_note table in the keycloak database
type UserSessionNote struct {
	//[ 0] user_session                                   VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	UserSession string `gorm:"primary_key;column:user_session;type:VARCHAR(36);size:36;" json:"user_session"`
	//[ 1] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Name string `gorm:"primary_key;column:name;type:VARCHAR(255);size:255;" json:"name"`
	//[ 2] value                                          VARCHAR(2048)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 2048    default: []
	Value sql.NullString `gorm:"column:value;type:VARCHAR(2048);size:2048;" json:"value"`
}

var user_session_noteTableInfo = &TableInfo{
	Name: "user_session_note",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "user_session",
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
			GoFieldName:        "UserSession",
			GoFieldType:        "string",
			JSONFieldName:      "user_session",
			ProtobufFieldName:  "user_session",
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
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *UserSessionNote) TableName() string {
	return "user_session_note"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *UserSessionNote) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *UserSessionNote) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *UserSessionNote) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *UserSessionNote) TableInfo() *TableInfo {
	return user_session_noteTableInfo
}
