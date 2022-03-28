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


Table: federated_identity
[ 0] identity_provider                              VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] realm_id                                       VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 2] federated_user_id                              VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] federated_username                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] token                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] user_id                                        VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "identity_provider": "SdPxPjwhfdwnGkFHSLwHRVnJe",    "realm_id": "oyWVkqaHYJplKCmTLTUdJLqxB",    "federated_user_id": "ZxTkgnKhnbIhSTkgxdBTadRaW",    "federated_username": "fTsAvLvgWkwNEAehGSkDuRWkA",    "token": "CJuOMYQtCOhnxibEPHMqpNxKc",    "user_id": "TeojvKKpqZmUFAVfXmUASvKcw"}



*/

// FederatedIdentity struct is a row record of the federated_identity table in the keycloak database
type FederatedIdentity struct {
	//[ 0] identity_provider                              VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	IdentityProvider string `gorm:"primary_key;column:identity_provider;type:VARCHAR(255);size:255;" json:"identity_provider"`
	//[ 1] realm_id                                       VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID sql.NullString `gorm:"column:realm_id;type:VARCHAR(36);size:36;" json:"realm_id"`
	//[ 2] federated_user_id                              VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	FederatedUserID sql.NullString `gorm:"column:federated_user_id;type:VARCHAR(255);size:255;" json:"federated_user_id"`
	//[ 3] federated_username                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	FederatedUsername sql.NullString `gorm:"column:federated_username;type:VARCHAR(255);size:255;" json:"federated_username"`
	//[ 4] token                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Token sql.NullString `gorm:"column:token;type:TEXT;" json:"token"`
	//[ 5] user_id                                        VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	UserID string `gorm:"primary_key;column:user_id;type:VARCHAR(36);size:36;" json:"user_id"`
}

var federated_identityTableInfo = &TableInfo{
	Name: "federated_identity",
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
			Name:               "realm_id",
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
			GoFieldName:        "RealmID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "realm_id",
			ProtobufFieldName:  "realm_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "federated_user_id",
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
			GoFieldName:        "FederatedUserID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "federated_user_id",
			ProtobufFieldName:  "federated_user_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "federated_username",
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
			GoFieldName:        "FederatedUsername",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "federated_username",
			ProtobufFieldName:  "federated_username",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "user_id",
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
			GoFieldName:        "UserID",
			GoFieldType:        "string",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (f *FederatedIdentity) TableName() string {
	return "federated_identity"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *FederatedIdentity) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *FederatedIdentity) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *FederatedIdentity) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *FederatedIdentity) TableInfo() *TableInfo {
	return federated_identityTableInfo
}
