package gener

import (
	"../../common"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	MaxiiotBillingInvoices_TableName = ""
)

type MaxiiotBillingInvoices struct {
	Id         int64
	UserId     int64     `xorm:"user_id not null BIGINT(20)"`
	Recipient  string    `xorm:"recipient VARCHAR(255)"`
	TotalPrice int64     `xorm:"total_price not null BIGINT(20)"`
	Status     int       `xorm:"status default 1 TINYINT(1)"`
	PayWay     string    `xorm:"pay_way VARCHAR(255)"`
	OrderId    string    `xorm:"order_id VARCHAR(255)"`
	Expired    time.Time `xorm:"expired not null DATETIME"`
	Remark     string    `xorm:"remark VARCHAR(255)"`
	Created    time.Time `xorm:"created not null DATETIME"`
	Updated    time.Time `xorm:"updated not null DATETIME"`
}

func (self *MaxiiotBillingInvoices) TableName() string {
	return MaxiiotBillingInvoices_TableName
}

//插入数据
func (self *MaxiiotBillingInvoices) Create() *MaxiiotBillingInvoices {
	_, err := common.Engine.InsertOne(self)
	if err != nil {
		log.Fatalln("Create MaxiiotBillingInvoices Err:", err)
		return nil
	}

	return self
}

//创建数据表
func (self *MaxiiotBillingInvoices) CreateTable() error {
	err := common.Engine.CreateTables(self)
	if err != nil {
		log.Fatalln("Create MaxiiotBillingInvoices Table Error:", err)
	}
	return err
}

// 删除数据
func (u *MaxiiotBillingInvoices) Delete() {
	_, err := common.Engine.Delete(u)
	if err != nil {
		log.Fatalln(" Delete MaxiiotBillingInvoices Error:", err)
	}
}

//更新MaxiiotBillingInvoices
func (self *MaxiiotBillingInvoices) Update() {
	common.Engine.Id(self.Id).Update(self)
}

//获取MaxiiotBillingInvoices列表
//Find: 获取多条数据
func GetMaxiiotBillingInvoicesList(page, perPage int, whereStr, orderStr string) []*MaxiiotBillingInvoices {
	var list []*MaxiiotBillingInvoices
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
		log.Fatalln("GetMaxiiotBillingInvoicesList Error:", err)
		return nil
	}
	return list
}

//MaxiiotBillingInvoices总数量
func GetMaxiiotBillingInvoicesCount(whereStr string) int64 {
	sql := "id is not null"
	if whereStr != "" {
		sql += fmt.Sprintf(" %s", whereStr)
	}
	total, err := common.Engine.Where(sql).Count(new(MaxiiotBillingInvoices))
	if err != nil {
		log.Fatalln("GetMaxiiotBillingInvoicesCount Error:", err)
		return 0
	}
	return total
}

//通过ID获取MaxiiotBillingInvoices
func GetMaxiiotBillingInvoicesById(id int64) *MaxiiotBillingInvoices {
	item := new(MaxiiotBillingInvoices)
	_, err := common.Engine.Where("id = ?", id).Get(item)
	// has, err := x.Id(id).Get(a)
	if err != nil {
		log.Fatalln("GetMaxiiotBillingInvoicesById Error:", err)
		return nil
	}
	return item
}

//通过原生sql获取
func GetMaxiiotBillingInvoicesBySql(sql string) []*MaxiiotBillingInvoices {
	var list []*MaxiiotBillingInvoices
	err := common.Engine.Sql(sql).Find(&list)
	if err != nil {
		log.Fatalln("GetMaxiiotBillingInvoicesBySql Error:", err)
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
func CreateMaxiiotBillingInvoicesTable(c *gin.Context) {
	err := new(****_model.MaxiiotBillingInvoices).CreateTable()
	c.JSON(200, gin.H{
		"error": err,
	})
}
//获取列表
func GetMaxiiotBillingInvoicesList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "5"))
	list := ****_model.GetMaxiiotBillingInvoicesList(page, perPage, "", "created desc")
	count := ****_model.GetMaxiiotBillingInvoicesCount("")
	obj := gin.H{
		"list":       list,
		"list_count": count,
	}
	c.JSON(200, obj)
}
//获取单条数据
func GetMaxiiotBillingInvoicesById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	item := ****_model.GetMaxiiotBillingInvoicesById(id)
	c.JSON(200, item)
}

//创建数据
func CreateMaxiiotBillingInvoices(c *gin.Context) {
	item := new(****_model.MaxiiotBillingInvoices)
	obj := item.Create()
	c.JSON(200, obj)
}
*/
