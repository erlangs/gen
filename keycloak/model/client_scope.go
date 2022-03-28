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


Table: client_scope
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] realm_id                                       VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 3] description                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] protocol                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "egQUicPaYYIAbOHqukjbrTtWo",    "name": "lCtHtiHQaltFCnSWcxjAAaggM",    "realm_id": "uWTVHBjrosijwDVUvjuFInRyw",    "description": "cllGmJtCRTlxlxrObFbmEwZMo",    "protocol": "urFDvRuKFotlMkdwogcosZJeH"}



*/

// ClientScope struct is a row record of the client_scope table in the keycloak database
type ClientScope struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR(36);size:36;" json:"id"`
	//[ 1] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Name sql.NullString `gorm:"column:name;type:VARCHAR(255);size:255;" json:"name"`
	//[ 2] realm_id                                       VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID sql.NullString `gorm:"column:realm_id;type:VARCHAR(36);size:36;" json:"realm_id"`
	//[ 3] description                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Description sql.NullString `gorm:"column:description;type:VARCHAR(255);size:255;" json:"description"`
	//[ 4] protocol                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Protocol sql.NullString `gorm:"column:protocol;type:VARCHAR(255);size:255;" json:"protocol"`
}

var client_scopeTableInfo = &TableInfo{
	Name: "client_scope",
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
			Name:               "name",
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
			GoFieldName:        "Name",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "description",
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
			GoFieldName:        "Description",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "description",
			ProtobufFieldName:  "description",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "protocol",
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
			GoFieldName:        "Protocol",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "protocol",
			ProtobufFieldName:  "protocol",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *ClientScope) TableName() string {
	return "client_scope"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ClientScope) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ClientScope) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ClientScope) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ClientScope) TableInfo() *TableInfo {
	return client_scopeTableInfo
}
