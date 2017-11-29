package gener

import (
	"fmt"
	"github.com/maxiiot/LoRaWan/common"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	MaxiiotBalances_TableName = ""
)

type MaxiiotBalances struct {
	Id      int64
	OrderId string    `xorm:"order_id not null VARCHAR(255)"`
	Amount  int64     `xorm:"amount BIGINT(20)"`
	Status  int       `xorm:"status default 0 TINYINT(4)"`
	Created time.Time `xorm:"created DATETIME"`
	Updated time.Time `xorm:"updated DATETIME"`
}

func (self *MaxiiotBalances) TableName() string {
	return MaxiiotBalances_TableName
}

//插入数据
func (self *MaxiiotBalances) Create() *MaxiiotBalances {
	_, err := common.Engine.InsertOne(self)
	if err != nil {
		log.Fatalln("Create MaxiiotBalances Err:", err)
		return nil
	}

	return self
}

//创建数据表
func (self *MaxiiotBalances) CreateTable() error {
	err := common.Engine.CreateTables(self)
	if err != nil {
		log.Fatalln("Create MaxiiotBalances Table Error:", err)
	}
	return err
}

// 删除数据
func (u *MaxiiotBalances) Delete() {
	_, err := common.Engine.Delete(u)
	if err != nil {
		log.Fatalln(" Delete MaxiiotBalances Error:", err)
	}
}

//更新MaxiiotBalances
func (self *MaxiiotBalances) Update() {
	common.Engine.Id(self.Id).Update(self)
}

//获取MaxiiotBalances列表
//Find: 获取多条数据
func GetMaxiiotBalancesList(page, perPage int, whereStr, orderStr string) []*MaxiiotBalances {
	var list []*MaxiiotBalances
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
		log.Fatalln("GetMaxiiotBalancesList Error:", err)
		return nil
	}
	return list
}

//MaxiiotBalances总数量
func GetMaxiiotBalancesCount(whereStr string) int64 {
	sql := "id is not null"
	if whereStr != "" {
		sql += fmt.Sprintf(" %s", whereStr)
	}
	total, err := common.Engine.Where(sql).Count(new(MaxiiotBalances))
	if err != nil {
		log.Fatalln("GetMaxiiotBalancesCount Error:", err)
		return 0
	}
	return total
}

//通过ID获取MaxiiotBalances
func GetMaxiiotBalancesById(id int64) *MaxiiotBalances {
	item := new(MaxiiotBalances)
	_, err := common.Engine.Where("id = ?", id).Get(item)
	// has, err := x.Id(id).Get(a)
	if err != nil {
		log.Fatalln("GetMaxiiotBalancesById Error:", err)
		return nil
	}
	return item
}

//通过原生sql获取
func GetMaxiiotBalancesBySql(sql string) []*MaxiiotBalances {
	var list []*MaxiiotBalances
	err := common.Engine.Sql(sql).Find(&list)
	if err != nil {
		log.Fatalln("GetMaxiiotBalancesBySql Error:", err)
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
func CreateMaxiiotBalancesTable(c *gin.Context) {
	err := new(****_model.MaxiiotBalances).CreateTable()
	obj := gin.H{
		"error": err,
	}
	controllers.ResponseJSON(c, http.StatusOK, "sucess", obj)
}
//获取列表
func GetMaxiiotBalancesList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "5"))
	list := ****_model.GetMaxiiotBalancesList(page, perPage, "", "created desc")
	count := ****_model.GetMaxiiotBalancesCount("")
	obj := gin.H{
		"list":       list,
		"list_count": count,
	}
	controllers.ResponseJSON(c, http.StatusOK, "sucess", obj)
}

//获取单条数据
func GetMaxiiotBalancesById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	obj := ****_model.GetMaxiiotBalancesById(id)
	controllers.ResponseJSON(c, http.StatusOK, "sucess", obj)
}

//创建数据
func CreateMaxiiotBalances(c *gin.Context) {
	item := new(****_model.MaxiiotBalances)
	obj := item.Create()
	controllers.ResponseJSON(c, http.StatusOK, "sucess", obj)
}

    "/list": {
      "get": {
        "summary": "获取列表",
        "operationId": "MaxiiotBalancesList",
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
        "operationId": "GetMaxiiotBalancesById",
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
        "operationId": "CreateMaxiiotBalances",
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
