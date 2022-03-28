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


Table: resource_server_resource
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] name                                           VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] type                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] icon_uri                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] owner                                          VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 5] resource_server_id                             VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 6] owner_managed_access                           BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[ 7] display_name                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "WOeeJVLfRdsPUeuGWpRXSFWgn",    "name": "FsdsvncdNVgbxscIZNjBtyJgL",    "type": "lwwkrGZBiQTSgiLvMPEbHApTl",    "icon_uri": "jgGXyOwJntbDLtvwYIeCqWXbh",    "owner": "uyAQDLjiCwWpJZRTSIsrgsUgk",    "resource_server_id": "EqOaksmdjWjMYMcLCqIsKdiZG",    "owner_managed_access": false,    "display_name": "mnmYOygSGOVUQtwslDaRpXRsK"}



*/

// ResourceServerResource struct is a row record of the resource_server_resource table in the keycloak database
type ResourceServerResource struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR(36);size:36;" json:"id"`
	//[ 1] name                                           VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Name string `gorm:"column:name;type:VARCHAR(255);size:255;" json:"name"`
	//[ 2] type                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Type sql.NullString `gorm:"column:type;type:VARCHAR(255);size:255;" json:"type"`
	//[ 3] icon_uri                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	IconURI sql.NullString `gorm:"column:icon_uri;type:VARCHAR(255);size:255;" json:"icon_uri"`
	//[ 4] owner                                          VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Owner string `gorm:"column:owner;type:VARCHAR(255);size:255;" json:"owner"`
	//[ 5] resource_server_id                             VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ResourceServerID string `gorm:"column:resource_server_id;type:VARCHAR(36);size:36;" json:"resource_server_id"`
	//[ 6] owner_managed_access                           BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	OwnerManagedAccess bool `gorm:"column:owner_managed_access;type:BOOL;default:false;" json:"owner_managed_access"`
	//[ 7] display_name                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	DisplayName sql.NullString `gorm:"column:display_name;type:VARCHAR(255);size:255;" json:"display_name"`
}

var resource_server_resourceTableInfo = &TableInfo{
	Name: "resource_server_resource",
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
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "type",
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
			GoFieldName:        "Type",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "icon_uri",
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
			GoFieldName:        "IconURI",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "icon_uri",
			ProtobufFieldName:  "icon_uri",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "owner",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Owner",
			GoFieldType:        "string",
			JSONFieldName:      "owner",
			ProtobufFieldName:  "owner",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "resource_server_id",
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
			GoFieldName:        "ResourceServerID",
			GoFieldType:        "string",
			JSONFieldName:      "resource_server_id",
			ProtobufFieldName:  "resource_server_id",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "owner_managed_access",
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
			GoFieldName:        "OwnerManagedAccess",
			GoFieldType:        "bool",
			JSONFieldName:      "owner_managed_access",
			ProtobufFieldName:  "owner_managed_access",
			ProtobufType:       "bool",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "display_name",
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
			GoFieldName:        "DisplayName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "display_name",
			ProtobufFieldName:  "display_name",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *ResourceServerResource) TableName() string {
	return "resource_server_resource"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *ResourceServerResource) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *ResourceServerResource) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *ResourceServerResource) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *ResourceServerResource) TableInfo() *TableInfo {
	return resource_server_resourceTableInfo
}
