package mail_controllers

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
	"log"
)

func SentEmail(c *gin.Context) {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "arvin.wong@maxiiot.com", "arvin.wong@maxiiot.com")
	m.SetHeader("To", m.FormatAddress("982073560@qq.com", "arvin"))
	m.SetHeader("Subject", "测试")
	m.SetBody("text/html", "Hello <a href = \"http://blog.csdn.net/liang19890820\">一去丶二三里</a>")
	// m.Attach("附件")
	d := gomail.NewPlainDialer("mail.maxiiot.com", 465, "arvin.wong@maxiiot.com", "arvin171017")
	if err := d.DialAndSend(m); err != nil {
		log.Fatalln("==================send failed===========", err)
		c.JSON(200, "failed")
	} else {
		c.JSON(200, "successful")
	}
}
