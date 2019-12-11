package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"iceberg/frame/icelog"
	"laoyuegou.pb/godgame/model"
	model2 "laoyuegou.pb/plorder/model"
	"os"
	"path/filepath"
	"testgo/config"
	"time"
)

func main() {
	CalculateScore(206603129, 4, 200)
	//TestSql()
	// res := GetGodPotentialLevel(10592941, 15)
	// icelog.Infof("%+v", res)
}

type FundWater struct {
	Prices     int
	Discounts  int
	TotalScore int     // 总分数
	Repurchase float32 // 复购率
	TotalWater int     // 总流水
}

func TestSql() {
	dao := ConnMysql()

	var orders []model2.Order
	err := dao.dbr.Table("play_order").Where("buyer = ?", 1896).
		Where("state in (?)", []int64{1, 2, 3, 4, 5, 7, 8}).First(1).Error
	// Limit(1).
	// Find(&orders).Error
	if err != nil {
		panic(err.Error())
	}
	icelog.Info(orders)
	return
}

// 计算分数
func CalculateScore(godId, gameId, days int64) model.StatisticsLevel {
	dao := ConnMysql()

	days = 0 - days
	end_time := time.Now().AddDate(0, 0, int(days)).Unix()
	// 获取总流水
	var TotalWater int
	var Water model.StatisticsLevel
	err := dao.dbr.Table("play_order").
		Select("sum(price) as  prices,sum(discount) as discounts").
		Where("god=? and game_id=? and state=? and create_time > ?", godId, gameId, 8, end_time).
		First(&Water).Error
	if err == nil {
		TotalWater = Water.Prices - Water.Discounts
	}

	// 复购率
	var OrderBuy []model.StatisticsLevel
	err = dao.dbr.Table("play_order").
		Select("count(*) as prices,buyer as discounts").
		Where("god=? and game_id=?  and state=? and create_time > ?", godId, gameId, 8, end_time).
		Group("buyer").
		Find(&OrderBuy).Error
	if err == nil {

	}
	// 接单人数
	number := len(OrderBuy)
	UserNum1, UserNum2, totalMoney := 0, 0, 0
	for i := 0; i < number; i++ {
		UserNum2++
		if OrderBuy[i].Prices > 1 {
			UserNum1++
		}
		totalMoney += OrderBuy[i].Prices
	}
	repurchase := (float32(UserNum1) / float32(UserNum2)) * 100
	icelog.Info(repurchase, UserNum1, UserNum2)
	icelog.Infof("%+v", OrderBuy)

	if repurchase < 0 || UserNum1 == 0 || UserNum2 == 0 {
		icelog.Info(repurchase, "^^^^^&&")
		repurchase = 0
	}
	// 历史评分
	var Score model.StatisticsLevel
	err = dao.dbr.Table("play_order_comment").
		Select("sum(score) as total_score").
		Where("god_id=? and create_time > ?", godId, end_time).
		First(&Score).Error
	return model.StatisticsLevel{
		TotalScore:  Score.TotalScore,
		Repurchase:  int(repurchase),
		TotalWater:  TotalWater,
		TotalNumber: number,
	}

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
	cfgFile  = flag.String("config-path", "/Users/darren/go/src/testgo/config.json", "config file")
	logLevel = flag.String("loglevel", "DEBUG", "log level")
)
