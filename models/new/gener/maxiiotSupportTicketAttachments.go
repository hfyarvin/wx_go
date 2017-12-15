package gener

import (
	"../../common"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	MaxiiotSupportTicketAttachments_TableName = ""
)

type MaxiiotSupportTicketAttachments struct {
	Id       int64
	Url      string    `xorm:"url not null VARCHAR(255)"`
	TicketId int64     `xorm:"ticket_id not null BIGINT(20)"`
	Created  time.Time `xorm:"created not null DATETIME"`
	Updated  time.Time `xorm:"updated not null DATETIME"`
}

func (self *MaxiiotSupportTicketAttachments) TableName() string {
	return MaxiiotSupportTicketAttachments_TableName
}

//插入数据
func (self *MaxiiotSupportTicketAttachments) Create() *MaxiiotSupportTicketAttachments {
	_, err := common.Engine.InsertOne(self)
	if err != nil {
		log.Fatalln("Create MaxiiotSupportTicketAttachments Err:", err)
		return nil
	}

	return self
}

//创建数据表
func (self *MaxiiotSupportTicketAttachments) CreateTable() error {
	err := common.Engine.CreateTables(self)
	if err != nil {
		log.Fatalln("Create MaxiiotSupportTicketAttachments Table Error:", err)
	}
	return err
}

// 删除数据
func (u *MaxiiotSupportTicketAttachments) Delete() {
	_, err := common.Engine.Delete(u)
	if err != nil {
		log.Fatalln(" Delete MaxiiotSupportTicketAttachments Error:", err)
	}
}

//更新MaxiiotSupportTicketAttachments
func (self *MaxiiotSupportTicketAttachments) Update() {
	common.Engine.Id(self.Id).Update(self)
}

//获取MaxiiotSupportTicketAttachments列表
//Find: 获取多条数据
func GetMaxiiotSupportTicketAttachmentsList(page, perPage int, whereStr, orderStr string) []*MaxiiotSupportTicketAttachments {
	var list []*MaxiiotSupportTicketAttachments
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
		log.Fatalln("GetMaxiiotSupportTicketAttachmentsList Error:", err)
		return nil
	}
	return list
}

//MaxiiotSupportTicketAttachments总数量
func GetMaxiiotSupportTicketAttachmentsCount(whereStr string) int64 {
	sql := "id is not null"
	if whereStr != "" {
		sql += fmt.Sprintf(" %s", whereStr)
	}
	total, err := common.Engine.Where(sql).Count(new(MaxiiotSupportTicketAttachments))
	if err != nil {
		log.Fatalln("GetMaxiiotSupportTicketAttachmentsCount Error:", err)
		return 0
	}
	return total
}

//通过ID获取MaxiiotSupportTicketAttachments
func GetMaxiiotSupportTicketAttachmentsById(id int64) *MaxiiotSupportTicketAttachments {
	item := new(MaxiiotSupportTicketAttachments)
	_, err := common.Engine.Where("id = ?", id).Get(item)
	// has, err := x.Id(id).Get(a)
	if err != nil {
		log.Fatalln("GetMaxiiotSupportTicketAttachmentsById Error:", err)
		return nil
	}
	return item
}

//通过原生sql获取
func GetMaxiiotSupportTicketAttachmentsBySql(sql string) []*MaxiiotSupportTicketAttachments {
	var list []*MaxiiotSupportTicketAttachments
	err := common.Engine.Sql(sql).Find(&list)
	if err != nil {
		log.Fatalln("GetMaxiiotSupportTicketAttachmentsBySql Error:", err)
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
func CreateMaxiiotSupportTicketAttachmentsTable(c *gin.Context) {
	err := new(****_model.MaxiiotSupportTicketAttachments).CreateTable()
	c.JSON(200, gin.H{
		"error": err,
	})
}
//获取列表
func GetMaxiiotSupportTicketAttachmentsList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "5"))
	list := ****_model.GetMaxiiotSupportTicketAttachmentsList(page, perPage, "", "created desc")
	count := ****_model.GetMaxiiotSupportTicketAttachmentsCount("")
	obj := gin.H{
		"list":       list,
		"list_count": count,
	}
	c.JSON(200, obj)
}
//获取单条数据
func GetMaxiiotSupportTicketAttachmentsById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	item := ****_model.GetMaxiiotSupportTicketAttachmentsById(id)
	c.JSON(200, item)
}

//创建数据
func CreateMaxiiotSupportTicketAttachments(c *gin.Context) {
	item := new(****_model.MaxiiotSupportTicketAttachments)
	obj := item.Create()
	c.JSON(200, obj)
}
*/
