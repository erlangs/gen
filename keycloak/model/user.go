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


Table: user
[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
[ 1] userName                                       TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] account                                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 3] nickname                                       TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 4] password                                       VARCHAR(512)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 512     default: []
[ 5] salt                                           VARCHAR(512)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 512     default: []
[ 6] email                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 7] state                                          INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 8] registration_date                              INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": 2,    "user_name": "qjkYehGDIOpGfaAnQOHSXOtNn",    "account": "GSRErVYejmGapZkgoQwGaIYdo",    "nickname": "bhUPjoKdKLhMmaBERrviaebNr",    "password": "xNFCyaCCTmLBbksOFZQwDGCED",    "salt": "XuKLMqPNBnTGcPKPxyImWMoWh",    "email": "IdcGXpXvaeRwOhaQyNtbtVoXD",    "state": 42,    "registration_date": 82}



*/

// User struct is a row record of the user table in the keycloak database
type User struct {
	//[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	//[ 1] userName                                       TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	UserName sql.NullString `gorm:"column:userName;type:TEXT;" json:"user_name"`
	//[ 2] account                                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Account sql.NullString `gorm:"column:account;type:TEXT;" json:"account"`
	//[ 3] nickname                                       TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Nickname sql.NullString `gorm:"column:nickname;type:TEXT;" json:"nickname"`
	//[ 4] password                                       VARCHAR(512)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 512     default: []
	Password sql.NullString `gorm:"column:password;type:VARCHAR(512);size:512;" json:"password"`
	//[ 5] salt                                           VARCHAR(512)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 512     default: []
	Salt sql.NullString `gorm:"column:salt;type:VARCHAR(512);size:512;" json:"salt"`
	//[ 6] email                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Email sql.NullString `gorm:"column:email;type:TEXT;" json:"email"`
	//[ 7] state                                          INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	State int64 `gorm:"column:state;type:INT8;" json:"state"`
	//[ 8] registration_date                              INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	RegistrationDate sql.NullInt64 `gorm:"column:registration_date;type:INT8;" json:"registration_date"`
}

var userTableInfo = &TableInfo{
	Name: "user",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "userName",
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
			GoFieldName:        "UserName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "user_name",
			ProtobufFieldName:  "user_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "account",
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
			GoFieldName:        "Account",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "account",
			ProtobufFieldName:  "account",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "nickname",
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
			GoFieldName:        "Nickname",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "nickname",
			ProtobufFieldName:  "nickname",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "password",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(512)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       512,
			GoFieldName:        "Password",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "password",
			ProtobufFieldName:  "password",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "salt",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(512)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       512,
			GoFieldName:        "Salt",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "salt",
			ProtobufFieldName:  "salt",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "email",
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
			GoFieldName:        "Email",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "email",
			ProtobufFieldName:  "email",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "state",
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
			GoFieldName:        "State",
			GoFieldType:        "int64",
			JSONFieldName:      "state",
			ProtobufFieldName:  "state",
			ProtobufType:       "int64",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "registration_date",
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
			GoFieldName:        "RegistrationDate",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "registration_date",
			ProtobufFieldName:  "registration_date",
			ProtobufType:       "int64",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "user"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *User) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *User) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *User) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *User) TableInfo() *TableInfo {
	return userTableInfo
}
