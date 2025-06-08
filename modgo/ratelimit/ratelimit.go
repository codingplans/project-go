package ratelimit

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 令牌桶算法实现
type TokenBucket struct {
	capacity   int64      // 桶容量
	tokens     int64      // 当前令牌数
	refillRate int64      // 每秒填充率
	lastRefill time.Time  // 上次填充时间
	mutex      sync.Mutex // 互斥锁
}

// 创建令牌桶
func NewTokenBucket(capacity, refillRate int64) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity, // 初始时桶是满的
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

// 尝试获取令牌
func (tb *TokenBucket) TryAcquire(tokens int64) bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	// 计算需要添加的令牌数
	tb.refill()

	if tb.tokens >= tokens {
		tb.tokens -= tokens
		return true
	}
	return false
}

// 填充令牌
func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)

	// 计算应该添加的令牌数
	tokensToAdd := int64(elapsed.Seconds()) * tb.refillRate
	if tokensToAdd > 0 {
		tb.tokens = min(tb.capacity, tb.tokens+tokensToAdd)
		tb.lastRefill = now
	}
}

// 获取当前令牌数
func (tb *TokenBucket) GetTokens() int64 {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	tb.refill()
	return tb.tokens
}

// 熔断器状态
type CircuitState int32

const (
	StateClosed CircuitState = iota
	StateOpen
	StateHalfOpen
)

func (s CircuitState) String() string {
	switch s {
	case StateClosed:
		return "CLOSED"
	case StateOpen:
		return "OPEN"
	case StateHalfOpen:
		return "HALF_OPEN"
	default:
		return "UNKNOWN"
	}
}

// 熔断器配置
type CircuitBreakerConfig struct {
	MaxRequests   uint32                                                // 半开状态下最大请求数
	Interval      time.Duration                                         // 统计窗口时间
	Timeout       time.Duration                                         // 熔断器打开后的超时时间
	ReadyToTrip   func(counts Counts) bool                              // 判断是否应该打开熔断器
	OnStateChange func(name string, from CircuitState, to CircuitState) // 状态变化回调
}

// 统计信息
type Counts struct {
	Requests             uint32
	TotalSuccesses       uint32
	TotalFailures        uint32
	ConsecutiveSuccesses uint32
	ConsecutiveFailures  uint32
}

// 熔断器实现
type CircuitBreaker struct {
	name          string
	maxRequests   uint32
	interval      time.Duration
	timeout       time.Duration
	readyToTrip   func(counts Counts) bool
	onStateChange func(name string, from CircuitState, to CircuitState)

	mutex      sync.Mutex
	state      CircuitState
	generation uint64
	counts     Counts
	expiry     time.Time
}

// 创建熔断器
func NewCircuitBreaker(name string, config CircuitBreakerConfig) *CircuitBreaker {
	cb := &CircuitBreaker{
		name:          name,
		maxRequests:   config.MaxRequests,
		interval:      config.Interval,
		timeout:       config.Timeout,
		readyToTrip:   config.ReadyToTrip,
		onStateChange: config.OnStateChange,
	}

	cb.toNewGeneration(time.Now())
	return cb
}

// 执行请求
func (cb *CircuitBreaker) Execute(req func() (interface{}, error)) (interface{}, error) {
	generation, err := cb.beforeRequest()
	if err != nil {
		return nil, err
	}

	defer func() {
		e := recover()
		if e != nil {
			cb.afterRequest(generation, false)
			panic(e)
		}
	}()

	result, err := req()
	cb.afterRequest(generation, err == nil)
	return result, err
}

// 请求前检查
func (cb *CircuitBreaker) beforeRequest() (uint64, error) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	now := time.Now()
	state, generation := cb.currentState(now)

	if state == StateOpen {
		return generation, errors.New("circuit breaker is open")
	} else if state == StateHalfOpen && cb.counts.Requests >= cb.maxRequests {
		return generation, errors.New("circuit breaker is half-open and max requests reached")
	}

	cb.counts.onRequest()
	return generation, nil
}

// 请求后处理
func (cb *CircuitBreaker) afterRequest(before uint64, success bool) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	now := time.Now()
	state, generation := cb.currentState(now)
	if generation != before {
		return
	}

	if success {
		cb.onSuccess(state, now)
	} else {
		cb.onFailure(state, now)
	}
}

// 获取当前状态
func (cb *CircuitBreaker) currentState(now time.Time) (CircuitState, uint64) {
	switch cb.state {
	case StateClosed:
		if !cb.expiry.IsZero() && cb.expiry.Before(now) {
			cb.toNewGeneration(now)
		}
	case StateOpen:
		if cb.expiry.Before(now) {
			cb.setState(StateHalfOpen, now)
		}
	}
	return cb.state, cb.generation
}

// 成功处理
func (cb *CircuitBreaker) onSuccess(state CircuitState, now time.Time) {
	cb.counts.onSuccess()

	if state == StateHalfOpen {
		cb.setState(StateClosed, now)
	}
}

// 失败处理
func (cb *CircuitBreaker) onFailure(state CircuitState, now time.Time) {
	cb.counts.onFailure()

	switch state {
	case StateClosed:
		if cb.readyToTrip(cb.counts) {
			cb.setState(StateOpen, now)
		}
	case StateHalfOpen:
		cb.setState(StateOpen, now)
	}
}

// 设置状态
func (cb *CircuitBreaker) setState(state CircuitState, now time.Time) {
	if cb.state == state {
		return
	}

	prev := cb.state
	cb.state = state

	cb.toNewGeneration(now)

	if cb.onStateChange != nil {
		cb.onStateChange(cb.name, prev, state)
	}
}

// 新的统计周期
func (cb *CircuitBreaker) toNewGeneration(now time.Time) {
	cb.generation++
	cb.counts.clear()

	var zero time.Time
	switch cb.state {
	case StateClosed:
		if cb.interval == 0 {
			cb.expiry = zero
		} else {
			cb.expiry = now.Add(cb.interval)
		}
	case StateOpen:
		cb.expiry = now.Add(cb.timeout)
	default:
		cb.expiry = zero
	}
}

// 获取状态
func (cb *CircuitBreaker) State() CircuitState {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	now := time.Now()
	state, _ := cb.currentState(now)
	return state
}

// 统计信息方法
func (c *Counts) onRequest() {
	c.Requests++
}

func (c *Counts) onSuccess() {
	c.TotalSuccesses++
	c.ConsecutiveSuccesses++
	c.ConsecutiveFailures = 0
}

func (c *Counts) onFailure() {
	c.TotalFailures++
	c.ConsecutiveFailures++
	c.ConsecutiveSuccesses = 0
}

func (c *Counts) clear() {
	c.Requests = 0
	c.TotalSuccesses = 0
	c.TotalFailures = 0
	c.ConsecutiveSuccesses = 0
	c.ConsecutiveFailures = 0
}

// 限流熔断服务
type RateLimitService struct {
	tokenBucket    *TokenBucket
	circuitBreaker *CircuitBreaker
	requestCount   int64
}

// 创建限流熔断服务
func NewRateLimitService() *RateLimitService {
	// 创建令牌桶：容量10，每秒填充5个令牌
	tokenBucket := NewTokenBucket(10, 5)

	// 创建熔断器
	circuitBreaker := NewCircuitBreaker("demo-service", CircuitBreakerConfig{
		MaxRequests: 3,
		Interval:    10 * time.Second,
		Timeout:     5 * time.Second,
		ReadyToTrip: func(counts Counts) bool {
			// 连续失败3次或失败率超过50%就打开熔断器
			return counts.ConsecutiveFailures >= 3 ||
				(counts.Requests >= 5 && counts.TotalFailures*2 > counts.Requests)
		},
		OnStateChange: func(name string, from CircuitState, to CircuitState) {
			log.Printf("熔断器 [%s] 状态变化: %s -> %s", name, from, to)
		},
	})

	return &RateLimitService{
		tokenBucket:    tokenBucket,
		circuitBreaker: circuitBreaker,
	}
}

// 处理请求
func (rls *RateLimitService) HandleRequest(ctx context.Context) (string, error) {
	// 首先检查令牌桶限流
	if !rls.tokenBucket.TryAcquire(1) {
		return "", errors.New("request rate limited")
	}

	// 通过熔断器执行请求
	result, err := rls.circuitBreaker.Execute(func() (interface{}, error) {
		return rls.simulateBusinessLogic(ctx)
	})

	if err != nil {
		return "", err
	}

	return result.(string), nil
}

// 模拟业务逻辑
func (rls *RateLimitService) simulateBusinessLogic(ctx context.Context) (string, error) {
	requestID := atomic.AddInt64(&rls.requestCount, 1)

	// 模拟处理时间
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	// 模拟30%的失败率
	if rand.Float32() < 0.3 {
		return "", fmt.Errorf("business logic failed for request %d", requestID)
	}

	return fmt.Sprintf("success response for request %d", requestID), nil
}

// 获取服务状态
func (rls *RateLimitService) GetStatus() map[string]interface{} {
	return map[string]interface{}{
		"tokens_available":      rls.tokenBucket.GetTokens(),
		"circuit_breaker_state": rls.circuitBreaker.State().String(),
		"total_requests":        atomic.LoadInt64(&rls.requestCount),
	}
}

// 演示程序
func Tcase() {
	service := NewRateLimitService()

	log.Println("=== 限流熔断演示开始 ===")

	// 启动状态监控
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				status := service.GetStatus()
				log.Printf("服务状态 - 可用令牌: %v, 熔断器状态: %v, 总请求数: %v",
					status["tokens_available"],
					status["circuit_breaker_state"],
					status["total_requests"])
			}
		}
	}()

	// 模拟并发请求
	var wg sync.WaitGroup

	// 第一波请求：正常频率
	log.Println("\n--- 第一波：正常请求频率 ---")
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			ctx := context.Background()
			result, err := service.HandleRequest(ctx)
			if err != nil {
				log.Printf("请求 %d 失败: %v", id, err)
			} else {
				log.Printf("请求 %d 成功: %s", id, result)
			}
		}(i)

		time.Sleep(200 * time.Millisecond)
	}

	wg.Wait()
	time.Sleep(3 * time.Second)

	// 第二波请求：高频率触发限流
	log.Println("\n--- 第二波：高频率请求（触发限流）---")
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			ctx := context.Background()
			result, err := service.HandleRequest(ctx)
			if err != nil {
				log.Printf("高频请求 %d 失败: %v", id, err)
			} else {
				log.Printf("高频请求 %d 成功: %s", id, result)
			}
		}(i)

		time.Sleep(50 * time.Millisecond) // 快速请求
	}

	wg.Wait()
	time.Sleep(5 * time.Second)

	// 第三波请求：等待令牌恢复后继续
	log.Println("\n--- 第三波：令牌恢复后的请求 ---")
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			ctx := context.Background()
			result, err := service.HandleRequest(ctx)
			if err != nil {
				log.Printf("恢复请求 %d 失败: %v", id, err)
			} else {
				log.Printf("恢复请求 %d 成功: %s", id, result)
			}
		}(i)

		time.Sleep(300 * time.Millisecond)
	}

	wg.Wait()

	// 等待观察熔断器状态变化
	time.Sleep(10 * time.Second)

	log.Println("\n=== 演示结束 ===")
	log.Printf("最终状态: %+v", service.GetStatus())
}
