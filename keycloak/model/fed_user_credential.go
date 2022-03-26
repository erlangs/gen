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


Table: fed_user_credential
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] salt                                           BYTEA                null: true   primary: false  isArray: false  auto: false  col: BYTEA           len: -1      default: []
[ 2] type                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] created_date                                   INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 4] user_id                                        VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 5] realm_id                                       VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 6] storage_provider_id                            VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 7] user_label                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] secret_data                                    TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 9] credential_data                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[10] priority                                       INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "esNIMlgvCiMySkNrFUoGcZKRB",    "salt": "vAlKtJoEFWNEPfVWkMgoMJgDa",    "type": "wwNHwPjZeYPuCgZQfRugAVEVp",    "created_date": 79,    "user_id": "TAWWuxBbnHkvApRlCcJQwLWGH",    "realm_id": "NscUUJjUkmZKrVRxiAJjyXafY",    "storage_provider_id": "sBxyMrIxVBFYYWQIMdNCoYIPu",    "user_label": "tpGWYfmfTruekUQBQNruOtSQh",    "secret_data": "lQuFkmqWNbjYnvWaXbKljRbAc",    "credential_data": "QsbBnSLfXcStKFIIdLNNIiQtW",    "priority": 88}



*/

// FedUserCredential struct is a row record of the fed_user_credential table in the keycloak database
type FedUserCredential struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR;size:36;" json:"id"`
	//[ 1] salt                                           BYTEA                null: true   primary: false  isArray: false  auto: false  col: BYTEA           len: -1      default: []
	Salt sql.NullString `gorm:"column:salt;type:BYTEA;" json:"salt"`
	//[ 2] type                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Type sql.NullString `gorm:"column:type;type:VARCHAR;size:255;" json:"type"`
	//[ 3] created_date                                   INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	CreatedDate sql.NullInt64 `gorm:"column:created_date;type:INT8;" json:"created_date"`
	//[ 4] user_id                                        VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	UserID string `gorm:"column:user_id;type:VARCHAR;size:255;" json:"user_id"`
	//[ 5] realm_id                                       VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID string `gorm:"column:realm_id;type:VARCHAR;size:36;" json:"realm_id"`
	//[ 6] storage_provider_id                            VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	StorageProviderID sql.NullString `gorm:"column:storage_provider_id;type:VARCHAR;size:36;" json:"storage_provider_id"`
	//[ 7] user_label                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	UserLabel sql.NullString `gorm:"column:user_label;type:VARCHAR;size:255;" json:"user_label"`
	//[ 8] secret_data                                    TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	SecretData sql.NullString `gorm:"column:secret_data;type:TEXT;" json:"secret_data"`
	//[ 9] credential_data                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	CredentialData sql.NullString `gorm:"column:credential_data;type:TEXT;" json:"credential_data"`
	//[10] priority                                       INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Priority sql.NullInt32 `gorm:"column:priority;type:INT4;" json:"priority"`
}

var fed_user_credentialTableInfo = &TableInfo{
	Name: "fed_user_credential",
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
			Name:               "salt",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "BYTEA",
			DatabaseTypePretty: "BYTEA",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BYTEA",
			ColumnLength:       -1,
			GoFieldName:        "Salt",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "salt",
			ProtobufFieldName:  "salt",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "type",
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
			GoFieldName:        "Type",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "user_label",
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
			GoFieldName:        "UserLabel",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "user_label",
			ProtobufFieldName:  "user_label",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "secret_data",
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
			GoFieldName:        "SecretData",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "secret_data",
			ProtobufFieldName:  "secret_data",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "credential_data",
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
			GoFieldName:        "CredentialData",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "credential_data",
			ProtobufFieldName:  "credential_data",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "priority",
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
			GoFieldName:        "Priority",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "priority",
			ProtobufFieldName:  "priority",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (f *FedUserCredential) TableName() string {
	return "fed_user_credential"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *FedUserCredential) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *FedUserCredential) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *FedUserCredential) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *FedUserCredential) TableInfo() *TableInfo {
	return fed_user_credentialTableInfo
}
