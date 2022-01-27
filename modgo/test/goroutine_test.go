package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

//  多消费多生产 场景 下通过 wg 控制生产者关闭 和消费
func TestGoroutine(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// runtime.GOMAXPROCS(1)
	buffer := make(chan int, 100)
	for i := 0; i < 40; i++ {
		go Producer(buffer, i)
		wg.Add(1)
	}
	for i := 0; i < 40; i++ {
		go Consumer(buffer, i)
		wgCon.Add(1)
	}

	time.Sleep(3 * time.Second)
	CloseBuffer(buffer)

}

var wg sync.WaitGroup
var wgCon sync.WaitGroup
var isClose bool
var count int32

func CloseBuffer(buffer chan int) {
	isClose = true
	time.Sleep(2 * time.Second)
	close(buffer)
	fmt.Println("all is ready")

	wg.Wait()
	wgCon.Wait()
	fmt.Println("all is ok")
}

// 生产者速率  100ms/op
func Producer(buffer chan int, num int) {
	for !isClose {
		atomic.AddInt32(&count, 1)
		buffer <- int(count)
		time.Sleep(time.Microsecond * 10)
	}
	fmt.Println("producer will done", num)
	wg.Done()

}

// 消费者速率  1s/op
func Consumer(buffer chan int, num int) {
	for v := range buffer {
		fmt.Println(v)
		time.Sleep(time.Second)
	}
	fmt.Println("consumer will done", num)
	wgCon.Done()

}
