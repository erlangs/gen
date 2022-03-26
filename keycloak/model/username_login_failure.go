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


Table: username_login_failure
[ 0] realm_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] username                                       VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] failed_login_not_before                        INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 3] last_failure                                   INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 4] last_ip_failure                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 5] num_failures                                   INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []


JSON Sample
-------------------------------------
{    "realm_id": "UTKWvlRGyeSJsupDwvUGbusoF",    "username": "gCBkmuYbXUPcLTYjQBClsDqto",    "failed_login_not_before": 30,    "last_failure": 64,    "last_ip_failure": "TrmeDvZjyIAfyHMecfxEfMhLw",    "num_failures": 37}



*/

// UsernameLoginFailure struct is a row record of the username_login_failure table in the keycloak database
type UsernameLoginFailure struct {
	//[ 0] realm_id                                       VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID string `gorm:"primary_key;column:realm_id;type:VARCHAR;size:36;" json:"realm_id"`
	//[ 1] username                                       VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Username string `gorm:"primary_key;column:username;type:VARCHAR;size:255;" json:"username"`
	//[ 2] failed_login_not_before                        INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	FailedLoginNotBefore sql.NullInt32 `gorm:"column:failed_login_not_before;type:INT4;" json:"failed_login_not_before"`
	//[ 3] last_failure                                   INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	LastFailure sql.NullInt64 `gorm:"column:last_failure;type:INT8;" json:"last_failure"`
	//[ 4] last_ip_failure                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	LastIPFailure sql.NullString `gorm:"column:last_ip_failure;type:VARCHAR;size:255;" json:"last_ip_failure"`
	//[ 5] num_failures                                   INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	NumFailures sql.NullInt32 `gorm:"column:num_failures;type:INT4;" json:"num_failures"`
}

var username_login_failureTableInfo = &TableInfo{
	Name: "username_login_failure",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
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
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "username",
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
			GoFieldName:        "Username",
			GoFieldType:        "string",
			JSONFieldName:      "username",
			ProtobufFieldName:  "username",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "failed_login_not_before",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "FailedLoginNotBefore",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "failed_login_not_before",
			ProtobufFieldName:  "failed_login_not_before",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "last_failure",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "LastFailure",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "last_failure",
			ProtobufFieldName:  "last_failure",
			ProtobufType:       "int64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "last_ip_failure",
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
			GoFieldName:        "LastIPFailure",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "last_ip_failure",
			ProtobufFieldName:  "last_ip_failure",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "num_failures",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "NumFailures",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "num_failures",
			ProtobufFieldName:  "num_failures",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *UsernameLoginFailure) TableName() string {
	return "username_login_failure"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *UsernameLoginFailure) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *UsernameLoginFailure) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *UsernameLoginFailure) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *UsernameLoginFailure) TableInfo() *TableInfo {
	return username_login_failureTableInfo
}
