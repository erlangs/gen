package model

import (
	//"database/sql"
	"time"

	//"github.com/satori/go.uuid"

	"gorm.io/gorm"
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `failed_jobs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `connection` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `queue` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `payload` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `exception` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `failed_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `failed_jobs_uuid_unique` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "id": 10,    "uuid": "MYdrQNGchVxCfcqeHaXErhABB",    "connection": "uDoNQMCYdjyNiLXwkvaGltftm",    "queue": "ZrVGOnjjDowIZdvUtxQEJyqug",    "payload": "joHQdNlKjIVnWdbTobRnJthCF",    "exception": "evlNvHKwiKyAqvlxccZVWyZib",    "failed_at": "2175-07-21T22:32:28.124048268+08:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// FailedJobs struct is a row record of the failed_jobs table in the test1 database
type FailedJobs struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] uuid                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	UUID string `gorm:"column:uuid;type:varchar;size:255;" json:"uuid"`
	//[ 2] connection                                     text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Connection string `gorm:"column:connection;type:text;size:65535;" json:"connection"`
	//[ 3] queue                                          text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Queue string `gorm:"column:queue;type:text;size:65535;" json:"queue"`
	//[ 4] payload                                        text(4294967295)     null: false  primary: false  isArray: false  auto: false  col: text            len: 4294967295 default: []
	Payload string `gorm:"column:payload;type:text;size:4294967295;" json:"payload"`
	//[ 5] exception                                      text(4294967295)     null: false  primary: false  isArray: false  auto: false  col: text            len: 4294967295 default: []
	Exception string `gorm:"column:exception;type:text;size:4294967295;" json:"exception"`
	//[ 6] failed_at                                      timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	FailedAt time.Time `gorm:"column:failed_at;type:timestamp;default:CURRENT_TIMESTAMP;" json:"failed_at"`
}

var failed_jobsTableInfo = &TableInfo{
	Name: "failed_jobs",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "uuid",
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
			GoFieldName:        "UUID",
			GoFieldType:        "string",
			JSONFieldName:      "uuid",
			ProtobufFieldName:  "uuid",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "connection",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Connection",
			GoFieldType:        "string",
			JSONFieldName:      "connection",
			ProtobufFieldName:  "connection",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "queue",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Queue",
			GoFieldType:        "string",
			JSONFieldName:      "queue",
			ProtobufFieldName:  "queue",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "payload",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(4294967295)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       4294967295,
			GoFieldName:        "Payload",
			GoFieldType:        "string",
			JSONFieldName:      "payload",
			ProtobufFieldName:  "payload",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "exception",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(4294967295)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       4294967295,
			GoFieldName:        "Exception",
			GoFieldType:        "string",
			JSONFieldName:      "exception",
			ProtobufFieldName:  "exception",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "failed_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "FailedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "failed_at",
			ProtobufFieldName:  "failed_at",
			ProtobufType:       "uint64",
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (f *FailedJobs) TableName() string {
	return "failed_jobs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *FailedJobs) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *FailedJobs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *FailedJobs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *FailedJobs) TableInfo() *TableInfo {
	return failed_jobsTableInfo
}
