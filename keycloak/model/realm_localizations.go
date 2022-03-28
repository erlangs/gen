package model

import (
	//"database/sql"
	//"time"

	//"github.com/satori/go.uuid"

	"gorm.io/gorm"
)

/*
DB Table Details
-------------------------------------


Table: realm_localizations
[ 0] realm_id                                       VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] locale                                         VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] texts                                          TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "realm_id": "dTTDoUXFePblruSNOGCiKAQnY",    "locale": "eXTVZcFrUNumAMTfJewfawlhT",    "texts": "iIVlEwbiGPgykbyZMLhsPtlgT"}



*/

// RealmLocalizations struct is a row record of the realm_localizations table in the keycloak database
type RealmLocalizations struct {
	//[ 0] realm_id                                       VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	RealmID string `gorm:"primary_key;column:realm_id;type:VARCHAR(255);size:255;" json:"realm_id"`
	//[ 1] locale                                         VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Locale string `gorm:"primary_key;column:locale;type:VARCHAR(255);size:255;" json:"locale"`
	//[ 2] texts                                          TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Texts string `gorm:"column:texts;type:TEXT;" json:"texts"`
}

var realm_localizationsTableInfo = &TableInfo{
	Name: "realm_localizations",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "realm_id",
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
			GoFieldName:        "RealmID",
			GoFieldType:        "string",
			JSONFieldName:      "realm_id",
			ProtobufFieldName:  "realm_id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "locale",
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
			GoFieldName:        "Locale",
			GoFieldType:        "string",
			JSONFieldName:      "locale",
			ProtobufFieldName:  "locale",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "texts",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Texts",
			GoFieldType:        "string",
			JSONFieldName:      "texts",
			ProtobufFieldName:  "texts",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *RealmLocalizations) TableName() string {
	return "realm_localizations"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RealmLocalizations) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RealmLocalizations) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RealmLocalizations) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *RealmLocalizations) TableInfo() *TableInfo {
	return realm_localizationsTableInfo
}
