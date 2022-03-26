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


Table: user_role_mapping
[ 0] role_id                                        VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] user_id                                        VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "role_id": "QHAelmhZFaaHBXfUDwfYgNHNA",    "user_id": "fIFkfLqgMgVSMZfsOlyCZaGZy"}



*/

// UserRoleMapping struct is a row record of the user_role_mapping table in the keycloak database
type UserRoleMapping struct {
	//[ 0] role_id                                        VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	RoleID string `gorm:"primary_key;column:role_id;type:VARCHAR;size:255;" json:"role_id"`
	//[ 1] user_id                                        VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	UserID string `gorm:"primary_key;column:user_id;type:VARCHAR;size:36;" json:"user_id"`
}

var user_role_mappingTableInfo = &TableInfo{
	Name: "user_role_mapping",
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
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *UserRoleMapping) TableName() string {
	return "user_role_mapping"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *UserRoleMapping) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *UserRoleMapping) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *UserRoleMapping) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *UserRoleMapping) TableInfo() *TableInfo {
	return user_role_mappingTableInfo
}
