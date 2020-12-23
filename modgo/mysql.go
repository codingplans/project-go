package main

import (
	"fmt"
	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/container/pool"
	xtime "github.com/go-kratos/kratos/pkg/time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
	"strconv"
	"strings"
	"testgo/modgo/model"
	"time"
)

var DB *gorm.DB
var DBB gorm.SQLCommon
var RS *redis.Redis

func main() {
	// GetRecommend(5)
	// RS.Kdo("set", "www", 123)
	// aa, _ := redis.String(RS.Kdo("get", "www"))
	// var total1 float64

	Rkwithdrew2()
}

func Rkwithdrew2() {
	tt := time.Now().Format("20060102")
	uid := int64(3505)
	keys := []string{}
	key := RedisHashGoodsOnce(uid, 0)
	// key2 := RedisHashGoodsOnce(8811, 0)
	keys = append(keys, key, "day_withdraw_num_"+tt+"_com.money.calendarweather.lite")
	// keys = append(keys, key, "idtf_gods_5fb8489e-7fc2-11ea-899f-dca90496ac47")
	keys = append(keys, key, "idtf_gods_5fb8489e-7fc2-11ea-899f-dca90496ac471")
	keys = append(keys, key, "bank_acnt_com.money.calendarweather.lite_origin_wxpay_o1OQt51C0dMUKyY8vouUrGKKr7NM")
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
		_, erre := RS.Kdo("del", v)
		println(erre)
	}
	//
	// db := DB.Table("orders_202011")
	// db.First(&Msgs{}, "uid=?", uid)
	// // DB.Table("orders_202011").Delete(Msgs{}, "uid=?", uid)
}

type Msgs struct {
	SiteId    int64  `gorm:"column:site_id" json:"site_id"`
	CoinID    int64  `gorm:"column:coin_id" json:"coin_id"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	DealID    int64  `gorm:"column:deal_id" json:"deal_id"`
	ID        int64  `gorm:"column:id;primary_key" json:"id;primary_key"`
	OrderNo   string `gorm:"column:order_no" json:"order_no"`
	Remark    string `gorm:"column:remark" json:"remark"`
	OrderId   int64  `gorm:"column:order_id" json:"order_id"`
	ToID      int64  `gorm:"column:to_id" json:"to_id"`
	FromID    int64  `gorm:"column:from_id" json:"from_id"`
	Type      int64  `gorm:"column:type" json:"type"`
	MsgType   int64  `gorm:"column:msg_type" json:"msg_type"`
	IsRead    int64  `gorm:"column:is_read" json:"is_read"`
	Username  string `gorm:"column:username" json:"username"`
	Msg       string `gorm:"column:msg" json:"msg"`
}

// 获取上级承兑商 账户
func GetSuperiorUsers(id int64) []int64 {
	var data model.MemberLevel
	DB.Model(data.TableName()).Where("member_id=?", id).First(&data)
	// 拆分上两级用户 id 取对应账户信息
	users := strings.Split(data.Records, "/")
	l := len(users)
	user_ids := make([]int64, 0)
	fmt.Printf("%v", users)
	if l > 2 {
		u1, _ := strconv.ParseInt(users[len(users)-2], 10, 64)
		u2, _ := strconv.ParseInt(users[len(users)-3], 10, 64)
		user_ids = append(user_ids, u2)
		user_ids = append(user_ids, u1)
	} else if l > 1 {
		u1, _ := strconv.ParseInt(users[len(users)-2], 10, 64)
		user_ids = append(user_ids, u1)
	}

	return user_ids
}
func txx(err error) {
	updateColumns := make(map[string]interface{})
	updateColumns["version"] = 10
	updateColumns["update_at"] = time.Now()
	updateColumns["balance"] = gorm.Expr("balance+?", 10)
	updateColumns["remain_amount"] = gorm.Expr("remain_amount+?", 10)
	tx := DB.Begin()
	aa := tx.Table("member_account").
		Where("version=?", 10).
		Where("id=?", 28).
		UpdateColumns(updateColumns).RowsAffected

	log.Info(err, aa)

	tx.Commit()
	log.Info(err, aa)
}

func othersql() {
	// DB.Raw("SELECT FLOOR( MAX(id) * RAND()) FROM `member`").Scan(&aa)
	// DB.Raw("SELECT count(*) FROM `member`").Scan(&aa)

	// DB.Exec("SELECT count(*) as age FROM `member` ").Scan(&aas)

	// DBB.QueryRow("SELECT CURRENT_DATABASE()").Scan(&aa)

	// DB.Where("member_account.status=1").Joins("JOIN member ON member_account.member_id = member.id AND member.switch_order=1 AND  member.type=?", 3).Take(&list)

	// var result Result
	// DB.Raw("SELECT username, id FROM member WHERE username = ?", 3).Scan(&result)
	//
	// if is_exist > 0 {
	// 	println(123, is_exist)
	// 	return
	// }
}

func asdasdas() *model.Member {
	loc, _ := time.LoadLocation("UTC")
	fmt.Printf("%+v", time.Now().In(loc))
	var list model.Member

	DB.Table("member").
		// Where("switch_order=1").
		Where("id=11").
		// Order("RAND()*max(20)").
		Take(&list)
	return &list
}

// Scan
type Result struct {
	Name string
	Age  int
}

func asdas() {
	dd := new(model.MemberProfiles)
	// err := DB.Table("member_profiles").
	// 	// Where("id=?", 1).
	// 	Where("def=?", 0).
	// 	Where("user_id=?", 1).
	// 	Take(&dd)

	DB.Table("member_profiles").Take(dd, "id=?", 10)
	dd.Remark = "122222222223123"
	dd.BackImage = "123122222223"
	dd.RealName = "22223"
	err := DB.Table("member_profiles").
		Save(dd).Error

	log.Info(dd, err)
}

func createdata() {
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
	NewRedis2()
}

func ConnMysql() {
	// dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	// os.Chdir(dir)
	// flag.Parse()
	// cfg := "root:root@tcp(127.0.0.1:3306)/newu?timeout=1s&readTimeout=3s&writeTimeout=3s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
	// cfg := "cootek:cootek@tcp(121.52.250.37:3306)/user_center?timeout=1s&readTimeout=3s&writeTimeout=3s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
	cfg := "cootek:cootek@tcp(121.52.250.37:3306)/withdrew?timeout=1s&readTimeout=3s&writeTimeout=3s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
	// var err error
	// DB, err = gorm.Open("mysql", cfg)
	// DBB = DB.CommonDB()
	//
	// if err != nil {
	// 	panic(err.Error())
	// }
	// DB.LogMode(true)
	dbHandle, _ = NewMysql(cfg)
}

var dbHandle *gorm.DB

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
func NewMysql(key string) (*gorm.DB, error) {
	if dbHandle != nil {
		return dbHandle, nil
	}
	dbHandle, err := gorm.Open("mysql", key)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	dbHandle.DB().SetMaxOpenConns(10)
	dbHandle.DB().SetMaxIdleConns(10)
	dbHandle.DB().SetConnMaxLifetime(1 * time.Minute)
	dbHandle.LogMode(true)
	err = dbHandle.DB().Ping()
	return dbHandle, nil
}

func NewRedis2() (r *redis.Redis, cf func(), err error) {
	// if err = paladin.Get("redis.toml").Unmarshal(&ct); err != nil {
	// 	return
	// }
	// if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
	// 	return
	// }

	cfg := redis.Config{
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
	// err = ping(RS)
	return RS, cf, err
}
