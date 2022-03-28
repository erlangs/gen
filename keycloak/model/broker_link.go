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


Table: broker_link
[ 0] identity_provider                              VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] storage_provider_id                            VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] realm_id                                       VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 3] broker_user_id                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] broker_username                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 5] token                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 6] user_id                                        VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "identity_provider": "QGqTqRABuXiQxYUgHTnBfFjpF",    "storage_provider_id": "QLVenRoRdqNOVYMKqDAKWFLEg",    "realm_id": "YlMVoxAVcKWxZvmxpIPtBmdbx",    "broker_user_id": "IKeoAslXsRjoFVKVVtFeTQMSA",    "broker_username": "TdgJeupTDleySgqmfFFBmDaLX",    "token": "MQKXjLPlMfqmegSpbOnClvBnG",    "user_id": "BsXRmCcYhpaOameseETNvUcje"}



*/

// BrokerLink struct is a row record of the broker_link table in the keycloak database
type BrokerLink struct {
	//[ 0] identity_provider                              VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	IdentityProvider string `gorm:"primary_key;column:identity_provider;type:VARCHAR(255);size:255;" json:"identity_provider"`
	//[ 1] storage_provider_id                            VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	StorageProviderID sql.NullString `gorm:"column:storage_provider_id;type:VARCHAR(255);size:255;" json:"storage_provider_id"`
	//[ 2] realm_id                                       VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID string `gorm:"column:realm_id;type:VARCHAR(36);size:36;" json:"realm_id"`
	//[ 3] broker_user_id                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	BrokerUserID sql.NullString `gorm:"column:broker_user_id;type:VARCHAR(255);size:255;" json:"broker_user_id"`
	//[ 4] broker_username                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	BrokerUsername sql.NullString `gorm:"column:broker_username;type:VARCHAR(255);size:255;" json:"broker_username"`
	//[ 5] token                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Token sql.NullString `gorm:"column:token;type:TEXT;" json:"token"`
	//[ 6] user_id                                        VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	UserID string `gorm:"primary_key;column:user_id;type:VARCHAR(255);size:255;" json:"user_id"`
}

var broker_linkTableInfo = &TableInfo{
	Name: "broker_link",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "identity_provider",
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
			GoFieldName:        "IdentityProvider",
			GoFieldType:        "string",
			JSONFieldName:      "identity_provider",
			ProtobufFieldName:  "identity_provider",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "storage_provider_id",
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
			GoFieldName:        "StorageProviderID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "storage_provider_id",
			ProtobufFieldName:  "storage_provider_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "broker_user_id",
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
			GoFieldName:        "BrokerUserID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "broker_user_id",
			ProtobufFieldName:  "broker_user_id",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "broker_username",
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
			GoFieldName:        "BrokerUsername",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "broker_username",
			ProtobufFieldName:  "broker_username",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "token",
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
			GoFieldName:        "Token",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "token",
			ProtobufFieldName:  "token",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "user_id",
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
			GoFieldName:        "UserID",
			GoFieldType:        "string",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *BrokerLink) TableName() string {
	return "broker_link"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BrokerLink) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BrokerLink) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BrokerLink) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BrokerLink) TableInfo() *TableInfo {
	return broker_linkTableInfo
}
