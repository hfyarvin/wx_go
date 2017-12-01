package tool_controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetRedirectUrl(c *gin.Context) {
	fmt.Println("this is a redirect===========================================")
	go func() {
		cCp1 := c.Copy()
		time.Sleep(5 * time.Second)
		fmt.Println("Done! in path ", cCp1.Request.URL.Path)
		fmt.Println(c.MustGet("example").(string))
	}()
	// c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	time.Sleep(5 * time.Second)
	c.JSON(http.StatusOK, "ok")
}
