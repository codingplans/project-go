package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"testgo/modgo/pprofin"
	"testgo/modgo/pprofin/allocimpl"
	"testgo/modgo/pprofin/blockimpl"
	"testgo/modgo/pprofin/groutineimpl"
	"testgo/modgo/pprofin/profileimpl"
	"testing"
	"time"
)

// go run pprof_func.go
func TestPprofFunc(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)

	runtime.GOMAXPROCS(1)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		if err := http.ListenAndServe(":8888", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	// 初始化测试内
	pprofs := []pprofin.Pprof{
		&allocimpl.PprofAlloc{Buf: make([][]byte, 0)}, // 内存优化
		&groutineimpl.PprofGoroutine{},                // 协程优化
		&profileimpl.PprofProfile{},                   // CPU优化
		&blockimpl.PproBlock{},                        // 锁阻塞优化
	}
	fmt.Printf("%s\n", "---start---")
	for {
		for _, p := range pprofs {
			p.DoPprof()
			time.Sleep(time.Second)
		}
	}
}
