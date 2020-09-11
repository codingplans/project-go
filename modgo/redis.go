package main

import (
	"context"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
)

var RS *redis.Redis

func main() {
	NewRedis()

	// RS

}

var (
	cfg redis.Config
	ct  paladin.Map
)

func NewRedis() (r *redis.Redis, cf func(), err error) {
	if err = paladin.Get("redis.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
		return
	}
	RS = redis.NewRedis(&cfg)
	cf = func() { r.Close() }
	err = ping(r)
	return
}

func ping(r *redis.Redis) (err error) {
	if _, err = r.Do(context.TODO(), "SET", "ping", "pong"); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}
