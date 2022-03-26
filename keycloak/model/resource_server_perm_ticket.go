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


Table: resource_server_perm_ticket
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] owner                                          VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] requester                                      VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] created_timestamp                              INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 4] granted_timestamp                              INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 5] resource_id                                    VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 6] scope_id                                       VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 7] resource_server_id                             VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 8] policy_id                                      VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []


JSON Sample
-------------------------------------
{    "id": "hscerupBAGtlfSAMsUpjSOixy",    "owner": "akLMiyRGUthpLibtFHhgfPiWk",    "requester": "kBrlQVVJsBrnrFbFHHpFOKYDq",    "created_timestamp": 53,    "granted_timestamp": 25,    "resource_id": "SWWCMCrpZCTiifVdMGnHoWkni",    "scope_id": "pSgoJlqpBbqOEtDPgfOKdTwQK",    "resource_server_id": "DxPFnXhTuSCyYSlypbORijnhP",    "policy_id": "GRCTAqnjBMWRsKMELWAUZpgQD"}



*/

// ResourceServerPermTicket struct is a row record of the resource_server_perm_ticket table in the keycloak database
type ResourceServerPermTicket struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR;size:36;" json:"id"`
	//[ 1] owner                                          VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Owner string `gorm:"column:owner;type:VARCHAR;size:255;" json:"owner"`
	//[ 2] requester                                      VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Requester string `gorm:"column:requester;type:VARCHAR;size:255;" json:"requester"`
	//[ 3] created_timestamp                              INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	CreatedTimestamp int64 `gorm:"column:created_timestamp;type:INT8;" json:"created_timestamp"`
	//[ 4] granted_timestamp                              INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	GrantedTimestamp sql.NullInt64 `gorm:"column:granted_timestamp;type:INT8;" json:"granted_timestamp"`
	//[ 5] resource_id                                    VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ResourceID string `gorm:"column:resource_id;type:VARCHAR;size:36;" json:"resource_id"`
	//[ 6] scope_id                                       VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ScopeID sql.NullString `gorm:"column:scope_id;type:VARCHAR;size:36;" json:"scope_id"`
	//[ 7] resource_server_id                             VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ResourceServerID string `gorm:"column:resource_server_id;type:VARCHAR;size:36;" json:"resource_server_id"`
	//[ 8] policy_id                                      VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	PolicyID sql.NullString `gorm:"column:policy_id;type:VARCHAR;size:36;" json:"policy_id"`
}

var resource_server_perm_ticketTableInfo = &TableInfo{
	Name: "resource_server_perm_ticket",
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
			Name:               "owner",
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
			GoFieldName:        "Owner",
			GoFieldType:        "string",
			JSONFieldName:      "owner",
			ProtobufFieldName:  "owner",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "requester",
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
			GoFieldName:        "Requester",
			GoFieldType:        "string",
			JSONFieldName:      "requester",
			ProtobufFieldName:  "requester",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "created_timestamp",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "CreatedTimestamp",
			GoFieldType:        "int64",
			JSONFieldName:      "created_timestamp",
			ProtobufFieldName:  "created_timestamp",
			ProtobufType:       "int64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "granted_timestamp",
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
			GoFieldName:        "GrantedTimestamp",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "granted_timestamp",
			ProtobufFieldName:  "granted_timestamp",
			ProtobufType:       "int64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "resource_id",
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
			GoFieldName:        "ResourceID",
			GoFieldType:        "string",
			JSONFieldName:      "resource_id",
			ProtobufFieldName:  "resource_id",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "scope_id",
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
			GoFieldName:        "ScopeID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "scope_id",
			ProtobufFieldName:  "scope_id",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "resource_server_id",
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
			GoFieldName:        "ResourceServerID",
			GoFieldType:        "string",
			JSONFieldName:      "resource_server_id",
			ProtobufFieldName:  "resource_server_id",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "policy_id",
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
			GoFieldName:        "PolicyID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "policy_id",
			ProtobufFieldName:  "policy_id",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *ResourceServerPermTicket) TableName() string {
	return "resource_server_perm_ticket"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *ResourceServerPermTicket) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *ResourceServerPermTicket) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *ResourceServerPermTicket) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *ResourceServerPermTicket) TableInfo() *TableInfo {
	return resource_server_perm_ticketTableInfo
}
