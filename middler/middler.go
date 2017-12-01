package middler

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

//定义一个中间件
func LogTime() gin.HandlerFunc {
	return func(c *gin.Context) {

		now := time.Now()
		//在gin上下文定义变量
		c.Set("example", "12345")

		//请求前
		c.Next()

		//请求后
		latecy := time.Since(now)
		log.Print(latecy)

		//access the status we are sending
		// status := c.Writer.Status()
		// log.Fatalln(status)
	}
}
