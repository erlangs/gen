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


Table: databasechangelog
[ 0] id                                             VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] author                                         VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] filename                                       VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] dateexecuted                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 4] orderexecuted                                  INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 5] exectype                                       VARCHAR(10)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 10      default: []
[ 6] md5sum                                         VARCHAR(35)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 35      default: []
[ 7] description                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] comments                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 9] tag                                            VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[10] liquibase                                      VARCHAR(20)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 20      default: []
[11] contexts                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[12] labels                                         VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[13] deployment_id                                  VARCHAR(10)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 10      default: []


JSON Sample
-------------------------------------
{    "id": "prOijFcFRbcnRyYcLoUFsYXpU",    "author": "ffpVNOFtcQJwmOLYARmWTXHLA",    "filename": "IAcyaWNmtItRmfEtHuRHEXqTG",    "dateexecuted": "2070-05-07T06:25:48.610024374+08:00",    "orderexecuted": 96,    "exectype": "ItsAgvbyRpmTXHSNxPuhpDLlE",    "md_5_sum": "TkdqHTLOCHogjUroCCtylwQVV",    "description": "rkBJuJeLWAFFCuNIQoRCaRZTh",    "comments": "sjbQJcJprSJcxEAjKoXVVouYh",    "tag": "RZSbhawoFZTfEOQQqSbAjUnrb",    "liquibase": "WeKaipypDUWTTApgburuNUfQY",    "contexts": "VJeKMeuivUZeIPFjKVyhOYYmy",    "labels": "lqeUBcghBWbBLjGfZvGdDtsDo",    "deployment_id": "nCfJUqKqDaUuyKyKVigSqatHI"}


Comments
-------------------------------------
[ 0] Warning table: databasechangelog does not have a primary key defined, setting col position 1 id as primary key




*/

// Databasechangelog struct is a row record of the databasechangelog table in the keycloak database
type Databasechangelog struct {
	//[ 0] id                                             VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ID string `gorm:"primary_key;column:id;type:VARCHAR(255);size:255;" json:"id"`
	//[ 1] author                                         VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Author string `gorm:"column:author;type:VARCHAR(255);size:255;" json:"author"`
	//[ 2] filename                                       VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Filename string `gorm:"column:filename;type:VARCHAR(255);size:255;" json:"filename"`
	//[ 3] dateexecuted                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	Dateexecuted time.Time `gorm:"column:dateexecuted;type:TIMESTAMP;" json:"dateexecuted"`
	//[ 4] orderexecuted                                  INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Orderexecuted int32 `gorm:"column:orderexecuted;type:INT4;" json:"orderexecuted"`
	//[ 5] exectype                                       VARCHAR(10)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 10      default: []
	Exectype string `gorm:"column:exectype;type:VARCHAR(10);size:10;" json:"exectype"`
	//[ 6] md5sum                                         VARCHAR(35)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 35      default: []
	Md5sum sql.NullString `gorm:"column:md5sum;type:VARCHAR(35);size:35;" json:"md_5_sum"`
	//[ 7] description                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Description sql.NullString `gorm:"column:description;type:VARCHAR(255);size:255;" json:"description"`
	//[ 8] comments                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Comments sql.NullString `gorm:"column:comments;type:VARCHAR(255);size:255;" json:"comments"`
	//[ 9] tag                                            VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Tag sql.NullString `gorm:"column:tag;type:VARCHAR(255);size:255;" json:"tag"`
	//[10] liquibase                                      VARCHAR(20)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 20      default: []
	Liquibase sql.NullString `gorm:"column:liquibase;type:VARCHAR(20);size:20;" json:"liquibase"`
	//[11] contexts                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Contexts sql.NullString `gorm:"column:contexts;type:VARCHAR(255);size:255;" json:"contexts"`
	//[12] labels                                         VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Labels sql.NullString `gorm:"column:labels;type:VARCHAR(255);size:255;" json:"labels"`
	//[13] deployment_id                                  VARCHAR(10)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 10      default: []
	DeploymentID sql.NullString `gorm:"column:deployment_id;type:VARCHAR(10);size:10;" json:"deployment_id"`
}

var databasechangelogTableInfo = &TableInfo{
	Name: "databasechangelog",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: ``,
			Notes: `Warning table: databasechangelog does not have a primary key defined, setting col position 1 id as primary key
`,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "author",
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
			GoFieldName:        "Author",
			GoFieldType:        "string",
			JSONFieldName:      "author",
			ProtobufFieldName:  "author",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "filename",
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
			GoFieldName:        "Filename",
			GoFieldType:        "string",
			JSONFieldName:      "filename",
			ProtobufFieldName:  "filename",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "dateexecuted",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "Dateexecuted",
			GoFieldType:        "time.Time",
			JSONFieldName:      "dateexecuted",
			ProtobufFieldName:  "dateexecuted",
			ProtobufType:       "uint64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "orderexecuted",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Orderexecuted",
			GoFieldType:        "int32",
			JSONFieldName:      "orderexecuted",
			ProtobufFieldName:  "orderexecuted",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "exectype",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       10,
			GoFieldName:        "Exectype",
			GoFieldType:        "string",
			JSONFieldName:      "exectype",
			ProtobufFieldName:  "exectype",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "md5sum",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(35)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       35,
			GoFieldName:        "Md5sum",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "md_5_sum",
			ProtobufFieldName:  "md_5_sum",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "comments",
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
			GoFieldName:        "Comments",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "comments",
			ProtobufFieldName:  "comments",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "tag",
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
			GoFieldName:        "Tag",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "tag",
			ProtobufFieldName:  "tag",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "liquibase",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       20,
			GoFieldName:        "Liquibase",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "liquibase",
			ProtobufFieldName:  "liquibase",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "contexts",
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
			GoFieldName:        "Contexts",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "contexts",
			ProtobufFieldName:  "contexts",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "labels",
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
			GoFieldName:        "Labels",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "labels",
			ProtobufFieldName:  "labels",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "deployment_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       10,
			GoFieldName:        "DeploymentID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "deployment_id",
			ProtobufFieldName:  "deployment_id",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *Databasechangelog) TableName() string {
	return "databasechangelog"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *Databasechangelog) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *Databasechangelog) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *Databasechangelog) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *Databasechangelog) TableInfo() *TableInfo {
	return databasechangelogTableInfo
}
