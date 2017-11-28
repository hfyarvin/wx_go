package router

import (
	"../controllers"
	"../controllers/file_controllers"
	"../controllers/mail_controllers"
	"../controllers/pay_controllers"
	"../controllers/wx_controllers"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.GET("/wx", controllers.WxGet)
	r.POST("/wx", controllers.WxPost)
	wxGroup := r.Group("/wx")
	{
		//获取access_token
		wxGroup.GET("/access_token", wx_controllers.GetAccess)
		//token
		wxGroup.GET("/token/get", controllers.TokenGet)
		//wxpay
		wxGroup.POST("/pay/notify", controllers.PayNotify)
	}
	r.GET("/test", controllers.Test)
	r.POST("/default/post/info", controllers.GetPostInfo)
	paypalGroup := r.Group("/paypal")
	{
		paypalGroup.GET("/", pay_controllers.PaypalGet)
		paypalGroup.POST("/", pay_controllers.PaypalPost)
		paypalGroup.GET("/new/client", pay_controllers.ShowClientInfo) //had changed
		paypalGroup.POST("/direct/payment", pay_controllers.DirectPaypalPaymentTest)
	}

	r.GET("/index", controllers.GetIndexPage)
	r.GET("/upload/page", controllers.GetUploadPage)
	r.Any("/upload/test", file_controllers.UploadTest)

	// r.MaxMultipartMemory = 8 << 20
	// r.Static("/", "../controllers/file_controllers")

	//上传附件
	fileGroup := r.Group("/file")
	{
		fileGroup.POST("/single/upload", file_controllers.UploadSingleFile)
	}

	//邮件
	mailGroup := r.Group("/mail")
	{
		mailGroup.GET("/", mail_controllers.SentEmail)
	}
}

// api_base: https://api.sandbox.paypal.com
