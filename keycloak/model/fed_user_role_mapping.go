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


Table: fed_user_role_mapping
[ 0] role_id                                        VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] user_id                                        VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] realm_id                                       VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 3] storage_provider_id                            VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "role_id": "lSIJDiQQHVevYxnVGwTCDxNIK",    "user_id": "amEObiIfhwXTavFFRhVxFAuXj",    "realm_id": "XhLVJRWcsLhykSLBoykBrlFHX",    "storage_provider_id": "NIsDKjqQAMSKHcXFmeLqeSSeX"}



*/

// FedUserRoleMapping struct is a row record of the fed_user_role_mapping table in the keycloak database
type FedUserRoleMapping struct {
	//[ 0] role_id                                        VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RoleID string `gorm:"primary_key;column:role_id;type:VARCHAR;size:36;" json:"role_id"`
	//[ 1] user_id                                        VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	UserID string `gorm:"primary_key;column:user_id;type:VARCHAR;size:255;" json:"user_id"`
	//[ 2] realm_id                                       VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID string `gorm:"column:realm_id;type:VARCHAR;size:36;" json:"realm_id"`
	//[ 3] storage_provider_id                            VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	StorageProviderID sql.NullString `gorm:"column:storage_provider_id;type:VARCHAR;size:36;" json:"storage_provider_id"`
}

var fed_user_role_mappingTableInfo = &TableInfo{
	Name: "fed_user_role_mapping",
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
			Name:               "user_id",
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
			GoFieldName:        "UserID",
			GoFieldType:        "string",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "realm_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(36)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       36,
			GoFieldName:        "RealmID",
			GoFieldType:        "string",
			JSONFieldName:      "realm_id",
			ProtobufFieldName:  "realm_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "storage_provider_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(36)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       36,
			GoFieldName:        "StorageProviderID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "storage_provider_id",
			ProtobufFieldName:  "storage_provider_id",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (f *FedUserRoleMapping) TableName() string {
	return "fed_user_role_mapping"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *FedUserRoleMapping) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *FedUserRoleMapping) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *FedUserRoleMapping) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *FedUserRoleMapping) TableInfo() *TableInfo {
	return fed_user_role_mappingTableInfo
}
