package main

import (
	"flag"
	"github.com/gomodule/redigo/redis"
	"iceberg/frame/icelog"
	"laoyuegou.com/cache"
	"os"
	"path/filepath"
	"plorder/key"
	"testgo/config"
)

// CacheStore 缓存组件封装
type CacheStore struct {
	Redis *cache.RedisPool
}

var (
	cfgFile2  = flag.String("config-path", "/Users/darren/go/src/testgo/config.json", "config file")
	logLevel2 = flag.String("loglevel", "DEBUG", "log level")
)

func main() {
	TestRedis()
}

func TestRedis() {
	re := connRedis()
	c := re.Redis.Get()
	res := false
	orderId := "201912111134024854111189511"
	// 通过哈希获取订单映射关系
	quickOrderId, _ := redis.String(c.Do("HGET", key.RKPlOrderIdQuickOrderIdMap(), orderId))
	// quickOrderId, _ := redis.String(c.Do("HGET", "PHP:PlOrderId:QuickOrderId:Map", orderId))
	if quickOrderId != "" {
		res, _ = redis.Bool(c.Do("sismember", key.RKIsAutoGrabOrder(quickOrderId), 1896))
	}

	icelog.Infof("%+v,%+v,%+v,%+v,%+v  &&&&&&&&&&&&", quickOrderId, res, key.RKPlOrderIdQuickOrderIdMap(), key.RKIsAutoGrabOrder(quickOrderId), orderId)

	// isGrabOrder, _ := redis.String(c.Do("hget", "PHP:PlOrderId:QuickOrderId:Map", "20191210204042485411114551"))
	// isGrabOrder, _ := redis.Bool(c.Do("sismember", core.RKGodAutoGrabGames(10593099), 4))
	// icelog.Info(isGrabOrder)

	return
	// data, err := R.Do("Set", "qqq", 1)
	// icelog.Info(data, err)

	// var arr []int64
	//
	// arr = append(arr, 1)
	// arr = append(arr, 2)
	// arr = append(arr, 121)
	// arr = append(arr, 11)
	//
	// R.Do("Hset", "testh", "no1", 1)
	// dd, err := redis.Int64(R.Do("HGET", "testh", "no1"))
	//
	// // RKQuickOrder
	// key := key2.RKQuickOrder()
	// level1, err := json.Marshal(data.GodPotentialLevel1)
	// level2, err := json.Marshal(data.GodPotentialLevel2)
	// level3, err := json.Marshal(data.GodPotentialLevel3)
	// level4, err := json.Marshal(data.GodPotentialLevel4)
	// level5, err := json.Marshal(data.GodPotentialLevel5)
	// if err == nil {
	// 	re.Do("HSET", keyQuickOrder, "god_potential_level1", level1)
	// 	re.Do("HSET", keyQuickOrder, "god_potential_level2", level2)
	// 	re.Do("HSET", keyQuickOrder, "god_potential_level3", level3)
	// 	re.Do("HSET", keyQuickOrder, "god_potential_level4", level4)
	// 	re.Do("HSET", keyQuickOrder, "god_potential_level5", level5)
	// ff, err := json.Marshal(arr)
	// R.Do("Hset", key, "god_level_time_range", 200)
	//
	// tt, err := R.Do("HGET", key, "god_level_time_range")
	// aa, err := redis.Bytes(tt, err)
	//
	// ee := []int64{1, 2, 3, 4}
	// json.Unmarshal(aa, &ee)
	// icelog.Info(err, dd, aa, tt, ee)

}

func connRedis() *CacheStore {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	os.Chdir(dir)
	flag.Parse()
	config.Init(*cfgFile2)
	icelog.SetLevel(*logLevel2)

	cacheStore := new(CacheStore)
	icelog.Info(config.DefaultConfig.RedisAuth.Addr, config.DefaultConfig.RedisAuth.Psw)
	cacheStore.Redis = cache.NewRedisPool(config.DefaultConfig.RedisAuth.Addr, config.DefaultConfig.RedisAuth.Psw, 3, 300)
	return cacheStore
}
