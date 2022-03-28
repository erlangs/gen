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


Table: user_required_action
[ 0] user_id                                        VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] required_action                                VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: [ ]


JSON Sample
-------------------------------------
{    "user_id": "JRZvhHGkBxibPLTPQAXBtogau",    "required_action": "RPSoynwJRFXZqDBLeaNXTuGvn"}



*/

// UserRequiredAction struct is a row record of the user_required_action table in the keycloak database
type UserRequiredAction struct {
	//[ 0] user_id                                        VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	UserID string `gorm:"primary_key;column:user_id;type:VARCHAR(36);size:36;" json:"user_id"`
	//[ 1] required_action                                VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: [ ]
	RequiredAction string `gorm:"primary_key;column:required_action;type:VARCHAR(255);size:255;default: ;" json:"required_action"`
}

var user_required_actionTableInfo = &TableInfo{
	Name: "user_required_action",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "user_id",
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
			GoFieldName:        "UserID",
			GoFieldType:        "string",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "required_action",
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
			GoFieldName:        "RequiredAction",
			GoFieldType:        "string",
			JSONFieldName:      "required_action",
			ProtobufFieldName:  "required_action",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *UserRequiredAction) TableName() string {
	return "user_required_action"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *UserRequiredAction) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *UserRequiredAction) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *UserRequiredAction) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *UserRequiredAction) TableInfo() *TableInfo {
	return user_required_actionTableInfo
}
