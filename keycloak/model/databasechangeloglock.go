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


Table: databasechangeloglock
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] locked                                         BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
[ 2] lockgranted                                    TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 3] lockedby                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": 93,    "locked": false,    "lockgranted": "2273-01-15T15:18:15.108612853+08:00",    "lockedby": "fjLtQSKplJNgtQiIVkCnIfTtG"}



*/

// Databasechangeloglock struct is a row record of the databasechangeloglock table in the keycloak database
type Databasechangeloglock struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;column:id;type:INT4;" json:"id"`
	//[ 1] locked                                         BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	Locked bool `gorm:"column:locked;type:BOOL;" json:"locked"`
	//[ 2] lockgranted                                    TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	Lockgranted time.Time `gorm:"column:lockgranted;type:TIMESTAMP;" json:"lockgranted"`
	//[ 3] lockedby                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Lockedby sql.NullString `gorm:"column:lockedby;type:VARCHAR(255);size:255;" json:"lockedby"`
}

var databasechangeloglockTableInfo = &TableInfo{
	Name: "databasechangeloglock",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "locked",
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
			GoFieldName:        "Locked",
			GoFieldType:        "bool",
			JSONFieldName:      "locked",
			ProtobufFieldName:  "locked",
			ProtobufType:       "bool",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "lockgranted",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "Lockgranted",
			GoFieldType:        "time.Time",
			JSONFieldName:      "lockgranted",
			ProtobufFieldName:  "lockgranted",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "lockedby",
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
			GoFieldName:        "Lockedby",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "lockedby",
			ProtobufFieldName:  "lockedby",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *Databasechangeloglock) TableName() string {
	return "databasechangeloglock"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *Databasechangeloglock) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *Databasechangeloglock) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *Databasechangeloglock) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *Databasechangeloglock) TableInfo() *TableInfo {
	return databasechangeloglockTableInfo
}
