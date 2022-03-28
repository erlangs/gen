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


Table: user_federation_provider
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] changed_sync_period                            INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] display_name                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] full_sync_period                               INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 4] last_sync                                      INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 5] priority                                       INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 6] provider_name                                  VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 7] realm_id                                       VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "id": "BbSlQGLhksekDllPKmLucpLeG",    "changed_sync_period": 19,    "display_name": "DUYmKTQyOddgfdoxRlJlORBuo",    "full_sync_period": 69,    "last_sync": 92,    "priority": 41,    "provider_name": "hHPhjpFYFwjKRhHgtNMNesTuB",    "realm_id": "iwqXbytQcuNSbpbrSFOrGMVxd"}



*/

// UserFederationProvider struct is a row record of the user_federation_provider table in the keycloak database
type UserFederationProvider struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR(36);size:36;" json:"id"`
	//[ 1] changed_sync_period                            INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	ChangedSyncPeriod sql.NullInt32 `gorm:"column:changed_sync_period;type:INT4;" json:"changed_sync_period"`
	//[ 2] display_name                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	DisplayName sql.NullString `gorm:"column:display_name;type:VARCHAR(255);size:255;" json:"display_name"`
	//[ 3] full_sync_period                               INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	FullSyncPeriod sql.NullInt32 `gorm:"column:full_sync_period;type:INT4;" json:"full_sync_period"`
	//[ 4] last_sync                                      INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	LastSync sql.NullInt32 `gorm:"column:last_sync;type:INT4;" json:"last_sync"`
	//[ 5] priority                                       INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Priority sql.NullInt32 `gorm:"column:priority;type:INT4;" json:"priority"`
	//[ 6] provider_name                                  VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ProviderName sql.NullString `gorm:"column:provider_name;type:VARCHAR(255);size:255;" json:"provider_name"`
	//[ 7] realm_id                                       VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID sql.NullString `gorm:"column:realm_id;type:VARCHAR(36);size:36;" json:"realm_id"`
}

var user_federation_providerTableInfo = &TableInfo{
	Name: "user_federation_provider",
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
			Name:               "changed_sync_period",
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
			GoFieldName:        "ChangedSyncPeriod",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "changed_sync_period",
			ProtobufFieldName:  "changed_sync_period",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "display_name",
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
			GoFieldName:        "DisplayName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "display_name",
			ProtobufFieldName:  "display_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "full_sync_period",
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
			GoFieldName:        "FullSyncPeriod",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "full_sync_period",
			ProtobufFieldName:  "full_sync_period",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "last_sync",
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
			GoFieldName:        "LastSync",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "last_sync",
			ProtobufFieldName:  "last_sync",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "provider_name",
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
			GoFieldName:        "ProviderName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "provider_name",
			ProtobufFieldName:  "provider_name",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *UserFederationProvider) TableName() string {
	return "user_federation_provider"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *UserFederationProvider) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *UserFederationProvider) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *UserFederationProvider) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *UserFederationProvider) TableInfo() *TableInfo {
	return user_federation_providerTableInfo
}
