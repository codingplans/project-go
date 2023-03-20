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
	"github.com/tidwall/gjson"
	"io"
	"log"
	"math"
	"math/big"
	"math/rand"
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
	"testing"
	"time"
	"unsafe"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
	"testgo/modgo/crypto"

	"github.com/Darrenzzy/person-go/structures"
	jsoniter "github.com/json-iterator/go"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.uber.org/goleak"
)

type w2 struct {
	q int
}

type baz struct {
	bar int
	foo int
}

type baz2 struct {
	bar int
	foo int
	fzz []int
}
type arrStruct []baz

func TestGoCtx(t *testing.T) {
	var l chan struct{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	ctx = context.WithValue(ctx, "key", "value")
	go func() {
		time.Sleep(4 * time.Second)
		GetA(ctx, "1")
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ctxx, cancel := context.WithTimeout(ctx, 2*time.Second)
		defer cancel()
		GetA(ctxx, "2")
	}()
	<-l

}
func GetA(ctx context.Context, f string) {
	log.Println(f, ctx.Value("key"))
	select {
	case <-ctx.Done():
		println("done")
	}
}

func RecoverGO(f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("%+v\n", errors.Errorf("%+v", r))
			}
		}()
		f()
	}()
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
	o.InfoLogger = new(foo)
	o.InfoLogger.Infof("test")
}

type foo struct {
}

func (*foo) Infof(s string, any ...interface{}) {
	println(s)

}

func TestSliceContains(t *testing.T) {
	sql := "11,3222,33,2233"
	sqls := strings.Split(sql, ",")
	t.Log(slices.Contains(sqls, "22"))
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
	// sleepHour := 1
	ts := time.Now()
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
	a := &baz2{bar: 1, foo: 1, fzz: []int{1, 2, 3}}
	b := &baz2{bar: 3}
	*b = *a
	a.bar = 2
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

	if ok1 != nil {
		return
	}
	v, ok := f.Float64()
	t.Log(v, ok, ok1)
}

func TestSliceFunc(t *testing.T) {

	// Fn := func(arr []int) {
	//
	// }

	m := lo.Map[int64, string]([]int64{1, 2, 3, 4}, func(x int64, index int) string {
		println(index)
		return strconv.FormatInt(x+1, 10)
	})
	t.Log(len(m), m)
}

func TestStructIsNil(t *testing.T) {
	for i := 0; i < 10; i++ {
		var foo []*baz
		t.Log(foo == nil)
		foo = append(foo, &baz{bar: 1, foo: 1})
		t.Log(foo == nil)
	}
}

func TestGoGroutines(t *testing.T) {
	arr := make(arrStruct, 0, 1000)
	for i := 1; i < 1000; i++ {
		arr = append(arr, baz{bar: i, foo: i})
	}

	fn := func(k int, v *baz, list arrStruct) {
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
	Bar int `json:"bar"`
	Foo int `json:"foo"`
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
	js := `{"a":"2","b":"2","bar":1,"foo":10}`
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
		x   int32
		v   int32
		vv  int32
		vvv int32
	}
	type C struct {
		// _ [0]atomic.Int64
		x int32
		// v int64
	}

	t.Log(unsafe.Sizeof(B{}))
	i := int32(1)
	t.Log(unsafe.Sizeof(i))
	t.Log(unsafe.Sizeof(C{}))
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

func TestBateInt(t *testing.T) {
	t.Log(unsafe.Sizeof(int64(1000000)))
	t.Log(unsafe.Sizeof(int32(100)))
	t.Log(unsafe.Sizeof(int16(100)))
	t.Log(unsafe.Sizeof(int8(100)))
	t.Log(unsafe.Sizeof(100))
	t.Log(unsafe.Sizeof(string("1")))
	// 	 tmain_test.go: 8
	//    tmain_test.go: 4
	//    tmain_test.go: 2
	//    tmain_test.go: 1
	//    tmain_test.go: 8
	//    tmain_test.go: 16
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

func TestParseInt(t *testing.T) {
	s := "1048576"
	i, err := strconv.Atoi(s)
	t.Log(i, err)
	ii, err := strconv.ParseInt(s, 10, 128)
	t.Log(ii, err)
	ii, err = strconv.ParseInt(s, 5, 64)
	t.Log(ii, err)

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
	t.Log(time.Now().String())
	t.Log(time.Now().Year())
	t.Log(time.Now().GoString())
	now := time.Now().Unix()
	t.Log(now, now+3600*24)
	tt := time.Unix(1697839585, 0)
	t.Log(tt)
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
	if b[i].bar < b[j].bar {
		return true
	}

	if b[i].bar == b[j].bar {
		return b[i].foo > b[j].foo
	}

	return false
}

func (b arrStruct) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func TestSortSlice(t *testing.T) {
	s := []baz{
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

func Example_Print() {
	score := []int{1, 2, 3}
	fmt.Println(score)
	// Output: [1 2 3]
}

func Test_structChan(t *testing.T) {

	var v BigBar
	m := make(chan struct{})
	t.Log(unsafe.Sizeof(v))
	// 大于2M的变量 协程会不执行:
	// fatal error: newproc: function arguments too large for new goroutine
	// go func(a BigBar) {
	// 	_ = a
	// 	m <- struct{}{}
	// }(v)

	// // 指针传递 不会造成异常
	// go func(a *BigBar) {
	// 	_ = a
	// 	m <- struct{}{}
	// }(&v)

	go func() {
		println(1232)
		time.Sleep(time.Minute * 1)
		<-m
		println(123)
	}()
	// // 正常参数传递 没有大小限制
	// task3 := func(a BigBar) (err error) {
	// 	_ = a
	// 	m <- struct{}{}
	// 	return nil
	// }
	// t.Log(task3(v))

	// 协程参数变量尽量用指针
	<-m
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
	isClose = false
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

	for i := int(1); i < 1000; i++ {
		md[i] = i
		wgs.Add(1)
		go func(i int) {
			md[int(k)] = int(i)
			atomic.AddUint64(&k, 1)
			wgs.Done()

		}(i)
	}

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

func TestFloat32To64(t *testing.T) {

	f := float64(44.532424234234)
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

	// ss := aa[2:3000]
	ss := aa[2:2]

	var name string
	name = "asSSd"
	aaaa := strings.Count(name, "")

	t.Log(strings.ToUpper(name))
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

func TestForRange(t *testing.T) {

	i := 0
	for {
		t.Log(time.Now())
		timer := time.NewTimer(time.Second)
		go func() {
			defer func() {
				if r := recover(); r != nil {
					t.Error("Recovered in f", r)
				}
			}()
			select {
			case <-timer.C:
				t.Log(time.Now(), "###", i)
				i++
			}

			if i > 500 {
				panic("444")
			}
		}()

	}

}

func TestTimeMicr(t *testing.T) {
	t.Log(strconv.FormatInt(time.Now().Unix(), 10))
	t.Log(time.Now().Unix() * 1000)
}

func TestAdd(t *testing.T) {
	s := Add(url)
	if s == "" {
		t.Errorf("Test.Add error!")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(url)
	}
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

func TestRegex(t *testing.T) {

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

func HeapSoft(arr []int) {
	l := len(arr)
	fmt.Println(arr)

	for i := l / 2; i >= 0; i-- {
		BuildHeapV2(arr, i, l)
	}

	l--
	for i := l; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		BuildHeapV2(arr, 0, l)
		l--
	}
	fmt.Println(arr)
}

func BuildHeapV2(arr []int, n, lens int) {
	k := n
	for n < lens {
		i := n*2 + 1
		j := i + 1
		if i < lens && arr[k] < arr[i] {
			k = i
		}
		if j < lens && arr[k] < arr[j] {
			k = j
		}
		if k != n {
			arr[k], arr[n] = arr[n], arr[k]
			n = k
		} else {
			// n = n * 2
			break
		}
		// println(n, j, i)
	}

}

func xRuntime() {
	runtime.Gosched()                                    // 切换任务
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

func TestIsValid(t *testing.T) {
	// t.Log(isValid("[](){}"))
	// t.Log(isValid("[]({}[]{{{}}}){}"))
	// t.Log(isValid("{[}]"))
	// t.Log(isValid("]"))
	// t.Log(Fibonacci(4))
	// t.Log(FibonacciV2(4))
	// t.Log(Fibonacci(10))
	// t.Log(FibonacciV2(30))
	// t.Log(search([]int{1, 2, 2, 3, 3, 6, 8, 8, 8, 9, 9, 9}, 6))
	// t.Log(search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6))
	// t.Log(search([]int{1, 1, 1, 1, 6, 6, 6, 6, 6, 7, 8, 9, 10, 10, 101}, 101))
	// t.Log(search([]int{-2, 1, 2}, -2))
	t.Log(LRU([][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}, 3))
}

type Lru struct {
	md   map[int]*node
	buf  int
	max  int
	head *node
	tail *node
}

type node struct {
	pre, next *node
	val, key  int
}

func LRU(operators [][]int, k int) []int {
	// write code here
	lru := initLru(k)
	res := []int{}
	for i := range operators {
		if operators[i][0] == 1 {
			lru.set(operators[i][1], operators[i][2])

		} else {
			res = append(res, lru.get(operators[i][1]))
		}
	}
	return res
}
func initLru(k int) Lru {
	return Lru{
		md:  make(map[int]*node),
		max: k,
	}
}

func (this *Lru) get(k int) int {
	if v, ok := this.md[k]; ok {
		this.remove(v)
		this.add(v)
		return v.val
	}
	return -1
}
func (this *Lru) set(k, x int) {
	if v, ok := this.md[k]; ok {
		this.remove(v)
		this.add(v)
		return
	} else {
		n := &node{val: x, key: k}
		this.md[k] = n
		this.add(n)

	}

	if len(this.md) > this.max {
		delete(this.md, this.tail.key)
		this.remove(this.tail)
	}

}

func (this *Lru) remove(n *node) {
	if this.head == n {
		this.head = n.next
		this.head.pre = nil
		return
	}

	if this.tail == n {
		this.tail = n.pre
		n.pre.next = nil
		n.pre = nil
		return

	}

	n.pre.next = n.next
	n.next.pre = n.pre

	return

}

func (this *Lru) add(n *node) {
	n.next = this.head
	if this.head != nil {
		this.head.pre = n
	}
	this.head = n
	if this.tail == nil {
		this.tail = n
		this.tail.next = nil
	}

	return
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 小练归并排序
func TestMergeArr(t *testing.T) {
	arr1 := []int{1, 3, 5, 7, 9}
	arr2 := []int{2, 4, 6, 8, 10}
	i, j := 0, 0
	for j < len(arr2) && i < len(arr1) {
		if arr1[i] <= arr2[j] {
			i++
		} else {
			arr1 = append(arr1[:i], append([]int{arr2[j]}, arr1[i:]...)...)
			i++
			j++
		}
	}
	if j < len(arr2) {
		arr1 = append(arr1, arr2[j])
		j++
	}

	fmt.Println(arr1)
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

func TestQuickSoft(t *testing.T) {
	// arr := []int{4, 3, 5, 1, 2, 6, 33, 12, 1, 55, 3, 2, 111, 57, 7, 5}
	// arr := []int{2, 9, 3, 333, 8, 11, 4, 7, 6, 5, 22, 115, 3}
	arr := []int{3, 5, 1, 9, 8, 344, 1, 5555, 44, 2}

	fmt.Println(arr)
	QuickSoft(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// 堆排序小练
func TestHeapSort(t *testing.T) {
	// arrStruct := []int{4, 3, 5, 1, 2, 6, 7}
	arr := []int{1, 4, 3, 2, 6, 5, 8, 7, 9, 0}

	fmt.Println(arr)

	BuildHeap(arr, len(arr))
	fmt.Println(arr)
}

// 随便练一下 二叉树排序 =》堆排序
func Test2TreeSoft(t *testing.T) {
	// arrStruct := []int{1, 2, 3, 4, 5, 6, 7, 8}
	arr := []int{4, 3, 5, 1, 2, 6, 7}

	ts := []int{0, 0, 0, 1}
	ss := copy(ts, arr)
	fmt.Println(len(ts), cap(ts), arr, ts, ss)
	fmt.Printf("%v,%p,%p,", ss, ts, arr)

	node := &structures.TreeNode{Val: 4}
	for v := range arr[:6] {
		node = CreateTree(node, arr[v+1])
	}
	fmt.Println("begin")
	Travel(node)
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
	DfsTreeV2(tree)
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

// 深度遍历 压栈处理
func DfsTreeV2(tree *structures.TreeNode) {

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
	t.Helper()
	aa := []*PayWay{}

	aa = append(aa, &PayWay{
		Id:  123,
		Ids: 123,
	})
	aa = append(aa, &PayWay{
		Id:  222,
		Ids: 222,
	})
	aa = append(aa, &PayWay{
		Id:  333,
		Ids: 333,
	})

	for k, v := range aa {
		fmt.Println(v, k)
	}

}

func TestPanicdefer(t *testing.T) {
	t.Log(DeferTest())
}

func DeferTest() int {
	a := 1
	b := 2
	defer calc(a, b, "1")
	// defer calc(a, calc(a, b, "0"), "1")
	// a = 0
	// defer calc(a, calc(a, b, "3"), "2")
	return a + 1
}

func calc(x, y int, s string) int {
	fmt.Println(s)
	fmt.Println(x, y, x+y)
	return x + y
}

func TestZhengzebiaoda(t *testing.T) {
	text := "fff${LastDateOfMonth(3)}ffff aa2021年02月30日aaa${LastDateOfMonth(123)}aaa     "
	mach := "\\$\\{LastDateOfMonth.([0-9]+.)\\}"
	re, _ := regexp.Compile(mach)

	// 取出所有符合规则日期
	list := re.FindAllString(text, -1)
	re1, _ := regexp.Compile("[0-9]+")
	t.Log("替换前：", text, "\n")

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
	t.Log("替换后：", text, "\n")
	ts := time.Now()
	tm1 := time.Date(ts.Year(), ts.Month(), ts.Day()+1, 0, 0, 0, 0, ts.Location())
	tm2 := tm1.AddDate(0, 0, 1)
	t.Log(tm1, tm2)

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
