package main

import (
	"crypto/hmac"
	"crypto/md5"
	b64 "encoding/base64"
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
	"regexp"
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
	text := "fff${LastDateOfMonth(3)}ffff aa2021年02月30日aaa${LastDateOfMonth(123)}aaa     "
	mach := "\\$\\{LastDateOfMonth.([0-9]+.)\\}"
	re, _ := regexp.Compile(mach)

	// 取出所有符合规则日期
	list := re.FindAllString(text, -1)
	re1, _ := regexp.Compile("[0-9]+")
	log.Info("替换前：", text, "\n")

	// 遍历替换不同日期
	for _, v := range list {
		dayString := re1.Find([]byte(v))
		days, _ := strconv.Atoi(string(dayString))
		// 获取目标日期
		targetDate := LastDateOfMonth(days, time.Now())
		// 整合当前替换规则
		curDate := "\\$\\{LastDateOfMonth.(" + string(dayString) + ".)\\}"
		// 生成当前替换规则
		re1, _ := regexp.Compile(curDate)
		// 执行替换
		text = re1.ReplaceAllString(text, targetDate)
	}

	log.Info("替换后：", text, "\n")

	<-ch
}

// param: days 为多少天以后
// return: 今天+days 天之后的日期,所在月的最后一天, 按"2006年01月02日"格式化
func LastDateOfMonth(days int, ct time.Time) string {
	d := ct.AddDate(0, 0, days)              // time.Now()可以换成支持测试环境调时间的方法
	firstDate := d.AddDate(0, 0, -d.Day()+1) // 当月的第一天
	lastDate := firstDate.AddDate(0, 1, -1)  // 当月的最后一天
	return lastDate.Format("2006年01月02日")
}

// genHMACmd5 generates a hash signature
func genHMACMD5(ciphertext, key []byte) []byte {
	mac := hmac.New(md5.New, key)
	mac.Write([]byte(ciphertext))
	hmac := mac.Sum(nil)
	return hmac
}

func GetAccessKey(publicKey, secret string) string {
	hmap := genHMACMD5([]byte(publicKey), []byte(secret))
	stringHmac := b64.StdEncoding.EncodeToString(hmap)
	return stringHmac
}
func encodeFullUrl(url string) string {
	// url = urlencode(url)
	// url = str_replace("%2F", "/", url)
	// url = str_replace("%3A", ":", url)
	return url
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
