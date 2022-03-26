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


Table: migration_model
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] version                                        VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 2] update_time                                    INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: [0]


JSON Sample
-------------------------------------
{    "id": "fWYvJTanFtaXGLJjjoolqFpwm",    "version": "kjpqkPyHIlpkesALLaotYKJCH",    "update_time": 37}



*/

// MigrationModel struct is a row record of the migration_model table in the keycloak database
type MigrationModel struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR;size:36;" json:"id"`
	//[ 1] version                                        VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	Version sql.NullString `gorm:"column:version;type:VARCHAR;size:36;" json:"version"`
	//[ 2] update_time                                    INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: [0]
	UpdateTime int64 `gorm:"column:update_time;type:INT8;default:0;" json:"update_time"`
}

var migration_modelTableInfo = &TableInfo{
	Name: "migration_model",
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
			Name:               "version",
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
			GoFieldName:        "Version",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "version",
			ProtobufFieldName:  "version",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "update_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "UpdateTime",
			GoFieldType:        "int64",
			JSONFieldName:      "update_time",
			ProtobufFieldName:  "update_time",
			ProtobufType:       "int64",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *MigrationModel) TableName() string {
	return "migration_model"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *MigrationModel) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *MigrationModel) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *MigrationModel) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *MigrationModel) TableInfo() *TableInfo {
	return migration_modelTableInfo
}
