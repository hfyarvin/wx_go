package db_controller

import (
	// "../../db"
	"../../lib/tool/db"
	"fmt"
	"github.com/gin-gonic/gin"
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

func GenerateTableFile(table_name string) {

}
