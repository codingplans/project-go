package main

import (
	"context"
	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/container/pool"
	xtime "github.com/go-kratos/kratos/pkg/time"
	"github.com/prometheus/common/log"
	"time"
)

var RS *redis.Redis

// func main() {
// 	RS.Kdo("set", "www", 123)
// 	aa, _ := redis.String(RS.Kdo("get", "www"))
//
// 	log.Info(aa)
// 	// RS
//
// }

var (
	cfg redis.Config
	ct  paladin.Map
)

// func init() {
// 	NewRedis()
// }

func NewRedis() (r *redis.Redis, cf func(), err error) {
	// if err = paladin.Get("redis.toml").Unmarshal(&ct); err != nil {
	// 	return
	// }
	// if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
	// 	return
	// }

	cfg = redis.Config{
		Addr:         "127.0.0.1:6379",
		DialTimeout:  xtime.Duration(90 * time.Second),
		ReadTimeout:  xtime.Duration(90 * time.Second),
		WriteTimeout: xtime.Duration(90 * time.Second),
		SlowLog:      xtime.Duration(90 * time.Second),
		Name:         "user_center",
		Proto:        "tcp",
		// *pool.Config: &pool.Config{Active: 12},
	}
	cfg.Config = &pool.Config{
		Active:      10,
		Idle:        5,
		IdleTimeout: xtime.Duration(90 * time.Second),
	}
	RS = redis.NewRedis(&cfg)
	cf = func() { r.Close() }
	err = ping(RS)
	return RS, cf, err
}

func ping(r *redis.Redis) (err error) {
	if _, err = r.Do(context.TODO(), "SET", "ping", "pong"); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}
