package controllers

import (
	"crypto/sha1"
	// "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

const (
	Token = "arvin_wong_token"
)

func WxGet(c *gin.Context) {
	s, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(s))
	signture := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")
	genSign := makeSignature(timestamp, nonce)
	fmt.Println(signture)
	fmt.Println(genSign)
	if genSign == signture {
		fmt.Println("equal")
		c.String(200, echostr)
	} else {
		fmt.Println("not equal")
		c.JSON(200, "not ok")
	}
}

//微信接入时，生成签名
func makeSignature(timestamp, nonce string) string {
	s1 := []string{Token, timestamp, nonce}
	sort.Strings(s1)
	s := sha1.New()
	io.WriteString(s, strings.Join(s1, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func WxPost(c *gin.Context) {
	s, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(s))
	c.JSON(200, "post ok")
}

func Test(c *gin.Context) {
	c.JSON(200, "Hello")
}

func PayNotify(c *gin.Context) {
	s, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(s))
	c.JSON(200, "post ok")
}

func TokenGet(c *gin.Context) {
	s, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(s)
	res, err := http.Post("http://127.0.0.1:8000/billing/invoices/list?page=1", "application/x-www-form-urlencoded", strings.NewReader(""))
	if err != nil {
		fmt.Println("post error", err)
	}
	body, bErr := ioutil.ReadAll(res.Body)
	if bErr != nil {
		fmt.Println("body error", bErr)
	}
	fmt.Println(string(body))
	defer res.Body.Close()
	var result interface{}
	tranErr := json.Unmarshal(body, &result)
	if tranErr != nil {
		fmt.Println("transErr:", tranErr)
	}
	c.JSON(200, result)
}
