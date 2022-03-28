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


Table: realm_attribute
[ 0] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] realm_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 2] value                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "name": "xwCNwgeJRiWNeiSZKbeZMkCyb",    "realm_id": "sxbkmvZSbExAFdxdOASCgnhgT",    "value": "cFRQVbsvPVoxlqVusTysCpZsN"}



*/

// RealmAttribute struct is a row record of the realm_attribute table in the keycloak database
type RealmAttribute struct {
	//[ 0] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Name string `gorm:"primary_key;column:name;type:VARCHAR(255);size:255;" json:"name"`
	//[ 1] realm_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID string `gorm:"primary_key;column:realm_id;type:VARCHAR(36);size:36;" json:"realm_id"`
	//[ 2] value                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Value sql.NullString `gorm:"column:value;type:TEXT;" json:"value"`
}

var realm_attributeTableInfo = &TableInfo{
	Name: "realm_attribute",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
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
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "realm_id",
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
			GoFieldName:        "RealmID",
			GoFieldType:        "string",
			JSONFieldName:      "realm_id",
			ProtobufFieldName:  "realm_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "value",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
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
func (r *RealmAttribute) TableName() string {
	return "realm_attribute"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RealmAttribute) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RealmAttribute) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RealmAttribute) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *RealmAttribute) TableInfo() *TableInfo {
	return realm_attributeTableInfo
}
