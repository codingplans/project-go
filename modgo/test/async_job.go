package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AsyncJob struct {
	ch        chan interface{}
	waitGroup sync.WaitGroup
	isStop    int32
	errs      []error
}

func NewAsyncJob() *AsyncJob {
	return &AsyncJob{
		ch: make(chan interface{}, 10000),
	}
}

func (j *AsyncJob) Push(i interface{}) {
	j.ch <- i
}

func (j *AsyncJob) RunParallel(total int, fn func(v interface{})) {
	if total == 0 {
		total = 1
	}
	j.waitGroup.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			defer j.waitGroup.Done()
			for v := range j.ch {
				func() {
					defer func() {
						if err := recover(); err != nil {
							j.errs = append(j.errs, fmt.Errorf("Async job panic: %v", err))
						}
					}()
					if atomic.LoadInt32(&j.isStop) == 1 {
						return
					}
					fn(v)
				}()
			}
		}()
	}
}

func (j *AsyncJob) Stop() {
	atomic.SwapInt32(&j.isStop, 1)
}

func (j *AsyncJob) Wait() {
	close(j.ch)
	j.waitGroup.Wait()
}

func (j *AsyncJob) Errors() []error {
	return j.errs
}
