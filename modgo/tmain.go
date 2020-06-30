package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"strconv"
	"time"
)

type PayWay struct {
	//    支付id
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// 支付名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

// 币种
type Fiats struct {
	//    名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//    币id
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

// 合并接口 数据包
type SyncResp struct {
	//    交易方式
	PayWay []*PayWay `protobuf:"bytes,1,rep,name=pay_way,json=payWay,proto3" json:"pay_way,omitempty"`
	//    当前交易所流通货币
	Symbol []string `protobuf:"bytes,2,rep,name=symbol,proto3" json:"symbol,omitempty"`
	// 法币币种
	Fiats []*Fiats `protobuf:"bytes,3,rep,name=fiats,proto3" json:"fiats,omitempty"`
}

type RRR struct {
}

type RespData struct {
	Code    int64  `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func main() {

	aa, _ := strconv.ParseFloat("152413.234452", 64)
	log.Infof("%+v ", aa, aa)

	// resp, err := http.Get("http://127.0.0.1:8000/exchange/usercenter/sync")
	// // resp, err := http.Get("http://192.168.0.125:9526/rpc/address/99990")
	// if err != nil {
	// 	return
	// }
	// defer resp.Body.Close()
	//
	// body, err := ioutil.ReadAll(resp.Body)
	//
	// response := &resaa{}
	//
	// json.Unmarshal(body, response)

	// if err := decoder.Decode(&response); err != nil && err != io.EOF {
	// 	return
	// }

	// log.Infof("%+v %v", resp, string(body))
	// log.Infof("%+v ********** %+v", resp.StatusCode, response.Data.Fiats[1])

	// switchs()

	// floattostr()

	// floattest()

	// maptest()

	// gogo()
}

func switchs() {
	aa := 11
	switch aa {
	case 22:
		println(aa, 2)
	case 3:
		println(aa, 3)
	case 1:
		println(aa, 11)
		aa++
	case 2:
		println(aa, 2)
	default:
		println(aa, 111)
		return

	}
	println(aa)
}

func floattostr() {
	var ff float64

	ff = 3198000.2200020

	aa := strconv.FormatFloat(ff, 'f', -1, 64)
	log.Info(aa)
}

func gogo() {
	var ss string
	// if ss, ok = qqs(); ok {
	// 	println(123)
	// }

	var tt time.Time
	tt = time.Now()
	ch := make(chan int, 10)
	exits := make(chan int, 2)

	go qqs2(ch)
	go qqs(ch, exits)
	// go qqs3(ch)
	fmt.Printf("%+v  %s  %d", ss, tt, tt.Unix())

	for i := 0; i < 5; i++ {
		ch <- i
	}

	time.Sleep(2 * time.Second)

	close(exits)
	if dd, ok := <-exits; ok {
		println(123, dd)
	} else {
		println(333)

	}

	time.Sleep(20 * time.Second)
}

func qqs(c, exits chan int) (string, bool) {
	select {
	case aa, ok := <-c:
		if ok {
			println(aa, 111)
		}
	case <-exits:
		println(1111, 111)

	}
	println("over")

	return "12", true
}

func qqs2(c chan int) (string, bool) {
	for {
		select {
		case aa, ok := <-c:
			if ok {
				println(aa, 222)
			}
		}
	}
}

func qqs3(c chan int) (string, bool) {
	for {
		select {
		case aa, ok := <-c:
			if ok {
				println(aa, 333)
			}
		}
	}
}

func maptest() {
	mm := make(map[string]int64, 0)
	mm["aaaaaaaaaaaaaaaaaaa"] = 123
	mm["zzzzzzz"] = 123
	mm["cccccccccccccccccccccc"] = 123
	mm["bb"] = 123
	mm["kkkkkkkkkkkkkkkk"] = 123
	mm["rr"] = 123
	mm["zzzz"] = 123
	for k, v := range mm {
		println(k, v)
	}

	fmt.Printf("%+v", mm)
}

type ttt struct {
	Amount  string
	Timeout string
}

type ttt2 struct {
	Amount  string
	Timeout string
}

func floattest() {
	s := "123.129asdas"

	// var fff float64
	//
	// fff = 55.43453
	//
	// println(strconv.FormatFloat(fff, 'g', -1, 64))

	aa, err := strconv.ParseFloat(s, 8)
	if err == nil {
		fmt.Printf("%T,%v", aa, aa)
	} else {
		fmt.Printf("%T,%v ****", aa, aa)

		println(err.Error(), aa)

	}

	var ae uint8
	ae = 'a'

	aas := float64(123123.22)
	ww := strconv.FormatFloat(aas, 'g', -1, 64)

	ww = fmt.Sprintf("%d", 100) + "12"
	println(ww, ae)
	// fmt.Println(authorize("123", "123", "post", 123))
}
