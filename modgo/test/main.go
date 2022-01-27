package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"
)

func mai2n() {
	go func() {
		for {
			Add("hs")
			// fmt.Println()
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			fmt.Printf("%p %d,%d \n", datas, len(datas), cap(datas))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}

var datas []string

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}

func main() {

	ticker := time.NewTicker(time.Second)
	q := int64(18004800001)
	for {
		select {
		case <-ticker.C:
			a := q + rand.Int63n(10000)
			a -= 1000000000 * rand.Int63n(8)
			println(a)
			go postAcck(a)
			q++

		}

	}
}

func postAcck(m int64) {

	url := "http://bj2.linghu888.vip/index/save_customer"
	method := "POST"

	b0 := `name=%E6%E6%9D%9D%8E%E5%87%A4&phone=`
	b := `&type=2&channel_id=2&city=%E5%8C%97%E4%BA%AC+%E5%8C%97%E4%BA%AC%E5%B8%82+%E4%B8%9C%E5%9F%8E%E5%8C%BA%BA%BA%BA&address=%E4
%B8%9C%E5%9F%8E%E5%A4%A7%E6%A5%BC+201+%E5%8F%B7%B7%B7%B7&num=2`
	str := fmt.Sprintf("%s%d%s", string(b0), m, string(b))
	// println(str)
	payload := strings.NewReader(str)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Mac212intosh; Intel Mac OS X 10_15_7) AppleWebKi21212t/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Origin", "http://bj2.linghu888.vip")
	req.Header.Add("Referer", "http://bj2.linghu888.vip/?channel=1&type=3")
	req.Header.Add("Accept-Language", "zh-TW,zh;q=0.9,en-US;q=0.8,en;q=0.7,zh-CN;q=0.6")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
