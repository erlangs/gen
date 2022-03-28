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


Table: user_group_membership
[ 0] group_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] user_id                                        VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "group_id": "TmSHkecVMTfwAWKHNleFJnYxx",    "user_id": "reToaBvYyMPktXQIbTpFsaThY"}



*/

// UserGroupMembership struct is a row record of the user_group_membership table in the keycloak database
type UserGroupMembership struct {
	//[ 0] group_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	GroupID string `gorm:"primary_key;column:group_id;type:VARCHAR(36);size:36;" json:"group_id"`
	//[ 1] user_id                                        VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	UserID string `gorm:"primary_key;column:user_id;type:VARCHAR(36);size:36;" json:"user_id"`
}

var user_group_membershipTableInfo = &TableInfo{
	Name: "user_group_membership",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "group_id",
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
			GoFieldName:        "GroupID",
			GoFieldType:        "string",
			JSONFieldName:      "group_id",
			ProtobufFieldName:  "group_id",
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
func (u *UserGroupMembership) TableName() string {
	return "user_group_membership"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *UserGroupMembership) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *UserGroupMembership) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *UserGroupMembership) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *UserGroupMembership) TableInfo() *TableInfo {
	return user_group_membershipTableInfo
}
