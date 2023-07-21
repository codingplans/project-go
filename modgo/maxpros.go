package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("runtime.GOMAXPROCS:", runtime.GOMAXPROCS(2))
	go func() {
		for {
			fmt.Println("NumGoroutine", runtime.NumGoroutine())
			time.Sleep(time.Second * 2)
		}
	}()

	for i := 0; i < 1000; i++ {
		go timeSleep10(i)
	}
	time.Sleep(time.Second * 41)
}

func timeSleep10(i int) {
	fmt.Println(i)
	arr := []int{}
	for i := 0; i < 100; i++ {
		arr = append(arr, i)
		time.Sleep(time.Millisecond * time.Duration(i*10))
	}
	fmt.Println("ok", i)
}
