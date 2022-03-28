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


Table: group_role_mapping
[ 0] role_id                                        VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] group_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "role_id": "VJVmVRVXdrwUlPAcoWpOxvdvl",    "group_id": "QocXtbuqoHvWgseRdecPOQBNT"}



*/

// GroupRoleMapping struct is a row record of the group_role_mapping table in the keycloak database
type GroupRoleMapping struct {
	//[ 0] role_id                                        VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RoleID string `gorm:"primary_key;column:role_id;type:VARCHAR(36);size:36;" json:"role_id"`
	//[ 1] group_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	GroupID string `gorm:"primary_key;column:group_id;type:VARCHAR(36);size:36;" json:"group_id"`
}

var group_role_mappingTableInfo = &TableInfo{
	Name: "group_role_mapping",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "role_id",
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
			GoFieldName:        "RoleID",
			GoFieldType:        "string",
			JSONFieldName:      "role_id",
			ProtobufFieldName:  "role_id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (g *GroupRoleMapping) TableName() string {
	return "group_role_mapping"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *GroupRoleMapping) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *GroupRoleMapping) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *GroupRoleMapping) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *GroupRoleMapping) TableInfo() *TableInfo {
	return group_role_mappingTableInfo
}
