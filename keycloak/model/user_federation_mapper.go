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


Table: user_federation_mapper
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] name                                           VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] federation_provider_id                         VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 3] federation_mapper_type                         VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] realm_id                                       VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "id": "lCcCQgErtaJRvxHIeVLomqAyX",    "name": "enHBOIeXgTRgvdNeVLmmUFunK",    "federation_provider_id": "LIicMIOpltOiDdXnRrMChthiJ",    "federation_mapper_type": "DsgMnVrcbFctwetIUWmDGqBsg",    "realm_id": "sytsLbZDSqEaAMddoKmJvhkcP"}



*/

// UserFederationMapper struct is a row record of the user_federation_mapper table in the keycloak database
type UserFederationMapper struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR(36);size:36;" json:"id"`
	//[ 1] name                                           VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Name string `gorm:"column:name;type:VARCHAR(255);size:255;" json:"name"`
	//[ 2] federation_provider_id                         VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	FederationProviderID string `gorm:"column:federation_provider_id;type:VARCHAR(36);size:36;" json:"federation_provider_id"`
	//[ 3] federation_mapper_type                         VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	FederationMapperType string `gorm:"column:federation_mapper_type;type:VARCHAR(255);size:255;" json:"federation_mapper_type"`
	//[ 4] realm_id                                       VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID string `gorm:"column:realm_id;type:VARCHAR(36);size:36;" json:"realm_id"`
}

var user_federation_mapperTableInfo = &TableInfo{
	Name: "user_federation_mapper",
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
			Name:               "federation_provider_id",
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
			GoFieldName:        "FederationProviderID",
			GoFieldType:        "string",
			JSONFieldName:      "federation_provider_id",
			ProtobufFieldName:  "federation_provider_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "federation_mapper_type",
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
			GoFieldName:        "FederationMapperType",
			GoFieldType:        "string",
			JSONFieldName:      "federation_mapper_type",
			ProtobufFieldName:  "federation_mapper_type",
			ProtobufType:       "string",
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
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *UserFederationMapper) TableName() string {
	return "user_federation_mapper"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *UserFederationMapper) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *UserFederationMapper) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *UserFederationMapper) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *UserFederationMapper) TableInfo() *TableInfo {
	return user_federation_mapperTableInfo
}
