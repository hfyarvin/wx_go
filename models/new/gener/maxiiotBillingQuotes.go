package gener

import (
	"../../common"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	MaxiiotBillingQuotes_TableName = ""
)

type MaxiiotBillingQuotes struct {
	Id      int64
	Quote   int64     `xorm:"quote default 0 BIGINT(20)"`
	Created time.Time `xorm:"created DATETIME"`
	Updated time.Time `xorm:"updated DATETIME"`
}

func (self *MaxiiotBillingQuotes) TableName() string {
	return MaxiiotBillingQuotes_TableName
}

//插入数据
func (self *MaxiiotBillingQuotes) Create() *MaxiiotBillingQuotes {
	_, err := common.Engine.InsertOne(self)
	if err != nil {
		log.Fatalln("Create MaxiiotBillingQuotes Err:", err)
		return nil
	}

	return self
}

//创建数据表
func (self *MaxiiotBillingQuotes) CreateTable() error {
	err := common.Engine.CreateTables(self)
	if err != nil {
		log.Fatalln("Create MaxiiotBillingQuotes Table Error:", err)
	}
	return err
}

// 删除数据
func (u *MaxiiotBillingQuotes) Delete() {
	_, err := common.Engine.Delete(u)
	if err != nil {
		log.Fatalln(" Delete MaxiiotBillingQuotes Error:", err)
	}
}

//更新MaxiiotBillingQuotes
func (self *MaxiiotBillingQuotes) Update() {
	common.Engine.Id(self.Id).Update(self)
}

//获取MaxiiotBillingQuotes列表
//Find: 获取多条数据
func GetMaxiiotBillingQuotesList(page, perPage int, whereStr, orderStr string) []*MaxiiotBillingQuotes {
	var list []*MaxiiotBillingQuotes
	var err error
	sql := "id is not null"
	if whereStr != "" {
		sql += fmt.Sprintf(" %s", whereStr)
	}
	if orderStr != "" {
		err = common.Engine.Where(sql).OrderBy(orderStr).Limit(perPage, (page-1)*perPage).Find(&list)
	} else {
		err = common.Engine.Where(sql).Limit(perPage, (page-1)*perPage).Find(&list)
	}
	if err != nil {
		log.Fatalln("GetMaxiiotBillingQuotesList Error:", err)
		return nil
	}
	return list
}

//MaxiiotBillingQuotes总数量
func GetMaxiiotBillingQuotesCount(whereStr string) int64 {
	sql := "id is not null"
	if whereStr != "" {
		sql += fmt.Sprintf(" %s", whereStr)
	}
	total, err := common.Engine.Where(sql).Count(new(MaxiiotBillingQuotes))
	if err != nil {
		log.Fatalln("GetMaxiiotBillingQuotesCount Error:", err)
		return 0
	}
	return total
}

//通过ID获取MaxiiotBillingQuotes
func GetMaxiiotBillingQuotesById(id int64) *MaxiiotBillingQuotes {
	item := new(MaxiiotBillingQuotes)
	_, err := common.Engine.Where("id = ?", id).Get(item)
	// has, err := x.Id(id).Get(a)
	if err != nil {
		log.Fatalln("GetMaxiiotBillingQuotesById Error:", err)
		return nil
	}
	return item
}

//通过原生sql获取
func GetMaxiiotBillingQuotesBySql(sql string) []*MaxiiotBillingQuotes {
	var list []*MaxiiotBillingQuotes
	err := common.Engine.Sql(sql).Find(&list)
	if err != nil {
		log.Fatalln("GetMaxiiotBillingQuotesBySql Error:", err)
		return nil
	}
	return list
}

/*controller方法
package ****_controller

import (
	"../../models/****_model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)
//创建数据表
func CreateMaxiiotBillingQuotesTable(c *gin.Context) {
	err := new(****_model.MaxiiotBillingQuotes).CreateTable()
	c.JSON(200, gin.H{
		"error": err,
	})
}
//获取列表
func GetMaxiiotBillingQuotesList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "5"))
	list := ****_model.GetMaxiiotBillingQuotesList(page, perPage, "", "created desc")
	count := ****_model.GetMaxiiotBillingQuotesCount("")
	obj := gin.H{
		"list":       list,
		"list_count": count,
	}
	c.JSON(200, obj)
}
//获取单条数据
func GetMaxiiotBillingQuotesById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	item := ****_model.GetMaxiiotBillingQuotesById(id)
	c.JSON(200, item)
}

//创建数据
func CreateMaxiiotBillingQuotes(c *gin.Context) {
	item := new(****_model.MaxiiotBillingQuotes)
	obj := item.Create()
	c.JSON(200, obj)
}
*/
