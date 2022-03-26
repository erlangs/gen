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


Table: offline_client_session
[ 0] user_session_id                                VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] client_id                                      VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] offline_flag                                   VARCHAR(4)           null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 4       default: []
[ 3] timestamp                                      INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 4] data                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] client_storage_provider                        VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: [local]
[ 6] external_client_id                             VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: [local]


JSON Sample
-------------------------------------
{    "user_session_id": "aUwTFqXppFhOFwVFjedeeGgRR",    "client_id": "UpBHWYrCaeCXTbNJyREdrfEam",    "offline_flag": "htIOwVjdsGrGaiPOvlJVqnhFE",    "timestamp": 16,    "data": "qIlFCmmOoTTPdKHrwTCeNMfMO",    "client_storage_provider": "rNQaQRIAHupmrNXQwgXBThhgR",    "external_client_id": "AVEFdShRhkEATACSuEbuSkKto"}



*/

// OfflineClientSession struct is a row record of the offline_client_session table in the keycloak database
type OfflineClientSession struct {
	//[ 0] user_session_id                                VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	UserSessionID string `gorm:"primary_key;column:user_session_id;type:VARCHAR;size:36;" json:"user_session_id"`
	//[ 1] client_id                                      VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ClientID string `gorm:"primary_key;column:client_id;type:VARCHAR;size:255;" json:"client_id"`
	//[ 2] offline_flag                                   VARCHAR(4)           null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 4       default: []
	OfflineFlag string `gorm:"primary_key;column:offline_flag;type:VARCHAR;size:4;" json:"offline_flag"`
	//[ 3] timestamp                                      INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Timestamp sql.NullInt32 `gorm:"column:timestamp;type:INT4;" json:"timestamp"`
	//[ 4] data                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Data sql.NullString `gorm:"column:data;type:TEXT;" json:"data"`
	//[ 5] client_storage_provider                        VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: [local]
	ClientStorageProvider string `gorm:"primary_key;column:client_storage_provider;type:VARCHAR;size:36;default:local;" json:"client_storage_provider"`
	//[ 6] external_client_id                             VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: [local]
	ExternalClientID string `gorm:"primary_key;column:external_client_id;type:VARCHAR;size:255;default:local;" json:"external_client_id"`
}

var offline_client_sessionTableInfo = &TableInfo{
	Name: "offline_client_session",
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
			Name:               "client_id",
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
			GoFieldName:        "ClientID",
			GoFieldType:        "string",
			JSONFieldName:      "client_id",
			ProtobufFieldName:  "client_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "client_storage_provider",
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
			GoFieldName:        "ClientStorageProvider",
			GoFieldType:        "string",
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
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ExternalClientID",
			GoFieldType:        "string",
			JSONFieldName:      "external_client_id",
			ProtobufFieldName:  "external_client_id",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *OfflineClientSession) TableName() string {
	return "offline_client_session"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OfflineClientSession) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OfflineClientSession) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OfflineClientSession) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *OfflineClientSession) TableInfo() *TableInfo {
	return offline_client_sessionTableInfo
}
