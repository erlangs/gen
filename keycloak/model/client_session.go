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


Table: client_session
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] client_id                                      VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 2] redirect_uri                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] state                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] timestamp                                      INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 5] session_id                                     VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 6] auth_method                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 7] realm_id                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] auth_user_id                                   VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 9] current_action                                 VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "id": "CjyAFAVjENaxdEEtnWDamHlWB",    "client_id": "xchEsRcEkEdprpdboOVtBxIPO",    "redirect_uri": "xCLHRuZXcIKfvZWsllkVAhYUU",    "state": "uENfRYodJwbrjPKXgHlhEJwnn",    "timestamp": 26,    "session_id": "nsZCsJhQQxWQItTQbCtikxXiv",    "auth_method": "IxFBLtrpogoSAdOMJFWSnUtqw",    "realm_id": "MddQUYQNRmxiRhICkWNUdYmmq",    "auth_user_id": "yHbDCBgEnROaOqUBhAXACCaWW",    "current_action": "kTOTtYYfWQpcaUoELdWgmIpCf"}



*/

// ClientSession struct is a row record of the client_session table in the keycloak database
type ClientSession struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR;size:36;" json:"id"`
	//[ 1] client_id                                      VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ClientID sql.NullString `gorm:"column:client_id;type:VARCHAR;size:36;" json:"client_id"`
	//[ 2] redirect_uri                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	RedirectURI sql.NullString `gorm:"column:redirect_uri;type:VARCHAR;size:255;" json:"redirect_uri"`
	//[ 3] state                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	State sql.NullString `gorm:"column:state;type:VARCHAR;size:255;" json:"state"`
	//[ 4] timestamp                                      INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Timestamp sql.NullInt32 `gorm:"column:timestamp;type:INT4;" json:"timestamp"`
	//[ 5] session_id                                     VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	SessionID sql.NullString `gorm:"column:session_id;type:VARCHAR;size:36;" json:"session_id"`
	//[ 6] auth_method                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	AuthMethod sql.NullString `gorm:"column:auth_method;type:VARCHAR;size:255;" json:"auth_method"`
	//[ 7] realm_id                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	RealmID sql.NullString `gorm:"column:realm_id;type:VARCHAR;size:255;" json:"realm_id"`
	//[ 8] auth_user_id                                   VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	AuthUserID sql.NullString `gorm:"column:auth_user_id;type:VARCHAR;size:36;" json:"auth_user_id"`
	//[ 9] current_action                                 VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	CurrentAction sql.NullString `gorm:"column:current_action;type:VARCHAR;size:36;" json:"current_action"`
}

var client_sessionTableInfo = &TableInfo{
	Name: "client_session",
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
			DatabaseTypePretty: "VARCHAR(36)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       36,
			GoFieldName:        "ClientID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "client_id",
			ProtobufFieldName:  "client_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "redirect_uri",
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
			GoFieldName:        "RedirectURI",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "redirect_uri",
			ProtobufFieldName:  "redirect_uri",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "state",
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
			GoFieldName:        "State",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "state",
			ProtobufFieldName:  "state",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "timestamp",
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
			GoFieldName:        "Timestamp",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "timestamp",
			ProtobufFieldName:  "timestamp",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "session_id",
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
			GoFieldName:        "SessionID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "session_id",
			ProtobufFieldName:  "session_id",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "auth_user_id",
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
			GoFieldName:        "AuthUserID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "auth_user_id",
			ProtobufFieldName:  "auth_user_id",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "current_action",
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
			GoFieldName:        "CurrentAction",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "current_action",
			ProtobufFieldName:  "current_action",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *ClientSession) TableName() string {
	return "client_session"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ClientSession) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ClientSession) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ClientSession) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ClientSession) TableInfo() *TableInfo {
	return client_sessionTableInfo
}
