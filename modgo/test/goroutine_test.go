package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

//  多消费多生产 场景 下通过 wg 控制生产者关闭 和消费
func TestGoroutine(t *testing.T) {
	buffer := make(chan int, 100)
	for i := 0; i < 40; i++ {
		go Producer(buffer)
		wg.Add(1)
	}
	for i := 0; i < 40; i++ {
		go Consumer(buffer)
		wgCon.Add(1)
	}

	time.Sleep(3 * time.Second)
	CloseBuffer(buffer)

}

var wg sync.WaitGroup
var wgCon sync.WaitGroup
var isClose bool

func CloseBuffer(buffer chan int) {
	isClose = true
	time.Sleep(2 * time.Second)
	close(buffer)
	println("all is ready")

	wg.Wait()
	wgCon.Wait()
	println("all is ok")
}
func Producer(buffer chan int) {
	for !isClose {
		buffer <- rand.Int()
		time.Sleep(time.Microsecond * 100)
	}
	wg.Done()
}

func Consumer(buffer chan int) {
	for v := range buffer {
		println(v)
		time.Sleep(time.Second)
	}
	wgCon.Done()
	println("consumer will done")

}
