package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"git.digittraders.com/exchange/pkg/lib"
	"github.com/google/uuid"
	"github.com/jordan-wright/email"
	"github.com/prometheus/common/log"
	"html/template"
	"math"
	"math/rand"
	"net/smtp"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testgo/modgo/model"
	"time"
)

var DATA *PayWay

type PayWay struct {
	//    支付id
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// 支付名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

type ll struct {
	List []*PayWay
}

func main() {

	DATA = &PayWay{
		Id:   222,
		Name: "222",
	}

	ch := make(chan int, 0)

	// var ss sync.Map

	// ss.LoadOrStore("qqq", "1")
	// ss.LoadOrStore("www", "2")
	// ss.LoadOrStore("ddd", "3")
	// aa, err := ss.Load("qqq")
	// log.Info(aa, err)
	// ExampleGmail()
	// t := time.Now().Format("2006-01-02")
	// ts := t.Unix() - int64(t.Minute()) - int64(t.Second()) - int64(t.Sub())
	// timeStr := time.Now().Format("2006-01-02")
	timeStr := time.Now().Format("2006-01")
	fmt.Println("timeStr:", timeStr)
	t, _ := time.Parse("2006-01", timeStr)
	timeNumber := t.Unix()
	fmt.Println("timeNumber:", timeNumber)

	// aa := t.AddDate(0, 0, -1)
	// log.Info(ts)

	<-ch
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

func funcName() {
	obj := Constructor()
	for i := 0; i < 10; i++ {
		obj.AddAtTail(2*i + 1)
	}
	obj.AddAtHead(22)

	obj.AddAtTail(99)

	obj.AddAtIndex(5, 323)
	obj.DeleteAtIndex(1)
	list := obj.Header
	log.Info(list.Val)
	for list.Next != nil {
		list = list.Next
		log.Info(list.Val)
	}

	list = obj.Tail
	log.Info(list.Val, "tail")
	for list.Next != nil {
		list = list.Next
		log.Info(list.Val, "tail")
	}
}

// 定义一个 链表结构
type ListNode struct {
	Val  int       `json:"val"`
	Next *ListNode `json:"next"`
}

// 为了方便计算,我们可以定义一个header,存header节点
// 定义一个tail 存 尾节点 (不需要遍历到结尾...)
// 定义一个长度,记录链表长度 (不需要每次遍历计算)
type MyLinkedList struct {
	Header *ListNode `json:"header"`
	Tail   *ListNode `json:"tail"`
	Lens   int       `json:"lens"`
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		Header: nil,
		Tail:   nil,
		Lens:   0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	// 如果获取的位置小于0或者等于链表长度,直接返回-1(注意链表下标从0开始,所以这地方可以等于)
	if index < 0 || index >= this.Lens {
		return -1
	}

	// 如果index等于0,直接返回头节点的值
	if index == 0 {
		return this.Header.Val
	}

	// 遍历一下,找到index节点的值
	node := this.Header
	for node.Next != nil {
		// 因为0的情况一排除,所以直接先减掉
		index--
		// node指针往下移动一位
		if node.Next != nil {
			node = node.Next
		}
		// 当index递减等于0的时候, 返回其值就可以了
		if index == 0 {
			return node.Val
		}
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	// 在头节点加入一个节点,那么这个节点就是以后的头节点了.. 而且这个节点的next指向以前的头节点...
	this.Header = &ListNode{
		Val:  val,
		Next: this.Header,
	}

	// 如果当前链表为空,那么增加一个节点,这个节点既是头节点又是尾节点
	if this.Lens == 0 {
		this.Tail = this.Header
	}
	// 因为增加了节点,所以链表长度+1
	this.Lens++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	// 如果当前链表为空,那么增加尾部,也就是加个头部..
	if this.Lens == 0 {
		this.Tail = &ListNode{
			Val:  val,
			Next: nil,
		}
		this.Header = this.Tail
		this.Lens++
		return
	}
	// 尾节点本来next等于nil,现在加一个,next等于这个节点
	this.Tail.Next = &ListNode{
		Val:  val,
		Next: nil,
	}

	// 所以以后新的尾节点就是之前的next节点了..
	this.Tail = this.Tail.Next

	// 新增节点,链表长度+1
	this.Lens++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {

	//   如果 index小于0，则在头部插入节点。
	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	//   如果 index 大于链表长度，则不会插入节点。
	if index > this.Lens {
		return
	}

	//   如果 index 等于链表的长度，则该节点将附加到链表的末尾。
	if index == this.Lens {
		this.AddAtTail(val)
		return
	}

	node := this.Header
	for node.Next != nil {
		index--
		// 当index == 0的时候,说明找到了这个节点,往这节点之前插入节点
		if index == 0 {
			newNode := &ListNode{
				Val:  val,
				Next: node.Next,
			}
			node.Next = newNode
			// 记得长度+1
			this.Lens++
			// 记得要返回..
			return
		}

		node = node.Next
	}

}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	// 如果index小于0或者大于等于长度,直接返回
	if index < 0 || index >= this.Lens {
		return
	}

	// 如果等于0,就是删除头节点,记得链表长度-1
	if index == 0 && this.Header.Next != nil {
		this.Header = this.Header.Next
		this.Lens--
	}

	node := this.Header
	for node.Next != nil {
		index--

		if index == 0 {
			// 如果node.Next.Next == nil 说明到最后一个节点了.相当于删除最后一个节点
			if node.Next.Next == nil {
				node.Next = nil
				this.Tail = node
				this.Lens--
				return
			}
			// 其他情况就是删除中间一个节点(A->B->C),操作就是  A 直接指向 C 就行 (A->C)
			node2 := node.Next.Next
			node.Next = node2
			this.Lens--
			return
		}
		node = node.Next
	}

}

// bb := twoSum([]int{0, 2522, 75}, 100)
func twoSum(numbers []int, target int) []int {

	l := len(numbers)
	var i, j int = 0, l - 1
	for i < j {

		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}

		if numbers[i]+numbers[j] < target {
			i++
		}
		if numbers[i]+numbers[j] > target {
			j--
		}
	}
	return []int{}
}

func mergeSoft() {
	// 	归并排序
	arr1 := []int{1, 3, 5, 7, 9}
	arr2 := []int{2, 4, 5, 6, 8, 99}

	arr3 := []int{}
	l := len(arr1)
	l2 := len(arr2)
	var j1, j2 int
	for j1 < l && j2 < l2 {
		if arr1[j1] < arr2[j2] {
			arr3 = append(arr3, arr1[j1])
			j1++
		} else {
			arr3 = append(arr3, arr2[j2])
			j2++
		}

	}

	for j1 < l {
		arr3 = append(arr3, arr1[j1])
		j1++
	}
	for j2 < l2 {
		arr3 = append(arr3, arr2[j2])
		j2++
	}
	log.Info(arr3)

}

// reverseString([]string{"A", " ", "m", "a", "n", ",", " ", "a", " ", "p", "l", "a", "n", ",", " ", "a", " ", "c", "a", "n", "a", "l", ":", " ", "P", "a", "n", "a", "m", "a"})
func reverseString(arr []string) {

	l := len(arr) - 1
	if l < 0 {
		return
	}
	for i := 0; i <= l/4; i++ {
		arr[i], arr[l-i] = arr[l-i], arr[i]
		arr[l/2+i], arr[l/2-i] = arr[l/2-i], arr[l/2+i]
		println(arr[l/2+i], arr[l/2-i], 00000)
		println(arr[i], arr[l-i])
	}
	log.Info(arr)
	log.Info([]string{"a", "m", "a", "n", "a", "P", " ", ":", "l", "a", "n", "a", "c", " ", "a", " ", ",", "n", "a", "l", "p", " ", "a", " ", ",", "n", "a", "m", " ", "A"})
	log.Info([]string{"A", " ", "m", "a", "n", ",", " ", "a", " ", "p", "l", "a", "n", ",", " ", "a", " ", "c", "a", "n", "a", "l", ":", " ", "P", "a", "n", "a", "m", "a"})

}

// s1 = "abaaacbacaabb"
//	s2 = "bcccaababcacc"
//	s3 = "bcccabaaaaabcbacaababcacbc"
func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if (n + m) != t {
		return false
	}
	f := make([]bool, m+1)
	f[0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			log.Info(i, j, i+j, f)
			p := i + j - 1
			if i > 0 {
				f[j] = f[j] && s1[i-1] == s3[p]
			}
			if j > 0 {
				f[j] = f[j] || f[j-1] && s2[j-1] == s3[p]
			}
		}
	}
	return f[m]
}

func isInterleave2(s1 string, s2 string, s3 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	if l1 == 0 || l2 == 0 {
		return s3 == s1+s2
	}
	return dfs(s1, s2, s3)
}

// dfs + lru  深度优先+ 淘汰策略
func dfs(s1, s2, s3 string) bool {
	max++
	println(s1, s2, s3, max)
	flag1 := (s1[0] == s3[0] && isInterleave(s1[1:], s2, s3[1:]))
	flag2 := (s2[0] == s3[0] && isInterleave(s1, s2[1:], s3[1:]))
	return flag1 || flag2
}

// aa := searchInsert([]int{1, 4, 6, 7, 8, 10}, 6)
func searchInsert(nums []int, target int) int {
	var l, h, mid int
	h = len(nums)
	for l < h {
		mid = (h-l)/2 + l
		println(l, mid, h)
		if target < nums[mid] {
			h = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			println(l, mid, h, "****")

			return mid
		}
	}
	println(l, mid, h, "$$$$$")

	return mid
	// l := len(nums)
	// for k, v := range nums {
	// 	if v >= target {
	// 		return k
	// 	}
	// 	if k+1 < l && nums[k+1] > target {
	// 		return k + 1
	// 	}
	// }
	// return l
}

var arr []string
var sss string
var max int

func longestPalindrome(s string) string {
	l := len(s)
	for i := 0; i < l; i++ {

		longestPalindrome2(s, i, i)
		longestPalindrome2(s, i, i+1)
	}
	return sss
}

func longestPalindrome2(s string, left, right int) {
	l := len(s)
	for left >= 0 && right < l && s[left] == s[right] {
		temp := right - left + 1
		if temp > max && left != right {
			max = temp
			sss = s[left : right+1]
		}
		left--
		right++
	}
}

func longestPalindrome11(s string) string {
	lenth := len(s)

	if lenth <= 1 {
		return s
	}

	dp := make([][]bool, lenth)

	start := 0
	maxlen := 1

	for r := 0; r < lenth; r++ {
		dp[r] = make([]bool, lenth)
		dp[r][r] = true
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
			} else {
				dp[l][r] = false
			}

			if dp[l][r] {
				curlen := r - l + 1
				if curlen > maxlen {
					maxlen = curlen
					start = l
				}
			}
		}
	}
	return s[start : start+maxlen]
}

// 回文串
func palindrome(ss string) bool {
	println(ss)
	ss = strings.ToLower(ss)
	l := len(ss)
	var n, m int = 0, 0
	m = l - 1
	for n <= l/2 && m > 0 {
		if !rands(ss[n]) {
			n++
			continue
		}
		if !rands(ss[m]) {
			m--
			continue
		}

		if ss[n] != ss[m] {
			return false
		}
		n++
		m--

	}
	return true
}

func rands(ss uint8) bool {
	switch {
	case ss <= 90 && ss >= 65:
		return true
	case ss <= 57 && ss >= 48:
		return true
	case ss <= 122 && ss >= 97:
		return true
	default:
		return false
	}
}

func wiatgo() {
	cc := make(chan int, 2)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 1; i < 10; i++ {
			wg.Add(1)
			cc <- i
			log.Info("111111")
		}
	}()

	go func() {
		for {
			aa, ok := <-cc
			wg.Done()
			log.Info(aa, ok)
		}

	}()

	wg.Wait()
}

func jsonmap() {
	body := new(PayWay)
	body.Id = 1
	body.Name = "12"

	bs, err := json.Marshal(body)
	log.Info(string(bs), err)

	ss := &PayWay{}
	json.Unmarshal(bs, ss)
	log.Infof("%+v", ss)
}

func amountverify() error {
	ss1 := "^(0|[1-9]\\d{0,7})(\\.\\d{1,8})?$"
	aa, _ := regexp.MatchString(ss1, "99999886.20888966")
	re := regexp.MustCompile("^(0|[1-9]+)(\\.\\d+)?$")
	if re.MatchString("1e20") {
		println(123, aa)
	}
	println(123, aa)
	return nil
}

func printss() {
	var aa uint32
	aa = 3342247978
	aa = aa / 1000

	fmt.Printf("%+v,%+v,%+v", aa, int64(aa), 11)
}

func delmap() {
	aa := new(PayWay)
	aa.Id = 11
	aa.Name = "12313"

	DATA = aa

	aa.Id = 12

	ss := make(map[int]int, 0)

	ss[1] = 22
	ss[11] = 22
	ss[3] = 22

	fmt.Printf("%v", ss)

	delete(ss, 13122323)
	fmt.Printf("%v", ss)
	delete(ss, 13)
	fmt.Printf("%v", ss)

	println(aa.Id)
}

func randst() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	println(rnd.Int31n(1000000))
}

func semail() {
	// 定义收件人
	mailTo := []string{
		"darrenzzy@126.com",
	}

	// to := "darrenzzy@126.com"
	// 邮件主题为"Hello"
	subject := "Hello by golang gomail from exmail.qq.com"
	// 邮件正文
	// body := "Hello,by gomail sent"

	err := SendMail1(mailTo, subject)
	if err != nil {
		log.Info(err.Error())
		fmt.Println("send fail")
		return
	}

	fmt.Println("send successfully")
}
func SendMail1(toUsers []string, subject string) error {
	// NewEmail返回一个email结构体的指针
	e := email.NewEmail()
	// 发件人
	e.From = "1376161485@qq.com"
	// 收件人(可以有多个)
	e.To = toUsers
	// 邮件主题
	e.Subject = subject

	// 解析html模板
	t, err := template.ParseFiles("email-template.html")
	if err != nil {
		return err
	}
	// Buffer是一个实现了读写方法的可变大小的字节缓冲
	body := new(bytes.Buffer)
	// Execute方法将解析好的模板应用到匿名结构体上，并将输出写入body中
	t.Execute(body, struct {
		FromUserName string
		ToUserName   string
		TimeDate     string
		Message      string
	}{
		FromUserName: "go语言",
		ToUserName:   "Sixah",
		TimeDate:     time.Now().Format("2006/01/02"),
		Message:      "golang是世界上最好的语言！",
	})
	// html形式的消息
	e.HTML = body.Bytes()
	// 从缓冲中将内容作为附件到邮件中
	e.Attach(body, "email-template.html", "text/html")
	// 以路径将文件作为附件添加到邮件中
	// e.AttachFile("~/go/src/email/main.go")
	// 发送邮件(如果使用QQ邮箱发送邮件的话，passwd不是邮箱密码而是授权码)
	return e.Send("smtp.qq.com:587", smtp.PlainAuth("", e.From, "tunsqnkhlgmmgbbc", "smtp.qq.com"))
}

func ccc() {
	aa.cc = "121"
	aa.bb = []string{"222222"}
}

func mapsliceadd() {
	SS := make(map[int64][]*model.Member, 0)

	SS[2] = append(SS[2], &model.Member{ID: 12})
	SS[2] = append(SS[2], &model.Member{ID: 122})
	SS[1] = append(SS[2], &model.Member{ID: 1223})
	log.Info(SS)
}

func floatmm() {
	aa := float64(5000)

	bb := float64(7)
	cc := float64(6.8)

	sss := lib.FloatRound(aa/cc, 4)

	log.Info(aa/bb, aa/cc, sss)
}

func uuidss() {
	aa := fmt.Sprintf("%010v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(100000000000))

	bb := uuid.New().ID()
	log.Infof("%+v  %v,%v", aa, bb, int64(bb))
}

// 币种
type Fiats struct {
	//    名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//    币id
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

// 合并接口 数据包
type SyncResp struct {
	//    交易方式
	PayWay []*PayWay `protobuf:"bytes,1,rep,name=pay_way,json=payWay,proto3" json:"pay_way,omitempty"`
	//    当前交易所流通货币
	Symbol []string `protobuf:"bytes,2,rep,name=symbol,proto3" json:"symbol,omitempty"`
	// 法币币种
	Fiats []*Fiats `protobuf:"bytes,3,rep,name=fiats,proto3" json:"fiats,omitempty"`
}

type RRR struct {
}

type RespData struct {
	Code    int64  `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func switchs() {
	aa := 11
	switch aa {
	case 22:
		println(aa, 2)
	case 3:
		println(aa, 3)
	case 1:
		println(aa, 11)
		aa++
	case 2:
		println(aa, 2)
	default:
		println(aa, 111)
		return

	}
	println(aa)
}

func floattostr() {
	var ff float64

	ff = 3198000.2200020

	aa := strconv.FormatFloat(ff, 'f', -1, 64)
	log.Info(aa)
}

func gogo() {
	var ss string
	// if ss, ok = qqs(); ok {
	// 	println(123)
	// }

	var tt time.Time
	tt = time.Now()
	ch := make(chan int, 10)
	exits := make(chan int, 2)

	go qqs2(ch)
	go qqs(ch, exits)
	// go qqs3(ch)
	fmt.Printf("%+v  %s  %d", ss, tt, tt.Unix())

	for i := 0; i < 5; i++ {
		ch <- i
	}

	time.Sleep(2 * time.Second)

	close(exits)
	if dd, ok := <-exits; ok {
		println(123, dd)
	} else {
		println(333)

	}

	time.Sleep(20 * time.Second)
}

func qqs(c, exits chan int) (string, bool) {
	select {
	case aa, ok := <-c:
		if ok {
			println(aa, 111)
		}
	case <-exits:
		println(1111, 111)

	}
	println("over")

	return "12", true
}

func qqs2(c chan int) (string, bool) {
	for {
		select {
		case aa, ok := <-c:
			if ok {
				println(aa, 222)
			}
		}
	}
}

func qqs3(c chan int) (string, bool) {
	for {
		select {
		case aa, ok := <-c:
			if ok {
				println(aa, 333)
			}
		}
	}
}

func maptest() {
	mm := make(map[string]int64, 0)
	mm["aaaaaaaaaaaaaaaaaaa"] = 123
	mm["zzzzzzz"] = 123
	mm["cccccccccccccccccccccc"] = 123
	mm["bb"] = 123
	mm["kkkkkkkkkkkkkkkk"] = 123
	mm["rr"] = 123
	mm["zzzz"] = 123
	for k, v := range mm {
		println(k, v)
	}

	fmt.Printf("%+v", mm)
}

type ttt struct {
	Amount  string
	Timeout string
}

type ttt2 struct {
	Amount  string
	Timeout string
}

func floattest() {
	s := "123.129asdas"

	// var fff float64
	//
	// fff = 55.43453
	//
	// println(strconv.FormatFloat(fff, 'g', -1, 64))

	aa, err := strconv.ParseFloat(s, 8)
	if err == nil {
		fmt.Printf("%T,%v", aa, aa)
	} else {
		fmt.Printf("%T,%v ****", aa, aa)

		println(err.Error(), aa)

	}

	var ae uint8
	ae = 'a'

	aas := float64(123123.22)
	ww := strconv.FormatFloat(aas, 'g', -1, 64)

	ww = fmt.Sprintf("%d", 100) + "12"
	println(ww, ae)
	// fmt.Println(authorize("123", "123", "post", 123))
}

var aa struct {
	cc string
	bb []string
}
var staticMachine = getMachineHash()
var staticIncrement int64
var staticPid = int32(os.Getpid())

func getMachineHash() int32 {
	machineName, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	buf := md5.Sum([]byte(machineName))
	return (int32(buf[0])<<0x10 + int32(buf[1])<<8) + int32(buf[2])
}
func GenerateID() (id int64) {
	timeStr := time.Now().Format("0102150405")
	ss := timeStr + fmt.Sprint(staticMachine) +
		fmt.Sprint(staticPid) + fmt.Sprint(atomic.AddInt64(&staticIncrement, 1))
	println(ss)
	id, _ = strconv.ParseInt(ss, 10, 64)
	return
}
