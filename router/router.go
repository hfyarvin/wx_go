package router

import (
	"../controllers"
	"../controllers/db_controller"
	"../controllers/file_controllers"
	"../controllers/mail_controllers"
	"../controllers/pay_controllers"
	"../controllers/tool_controllers"
	"../controllers/wx_controllers"
	"../middler"
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
		paypalGroup.GET("/direct/payment", pay_controllers.DirectPaypalPaymentTest)
		paypalGroup.GET("/custom/payment", pay_controllers.CreateCustomPayment)
		paypalGroup.GET("/index", pay_controllers.GetPaypalIndexPage)
		buttonGroup := paypalGroup.Group("/button")
		{
			buttonGroup.GET("/create", pay_controllers.PaypalButtonIndex)
		}
	}

	r.GET("/index", controllers.GetIndexPage)
	r.GET("/upload/page", controllers.GetUploadPage)
	r.Any("/upload/test", file_controllers.UploadTest)

	// r.MaxMultipartMemory = 8 << 20
	// r.Static("/", "../controllers/file_controllers")

	//上传附件
	fileGroup := r.Group("/file")
	{
		fileGroup.GET("/upload", controllers.GetUploadPage)
		fileGroup.POST("/single/upload", file_controllers.UploadSingleFile)
		fileGroup.POST("/upload/one", file_controllers.UploadOneFile)
		fileGroup.POST("/upload/multiple", file_controllers.UploadMultipleFiles)
		fileGroup.GET("/upload/multiple/page", file_controllers.UploadMultipleFilesPage)
	}

	//邮件
	mailGroup := r.Group("/mail")
	{
		mailGroup.GET("/", mail_controllers.SentEmail)
	}

	//paydollar
	paydollarGroup := r.Group("paydollar")
	{
		paydollarGroup.GET("/index", pay_controllers.GetPaydollarIndexPage)
		paydollarGroup.GET("/cancel", pay_controllers.CancelPaydollar)
		paydollarGroup.GET("/success", pay_controllers.SuccessPaydollar)
		paydollarGroup.GET("/fail", pay_controllers.FailPaydollar)
	}
	//重定向
	toolGroup := r.Group("/tool")
	toolGroup.Use(middler.LogTime())
	{
		toolGroup.GET("/redirect_url", tool_controllers.GetRedirectUrl)
		toolGroup.GET("/base64/test", tool_controllers.Base64Test)
	}

	//table
	r.GET("tables/cols", db_controller.AllTabelColumns)
	r.GET("tables", db_controller.AllTabels)
	r.GET("gen", db_controller.Gen)
}

// api_base: https://api.sandbox.paypal.com
