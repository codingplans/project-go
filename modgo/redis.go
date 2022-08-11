package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/container/pool"
	xtime "github.com/go-kratos/kratos/pkg/time"
	"github.com/prometheus/common/log"
)

func main() {
	// 定义命令行参数方式1
	patten := ""
	flag.StringVar(&connDrr, "addr", "", "redis 连接地址")
	flag.StringVar(&patten, "key", "name", "关键字")
	// 解析命令行参数
	flag.Parse()
	NewRedis()
	// flag.Parse()
	if patten != "" {
		res, _ := SCAN(patten)
		fmt.Println(res)
	}
	res, _ := SCAN(patten)
	fmt.Println("总共结果：", len(res))

	for _, v := range res {
		fmt.Println(v)
	}
}

var (
	cfg redis.Config
	ct  paladin.Map
	RSs *redis.Redis
)

var connDrr string

func NewRedis() (r *redis.Redis, cf func(), err error) {
	cfg = redis.Config{
		Addr:         "127.0.0.1:6379",
		DialTimeout:  xtime.Duration(90 * time.Second),
		ReadTimeout:  xtime.Duration(90 * time.Second),
		WriteTimeout: xtime.Duration(90 * time.Second),
		SlowLog:      xtime.Duration(90 * time.Second),
		Proto:        "tcp",
		// *pool.Config: &pool.Config{Active: 12},
	}
	if connDrr != "" {
		println("当前连接：", connDrr)
		cfg.Addr = connDrr
	}
	cfg.Config = &pool.Config{
		Active:      10,
		Idle:        5,
		IdleTimeout: xtime.Duration(90 * time.Second),
	}
	RSs = redis.NewRedis(&cfg)
	cf = func() { r.Close() }
	err = ping(RSs)
	return RSs, cf, err
}

func ping(r *redis.Redis) (err error) {
	if _, err = r.Do(context.TODO(), "SET", "ping", "pong"); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}

// SCAN 获取大量key  代替KEYS命令
func SCAN(patten string) ([]string, error) {
	var out []string
	var cursor uint64 = 0xffffff
	isfirst := true
	for cursor != 0 {
		if isfirst {
			cursor = 0
			isfirst = false
		}
		arr, err := RSs.Do(context.TODO(), "SCAN", cursor, "MATCH", patten, "COUNT", 100)
		if err != nil {
			return out, err
		}
		switch arr := arr.(type) {
		case []interface{}:
			cursor, _ = redis.Uint64(arr[0], nil)
			it, _ := redis.Strings(arr[1], nil)
			out = append(out, it...)
		}
	}
	out = ArrayDuplice(out)
	return out, nil
}

// ArrayDuplice 数组去重
func ArrayDuplice(arr []string) []string {
	var out []string
	tmp := make(map[string]byte)
	for _, v := range arr {
		tmplen := len(tmp)
		tmp[v] = 0
		if len(tmp) != tmplen {
			out = append(out, v)
		}
	}
	return out
}
