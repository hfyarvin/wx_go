package pay_controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func GetPaydollarIndexPage(c *gin.Context) {
	obj := gin.H{
		"title": "Main",
	}
	c.HTML(http.StatusOK, "paydollarPay.tmpl", obj)
}

func CancelPaydollar(c *gin.Context) {
	fmt.Println("------the request method-----------:", c.Request.Method)
	req := c.Request.Body
	body, _ := ioutil.ReadAll(req)
	var (
		res interface{}
	)
	err := json.Unmarshal(body, &res)
	if err != nil {

		c.JSON(200, res)
	} else {
		c.JSON(403, "fail")
	}
}

func SuccessPaydollar(c *gin.Context) {
	req := c.Request.Body
	body, _ := ioutil.ReadAll(req)
	var (
		res interface{}
	)
	fmt.Println("------the request method-----------:", c.Request.Method)
	err := json.Unmarshal(body, &res)
	if err != nil {
		c.JSON(200, res)
	} else {
		c.JSON(403, "fail")
	}
}

func FailPaydollar(c *gin.Context) {
	fmt.Println("------the request method-----------:", c.Request.Method)
	req := c.Request.Body
	body, _ := ioutil.ReadAll(req)
	var (
		res interface{}
	)
	err := json.Unmarshal(body, &res)
	if err != nil {
		c.JSON(200, res)
	} else {
		c.JSON(403, "fail")
	}
}
