package model

import (
	"time"
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `password_resets` (
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  KEY `password_resets_email_index` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "email": "EpyHmcXkehPsmjMArSDdiwojY",    "token": "dHsSiRuFMhMVjieaVbpwEAGUE",    "created_at": "2150-05-30T00:46:27.095525744+08:00"}


Comments
-------------------------------------
[ 0] Warning table: password_resets does not have a primary key defined, setting col position 1 email as primary key




*/

// PasswordResets struct is a row record of the password_resets table in the test1 database
type PasswordResets struct {
	//[ 0] email                                          varchar(255)         null: false  primary: true   isArray: false  auto: false  col: varchar         len: 255     default: []
	Email string `gorm:"primary_key;column:email;type:varchar;size:255;" json:"email"`
	//[ 1] token                                          varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Token string `gorm:"column:token;type:varchar;size:255;" json:"token"`
	//[ 2] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
}

var password_resetsTableInfo = &TableInfo{
	Name: "password_resets",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "email",
			Comment: ``,
			Notes: `Warning table: password_resets does not have a primary key defined, setting col position 1 email as primary key
`,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Email",
			GoFieldType:        "string",
			JSONFieldName:      "email",
			ProtobufFieldName:  "email",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "token",
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
			GoFieldName:        "Token",
			GoFieldType:        "string",
			JSONFieldName:      "token",
			ProtobufFieldName:  "token",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *PasswordResets) TableName() string {
	return "password_resets"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PasswordResets) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PasswordResets) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PasswordResets) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PasswordResets) TableInfo() *TableInfo {
	return password_resetsTableInfo
}
