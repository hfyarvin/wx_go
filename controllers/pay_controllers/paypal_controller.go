package pay_controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	paypal "github.com/logpacker/PayPal-Go-SDK"
	"net/http"
	"os"
)

const (
	live_account     = "hfyarvin@gmail.com"
	access_token     = "access_token$production$rvfnyxz95j4kmsrh$1c086825b9952b82f797ea3c6128b856"
	expiry_date      = "22 Nov 2027"
	SANDBOX_ACCOUNT  = "982073560.wong@qq.com"
	SANDBOX_CLIENTID = "ATtwdeAOKDgyonrn6iQ2l0B_v8fRvYn1qDoJldPmGjlaGPNnLSCaSR2xlUtDEAYreltdtUR7boglX0NR"
	SANDBOX_SECRET   = "EHEtF6rBt_CO9X-45uLnTHyANLVYKBhnbcPaTYZ9FiJrX8njdDerjC8Z4Z8ADHckQlnNKf2yXydr4KBV"
)

var (
	MY_ACCESS_TOKEN = map[string]interface{}{
		"refresh_token": "",
		"access_token":  "A21AAGnMjmDtDHOqAz1riJ4bdXu9uQmYnw1bF2iSpgcQmpFOkJ2_UnF24Dg0jgwYiC60yZEABPAMobhhkYYeFnSJj8-jiQD3Q",
		"token_type":    "Bearer",
		"expires_in":    32382, //过期时间
	}
)

func GetNewClient() *paypal.Client {
	fmt.Println("==============Get New Client==================")
	clientId := SANDBOX_CLIENTID
	secret := SANDBOX_SECRET
	client, err := paypal.NewClient(clientId, secret, paypal.APIBaseSandBox)
	if err != nil {
		return nil
		fmt.Println("======create client err:=====", err)
	}
	client.SetLog(os.Stdout)
	//再次获取token
	_, err = client.GetAccessToken()
	if err != nil {
		return nil
		fmt.Println("=====access token err:=====", err)
	}
	return client
}

//获取刚创建的客户端的信息
func ShowClientInfo(c *gin.Context) {
	client := GetNewClient()
	if client != nil {
		obj := gin.H{
			"api_base":  client.APIBase,
			"client_id": client.ClientID,
			"secret":    client.Secret,
			"token":     client.Token,
		}
		c.JSON(200, obj)
	} else {
		c.JSON(403, "not ok")
	}
}
func DirectPaypalPaymentTest(c *gin.Context) {
	client := GetNewClient()
	if client == nil {
		c.JSON(200, "client is nil")
	} else {
		amount := paypal.Amount{
			Total:    "1.00",
			Currency: "USD",
		}
		redirectURL := "http://arvin-wong.natapp1.cc/paydollar/success"
		cancelURL := "http://arvin-wong.natapp1.cc/paydollar/cancel"
		desc := "Description for this direct paypal payment"
		paymentRes, err := client.CreateDirectPaypalPayment(amount, redirectURL, cancelURL, desc)
		if err != nil {
			fmt.Println("payments Result Error:", err)
		}
		c.JSON(200, paymentRes)
	}
}

func PaypalGet(c *gin.Context) {
	c.JSON(200, "get ok")
}

func PaypalPost(c *gin.Context) {
	c.JSON(200, "post ok")
}

func GetPaypalIndexPage(c *gin.Context) {
	obj := gin.H{
		"title": "Main",
	}
	c.HTML(http.StatusOK, "paypal.tmpl", obj)
}

func CreateCustomPayment(c *gin.Context) {
	client := GetNewClient()
	p := paypal.Payment{
		Intent: "sale",
		Payer: &paypal.Payer{
			PaymentMethod: "credit_card",
			FundingInstruments: []paypal.FundingInstrument{paypal.FundingInstrument{
				//信用卡信息
				CreditCard: &paypal.CreditCard{
					Number:      "4111111111111111",
					Type:        "visa",
					ExpireMonth: "11",
					ExpireYear:  "2020",
					CVV2:        "777",
					FirstName:   "John",
					LastName:    "Doe",
				},
			}},
		},
		Transactions: []paypal.Transaction{paypal.Transaction{
			Amount: &paypal.Amount{
				Currency: "USD",
				Total:    "7.00",
			},
			Description: "My Payment",
		}},
		RedirectURLs: &paypal.RedirectURLs{
			ReturnURL: "http://arvin-wong.natapp1.cc/paydollar/success",
			CancelURL: "http://arvin-wong.natapp1.cc/paydollar/cancel",
		},
	}
	paymentResponse, err := client.CreatePayment(p)
	if err != nil {
		c.JSON(200, "the pay is failed")
	} else {
		c.JSON(200, paymentResponse.Payment)
	}
}
