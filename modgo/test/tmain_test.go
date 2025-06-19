package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/md5"
	crand "crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"net/http/httptest"
	urls "net/url"
	"os"
	"os/signal"
	"reflect"
	"regexp"
	"runtime"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"testing"
	"time"
	"unicode/utf8"
	"unsafe"

	"github.com/Darrenzzy/person-go/structures"
	"github.com/jinzhu/copier"
	jsoniter "github.com/json-iterator/go"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	// _ "go.uber.org/automaxprocs"
	"go.uber.org/goleak"
	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"
	"testgo/modgo/crypto"
)

// 给定一个正整数数组a，以及一个目标数字x，使用你最熟悉的语言，编写一个函数，找出数组中所有可以使数字之和等于x的组合g（组合不能重复）。注：数组中的数字可以无限制被重复选取。
func TestSumToG(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	target := 7
	result := sumToG(arr, target)
	fmt.Println(result)

}

func sumToG(arr []int, target int) [][]int {
	result := make([][]int, 0)
	dfs3(arr, target, 0, []int{}, &result)
	fmt.Printf("所有组合为: %v\n", result)

	return nil
}

func dfs3(arr []int, target int, start int, path []int, result *[][]int) {
	// 循环从start开始，到最后，一层层记录path，并且递归，减少target
	pathItem := make([]int, len(path))
	copy(pathItem, path)
	if target == 0 {
		*result = append(*result, pathItem)
	}

	for i := start; i < len(arr); i++ {
		if arr[start] > target {
			return
		}
		pathItem = append(pathItem, arr[i])
		dfs3(arr, target-arr[i], i, pathItem, result)
		pathItem = pathItem[:len(pathItem)-1]
	}
}

func TestForRangex(t *testing.T) {
	arr := make([]*int, 0, 10)
	// arr := make([]*Baz, 0, 10)
	for i := 0; i < 10; i++ {
		arr = append(arr, &i)
		// arr = append(arr, &Baz{
		// 	Bar: &i,
		// })
	}
	// for i := 0; i < 10; i++ {
	// 	print(&i)
	// }
	fmt.Println(*arr[0], *arr[3])

	// for m, n := range arr {
	// 	fmt.Println(m, n)
	// }

}

func TestSumDiff(t *testing.T) {
	res := sumDiff(13112)
	print(res)
}

// 给你一个整数 n，请你帮忙计算并返回该整数「各位数字之积」与「各位数字之和」的差。
func sumDiff(n int) int {
	sum := 0
	product := 1
	for n != 0 {
		xxx := n % 10
		n = n / 10
		sum += xxx
		product *= xxx
	}
	return product - sum
}

// 方案1: 在插入数据库前清理无效的UTF-8字符
func cleanInvalidUTF8(input string) string {
	// 1. 移除常见的十六进制表示
	hexPattern := regexp.MustCompile(`<[0-9A-F]{2}>`)
	cleaned := hexPattern.ReplaceAllString(input, "")

	// 2. 确保是有效的 UTF-8
	cleaned = strings.ToValidUTF8(cleaned, "")

	// // 3. 移除可能的末尾特殊字符
	// cleaned = strings.TrimRight(cleaned, "}")
	// cleaned = cleaned + "}" // 重新添加结束括号

	return cleaned
}

func TestGetConfig(t *testing.T) {
	// 定义日志行
	// line := "[04-14 16:49:04.974089][WARNING ][indicator_adapter][prepare_data        :  245][20073][-] {monitor_indicator_result_alert use config in params , error:no indicator_alert_prepare_data}"
	line := `[11-27 09:05:01.770200] [info    ] [main        ] [main            ] [Version     ] [-] {WCTraderVersion:-,GIT_COMMIT_MSG:\"\",WCCommonVersion:1.7.2,WCCommonUpdate:Add config.h which originally is in Repo of WCTrader)}`
	// line := "[04-14 16:49:05.872243][INFO    ][base_group_log][package_alarm       :  104][20073][-] { rows: {'params': {'config': {'name': 'monitor_indicator_result_alert', 'alias': '指标计算回测对比结果', 'module_path': 'filter_dispatcher/dispatcher/alert/monitor_indicator_result_alert/monitor_indicator_result_alert.py', 'alert_time_range': {'date': '0-4', 'time': [['04:00:00', '04:04:00']]}, 'alert_calculation_period': 300, 'uniq_index': ['alert_name', 'indicator_msg'], 'alarm_group': 'Dev', 'alarm_manager': 'zhangzhenyu', 'alarm_uniq_period': 300, 'alarm_notice_type': ['msg']}}, 'data': {}}}"

	// 定义正则表达式
	re := regexp.MustCompile(`^\[(\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}\.\d{6})\]\s*\[([^\]]+)\]\s*\[([^\]]+)\]\s*\[([^\]]+)\]\s*\[([^\]]+)\]\s*\[([^\]]+)\]\s*(.*)$`)

	// 执行正则匹配
	matches := re.FindStringSubmatch(line)

	// 打印匹配结果
	if len(matches) > 0 {
		for i, match := range matches {
			fmt.Printf("Match %d: %s\n", i, match)
		}
	} else {
		fmt.Println("No matches found.")
	}
}

func TestUtf8Case(t *testing.T) {
	// 使用示例
	aas := "[01-02 14:31:05.185455] [error   ] [AdpBrokerSpi] [OnOrderEvent    ] [Rejected    ] [-] {msg:平仓仓位或配股可用权证不足,可用:0<EF><BC>}"
	msg2 := cleanInvalidUTF8(aas)
	t.Log(aas)
	t.Log(msg2)

}

func TestGetDaysInMonth(t *testing.T) {
	value := "08.9...93673.67"

	t.Log(strings.SplitN(value, ".", 2))
	t.Log(cast.ToString(value))
	arr := []int{1, 2}
	t.Log(arr[:0])
	t.Log(arr[:2])
}
func TestPubPrefix(t *testing.T) {
	pid := os.Getpid()
	fmt.Printf("Current process PID: %d\n", pid)
	strsarr := []string{
		"/xxxxxxxxx/Configxxxxx/",
		"WCTrader.pid.471765.log",
	}

	// prefix := commonPrefix(strsarr)
	// fmt.Println("Common Prefix:", prefix)
	keyword(strsarr)
}

func TestMapStore(t *testing.T) {
	MM := sync.Map{}
	t.Log(MM.LoadOrStore("key", "value"))
	t.Log(MM.LoadOrStore("key2", "value2"))
	t.Log(MM.Load("key2"))
	t.Log(MM.Load("key"))
	t.Log(MM.LoadOrStore("key", "value11"))
	t.Log(MM.LoadOrStore("key2", "value222"))

}

// 保留的SQL关键字列表
var sqlKeywords = map[string]bool{
	"SELECT": true, "FROM": true, "WHERE": true, "INSERT": true, "UPDATE": true, "DELETE": true,
	// 添加其他SQL关键字...
}

// 判断字符串是否为有效的SQL表名称
func isValidTableName(name string) bool {
	// 检查是否为空
	if len(name) == 0 {
		return false
	}
	// 检查是否为SQL关键字
	if _, exists := sqlKeywords[strings.ToUpper(name)]; exists {
		return false
	}
	// 使用正则表达式检查格式
	match, _ := regexp.MatchString(`^[a-zA-Z_][a-zA-Z0-9_]*$`, name)
	return match
}

func TestTableName(t *testing.T) {
	t.Log(isValidTableName("SELECT"))
	t.Log(isValidTableName("aaa-test"))
	t.Log(isValidTableName("aaa_asd3"))
}

// ConvertPattern 将带有*的模式转换为正则表达式
func ConvertPattern(pattern string) string {
	pattern = strings.Replace(pattern, ".", "\\.", -1)
	pattern = strings.Replace(pattern, "*", ".*", -1)
	return "^" + pattern + "$"
}

func keyword(stringsList []string) {

	// 定义匹配模式
	pattern := "WCTrader.pid*.log"

	// 将匹配模式转换为正则表达式
	regexPattern := ConvertPattern(pattern)

	// 编译正则表达式
	re := regexp.MustCompile(regexPattern)

	// 迭代字符串数组并打印匹配的字符串
	for _, str := range stringsList {
		if re.MatchString(str) {
			fmt.Println(str)
		}
	}
}

// Function to find the common prefix of a slice of strings
func commonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Split the first string by "/" to get the initial prefix parts
	prefixParts := strings.Split(strs[0], "/")

	// Iterate over the other strings
	for _, str := range strs[1:] {
		parts := strings.Split(str, "/")
		for i := range prefixParts {
			// If we reach a part that doesn't match, truncate the prefixParts
			if i >= len(parts) || prefixParts[i] != parts[i] {
				prefixParts = prefixParts[:i]
				break
			}
		}
	}

	// Join the common prefix parts back together
	return strings.Join(prefixParts, "/")
}

// getDaysInMonth 返回给定时间所在月份的天数
func getDaysInMonth(date time.Time) int {
	// 当月第一天
	firstDayOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())

	// 下个月第一天
	firstDayOfNextMonth := firstDayOfMonth.AddDate(0, 1, 0)

	// 计算两个日期之间的差值
	daysInMonth := firstDayOfNextMonth.Sub(firstDayOfMonth).Hours() / 24
	return int(daysInMonth)
}

func getdateProgress() {
	// 获取当前时间
	now := time.Now()

	// 获取当前年份
	year := now.Year()

	// 计算年初和年末的时间
	startOfYear := time.Date(year, time.January, 1, 0, 0, 0, 0, now.Location())
	endOfYear := time.Date(year, time.December, 31, 23, 59, 59, 999999999, now.Location())

	// 计算当前时间是一年中的第几天
	dayOfYear := now.YearDay()

	// 计算全年总天数
	totalDays := endOfYear.Sub(startOfYear).Hours()/24 + 1

	// 计算进度百分比
	progress := float64(dayOfYear) / totalDays * 100

	fmt.Printf("当前日期：%s\n", now.Format("2006-01-02"))
	fmt.Printf("今年是：%d\n", year)
	fmt.Printf("今天是一年中的第%d天\n", dayOfYear)
	fmt.Printf("全年共有%.0f天\n", totalDays)
	fmt.Printf("今年的时间进度为：%.2f%%\n", progress)
}

func getDateQuarterProgress() {
	// 获取当前时间
	now := time.Now()

	// 获取当前年份和月份
	year, month, _ := now.Date()

	// 计算当前季度的开始月份
	startMonth := month - time.Month((int(month)-1)%3)

	// 计算当前季度的开始和结束时间
	startOfQuarter := time.Date(year, startMonth, 1, 0, 0, 0, 0, now.Location())
	endOfQuarter := startOfQuarter.AddDate(0, 3, 0).Add(-time.Nanosecond)

	// 计算当前日期是该季度的第几天
	dayOfQuarter := now.Sub(startOfQuarter).Hours()/24 + 1

	// 计算整个季度的天数
	totalDaysInQuarter := endOfQuarter.Sub(startOfQuarter).Hours()/24 + 1

	// 计算进度百分比
	progressInQuarter := (dayOfQuarter / totalDaysInQuarter) * 100

	fmt.Printf("当前日期：%s\n", now.Format("2006-01-02"))
	fmt.Printf("当前季度开始于：%s\n", startOfQuarter.Format("2006-01-02"))
	fmt.Printf("当前季度结束于：%s\n", endOfQuarter.Format("2006-01-02"))
	fmt.Printf("今天是这个季度的第%.0f天\n", dayOfQuarter)
	fmt.Printf("这个季度共有%.0f天\n", totalDaysInQuarter)
	fmt.Printf("当前日期在本季度的进度为：%.2f%%\n", progressInQuarter)
}

func TestDefers(t *testing.T) {
	start := "20240103"
	getdateProgress()
	getDateQuarterProgress()

	t.Log(fmt.Sprintf("%s-%s-%s", start[:4], start[4:6], start[6:]))
	ss := time.Now().Add(-time.Hour * 24 * 30)
	t.Log(getDaysInMonth(ss))
	for counter, n := 0, 2; n >= 0; n-- {
		defer func() {
			print(counter)
			counter++
		}()
	}
}

func TestWeeks(t *testing.T) {
	xAxis := []string{
		"01",
		"02",
		"03",
		"04",
		"05",
		"06",
		"07",
		"08",
		"09",
		"10",
		"11",
		"12"}
	for _, axi := range xAxis {
		week, err := strconv.Atoi(axi)
		if err != nil {
			// log.Error(err.Error())
		}
		// t.Log(strconv.Atoi(axi))
		ts := time.Now()

		// 获取年份的第一天
		firstDayOfYear := time.Date(ts.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
		// 计算年份的第一天是星期几
		firstDayOfWeek := int(firstDayOfYear.Weekday())

		// 计算偏移量以确定第一个周的起始日期
		offset := (8 - firstDayOfWeek) % 7

		// 计算指定周的起始日期
		startDate := firstDayOfYear.AddDate(0, 0, offset+(week-1)*7)
		log.Println(startDate.Format("01-02"))
	}

}

// 根据日期计算所在季度的起始日期和结束日期
func TestQuarterDates(t *testing.T) {
	// 获取当前日期
	now := time.Now()

	// 获取当前季度
	quarter := (now.Month()-1)/3 + 1

	// 获取当前季度开始日期
	quarterStart := time.Date(now.Year(), time.Month(quarter*3-2), 1, 0, 0, 0, 0, now.Location())

	// 计算当前季度已过去的天数
	daysPassed := int(now.Sub(quarterStart).Hours() / 24)

	// 计算当前季度总天数
	daysTotal := int(time.Date(now.Year(), time.Month(quarter*3+1), 0, 0, 0, 0, 0, now.Location()).Sub(quarterStart).Hours() / 24)

	// 计算当前季度进度
	progress := float64(daysPassed) / float64(daysTotal)

	// 输出进度
	fmt.Println("当前季度进度：", progress)
}

// 整形 转换特殊case
func TestStringToInt(t *testing.T) {
	t.Log(cast.ToInt64("09"))
	t.Log(cast.ToInt64("02"))
	t.Log(cast.ToInt64("03"))
	t.Log(cast.ToInt64("029"))
	t.Log(cast.ToInt64("079"))
	t.Log(strconv.Atoi("079"))
	t.Log(strconv.Atoi("079"))
	t.Log(strconv.Atoi("029"))
	s := "1048576"
	i, err := strconv.Atoi(s)
	t.Log(i, err)
	ii, err := strconv.ParseInt(s, 10, 128)
	t.Log(ii, err)
	ii, err = strconv.ParseInt(s, 5, 64)
	t.Log(ii, err)
	assert.Error(t, err)
	assert.Equal(t, i, 1048576)

}

func TestCheckInterfaceNil(t *testing.T) {
	aa := getintface()
	t.Log(aa == nil)
	nn := GetNil()
	t.Log(nn == nil)
	t.Log(reflect.ValueOf(nn).IsNil())

}

func GetNil() *Baz {
	return nil
}
func getintface() any {
	return GetNil()
}

func TestMoveBit(t *testing.T) {
	sec := uint64(1234567890)
	shifted := uint64(sec) >> 33  // 0
	fmt.Println(uint64(sec) >> 1) // 308641972.5
	fmt.Println(uint64(sec) >> 2) // 617283945
	fmt.Println(shifted)          // 0
	aa := "03"
	fmt.Println(strconv.Atoi(aa))
	ww := "33"
	fmt.Println(strconv.Atoi(ww))
	w := 3
	fmt.Printf("%02d", w)

}

// tcp拥塞控制实现
func TestTcp(t *testing.T) {
	congestionWindow := initialWindowSize
	slowStart := true

	for round := 1; round <= 20; round++ {
		// Simulate data transmission and reception
		fmt.Printf("Round %d - Congestion Window: %d\n", round, congestionWindow)

		// Simulate successful data transmission and reception
		ackReceived := simulateTransmission()

		// Adjust congestion window based on acknowledgment
		if ackReceived {
			if slowStart {
				// 当处于慢启动阶段，就指数级扩大窗口，直到达到慢启动阈值
				congestionWindow *= 2
				if congestionWindow >= slowStartThreshold {
					slowStart = false
					fmt.Println("Entering Congestion Avoidance phase")
				}
			} else {
				// 当处于大于启动阈值后，窗口大小依次递增
				congestionWindow += 1
			}
		} else {
			// 当发生传输失败情况，就把窗口大小减半，重新进入慢启动阶段
			fmt.Println("Timeout or packet loss detected. Reducing Congestion Window.")
			congestionWindow /= 2
			slowStart = true
		}

		// Introduce a delay to simulate network conditions
		time.Sleep(time.Millisecond * 100)
	}
}

func simulateTransmission() bool {
	// Simulate successful transmission with a 90% probability
	randNum := rand.Float64()
	return randNum < 0.9
}

// RemoveANSIEscapeCodes 函数用于移除字符串中的 ANSI 转义序列
func RemoveANSIEscapeCodes(s string) string {
	// ANSI 转义序列的正则表达式
	re := regexp.MustCompile(`\x1B\[[0-?]*[ -/]*[@-~]`)
	// 使用正则表达式替换掉所有匹配的 ANSI 转义序列

	return re.ReplaceAllString(s, "")
}

// ContainsANSIEscapeCodes 函数检查字符串中是否包含 ANSI 转义序列
func ContainsANSIEscapeCodes(s string) bool {
	// ANSI 转义序列的正则表达式
	re := regexp.MustCompile(`\x1B\[[0-?]*[ -/]*[@-~]`)
	// 检查字符串是否匹配正则表达式

	return re.MatchString(s)
}

func TestStringsTrim(t *testing.T) {
	src := "/da/sd/as/d/"
	t.Log(strings.Trim(src, "/"))
	t.Log(strings.Index(src, "/d/"))
	t.Log(len(src))
}

func TestAnsi(t *testing.T) {
	ansiStr := `[01m[K/src/WCTrader/include/Brokers/KMax/AdaptedBrokerApiTypes.h:50:56:[m[K [01;31m[Kerror: [m[K'[01m[KAlgoType[m[K' in namespace '[01m[Kcom::kmax::api[m[K' does not name a type`
	ansiStr = `go build build_parse^[[01;31m^[[Kerror: ^[[m^[[K`
	println(ContainsANSIEscapeCodes(ansiStr))
	// 移除 ANSI 转义序列
	normalStr := RemoveANSIEscapeCodes(ansiStr)

	// 输出结果
	println(normalStr)
}

func TestCopyBaz(t *testing.T) {
	info := &Baz{
		Bar: 1,
		Foo: 1,
	}
	info2 := &baz2{}

	err := copier.Copy(&info2, &info)
	if err != nil {
		t.Log(err)
	}
	t.Log(info2)
}

func TestAppendWhenCloseSlice(t *testing.T) {
	arr := make([]Baz, 0, 10)
	arr2 := make([]Baz, 0, 10)
	copy(arr2, arr)
	arr = append(arr, Baz{1, 2})
	arr = append(arr, Baz{1, 2})
	arr = append(arr, Baz{1, 2})
	t.Log(len(arr), cap(arr))
	// arr = arr[0:0]
	arr = append(arr, Baz{1, 2})

	arr = append(arr, Baz{1, 22})
	arr2 = append(arr2, Baz{1, 44})

	t.Log(len(arr), cap(arr))
	t.Log(len(arr2), cap(arr2))
	t.Log(arr)
	t.Log(arr2)

}

func TestAAA(t *testing.T) {
	t.Log(parseTimestamp(93428320))
	t.Log(parseTimestamp(91500260))
}

// parseTimestamp 将HHMMSSsss格式的时间戳转换为自午夜以来的总毫秒数
func parseTimestamp(ts uint32) int {
	hours := ts / 10000000
	minutes := (ts % 10000000) / 100000
	seconds := (ts % 100000) / 1000
	milliseconds := ts % 1000
	return int(hours)*3600000 + int(minutes)*60000 + int(seconds)*1000 + int(milliseconds)
}

// formatTimestamp 将自午夜以来的总毫秒数转换回HHMMSSsss格式的时间戳
func formatTimestamp(ms int) uint32 {
	hours := ms / 3600000
	minutes := (ms % 3600000) / 60000
	seconds := (ms % 60000) / 1000
	milliseconds := ms % 1000
	return uint32(hours)*10000000 + uint32(minutes)*100000 + uint32(seconds)*1000 + uint32(milliseconds)
}

func maxSum(a []int, q [][]int) int {
	// n := len(a)
	m := len(q)
	sums := make([]int, m)

	for i := 0; i < m; i++ {
		queries := q[i]
		start, end := queries[0], queries[1]

		// 枚举所有可能的排序
		for _, permutation := range permutationsMax(a) {
			// 计算当前排序下的部分和
			currentSum := 0
			for j := start; j < end; j++ {
				currentSum += permutation[j]
			}

			// 更新最大部分和
			if currentSum > sums[i] {
				sums[i] = currentSum
			}
		}
	}

	// 计算所有查询的总和
	totalSum := 0
	for _, s := range sums {
		totalSum += s
	}

	return totalSum
}

// 生成数组的所有可能排列
func permutationsMax(a []int) [][]int {
	var result [][]int
	var backtrackitem func(start int)
	backtrackitem = func(start int) {
		if start == len(a)-1 {
			tmp := make([]int, len(a))
			copy(tmp, a)
			result = append(result, tmp)
			return
		}
		for i := start; i < len(a); i++ {
			a[start], a[i] = a[i], a[start]
			backtrackitem(start + 1)
			a[start], a[i] = a[i], a[start]
		}
	}
	backtrackitem(0)
	return result
}
func TestMaxSum(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	intervals := [][]int{{2, 4}, {1, 3}, {5, 6}}
	// 2, 3, 4
	// 1, 2, 3
	// 5, 6
	result := maxSum(a, intervals)
	fmt.Println("Maximum sum:", result)
}

func TestFilename(t *testing.T) {
	// confName := "/Users/darren/go/src/person-go/modgo/test/func.go"
	confName := "./func.go"
	data, err := os.ReadFile(confName)
	t.Log(string(data), err)
}

// 校验切片是否被安全释放
func TestSliceCut(t *testing.T) {
	sk := make([]string, 20)
	for i := 1; i < 800; i++ {
		if i%10 == 0 {
			runtime.GC()
		}
		if i%50 == 0 {
			println(len(sk), cap(sk))
		}
		if i < 20 {
			sk = append(sk, "test")
		}
		if i < 80 && i >= 20 {
			if len(sk) >= 2 {
				sk = sk[2:]
			}
		}
		if i > 80 && i < 200 {
			sk = append(sk, "test")
		}
		if i < 300 && i >= 200 {
			if len(sk) >= 2 {
				sk = sk[2:]
			}
		}
		if i < 800 && i >= 300 {
			sk = append(sk, "test")
		}

	}
	println("len(sk), cap(sk)")
	println(len(sk), cap(sk))
	time.Sleep(time.Second)
	runtime.GC()
	println(len(sk), cap(sk))

}

func TestHttpMax(t *testing.T) {
	targetHost := "http://172.16.54.182:8081/debug"
	concurrency := 2000
	maxIdleConnsPerHost := 3

	transport := &http.Transport{
		MaxIdleConnsPerHost: maxIdleConnsPerHost,
	}

	client := &http.Client{Transport: transport}

	var wg sync.WaitGroup
	incr := atomic.Int64{}
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			req, err := http.NewRequest("POST", targetHost, nil)
			if err != nil {
				fmt.Printf("Error creating request: %v\n", err)
				return
			}

			// // 禁用连接复用
			// req.Close = true
			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("Error sending request: %v\n", err)
				return
			}
			defer resp.Body.Close()
			incr.Add(1)
			fmt.Printf("Request %d completed\n", i)
		}(i)
	}

	wg.Wait()
	fmt.Println("All requests completed.", incr.Load())
}

func Test41600(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Length", "0")
	}))
	defer ts.Close()

	client := ts.Client()
	transport := client.Transport.(*http.Transport)
	transport.MaxIdleConns = 1
	transport.MaxConnsPerHost = 1

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for ctx.Err() == nil {
				reqctx, reqcancel := context.WithCancel(context.Background())
				go reqcancel()
				req, _ := http.NewRequestWithContext(reqctx, http.MethodGet, ts.URL, nil)
				rsp, err := client.Do(req)
				if err == nil {
					defer rsp.Body.Close()
				}
			}
		}()
	}

	for {
		req, _ := http.NewRequest(http.MethodGet, ts.URL, nil)
		if rsp, err := client.Do(req); err != nil {
			t.Errorf("unexpected: %p %v", req, err)
			break
		} else {
			rsp.Body.Close()
		}
	}

	cancel()
	wg.Wait()
}

func (c *ConfigOne) Strings() string {
	return fmt.Sprintf("print: %v", c)
}

func TestErrPrint(t *testing.T) {

	s := &ConfigOne{}
	s.Strings()
}

// === RUN   TestErrPrint
//    tmain_test.go:74: 1 1
//    tmain_test.go:76: 2 <nil>
//    tmain_test.go:79: 2 <nil>
//    tmain_test.go:82: 1 1
// --- PASS: TestErrPrint (0.00s)
// PASS

func TestForTime(t *testing.T) {

	ts := time.Now()
	i := -3

	ss, ww := json.Marshal(i)
	t.Log(ss, ww)

	www := []byte{44, 45, 97, 99}
	t.Log(string(www))
	for ts.Sub(ts.AddDate(0, 0, i)).Hours() >= 24 {
		log.Println(ts.AddDate(0, 0, i).Unix())
		// log.Println(ts.AddDate(0, 0, -1).Unix())
		i++
	}
}

// 指针赋值
func TestPointSet(t *testing.T) {
	t.Logf("%.f", 21212.13213)

	type Person struct {
		Name string
		Age  int
	}
	// 创建一个 Person 实例
	person := Person{Name: "Alice", Age: 30}
	var p = &person
	height := unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.Age))
	*((*int)(height)) = 100 // 将height的值改为100
	fmt.Println(p.Age)

	height2 := unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.Name))
	*((*string)(height2)) = "darren" // 将height的值改为100
	fmt.Println(p.Name)
	fmt.Println(person)

}

//go:linkname FastRand runtime.fastrand
func FastRand() uint32

//go:linkname FastRandN runtime.fastrandn
func FastRandN(n uint32) uint32
func TestFastRand(t *testing.T) {
	for i := 0; i < 1000; i++ {
		t.Log(FastRand())
		t.Log(FastRandN(3))
	}
}

// 求最大子序列
func TestDpArr(t *testing.T) {
	arr := []int{199, -500, 2, 10, 30, 50, -13}
	l := len(arr)
	dp := make([]int, l)
	dp[0] = arr[0]
	index := 0
	for i := 1; i < l; i++ {
		maxs := dp[i-1]
		if maxs > dp[i-1]+arr[i] {
			if arr[i] < 0 {
				dp[i] = 0
				continue
			}
			dp[i] = arr[i]
		} else {
			dp[i] = maxs + arr[i]
			index = i
		}
	}

	sum := dp[index]
	left := index
	for sum != 0 && left >= 0 {
		sum -= arr[left]
		left--
	}
	fmt.Println(arr[left+1 : index+1])

}

// 这个并发例子说明 提前声明好map所有容量后，可以并发写入，不会报错，因为不会发生扩容。 如果填写m[n%9] 大于原始规格，就可能发生：
// fatal error: concurrent map read and map write
// goroutine 5091828 [running]:
func BenchmarkMapWrite(b *testing.B) {
	m := make(map[int]*baz2)
	m[0] = &baz2{Bar: 1, Foo: 1, fzz: []int{1, 2, 3}}
	m[1] = &baz2{Bar: 1, Foo: 1, fzz: []int{1, 2, 3}}
	m[2] = &baz2{Bar: 1, Foo: 1, fzz: []int{1, 2, 3}}
	for n := b.N; n > 0; n-- {
		n := n
		go func() {
			if _, ok := m[n%3]; !ok {
				// return
				m[n%3] = &baz2{Bar: 1, Foo: 1, fzz: []int{1, 2, 3}}
			}
			m[n%3].Bar = n
			m[n%3].Foo = n
		}()
	}

}

func TestYushu(t *testing.T) {
	fastrand := uint64(717)
	n := uint64(4)
	t.Log(uint32(uint64(fastrand) * uint64(n) >> 32))
	t.Log(uint32(fastrand % n))
}

/*
对比注释runtime.Gosched() 前后的输出结果，可以看出runtime.Gosched()的作用是让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。
=== RUN   TestGosched
tmain_test.go:79: after gosched1
tmain_test.go:91: after gosched2
tmain_test.go:75: test1
tmain_test.go:82: test2
tmain_test.go:87: test3
tmain_test.go:104: ctx done
=== RUN   TestGosched
tmain_test.go:75: test1
tmain_test.go:79: after gosched1
tmain_test.go:82: test2
tmain_test.go:91: after gosched2
tmain_test.go:87: test3
tmain_test.go:104: ctx done
--- PASS: TestGosched (2.00s)
*/
func TestGosched(t *testing.T) {
	ctx, caa := context.WithTimeout(context.Background(), 4*time.Second)
	runtime.GOMAXPROCS(1)
	stopGroup, ctx := errgroup.WithContext(ctx)
	stopGroup.Go(func() error {
		t.Log(errors.New("test1"))
		return nil
	})
	runtime.Gosched()
	t.Log("after gosched1")

	stopGroup.Go(func() error {
		t.Log(errors.New("test2"))
		return nil
	})
	stopGroup.Go(func() error {
		time.Sleep(2 * time.Second)
		t.Log(errors.New("test3"))
		return nil
	})
	runtime.Gosched()
	t.Log("after gosched2")

	go func() {
		err := stopGroup.Wait()
		if err != nil {
			caa()
			t.Log(err)
			return
		}
	}()

	select {
	case <-ctx.Done():
		t.Log("ctx done")
	case <-time.After(3 * time.Second):
		t.Log("time after")
	}

}

// 主要测试一下，runtime.Gosched() 在多个协程中的作用，是否会让出时间片。其次就是ctx超时退出机制要在协程中时刻监听才有效
func TestErrGroup(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	runtime.GOMAXPROCS(1)
	stopGroup, ctx := errgroup.WithContext(ctx)
	fn := func(ctx context.Context, i int) error {
		for {
			select {
			case <-ctx.Done():
				return nil
			default:
				// time.Sleep(time.Millisecond * 10)
				log.Println("test", i)
				runtime.Gosched()
			}
		}
	}
	stopGroup.Go(func() error {
		return fn(ctx, 1)
	})
	stopGroup.Go(func() error {
		return fn(ctx, 2)
	})
	time.Sleep(5 * time.Second)

}

// 控制多协程作业，超时退出
func TestMultiGoroutine(t *testing.T) {
	target := int64(9910)
	ch := make(chan struct{}, 1)
	var wg sync.WaitGroup
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)
	fn := func(c context.Context, data []int64) {
		defer wg.Done()
		for _, i := range data {
			select {
			case <-c.Done():
				return
			default:
				if i == target {
					log.Println("find it222", c.Err())
					ch <- struct{}{}
					return
				}
			}
			time.Sleep(time.Millisecond * 1800)
		}
		return
	}
	arr := make([]int64, 0, 10000)
	for i := 0; i < 100000; i++ {
		arr = append(arr, int64(i))
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	go func() {
		<-c
		cancel()
	}()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go fn(ctx, arr[(2000*(i)):(2000*(i+1))-1])
	}
	go func() {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				t.Log("timeouted ok")
			}
			t.Log("主动结束 ok222")
		case <-ch:
			cancel()
			t.Log("find it", ctx.Err())
		}
	}()
	wg.Wait()
	close(ch)
}

// 读取并转换900KB变量
func TestBufferToByte(t *testing.T) {
	// 创建一个900KB大小的变量
	dataB := make([]byte, 900*1024) // 900KB
	for i := 0; i < 900*1024; i++ {
		dataB[i] = byte('1' + (i % 9)) // '1' to '9'
	}
	// 假设把这个数据转移到一个新的结构上
	type NewStruct struct {
		Data []byte
	}
	var ns NewStruct

	// 创建一个bytes.Reader对象
	reader := bytes.NewReader(dataB)

	// 使用ioutil.ReadAll读取Reader对象的所有字节
	readData, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Error reading dataB:", err)
		return
	}
	// 将读取到的字节转移到新的结构上
	ns.Data = readData

	t.Log(string(ns.Data[:20]), string(dataB[:20]))
	// 验证转移是否成功
	fmt.Println(len(ns.Data) == len(dataB)) // 应该打印 "true"

}

func TestPrintNum(t *testing.T) {
	even := make(chan bool)
	odd := make(chan bool)

	fn := func(self, other chan bool, start int) {
		for i := start; i <= 100; i += 2 {
			<-self // 等待自己的轮次
			fmt.Println(i)
			other <- true // 唤醒另一个协程
		}
	}

	go fn(even, odd, 1) // 协程打印奇数
	go fn(odd, even, 2) // 协程打印偶数
	even <- true        // 启动第一个协程
	wg.Add(1)
	wg.Wait()
	// even <- true        // 启动第一个协程
	// var input string
	fmt.Scanln() // 等待两个协程结束
}

// 协程交叉打印数字 V2
func TestPrintNumV2(t *testing.T) {
	var m sync.Mutex
	cond := sync.NewCond(&m)

	var turn = 1
	upperLimit := 100

	printOdd := func() {
		for i := 1; i <= upperLimit; i += 2 {
			m.Lock()
			for turn != 1 {
				cond.Wait()
			}

			fmt.Println(i)
			turn = 2
			cond.Signal()
			m.Unlock()
		}
	}

	printEven := func() {
		for i := 2; i <= upperLimit; i += 2 {
			m.Lock()
			for turn != 2 {
				cond.Wait()
			}

			fmt.Println(i)
			turn = 1
			cond.Signal()
			m.Unlock()
		}
	}

	go printOdd()
	go printEven()

	time.Sleep(time.Second * 10)
	// 等待协程完成
	scanln, err := fmt.Scanln()
	t.Log(scanln)
	if err != nil {
		return
	}
}

func TestRandtime(t *testing.T) {
	// rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		t.Log(rand.Intn(100))
	}

}

func removeQuotes(s string) string {
	// Remove both single and double quotes
	return strings.Trim(s, `"'`)
}
func TestIoAll(t *testing.T) {
	sc := `{"a":1,"b":2}`
	reader := strings.NewReader(sc)
	var sss strings.Builder
	verb := make([]byte, 3)
	for {
		full, err := reader.Read(verb)
		if err != nil {
			t.Log(err)
			break
		}
		if full == 0 {
			t.Log("full==0")
			break
		}

		sss.Write(verb)

		t.Log(string(verb))
	}
	t.Log(sss.String(), sss.Len())

}

func BenchmarkInt63Threadsafe(b *testing.B) {
	for n := b.N; n > 0; n-- {
		b.Log(rand.Intn(10))
		// b.Log(rand.Int())
	}
}

func findSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	totalWords := len(words)
	totalLen := wordLen * totalWords
	result := make([]int, 0)

	if len(s) < totalLen {
		return result
	}

	wordCounts := make(map[string]int)
	for _, w := range words {
		wordCounts[w]++
	}

	for i := 0; i <= len(s)-totalLen; i++ {
		seen := make(map[string]int)
		j := 0
		for ; j < totalWords; j++ {
			sub := s[i+j*wordLen : i+(j+1)*wordLen]
			if _, ok := wordCounts[sub]; !ok {
				break
			}
			seen[sub]++
			if seen[sub] > wordCounts[sub] {
				break
			}
		}
		if j == totalWords {
			result = append(result, i)
		}
	}

	return result
}
func TestReverssss(t *testing.T) {
	aa := findSubstring("barfoofoobarthefoobarman", []string{"Bar", "Foo", "the"})
	t.Logf("%v", aa)
	t.Log(99 % 9)
}

func TestGroutineDFS(t *testing.T) {
	ctx := context.TODO()
	ctx, cancel := context.WithCancel(ctx)
	GOCTX(ctx)
	contexts = append(contexts, ctx)
	defer cancel()
	for _, ctx2 := range contexts {
		if ctx2.Err() != nil {
			t.Log(ctx2.Err())
		}
	}
}

var contexts []context.Context

type otherContext struct {
	context.Context
}

func GOCTX(ctx context.Context) {
	o := otherContext{ctx}
	c, _ := context.WithCancel(o)
	fmt.Println(c.Err())
	contexts = append(contexts, c)
}

func ff() {
	fmt.Println("ff")
}
func TestForChan(t *testing.T) {
	ch := make(chan func(), 1)
	var f3 func() = ff
	go func() {
		for {
			ch <- f3
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		select {
		case f := <-ch:
			f()

		}
	}
	println("ok  ")

}

func TestBigSlice(t *testing.T) {
	arr := []int{0, 5: 0, 2}
	t.Log(len(arr))
}

var x int

func fff() bool {
	x++
	return x < 5
}

func TestFors(t *testing.T) {
	for fff(); fff(); fff() {
		println("ok", x)
	}
}

func TestSignal(t *testing.T) {
	hookSignals()
}

// 用于阻塞主进程，等待信号
func hookSignals() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(
		sigChan,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	switch <-sigChan {
	case syscall.SIGQUIT: // terminate now
		log.Printf("Server Stop")
	case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
		log.Println("Server GracefulStop")
	default:
		log.Println("Server Unknown")
	}
}

// 这个例子说明 即便存在协程错误，都需要等全部协程结束后，才会返回错误，会有阻塞风险问题，建议使用ctx来控制协程超时，这样不影响主流程等待阻塞。
func TestErrGroup1(t *testing.T) {
	c, _ := context.WithTimeout(context.Background(), 1*time.Second)
	cc, _ := errgroup.WithContext(c)
	cc.Go(func() error {
		t.Log("ok", "aa")
		time.Sleep(2 * time.Second)
		return nil
	})
	cc.Go(func() error {
		t.Log("ok", "aa")
		time.Sleep(1 * time.Second)
		return errors.New("err")
	})
	cc.Go(func() error {
		t.Log("ok", "aa")
		time.Sleep(13 * time.Second)
		return nil
	})

	err := cc.Wait()
	if err != nil {
		t.Log("err", err)
		return
	}

}

func TestGoCtx(t *testing.T) {
	// ctx1, cancel1 := context.WithTimeout(context.Background(), 10*time.Second)
	ctx1 := context.Background()
	// defer cancel1()
	ctx := context.WithValue(ctx1, "key", "value")
	ctx, cancel := context.WithCancelCause(ctx)

	go func() {
		time.Sleep(4 * time.Second)
		cancel(errors.New("不会再打印了"))
	}()

	go func() {
		time.Sleep(1 * time.Second)
		cancel(errors.New("test"))
	}()
	for {
		t.Log(ctx.Err())
		t.Log(context.Cause(ctx))
		t.Log(context.Cause(ctx1))

		time.Sleep(500 * time.Millisecond)
	}
}

func betch() []int {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	ctx = context.WithValue(ctx, "key", "value")
	flows := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var w sync.WaitGroup
	res := []int{}
	for _, flow := range flows {
		flow := flow
		w.Add(1)
		RecoverGO(func() {
			defer w.Done()
			// ctxx, cel := context.WithTimeout(ctx, 4*time.Second)
			// defer cel()
			ctxx := ctx
			a := GetA(ctxx, flow)
			res = append(res, a)
		})
	}
	RecoverGO(func() {
		w.Wait()
		cancel()
	})
	select {
	case <-ctx.Done():
		log.Println("Timeout  goroutines occurred.")
	}
	return res
}

func GetA(ctx context.Context, f int) int {
	time.Sleep(time.Duration((f)) * time.Second)
	log.Println(f, ctx.Value("key"))
	if ctx.Err() != nil {
		log.Println(ctx.Err().Error())
	}
	return f
}

func TestMapStruct(t *testing.T) {
	foos := make(map[int]*w2)
	foos[0] = &w2{q: 1}
	m1 := make(map[int]decimal.Decimal)
	m1[0].Add(decimal.NewFromFloat(2.0))
	fmt.Printf("m1: %v", foos[0])
	fmt.Printf("m1: %v", m1[0])
}

func TestGjson(t *testing.T) {
	jsonBs := `[
      {
      "first": "last",
      "last": "Prichard"
    },
    {
      "first": "Janet",
      "last": "Prichard"
    }
]`
	value := gjson.Get(jsonBs, "name.last")
	println(value.String())
	value = gjson.Get(jsonBs, "#.first")
	println(value.String())
	println(gjson.Get(jsonBs, "#").Int())
	t.Log(gjson.Get(jsonBs, "#.last").Array()[1])

}

func TestInterfaceV(t *testing.T) {
	type options struct {
		InfoLogger interface{ Infof(string, ...any) }
	}
	o := &options{}
	o.InfoLogger = new(Foo)
	o.InfoLogger.Infof("test")
}

type Foo struct {
}

func (*Foo) Infof(s string, any ...interface{}) {
	println(s)

}

func TestSliceContains(t *testing.T) {
	sql := "11,3222,33,2233"
	sql2 := "’11‘“看到痰"
	t.Log(utf8.RuneCountInString(sql2), len(sql2))
	sqls := strings.Split(sql, ",")
	t.Log(slices.Contains(sqls, "22"))
	Start := "2029-09-11"
	t.Log(strings.Replace(Start, "-", "", -1))
}

// 并发写入ctx， 需要加锁
func TestWatchContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	var r sync.Mutex
	var s int32
	for i := 0; i < 100; i++ {
		i := i
		go func() {
			r.Lock()
			defer r.Unlock()
			if atomic.AddInt32(&s, 1) == 70 {
				cancel()
			}
			select {
			case <-ctx.Done():
				// t.Log(i, "主动退")
				return
			default:
			}

			ctx = context.WithValue(ctx, fmt.Sprintf("key %d", i), fmt.Sprintf("key %d", i))
			t.Log(i, ctx.Err())

		}()
	}
	time.Sleep(2 * time.Second)
	for i := 0; i < 100; i++ {
		if v := ctx.Value(fmt.Sprintf("key %d", i)); v != nil {
			t.Log(v)
		}
	}
	return
}

func TestWatchCan(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "【监控1】")
	go watch(ctx, "【监控2】")
	go watch(ctx, "【监控3】")

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	// 为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return

		}
	}
}
func TestRandTimeMin(t *testing.T) {
	ts := time.Now()

	ttt, _ := time.ParseInLocation("2006-01-02 15:04:05", "2024-09-23 10:00:00", time.Local)
	t.Log(ts.Unix())
	t.Log(ttt.Unix())
	t.Log(ts.Sub(ttt).Seconds())
	sleepHour := 28 - ts.Hour()
	a := rand.NewSource(1)
	a.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {

		a := time.Duration(sleepHour)*time.Minute*60 - time.Duration(rand.Intn(20))*time.Minute
		// a := time.Duration(sleepHour)*time.Minute*60 - time.Duration(rand.Intn(20))*time.Minute
		t.Log(a.Minutes())
		t.Log(a)
	}
}
func TestFuncnil(t *testing.T) {
	fn := func(x int) { print(x) }
	pn := func() int {
		fn = nil
		return 1
	}
	defer fn(pn())
}
func TestFuncnil2(t *testing.T) {
	fn := func(x int) { print(x) }
	pn := func() int {
		fn = nil
		return 2
	}
	fn(1)
	fn(pn())
	pn()
}

// 拷贝指针结构体
func TestCopystruct(t *testing.T) {
	a := &baz2{Bar: 1, Foo: 1, fzz: []int{1, 2, 3}}
	b := &baz2{Bar: 3}
	*b = *a
	a.Bar = 2
	t.Log(a, b)
}

func xielou() { // 待测试的方法
	// ch := make(chan struct{}, 1)
	ch := make(chan struct{})
	go func() {
		ch <- struct{}{}
	}()
	// <-ch
}

func TestXielou(t *testing.T) {
	// defer goleak.VerifyNone(t)
	xielou()
	goLeakCheck()
	// time.Sleep(3 * time.Second)
}

func goLeakCheck() {
	ticker := time.NewTimer(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			RecoverGO(
				func() {
					err := goleak.Find()
					if err != nil {
						fmt.Println("定时打印当前进程中存在内存泄露风险的代码段：goleak check")
						fmt.Println(err)
					}
				},
			)
		}
	}

}

func TestAppendToSlice(t *testing.T) {

	fn := func(res *[][]int) {
		r := *res
		r = append(r, []int{1})

		// *res = r
	}
	rr := make([][]int, 0)
	fn(&rr)
	fmt.Println("\n", len(rr))

	type A struct {
		Arr [][]int
	}
	funoo := func(k *A) {
		k.Arr = append(k.Arr, []int{1})
	}
	a := A{}
	funoo(&a)
	fmt.Println("\n", len(a.Arr))

}

func TestFloatString(t *testing.T) {
	f, ok1 := decimal.NewFromString("3421.12")

	dd := f
	dd2 := decimal.NewFromInt(4444)
	a := dd.Mul(decimal.NewFromInt(100)).Div(dd2).Sub(decimal.NewFromInt(1)).String()
	t.Log(a)
	t.Log(dd.Div(dd2).Sub(decimal.NewFromInt(1)).Mul(decimal.NewFromInt(100)))
	t.Log(dd.Mul(decimal.NewFromInt(100)).Div(dd2))
	t.Log(dd.Mul(decimal.NewFromInt(100)))
	aa := float64(0.00)
	t.Log(decimal.NewFromFloat(aa).IsZero(), 33333333)
	if ok1 != nil {
		return
	}
	v, ok := f.Float64()
	t.Log(v, ok, ok1)
}

func TestSliceFunc(t *testing.T) {

	// arr := []int{1, 2, 3}
	// lo.KeyBy()

	m := lo.Map[int64, string]([]int64{1, 2, 3, 4}, func(x int64, index int) string {
		println(index)
		return strconv.FormatInt(x+1, 10)
	})
	t.Log(len(m), m)
}

func TestLoAssociate(t *testing.T) {
	list := []string{"a", "aa", "aaa"}

	result := lo.SliceToMap(list, func(str string) (string, int) {
		return str, len(str)
	})

	fmt.Printf("%v", result)
}

func TestLoFindKeyBy(t *testing.T) {
	// lo.Find()
	list := make(map[string]*ConfigOne, 0)
	list["22"] = &ConfigOne{
		Daemon: "sss",
	}
	list["233"] = &ConfigOne{
		Daemon: "ssa",
	}
	list["54"] = &ConfigOne{
		Daemon: "ssa",
	}

	result, ee := lo.FindKeyBy(list, func(str string, v *ConfigOne) bool {
		if v.Daemon == "sss" {
			return false
		}
		return true
	})

	fmt.Printf("%v===", result)
	fmt.Printf("%v---", ee)
}

func TestStructIsNil(t *testing.T) {
	for i := 0; i < 10; i++ {
		var Foo []*Baz
		t.Log(Foo == nil)
		Foo = append(Foo, &Baz{Bar: 1, Foo: 1})
		t.Log(Foo == nil)
	}
}

func TestGoGroutines(t *testing.T) {
	arr := make(arrStruct, 0, 1000)
	for i := 1; i < 1000; i++ {
		arr = append(arr, Baz{Bar: i, Foo: i})
	}

	fn := func(k int, v *Baz, list arrStruct) {
		for i := range list {
			if i == k {
				continue
			}
			if list[i] == *v {
				fmt.Println("not match ", list[i], v, i)
			}
		}

	}
	wg := sync.WaitGroup{}
	wg.Add(len(arr))

	for i, b := range arr {
		go fn(i, &b, arr)
		wg.Done()
	}
	wg.Wait()

}

func TestSigns(t *testing.T) {
	// 用非阻塞方案 可以提前保留信号，然后到时间后再处理
	c := make(chan os.Signal, 1)

	// 用阻塞方案 在sleep期间，信号会被阻塞，直到sleep结束，所以需要再按一次
	// c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	time.Sleep(5 * time.Second)

	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
}

// 错误的调用示例
func TestPanics(t *testing.T) {
	fn := func() {
		fmt.Println("recobered: ", recover())

	}
	defer func() {
		fn() // 错误方式

		// fmt.Println("recobered: ", recover()) // 正确方式

	}()
	panic("not good")
	// time.Sleep(time.Second)
}

func TestQuicksss(t *testing.T) {
	arr := []int{3, 5, 1, 9, 8, 344, 1, 5555, 44, 2}
	fmt.Println(arr)
	quicksss(arr, 0, len(arr)-1)
	fmt.Println(arr)

}

func quicksss(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := continaes(arr, l, r)
	quicksss(arr, l, mid-1)
	quicksss(arr, mid+1, r)
}
func continaes(arr []int, l, r int) int {
	if l >= r {
		return l
	}
	s := l
	for l < r {
		for arr[s] >= arr[l] && l < r {
			l++
		}
		for arr[s] <= arr[r] && l < r {
			r--
		}
		if arr[r] != arr[l] {
			arr[r], arr[l] = arr[l], arr[r]
		}
	}
	arr[l], arr[s] = arr[s], arr[l]

	return l
}

type Fn struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	D string `json:"d"`
	*bazs
}
type bazs struct {
	Bar int `json:"Bar"`
	Foo int `json:"Foo"`
}
type MacStruct interface {
	Open() int
	Close() bool
}
type M2 struct {
	OK    bool
	Other string
}

func (m *M2) Close() bool {
	m.OK = false
	return m.OK
}
func (m *M2) Open() int {
	return 1
}

func NewM2() MacStruct {
	return &M2{OK: true}
}

func TestInterfaceDesign(t *testing.T) {
	m := NewM2()
	t.Log(m.Open())
	t.Log(m.Close())

}

func TestStructEmbed(t *testing.T) {
	js := `{"a":"2","b":"2","Bar":1,"Foo":10}`
	bs := Fn{"a", "ww", "qq", "www", &bazs{
		Bar: 1,
		Foo: 2,
	}}

	bb, err := json.Marshal(bs)
	t.Log(string(bb))
	err = json.Unmarshal([]byte(js), &bs)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(bs.A, bs.Foo)

}

func TestSpinsWords(t *testing.T) {
	t.Log(reverse(-123456789))
	t.Log(reverse(123))
	t.Log(SpinWords("www abcdefg eeee"))
}
func reverse(x int) int {
	n := 0
	for x != 0 {
		n = n*10 + x%10
		x = x / 10
	}
	return n
}

func SpinWords(str string) string {
	space := " "
	str2 := ""
	// arr := []string{}
	m, n := 0, 0
	for i, k := range str {
		if string(k) == space {
			m = i
		}
		if m != n {
			str2 += " "
			// println(strings.Trim(str[n:m], " "))
			arr := strings.Trim(str[n:m], " ")
			if len(arr) <= 5 {
				for j := range arr {
					str2 += string(arr[j])
				}
			} else {
				for j := range arr {
					str2 += string(arr[len(arr)-1-j])
				}
			}
			n = m
		}
	}
	if n != len(str) {
		str2 += " "

		arr := strings.Trim(str[n:], " ")
		if len(arr) <= 5 {
			for j := range arr {
				str2 += string(arr[j])
			}
		} else {
			for j := range arr {
				str2 += string(arr[len(arr)-1-j])
			}
		}
	}
	// str2 += " "
	//
	// if len(arr) <= 5 {
	// 	for j := range arr {
	// 		str2 += arr[j]
	// 	}
	// 	arr = []string{}
	// } else {
	// 	for j := range arr {
	// 		str2 += arr[len(arr)-1-j]
	// 	}
	// 	arr = []string{}
	// }
	return strings.TrimLeft(str2, " ")
} // SpinWords

func TestMatrix(t *testing.T) {
	t.Log(Determinant([][]int{{1}}))
	t.Log(Determinant([][]int{{1, 3}, {2, 5}}))
	t.Log(Determinant([][]int{{2, 5, 3}, {1, -2, -1}, {1, 3, 4}}))
}

func Determinant(matrix [][]int) int {
	// your code here
	line := len(matrix[0])
	y := len(matrix)

	if line == 1 {
		return matrix[0][0]
	}

	if line == 2 {
		return cal(matrix)
	}
	temp := 0
	for i := 0; i < y; i++ {
		symbol := 1
		if i%2 != 0 {
			symbol = -1
		}
		temp += symbol * matrix[0][i] * Determinant(cal2(matrix, 0, i))
	}
	return temp
}

func cal(matrix [][]int) int {
	if len(matrix) == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}
	return 0
}

func cal2(matrix [][]int, x, y int) [][]int {
	var arr [][]int
	for i := 1; i < len(matrix); i++ {
		if i != x {
			var arr2 []int
			for j := 0; j < len(matrix[i]); j++ {
				if j != y {
					arr2 = append(arr2, matrix[i][j])
				}
			}
			arr = append(arr, arr2)
		}
	}
	fmt.Println(arr, x, y)
	return arr
}

func TestFileIo(t *testing.T) {
	f := &os.File{}
	w := bufio.NewWriter(f)
	a, s := w.WriteString("qwe \n")
	t.Log(a, s)

	w.Flush()
	f.Close()

}

func TestRanddd(t *testing.T) {
	a, er := randomString(32)
	t.Log(a, er)
}

// 生成随机字符串
func randomString(length int) (string, error) {
	rb := make([]byte, length)
	_, err := crand.Read(rb)
	if err != nil {
		return "", err
	}
	fmt.Println(base64.URLEncoding.EncodeToString(rb), rb)
	return base64.URLEncoding.EncodeToString(rb), nil
}

func TestPoolNew(t *testing.T) {
	// disable GC so we can control when it happens.
	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	runtime.GOMAXPROCS(1)

	i := 0
	p := sync.Pool{
		New: func() interface{} {
			i++
			return i
		},
	}
	if v := p.Get(); v != 1 {
		t.Fatalf("got %v; want 1", v)
	}
	if v := p.Get(); v != 2 {
		t.Fatalf("got %v; want 2", v)
	}
	p.Put(33)
	t.Log(p.Get())
	p.Put(44)
	runtime.GC()
	t.Log(p.Get())

	// Make sure that the goroutine doesn't migrate to another P
	// between Put and Get calls.
	p.Put(42)
	if v := p.Get(); v != 42 {
		t.Fatalf("got %v; want 42", v)
	}

	if v := p.Get(); v != 3 {
		t.Fatalf("got %v; want 3", v)
	}
}

func TestDeferPrint(t *testing.T) {

	a, b := &w2{q: 1}, &w2{q: 1}
	defer func() {
		printInt(10, a, printInt(100, a, b))
	}()
	a.q = 2
	defer func() {
		printInt(20, a, printInt(200, a, b))
	}()
}

func printInt(index int, a, b *w2) *w2 {
	println(index, a.q, b.q, a.q+b.q)
	a.q += b.q
	return a
}

// 查看结构体对象大小
func TestSizeOf(t *testing.T) {
	type T struct {
		// _ [0]atomic.Int64
		x int32
	}

	type A struct {
		T
		v int32
	}
	t.Log(unsafe.Sizeof(A{}))
	var ss string
	t.Log(unsafe.Sizeof(ss))

	type B struct {
		// // _ [0]atomic.Int64
		a int8  // 4
		b int16 // 4
		c int32 // 4
	}
	type C struct {
		// _ [0]atomic.Int64
		x int32 // 4
		// v int64
	}
	t.Log(unsafe.Sizeof(B{}))
	i := int32(1)
	t.Log(unsafe.Sizeof(i))
	t.Log(unsafe.Sizeof(C{}))

	t.Log(unsafe.Sizeof(int64(1000000)))
	t.Log(unsafe.Sizeof(int32(100)))
	t.Log(unsafe.Sizeof(int16(100)))
	t.Log(unsafe.Sizeof(int8(100)))
	t.Log(unsafe.Sizeof(100))
	t.Log(unsafe.Sizeof(string("1")))

}

func TestArranges(t *testing.T) {
	aa := []int64{1, 23, 45}

	t.Log(aa[0:3])

	l := 2
	for len(aa) != 0 {
		if len(aa) < l {
			l = len(aa)
		}
		t.Log(aa)
		aa = aa[l:]
	}
}

func TestArrslice(t *testing.T) {
	mm := make(map[int]int, 100)
	mm[1] = 1
	t.Log(len(mm))
	t.Logf("%p", &mm)
	t.Log(reflect.TypeOf(mm))
	t.Log(len(mm))

	for i := 0; i < 1000; i++ {
		mm[i] = i
	}
	t.Logf("%p", &mm)
	t.Log(len(mm))

	arr := [10]int{}
	arr[6] = 1
	arr[4] = 11
	t.Log(arr)
	m := make(map[int][10]int)
	m[1] = arr
	m[0] = arr
	t.Log(m)
	arr[2] = 1
	m[3] = arr
	t.Log(m)

}

func TestAssertions(t *testing.T) {
	// v := []int64{1, 2, 3}
	// if ht, ok := tr.(*http.Transport); ok {
	// 	stdreq = ht.Request()
	// }
	// a, ok := v.([]uint64)

	m := make(map[int64]uint64)
	for i := int64(-1110); i < 820000; i++ {
		// m[i] = uint64(i + 199)
		m[i] = uint64(i)
	}
	t.Log("over", len(m))

}

func TestStructTransfer(t *testing.T) {
	a := &Fn{}
	a.A = "3"
	a.C = "3"
	type b interface {
		mark1()
		mark2()
	}
	var m []b
	m = append(m, a)
	t.Logf("%+v", m[0])
	for _, b2 := range m {
		b2.mark1()
		t.Log(a)
		b2.mark2()
		t.Log(a)
		c, err := b2.(*Fn)
		d, err := b2.(interface{})

		t.Log(err, d)
		t.Log(err, c)
	}
}
func (a *Fn) mark1() {
	a.A = "1"
}
func (a *Fn) mark2() {
	a.A = "2"
	a.B = "2"
}

func TestHash(t *testing.T) {

	for _, TestString := range [][]int64{
		nil,
		{1000, 1001},
		{1000, 1101},
		{1001, 1101},
		{2, 1, 23, 124, 143, 4152, 412, 2, 1},
		{2, 1, 23, 124, 144, 4152, 412, 2, 1},
		{2, 1, 23, 4152, 412, 2, 1},
		{2, 1, 23, 124, 14, 4152, 412, 2, 22, 2, 2, 22},
	} {

		// 先 序列化 在md5 下 在转int64
		bs, _ := json.Marshal(TestString)

		Md5Inst := md5.New()
		Md5Inst.Write(bs)
		Result := Md5Inst.Sum(nil)
		hStr := fmt.Sprintf("%x", Result)
		t.Log(TestString, hStr)

		i := new(big.Int)
		i.SetBytes([]byte(hStr)) // octal
		fmt.Println(i)

		fmt.Println("******:", i.Int64())
	}

}

func TestDeepCopy(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	deepCopy := func(x, y interface{}) {
		b, err := json.Marshal(x)
		if err != nil {
			t.Log(err)
		}
		t.Log(string(b), b, x)
		fmt.Printf("%p \n", x)
		fmt.Printf("%p \n", b)
		err = json.Unmarshal(b, y)
		fmt.Printf("%p \n", y)
		if err != nil {
			t.Log(err)
		}
	}
	m := make([]int, 0)
	fmt.Printf("## %p \n", m)

	deepCopy(a, &m)
	fmt.Printf("## %p \n", m)
	t.Log(m)

}

// 测试闭包方法 可以让方法内部的变量不受外部的影响
func TestClosure(t *testing.T) {
	last := time.Unix(1657761812, 0).Format("2006-01-02 15:04:05")
	defer func(a string) {
		t.Log(1, a)
	}(last)

	last = "1"

	defer func(a *string) {
		t.Log(2, *a)
	}(&last)
	last = "2"

	defer func() {
		t.Log(3, last)
	}()
	last = "3"
	t.Log(last)
	return
}

func TestSliceInfo(t *testing.T) {

	s := fmt.Sprintf("%s 已经 %d 岁了。 \n ", "asd", 212)
	// t.Log(fmt.Sprintf("%.2f 已经 %d 岁了。 \n ", "22.222%", 212))
	t.Log(strings.TrimSuffix("22.3322%", "%"))
	io.WriteString(os.Stdout, s)

	sliceint := make([]int64, 45000)                                                     // 指向元素类型为int32的1000个元素的数组的切片
	fmt.Println("Size of []int32:", unsafe.Sizeof(sliceint))                             // 24
	fmt.Println("Size of [1000]int32:", unsafe.Sizeof([1000]int64{}))                    // 4000
	fmt.Println("Real size of s:", unsafe.Sizeof(sliceint)+unsafe.Sizeof([1000]int64{})) // 4024
}

func TestFileSecrie(t *testing.T) {

	err := backupSecret("./tmp/", "sss.conf",
		&PayWay{Id: 123,
			Name: "qqq"})
	if err != nil {
		t.Error(err)
	}

	cf := &PayWay{}
	err = loadExSecret(cf, "./tmp/sss.conf")
	if err != nil {
		t.Error(err)
	}

}

func loadExSecret(conf interface{}, confName string) error {
	data, err := os.ReadFile(confName)
	if err != nil {
		return err
	}
	raw, err := crypto.AesDecode(string(data))
	raws := trimComments([]byte(raw))
	if err = json.Unmarshal(raws, conf); err != nil {
		fmt.Printf("Parse conf %v failed: %v", string(raws), err)
		return err
	}
	return nil
}

func backupSecret(RestDir, filename string, conf *PayWay) error {
	confBytes, err := jsoniter.MarshalIndent(conf, "", "    ")
	if err != nil {
		return fmt.Errorf("runner config %v marshal failed, err is %v", conf, err)
	}
	// 判断默认备份文件夹是否存在，不存在就尝试创建
	if _, err := os.Stat(RestDir); err != nil {
		if os.IsNotExist(err) {
			if err = os.Mkdir(RestDir, 0755); err != nil && !os.IsExist(err) {
				return fmt.Errorf("rest default dir not exists and make dir failed, err is %v", err)
			}
		}
	}

	// 配置文件加密
	cpted, err := crypto.AesEncode(string(confBytes))
	if err != nil {
		return fmt.Errorf("runner config %v crypto failed, err is %v", conf, err)
	}

	return os.WriteFile(RestDir+filename, []byte(cpted), 0644)
}

func TestStringToByte(t *testing.T) {
	var r io.Reader

	// r = &bytes.Buffer{}
	r = bytes.NewBufferString("#$12fddd11")
	// r = bytes.NewBufferString("sdasdqw")
	fmt.Println(r)

	buf := make([]byte, 4)
	n, err := r.Read(buf)
	t.Log(n, err, buf)

	if err != nil {
		panic(err)
	}
	size := binary.BigEndian.Uint32(buf)
	if size == 0 {
		t.Error("nil!!!", size)
	}
	data := make([]byte, size)
	a, err := r.Read(data)
	t.Log(err, a, size, len(data))

}

func TestSelectToGo(t *testing.T) {

	a := 2

dodo:
	println(0)

	switch {
	case a < 5:
		println(5)

	case 1 == a:
		println(1)
	case 2 == a:
		a = 1
		println(2)
		goto dodo
	default:
	}

}

func TestMapDel(t *testing.T) {
	var p map[string]string
	p = nil
	delete(p, "a")

	ii := int64(123_2123)
	println(ii)
}

// 洗牌算法
func TestRandRange(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	t.Log(math.Round(922337203685477580))
	t.Log(math.Round(9223372036854775807.9111))
	t.Log(math.Round(9223372036854775807.1111))
	for i := 0; i < len(arr)-1; i++ {
		t.Log(rand.Intn(10))

		rand.Seed(time.Now().UnixNano())
		a := rand.Intn(len(arr) - i)
		arr[i], arr[a] = arr[a], arr[i]
	}

	t.Log(arr)
}

func TestUrlname(t *testing.T) {

	name := "%E6%98%A5%E5%A4%A9%E5%8F%AF%E7%9C%9F%E6%98%AF%E4%B8%AA%E5%B0%8F%E8%AE%A8%E5%8E%8C%E9%AC%BC%EF%BC%8C%E5%9C%A8%E6%88%91%E5%BF%83%E9%87%8C%E5%81%B7%E5%81%B7%E5%85%BB%E4%BA%86%E4%B8%80%E5%8F%AA%E5%B0%8F%E9%B9%BF%EF%BC%8C%E5%B0%B1%E6%92%92%E6%89%8B%E4%B8%8D%E7%AE%A1%E4%BA%86%E2%9C%A8%E2%9C%A8%E2%9C%A8"
	// escapeUrl := urls.QueryEscape(name)

	t.Log(urls.QueryUnescape(name))
	name = "darren is  iphone"
	t.Log(urls.QueryUnescape(name))

}

func TestBuInterface(t *testing.T) {
	arr := []Buiding{House{}, Shop{}, Toilet{}}

	for _, v := range arr {
		v.Builds()
	}
}

func TestTickerFor(t *testing.T) {

	tt := time.NewTicker(time.Second * 1)
	i := 1
	for ; true; <-tt.C {
		fmt.Println("tick", i)
		i++
	}

}

func TestByte(t *testing.T) {
	ss := "dHJ1ZQ=="
	ss = "d"
	var vv interface{}
	err := json.Unmarshal([]byte(ss), &vv)
	t.Log(vv, err)
}
func TestJsonByte(t *testing.T) {
	ss := NewA()
	b, _ := json.Marshal(ss)
	t.Log(b, string(b))
	t.Log(b, *(*string)(unsafe.Pointer(&b)))
	b = nil
	t.Log(b, 22, *(*string)(unsafe.Pointer(&b)))

}

// -gcflags '-m -l'
func TestAddPoint(t *testing.T) {
	println(addV1(1, 2))
	println(addV2(1, 2))
}
func addV2(a, b int) (add *int) {
	add = new(int)
	*add = a + b
	return add
}
func addV1(a, b int) *int {
	add := 0
	add = a + b
	return &add
}

func TestGoPanic(t *testing.T) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				t.Log("panic 11111", err)
			}
		}()
		i := 1
		for i < 10 {
			i++
			if i == 4 {
				panic(2)
			}
			time.Sleep(time.Second)
		}
		t.Log("end 1")

	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				t.Log("panic 99999", err)
			}
		}()
		i := 1
		for i < 5 {
			i++
			t.Log(i)
			time.Sleep(time.Second)
		}
		t.Log("end 2")
	}()
	ch2 := make(chan int)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				t.Log("panic 2222", err)
			}
		}()
		i := 1
		for {
			if i == 4 {
				panic("2222")
			}
			ch2 <- i
			i++
			t.Log(222)
			time.Sleep(time.Second * 2)

		}
	}()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				t.Log("panic 3333", err)
			}
		}()
		i := 1
		for {
			t.Log(<-ch2, i)
		}
	}()

	av := func() {
		t.Log(1)
		go func() {
			defer func() {
				if err := recover(); err != nil {
					t.Log("panic 4444", err)
				}
			}()
			panic(1)
		}()
		t.Log(2)

	}

	go av()
	time.Sleep(time.Hour * 10)

}

func TestTimeDaysAdd(t *testing.T) {
	st, _ := time.Parse("2006-01-02 15:04:05", "2024-06-19 14:30:52")
	et, _ := time.Parse("2006-01-02 15:04:05", "2024-06-19 14:34:52")
	ct, _ := time.Parse("2006-01-02 15:04:05", "2024-06-19 14:33:52")

	var zt time.Time
	t.Log(zt.IsZero())

	t.Log(ct.Before(et))
	t.Log(ct.After(st))
	t.Log(time.Now().String())
	t.Log(time.Now().Year())
	t.Log(time.Now().GoString())
	now := time.Now().Unix()
	t.Log(now, now+3600*24)
	tt := time.Unix(1697839585, 0)
	t.Log(tt.Format("2006-01-02 15:04:05"))
	t.Log(tt.Day())
	ts := time.Now()
	tm1 := time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, ts.Location())
	// tm2 := tm1.AddDate(0, 0, 1)
	t.Log(tm1, tm1.Unix())

	// 7天前0点
	day7 := ts.AddDate(0, 0, -7)
	before7Day := strconv.FormatInt(time.Date(day7.Year(), day7.Month(), day7.Day(), 0, 0, 0, 0, ts.Location()).Unix(), 10)
	t.Log(before7Day)

}

func TestSleepSpeed(t *testing.T) {
	t.Log("sleep speed test")
	start := time.Now()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 300)
		// time.Sleep(time.Millisecond * 2)
	}
	t.Log(time.Now().Sub(start))
}

func TestContains(t *testing.T) {
	// t.Log(strings.Contains(strings.ToLower("shenZhou"), "shenzhou"))
	// t.Log(strings.Contains("darren91231231i1892", "darren"))
	t.Log(getLatestIpAddr("127.0.0.1, 115.171.170.95, 10.5.12.212"))
	str := "\n     \nHello, World!\n  \n"
	trimmedStr := strings.TrimSpace(str)
	fmt.Println("原字符串:", str)
	fmt.Println("原字符串:", trimmedStr)
}

func getLatestIpAddr(clientIP string) string {
	if index := strings.LastIndex(clientIP, ","); index >= 0 {
		clientIP = clientIP[index+1:]
	}
	clientIP = strings.TrimSpace(clientIP)
	if len(clientIP) > 0 {
		return clientIP
	}
	return ""
}

func (b arrStruct) Len() int {
	return len(b)
}

func (b arrStruct) Less(i, j int) bool {
	if b[i].Bar < b[j].Bar {
		return true
	}

	if b[i].Bar == b[j].Bar {
		return b[i].Foo > b[j].Foo
	}

	return false
}

func (b arrStruct) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func TestSortSlice(t *testing.T) {
	s := []Baz{
		{5, 4},
		{6, 7},
		{2, 3},
		{6, 4},
	}

	fmt.Printf("%+v\n", s)
	sort.Sort(arrStruct(s))
	fmt.Printf("%+v\n", s)
}

func TestWriteSlice(t *testing.T) {
	app := make([]int64, 0, 1000)
	var lock sync.RWMutex
	var lock2 sync.RWMutex

	a := 2
	wg.Add(1)

	go func(a *int) {
		lock2.RLock()
		time.Sleep(time.Second * 5)
		*a = 3
		lock2.RUnlock()
		println("成功释放")
		wg.Done()
	}(&a)

	for i := 0; i < 10; i++ {
		go func(i int) {
			lock2.Lock()
			a = i
			println(a)
			lock2.Unlock()
		}(i)
	}
	// 以上例子表明 当读锁时候 ，写锁会被阻塞
	wg.Wait()
	wg.Add(1)
	go func() {
		lock2.Lock()
		a = 1
		println("写锁中", a, time.Now().Unix())
		time.Sleep(time.Second * 5)
		lock2.Unlock()
		wg.Done()
	}()

	for i := 0; i < 10; i++ {
		go func(i int) {
			lock2.RLock()
			println("读取变量：", i, time.Now().Unix())
			lock2.RUnlock()

		}(a)
	}
	// 以上例子表明 当写锁时候 ，重复多次读锁会被阻塞 ，直到写锁释放
	wg.Wait()

	wg.Add(1)
	for i := 0; i < 10; i++ {
		go func(i int) {
			lock2.RLock()
			println("没有写锁时：读取变量：", i, time.Now().Unix())
			lock2.RUnlock()
		}(i)
	}
	// 以上例子表明 当没有写锁时候 ，可以重复多次读锁
	wg.Wait()

	go func() {
		i := 0
		for {
			lock.Lock()
			app = append(app, rand.Int63())
			lock.Unlock()
			if i%10000000 == 0 {
				t.Log(len(app))
			}
			i++
		}
	}()
	wg.Wait()

	go func() {
		tic := time.NewTicker(100 * time.Microsecond)
		for {
			select {
			case <-tic.C:
				println(len(app), 99999999)
				lock.Lock()
				app = make([]int64, 0, 1000)
				lock.Unlock()

				println(len(app), 9888888888)

			}
		}
	}()
	wg.Wait()

}

func TestChanCtx(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	ctx = context.WithValue(ctx, "key", "darren")
	wgsync := new(sync.Once)
	handle2(ctx, time.Second*1, wgsync)
	go handle(ctx, time.Second*1)
	t.Log(ctx.Value("key"))
	t.Log(ctx.Err())
	t.Log(time.Now().Format("2006-01-02 15:04:05"))
	cancel()
	select {
	case <-time.After(time.Second * 10):
		fmt.Println("over")
	}
}

// 实现一个在执行的任务，超时时必须结束，停止任务
func handle2(ctx context.Context, duration time.Duration, wgsync *sync.Once) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("handle2  case <-ctx.Done():  with", duration)
			return
		case <-time.After(time.Millisecond):
			fmt.Println("handle2  log  with", duration)
			wgsync.Do(func() {
				time.Sleep(time.Second * 40)
				fmt.Println("handle2  sssssssssss  with", duration)
			})
			fmt.Println("handle2  log end  with", duration)
		default:

		}
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func TestBoolSize(t *testing.T) {
	aa := make([]bool, 0, 10)
	aa = append(aa, true)
	aa = append(aa, true)
	println(len(aa), cap(aa), unsafe.Sizeof(aa), &aa)
	aa = aa[0:1]
	println(len(aa), cap(aa), unsafe.Sizeof(aa), &aa)
	aa = aa[0:2]
	println(len(aa), cap(aa), unsafe.Sizeof(aa), &aa)

	var b = make(map[int64]bool, 0)
	b[1] = true
	t.Log(b[2], b[1])
	if unsafe.Sizeof(b) != 1 {
		t.Errorf("bool size is %d, want 1", unsafe.Sizeof(b))
	}
}

type CommBody interface {
	GetA() string
	keys() []string
}

func (fn Fn) GetA() string {
	return fn.A
}
func (fn Fn) keys() []string {
	return []string{"A", "B", "C", "D"}
}

func DefClass(param CommBody) {
	ms, _ := json.Marshal(param)
	fmt.Println(string(ms))

}

func TestClassDef(t *testing.T) {
	var fn = Fn{A: "a", B: "b", C: "c", D: "d"}
	DefClass(fn)
}

func BenchmarkDirConcat(b *testing.B) {
	var s37 = []byte{36: 'x'} // len(s37) == 37
	var str string
	for i := 0; i < b.N; i++ {
		str = string(s37) + string(s37)
	}
	_ = str
}

func BenchmarkSplitedConcat(b *testing.B) {
	var s37 = []byte{36: 'x'} // len(s37) == 37
	var str string

	for i := 0; i < b.N; i++ {
		str = string(s37[:32]) +
			string(s37[32:]) +
			string(s37[:32]) +
			string(s37[32:])
	}
	_ = str
}

func TestFloatTostring(t *testing.T) {
	f := float64(23.434532)
	t.Logf("%f", f)
	t.Logf("%.2f", f)
}

func TestCompareInterface(t *testing.T) {
	var x interface{} = []int{1, 2}
	var y interface{} = map[string]int{"aa": 1, "bb": 1}
	var z interface{} = func() {}

	// The lines all print false.
	println(x == y)
	println(x == z)
	println(x == nil)
	t.Log(x, y, z)

	// Each of these line could produce a panic.
	// println(x == x)
	// println(y == y)
	// println(z == z)
}

func TestSortInt(t *testing.T) {

	arr := []int{1, 44, 2, 77, 3, 4, 5}

	s := sort.SearchInts(arr, 22)
	t.Log(s, arr)

}

func TestSliss(t *testing.T) {
	arr := []int{1}

	go func() {
		for i := 100; i > 0; i-- {
			println(len(arr))
			arr = append(arr, 2)
		}
	}()
	time.Sleep(time.Second)
	for len(arr) > 0 {
		l := 10
		if l > len(arr) {
			l = len(arr)
		}
		_ = arr[:l]
		t.Log(len(arr), l)
		arr = arr[l:]

	}

	t.Log(arr)

}

func TestAppends(t *testing.T) {

	arr := []int{}
	ch := make(chan struct{}, 0)
	isClose := false
	i := int32(0)
	maxCount := int32(5000)
	println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	go func(i int32) {
		for {
			if i > maxCount {
				ch <- struct{}{}
				return
			}
			atomic.AddInt32(&i, 1)
			arr = append(arr, int(i))
			if len(arr)%100 == 0 {
				fmt.Println(i)
			}
		}
	}(i)
	go func(i int32) {
		for {
			if i > maxCount {
				ch <- struct{}{}
				return
			}
			atomic.AddInt32(&i, 1)
			arr = append(arr, int(i))
			if len(arr)%100 == 0 {
				fmt.Println(i)
			}
		}
	}(i)
	go func(i int32) {
		for {
			if i > maxCount {
				ch <- struct{}{}
				return
			}
			atomic.AddInt32(&i, 1)
			arr = append(arr, int(i))
			if len(arr)%100 == 0 {
				fmt.Println(i)
			}
		}
	}(i)
	for {
		select {
		case <-ch:
			isClose = true
		default:
			if len(arr) > 20 {
				_ = arr[:20]
				// fmt.Println(ss)
				arr = arr[20:]
			}

		}

		if isClose && len(arr) == 0 {
			return
		} else if isClose && len(arr) < 20 {
			ss := arr[:]
			fmt.Println("***", ss)
			arr = arr[len(arr):]
		}

	}

}

func TestWgAdd(t *testing.T) {
	var wgs sync.WaitGroup

	k := uint64(1)
	md := make([]int, 1000)

	wgs.Add(1)

	// for i := int(1); i < 1000; i++ {
	//	md[i] = i
	//	wgs.Add(1)
	//	go func(i int) {
	//		md[int(k)] = int(i)
	//		atomic.AddUint64(&k, 1)
	//		wgs.Done()
	//
	//	}(i)
	// }
	wgs.Done()
	// wgs.Done()
	// wgs.Done()
	time.Sleep(time.Second)
	wgs.Wait()
	t.Log(md, len(md))
	md = md[:k]
	t.Log(md, len(md))
}

func TestMapRead(t *testing.T) {
	m := map[string]int{}
	m1 := map[string]int{}
	go func() {
		for {
			m1["key1"] = 1
			m = m1
		}
	}()
	for {
		_ = m["key2"]
	}
}

func TestParams(t *testing.T) {

	B := Fn{}
	B.A = "qqqq"
	P(B)
	B.B = "qqqq"
	t.Log(B.A, B)

}
func P(f Fn) {
	f.A = "aaa"
	f.B = "aaa"

}

func TestDeferSort(t *testing.T) {
	x := 1
	y := AddD(&x)
	println(*y, x)

	x1 := 1
	y1 := AddE(&x1)
	println(x1, y1)

}
func AddD(a *int) *int {

	println(a, *a, 111)
	defer func() {
		*a = *a + 1
	}()
	println(a, *a, 333)

	return a
}

func AddE(a *int) *int {

	println(a, *a, 111)
	defer func() {
		*a = *a + 1
	}()
	println(a, *a, 333)

	return nil
}

func TestSliceEqual(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 3, 2, 4}
	c := []int{1, 2, 3, 4}
	d := []byte{1, 2, 3, 4}
	fmt.Println(reflect.DeepEqual(a, b))
	fmt.Println(reflect.DeepEqual(a, c))
	fmt.Println(reflect.DeepEqual(a, d))

}

func TestStructNil(t *testing.T) {
	D := new(PayWay)
	t.Log(D.Test)
	t.Log(D.Test2)
	t.Log(D.Test.S)
	t.Log(D.Test.S.Alias)
	// 指针 就是nil  会panic！
	if D.Test2 != nil {
		t.Log(D.Test2.Alias)
	}

}

func trimZeroDecimal(s string) string {
	var foundZero bool
	flag := 0
	for i := len(s); i > 0; i-- {
		switch s[i-1] {
		case '.':
			if foundZero {
				return s[:i-1]
			}
		case '0':
			foundZero = true
		default:

			flag = i
		}
	}
	return s[:flag]
}

func TestFloat32To64(t *testing.T) {

	f := float64(44.532424234234)
	// sss := "44.532424234234"

	// t.Log(fmt.Sprintf("%.f", sss))
	t.Log(float32(f))

	f = float64(445324242342.34)
	t.Log(float32(f))

	f = float64(117.11166743741192)
	t.Log(float32(f))

	f = float64(99999.532424234234)
	t.Log(float32(f))
}

func TestSliceSplit(t *testing.T) {

	aa := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	t.Log(len(aa), cap(aa)) // ss := aa[2:3000]
	// 表明了只亲空长度，但是原有容量保存
	aa = aa[:0]
	t.Log(len(aa), cap(aa)) // ss := aa[2:3000]
	ss := aa[2:2]

	var name string
	name = "asSSd"
	aaaa := strings.Count(name, "")

	t.Log(strings.ToUpper(name))
	t.Log(strings.ToUpper(" 免费-iOS-底部-1 "))
	t.Log(aaaa)
	t.Log(ss, 2222)
}

func TestAaa(t *testing.T) {
	// // target := 8
	// target := 6
	// // target := 1
	// // arrStruct := []int{1, 2, 3, 3, 3, 3, 3, 3, 9, 10}
	// arrStruct := []int{5, 7, 7, 8, 8, 10}
	// // arrStruct := []int{1}
	// fmt.Println(searchRange(arrStruct, target))
	// beginningOfTime := time.Unix(time.Now().Unix(), 0)
	beginningOfTime := time.Unix(99999123123, 0)
	fmt.Println(beginningOfTime.Unix())
}

const url = "https://github.com/EDDYCJY"

// 持续监听消息，在5秒内没有新消息，则认为消息接收超时 5秒内有新消息，则重新计时
func TestForRange(t *testing.T) {
	boxLine := make(chan string)
	// 在一个新的goroutine中模拟接收消息
	go func() {
		time.Sleep(3 * time.Second) // 模拟消息接收的延迟
		boxLine <- "Message 1"
		t.Log("Message 1")
		time.Sleep(2 * time.Second)
		boxLine <- "Message 2"
		t.Log("Message 2")
		time.Sleep(7 * time.Second) // 超过5秒没有新消息，模拟超时情况
		boxLine <- "Message 3"
		t.Log("Message 3")
	}()
	t.Log(time.Now())
	timer := time.NewTimer(time.Second * 5)
	cs := make(chan struct{})
	go func() {
		defer func() {
			if r := recover(); r != nil {
				t.Error("Recovered in f", r)
			}
		}()
		for {
			select {
			case aas, ok := <-boxLine:
				t.Log(aas, ok, time.Now())
				time.Sleep(time.Second * 10)
				timer.Reset(time.Second * 5)
			case <-timer.C:
				t.Log(time.Now(), "###")
				cs <- struct{}{}
				return
			}
		}
	}()
	<-cs

}

func TestTimeMicr(t *testing.T) {
	t.Log(strconv.FormatInt(time.Now().Unix(), 10))
	t.Log(time.Now().Unix() * 1000)
}

func TestSpilt(t *testing.T) {
	msg := "aaskjdhakshdlkhsada审"
	t.Log(len(strings.Split(msg, "审批于")))
	t.Log(strings.Contains("15618802115 18658852500", "18658852500"))

}

func TestPanicV4(t *testing.T) {

	type R struct {
		S *int64
		K string
	}
	t.Logf("%+v", errors.New("123"))
	w := int64(2)
	aa := R{
		S: &w,
		K: "123",
	}
	if aa.S == nil || *aa.S == 0 {
		println(123)
	}
	aa.S = nil
	bb, err := aa, errors.New("123")
	if err == nil || bb.K == "123" {
		println("ppppp")
	}

}

func TestRangeNil(t *testing.T) {
	obj := make([]string, 0)
	obj = nil
	// obj = append(obj, "123")
	for v := range obj {
		println(v)
	}
}

func TestMapv2(t *testing.T) {
	m := make(map[string][]int)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%p---%v\n", m, m)
	m["test"] = s
	fmt.Printf("%p---%v\n", s, s)
	fmt.Printf("%p---%v\n", m["test"], m["test"])
	fmt.Printf("%p---%v\n", m, m)
	clear(m)
	fmt.Printf("%p---%v\n", m, m)
	m["sss"] = s
	fmt.Printf("%p---%v\n", m, m)

}

func TestSliceV2(t *testing.T) {
	s := make([]int, 1)
	s[0], s, s[0] = 333, []int{1, 2, 3}, 222
	t.Log(s)
}

func TestFnLoop(t *testing.T) {

	aa := new(Fn)
	aa.Geta()
	aa.Getb()
	aa.Getc()
	aa.Getd()
	aa.Geta()

}
func (f *Fn) Geta() string {
	return f.A
}

func (f *Fn) Getb() string {
	return f.B
}
func (f *Fn) Getc() string {
	return f.C
}
func (f *Fn) Getd() string {
	return f.D
}

func TestNilFun(t *testing.T) {

	a := NewA()
	c := context.TODO()
	fmt.Println(a.GetName(&c, "222"))

}

type A struct {
	name  string
	Alias string
}

func NewA() *A {
	return &A{
		name: "111",
	}
}

func (a *A) GetName(ctx *context.Context, name string) string {
	a.name = name
	return a.name
}

func WhichIsBest() int {
	a, b, c, d := rand.Intn(10), rand.Intn(10), rand.Intn(10), 0
	switch {
	case a == 1:
		d = 1
	case b == 1:
		d = 2
	case c == 1:
		d = 3
	default:
		d = 4
	}
	return d
}

func WhichIsBestV2() int {
	a, b, c, d := rand.Intn(10), rand.Intn(10), rand.Intn(10), 0
	switch {
	case a == 1:
		d = 1
		return d
	case b == 1:
		d = 2
		return d
	case c == 1:
		d = 3
		return d
	}
	return d
}

type SR string

func TestPoint(t *testing.T) {
	var vv = SR("初始值")
	d := vv
	d.myVal()

	d.Get1()
	d.myVal()

	vv = "2次"
	d = vv
	d.Get2()
	d.myVal()

	d.Get4()
	d.myVal()

	d.Get3()
	d.myVal()
}

func (s *SR) Get1() {
	// s = nil
	// s.myVal()
}

func (s SR) Get2() {
	s = SR("期望的值2")
	s.myVal()

}
func (s *SR) Get3() {
	v := SR("期望的值3")
	s = &v
	s.myVal()
}
func (s *SR) Get4() {
	v := SR("期望的值4")
	*s = v
	s.myVal()
}

func (s *SR) myVal() {
	fmt.Printf("the val : %p %s \n", s, *s)
}

func TestAffirmation(t *testing.T) {
	var a = uint8(90)
	println(int64(a))
	println(int8(a))
	var m interface{}
	m = a

	if s, ok := m.(int64); !ok {
		println(s)
	}

}

func TestSwitchs(t *testing.T) {
	aa := structures.Interval{Start: 123, End: 333}
	switch {
	case aa.End == 333:
		println(2)
	case aa.Start == 23:
		println(3)
	case aa.End == 123:
		println(4)
	default:
		println(5)

	}
}

func regexMatch(regex string, operation string) bool {
	r, err := regexp.Compile(regex)
	if err != nil {
		return false
	}

	return r.FindString(operation) == operation
}

// log := `[20241008-14:43:53]#[EVENT  ]#[normal_opera]#tid[978102]#log_id[002300100036]# { "sz_market_producer":{"sz_session1":[ {"producer":"hot_0", "tick_msg_count":324875618, "tick_msg_discard":0,"tick_msg_disorder":0 }, {"producer":"hot_1", "tick_msg_count":324875619,"tick_msg_discard":0,"tick_msg_disorder":0 }]}}`
//
//	re := regexp.MustCompile(`$begin:math:display$(\\d{8}-\\d{2}:\\d{2}:\\d{2})$end:math:display$.*?#(.*)$`)

//
//	matches := re.FindStringSubmatch(log)
//	if len(matches) > 2 {
//		fmt.Println("Date:", matches[1])
//		fmt.Println("JSON:", matches[2])
//	} else {
//		fmt.Println("No match found")
//	}

func TestRegex(t *testing.T) {
	// println(regexMatch(`$begin:math:display$(\\d{8}-\\d{2}:\\d{2}:\\d{2})$end:math:display$.*?#(.*)$`, "login"))

	// println(regexMatch(`^(\[\d{8}-\d{2}:\d{2}:\d{2}\])#.*?#(.*)$`, "login"))
	// log := `[20241008-14:43:53]#[EVENT  ]#[normal_opera]#tid[978102]#log_id[002300100036]# { "sz_market_producer":{"sz_session1":[ {"producer":"hot_0", "tick_msg_count":324875618, "tick_msg_discard":0,"tick_msg_disorder":0 }, {"producer":"hot_1", "tick_msg_count":324875619,"tick_msg_discard":0,"tick_msg_disorder":0 }]}}`
	//
	// re := regexp.MustCompile(`$begin:math:display$(\\d{8}-\\d{2}:\\d{2}:\\d{2})$end:math:display$.*?#(.*)$`)
	//
	// matches := re.FindStringSubmatch(log)
	// if len(matches) > 2 {
	// 	fmt.Println("Date:", matches[1])
	// 	fmt.Println("JSON:", matches[2])
	// } else {
	// 	fmt.Println("No match found")
	// }

	logLines := []string{
		`[20241008-14:30:20]#[EVENT  ]#[normal_opera]#tid[978102]#log_id[002300100034]# { "sz_market":[ {"session":"sz_session1", "session_msg_count":326252314, "session_tick_msg_count":309270624, "session_tick_msg_discard":0,"session_tick_msg_disorder":0, "session_opt_msg_count":0, "msg_queue[0]_remain_len":16777216, "msg_queue[1]_remain_len":16777216 },{"mddp_fix":"sz_session1_mddp_fix_1_0", "ok":4, "timeout":0, "cacheover":0, "q_max":111, "empty_max":0}] }`,
		`[20241008-14:30:20][20241008-23:30:20]#[EVENT  ]#[normal_opera]#tid[978102]#log_id[002300100036]# { "sz_market_producer":{"sz_session1":[ {"producer":"hot_0", "tick_msg_count":309270624, "tick_msg_discard":0,"tick_msg_disorder":0 }, {"producer":"hot_1", "tick_msg_count":309270624, "tick_msg_discard":0,"tick_msg_disorder":0 }]}}`,
	}

	// 正则表达式提取日期部分
	// datePattern := `$begin:math:display$(\\d{8}-\\d{2}:\\d{2}:\\d{2})$end:math:display$`
	// datePattern := `\[(.*?)\]`
	datePattern := `\[(.*?)\]`
	// datePattern := `\[([^\]]+)\].*?\[([^\]]+)\]`
	// datePattern := `$begin:math:display$(.*?)$end:math:display$`
	// 正则表达式提取 JSON 内容
	jsonPattern := `\[(.*?)\].*?#\s*({.*})`
	// jsonPattern := `#\s*({.*})`

	dateRegex := regexp.MustCompile(datePattern)
	jsonRegex := regexp.MustCompile(jsonPattern)

	for _, logs := range logLines {
		dateMatch := dateRegex.FindStringSubmatch(logs)
		jsonMatch := jsonRegex.FindStringSubmatch(logs)

		if len(dateMatch) > 1 {
			fmt.Println("Date:", dateMatch[1])
		}
		if len(jsonMatch) > 1 {
			mapdata := make(map[string]any)
			_ = json.Unmarshal([]byte(jsonMatch[1]), &mapdata)
			fmt.Println("JSON:", mapdata)
			// fmt.Println("JSON:", jsonMatch[1])
		}
	}

	println(regexMatch(`^.*login.*$`, "asdkalogin/hsj"))
	println(regexMatch(`^.*login.*$`, "1qweqwi"))
	println(regexMatch(`.*2014.*$`, "1qwe[2014]qwi"))
	println(regexMatch(`^.*TenantSso/Login/.*$`, "/helloworld/aaa?asdasdjk"))
}

func TestArrayGroup(t *testing.T) {
	// 	原来 arrStruct[ "qwe","weq","wqe","abc","cba"]
	// 	期望 arrStruct[["qwe","weq","wqe"],["abc","cba"]]

	arr := []string{"qwe", "weq", "wqe", "abc", "cba"}

	// pp := make(map[int32][][26]int, 10)

	println(arr, 'a'-97)

	for _, v := range arr {
		sum := int32(0)
		m := [26]int{}
		for _, vv := range v {
			sum += vv
			m[vv-97]++
		}

	}

	// fmt.Println(m, pp)
	// p := make(map[byte][]map[byte]bool, 0)
	//
	// for _, v := range arrStruct {
	// 	m := make(map[byte]bool, 0)
	// 	var s byte
	// 	for k := range v {
	// 		m[v[k]] = true
	// 		s += v[k]
	// 	}
	// 	if vv, ok := p[s]; ok {
	// 		vv = append(vv, m)
	// 		p[s] = vv
	// 	} else {
	// 		p[s] = []map[byte]bool{m}
	// 		println(s)
	// 	}
	// }
	// fmt.Println(p)

}

// 当前用例
//
//	期望： 1,8，2，7，3，6,4，5
func TestReverseV3(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	list := structures.Ints2List(arr)
	f := list
	s := list
	// 快慢指针取到中点
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	// 翻转慢 list
	mm := structures.Reverse(s)
	aaa := &structures.ListNode{}
	res := aaa
	i := 0
	// 交替 append 到新链表中，直到完成
	for list != nil && mm != nil {
		if i&1 == 0 {
			res.Next = &structures.ListNode{Val: list.Val}
			list = list.Next
		} else {
			res.Next = &structures.ListNode{Val: mm.Val}
			mm = mm.Next
		}
		res = res.Next
		i++
	}
	structures.Travel(aaa.Next)
}
func TestZigzagLevelOrderV2(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	tree := structures.Ints2TreeNode(arr)

	stack := []*structures.TreeNode{tree}

	aas := make([][]int, 0)
	isSS := false
	level := 0
	for len(stack) != 0 {
		i := 0

		if len(aas) <= level {
			aas = append(aas, []int{})
		}
		l := len(stack)
		for l > i {
			tt := stack[0]
			if !isSS {
				aas[level] = append(aas[level], tt.Val)
			} else {
				aas[level] = append([]int{tt.Val}, aas[level]...)
			}

			stack = stack[1:]
			if tt.Left != nil {
				stack = append(stack, tt.Left)
			}
			if tt.Right != nil {
				stack = append(stack, tt.Right)
			}
			i++
		}
		if isSS {
			isSS = false
		} else {
			isSS = true
		}
		level++
	}
	fmt.Println(aas)
}

func TestArraySum(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 0, 7}
	B := []int{6, 7, 0}

	la := len(A) - 1
	lb := len(B) - 1
	if la > lb {

	}
	x := 0
	for k := range A {
		if lb-k >= 0 {
			A[la-k] += B[lb-k]
		}
		A[la-k] = A[lb-k] + x
		x = 0
		if A[la-k] >= 10 {
			A[la-k] %= 10
			x = 1
		}
	}
	if x > 0 {
		for i := la; i <= lb; i++ {
			B[lb-i] += x
			x = 0
			if B[lb-i] >= 10 {
				B[lb-i] %= 10
				x = 1
			}
		}
	}

	if x > 0 {

	}
}

func TestPractice(t *testing.T) {
	// HeapSoft([]int{1, 888, 11, 2, 44, 3, 777, 4, 55, 5, 67})
	singleNumber([]int{1, 2, 3, 4, 3, 2, 4, 1})
}
func singleNumber(nums []int) int {
	// bit:=
	a := 0
	for v := range nums {
		a ^= nums[v]
	}
	println(a)
	return a
}

func TestXruntime(t *testing.T) {
	fmt.Println(runtime.GOMAXPROCS(2))
	go func() {
		for {
			fmt.Println("NumGoroutine", runtime.NumGoroutine())
			time.Sleep(time.Second * 1)
		}
	}()

	for i := 0; i < 10000; i++ {
		go timeSleep10(i)
	}
	time.Sleep(time.Second * 41)
}

func timeSleep10(i int) {
	fmt.Println(i)
	time.Sleep(time.Second * time.Duration(i+5))
	fmt.Println("ok", i)
}

func TestXRuntime(t *testing.T) {
	// runtime.Gosched()
	// 切换任务
	fmt.Println("cpus:", runtime.GOMAXPROCS(0))          // 返回当前系统的CPU核数量
	fmt.Println("cpus:", runtime.NumCPU())               // 返回当前系统的CPU核数量
	fmt.Println("goroot:", runtime.GOROOT())             //
	fmt.Println("NumGoroutine:", runtime.NumGoroutine()) // 返回真该执行和排队的任务总数
	fmt.Println("archive:", runtime.GOOS)                // 目标操作系统
}

func smallestDistancePair(nums []int, k int) int {
	keys := make(map[int]int, 0)
	arr := make([]int, 0)
	l := len(nums)
	for i := range nums {
		for j := i + 1; j < l; j++ {
			diff := nums[j] - nums[i]
			if diff < 0 {
				diff = ^diff + 1
			}
			keys[diff]++
			// arrStruct = mergeappend(arrStruct, diff)

		}
	}
	fmt.Println(arr, k, keys)
	for m := range keys {
		arr = mergeappend(arr, m)
	}

	if len(arr) > 0 {

	}

	return 0
}

func mergeappend(arr []int, r int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] >= r {
			arr = append(arr[:i], append([]int{r}, arr[i:]...)...)
			return arr
		}
	}
	arr = append(arr, r)

	return arr
}

func TestCase(t *testing.T) {

	x := 24
	a := 14
	b := -10
	// s:=b^b
	println(a|b, x^a^b, ^b+1)
	fmt.Printf("%b \n", x)
	fmt.Printf("%b \n", a)
	fmt.Printf("%b \n", b)
	fmt.Printf("%b \n", a&b)

}

func TestTwoSum(t *testing.T) {
	res := twoSum([]int{20, 70, 20, 150}, 220)
	t.Log(res)
}

func TestMaxLeng(t *testing.T) {
	// t.Log(maxLength([]int{2, 3, 4, 1, 5}))
	// t.Log(maxLength([]int{2, 2, 3, 4, 1, 5}))
	t.Log(maxLength([]int{1, 1, 1, 2, 2, 3, 4, 5, 6, 7, 7, 8, 9}))
}

func search(nums []int, target int) int {
	// write code here
	l := len(nums) - 1
	i := 0
	mid := 0
	for i <= l {
		mid = int(uint(i+l) >> 1)
		// fmt.Println(mid, i, l)
		if nums[mid] == target {
			for mid > 1 && nums[mid-1] == target {
				mid--
			}
			return mid
		}
		if nums[mid] < target {
			i = mid + 1
		} else {
			l = mid - 1

		}
	}

	return -1
}

func FibonacciV2(n int) int {

	if n > 40 || n < 0 {
		return n
	}

	arr := [40]int{}
	arr[0] = 0
	arr[1] = 1

	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]

}
func Fibonacci(n int) int {
	if n >= 2 {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}
	return 0
	// write code here
}

// 二分查找
func TestSearch(t *testing.T) {
	arr := []int{1, 2, 3, 3, 3, 3, 5, 6, 7, 8, 9, 9, 9, 9, 99}
	target := 9
	low, fast := 0, len(arr)-1
	for low <= fast {
		mid := len(arr) - (fast-low)>>1
		if target > arr[mid] {
			low = mid
		} else if target < arr[mid] {
			fast = mid
		} else {
			for mid < len(arr)-2 && arr[mid+1] == target {
				mid++
			}
			println(arr[mid], mid)
			break
		}
	}
}

func isValid(s string) bool {
	mp := make(map[uint8]uint8, 3)
	mp['['] = ']'
	mp['{'] = '}'
	mp['('] = ')'
	stack := make([]uint8, 0)

	for i := range s {
		if v, ok := mp[s[i]]; ok {
			stack = append(stack, v)
		} else {
			if len(stack) > 0 && stack[len(stack)-1] == s[i] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
		fmt.Println(stack)
	}

	return len(stack) == 0
}
func maxLength(arr []int) int {
	long, i, r := 0, 0, 0
	l := len(arr)
	as := [256]byte{}
	for i < l {
		if r == l {
			return long
		}
		if as[arr[r]] == 0 {
			as[arr[r]]++
			r++
		} else {
			as[arr[i]]--
			i++
		}
		long = max(long, r-i)
		fmt.Println(i, r, long)
	}

	return long
}
func twoSum(numbers []int, target int) []int {

	l := len(numbers)
	for k := range numbers {
		for j := 1; j < l; j++ {
			if k != j && target == numbers[k]+numbers[j] {
				return []int{k + 1, j + 1}
			}
		}

	}
	return []int{}

	// write code here
}

func TestMergeArr2(t *testing.T) {
	arr1 := []int{1, 3, 3, 3, 4, 4, 5, 7, 9}
	arr2 := []int{2, 4}
	arr3 := append(arr1, make([]int, len(arr2))...) // 扩展 arr1 的长度
	merge(arr3, len(arr1), arr2, len(arr2))
}

// 归并排序 不用额外空间，改变原来数组
func merge(A []int, m int, B []int, n int) {
	var a = m - 1
	var b = n - 1
	var i int
	for i = m + n - 1; a >= 0 && b >= 0; i-- {
		if A[a] >= B[b] {
			A[i] = A[a]
			a--
		} else {
			A[i] = B[b]
			b--
		}
	}
	if a < 0 {
		for ; i >= 0; i-- {
			A[i] = B[b]
			b--
		}
	}
	fmt.Println(A)
}

// 最小路径和
func minPathSum(matrix [][]int) int {
	n := len(matrix)
	m := len(matrix[0])

	dp := make([][]int, n)
	for k := range matrix {
		if dp[k] == nil {
			dp[k] = make([]int, m)
			dp[0][0] = matrix[0][0]
		}
		if k < 1 {
			continue
		}
		dp[k][0] = matrix[k][0] + dp[k-1][0]
	}

	for k := range matrix[0] {
		if k > 0 {
			dp[0][k] = matrix[0][k] + dp[0][k-1]
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + matrix[i][j]
		}
	}
	fmt.Println(matrix)
	fmt.Println(dp)
	return dp[n-1][m-1]
}

// 最小路径和 用原来数组不需要创建
func minPathSumV2(matrix [][]int) int {
	n := len(matrix)
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == 0 {
				if i > 0 {
					matrix[i][j] += matrix[i-1][j]
				}
				continue
			}
			if i == 0 {
				matrix[i][j] += matrix[i][j-1]
				continue
			}
			matrix[i][j] = min(matrix[i-1][j], matrix[i][j-1]) + matrix[i][j]
		}
	}
	fmt.Println(matrix)
	return matrix[n-1][m-1]
}

func TestLeetcode(t *testing.T) {

	// k := getLongestPalindrome("ab1234321abcvbnmmnbvcba1", 24)
	// k := minPathSum([][]int{[]int{1, 3, 5, 9}, []int{8, 1, 3, 4}, []int{5, 0, 6, 1}, []int{8, 8, 4, 0}})
	k := minPathSumV2([][]int{{1, 3, 5, 9}, {8, 1, 3, 4}, {5, 0, 6, 1}, {8, 8, 4, 0}})
	// k := minPathSum([][]int{[]int{1, 1, 5, 9}, []int{8, 1, 3, 4}, []int{5, 0, 6, 1}, []int{8, 1, 1, 0}})
	// k := minPathSumV2([][]int{[]int{1, 1, 5, 9}, []int{8, 1, 3, 4}, []int{5, 0, 6, 1}, []int{8, 1, 1, 0}})

	println(k)
}
func getLongestPalindrome(A string, n int) int {

	k := 0
	for i := 0; i < n; i++ {
		// 两种情况： 一种是 aba  一种是：aa 所以用 2 个 for 循环
		for j := 0; i-j >= 0 && j+i < n; j++ {
			if A[i-j] != A[i+j] {
				break
			}
			k = max(k, 2*j+1)

		}
		for j := 0; i-j >= 0 && j+i+1 < n; j++ {
			if A[i-j] != A[i+j+1] {
				break
			}
			k = max(k, 2*(j+1))
		}
	}
	return max(k, 0)
}

func TestChanV2(t *testing.T) {

	ob := &S{
		cl:     make(chan struct{}),
		notity: make(chan int, 1),
	}
	ob.wg.Add(1)

	// 写线程
	go func() {
		ob.wg.Add(1)

		i := 0
		for {
			select {
			case <-ob.cl:
				ob.wg.Done()
				return
			case ob.notity <- i:

			}
			ob.num++
			i++
			time.Sleep(time.Second)
		}
	}()

	// 写线程
	go func() {
		ob.wg.Add(1)

		i := 500

		for {
			select {
			case <-ob.cl:
				ob.wg.Done()
				return
			case ob.notity <- i:

			}
			ob.num--
			i--
			time.Sleep(time.Second)
		}
	}()

	// 写线程
	go func() {
		ob.wg.Add(1)

		i := 200

		for {
			select {
			case <-ob.cl:
				ob.wg.Done()
				fmt.Println("推出 chan")
				return
			case ob.notity <- i:

			}
			ob.num--
			i--
			time.Sleep(time.Second)
		}
	}()

	// 读线程
	go func() {
		for v := range ob.notity {
			fmt.Println("读取", v, ob.num)
		}
		// for {
		// 	select {
		// 	case v, ok := <-ob.notity:
		// 		if ok {
		// 			fmt.Println(v, ob.num)
		// 		} else {
		// 			fmt.Println(999)
		// 			return
		// 		}
		// 	case <-ob.cl:
		// 		return
		// 	}
		// }
	}()

	go func() {
		time.Sleep(time.Second * 10)
		ob.wg.Done()
		close(ob.cl)
	}()

	println(1222)
	ob.wg.Wait()
	println(333)

	time.Sleep(time.Hour)
	return

}

// 堆排序小练
func TestHeapSort(t *testing.T) {
	// arrStruct := []int{4, 3, 5, 1, 2, 6, 7}
	arr := []int{1, 4, 3, 2, 6, 5, 8, 7, 9, 0}
	// arr := []int{1, 2}
	// arr := []int{5, 4}
	// fmt.Println(arr)

	heapSort(arr)
	// BuildHeap(arr, len(arr))
	fmt.Println(arr)
}

func heapSort(arr []int) {
	n := len(arr)

	// 构建大顶堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify22(arr, n, i)
	}

	// 一个个将堆顶元素交换到末尾，再重新构建堆
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // 交换最大值到末尾
		heapify22(arr, i, 0)            // 重新构建堆
	}
}

// 维护堆结构
func heapify22(arr []int, n, i int) {
	fmt.Println(arr)
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify22(arr, n, largest)
	}
}

func TestBlocking(t *testing.T) {
	ch := make(chan struct{})

	// var x interface{} = nil
	// var y *int = nil
	// interfaceIsNil(x)
	// interfaceIsNil(y)

	aa1 := "aaa" + "222你好"
	var aa2 strings.Builder
	aa2.WriteString(aa1)
	aa2.WriteString("24444")
	fmt.Println(aa2.String())
	go func() {
		time.Sleep(time.Hour)
		ch <- struct{}{}
	}()
	<-ch

}

// 无缓冲 buf chan
func TestChanNoBuf(t *testing.T) {

	ch := make(chan int)
	timeout := make(chan struct{})
	go func() {
		i := 0
		for {
			i++
			select {
			case <-timeout:
				return
			default:
				ch <- i
			}
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		i := 1
		for {
			aa, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println(aa, ok, i)
			i++
			if i == 3 {
				timeout <- struct{}{}
				close(ch)
				return
			}

		}
	}()

	ww := sync.WaitGroup{}
	// ww.Add(10)
	go func() {
		for {
			time.Sleep(10 * time.Second)
			ww.Done()
		}
	}()

	ww.Wait()

}

func TestBfs(t *testing.T) {
	// 初始化树
	tree := structures.Ints2TreeNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	// BfsTree(tree)
	// DfsTreeV2(tree)
	// DfsTree(tree)
	BfsTree(tree)
}

// 深度遍历
func DfsTree(tree *structures.TreeNode) {
	if tree == nil {
		return
	}
	fmt.Println(tree.Val)
	if tree.Left != nil {
		DfsTree(tree.Left)
	}
	if tree.Right != nil {
		DfsTree(tree.Right)
	}
}

// 广度遍历 队列实现
func BfsTree(tree *structures.TreeNode) {
	if tree == nil {
		return
	}
	var node []*structures.TreeNode
	node = []*structures.TreeNode{tree}
	for len(node) != 0 {

		t := node[0]
		fmt.Println(t.Val, len(node))
		node = node[1:]

		l := t.Left
		if l != nil {
			node = append(node, l)
		}
		r := t.Right
		if r != nil {
			node = append(node, r)
		}
	}
}

func TestWeiyi(t *testing.T) {
	// 000011
	aa := 20
	// 0000100
	bb := 4
	t.Log(aa & (aa - 1))
	t.Log(aa >> 1)
	t.Log(aa << 10)
	// 10亿
	t.Log(1 << 30)
	// 取相反数
	t.Log(^99)
	// 异位或
	t.Log(15 | 20)
	t.Log(99 | 91)
	// 判断奇偶
	// 异位与
	t.Log(bb & 1)
	//
	t.Log(aa ^ bb)
	t.Log(aa | bb)

}

func TestArrEq(t *testing.T) {
	aa := []byte{1, 2, 3}
	bb := []byte{1, 2, 3}
	cc := []byte{1, 3, 2}
	dd := []int{1, 3, 2}

	println(bytes.Equal(aa, bb))
	println(reflect.DeepEqual(aa, cc))
	println(reflect.DeepEqual(dd, cc))
	println(reflect.DeepEqual(aa, bb))
}

func TestSliceRange(t *testing.T) {
	aa := []*PayWay{}
	n := int64(123)
	for i := 0; i < 10; i++ {
		aa = append(aa, &PayWay{
			Id:  n * int64(i+1),
			Ids: n * int64(i+1),
		})
	}

	for k, v := range aa {
		// v := v
		go func() {
			t.Log(k, v.Ids)
		}()
	}
	time.Sleep(time.Second * 3)

}

func TestZhengzebiaoda(t *testing.T) {
	// text := "[08-14 09:02:08.187535][IN1FO 1231       ][rl_mode123l_agent]  123   "
	// mach := `^\[(\d{2}-\d{2}\s*\d{2}:\d{2}:\d{2}\.\d{6})\]`
	// re, _ := regexp.Compile(mach)
	// t.Log(re.MatchString(text))
	text := "2025/03/20 14:36:16 [2839218] .d..t...... Data/Model/20250320/asdasd2025/03/20 11:36:16 [2839218] .d..t...... Data/Mod2025/03/20 14:36:16 [2839218] el/20250320/"
	mach := `(\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2}) `
	re, _ := regexp.Compile(mach)
	aas := re.FindAllStringIndex(text, -1)
	t.Log(aas)
	t.Log(re.MatchString(text))
}

// param: days 为多少天以后
// return: 今天+days 天之后的日期,所在月的最后一天, 按"2006年01月02日"格式化
func LastDateOfMonth(days int, ct time.Time) string {
	d := ct.AddDate(0, 0, days)              // time.Now()可以换成支持测试环境调时间的方法
	firstDate := d.AddDate(0, 0, -d.Day()+1) // 当月的第一天
	lastDate := firstDate.AddDate(0, 2, -1)
	// lastDate.Unix()
	// 当月的最后一天
	return lastDate.Format("2006年01月02日")
}

type PayWay struct {
	//    支付id
	Id  int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Ids int64 `protobuf:"varint,2,opt,name=id,proto3" json:"ids,omitempty"`
	// 支付名称
	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Test  B
	Test2 *A
}

type B struct {
	S A
}

type S struct {
	cl     chan struct{}
	num    int
	notity chan int
	wg     sync.WaitGroup
	sync.Mutex
}
