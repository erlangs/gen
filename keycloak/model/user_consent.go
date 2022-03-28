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


Table: user_consent
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] client_id                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] user_id                                        VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 3] created_date                                   INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 4] last_updated_date                              INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 5] client_storage_provider                        VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 6] external_client_id                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "IZyviYHJTBUQTwyrDfxmQxZTM",    "client_id": "MxkxyahjivaSuyiKrxAVQUDTJ",    "user_id": "VpTapSpTfdIUdgeHEupUHgPoM",    "created_date": 16,    "last_updated_date": 22,    "client_storage_provider": "sniheGVErcEFSwMLVYHDdNNux",    "external_client_id": "KbnvKSRZIVLakMsYRwnrmVqPy"}



*/

// UserConsent struct is a row record of the user_consent table in the keycloak database
type UserConsent struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR(36);size:36;" json:"id"`
	//[ 1] client_id                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ClientID sql.NullString `gorm:"column:client_id;type:VARCHAR(255);size:255;" json:"client_id"`
	//[ 2] user_id                                        VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	UserID string `gorm:"column:user_id;type:VARCHAR(36);size:36;" json:"user_id"`
	//[ 3] created_date                                   INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	CreatedDate sql.NullInt64 `gorm:"column:created_date;type:INT8;" json:"created_date"`
	//[ 4] last_updated_date                              INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	LastUpdatedDate sql.NullInt64 `gorm:"column:last_updated_date;type:INT8;" json:"last_updated_date"`
	//[ 5] client_storage_provider                        VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ClientStorageProvider sql.NullString `gorm:"column:client_storage_provider;type:VARCHAR(36);size:36;" json:"client_storage_provider"`
	//[ 6] external_client_id                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ExternalClientID sql.NullString `gorm:"column:external_client_id;type:VARCHAR(255);size:255;" json:"external_client_id"`
}

var user_consentTableInfo = &TableInfo{
	Name: "user_consent",
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
			DatabaseTypePretty: "VARCHAR(36)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       36,
			GoFieldName:        "UserID",
			GoFieldType:        "string",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *UserConsent) TableName() string {
	return "user_consent"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *UserConsent) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *UserConsent) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *UserConsent) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *UserConsent) TableInfo() *TableInfo {
	return user_consentTableInfo
}
