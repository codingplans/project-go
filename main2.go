package main

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"runtime"
	"sync"
	"time"
)

type student struct {
	name   string
	number int64
}

func main() {
	// array()
	// syncpool()
	// stringtojson()
	// wgs()
	// println(reverse(508200))
	// chanss()
	// synconce()
	println(333)
	time.Sleep(2 * time.Second)
}

func synconce() {
	var con sync.Once
	var temp int
	println(temp)
	go con.Do(func() {
		println(222)
		temp = 1
		time.Sleep(3 * time.Second)
		println(25555)

	})
	time.Sleep(1 * time.Second)
	println(temp)
	time.Sleep(3 * time.Second)
	println(temp)
}

func chanss() {
	c := make(chan int)
	defer close(c)
	go func() {
		// time.Sleep(time.Second * 3)

		c <- 4
		c <- 3
	}()

	println(123)

	i := <-c

	println(cap(c), i)
	println(cap(c), <-c)
	println(len(c))

}

func reverse(x int32) int32 {
	var num int64
	x64 := int64(x)
	println(x64)
	for x64 != 0 {
		num = num*10 + x64%10
		x64 = x64 / 10
		println(num*10, "www", x64%10)
		println(x64)
		println(num)

	}
	// 使用 math 包中定义好的最大最小值
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}
	return int32(num)
}

func wgs() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}

	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func locks() {
	var L *sync.RWMutex
	L = new(sync.RWMutex)
	L.Lock()

}
func stringtojson() {
	jstring := `{"id":1,"name":"zhang"}`
	jMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(jstring), &jMap)
	fmt.Printf("%+v", jMap)
}

func syncpool() {

	chi := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	cc := &sync.Pool{
		New: func() interface{} {
			return student{
				name:   "",
				number: 0,
			}
		},
	}

	cc.Put(student{
		name:   "12",
		number: 333,
	})
	chi.Put(student{
		name:   "12",
		number: 333,
	})
	chi.Put(100)
	chi.Put(200)
	chi.Put(600)
	chi.Put(400)
	chi.Put(500)
	chi.Put(50)
	chi.Put(50)
	chi.Put(50)
	chi.Put(50)
	chi.Put(50)
	defer defertime(3)
	defer defertime(4)
	fmt.Printf("%+v \n", chi.Get())
	fmt.Printf("%+v \n", chi.Get())
	fmt.Printf("%+v \n", chi.Get())
	fmt.Printf("%+v \n", chi.Get())
	fmt.Printf("%+v \n", chi.Get())
	fmt.Printf("%+v \n", chi.Get())
	fmt.Printf("%+v \n", chi.Get())
	fmt.Printf("%+v \n", chi.Get())
	fmt.Printf("%+v \n", chi.Get())
	fmt.Printf("%+v \n", chi.Get())
	fmt.Printf("%+v 1212\n", cc.Get())
	defer defertime(1)
	defer defertime(2)
	return
}

func defertime(i int) {
	println(time.Now().Unix(), i)
}

func array() {
	var arrs []int
	arrs = append(arrs, 1)
	arrs = append(arrs, 1)
	arrs = append(arrs, 1)
	fmt.Printf("%+v \n", cap(arrs))

	arr := [10]int{1, 12, 22}
	arr[5] = 1
	arr[9] = arr[0]
	fmt.Printf("%d", cap(arr))

}
func tickers() {
	// 打点器和定时器的机制有点相似：一个通道用来发送数据。
	// 这里我们在这个通道上使用内置的 `range` 来迭代值每隔
	// 500ms 发送一次的值。
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// 打点器可以和定时器一样被停止。一旦一个打点停止了，
	// 将不能再从它的通道中接收到值。我们将在运行后 1600ms
	// 停止这个打点器。
	time.Sleep(time.Millisecond * 21600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
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
