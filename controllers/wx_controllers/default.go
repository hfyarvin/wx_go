package wx_controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type AccessTokenInfo struct{
	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
	Errcode int64 `json:"errcode"`
	Errmsg string `json:"errmsg"`
}

func GetAccessToken( ) *AccessTokenInfo{
	appId := "wx5f06a7e4356f0b89"
	secret := "4bc31d0c85aa880b2d33124bf2a35fd1"
	url	:= fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appId, secret)
	res, err := http.Get(url)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("==============Read All err===================")
		fmt.Println(err)
		return nil
	}
	access := new(AccessTokenInfo)
	err = json.Unmarshal(body,&access)
	if err != nil {
		fmt.Println("==============Unmarshal err===================")
		fmt.Println(err)
		return nil
	}
	return access
}

func GetAccess(c *gin.Context) {
	access := GetAccessToken()
	fmt.Println("...token...",access.AccessToken)
	c.JSON(200,access)
}