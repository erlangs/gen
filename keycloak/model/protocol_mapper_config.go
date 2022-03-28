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


Table: protocol_mapper_config
[ 0] protocol_mapper_id                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] value                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "protocol_mapper_id": "qeDgFmOJvRpoGYBhCCbamBXEE",    "value": "yVyaRWSDpbclIfpYjvoGrAPUt",    "name": "YtSEdYTqAfQBXaPgGZmLDRvSx"}



*/

// ProtocolMapperConfig struct is a row record of the protocol_mapper_config table in the keycloak database
type ProtocolMapperConfig struct {
	//[ 0] protocol_mapper_id                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ProtocolMapperID string `gorm:"primary_key;column:protocol_mapper_id;type:VARCHAR(36);size:36;" json:"protocol_mapper_id"`
	//[ 1] value                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Value sql.NullString `gorm:"column:value;type:TEXT;" json:"value"`
	//[ 2] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Name string `gorm:"primary_key;column:name;type:VARCHAR(255);size:255;" json:"name"`
}

var protocol_mapper_configTableInfo = &TableInfo{
	Name: "protocol_mapper_config",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "protocol_mapper_id",
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
			GoFieldName:        "ProtocolMapperID",
			GoFieldType:        "string",
			JSONFieldName:      "protocol_mapper_id",
			ProtobufFieldName:  "protocol_mapper_id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "value",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Value",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "name",
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
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *ProtocolMapperConfig) TableName() string {
	return "protocol_mapper_config"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *ProtocolMapperConfig) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *ProtocolMapperConfig) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *ProtocolMapperConfig) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *ProtocolMapperConfig) TableInfo() *TableInfo {
	return protocol_mapper_configTableInfo
}
