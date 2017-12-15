package db

import (
	"../../../db"
	"fmt"
	"strings"
)

type Table struct {
	TableName string `json:"table_name"`
}

// 获取数据库所有表名
func GetAllTabelName(dbName string) []*Table {
	sql := fmt.Sprintf("SELECT table_name FROM information_schema.TABLES WHERE TABLE_SCHEMA='%s';", dbName)
	var ret []*Table
	results, err := db.Engine.Query(sql)
	if err != nil {
		return nil
	} else {
		for _, t := range results {
			r := new(Table)
			r.TableName = string(t["table_name"])
			ret = append(ret, r)
		}
		return ret
	}
}

//数据表字段属性表
type ColAttribute struct {
	TableName              string `json:"table_name"`
	ColumnName             string `json:"column_name"`
	OrdinalPosition        string `json:"ordinal_position"`
	ColumnDefault          string `json:"column_default"`
	IsNullable             string `json:"is_nullable"`
	DataType               string `json:"data_type"`
	CharacterMaximumLength string `json:"character_maximum_length"`
	CharacterOctetLength   string `json:"character_octet_length"`
	NumericPrecision       string `json:"numeric_precision"`
	NumericScale           string `json:"numeric_scale"`
	DateTimePrecision      string `json:"date_time_precision"`
	CharacterSetName       string `json:"character_set_name"`
	CollationName          string `json:"collation_name"`
	ColumnType             string `json:"column_type"`
	ColumnKey              string `json:"column_key"`
	Extra                  string `json:"extra"`
	Privileges             string `json:"privileges"`
	ColumnComment          string `json:"column_comment"`
	GenerationExpression   string `json:"generation_expression"`
	Tag                    string `tag`
}

//获取数据库表所有字段名
func GetAllColNameByTableName(table string) []*ColAttribute {
	var ret []*ColAttribute
	sql := fmt.Sprintf("select * from information_schema.COLUMNS where table_name='%s'", table)
	results, err := db.Engine.Query(sql)
	if err != nil {
		return nil
	} else {
		for _, item := range results {
			r := new(ColAttribute)
			r.ColumnName = string(item["COLUMN_NAME"])
			r.DataType = string(item["DATA_TYPE"])
			r.TableName = string(item["TABLE_NAME"])
			r.IsNullable = string(item["IS_NULLABLE"])
			r.ColumnKey = string(item["COLUMN_KEY"])
			r.ColumnDefault = string(item["COLUMN_DEFAULT"])
			r.Tag = r.GetStructStrByCol()
			ret = append(ret, r)
		}
		return ret
	}
}

//获取构造结构体字段
func (self *ColAttribute) GetStructStrByCol() string {
	importStr := "import\t {\n"
	str := ""
	upperColName := colUpper(self.ColumnName)
	dateType := ""
	defaultStr := self.ColumnDefault
	switch self.DataType {
	case "tinyint", "smallint", "mediumint", "int":
		dateType = "int"
	case "bigint":
		dateType = "int64"
	case "float":
		dateType = "int32"
	case "double":
		dateType = "int64"
	case "decimal":
		dateType = "string"
	case "timestamp":
		dateType = "time.Time"
		importStr = fmt.Sprintf("%s'%s'\n", importStr, "time")
	case "datetime":
		dateType = "time.Time"
		importStr = fmt.Sprintf("%s'%s'\n", importStr, "time")
	case "char", "varchar", "tinytext", "text", "mediumtext", "longtext":
		dateType = "string"
		defaultStr = fmt.Sprintf("'%s'", defaultStr)
	default:
		dateType = "string"
	}
	// Created time.Time `xorm:"created DATETIME(0) NOTNULL"`
	if self.ColumnKey == "PRI" {
		return fmt.Sprintf("%s\t%s", upperColName, dateType)
	}
	xormStr := fmt.Sprintf("%s", self.ColumnName)
	if self.ColumnDefault != "" {
		xormStr := fmt.Sprintf("%s\tdefault %s", defaultStr)
	}
	var attr []string
	if self.Extra == "auto_increment" {
		attr = append(attr, "autoincr")
	}
	str = fmt.Sprintf("%s\t%s\t`xorm:'%s'\tjson:'%s'`", upperColName, dateType, xormStr, self.ColumnName)
	importStr = fmt.Sprintf("%s\n}", importStr)
	fmt.Println(importStr)
	return str
}

// 数据字段变为首字母大写
func colUpper(col string) string {
	temp := strings.Split(col, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	return upperStr
}
