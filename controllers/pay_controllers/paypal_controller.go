package pay_controllers

import(
	paypal "github.com/logpacker/PayPal-Go-SDK"
	"github.com/gin-gonic/gin"
	"os"
	"fmt"
)

const(
	live_account = "hfyarvin@gmail.com"
	access_token = "access_token$production$rvfnyxz95j4kmsrh$1c086825b9952b82f797ea3c6128b856"
	expiry_date = "22 Nov 2027"
	SANDBOX_ACCOUNT = "hfyarvin-facilitator@gmail.com"
	SANDBOX_CLIENTID = "AW_AhlQXzMYg4Qtb3kkXnS69k5ZL1MYJBq0Prkv79f7FoIytq_oXYxoGhLwjOuvTuRIMM-UmTMdFpFD8"
	SANDBOX_SECRET = "EAZqkGEHyp16kgaFU6UouZC8KFJesS0pyYtcyEU-M6F-1nnXz1Q-q648tgvWyR8C99GU5KU3XpF0Qyhz"
)

var (
	MY_ACCESS_TOKEN = map[string]interface{}{
		"refresh_token": "",
		"access_token": "A21AAGnMjmDtDHOqAz1riJ4bdXu9uQmYnw1bF2iSpgcQmpFOkJ2_UnF24Dg0jgwYiC60yZEABPAMobhhkYYeFnSJj8-jiQD3Q",
		"token_type": "Bearer",
		"expires_in": 32382,
	}
)

func GetNewClient() *paypal.Client{
	fmt.Println(MY_ACCESS_TOKEN["expires_in"])
	fmt.Println("==================Func Get new Client============================================")
	clientId := SANDBOX_CLIENTID
	secret := SANDBOX_SECRET
	fmt.Println("==================Func create new Client============================================")
	client, err := paypal.NewClient(clientId,secret,paypal.APIBaseSandBox)
	if err != nil {
		fmt.Println("======create client err:=====", err)
	}
	fmt.Println(client.Token)
	client.SetLog(os.Stdout)
	fmt.Println("==================Func Get Access Token============================================")
	accessToken,err := client.GetAccessToken()
	if err!=nil {
		fmt.Println("=====access token err:=====", err)
	}
	fmt.Println("==================Print Access Token============================================")
	fmt.Println(accessToken.Token)
	fmt.Println(MY_ACCESS_TOKEN["access_token"])
	// c.JSON(200, accessToken)
	return client
}

func DirectPaypalPaymentTest(c *gin.Context) {
	client := GetNewClient()
	if client == nil {
		c.JSON(200,"not ok")
	} else {
		amount := paypal.Amount{
			Total: "1.00",
			Currency: "USD",
		}
		redirectURL := "http://arvin-wong.natapp1.cc/default/post/info"
		cancelURL := "http://arvin-wong.natapp1.cc/default/post/info"
		desc := "Description for this direct paypal payment"
		paymentRes, err := client.CreateDirectPaypalPayment(amount, redirectURL, cancelURL, desc)
		if err !=nil {
			fmt.Println("payments Result Error:", err)
		}
		c.JSON(200,paymentRes)
	}
}

func ShowClientInfo(c *gin.Context) {
	client := GetNewClient()
	if client != nil {
		obj := gin.H{
			// ClientID
			// Secret         string
			// APIBase        string
			// Token          *TokenResponse
			// tokenExpiresAt time.Time
			"client_id": client.ClientID,
			"secret": client.Secret,
			"api_base": client.APIBase,
			"token": client.Token,
		}
		c.JSON(200,obj)
	} else {
		c.JSON(403,"not ok")
	}
}

func PaypalGet(c *gin.Context) {
	c.JSON(200, "get ok")
}

func PaypalPost(c *gin.Context) {
	c.JSON(200, "post ok")
}