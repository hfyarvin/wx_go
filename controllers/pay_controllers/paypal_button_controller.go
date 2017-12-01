package pay_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PaypalButtonIndex(c *gin.Context) {
	obj := gin.H{
		"title": "Main",
	}
	c.HTML(http.StatusOK, "paypal_button.tmpl", obj)
}
