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


CREATE TABLE `personal_access_tokens` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tokenable_type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tokenable_id` bigint unsigned NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `abilities` text COLLATE utf8mb4_unicode_ci,
  `last_used_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `personal_access_tokens_token_unique` (`token`),
  KEY `personal_access_tokens_tokenable_type_tokenable_id_index` (`tokenable_type`,`tokenable_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "id": 48,    "tokenable_type": "EcgZIflYYERYpnBWDQvZXDFYW",    "tokenable_id": 11,    "name": "oEXORBPrrCsZcZhBApEGabxVH",    "token": "ABIVwrkooVbZMtVSZRsvOjqRq",    "abilities": "uimoZZSFhquEMTEOeSEhwQyfa",    "last_used_at": "2200-01-28T23:06:05.96625396+08:00",    "created_at": "2169-09-12T09:06:11.632004116+08:00",    "updated_at": "2206-01-19T08:18:21.66445864+08:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned



*/

// PersonalAccessTokens struct is a row record of the personal_access_tokens table in the test1 database
type PersonalAccessTokens struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] tokenable_type                                 varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	TokenableType string `gorm:"column:tokenable_type;type:varchar;size:255;" json:"tokenable_type"`
	//[ 2] tokenable_id                                   ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	TokenableID uint64 `gorm:"column:tokenable_id;type:ubigint;" json:"tokenable_id"`
	//[ 3] name                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name string `gorm:"column:name;type:varchar;size:255;" json:"name"`
	//[ 4] token                                          varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	Token string `gorm:"column:token;type:varchar;size:64;" json:"token"`
	//[ 5] abilities                                      text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Abilities sql.NullString `gorm:"column:abilities;type:text;size:65535;" json:"abilities"`
	//[ 6] last_used_at                                   timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	LastUsedAt time.Time `gorm:"column:last_used_at;type:timestamp;" json:"last_used_at"`
	//[ 7] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[ 8] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var personal_access_tokensTableInfo = &TableInfo{
	Name: "personal_access_tokens",
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
			Name:               "tokenable_type",
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
			GoFieldName:        "TokenableType",
			GoFieldType:        "string",
			JSONFieldName:      "tokenable_type",
			ProtobufFieldName:  "tokenable_type",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "tokenable_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "TokenableID",
			GoFieldType:        "uint64",
			JSONFieldName:      "tokenable_id",
			ProtobufFieldName:  "tokenable_id",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "name",
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
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "token",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "Token",
			GoFieldType:        "string",
			JSONFieldName:      "token",
			ProtobufFieldName:  "token",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "abilities",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Abilities",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "abilities",
			ProtobufFieldName:  "abilities",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "last_used_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "LastUsedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "last_used_at",
			ProtobufFieldName:  "last_used_at",
			ProtobufType:       "uint64",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "created_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "uint64",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "updated_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "uint64",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *PersonalAccessTokens) TableName() string {
	return "personal_access_tokens"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PersonalAccessTokens) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PersonalAccessTokens) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PersonalAccessTokens) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PersonalAccessTokens) TableInfo() *TableInfo {
	return personal_access_tokensTableInfo
}
