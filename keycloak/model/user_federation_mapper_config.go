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


Table: user_federation_mapper_config
[ 0] user_federation_mapper_id                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] value                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "user_federation_mapper_id": "TtAYPInunoEuQpAvKaTrrHghB",    "value": "YGHLYLLZZZlpYeQkoRubxxBYN",    "name": "OpfkLvtDoHNQnebMqVrokepEc"}



*/

// UserFederationMapperConfig struct is a row record of the user_federation_mapper_config table in the keycloak database
type UserFederationMapperConfig struct {
	//[ 0] user_federation_mapper_id                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	UserFederationMapperID string `gorm:"primary_key;column:user_federation_mapper_id;type:VARCHAR(36);size:36;" json:"user_federation_mapper_id"`
	//[ 1] value                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Value sql.NullString `gorm:"column:value;type:VARCHAR(255);size:255;" json:"value"`
	//[ 2] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Name string `gorm:"primary_key;column:name;type:VARCHAR(255);size:255;" json:"name"`
}

var user_federation_mapper_configTableInfo = &TableInfo{
	Name: "user_federation_mapper_config",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "user_federation_mapper_id",
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
			GoFieldName:        "UserFederationMapperID",
			GoFieldType:        "string",
			JSONFieldName:      "user_federation_mapper_id",
			ProtobufFieldName:  "user_federation_mapper_id",
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
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Value",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *UserFederationMapperConfig) TableName() string {
	return "user_federation_mapper_config"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *UserFederationMapperConfig) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *UserFederationMapperConfig) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *UserFederationMapperConfig) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *UserFederationMapperConfig) TableInfo() *TableInfo {
	return user_federation_mapper_configTableInfo
}
