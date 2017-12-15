package gener

import (
	"../../common"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	MaxiiotUsers_TableName = ""
)

type MaxiiotUsers struct {
	Id      int64
	Name    string    `xorm:"name VARCHAR(255)"`
	Sex     int       `xorm:"sex TINYINT(1)"`
	Created time.Time `xorm:"created not null DATETIME"`
	Updated time.Time `xorm:"updated not null DATETIME"`
}

func (self *MaxiiotUsers) TableName() string {
	return MaxiiotUsers_TableName
}

//插入数据
func (self *MaxiiotUsers) Create() *MaxiiotUsers {
	_, err := common.Engine.InsertOne(self)
	if err != nil {
		log.Fatalln("Create MaxiiotUsers Err:", err)
		return nil
	}

	return self
}

//创建数据表
func (self *MaxiiotUsers) CreateTable() error {
	err := common.Engine.CreateTables(self)
	if err != nil {
		log.Fatalln("Create MaxiiotUsers Table Error:", err)
	}
	return err
}

// 删除数据
func (u *MaxiiotUsers) Delete() {
	_, err := common.Engine.Delete(u)
	if err != nil {
		log.Fatalln(" Delete MaxiiotUsers Error:", err)
	}
}

//更新MaxiiotUsers
func (self *MaxiiotUsers) Update() {
	common.Engine.Id(self.Id).Update(self)
}

//获取MaxiiotUsers列表
//Find: 获取多条数据
func GetMaxiiotUsersList(page, perPage int, whereStr, orderStr string) []*MaxiiotUsers {
	var list []*MaxiiotUsers
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
		log.Fatalln("GetMaxiiotUsersList Error:", err)
		return nil
	}
	return list
}

//MaxiiotUsers总数量
func GetMaxiiotUsersCount(whereStr string) int64 {
	sql := "id is not null"
	if whereStr != "" {
		sql += fmt.Sprintf(" %s", whereStr)
	}
	total, err := common.Engine.Where(sql).Count(new(MaxiiotUsers))
	if err != nil {
		log.Fatalln("GetMaxiiotUsersCount Error:", err)
		return 0
	}
	return total
}

//通过ID获取MaxiiotUsers
func GetMaxiiotUsersById(id int64) *MaxiiotUsers {
	item := new(MaxiiotUsers)
	_, err := common.Engine.Where("id = ?", id).Get(item)
	// has, err := x.Id(id).Get(a)
	if err != nil {
		log.Fatalln("GetMaxiiotUsersById Error:", err)
		return nil
	}
	return item
}

//通过原生sql获取
func GetMaxiiotUsersBySql(sql string) []*MaxiiotUsers {
	var list []*MaxiiotUsers
	err := common.Engine.Sql(sql).Find(&list)
	if err != nil {
		log.Fatalln("GetMaxiiotUsersBySql Error:", err)
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
func CreateMaxiiotUsersTable(c *gin.Context) {
	err := new(****_model.MaxiiotUsers).CreateTable()
	c.JSON(200, gin.H{
		"error": err,
	})
}
//获取列表
func GetMaxiiotUsersList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "5"))
	list := ****_model.GetMaxiiotUsersList(page, perPage, "", "created desc")
	count := ****_model.GetMaxiiotUsersCount("")
	obj := gin.H{
		"list":       list,
		"list_count": count,
	}
	c.JSON(200, obj)
}
//获取单条数据
func GetMaxiiotUsersById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	item := ****_model.GetMaxiiotUsersById(id)
	c.JSON(200, item)
}

//创建数据
func CreateMaxiiotUsers(c *gin.Context) {
	item := new(****_model.MaxiiotUsers)
	obj := item.Create()
	c.JSON(200, obj)
}
*/
