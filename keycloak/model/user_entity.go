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


Table: user_entity
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] email                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] email_constraint                               VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] email_verified                                 BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[ 4] enabled                                        BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[ 5] federation_link                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] first_name                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 7] last_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] realm_id                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 9] username                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[10] created_timestamp                              INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[11] service_account_client_link                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[12] not_before                                     INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]


JSON Sample
-------------------------------------
{    "id": "hXaAuMRkfoSTwdZtPXFdtbETT",    "email": "WcMwEckuMbOYSBdLnxVBNUQgh",    "email_constraint": "tKrRGKhlfskFYWExyMTAMARxe",    "email_verified": true,    "enabled": true,    "federation_link": "MQZBhyMauTUmepXeYVOevFdwg",    "first_name": "mkYNQWlWBixchMOWekuhslAZo",    "last_name": "oAWxleewsPewsaJtqcQTXWmvt",    "realm_id": "YpvLaSYsBkLRahZsLxNIGBmKO",    "username": "JAQDEsHCdHXGnYDQbOwnsDhbp",    "created_timestamp": 82,    "service_account_client_link": "hcjNVwOJkKVLLcOAxJUrfassh",    "not_before": 57}



*/

// UserEntity struct is a row record of the user_entity table in the keycloak database
type UserEntity struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR(36);size:36;" json:"id"`
	//[ 1] email                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Email sql.NullString `gorm:"column:email;type:VARCHAR(255);size:255;" json:"email"`
	//[ 2] email_constraint                               VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	EmailConstraint sql.NullString `gorm:"column:email_constraint;type:VARCHAR(255);size:255;" json:"email_constraint"`
	//[ 3] email_verified                                 BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	EmailVerified bool `gorm:"column:email_verified;type:BOOL;default:false;" json:"email_verified"`
	//[ 4] enabled                                        BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	Enabled bool `gorm:"column:enabled;type:BOOL;default:false;" json:"enabled"`
	//[ 5] federation_link                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	FederationLink sql.NullString `gorm:"column:federation_link;type:VARCHAR(255);size:255;" json:"federation_link"`
	//[ 6] first_name                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	FirstName sql.NullString `gorm:"column:first_name;type:VARCHAR(255);size:255;" json:"first_name"`
	//[ 7] last_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	LastName sql.NullString `gorm:"column:last_name;type:VARCHAR(255);size:255;" json:"last_name"`
	//[ 8] realm_id                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	RealmID sql.NullString `gorm:"column:realm_id;type:VARCHAR(255);size:255;" json:"realm_id"`
	//[ 9] username                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Username sql.NullString `gorm:"column:username;type:VARCHAR(255);size:255;" json:"username"`
	//[10] created_timestamp                              INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	CreatedTimestamp sql.NullInt64 `gorm:"column:created_timestamp;type:INT8;" json:"created_timestamp"`
	//[11] service_account_client_link                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ServiceAccountClientLink sql.NullString `gorm:"column:service_account_client_link;type:VARCHAR(255);size:255;" json:"service_account_client_link"`
	//[12] not_before                                     INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
	NotBefore int32 `gorm:"column:not_before;type:INT4;default:0;" json:"not_before"`
}

var user_entityTableInfo = &TableInfo{
	Name: "user_entity",
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
			Name:               "email",
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
			GoFieldName:        "Email",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "email",
			ProtobufFieldName:  "email",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "email_constraint",
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
			GoFieldName:        "EmailConstraint",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "email_constraint",
			ProtobufFieldName:  "email_constraint",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "email_verified",
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
			GoFieldName:        "EmailVerified",
			GoFieldType:        "bool",
			JSONFieldName:      "email_verified",
			ProtobufFieldName:  "email_verified",
			ProtobufType:       "bool",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "enabled",
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
			GoFieldName:        "Enabled",
			GoFieldType:        "bool",
			JSONFieldName:      "enabled",
			ProtobufFieldName:  "enabled",
			ProtobufType:       "bool",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "federation_link",
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
			GoFieldName:        "FederationLink",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "federation_link",
			ProtobufFieldName:  "federation_link",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "first_name",
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
			GoFieldName:        "FirstName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "first_name",
			ProtobufFieldName:  "first_name",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "last_name",
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
			GoFieldName:        "LastName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "last_name",
			ProtobufFieldName:  "last_name",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "realm_id",
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
			GoFieldName:        "RealmID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "realm_id",
			ProtobufFieldName:  "realm_id",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "username",
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
			GoFieldName:        "Username",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "username",
			ProtobufFieldName:  "username",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "created_timestamp",
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
			GoFieldName:        "CreatedTimestamp",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "created_timestamp",
			ProtobufFieldName:  "created_timestamp",
			ProtobufType:       "int64",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "service_account_client_link",
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
			GoFieldName:        "ServiceAccountClientLink",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "service_account_client_link",
			ProtobufFieldName:  "service_account_client_link",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "not_before",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "NotBefore",
			GoFieldType:        "int32",
			JSONFieldName:      "not_before",
			ProtobufFieldName:  "not_before",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *UserEntity) TableName() string {
	return "user_entity"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *UserEntity) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *UserEntity) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *UserEntity) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *UserEntity) TableInfo() *TableInfo {
	return user_entityTableInfo
}
