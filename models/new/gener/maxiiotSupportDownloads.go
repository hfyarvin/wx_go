package gener

import (
	"../../common"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	MaxiiotSupportDownloads_TableName = ""
)

type MaxiiotSupportDownloads struct {
	Id      int64
	Url     string    `xorm:"url VARCHAR(255)"`
	Title   string    `xorm:"title VARCHAR(255)"`
	Created time.Time `xorm:"created not null DATETIME"`
	Updated time.Time `xorm:"updated not null DATETIME"`
}

func (self *MaxiiotSupportDownloads) TableName() string {
	return MaxiiotSupportDownloads_TableName
}

//插入数据
func (self *MaxiiotSupportDownloads) Create() *MaxiiotSupportDownloads {
	_, err := common.Engine.InsertOne(self)
	if err != nil {
		log.Fatalln("Create MaxiiotSupportDownloads Err:", err)
		return nil
	}

	return self
}

//创建数据表
func (self *MaxiiotSupportDownloads) CreateTable() error {
	err := common.Engine.CreateTables(self)
	if err != nil {
		log.Fatalln("Create MaxiiotSupportDownloads Table Error:", err)
	}
	return err
}

// 删除数据
func (u *MaxiiotSupportDownloads) Delete() {
	_, err := common.Engine.Delete(u)
	if err != nil {
		log.Fatalln(" Delete MaxiiotSupportDownloads Error:", err)
	}
}

//更新MaxiiotSupportDownloads
func (self *MaxiiotSupportDownloads) Update() {
	common.Engine.Id(self.Id).Update(self)
}

//获取MaxiiotSupportDownloads列表
//Find: 获取多条数据
func GetMaxiiotSupportDownloadsList(page, perPage int, whereStr, orderStr string) []*MaxiiotSupportDownloads {
	var list []*MaxiiotSupportDownloads
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
		log.Fatalln("GetMaxiiotSupportDownloadsList Error:", err)
		return nil
	}
	return list
}

//MaxiiotSupportDownloads总数量
func GetMaxiiotSupportDownloadsCount(whereStr string) int64 {
	sql := "id is not null"
	if whereStr != "" {
		sql += fmt.Sprintf(" %s", whereStr)
	}
	total, err := common.Engine.Where(sql).Count(new(MaxiiotSupportDownloads))
	if err != nil {
		log.Fatalln("GetMaxiiotSupportDownloadsCount Error:", err)
		return 0
	}
	return total
}

//通过ID获取MaxiiotSupportDownloads
func GetMaxiiotSupportDownloadsById(id int64) *MaxiiotSupportDownloads {
	item := new(MaxiiotSupportDownloads)
	_, err := common.Engine.Where("id = ?", id).Get(item)
	// has, err := x.Id(id).Get(a)
	if err != nil {
		log.Fatalln("GetMaxiiotSupportDownloadsById Error:", err)
		return nil
	}
	return item
}

//通过原生sql获取
func GetMaxiiotSupportDownloadsBySql(sql string) []*MaxiiotSupportDownloads {
	var list []*MaxiiotSupportDownloads
	err := common.Engine.Sql(sql).Find(&list)
	if err != nil {
		log.Fatalln("GetMaxiiotSupportDownloadsBySql Error:", err)
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
func CreateMaxiiotSupportDownloadsTable(c *gin.Context) {
	err := new(****_model.MaxiiotSupportDownloads).CreateTable()
	c.JSON(200, gin.H{
		"error": err,
	})
}
//获取列表
func GetMaxiiotSupportDownloadsList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "5"))
	list := ****_model.GetMaxiiotSupportDownloadsList(page, perPage, "", "created desc")
	count := ****_model.GetMaxiiotSupportDownloadsCount("")
	obj := gin.H{
		"list":       list,
		"list_count": count,
	}
	c.JSON(200, obj)
}
//获取单条数据
func GetMaxiiotSupportDownloadsById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	item := ****_model.GetMaxiiotSupportDownloadsById(id)
	c.JSON(200, item)
}

//创建数据
func CreateMaxiiotSupportDownloads(c *gin.Context) {
	item := new(****_model.MaxiiotSupportDownloads)
	obj := item.Create()
	c.JSON(200, obj)
}
*/
