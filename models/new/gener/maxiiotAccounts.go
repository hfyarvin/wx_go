package gener

import (
	"../../common"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	MaxiiotAccounts_TableName = ""
)

type MaxiiotAccounts struct {
	Id      int64
	Name    string    `xorm:"name VARCHAR(255)"`
	Aaaa    time.Time `xorm:"aaaa TIMESTAMP"`
	Balance float64   `xorm:"balance not null DOUBLE(10)"`
	Version int       `xorm:"version default 1 INT(11)"`
	Created time.Time `xorm:"created not null DATETIME(4)"`
	Updated time.Time `xorm:"updated not null DATETIME"`
}

func (self *MaxiiotAccounts) TableName() string {
	return MaxiiotAccounts_TableName
}

//插入数据
func (self *MaxiiotAccounts) Create() *MaxiiotAccounts {
	_, err := common.Engine.InsertOne(self)
	if err != nil {
		log.Fatalln("Create MaxiiotAccounts Err:", err)
		return nil
	}

	return self
}

//创建数据表
func (self *MaxiiotAccounts) CreateTable() error {
	err := common.Engine.CreateTables(self)
	if err != nil {
		log.Fatalln("Create MaxiiotAccounts Table Error:", err)
	}
	return err
}

// 删除数据
func (u *MaxiiotAccounts) Delete() {
	_, err := common.Engine.Delete(u)
	if err != nil {
		log.Fatalln(" Delete MaxiiotAccounts Error:", err)
	}
}

//更新MaxiiotAccounts
func (self *MaxiiotAccounts) Update() {
	common.Engine.Id(self.Id).Update(self)
}

//获取MaxiiotAccounts列表
//Find: 获取多条数据
func GetMaxiiotAccountsList(page, perPage int, whereStr, orderStr string) []*MaxiiotAccounts {
	var list []*MaxiiotAccounts
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
		log.Fatalln("GetMaxiiotAccountsList Error:", err)
		return nil
	}
	return list
}

//MaxiiotAccounts总数量
func GetMaxiiotAccountsCount(whereStr string) int64 {
	sql := "id is not null"
	if whereStr != "" {
		sql += fmt.Sprintf(" %s", whereStr)
	}
	total, err := common.Engine.Where(sql).Count(new(MaxiiotAccounts))
	if err != nil {
		log.Fatalln("GetMaxiiotAccountsCount Error:", err)
		return 0
	}
	return total
}

//通过ID获取MaxiiotAccounts
func GetMaxiiotAccountsById(id int64) *MaxiiotAccounts {
	item := new(MaxiiotAccounts)
	_, err := common.Engine.Where("id = ?", id).Get(item)
	// has, err := x.Id(id).Get(a)
	if err != nil {
		log.Fatalln("GetMaxiiotAccountsById Error:", err)
		return nil
	}
	return item
}

//通过原生sql获取
func GetMaxiiotAccountsBySql(sql string) []*MaxiiotAccounts {
	var list []*MaxiiotAccounts
	err := common.Engine.Sql(sql).Find(&list)
	if err != nil {
		log.Fatalln("GetMaxiiotAccountsBySql Error:", err)
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
func CreateMaxiiotAccountsTable(c *gin.Context) {
	err := new(****_model.MaxiiotAccounts).CreateTable()
	c.JSON(200, gin.H{
		"error": err,
	})
}
//获取列表
func GetMaxiiotAccountsList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "5"))
	list := ****_model.GetMaxiiotAccountsList(page, perPage, "", "created desc")
	count := ****_model.GetMaxiiotAccountsCount("")
	obj := gin.H{
		"list":       list,
		"list_count": count,
	}
	c.JSON(200, obj)
}
//获取单条数据
func GetMaxiiotAccountsById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	item := ****_model.GetMaxiiotAccountsById(id)
	c.JSON(200, item)
}

//创建数据
func CreateMaxiiotAccounts(c *gin.Context) {
	item := new(****_model.MaxiiotAccounts)
	obj := item.Create()
	c.JSON(200, obj)
}
*/
