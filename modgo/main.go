package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"hash"
)

func main() {

	// s := "darren" + "123:129"

	fmt.Println(authorize("123", "123", "post", "123", "123", 100))

}

// 生成鉴权签名
func authorize(secret, method, bucket string, expire int64) string {
	var (
		content   string
		mac       hash.Hash
		signature string
	)
	content = fmt.Sprintf("%s\n%s\n%d\n", method, bucket, expire)
	mac = hmac.New(sha1.New, []byte(secret))
	mac.Write([]byte(content))
	signature = base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return signature
}
