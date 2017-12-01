package tool

import (
	"encoding/base64"
	"log"
)

//加密
func EncodeBase64(s string) string {
	input := []byte(s)
	encodeStr := base64.StdEncoding.EncodeToString(input)
	return encodeStr
}

// 解密
func DecodeBase64(s string) string {
	decodeBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatal("decode base64 error:", err)
		return ""
	}
	return string(decodeBytes)
}
