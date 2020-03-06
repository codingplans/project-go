package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"sync"
)

func main() {
	// gotuni()
	// reflects()
	// appendtest()
	appends()
}
func appends() {
	i := new([]int)
	i = &[]int{1}

	*i = append(*i, 1, 2, 3)

	aa := map[int]int{}
	aa[12] = 21
	aa[122] = 22
	aa[121] = 212
	if ss, ok := aa[22]; !ok {
		println(ok, ss)
	}
	fmt.Println(cap(*i), aa)
}

func appendtest() {
	idd := []int{12, 23, 12, 22, 11, 33}
	ids := append(idd, 12)
	ids = append(idd, 32)
	ids = append(idd, 32)
	ids = append(idd, 12)
	ids = append(idd, 2)
	fmt.Printf("%+v", ids)
	fmt.Printf("%d", len(ids))

	ids = RemoveRepByLoop(ids)
	fmt.Printf("%+v", ids)
	fmt.Printf("%d", len(ids))

}

// 通过两重循环过滤重复元素
func RemoveRepByLoop(slc []int) []int {
	result := []int{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

func gotuni() {
	runtime.GOMAXPROCS(20)
	wg := sync.WaitGroup{}

	i := uint16(1024)
	// j:=1

	// println(9999991 ^ 2)
	println(i >> 2)
	println(i << 2)
	return
	// if int(i+j)>>1
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func(v int) {
	// 		println(v)
	// 		wg.Done()
	// 		println(v, "###")
	//
	// 	}(i)
	// }

	// wg.Wait()
	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		println(i, "***")
		j := i
		go func() {
			println(j, "￥￥￥")
			wg.Done()
		}()
	}

	// println(j, "￥￥￥")

	println(123)
}

func reflects() {

	stu := student{
		name:   "qqq",
		number: 1,
	}

	t := reflect.TypeOf(stu)

	fmt.Printf("%+v\n", t)
	fmt.Printf("%+u", t)
}

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

type student struct {
	name   string
	number int64
}

func Testmap() bool {

	m := make(map[string]*student)

	data := []student{
		{name: "qwe", number: 123},
		{name: "sss", number: 222},
		{name: "ddd", number: 444},
	}

	for k, _ := range data {
		m[data[k].name] = &data[k]
		fmt.Println(data[k].number)
	}

	fmt.Printf("%+v,%+v", data, m["qwe"].number)
	return true
}

func other() {

	Testmap()
	return

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
