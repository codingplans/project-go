package main

import (
	"encoding/json"
	"fmt"
	"git.digittraders.com/exchange/pkg/lib"
	"github.com/jordan-wright/email"
	"github.com/prometheus/common/log"
	"math"
	"math/rand"
	"net/http"
	"net/smtp"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

var DATA *PayWay

type PayWay struct {
	//    支付id
	Id  int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Ids int64 `protobuf:"varint,2,opt,name=id,proto3" json:"ids,omitempty"`
	// 支付名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}
type ll struct {
	List []*PayWay
}

func main() {

	ch := make(chan int, 0)

	currentTime := time.Now()
	currentTime1 := time.Now().AddDate(0, 0, 10)
	aa := currentTime1.Sub(currentTime).Hours()
	lastTime, err := time.ParseInLocation("20060102", currentTime.AddDate(0, 0, -1).Format("20060102"), time.Local)
	log.Infof("%v  %v  %v", aa, lastTime, err)

	path := strings.Split("2020-11-20T15:00+08:00", "+")[0]

	timestamp := currentTime.Add(-time.Hour).Hour()

	rand.Seed(time.Now().UnixNano())
	fmt.Println(timestamp, rand.Int31n(20), "\n", 10/7)

	rate := "{\"id\":1000,\"ids\":100}"
	ww := &PayWay{}
	json.Unmarshal([]byte(rate), &ww)

	log.Info(fmt.Sprintf("%.4f,%.4f \n", 2123.12312, 111.21212), time.Now().Unix(), path, "\n", ww, timestamp)

	<-ch
}

func MsToTime(ms string) (time.Time, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	tm := time.Unix(0, msInt*int64(time.Millisecond))
	return tm, nil
}

func Post(path string, param map[string]interface{}) (content []byte, err error) {

	bs, _ := json.Marshal(param)
	body := strings.NewReader(string(bs))
	req, err := http.NewRequest("POST", path, body)
	// 此处还可以写req.Header.Set("User-Agent", "myClient")
	req.Header.Add("User-Agent", "myClient")

	clt := http.Client{}
	resp, err := clt.Do(req)

	log.Info(resp)
	return
}

func GenRedEnvelopeRain(prize_keys []string, coin_amount int64) (ret []Rain) {
	var do_limit bool
	if coin_amount >= 2500 {
		do_limit = true
	}
	amounts := []int64{}
	rand.Seed(time.Now().Unix())

	for _, v := range prize_keys {
		if v == "coin_red" {
			if !do_limit {
				println(333)
				for i := 0; i < 55; i++ {
					amounts = append(amounts, rand.Int63n(9))
				}
			} else {
				amounts = []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
				temp := make([]int64, 40, 40)
				amounts = append(amounts, temp...)
			}
		} else if v == "coin_golden" && !do_limit {
			rang := []int64{18, 28, 58}
			amounts = append(amounts, rang[rand.Intn(2)])
		}
		ret = append(ret, Rain{
			PrizeKey:     v,
			PrizeType:    "coin",
			PrizeAmounts: amounts,
		})

	}
	return ret
}

type Rain struct {
	PrizeAmounts []int64 `json:"prize_amounts"`
	PrizeKey     string  `json:"prize_key"`
	PrizeType    string  `json:"prize_type"`
}

// 判断时间是当年的第几周
func WeekByDate(t time.Time) string {
	yearDay := t.YearDay()
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
	firstDayInWeek := int(yearFirstDay.Weekday())

	// 今年第一周有几天
	firstWeekDays := 1
	if firstDayInWeek != 0 {
		firstWeekDays = 7 - firstDayInWeek + 1
	}
	var week int
	if yearDay <= firstWeekDays {
		week = 1
	} else {
		week = (yearDay-firstWeekDays)/7 + 2
	}
	return fmt.Sprintf("%d%d   %d", t.Year(), week, 123)
}
func closss(ch1 chan int, ch2 chan int) {
	// time.Sleep(10 * time.Second)
	// close(ch1)
	// close(ch2)
}

type T int

func IsClosed(ch <-chan int) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}
func chanaa(ch1 chan int, ch2 chan int) {
	for {

		ch1 <- 22
		// ch1 <- 21
		// if _, ok := <-ch2; ok {
		ch2 <- 44
		ch2 <- 31
		ch2 <- 34
		ch2 <- 32
		ch2 <- 33

		// _, ok := <-ch1

		// println(IsClosed(ch1))
		// ch1 <- 22
		// // if _, ok := <-ch2; ok {
		// ch2 <- 44
		// }
		println(999)
		// time.Sleep(3 * time.Second)
	}

}

func ExampleGmail() {
	println(222)
	e := email.NewEmail()
	e.From = "zzyphp@gmail.com"
	e.To = []string{"darrenzzy@126.com"}
	// e.Bcc = []string{"darrenzzy@126.com"}
	// e.Cc = []string{"darrenzzy@126.com"}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Text Body is, of course, supported!\n")
	e.HTML = []byte("<h1>Fancy Html is supported, too!</h1>\n")
	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", e.From, "facaiba123", "smtp.gmail.com"))
	if err != nil {
		log.Info(err.Error())

	}
	println(333333)
}

func cmdd() {
	// cmd := exec.Command("ls", "|grep", "go") // /查看当前目录下文件
	cmd := exec.Command("sh", "-c", "ls ../../../ ")
	// /查看当前目录下文件
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out), 444)
}

func aaaa() (float64, float64, float64) {
	aa := float64(158)
	fee := math.Round(20*float64(7)/1000*1000) / 1000
	level_two_fee := math.Round((aa-aa*float64(1)/1000)*1000) / 1000
	return aa, fee, level_two_fee
}

func authgoogle() {
	fmt.Println("-----------------开启二次认证----------------------")
	// user := "testxx1111@qq.com"
	// secret, code := lib.InitAuth(user)
	secret, code := "YTL5YDXZF5GOOALE5HYN2BH7LYYZOFXL", "981135"
	fmt.Println(secret, 8888, code)

	fmt.Println("-----------------信息校验----------------------")

	// secret最好持久化保存在
	// 验证,动态码(从谷歌验证器获取或者freeotp获取)
	bool, err := lib.NewGoogleAuth().VerifyCode(secret, code)
	if bool {
		fmt.Println("√")
	} else {
		fmt.Println("X", err)
	}
}

func ddddwg() {
	// funcName()
	var wg sync.WaitGroup
	wg.Add(11)
	go dddf()
	go dddf()

	// discov()
	wg.Wait()
}

//
// func discov() {
// 	c := &conf.Config{
// 		Env: &conf.Env{
// 			Region:    "",
// 			Zone:      "sh1",
// 			DeployEnv: "test",
// 			Host:      "test_server",
// 		},
// 		Nodes: []string{"127.0.0.1:7171"},
// 		HTTPServer: &xhttp.ServerConfig{
// 			Addr:    "127.0.0.1:7171",
// 			Timeout: xtime.Duration(time.Second * 1),
// 		},
// 		HTTPClient: &xhttp.ClientConfig{
// 			Timeout:   xtime.Duration(time.Second * 1),
// 			Dial:      xtime.Duration(time.Second),
// 			KeepAlive: xtime.Duration(time.Second * 1),
// 		},
// 	}
// 	_ = c.Fix()
// 	paladin.Init()
// 	dis, _ := discovery.New(c)
// 	println(123)
//
// 	http.Init(c, dis)
// }

func dddf() {
	for i := 1; i < 10; i++ {
		defer println(i)
	}
}
