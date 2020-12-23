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

func main() {
	NewRedis()
	Rkwithdrew()
}

func Rkwithdrew() {
	tt := time.Now().Format("20060102")
	uid := int64(3505)
	keys := make([]string, 20)
	key := RedisHashGoodsOnce(uid, 0)
	keys = append(keys, key, "day_withdraw_num_"+tt+"_com.money.calendarweather.lite")
	keys = append(keys, key, "idtf_gods_5fb8489e-7fc2-11ea-899f-dca90496ac471")
	keys = append(keys, key, "5fb8489e-7fc2-11ea-899f-dca90496ac471")
	keys = append(keys, key, "bank_acnt_com.money.calendarweather.lite_origin_wxpay_test") // 提现账户
	keys = append(keys, key, "day_withdraw_num_"+tt+"_com.money.calendarweather.lite")
	keys = append(keys, key, "user_goods_once_"+fmt.Sprintf("%d", uid)+"")
	keys = append(keys, key, "user_goods_daily_"+fmt.Sprintf("%d", uid)+"_"+tt+"")
	keys = append(keys, key, "account_"+fmt.Sprintf("%d", uid)+"")
	keys = append(keys, key, "user_ticket_"+fmt.Sprintf("%d", uid)+"")
	keys = append(keys, key, "notice_flow_com.money.calendarweather.lite")
	keys = append(keys, key, "daily_income_"+fmt.Sprintf("%d", uid)+"_"+tt+"")

	for _, v := range keys {
		_, erre := RSs.Kdo("del", v)
		println(erre)
	}
}

var (
	cfg redis.Config
	ct  paladin.Map
	RSs *redis.Redis
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

//
// // 提现某个商品的次数
// func RedisHashGoodsOnce(uid int64, group int64) string {
// 	if group == 0 {
// 		return fmt.Sprintf("user_goods_once_%d", uid)
// 	} else {
// 		return fmt.Sprintf("user_goods_once_%d_%d", uid, group)
// 	}
// }
//
// // 每日商品
// func RedisHashGoodsDaily(uid int64, day string, group int64) string {
// 	if group == 0 {
// 		return fmt.Sprintf("user_goods_daily_%d_%s", uid, day)
// 	} else {
// 		return fmt.Sprintf("user_goods_daily_%d_%s_%d", uid, day, group)
// 	}
// }
//
// // 商品每日被提现次数
// func RedisHashDayGoodsWithdrawNum(day string, appName string) string {
// 	// 2020/5/22 0:0:0开启
// 	if time.Now().Unix() >= 1590076800 {
// 		return fmt.Sprintf("day_withdraw_num_%s_%s", day, appName)
// 	} else {
// 		return fmt.Sprintf("day_withdraw_num_%s", day)
// 	}
// }
//
// // 设备下提现的商品数量
// func RedisHashIdentifierWithdrawGoodsNum(identifier string) string {
// 	return fmt.Sprintf("idtf_gods_%s", identifier)
// }
//
// // 提现账户对应的uid
// func RedisKvAppBankAccountUid(appName, bank, bankAccount string) string {
// 	return fmt.Sprintf("bank_acnt_%s_%s_%s", appName, bank, bankAccount)
// }
