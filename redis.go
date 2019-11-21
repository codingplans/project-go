package main

import (
	"encoding/json"
	"flag"
	key2 "game/key"
	"github.com/gomodule/redigo/redis"
	"iceberg/frame/icelog"
	"laoyuegou.com/cache"
	"os"
	"path/filepath"
	"test/config"
)

// CacheStore 缓存组件封装
type CacheStore struct {
	Redis *cache.RedisPool
}

var (
	cfgFile2  = flag.String("config-path", "/Users/darren/go/src/test/config.json", "config file")
	logLevel2 = flag.String("loglevel", "DEBUG", "log level")
)

func main() {
	TestRedis()
}

func TestRedis() {
	re := connRedis()
	R := re.Redis.Get()
	defer R.Close()

	// data, err := R.Do("Set", "qqq", 1)
	// icelog.Info(data, err)

	var arr []int64

	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 121)
	arr = append(arr, 11)

	R.Do("Hset", "testh", "no1", 1)
	dd, err := redis.Int64(R.Do("HGET", "testh", "no1"))

	// RKQuickOrder
	key := key2.RKQuickOrder()
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
	ff, err := json.Marshal(arr)
	R.Do("Hset", key, "god_potential_level1", ff)

	tt, err := R.Do("HGET", key, "god_potential_level2")
	aa, err := redis.Bytes(tt, err)

	ee := []int64{1, 2, 3, 4}
	json.Unmarshal(aa, &ee)
	icelog.Info(err, dd, aa, tt, ee)

}

func connRedis() *CacheStore {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	os.Chdir(dir)
	flag.Parse()
	config.Init(*cfgFile2)
	icelog.SetLevel(*logLevel2)

	cacheStore := new(CacheStore)
	cacheStore.Redis = cache.NewRedisPool(config.DefaultConfig.RedisAuth.Addr, config.DefaultConfig.RedisAuth.Psw, 3, 300)
	return cacheStore
}
