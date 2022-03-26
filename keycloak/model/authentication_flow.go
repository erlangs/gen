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


Table: authentication_flow
[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] alias                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] description                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] realm_id                                       VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 4] provider_id                                    VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: [basic-flow]
[ 5] top_level                                      BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[ 6] built_in                                       BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]


JSON Sample
-------------------------------------
{    "id": "SWRCYLbZDgEPtCyNPDfmmDwCL",    "alias": "eEmwAVUCoAlUEJihcqOkfLJKp",    "description": "hkeeLLoZWyPHOxTGblwbqYHih",    "realm_id": "eVkwNmPWCUPbXnHjPUaPPiHTO",    "provider_id": "wckFfviALrPGFmyihkTHqBfxI",    "top_level": true,    "built_in": true}



*/

// AuthenticationFlow struct is a row record of the authentication_flow table in the keycloak database
type AuthenticationFlow struct {
	//[ 0] id                                             VARCHAR(36)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR;size:36;" json:"id"`
	//[ 1] alias                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Alias sql.NullString `gorm:"column:alias;type:VARCHAR;size:255;" json:"alias"`
	//[ 2] description                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Description sql.NullString `gorm:"column:description;type:VARCHAR;size:255;" json:"description"`
	//[ 3] realm_id                                       VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
	RealmID sql.NullString `gorm:"column:realm_id;type:VARCHAR;size:36;" json:"realm_id"`
	//[ 4] provider_id                                    VARCHAR(36)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: [basic-flow]
	ProviderID string `gorm:"column:provider_id;type:VARCHAR;size:36;default:basic-flow;" json:"provider_id"`
	//[ 5] top_level                                      BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	TopLevel bool `gorm:"column:top_level;type:BOOL;default:false;" json:"top_level"`
	//[ 6] built_in                                       BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	BuiltIn bool `gorm:"column:built_in;type:BOOL;default:false;" json:"built_in"`
}

var authentication_flowTableInfo = &TableInfo{
	Name: "authentication_flow",
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
			Name:               "alias",
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
			GoFieldName:        "Alias",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "alias",
			ProtobufFieldName:  "alias",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "provider_id",
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
			GoFieldName:        "ProviderID",
			GoFieldType:        "string",
			JSONFieldName:      "provider_id",
			ProtobufFieldName:  "provider_id",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "top_level",
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
			GoFieldName:        "TopLevel",
			GoFieldType:        "bool",
			JSONFieldName:      "top_level",
			ProtobufFieldName:  "top_level",
			ProtobufType:       "bool",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "built_in",
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
			GoFieldName:        "BuiltIn",
			GoFieldType:        "bool",
			JSONFieldName:      "built_in",
			ProtobufFieldName:  "built_in",
			ProtobufType:       "bool",
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AuthenticationFlow) TableName() string {
	return "authentication_flow"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AuthenticationFlow) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AuthenticationFlow) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AuthenticationFlow) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AuthenticationFlow) TableInfo() *TableInfo {
	return authentication_flowTableInfo
}
