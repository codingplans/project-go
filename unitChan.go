package main

import (
	"fmt"
	"sync"
	"time"
)

// HandlerFunc 被执行的异步任务
type HandlerFunc func()

// Executor 异步任务执行器
type Executor interface {
	// Start 启动执行器
	Start()

	// Execute 执行一个异步任务
	Execute(handler HandlerFunc)

	// WaitAndStop 阻塞等待所有任务完成后，关闭执行器
	WaitAndStop()
}

func newExecutor(runnerCount int64) Executor {
	return asyncExecutor{
		pool:          make(chan HandlerFunc),
		stopChan:      make(chan bool),
		jobsWaitGroup: new(sync.WaitGroup),
		runnerCount:   runnerCount,
	}
}

type asyncExecutor struct {
	pool          chan HandlerFunc
	stopChan      chan bool
	jobsWaitGroup *sync.WaitGroup
	runnerCount   int64
}

func (e asyncExecutor) Start() {
	// TODO
	go func() {
	CC:
		for {
			println("当前运行数：", e.runnerCount)
			select {
			case obj := <-e.pool:
				obj()
				// e.runnerCount--
			case <-e.stopChan:
				println("close chan")
				break
			default:
				println("conti")
				goto CC
			}
		}
	}()

}

func (e asyncExecutor) Execute(handler HandlerFunc) {
	// TODO
	e.jobsWaitGroup.Add(1)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("err: %+v", r)
		}
	}()
	go func() {
		e.pool <- handler
		e.jobsWaitGroup.Done()
		e.runnerCount++
	}()

}

func (e asyncExecutor) WaitAndStop() {
	// TODO
	e.jobsWaitGroup.Wait()
	if e.runnerCount == 0 {
		e.stopChan <- true
	}
	close(e.pool)
	defer close(e.stopChan)
	println("over~")

}

func main() {
	executor := newExecutor(3)
	executor.Start()

	executor.Execute(func() {
		fmt.Println("async job")
	})

	executor.Execute(func() {
		time.Sleep(2 * time.Second)
		fmt.Println("time-consuming job 1")
	})

	executor.Execute(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("time-consuming job 2")
	})

	executor.WaitAndStop()

}
