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


Table: fed_user_consent
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] client_id                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] user_id                                        VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] realm_id                                       VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 4] storage_provider_id                            VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 5] created_date                                   INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 6] last_updated_date                              INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 7] client_storage_provider                        VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 8] external_client_id                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "PEHFNUnbSHXTEPFVVeijWUPwv",    "client_id": "xsXnhRDGWlnCXTuDwPqCsTbFN",    "user_id": "WrEcrPlKWKhLncTRWTPCLSsfr",    "realm_id": "RbkLYqWQVblWHaoqTlLQONDBe",    "storage_provider_id": "gCnqgnZovHThGqbLNOadXdHFM",    "created_date": 6,    "last_updated_date": 13,    "client_storage_provider": "UHWQmVtRAGhfpgtnYEqqWNwrh",    "external_client_id": "iHstJhowrynqGnkAZtsjKkjjS"}



*/

// FedUserConsent struct is a row record of the fed_user_consent table in the keycloak database
type FedUserConsent struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR;size:36;" json:"id"`
	//[ 1] client_id                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ClientID sql.NullString `gorm:"column:client_id;type:VARCHAR;size:255;" json:"client_id"`
	//[ 2] user_id                                        VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	UserID string `gorm:"column:user_id;type:VARCHAR;size:255;" json:"user_id"`
	//[ 3] realm_id                                       VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID string `gorm:"column:realm_id;type:VARCHAR;size:36;" json:"realm_id"`
	//[ 4] storage_provider_id                            VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	StorageProviderID sql.NullString `gorm:"column:storage_provider_id;type:VARCHAR;size:36;" json:"storage_provider_id"`
	//[ 5] created_date                                   INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	CreatedDate sql.NullInt64 `gorm:"column:created_date;type:INT8;" json:"created_date"`
	//[ 6] last_updated_date                              INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	LastUpdatedDate sql.NullInt64 `gorm:"column:last_updated_date;type:INT8;" json:"last_updated_date"`
	//[ 7] client_storage_provider                        VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ClientStorageProvider sql.NullString `gorm:"column:client_storage_provider;type:VARCHAR;size:36;" json:"client_storage_provider"`
	//[ 8] external_client_id                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ExternalClientID sql.NullString `gorm:"column:external_client_id;type:VARCHAR;size:255;" json:"external_client_id"`
}

var fed_user_consentTableInfo = &TableInfo{
	Name: "fed_user_consent",
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
			Name:               "client_id",
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
			GoFieldName:        "ClientID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "client_id",
			ProtobufFieldName:  "client_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "user_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "UserID",
			GoFieldType:        "string",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "storage_provider_id",
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
			GoFieldName:        "StorageProviderID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "storage_provider_id",
			ProtobufFieldName:  "storage_provider_id",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "created_date",
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
			GoFieldName:        "CreatedDate",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "created_date",
			ProtobufFieldName:  "created_date",
			ProtobufType:       "int64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "last_updated_date",
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
			GoFieldName:        "LastUpdatedDate",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "last_updated_date",
			ProtobufFieldName:  "last_updated_date",
			ProtobufType:       "int64",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "client_storage_provider",
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
			GoFieldName:        "ClientStorageProvider",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "client_storage_provider",
			ProtobufFieldName:  "client_storage_provider",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "external_client_id",
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
			GoFieldName:        "ExternalClientID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "external_client_id",
			ProtobufFieldName:  "external_client_id",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (f *FedUserConsent) TableName() string {
	return "fed_user_consent"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *FedUserConsent) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *FedUserConsent) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *FedUserConsent) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *FedUserConsent) TableInfo() *TableInfo {
	return fed_user_consentTableInfo
}
