package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
	"os"
	"path/filepath"
	"testgo/modgo/model"
	"time"
)

var DB *gorm.DB

func main() {

	userId := int64(1919)
	coinId := int64(userId)

	wallet := &model.WalletAddress{}
	DB.Table("wallet_address").
		Where("member_id=?", userId).
		Where("coin_id=?", coinId).
		First(&wallet)

	if wallet.ID == 0 {
		wallet.Address = "1212"
		wallet.CoinID = coinId
		wallet.CreateTime = time.Now()
		wallet.ProviderID = "12"
		wallet.MemberID = userId
		DB.Table("wallet_address").Create(&wallet)
	}

	println(wallet.ID)
	// formatss()

}

func formatss() {
	var mm []*model.OtcAdvertise
	ttt := int64(1593328073)
	aaa := time.Unix(ttt, 0).Format("2006-01-02 15:04:05")
	sss := time.Unix(ttt-800000, 0).Format("2006-01-02 15:04:05")
	err := DB.Table("otc_advertise").
		Where("create_at <=?", aaa).
		Where("create_at >=?", sss).
		Find(&mm).Error

	log.Info(len(mm), err, aaa)

	for _, v := range mm {
		log.Info(v.Type)

	}
}

func calldb(mm model.Member) (err error) {

	mm.Status = 12

	kk := mm
	dd := DB.Begin()
	err = dd.Table("member").Save(&kk).Error
	if err != nil {
		return
	}
	return dd.Commit().Error
}

func init() {
	ConnMysql()
}

func ConnMysql() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	os.Chdir(dir)
	flag.Parse()
	cfg := "dbuser:pass!23word@tcp(192.168.3.8:3306)/user_center?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
	// cfg := "root:root@tcp(127.0.0.1:23306)/user_center?timeout=1s&readTimeout=3s&writeTimeout=3s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
	var err error
	DB, err = gorm.Open("mysql", cfg)
	if err != nil {
		panic(err.Error())
	}
	DB.LogMode(true)
}

// SqlStore mysql读/写库客户端封装
type SqlStore struct {
	*gorm.DB
}
