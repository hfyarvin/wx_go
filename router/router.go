package router

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.GET("/wx", controllers.WxGet)
	r.POST("/wx", controllers.WxPost)
	wxGroup := r.Group("/wx")
	{
		//token
		wxGroup.GET("/token/get", controllers.TokenGet)
		//wxpay
		wxGroup.POST("/pay/notify", controllers.PayNotify)
	}
	r.GET("/test", controllers.Test)
}
