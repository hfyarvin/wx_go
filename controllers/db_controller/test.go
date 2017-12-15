package db_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func Gen(c *gin.Context) {
	//你的结构体定义
	type ColAttribute struct {
		TableName              string
		ColumnName             string
		OrdinalPosition        string
		ColumnDefault          string
		IsNullable             string
		DataType               string
		CharacterMaximumLength string
		CharacterOctetLength   string
		NumericPrecision       string
		NumericScale           string
		DateTimePrecision      string
		CharacterSetName       string
		CollationName          string
		ColumnType             string
		ColumnKey              string
		Extra                  string
		Privileges             string
		ColumnComment          string
		GenerationExpression   string
	}

	var s ColAttribute
	//为结构体中的变量,生成json的tag
	//把单词用下划线连接（通过大写字母来区分）
	str := ProduceStructTag(s, "json")
	fmt.Println(str)
	// writeStr := "type db_controller.MyStruct struct {\n\tName\tstring\t\t`json:\"name\"`\n\tMaxHeight\tint\t\t`json:\"max_height\"`\n}\n"
	err := writeAFile(str)
	if err != nil {
		c.JSON(200, err)
	} else {
		c.JSON(200, str)
	}
}

func writeAFile(str string) error {
	f, err := os.Create("./controllers/db_controller/s.go")
	if err != nil {
		return err
	}
	f.WriteString(str)
	f.Close()
	return nil
}
