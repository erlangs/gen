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


Table: default_client_scope
[ 0] realm_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] scope_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 2] default_scope                                  BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]


JSON Sample
-------------------------------------
{    "realm_id": "eRrKBHpKewOlmlcwcAmnJkSdB",    "scope_id": "pFEIcuOHIkAAEfnbvZgXUxhLC",    "default_scope": false}



*/

// DefaultClientScope struct is a row record of the default_client_scope table in the keycloak database
type DefaultClientScope struct {
	//[ 0] realm_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID string `gorm:"primary_key;column:realm_id;type:VARCHAR(36);size:36;" json:"realm_id"`
	//[ 1] scope_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ScopeID string `gorm:"primary_key;column:scope_id;type:VARCHAR(36);size:36;" json:"scope_id"`
	//[ 2] default_scope                                  BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	DefaultScope bool `gorm:"column:default_scope;type:BOOL;default:false;" json:"default_scope"`
}

var default_client_scopeTableInfo = &TableInfo{
	Name: "default_client_scope",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
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
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "scope_id",
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
func (d *DefaultClientScope) TableName() string {
	return "default_client_scope"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *DefaultClientScope) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *DefaultClientScope) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *DefaultClientScope) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *DefaultClientScope) TableInfo() *TableInfo {
	return default_client_scopeTableInfo
}
