package controllers

import (
	"crypto/sha1"
	// "encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
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
	type Ret struct {
		XMLName      xml.Name `xml:"xml"`
		MsgType      string   `xml:"MsgType"`
		Content      string   `xml:"Content"`
		ToUserName   string   `xml:"ToUserName"`
		CreateTime   int64    `xml:"CreateTime"`
		FromUserName string   `xml:"FromUserName"`
		PicUrl       string   `xml:"PicUrl"`
		MediaId      string   `xml:"MediaId"`
		MsgId        string   `xml:"MsgId"`
	}
	result := new(Ret)
	err := xml.Unmarshal(s, &result)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("result")
	fmt.Println(result)
	// err := json.Unmarshal(s,&result)
	// if err != nil {
	// 	fmt.Println("err:",err)
	// }
	// fmt.Println("result")
	// fmt.Println(result)
	r := new(Ret)
	r.MsgType = result.MsgType
	r.Content = "Hello"
	r.CreateTime = time.Now().Unix()
	r.FromUserName = result.ToUserName
	r.ToUserName = result.FromUserName
	r.PicUrl = "https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=4147765522,1884713051&fm=27&gp=0.jpg"
	r.MediaId = result.MediaId
	r.MsgId = result.MsgId
	obj := r
	c.XML(200, obj)
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
	res, err := http.Post("http://127.0.0.1:8000/billing/invoices/list?page=2", "application/x-www-form-urlencoded", strings.NewReader(""))
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

func GetWxToekn(c *gin.Context) {
	// appId := "wx0ed3b325349ac4df"
	// secret := ""
	// url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",appId,secret)
	// res, err := http.Get("http://127.0.0.1:8000/billing/invoices/list?page=2")
	// if err != nil {
	// 	fmt.Println("post error", err)
	// }
	// body, bErr := ioutil.ReadAll(res.Body)
	// if bErr != nil {
	// 	fmt.Println("body error", bErr)
	// }
	// fmt.Println(string(body))
	// defer res.Body.Close()
	// var result interface{}
	// tranErr := json.Unmarshal(body, &result)
	// if tranErr != nil {
	// 	fmt.Println("transErr:", tranErr)
	// }
	// c.JSON(200, result)
}

func GetPostInfo(c *gin.Context) {
	s, _ := ioutil.ReadAll(c.Request.Body)
	var result interface{}
	tranErr := json.Unmarshal(s, &result)
	if tranErr != nil {
		fmt.Println(tranErr)
	}
	fmt.Println("==================Print Post Info============================================")
	fmt.Println(result)
	c.JSON(200, result)
}

func GetIndexPage(c *gin.Context) {
	obj := gin.H{
		"title": "Main",
	}
	c.HTML(http.StatusOK, "index.tmpl", obj)
}

func GetUploadPage(c *gin.Context) {
	obj := gin.H{
		"title": "Main",
	}
	c.HTML(http.StatusOK, "upload.tmpl", obj)
}
