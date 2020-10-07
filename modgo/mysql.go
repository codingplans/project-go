package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/container/pool"
	xtime "github.com/go-kratos/kratos/pkg/time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
	"github.com/shopspring/decimal"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testgo/modgo/model"
	"time"
)

var DB *gorm.DB
var DBB gorm.SQLCommon

func main() {

	// GetRecommend(5)
	// RS.Kdo("set", "www", 123)
	// aa, _ := redis.String(RS.Kdo("get", "www"))
	// var total1 float64

	type res struct {
		Total1 decimal.Decimal `gorm:"column:total1"`
	}
	var Res1, Res2, Res3 res

	today := time.Now().Format("2006-01-02")
	month := time.Now().Format("2006-01")
	year := time.Now().Format("2006")

	query := DB.Table("member_transaction").
		Select("sum(amount) as total1").
		Where("member_id=?", 5).
		Where("amount>0").
		Where("type=?", 8)

	query.Where("create_time>?", today).
		Scan(&Res1)

	query.Where("create_time>?", month).
		Scan(&Res2)

	query.Where("create_time>?", year).
		Scan(&Res3)
	log.Info(Res1, Res2, Res3)
	return
	userId := 10172
	var data model.MemberLevel
	DB.Table("member_level").Take(&data, "member_id=?", userId)

	var num2 int64
	var list []*model.MemberLevel
	var like string
	like = fmt.Sprintf("%s%d/", data.Records, userId)
	DB.Table("member_level").
		Where("member_level.before_id!=?", userId).
		Where("member_level.records like ?", like+"%").
		Joins("JOIN member_level as b ON b.member_id = member_level.before_id AND b.before_id != ?", userId).
		Find(&list)

	// 遍历出 2 级用户数量
	for _, v := range list {
		arr := strings.Split(v.Records, like)
		ss := arr[len(arr)-1]
		repet := strings.Split(ss, "/")
		log.Info(ss, "****", v.ID, "&&&&7", repet, "   ", len(repet), num2)
		// 不是 2 级则跳过
		if len(repet) != 3 {
			continue
		}
		num2++
	}
	log.Info(num2)
	// RS
}

func GetRecommend(userId int64) {
	var data model.MemberLevel
	// data := d.GetMemberLevel(userId)
	DB.Table("member_transaction").Take(&data, "member_id=?", userId)
	// var num1, num2 int64
	// 找自己下面的一级人
	// DB.Table("member_transaction").Where("before_id=?", userId).Count(&num1)

	var ss []*model.MemberTransaction
	// like := fmt.Sprintf("%s%d/%s", data.Records, data.MemberID, "%")
	DB.Table("member_transaction").
		// Select("id,amount as records").
		// Where("before_id!=?", userId).
		// Where("records like ?", like).
		// Group("member_id").
		Find(&ss)

	bs, _ := json.Marshal(ss)
	RS.Kdo("set", "www", string(bs))
	ll, _ := redis.Bytes(RS.Kdo("get", "www"))

	json.Unmarshal(ll, &ss)
	for _, v := range ss {
		log.Info(v)
	}
	// log.Info(num1, num2, ll)
	// data.Records

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
	NewRedis1()

}

func ConnMysql() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	os.Chdir(dir)
	flag.Parse()
	cfg := "dbuser:pass!23word@tcp(192.168.3.8:3306)/user_center?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
	// cfg := "root:root@tcp(127.0.0.1:23306)/user_center?timeout=1s&readTimeout=3s&writeTimeout=3s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
	var err error
	DB, err = gorm.Open("mysql", cfg)
	DBB = DB.CommonDB()
	if err != nil {
		panic(err.Error())
	}
	DB.LogMode(true)
}

// SqlStore mysql读/写库客户端封装
type SqlStore struct {
	*gorm.DB
}

func NewRedis1() (r *redis.Redis, cf func(), err error) {
	// if err = paladin.Get("redis.toml").Unmarshal(&ct); err != nil {
	// 	return
	// }
	// if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
	// 	return
	// }

	cfg = redis.Config{
		Addr:         "127.0.0.1:6379",
		DialTimeout:  xtime.Duration(90 * time.Second),
		ReadTimeout:  xtime.Duration(90 * time.Second),
		WriteTimeout: xtime.Duration(90 * time.Second),
		SlowLog:      xtime.Duration(90 * time.Second),
		Name:         "user_center",
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
