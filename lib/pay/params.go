package pay

import ()

const (
	//货币编号
	CURRENCY_CODE_HKD = "344"
	CURRENCY_CODE_USD = "840"
	CURRENCY_CODE_SGD = "702"
	CURRENCY_CODE_CNY = "156" //RMB
	CURRENCY_CODE_JPY = "392"
	CURRENCY_CODE_TWD = "901"
	CURRENCY_CODE_AUD = "AUD"
	CURRENCY_CODE_EUR = "978"
	CURRENCY_CODE_GBP = "826"
	CURRENCY_CODE_CAD = "124"
	CURRENCY_CODE_MOP = "446"
	CURRENCY_CODE_PHP = "608"
	CURRENCY_CODE_THB = "764"
	CURRENCY_CODE_MYR = "458"
	CURRENCY_CODE_IDR = "360"
	CURRENCY_CODE_KRW = "410"
	CURRENCY_CODE_SAR = "682"
	CURRENCY_CODE_NZD = "554"
	CURRENCY_CODE_AED = "784"
	CURRENCY_CODE_BND = "096"
)

var (
	payMethod = []string{"ALL", "CC", "VISA", "Master", "JCB", "AMEX", "Diners", "PPS", "PAYPAL", "CHINAPAY", "ALIPAY", "TENPAY", "99BILL", "MEPS"}
)

func GetPaypalParams() {
	param := map[string]string{
		"orderRef":   "", //订单号
		"mpsMode":    "", //多货币处理服务,NIL没有货币转换
		"currCode":   "", //货币编号
		"amount":     "",
		"lang":       "", //"C":繁体中文；"X":简体中文;"E":英语
		"cancelUrl":  "",
		"failUrl":    "",
		"successUrl": "",
		"merchantId": "", //商户代码
		"payType":    "", //"N":消费交易;"H":预授权交易
		"payMethod":  "",
		//--------可选-----------------
		"remark":      "",
		"redirect":    "",
		"oriCountry":  "",   //源国家编码,344:"HK"
		"destCountry": "",   //目的国家编码
		"print":       "no", //支付结果页面，关闭打印功能
		"failRetry":   "no", //交易被拒绝，关闭重试功能
		"Ref":         "",   //订单参考号,取消页面跳转时作为Query参数返回
		"":            "",
		"":            "",
	}
}
