package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func Testmaps() {

	res := recursive(1)
	print(res)

}

func recursive(left int) int {

	print(left, "\n")

	if left == 10 {
		return (left + 1) * 2
	}

	return recursive((left + 1)) * 2
}

// func Test19(t *testing.T) {
// 	s := make([]map[string]int, 10)
// 	delete(s[1], "h")
// 	fmt.Println(s)
// }

func main() {

	var uniprice int64
	sadas := 0.8
	uniprice = 1900
	uniprice = uniprice * int64(sadas*100) / 100

	println(uniprice)
	return

	type User struct {
		UserId   int    `json:"user_id" bson:"123123"`
		UserName string `json:"user_name" bson:"3333"`
	}
	// 输出json格式
	u := &User{UserId: 1, UserName: "tony"}
	arr := make([]*User, 0)
	js, _ := json.Marshal(u)
	arr = append(arr, u)
	arr = append(arr, u)
	aa, _ := json.Marshal(arr)
	fmt.Println(string(js), aa)
	// 输出内容：{"user_id":1,"user_name":"tony"}

	var re User
	if err := json.Unmarshal(js, re); err != nil {
		fmt.Println(re.UserId)
	}

	// 获取tag中的内容
	t := reflect.TypeOf(u)
	fields := t.Elem()
	field := t.Elem().Field(1)
	fmt.Println(field.Tag.Get("json"), fields.Name())
	fmt.Println(field.Tag.Get("bson"))

	// Test19(nil)
	// Testmaps()
}
