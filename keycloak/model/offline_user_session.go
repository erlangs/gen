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


Table: offline_user_session
[ 0] user_session_id                                VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] user_id                                        VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] realm_id                                       VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 3] created_on                                     INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 4] offline_flag                                   VARCHAR(4)           null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 4       default: []
[ 5] data                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 6] last_session_refresh                           INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]


JSON Sample
-------------------------------------
{    "user_session_id": "DmXLgctOwjIpjTUyFhuaJAOdr",    "user_id": "fJSMDlBXSLESZbmTstTSoqAZy",    "realm_id": "TaquAGLHEiionRflMXXCaHfPH",    "created_on": 12,    "offline_flag": "iLXUmVRuspjdgnhHQoJlgRKgr",    "data": "BCCIsATtRiPfQeHdgPrAdgfSX",    "last_session_refresh": 36}



*/

// OfflineUserSession struct is a row record of the offline_user_session table in the keycloak database
type OfflineUserSession struct {
	//[ 0] user_session_id                                VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	UserSessionID string `gorm:"primary_key;column:user_session_id;type:VARCHAR(36);size:36;" json:"user_session_id"`
	//[ 1] user_id                                        VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	UserID string `gorm:"column:user_id;type:VARCHAR(255);size:255;" json:"user_id"`
	//[ 2] realm_id                                       VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID string `gorm:"column:realm_id;type:VARCHAR(36);size:36;" json:"realm_id"`
	//[ 3] created_on                                     INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	CreatedOn int32 `gorm:"column:created_on;type:INT4;" json:"created_on"`
	//[ 4] offline_flag                                   VARCHAR(4)           null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 4       default: []
	OfflineFlag string `gorm:"primary_key;column:offline_flag;type:VARCHAR(4);size:4;" json:"offline_flag"`
	//[ 5] data                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Data sql.NullString `gorm:"column:data;type:TEXT;" json:"data"`
	//[ 6] last_session_refresh                           INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
	LastSessionRefresh int32 `gorm:"column:last_session_refresh;type:INT4;default:0;" json:"last_session_refresh"`
}

var offline_user_sessionTableInfo = &TableInfo{
	Name: "offline_user_session",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "user_session_id",
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
			GoFieldName:        "UserSessionID",
			GoFieldType:        "string",
			JSONFieldName:      "user_session_id",
			ProtobufFieldName:  "user_session_id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			Name:               "created_on",
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
			GoFieldName:        "CreatedOn",
			GoFieldType:        "int32",
			JSONFieldName:      "created_on",
			ProtobufFieldName:  "created_on",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "offline_flag",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(4)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       4,
			GoFieldName:        "OfflineFlag",
			GoFieldType:        "string",
			JSONFieldName:      "offline_flag",
			ProtobufFieldName:  "offline_flag",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "data",
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
			GoFieldName:        "Data",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "data",
			ProtobufFieldName:  "data",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "last_session_refresh",
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
			GoFieldName:        "LastSessionRefresh",
			GoFieldType:        "int32",
			JSONFieldName:      "last_session_refresh",
			ProtobufFieldName:  "last_session_refresh",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *OfflineUserSession) TableName() string {
	return "offline_user_session"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OfflineUserSession) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OfflineUserSession) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OfflineUserSession) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *OfflineUserSession) TableInfo() *TableInfo {
	return offline_user_sessionTableInfo
}
