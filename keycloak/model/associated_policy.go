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


Table: associated_policy
[ 0] policy_id                                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] associated_policy_id                           VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "policy_id": "sflSbOiBHPDBkVqAhSpvgkYHI",    "associated_policy_id": "knwQIbMtAgaCjWXgryxAnoHte"}



*/

// AssociatedPolicy struct is a row record of the associated_policy table in the keycloak database
type AssociatedPolicy struct {
	//[ 0] policy_id                                      VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	PolicyID string `gorm:"primary_key;column:policy_id;type:VARCHAR;size:36;" json:"policy_id"`
	//[ 1] associated_policy_id                           VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	AssociatedPolicyID string `gorm:"primary_key;column:associated_policy_id;type:VARCHAR;size:36;" json:"associated_policy_id"`
}

var associated_policyTableInfo = &TableInfo{
	Name: "associated_policy",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
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
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "associated_policy_id",
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
			GoFieldName:        "AssociatedPolicyID",
			GoFieldType:        "string",
			JSONFieldName:      "associated_policy_id",
			ProtobufFieldName:  "associated_policy_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AssociatedPolicy) TableName() string {
	return "associated_policy"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AssociatedPolicy) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AssociatedPolicy) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AssociatedPolicy) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AssociatedPolicy) TableInfo() *TableInfo {
	return associated_policyTableInfo
}
