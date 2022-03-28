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


Table: realm_required_credential
[ 0] type                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] form_label                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] input                                          BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[ 3] secret                                         BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[ 4] realm_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "type": "NHsYVarclSoTtlrZZLUbsAdnW",    "form_label": "iYWbOwaVhWxhWlBWVuOyiCXEy",    "input": false,    "secret": true,    "realm_id": "QapekyECHnYfGPNcZGqrYgcvW"}



*/

// RealmRequiredCredential struct is a row record of the realm_required_credential table in the keycloak database
type RealmRequiredCredential struct {
	//[ 0] type                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Type string `gorm:"primary_key;column:type;type:VARCHAR(255);size:255;" json:"type"`
	//[ 1] form_label                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	FormLabel sql.NullString `gorm:"column:form_label;type:VARCHAR(255);size:255;" json:"form_label"`
	//[ 2] input                                          BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	Input bool `gorm:"column:input;type:BOOL;default:false;" json:"input"`
	//[ 3] secret                                         BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	Secret bool `gorm:"column:secret;type:BOOL;default:false;" json:"secret"`
	//[ 4] realm_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID string `gorm:"primary_key;column:realm_id;type:VARCHAR(36);size:36;" json:"realm_id"`
}

var realm_required_credentialTableInfo = &TableInfo{
	Name: "realm_required_credential",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "type",
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
			GoFieldName:        "Type",
			GoFieldType:        "string",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "form_label",
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
			GoFieldName:        "FormLabel",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "form_label",
			ProtobufFieldName:  "form_label",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "input",
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
			GoFieldName:        "Input",
			GoFieldType:        "bool",
			JSONFieldName:      "input",
			ProtobufFieldName:  "input",
			ProtobufType:       "bool",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "secret",
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
			GoFieldName:        "Secret",
			GoFieldType:        "bool",
			JSONFieldName:      "secret",
			ProtobufFieldName:  "secret",
			ProtobufType:       "bool",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *RealmRequiredCredential) TableName() string {
	return "realm_required_credential"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RealmRequiredCredential) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RealmRequiredCredential) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RealmRequiredCredential) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *RealmRequiredCredential) TableInfo() *TableInfo {
	return realm_required_credentialTableInfo
}
