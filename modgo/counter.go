package main

// 限流计数器

import (
	"log"
	"math"
	"sync"
	"time"
)

type Counter struct {
	rate  int           //计数周期内最多允许的请求数
	begin time.Time     //计数开始时间
	cycle time.Duration //计数周期
	count int           //计数周期内累计收到的请求数
	lock  sync.Mutex
}

func (l *Counter) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.count == l.rate-1 {
		now := time.Now()
		if now.Sub(l.begin) >= l.cycle {
			//速度允许范围内， 重置计数器
			l.Reset(now)
			return true
		} else {
			return false
		}
	} else {
		//没有达到速率限制，计数加1
		l.count++
		return true
	}
}
func (l *Counter) Set(r int, cycle time.Duration) {
	l.rate = r
	l.begin = time.Now()
	l.cycle = cycle
	l.count = 0
}
func (l *Counter) Reset(t time.Time) {
	l.begin = t
	l.count = 0
}
func main() {
	var wg sync.WaitGroup
	var lr Counter
	lr.Set(3, time.Second) // 1s内最多请求3次
	var succ, fail int
	for i := 0; i < 20; i++ {
		wg.Add(1)
		log.Println("创建请求:", i)
		go func(i int) {
			if lr.Allow() {
				log.Println("响应请求:", i)
				succ++
			} else {
				log.Println("拒绝:", i)
				fail++

			}
			wg.Done()
		}(i)
		time.Sleep(400 * time.Millisecond)
	}
	wg.Wait()
	log.Println("succ:", succ, "fail:", fail)
}

type TokenBucket struct {
	rate         int64 //固定的token放入速率, r/s
	capacity     int64 //桶的容量
	tokens       int64 //桶中当前token数量
	lastTokenSec int64 //桶上次放token的时间戳 s

	lock sync.Mutex
}

func (l *TokenBucket) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	now := time.Now().Unix()
	l.tokens = l.tokens + (now-l.lastTokenSec)*l.rate // 先添加令牌
	if l.tokens > l.capacity {
		l.tokens = l.capacity
	}
	l.lastTokenSec = now
	if l.tokens > 0 {
		// 还有令牌，领取令牌
		l.tokens--
		return true
	} else {
		// 没有令牌,则拒绝
		return false
	}
}
func (l *TokenBucket) Set(r, c int64) {
	l.rate = r
	l.capacity = c
	l.tokens = 0
	l.lastTokenSec = time.Now().Unix()
}

type LeakyBucket struct {
	rate       float64 //固定每秒出水速率
	capacity   float64 //桶的容量
	water      float64 //桶中当前水量
	lastLeakMs int64   //桶上次漏水时间戳 ms
	lock       sync.Mutex
}

func (l *LeakyBucket) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	now := time.Now().UnixNano() / 1e6
	eclipse := float64((now - l.lastLeakMs)) * l.rate / 1000 //先执行漏水
	l.water = l.water - eclipse                              //计算剩余水量
	l.water = math.Max(0, l.water)                           //桶干了
	l.lastLeakMs = now
	if (l.water + 1) < l.capacity {
		// 尝试加水,并且水还未满
		l.water++
		return true
	} else {
		// 水满，拒绝加水
		return false
	}
}
func (l *LeakyBucket) Set(r, c float64) {
	l.rate = r
	l.capacity = c
	l.water = 0
	l.lastLeakMs = time.Now().UnixNano() / 1e6
}
