package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"iceberg/frame/icelog"
	"os"
	"path/filepath"
	"test/config"
	"time"
)

func main() {
	TestSql()
}

type FundWater struct {
	Prices    int64
	Discounts int64
}

func TestSql() {
	sql := ConnMysql()

	end_time := time.Now().AddDate(0, 0, -140).Unix()
	// 获取总流水
	var TotalWater int64
	var Water FundWater
	err := sql.dbr.Table("play_order").
		Select("sum(price) as  prices,sum(discount) as discounts").
		Where("god=? and state=? and create_time > ?", 1896, 8, end_time).
		First(&Water).Error
	if err == nil {
		TotalWater = Water.Prices - Water.Discounts
	}

	icelog.Info(err, TotalWater, Water)

	return
}

func ConnMysql() *SqlStore {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	os.Chdir(dir)
	flag.Parse()
	config.Init(*cfgFile)
	icelog.SetLevel(*logLevel)
	mysql := config.DefaultConfig.MysqlAppMain
	sqlStore := new(SqlStore)
	dsnr := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysql.User, mysql.Psw, mysql.Host.Read, mysql.Port, mysql.DbName)
	var err error
	sqlStore.dbr, err = gorm.Open("mysql", dsnr)
	if err != nil {
		panic(err.Error())
	}
	dsnw := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysql.User, mysql.Psw, mysql.Host.Write, mysql.Port, mysql.DbName)
	sqlStore.dbw, err = gorm.Open("mysql", dsnw)
	if err != nil {
		panic(err.Error())
	}
	sqlStore.dbr.LogMode(mysql.LogMode)
	sqlStore.dbw.LogMode(mysql.LogMode)
	return sqlStore
}

// SqlStore mysql读/写库客户端封装
type SqlStore struct {
	dbr           *gorm.DB
	dbw           *gorm.DB
	db_r_app_main *gorm.DB
}

var (
	cfgFile  = flag.String("config-path", "/Users/darren/go/src/test/config.json", "config file")
	logLevel = flag.String("loglevel", "DEBUG", "log level")
)
