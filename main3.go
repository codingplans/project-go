package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type doc2 struct {
	Taa int64
	aaa int64
}
type doc struct {
	Taa int64
	aaa int64
	dw  doc2
}

func main() {
	//nullpoint()

	var ppFree = sync.Pool{
		New: func() interface{} { return new(doc) },
	}

	//d := getD(ppFree)
	//d2 := getD(ppFree)
	ppFree.Put(&doc{aaa: 22, Taa: 22})
	ppFree.Put(&doc{aaa: 11, Taa: 11})
	time.Sleep(time.Second * 5)
	d := ppFree.Get().(*doc)
	d1 := ppFree.Get().(*doc)
	d2 := ppFree.Get().(*doc)
	fmt.Println(d, d1, d2)
}

func (d *doc) putD() {
	d.aaa = 11
	d.Taa = 11

}

func getD(ppFree sync.Pool) *doc {
	d := ppFree.Get().(*doc)
	d.aaa = 123
	d.Taa = 222
	return d
}

func nullpoint() {
	// NumCPU 返回当前进程可以用到的逻辑核心数
	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	//var data *doc
	//var data doc
	var data = new(doc)

	//data := &doc{}

	data.aaa = 123
	data.Taa = 333

	fmt.Println(data.aaa)
	fmt.Println(&data.aaa)
	sma(data)
	fmt.Println(data.aaa)
}

func sma(data *doc) (re doc, ree error) {
	//func sma(data *doc) {
	data.aaa = 1
	fmt.Println(data.aaa)
	re = *data
	ree = nil
	return
}
