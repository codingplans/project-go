package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"iceberg/frame/icelog"
	"laoyuegou.pb/godgame/model"
	model3 "laoyuegou.pb/user/model"
	"os"
	"path/filepath"
	"testgo/config"
)

func main() {
	fmt.Println(HasGodGame(18961))
	// TestSql()

	// res := GetGodPotentialLevel(10592941, 15)
	// icelog.Infof("%+v", res)
}

func TestSql() {
	dao := ConnMysql()

	userId := int64(1896)
	var data model3.PrivacyCfg
	// data := model3.PrivacyCfg{
	// 	UserID:         userId,
	// 	IsShowSHB:      2,
	// 	IsShowChatroom: 2,
	// 	IsShowDistance: 2,
	// 	IsShowNear:     2,
	// }
	err := dao.dbr.Table("privacy_cfg").Where("user_id  = ?", userId).
		Assign("is_show_chatroom", 10).
		Assign("is_show_near", 10).
		FirstOrCreate(&data).Error
	if err != nil {
	}
	icelog.Infof("%+v", &data, err)
	//
	// var orders []model2.Order
	// err := dao.dbr.Table("play_order").Where("buyer = ?", 1896).
	// 	Where("state in (?)", []int64{1, 2, 3, 4, 5, 7, 8}).First(1).Error
	// // Limit(1).
	// // Find(&orders).Error
	// if err != nil {
	// 	panic(err.Error())
	// }
	// icelog.Info(orders)
	return
}

// 是否有申请通过的品类
func HasGodGame(godId int64) bool {
	dao := ConnMysql()
	var games model.GodGame
	err := dao.dbr.Table("play_god_games").Select("gameid").Where("userid=? AND status=?", godId, 1).First(&games).Error
	if err != nil {
		icelog.Errorf("Has godgame status [%d] error:%s", godId, err.Error())
		return false
	}
	return true
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
