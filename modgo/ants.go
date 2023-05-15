package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/panjf2000/ants/v2"
	"github.com/valyala/fasthttp"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"
	"syscall"
	"time"
)

var (
	ok int // 成功次数
	no int // 失败次数
)
var mu = sync.RWMutex{}  // 加锁
var mus = sync.RWMutex{} // 加锁

var tunnel = make(chan string, 20000)

var wgAnt sync.WaitGroup

// 请求数量
func count() {
	for i := 0; i < 1000000; i++ {
		tunnel <- strconv.Itoa(i)
		wgAnt.Add(1)
	}
}

var fail string

func sendGetRequest() {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://localhost:8081/")
	req.Header.SetMethod(fasthttp.MethodGet)
	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	//httpClient := http.Client{Timeout: 2 * time.Second}
	//req, _ := http.NewRequest("GET", "http://localhost:8081/", nil)
	// 获取响应
	//_, err := httpClient.Do(req)
	if err == nil {
		//fmt.Printf("DEBUG Response: %s\n", resp.Body())
		mu.Lock() // 上锁
		ok++
		mu.Unlock() // 上锁

	} else {
		mu.Lock()
		no++
		mu.Unlock()
		fmt.Fprintf(os.Stderr, "ERR Connection error: %v\n", err)
	}
}

// 处理请求
func dorequest(i interface{}) {
	//sendGetRequest()
	err := GetApiData("http://localhost:8081/", nil, nil)
	if err != nil {
		mu.Lock() // 上锁
		ok++
		mu.Unlock() // 上锁
	} else {
		mu.Lock()
		no++
		mu.Unlock()
		fmt.Fprintf(os.Stderr, "ERR Connection error: %v\n", err)
	}
	fmt.Println(i)
	wgAnt.Done()

}

var client *fasthttp.Client

var httpClient = &http.Client{
	Timeout:   time.Second * 60,
	Transport: &http.Transport{DisableKeepAlives: true},
}

// GetApiData 最小化请求 ，可复用于其他三方api 作为基础依赖
func GetApiData(path string, params, header map[string]string) []byte {
	// 勿删： req.Header.Add("Connection", "close")    // 等效的关闭方式
	method := "GET"
	client := httpClient
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		fmt.Println("readFromHttp", err.Error())
		return nil
	}
	// resp, err := s.httpClt.Get(path)
	//req.Close = true
	//req.Header.Add("Connection", "close")
	q := req.URL.Query()
	for s2, s3 := range params {
		q.Add(s2, s3)
	}

	for s2, s3 := range header {
		req.Header.Add(s2, s3)
	}
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error(), resp)
		return nil
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			return
		}
	}()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error(), resp)
	}
	return b
}

func main() {
	//readTimeout, _ := time.ParseDuration("500ms")
	//writeTimeout, _ := time.ParseDuration("500ms")
	//maxIdleConnDuration, _ := time.ParseDuration("1h")
	//client = &fasthttp.Client{
	//	MaxConnsPerHost:               2000,
	//	ReadTimeout:                   readTimeout,
	//	WriteTimeout:                  writeTimeout,
	//	MaxIdleConnDuration:           maxIdleConnDuration,
	//	NoDefaultUserAgentHeader:      true, // Don't send: User-Agent: fasthttp
	//	DisableHeaderNamesNormalizing: true, // If you set the case on your headers correctly you can enable this
	//	DisablePathNormalizing:        true,
	//	// increase DNS cache time to an hour instead of default minute
	//	Dial: (&fasthttp.TCPDialer{
	//		Concurrency:      4096,
	//		DNSCacheDuration: time.Hour,
	//	}).Dial,
	//}

	g := gin.New()
	g.GET("/", func(c *gin.Context) {
		time.Sleep(30 * time.Millisecond)
		c.String(200, "ok")
	})
	go g.Run(":8081")

	//使用所有CPU核心
	runtime.GOMAXPROCS(runtime.NumCPU())

	go count()

	time.Sleep(1 * time.Millisecond)

	chairPool, _ := ants.NewPoolWithFunc(50000, dorequest) // 声明有几把电椅
	defer chairPool.Release()

	for {
		select {
		case a := <-tunnel:
			_ = chairPool.Invoke(a)

		default:
			wgAnt.Wait()
			fmt.Println("所有请求已处理完毕", ok, no)
			fmt.Println(fail)
			return
		}
	}

}

func setRlimitNOFile(nofile uint64) error {
	if nofile == 0 {
		return nil
	}
	var lim syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &lim); err != nil {
		return err
	}
	if nofile <= lim.Cur {
		return nil
	}
	lim.Cur = nofile
	lim.Max = nofile
	return syscall.Setrlimit(syscall.RLIMIT_NOFILE, &lim)
}
