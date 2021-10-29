package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	println(123)
	go func() {
		for {
			fmt.Println(Add("https://github.com/EDDYCJY"))
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
