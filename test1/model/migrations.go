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


CREATE TABLE `migrations` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "id": 64,    "migration": "rmmKgFYxGwNJRtpUIFqtdPHaf",    "batch": 10}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// Migrations struct is a row record of the migrations table in the test1 database
type Migrations struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] migration                                      varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Migration string `gorm:"column:migration;type:varchar(255);size:255;" json:"migration"`
	//[ 2] batch                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Batch int32 `gorm:"column:batch;type:int;" json:"batch"`
}

var migrationsTableInfo = &TableInfo{
	Name: "migrations",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "migration",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Migration",
			GoFieldType:        "string",
			JSONFieldName:      "migration",
			ProtobufFieldName:  "migration",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "batch",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Batch",
			GoFieldType:        "int32",
			JSONFieldName:      "batch",
			ProtobufFieldName:  "batch",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *Migrations) TableName() string {
	return "migrations"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *Migrations) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *Migrations) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *Migrations) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *Migrations) TableInfo() *TableInfo {
	return migrationsTableInfo
}
