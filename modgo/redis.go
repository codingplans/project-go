package main

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/container/pool"
	xtime "github.com/go-kratos/kratos/pkg/time"
	"github.com/prometheus/common/log"
	"time"
)

var RS *redis.Redis

func main() {
	NewRedis()
	// RS.Kdo("set", "www", 123)
	// aa, _ := redis.String(RS.Kdo("get", "www"))
	//
	// log.Info(aa)

	uid := int64(8810)
	keys := []string{}
	key := RedisHashGoodsOnce(uid, 0)
	// key2 := RedisHashGoodsOnce(8811, 0)
	keys = append(keys, key, "day_withdraw_num_20201209_com.money.calendarweather.lite")
	for _, v := range keys {
		RS.Kdo("del", v)
	}
}

var (
	cfg redis.Config
	ct  paladin.Map
)

func NewRedis() (r *redis.Redis, cf func(), err error) {
	// if err = paladin.Get("redis.toml").Unmarshal(&ct); err != nil {
	// 	return
	// }
	// if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
	// 	return
	// }

	cfg = redis.Config{
		// Addr:         "127.0.0.1:6379",
		Addr:         "192.168.10.15:8101",
		DialTimeout:  xtime.Duration(90 * time.Second),
		ReadTimeout:  xtime.Duration(90 * time.Second),
		WriteTimeout: xtime.Duration(90 * time.Second),
		SlowLog:      xtime.Duration(90 * time.Second),
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

// 提现某个商品的次数
func RedisHashGoodsOnce(uid int64, group int64) string {
	if group == 0 {
		return fmt.Sprintf("user_goods_once_%d", uid)
	} else {
		return fmt.Sprintf("user_goods_once_%d_%d", uid, group)
	}
}

// 每日商品
func RedisHashGoodsDaily(uid int64, day string, group int64) string {
	if group == 0 {
		return fmt.Sprintf("user_goods_daily_%d_%s", uid, day)
	} else {
		return fmt.Sprintf("user_goods_daily_%d_%s_%d", uid, day, group)
	}
}

// 商品每日被提现次数
func RedisHashDayGoodsWithdrawNum(day string, appName string) string {
	// 2020/5/22 0:0:0开启
	if time.Now().Unix() >= 1590076800 {
		return fmt.Sprintf("day_withdraw_num_%s_%s", day, appName)
	} else {
		return fmt.Sprintf("day_withdraw_num_%s", day)
	}
}

// 设备下提现的商品数量
func RedisHashIdentifierWithdrawGoodsNum(identifier string) string {
	return fmt.Sprintf("idtf_gods_%s", identifier)
}

// 提现账户对应的uid
func RedisKvAppBankAccountUid(appName, bank, bankAccount string) string {
	return fmt.Sprintf("bank_acnt_%s_%s_%s", appName, bank, bankAccount)
}
