package gener

import (
	"fmt"
	"github.com/maxiiot/LoRaWan/common"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	MaxiiotSupportArticleCats_TableName = ""
)

type MaxiiotSupportArticleCats struct {
	Id      int64
	Titlt   string    `xorm:"titlt VARCHAR(255)"`
	Created time.Time `xorm:"created not null DATETIME"`
	Updated time.Time `xorm:"updated not null DATETIME"`
}

func (self *MaxiiotSupportArticleCats) TableName() string {
	return MaxiiotSupportArticleCats_TableName
}

//插入数据
func (self *MaxiiotSupportArticleCats) Create() *MaxiiotSupportArticleCats {
	_, err := common.Engine.InsertOne(self)
	if err != nil {
		log.Fatalln("Create MaxiiotSupportArticleCats Err:", err)
		return nil
	}

	return self
}

//创建数据表
func (self *MaxiiotSupportArticleCats) CreateTable() error {
	err := common.Engine.CreateTables(self)
	if err != nil {
		log.Fatalln("Create MaxiiotSupportArticleCats Table Error:", err)
	}
	return err
}

// 删除数据
func (u *MaxiiotSupportArticleCats) Delete() {
	_, err := common.Engine.Delete(u)
	if err != nil {
		log.Fatalln(" Delete MaxiiotSupportArticleCats Error:", err)
	}
}

//更新MaxiiotSupportArticleCats
func (self *MaxiiotSupportArticleCats) Update() {
	common.Engine.Id(self.Id).Update(self)
}

//获取MaxiiotSupportArticleCats列表
//Find: 获取多条数据
func GetMaxiiotSupportArticleCatsList(page, perPage int, whereStr, orderStr string) []*MaxiiotSupportArticleCats {
	var list []*MaxiiotSupportArticleCats
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
		log.Fatalln("GetMaxiiotSupportArticleCatsList Error:", err)
		return nil
	}
	return list
}

//MaxiiotSupportArticleCats总数量
func GetMaxiiotSupportArticleCatsCount(whereStr string) int64 {
	sql := "id is not null"
	if whereStr != "" {
		sql += fmt.Sprintf(" %s", whereStr)
	}
	total, err := common.Engine.Where(sql).Count(new(MaxiiotSupportArticleCats))
	if err != nil {
		log.Fatalln("GetMaxiiotSupportArticleCatsCount Error:", err)
		return 0
	}
	return total
}

//通过ID获取MaxiiotSupportArticleCats
func GetMaxiiotSupportArticleCatsById(id int64) *MaxiiotSupportArticleCats {
	item := new(MaxiiotSupportArticleCats)
	_, err := common.Engine.Where("id = ?", id).Get(item)
	// has, err := x.Id(id).Get(a)
	if err != nil {
		log.Fatalln("GetMaxiiotSupportArticleCatsById Error:", err)
		return nil
	}
	return item
}

//通过原生sql获取
func GetMaxiiotSupportArticleCatsBySql(sql string) []*MaxiiotSupportArticleCats {
	var list []*MaxiiotSupportArticleCats
	err := common.Engine.Sql(sql).Find(&list)
	if err != nil {
		log.Fatalln("GetMaxiiotSupportArticleCatsBySql Error:", err)
		return nil
	}
	return list
}

/*controller方法
package ****_controller

import (
	"github.com/maxiiot/LoRaWan/models/****_model"
	"net/http"
	"github.com/maxiiot/LoRaWan/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

//创建数据表
func CreateMaxiiotSupportArticleCatsTable(c *gin.Context) {
	err := new(****_model.MaxiiotSupportArticleCats).CreateTable()
	obj := gin.H{
		"error": err,
	}
	controllers.ResponseJSON(c, http.StatusOK, "sucess", obj)
}
//获取列表
func GetMaxiiotSupportArticleCatsList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "5"))
	list := ****_model.GetMaxiiotSupportArticleCatsList(page, perPage, "", "created desc")
	count := ****_model.GetMaxiiotSupportArticleCatsCount("")
	obj := gin.H{
		"list":       list,
		"list_count": count,
	}
	controllers.ResponseJSON(c, http.StatusOK, "sucess", obj)
}

//获取单条数据
func GetMaxiiotSupportArticleCatsById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	obj := ****_model.GetMaxiiotSupportArticleCatsById(id)
	controllers.ResponseJSON(c, http.StatusOK, "sucess", obj)
}

//创建数据
func CreateMaxiiotSupportArticleCats(c *gin.Context) {
	item := new(****_model.MaxiiotSupportArticleCats)
	obj := item.Create()
	controllers.ResponseJSON(c, http.StatusOK, "sucess", obj)
}

    "/list": {
      "get": {
        "summary": "获取列表",
        "operationId": "MaxiiotSupportArticleCatsList",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/apiEmptyResponse"
            }
          }
        },
        "parameters": [
          {
              "description":"列表页序号.",
              "format":"int32",
              "in":"query",
              "name":"page",
              "required":false,
              "type":"integer"
          },
          {
              "description":"每页展示数量.",
              "format":"int32",
              "in":"query",
              "name":"per_page",
              "required":false,
              "type":"integer"
          }
        ],
        "tags": [
          "Support"
        ]
      }
    },
    "/id/{id}": {
      "get": {
        "summary": "获取单个服务单",
        "operationId": "GetMaxiiotSupportArticleCatsById",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/apiEmptyResponse"
            }
          }
        },
        "parameters": [
          {
              "format":"int64",
              "in":"path",
              "name":"id",
              "required":true,
              "type":"string"
          }
        ],
        "tags": [
          ".."
        ]
      }
    },
    "/new": {
      "post": {
        "summary": "创建新服务单",
        "operationId": "CreateMaxiiotSupportArticleCats",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/apiEmptyResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/apiEmptyRequest"
            }
          }
        ],
        "tags": [
          "***"
        ]
      }
    }
*/
