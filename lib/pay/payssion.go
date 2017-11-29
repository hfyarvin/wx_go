package pay

import (
	"fmt"
)

const (
	ResultCode_Success     = 200
	ResultCode_ParamsError = 400
	ResultCode_MerchantId  = 401
	ResultCode_
	ResultCode_ApiSignError         = 402
	ResultCode_AppNameError         = 403
	ResultCode_PayWayError          = 405
	ResultCode_CurrencyError        = 406
	ResultCode_AmountError          = 407
	ResultCode_LanguageError        = 408
	ResultCode_UrlError             = 409
	ResultCode_SecretKeyError       = 411
	ResultCode_TeansactionIdError   = 412
	ResultCode_OrderRepetitionError = 413 //订单重复
	ResultCode_CountryError         = 414
	ResultCode_PayTypeError         = 415
)

func GetPayssionParams() {
	params := map[string]string{
		"api_key":     "",
		"pm_id":       "",
		"amount":      "",
		"currency":    "",
		"description": "",
		"order_id":    "",
		"api_sig":     "",
		"notify_url":  "",
		"return_url":  "",
		"success_url": "",
		"language":    "",
	}
	fmt.Println(params)
}

type PayssionReponse struct {
	ResultCode    int64       `json:"result_code"`
	Todo          string      `json:"todo"`
	ReturnUrl     string      `json:"return_url"`
	DeviceSupport string      `json:"device_support"`
	Transaction   interface{} `json:"transaction"`
	Bankaccount   interface{} `json:"bankaccount"`
}
