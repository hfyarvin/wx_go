package gener

import (
	"../../common"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	MaxiiotSupportTickets_TableName = ""
)

type MaxiiotSupportTickets struct {
	Id               int64
	UserId           int64     `xorm:"user_id not null BIGINT(20)"`
	Department       string    `xorm:"department not null VARCHAR(100)"`
	Subject          string    `xorm:"subject VARCHAR(255)"`
	Status           string    `xorm:"status VARCHAR(255)"`
	RelatedServiceId int64     `xorm:"related_service_id BIGINT(20)"`
	Priority         int       `xorm:"priority TINYINT(1)"`
	Message          string    `xorm:"message TEXT"`
	Attachments      string    `xorm:"attachments VARCHAR(255)"`
	Created          time.Time `xorm:"created not null DATETIME"`
	Updated          time.Time `xorm:"updated not null DATETIME"`
}

func (self *MaxiiotSupportTickets) TableName() string {
	return MaxiiotSupportTickets_TableName
}

//插入数据
func (self *MaxiiotSupportTickets) Create() *MaxiiotSupportTickets {
	_, err := common.Engine.InsertOne(self)
	if err != nil {
		log.Fatalln("Create MaxiiotSupportTickets Err:", err)
		return nil
	}

	return self
}

//创建数据表
func (self *MaxiiotSupportTickets) CreateTable() error {
	err := common.Engine.CreateTables(self)
	if err != nil {
		log.Fatalln("Create MaxiiotSupportTickets Table Error:", err)
	}
	return err
}

// 删除数据
func (u *MaxiiotSupportTickets) Delete() {
	_, err := common.Engine.Delete(u)
	if err != nil {
		log.Fatalln(" Delete MaxiiotSupportTickets Error:", err)
	}
}

//更新MaxiiotSupportTickets
func (self *MaxiiotSupportTickets) Update() {
	common.Engine.Id(self.Id).Update(self)
}

//获取MaxiiotSupportTickets列表
//Find: 获取多条数据
func GetMaxiiotSupportTicketsList(page, perPage int, whereStr, orderStr string) []*MaxiiotSupportTickets {
	var list []*MaxiiotSupportTickets
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
		log.Fatalln("GetMaxiiotSupportTicketsList Error:", err)
		return nil
	}
	return list
}

//MaxiiotSupportTickets总数量
func GetMaxiiotSupportTicketsCount(whereStr string) int64 {
	sql := "id is not null"
	if whereStr != "" {
		sql += fmt.Sprintf(" %s", whereStr)
	}
	total, err := common.Engine.Where(sql).Count(new(MaxiiotSupportTickets))
	if err != nil {
		log.Fatalln("GetMaxiiotSupportTicketsCount Error:", err)
		return 0
	}
	return total
}

//通过ID获取MaxiiotSupportTickets
func GetMaxiiotSupportTicketsById(id int64) *MaxiiotSupportTickets {
	item := new(MaxiiotSupportTickets)
	_, err := common.Engine.Where("id = ?", id).Get(item)
	// has, err := x.Id(id).Get(a)
	if err != nil {
		log.Fatalln("GetMaxiiotSupportTicketsById Error:", err)
		return nil
	}
	return item
}

//通过原生sql获取
func GetMaxiiotSupportTicketsBySql(sql string) []*MaxiiotSupportTickets {
	var list []*MaxiiotSupportTickets
	err := common.Engine.Sql(sql).Find(&list)
	if err != nil {
		log.Fatalln("GetMaxiiotSupportTicketsBySql Error:", err)
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
func CreateMaxiiotSupportTicketsTable(c *gin.Context) {
	err := new(****_model.MaxiiotSupportTickets).CreateTable()
	c.JSON(200, gin.H{
		"error": err,
	})
}
//获取列表
func GetMaxiiotSupportTicketsList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "5"))
	list := ****_model.GetMaxiiotSupportTicketsList(page, perPage, "", "created desc")
	count := ****_model.GetMaxiiotSupportTicketsCount("")
	obj := gin.H{
		"list":       list,
		"list_count": count,
	}
	c.JSON(200, obj)
}
//获取单条数据
func GetMaxiiotSupportTicketsById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	item := ****_model.GetMaxiiotSupportTicketsById(id)
	c.JSON(200, item)
}

//创建数据
func CreateMaxiiotSupportTickets(c *gin.Context) {
	item := new(****_model.MaxiiotSupportTickets)
	obj := item.Create()
	c.JSON(200, obj)
}
*/
