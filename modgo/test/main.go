package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
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
