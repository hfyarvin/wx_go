package gener

import (
	"../../common"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	Users_TableName = ""
)

type Users struct {
	Id           int64
	FirstName    string    `xorm:"first_name not null VARCHAR(50)"`
	LastName     string    `xorm:"last_name not null VARCHAR(50)"`
	PasswordHash string    `xorm:"password_hash not null VARCHAR(200)"`
	Email        string    `xorm:"email not null unique VARCHAR(100)"`
	Phone        string    `xorm:"phone not null VARCHAR(30)"`
	Company      string    `xorm:"company VARCHAR(200)"`
	Address      string    `xorm:"address VARCHAR(400)"`
	City         string    `xorm:"city VARCHAR(100)"`
	Regin        string    `xorm:"regin VARCHAR(100)"`
	ZipCode      string    `xorm:"zip_code VARCHAR(30)"`
	Country      string    `xorm:"country VARCHAR(100)"`
	IdentityCard string    `xorm:"identity_card VARCHAR(50)"`
	IsActive     int       `xorm:"is_active BIT(1)"`
	Created      time.Time `xorm:"created not null DATETIME"`
	Updated      time.Time `xorm:"updated not null DATETIME"`
}

func (self *Users) TableName() string {
	return Users_TableName
}

//插入数据
func (self *Users) Create() *Users {
	_, err := common.Engine.InsertOne(self)
	if err != nil {
		log.Fatalln("Create Users Err:", err)
		return nil
	}

	return self
}

//创建数据表
func (self *Users) CreateTable() error {
	err := common.Engine.CreateTables(self)
	if err != nil {
		log.Fatalln("Create Users Table Error:", err)
	}
	return err
}

// 删除数据
func (u *Users) Delete() {
	_, err := common.Engine.Delete(u)
	if err != nil {
		log.Fatalln(" Delete Users Error:", err)
	}
}

//更新Users
func (self *Users) Update() {
	common.Engine.Id(self.Id).Update(self)
}

//获取Users列表
//Find: 获取多条数据
func GetUsersList(page, perPage int, whereStr, orderStr string) []*Users {
	var list []*Users
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
		log.Fatalln("GetUsersList Error:", err)
		return nil
	}
	return list
}

//Users总数量
func GetUsersCount(whereStr string) int64 {
	sql := "id is not null"
	if whereStr != "" {
		sql += fmt.Sprintf(" %s", whereStr)
	}
	total, err := common.Engine.Where(sql).Count(new(Users))
	if err != nil {
		log.Fatalln("GetUsersCount Error:", err)
		return 0
	}
	return total
}

//通过ID获取Users
func GetUsersById(id int64) *Users {
	item := new(Users)
	_, err := common.Engine.Where("id = ?", id).Get(item)
	// has, err := x.Id(id).Get(a)
	if err != nil {
		log.Fatalln("GetUsersById Error:", err)
		return nil
	}
	return item
}

//通过原生sql获取
func GetUsersBySql(sql string) []*Users {
	var list []*Users
	err := common.Engine.Sql(sql).Find(&list)
	if err != nil {
		log.Fatalln("GetUsersBySql Error:", err)
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
func CreateUsersTable(c *gin.Context) {
	err := new(****_model.Users).CreateTable()
	c.JSON(200, gin.H{
		"error": err,
	})
}
//获取列表
func GetUsersList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "5"))
	list := ****_model.GetUsersList(page, perPage, "", "created desc")
	count := ****_model.GetUsersCount("")
	obj := gin.H{
		"list":       list,
		"list_count": count,
	}
	c.JSON(200, obj)
}
//获取单条数据
func GetUsersById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	item := ****_model.GetUsersById(id)
	c.JSON(200, item)
}

//创建数据
func CreateUsers(c *gin.Context) {
	item := new(****_model.Users)
	obj := item.Create()
	c.JSON(200, obj)
}
*/
