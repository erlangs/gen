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


Table: scope_policy
[ 0] scope_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] policy_id                                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "scope_id": "tXbLpCtcWTUXNHRRKKRWrXiIW",    "policy_id": "DglAcGpQwAlWmKeTcwQQkUrKB"}



*/

// ScopePolicy struct is a row record of the scope_policy table in the keycloak database
type ScopePolicy struct {
	//[ 0] scope_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ScopeID string `gorm:"primary_key;column:scope_id;type:VARCHAR;size:36;" json:"scope_id"`
	//[ 1] policy_id                                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	PolicyID string `gorm:"primary_key;column:policy_id;type:VARCHAR;size:36;" json:"policy_id"`
}

var scope_policyTableInfo = &TableInfo{
	Name: "scope_policy",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
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
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "policy_id",
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
			GoFieldName:        "PolicyID",
			GoFieldType:        "string",
			JSONFieldName:      "policy_id",
			ProtobufFieldName:  "policy_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *ScopePolicy) TableName() string {
	return "scope_policy"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *ScopePolicy) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *ScopePolicy) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *ScopePolicy) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *ScopePolicy) TableInfo() *TableInfo {
	return scope_policyTableInfo
}
