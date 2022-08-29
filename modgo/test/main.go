package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
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
			q++
		}

	}
}
