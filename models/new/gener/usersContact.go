package gener

const (
	UsersContact_TableName = ""
)

type UsersContact struct {
	Id        int64
	UserId    int64  `xorm:"user_id not null BIGINT(20)"`
	FirstName string `xorm:"first_name not null VARCHAR(100)"`
	LastName  string `xorm:"last_name not null VARCHAR(100)"`
	Company   string `xorm:"company not null VARCHAR(200)"`
	Email     string `xorm:"email not null VARCHAR(50)"`
	Phone     string `xorm:"phone VARCHAR(30)"`
	Address1  string `xorm:"address1 VARCHAR(200)"`
	Address2  string `xorm:"address2 VARCHAR(200)"`
	City      string `xorm:"city VARCHAR(50)"`
	Region    string `xorm:"region VARCHAR(100)"`
	ZipCode   string `xorm:"zip_code VARCHAR(30)"`
	Country   string `xorm:"country VARCHAR(100)"`
}

func (self *UsersContact) TableName() string {
	return UsersContact_TableName
}

//插入数据
func (self *UsersContact) Create() *UsersContact {
	_, err := common.Engine.InsertOne(self)
	if err != nil {
		log.Fatalln("Create UsersContact Err:", err)
		return nil
	}

	return self
}

//创建数据表
func (self *UsersContact) CreateTable() error {
	err := common.Engine.CreateTables(self)
	if err != nil {
		log.Fatalln("Create UsersContact Table Error:", err)
	}
	return err
}

// 删除数据
func (u *UsersContact) Delete() {
	_, err := common.Engine.Delete(u)
	if err != nil {
		log.Fatalln(" Delete UsersContact Error:", err)
	}
}

//更新UsersContact
func (self *UsersContact) Update() {
	common.Engine.Id(self.Id).Update(self)
}

//获取UsersContact列表
//Find: 获取多条数据
func GetUsersContactList(page, perPage int, whereStr, orderStr string) []*UsersContact {
	var list []*UsersContact
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
		log.Fatalln("GetUsersContactList Error:", err)
		return nil
	}
	return list
}

//UsersContact总数量
func GetUsersContactCount(whereStr string) int64 {
	sql := "id is not null"
	if whereStr != "" {
		sql += fmt.Sprintf(" %s", whereStr)
	}
	total, err := common.Engine.Where(sql).Count(new(UsersContact))
	if err != nil {
		log.Fatalln("GetUsersContactCount Error:", err)
		return 0
	}
	return total
}

//通过ID获取UsersContact
func GetUsersContactById(id int64) *UsersContact {
	item := new(UsersContact)
	_, err := common.Engine.Where("id = ?", id).Get(item)
	// has, err := x.Id(id).Get(a)
	if err != nil {
		log.Fatalln("GetUsersContactById Error:", err)
		return nil
	}
	return item
}

//通过原生sql获取
func GetUsersContactBySql(sql string) []*UsersContact {
	var list []*UsersContact
	err := common.Engine.Sql(sql).Find(&list)
	if err != nil {
		log.Fatalln("GetUsersContactBySql Error:", err)
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
func CreateUsersContactTable(c *gin.Context) {
	err := new(****_model.UsersContact).CreateTable()
	c.JSON(200, gin.H{
		"error": err,
	})
}
//获取列表
func GetUsersContactList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "5"))
	list := ****_model.GetUsersContactList(page, perPage, "", "created desc")
	count := ****_model.GetUsersContactCount("")
	obj := gin.H{
		"list":       list,
		"list_count": count,
	}
	c.JSON(200, obj)
}
//获取单条数据
func GetUsersContactById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	item := ****_model.GetUsersContactById(id)
	c.JSON(200, item)
}

//创建数据
func CreateUsersContact(c *gin.Context) {
	item := new(****_model.UsersContact)
	obj := item.Create()
	c.JSON(200, obj)
}
*/
