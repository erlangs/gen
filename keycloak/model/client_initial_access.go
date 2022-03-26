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


Table: client_initial_access
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] realm_id                                       VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 2] timestamp                                      INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 3] expiration                                     INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 4] count                                          INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 5] remaining_count                                INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "XfgSuOblDDKddsbDEZKXHBHVK",    "realm_id": "MvEMCcvMIOYCenolfiylrdbyY",    "timestamp": 28,    "expiration": 82,    "count": 29,    "remaining_count": 40}



*/

// ClientInitialAccess struct is a row record of the client_initial_access table in the keycloak database
type ClientInitialAccess struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR;size:36;" json:"id"`
	//[ 1] realm_id                                       VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID string `gorm:"column:realm_id;type:VARCHAR;size:36;" json:"realm_id"`
	//[ 2] timestamp                                      INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Timestamp sql.NullInt32 `gorm:"column:timestamp;type:INT4;" json:"timestamp"`
	//[ 3] expiration                                     INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Expiration sql.NullInt32 `gorm:"column:expiration;type:INT4;" json:"expiration"`
	//[ 4] count                                          INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Count sql.NullInt32 `gorm:"column:count;type:INT4;" json:"count"`
	//[ 5] remaining_count                                INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	RemainingCount sql.NullInt32 `gorm:"column:remaining_count;type:INT4;" json:"remaining_count"`
}

var client_initial_accessTableInfo = &TableInfo{
	Name: "client_initial_access",
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "expiration",
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
			GoFieldName:        "Expiration",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "expiration",
			ProtobufFieldName:  "expiration",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "count",
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
			GoFieldName:        "Count",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "count",
			ProtobufFieldName:  "count",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "remaining_count",
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
			GoFieldName:        "RemainingCount",
			GoFieldType:        "sql.NullInt32",
			JSONFieldName:      "remaining_count",
			ProtobufFieldName:  "remaining_count",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *ClientInitialAccess) TableName() string {
	return "client_initial_access"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ClientInitialAccess) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ClientInitialAccess) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ClientInitialAccess) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ClientInitialAccess) TableInfo() *TableInfo {
	return client_initial_accessTableInfo
}
