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


Table: keycloak_group
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] parent_group                                   VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 3] realm_id                                       VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "id": "ESjBwZutCEebPUpHCbrjNnHsE",    "name": "hgXgNNbXeBChIwHhBiWKdvvoM",    "parent_group": "OJDHPQbqQFXPNWkxBVcJDfNQk",    "realm_id": "sApfURnSoDnAosnmjneYTWiIv"}



*/

// KeycloakGroup struct is a row record of the keycloak_group table in the keycloak database
type KeycloakGroup struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR(36);size:36;" json:"id"`
	//[ 1] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Name sql.NullString `gorm:"column:name;type:VARCHAR(255);size:255;" json:"name"`
	//[ 2] parent_group                                   VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ParentGroup string `gorm:"column:parent_group;type:VARCHAR(36);size:36;" json:"parent_group"`
	//[ 3] realm_id                                       VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID sql.NullString `gorm:"column:realm_id;type:VARCHAR(36);size:36;" json:"realm_id"`
}

var keycloak_groupTableInfo = &TableInfo{
	Name: "keycloak_group",
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
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "parent_group",
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
			GoFieldName:        "ParentGroup",
			GoFieldType:        "string",
			JSONFieldName:      "parent_group",
			ProtobufFieldName:  "parent_group",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "realm_id",
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
			GoFieldName:        "RealmID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "realm_id",
			ProtobufFieldName:  "realm_id",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (k *KeycloakGroup) TableName() string {
	return "keycloak_group"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (k *KeycloakGroup) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (k *KeycloakGroup) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (k *KeycloakGroup) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (k *KeycloakGroup) TableInfo() *TableInfo {
	return keycloak_groupTableInfo
}
