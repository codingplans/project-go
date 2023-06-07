package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	_ "net/http/pprof"
	"sync"
	"time"
)

// StartWork模拟一个http请求 。
// 这个请求中会并发30次做个workFn， workFn类似第三方接口，返回响应时间不确定，但是3s内的响应都算成功，返回结果到主任务中。
// 主任务等待3s后，结束任务，继续后续流程
// 3s以外结束的workFn，后续会做惩戒处理，做好计数后，做限流操作  如： 1分钟/500 =》 1分钟/20
func main() {
	StartWork()
	log.Println("截止时间3s 结束3s内的任务，继续后续流程")
	time.Sleep(time.Second * 10)
	// 这里模拟后面流程，只做打印
	for _, i2 := range dsp.reply {
		log.Println(i2)
	}
}

type dspExchange struct {
	reply []string
	rLock sync.Mutex
	wg    sync.WaitGroup
}

var dsp *dspExchange = &dspExchange{
	reply: make([]string, 0, 30),
	wg:    sync.WaitGroup{},
}

// 启动30个协程， 分别作业，随机延时， 返回3s内成功的结果，到主任务中
func StartWork() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	for i := 0; i < 30; i++ {
		dsp.wg.Add(1)
		go workFn(ctx)
	}

	// 这里是阻塞等待所有workFn都在1s内结束，主动cancel CTX，结束StartWork
	go func() {
		dsp.wg.Wait()
		cancel()
	}()
	// 这里等待2种信号，一种是3s超时，一种是所有workFn都在1s内结束
	select {
	case <-ctx.Done():
		return
	}

}

func workFn(ctx2 context.Context) {
	// 设置上下文主动取消 增加原因用于惩戒判定
	ctx, pcancel := context.WithCancelCause(ctx2)
	// 小作业根据配置时间，设置超时时间，超时后，主动取消，此处默认设置2s
	ctx, cc := context.WithTimeout(ctx, 2*time.Second)
	defer func() {
		dsp.wg.Done()
		cc()
	}()

	// 模拟第三方接口，随机延时，返回结果错误原因塞入err中
	fc := func(ctx context.Context, ctl context.CancelCauseFunc) {
		ts := time.Millisecond * 100 * time.Duration(rand.Intn(80))
		// ts := time.Millisecond * 10 * time.Duration(rand.Intn(80))
		defer func() {
			var err error
			if ts.Seconds() > 1 {
				err = fmt.Errorf("超过1s且小于2s的记录下来，后面给予惩戒 %s", ts.String())
			} else {
				err = fmt.Errorf("小于1s的不记录，后面不给予惩戒 %s", ts.String())
			}
			ctl(err)
		}()
		// 模拟三方请求的随机响应时间
		time.Sleep(ts)

		if ts.Seconds() > 1 {
			log.Println("主work已结束，这是当前workFn结束时间打印下", ts.String())
		}
	}
	// 丢协程执行请求
	go fc(ctx, pcancel)

	// 阻塞等待ctx 信号
	select {
	case <-ctx.Done():
		dsp.rLock.Lock()
		defer dsp.rLock.Unlock()
		err := context.Cause(ctx)
		// 主动cancel不需要记录
		if err != nil && err != context.Canceled && err.Error() != context.DeadlineExceeded.Error() {
			// 如果ctx被超时，则结束函数的执行
			dsp.reply = append(dsp.reply, " 小作业结束："+err.Error())
		}

	}
}
