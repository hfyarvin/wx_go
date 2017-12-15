package gener

import (
	"../../common"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	MaxiiotSupportArticles_TableName = ""
)

type MaxiiotSupportArticles struct {
	Id           int64
	ArticleCatId int64     `xorm:"article_cat_id BIGINT(20)"`
	Title        string    `xorm:"title VARCHAR(255)"`
	Content      string    `xorm:"content TEXT"`
	ClickTimes   int       `xorm:"click_times default 0 INT(11)"`
	Created      time.Time `xorm:"created not null DATETIME"`
	Updated      time.Time `xorm:"updated not null DATETIME"`
}

func (self *MaxiiotSupportArticles) TableName() string {
	return MaxiiotSupportArticles_TableName
}

//插入数据
func (self *MaxiiotSupportArticles) Create() *MaxiiotSupportArticles {
	_, err := common.Engine.InsertOne(self)
	if err != nil {
		log.Fatalln("Create MaxiiotSupportArticles Err:", err)
		return nil
	}

	return self
}

//创建数据表
func (self *MaxiiotSupportArticles) CreateTable() error {
	err := common.Engine.CreateTables(self)
	if err != nil {
		log.Fatalln("Create MaxiiotSupportArticles Table Error:", err)
	}
	return err
}

// 删除数据
func (u *MaxiiotSupportArticles) Delete() {
	_, err := common.Engine.Delete(u)
	if err != nil {
		log.Fatalln(" Delete MaxiiotSupportArticles Error:", err)
	}
}

//更新MaxiiotSupportArticles
func (self *MaxiiotSupportArticles) Update() {
	common.Engine.Id(self.Id).Update(self)
}

//获取MaxiiotSupportArticles列表
//Find: 获取多条数据
func GetMaxiiotSupportArticlesList(page, perPage int, whereStr, orderStr string) []*MaxiiotSupportArticles {
	var list []*MaxiiotSupportArticles
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
		log.Fatalln("GetMaxiiotSupportArticlesList Error:", err)
		return nil
	}
	return list
}

//MaxiiotSupportArticles总数量
func GetMaxiiotSupportArticlesCount(whereStr string) int64 {
	sql := "id is not null"
	if whereStr != "" {
		sql += fmt.Sprintf(" %s", whereStr)
	}
	total, err := common.Engine.Where(sql).Count(new(MaxiiotSupportArticles))
	if err != nil {
		log.Fatalln("GetMaxiiotSupportArticlesCount Error:", err)
		return 0
	}
	return total
}

//通过ID获取MaxiiotSupportArticles
func GetMaxiiotSupportArticlesById(id int64) *MaxiiotSupportArticles {
	item := new(MaxiiotSupportArticles)
	_, err := common.Engine.Where("id = ?", id).Get(item)
	// has, err := x.Id(id).Get(a)
	if err != nil {
		log.Fatalln("GetMaxiiotSupportArticlesById Error:", err)
		return nil
	}
	return item
}

//通过原生sql获取
func GetMaxiiotSupportArticlesBySql(sql string) []*MaxiiotSupportArticles {
	var list []*MaxiiotSupportArticles
	err := common.Engine.Sql(sql).Find(&list)
	if err != nil {
		log.Fatalln("GetMaxiiotSupportArticlesBySql Error:", err)
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
func CreateMaxiiotSupportArticlesTable(c *gin.Context) {
	err := new(****_model.MaxiiotSupportArticles).CreateTable()
	c.JSON(200, gin.H{
		"error": err,
	})
}
//获取列表
func GetMaxiiotSupportArticlesList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "5"))
	list := ****_model.GetMaxiiotSupportArticlesList(page, perPage, "", "created desc")
	count := ****_model.GetMaxiiotSupportArticlesCount("")
	obj := gin.H{
		"list":       list,
		"list_count": count,
	}
	c.JSON(200, obj)
}
//获取单条数据
func GetMaxiiotSupportArticlesById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	item := ****_model.GetMaxiiotSupportArticlesById(id)
	c.JSON(200, item)
}

//创建数据
func CreateMaxiiotSupportArticles(c *gin.Context) {
	item := new(****_model.MaxiiotSupportArticles)
	obj := item.Create()
	c.JSON(200, obj)
}
*/
