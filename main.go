package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func getSig(appId string, ts int64, secretKey string) string {
	// 拼接 basestring
	baseString := fmt.Sprintf("%s%d", appId, ts)

	// 计算 MD5
	md5Bytes := md5.Sum([]byte(baseString))
	md5String := hex.EncodeToString(md5Bytes[:])

	// 计算 HmacSHA1
	key := []byte(secretKey)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(md5String))
	hmacBytes := h.Sum(nil)

	// base64 编码
	return base64.StdEncoding.EncodeToString(hmacBytes)
}

func main() {
	// 测试
	appId := "595f23df"
	ts := int64(1512041814) // 当前时间戳
	secretKey := "d9f4aa7ea6d94faca62cd88a28fd5234"

	sig := getSig(appId, ts, secretKey)
	fmt.Println(sig) // 输出计算得到的签名
}
