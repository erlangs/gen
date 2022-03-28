package model

import (
	"database/sql"
	//"github.com/satori/go.uuid"

	"gorm.io/gorm"
)

/*
DB Table Details
-------------------------------------


Table: admin_event_entity
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] admin_event_time                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 2] realm_id                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] operation_type                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] auth_realm_id                                  VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 5] auth_client_id                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] auth_user_id                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 7] ip_address                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] resource_path                                  VARCHAR(2550)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 2550    default: []
[ 9] representation                                 TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[10] error                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[11] resource_type                                  VARCHAR(64)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 64      default: []


JSON Sample
-------------------------------------
{    "id": "uqDCdSBKUlfRmvOptirarUAuD",    "admin_event_time": 1,    "realm_id": "jlMMybMBaBFwsloamclIUmaCK",    "operation_type": "prGHyFkDleYpdECtiRVEpudMc",    "auth_realm_id": "wHueHZgXRPdgkHrGgEXuqjIMu",    "auth_client_id": "geTFsrRFRbyxlLFccGjmhPWEr",    "auth_user_id": "AwIvWUdXVmCnfIRnvrJuiRhAb",    "ip_address": "vWZosPuxNRMNbcgilGWDcchwy",    "resource_path": "tejFwvSbNtqXraGMFNbyuRTfa",    "representation": "nocZxQAgaHNJMrlgjvKcxDnZN",    "error": "KJnDJHpqMAfoElqMNpjuJqjVv",    "resource_type": "kGhEGTUHAGxTQFgvSOUfHMobD"}



*/

// AdminEventEntity struct is a row record of the admin_event_entity table in the keycloak database
type AdminEventEntity struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR(36);size:36;" json:"id"`
	//[ 1] admin_event_time                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	AdminEventTime sql.NullInt64 `gorm:"column:admin_event_time;type:INT8;" json:"admin_event_time"`
	//[ 2] realm_id                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	RealmID sql.NullString `gorm:"column:realm_id;type:VARCHAR(255);size:255;" json:"realm_id"`
	//[ 3] operation_type                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	OperationType sql.NullString `gorm:"column:operation_type;type:VARCHAR(255);size:255;" json:"operation_type"`
	//[ 4] auth_realm_id                                  VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	AuthRealmID sql.NullString `gorm:"column:auth_realm_id;type:VARCHAR(255);size:255;" json:"auth_realm_id"`
	//[ 5] auth_client_id                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	AuthClientID sql.NullString `gorm:"column:auth_client_id;type:VARCHAR(255);size:255;" json:"auth_client_id"`
	//[ 6] auth_user_id                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	AuthUserID sql.NullString `gorm:"column:auth_user_id;type:VARCHAR(255);size:255;" json:"auth_user_id"`
	//[ 7] ip_address                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	IPAddress sql.NullString `gorm:"column:ip_address;type:VARCHAR(255);size:255;" json:"ip_address"`
	//[ 8] resource_path                                  VARCHAR(2550)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 2550    default: []
	ResourcePath sql.NullString `gorm:"column:resource_path;type:VARCHAR(2550);size:2550;" json:"resource_path"`
	//[ 9] representation                                 TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Representation sql.NullString `gorm:"column:representation;type:TEXT;" json:"representation"`
	//[10] error                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Error sql.NullString `gorm:"column:error;type:VARCHAR(255);size:255;" json:"error"`
	//[11] resource_type                                  VARCHAR(64)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 64      default: []
	ResourceType sql.NullString `gorm:"column:resource_type;type:VARCHAR(64);size:64;" json:"resource_type"`
}

var admin_event_entityTableInfo = &TableInfo{
	Name: "admin_event_entity",
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
			Name:               "admin_event_time",
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
			GoFieldName:        "AdminEventTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "admin_event_time",
			ProtobufFieldName:  "admin_event_time",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "operation_type",
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
			GoFieldName:        "OperationType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "operation_type",
			ProtobufFieldName:  "operation_type",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "auth_realm_id",
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
			GoFieldName:        "AuthRealmID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "auth_realm_id",
			ProtobufFieldName:  "auth_realm_id",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "auth_client_id",
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
			GoFieldName:        "AuthClientID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "auth_client_id",
			ProtobufFieldName:  "auth_client_id",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "auth_user_id",
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
			GoFieldName:        "AuthUserID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "auth_user_id",
			ProtobufFieldName:  "auth_user_id",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "resource_path",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(2550)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       2550,
			GoFieldName:        "ResourcePath",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "resource_path",
			ProtobufFieldName:  "resource_path",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "representation",
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
			GoFieldName:        "Representation",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "representation",
			ProtobufFieldName:  "representation",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "error",
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
			GoFieldName:        "Error",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "error",
			ProtobufFieldName:  "error",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "resource_type",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       64,
			GoFieldName:        "ResourceType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "resource_type",
			ProtobufFieldName:  "resource_type",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AdminEventEntity) TableName() string {
	return "admin_event_entity"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AdminEventEntity) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AdminEventEntity) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AdminEventEntity) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AdminEventEntity) TableInfo() *TableInfo {
	return admin_event_entityTableInfo
}
