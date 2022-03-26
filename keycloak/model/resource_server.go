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


Table: resource_server
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] allow_rs_remote_mgmt                           BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[ 2] policy_enforce_mode                            VARCHAR(15)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 15      default: []
[ 3] decision_strategy                              INT2                 null: false  primary: false  isArray: false  auto: false  col: INT2            len: -1      default: [1]


JSON Sample
-------------------------------------
{    "id": "pFtKSdHxPcvvJebpZXMHxjbeq",    "allow_rs_remote_mgmt": true,    "policy_enforce_mode": "syYgWgiZrYjIRvGZLwKdpPlWa",    "decision_strategy": 35}



*/

// ResourceServer struct is a row record of the resource_server table in the keycloak database
type ResourceServer struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR;size:36;" json:"id"`
	//[ 1] allow_rs_remote_mgmt                           BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	AllowRsRemoteMgmt bool `gorm:"column:allow_rs_remote_mgmt;type:BOOL;default:false;" json:"allow_rs_remote_mgmt"`
	//[ 2] policy_enforce_mode                            VARCHAR(15)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 15      default: []
	PolicyEnforceMode string `gorm:"column:policy_enforce_mode;type:VARCHAR;size:15;" json:"policy_enforce_mode"`
	//[ 3] decision_strategy                              INT2                 null: false  primary: false  isArray: false  auto: false  col: INT2            len: -1      default: [1]
	DecisionStrategy int32 `gorm:"column:decision_strategy;type:INT2;default:1;" json:"decision_strategy"`
}

var resource_serverTableInfo = &TableInfo{
	Name: "resource_server",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
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
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "allow_rs_remote_mgmt",
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
			GoFieldName:        "AllowRsRemoteMgmt",
			GoFieldType:        "bool",
			JSONFieldName:      "allow_rs_remote_mgmt",
			ProtobufFieldName:  "allow_rs_remote_mgmt",
			ProtobufType:       "bool",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "policy_enforce_mode",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(15)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       15,
			GoFieldName:        "PolicyEnforceMode",
			GoFieldType:        "string",
			JSONFieldName:      "policy_enforce_mode",
			ProtobufFieldName:  "policy_enforce_mode",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "decision_strategy",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT2",
			DatabaseTypePretty: "INT2",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT2",
			ColumnLength:       -1,
			GoFieldName:        "DecisionStrategy",
			GoFieldType:        "int32",
			JSONFieldName:      "decision_strategy",
			ProtobufFieldName:  "decision_strategy",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *ResourceServer) TableName() string {
	return "resource_server"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *ResourceServer) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *ResourceServer) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *ResourceServer) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *ResourceServer) TableInfo() *TableInfo {
	return resource_serverTableInfo
}
