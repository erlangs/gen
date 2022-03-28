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


Table: user_session
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] auth_method                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] ip_address                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] last_session_refresh                           INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 4] login_username                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 5] realm_id                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] remember_me                                    BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[ 7] started                                        INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 8] user_id                                        VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 9] user_session_state                             INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[10] broker_session_id                              VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[11] broker_user_id                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "oDFkCXbJSuSukqWrKXAXKeBHn",    "auth_method": "EcrgoFcLWPFYHqrmkjIvkevpO",    "ip_address": "gWJTLfIRedpsbDUOGpyZKvBcJ",    "last_session_refresh": 45,    "login_username": "DZVoqmstYoFTwUggeUdOidgJQ",    "realm_id": "OvTemRUSMbxuUCemFbtVyXnxN",    "remember_me": true,    "started": 91,    "user_id": "NFuwVWWKDmYQrFADDLrYHMpNH",    "user_session_state": 60,    "broker_session_id": "EvftXXfKKPjrFnJebKTIYiJBl",    "broker_user_id": "ELDGdHbrMUSjKFcnOQhaxKILA"}



*/

// UserSession struct is a row record of the user_session table in the keycloak database
type UserSession struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR(36);size:36;" json:"id"`
	//[ 1] auth_method                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	AuthMethod sql.NullString `gorm:"column:auth_method;type:VARCHAR(255);size:255;" json:"auth_method"`
	//[ 2] ip_address                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	IPAddress sql.NullString `gorm:"column:ip_address;type:VARCHAR(255);size:255;" json:"ip_address"`
	//[ 3] last_session_refresh                           INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	LastSessionRefresh sql.NullInt32 `gorm:"column:last_session_refresh;type:INT4;" json:"last_session_refresh"`
	//[ 4] login_username                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	LoginUsername sql.NullString `gorm:"column:login_username;type:VARCHAR(255);size:255;" json:"login_username"`
	//[ 5] realm_id                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	RealmID sql.NullString `gorm:"column:realm_id;type:VARCHAR(255);size:255;" json:"realm_id"`
	//[ 6] remember_me                                    BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	RememberMe bool `gorm:"column:remember_me;type:BOOL;default:false;" json:"remember_me"`
	//[ 7] started                                        INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Started sql.NullInt32 `gorm:"column:started;type:INT4;" json:"started"`
	//[ 8] user_id                                        VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	UserID sql.NullString `gorm:"column:user_id;type:VARCHAR(255);size:255;" json:"user_id"`
	//[ 9] user_session_state                             INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	UserSessionState sql.NullInt32 `gorm:"column:user_session_state;type:INT4;" json:"user_session_state"`
	//[10] broker_session_id                              VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	BrokerSessionID sql.NullString `gorm:"column:broker_session_id;type:VARCHAR(255);size:255;" json:"broker_session_id"`
	//[11] broker_user_id                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	BrokerUserID sql.NullString `gorm:"column:broker_user_id;type:VARCHAR(255);size:255;" json:"broker_user_id"`
}

var user_sessionTableInfo = &TableInfo{
	Name: "user_session",
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
			Name:               "auth_method",
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
			GoFieldName:        "AuthMethod",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "auth_method",
			ProtobufFieldName:  "auth_method",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "ip_address",
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
			GoFieldName:        "IPAddress",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "ip_address",
			ProtobufFieldName:  "ip_address",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "last_session_refresh",
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
			GoFieldName:        "LastSessionRefresh",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "last_session_refresh",
			ProtobufFieldName:  "last_session_refresh",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "login_username",
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
			GoFieldName:        "LoginUsername",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "login_username",
			ProtobufFieldName:  "login_username",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "remember_me",
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
			GoFieldName:        "RememberMe",
			GoFieldType:        "bool",
			JSONFieldName:      "remember_me",
			ProtobufFieldName:  "remember_me",
			ProtobufType:       "bool",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "started",
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
			GoFieldName:        "Started",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "started",
			ProtobufFieldName:  "started",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "user_id",
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
			GoFieldName:        "UserID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "user_session_state",
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
			GoFieldName:        "UserSessionState",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "user_session_state",
			ProtobufFieldName:  "user_session_state",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "broker_session_id",
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
			GoFieldName:        "BrokerSessionID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "broker_session_id",
			ProtobufFieldName:  "broker_session_id",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *UserSession) TableName() string {
	return "user_session"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *UserSession) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *UserSession) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *UserSession) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *UserSession) TableInfo() *TableInfo {
	return user_sessionTableInfo
}
