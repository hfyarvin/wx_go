package gener

import (
	"../../common"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	MaxiiotSupportAnnouncements_TableName = ""
)

type MaxiiotSupportAnnouncements struct {
	Id         int64
	Title      string    `xorm:"title not null default 'ss' VARCHAR(100)"`
	Content    string    `xorm:"content not null TEXT"`
	ClickTimes int       `xorm:"click_times default 1 INT(11)"`
	Created    time.Time `xorm:"created not null DATETIME"`
	Updated    time.Time `xorm:"updated not null DATETIME"`
}

func (self *MaxiiotSupportAnnouncements) TableName() string {
	return MaxiiotSupportAnnouncements_TableName
}

//插入数据
func (self *MaxiiotSupportAnnouncements) Create() *MaxiiotSupportAnnouncements {
	_, err := common.Engine.InsertOne(self)
	if err != nil {
		log.Fatalln("Create MaxiiotSupportAnnouncements Err:", err)
		return nil
	}

	return self
}

//创建数据表
func (self *MaxiiotSupportAnnouncements) CreateTable() error {
	err := common.Engine.CreateTables(self)
	if err != nil {
		log.Fatalln("Create MaxiiotSupportAnnouncements Table Error:", err)
	}
	return err
}

// 删除数据
func (u *MaxiiotSupportAnnouncements) Delete() {
	_, err := common.Engine.Delete(u)
	if err != nil {
		log.Fatalln(" Delete MaxiiotSupportAnnouncements Error:", err)
	}
}

//更新MaxiiotSupportAnnouncements
func (self *MaxiiotSupportAnnouncements) Update() {
	common.Engine.Id(self.Id).Update(self)
}

//获取MaxiiotSupportAnnouncements列表
//Find: 获取多条数据
func GetMaxiiotSupportAnnouncementsList(page, perPage int, whereStr, orderStr string) []*MaxiiotSupportAnnouncements {
	var list []*MaxiiotSupportAnnouncements
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
		log.Fatalln("GetMaxiiotSupportAnnouncementsList Error:", err)
		return nil
	}
	return list
}

//MaxiiotSupportAnnouncements总数量
func GetMaxiiotSupportAnnouncementsCount(whereStr string) int64 {
	sql := "id is not null"
	if whereStr != "" {
		sql += fmt.Sprintf(" %s", whereStr)
	}
	total, err := common.Engine.Where(sql).Count(new(MaxiiotSupportAnnouncements))
	if err != nil {
		log.Fatalln("GetMaxiiotSupportAnnouncementsCount Error:", err)
		return 0
	}
	return total
}

//通过ID获取MaxiiotSupportAnnouncements
func GetMaxiiotSupportAnnouncementsById(id int64) *MaxiiotSupportAnnouncements {
	item := new(MaxiiotSupportAnnouncements)
	_, err := common.Engine.Where("id = ?", id).Get(item)
	// has, err := x.Id(id).Get(a)
	if err != nil {
		log.Fatalln("GetMaxiiotSupportAnnouncementsById Error:", err)
		return nil
	}
	return item
}

//通过原生sql获取
func GetMaxiiotSupportAnnouncementsBySql(sql string) []*MaxiiotSupportAnnouncements {
	var list []*MaxiiotSupportAnnouncements
	err := common.Engine.Sql(sql).Find(&list)
	if err != nil {
		log.Fatalln("GetMaxiiotSupportAnnouncementsBySql Error:", err)
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
func CreateMaxiiotSupportAnnouncementsTable(c *gin.Context) {
	err := new(****_model.MaxiiotSupportAnnouncements).CreateTable()
	c.JSON(200, gin.H{
		"error": err,
	})
}
//获取列表
func GetMaxiiotSupportAnnouncementsList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "5"))
	list := ****_model.GetMaxiiotSupportAnnouncementsList(page, perPage, "", "created desc")
	count := ****_model.GetMaxiiotSupportAnnouncementsCount("")
	obj := gin.H{
		"list":       list,
		"list_count": count,
	}
	c.JSON(200, obj)
}
//获取单条数据
func GetMaxiiotSupportAnnouncementsById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	item := ****_model.GetMaxiiotSupportAnnouncementsById(id)
	c.JSON(200, item)
}

//创建数据
func CreateMaxiiotSupportAnnouncements(c *gin.Context) {
	item := new(****_model.MaxiiotSupportAnnouncements)
	obj := item.Create()
	c.JSON(200, obj)
}
*/
