package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 设置 CPU 的数量为 1
	runtime.GOMAXPROCS(1)

	// 创建一个死循环的 Goroutine
	go func() {
		for {
			fmt.Println("")
		}
	}()

	// 创建一个正常运行的 Goroutine
	go func() {
		fmt.Println("Hello, world!")
		panic(222)
	}()

	// 阻塞主线程
	fmt.Scanln()
}
