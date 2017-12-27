package db_controller

import (
	// "../../db"
	"../../lib/tool/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func AllTabelColumns(c *gin.Context) {
	fmt.Println("...")
	obj := db.GetAllColNameByTableName("maxiiot_billing_invoices")
	// str := ProduceStructTag(obj, "json")
	c.JSON(200, obj)
}

func AllTabels(c *gin.Context) {
	fmt.Println("...")
	obj := db.GetAllTabelName("maxiiot_test")
	c.JSON(200, obj)
}

/*
1.找到所有表格
2.分别处理单个表格
3.找到单个表格所有字段及属性
4.生成文档
*/

func GenerateTableFile(c *gin.Context) {
	packageStr := "package pkg_name\n"
	importStr := "import(\n"
	//
	con := fmt.Sprintf("const (\n tablename = \"%s\"\n)\n", "maxiiot_billing_invoices")
	// 获取所有字段属性
	cols := db.GetAllColNameByTableName("maxiiot_billing_invoices")
	structStr := "type Invoice struct {\n"
	for _, item := range cols {
		structStr += fmt.Sprintf("%s\n", item.Tag)
	}

	for _, item := range cols {
		if item.DataType == "datetime" || item.DataType == "timestamp" {
			importStr += fmt.Sprintf("\"time\"\n")
			break
		}
	}
	importStr += ")\n"
	structStr += fmt.Sprintf("}")
	str := fmt.Sprintf("%s%s%s%s", packageStr, importStr, con, structStr)
	err := WriteTableModel(str)
	if err != nil {
		c.JSON(200, err)
	} else {
		c.JSON(200, "ok")
	}
}

func WriteTableModel(str string) error {
	f, err := os.Create("./controllers/db_controller/aaa.go")
	if err != nil {
		return err
	}
	f.WriteString(str)
	f.Close()
	return nil
}
