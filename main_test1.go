package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"gopkg.in/olivere/elastic.v5"
	"iceberg/frame"
	"iceberg/frame/icelog"
	"iceberg/frame/protocol"
	"laoyuegou.com/http_api"
	"laoyuegou.com/util"
	"laoyuegou.pb/gameserver/pb"
	"laoyuegou.pb/gameserver/pb/gobang"
	"laoyuegou.pb/imapi/pb"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

const (
	BLACK_STONE    = int(1)  // 黑
	WHITE_STONE    = int(2)  // 白
	BOARD_SIZE     = int(15) // 棋盘
	BOARD_SIZE_jun = int(4)  // 棋盘
	EMPTY_STONE    = int(0)  //
	BOARD_SIZE1    = int(4)

	BLACK_FIVE = int(1)
	WHITE_FIVE = int(2)
)

func Testslice(T *testing.T) {

	arr := [50]int64{12, 12123, 23, 123, 1, 24, 124, 124, 14, 1234, 23, 4235, 23523}

	userIds := ""
	for _, userId := range arr {
		userIds += fmt.Sprintf("%d,", userId)
	}
	userIds = userIds[0 : len(userIds)-1]
	log.Printf(userIds)
}

func Testintstring(T *testing.T) {
	var ss int64
	ss = 111
	icelog.Infof("%+v", strconv.FormatInt(ss, 10))
}
func Testmapappend(T *testing.T) {

	aa := elastic.GeoPointFromLatLon(1, 2)

	log.Printf("%+v", aa)
}

func Testerror(T *testing.T) {
	var err error
	icelog.Info(err)
	err = fmt.Errorf("%v", 123123)
	icelog.Info(err.Error())
}

type Person struct {
	age int
}

func TestLists(T *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[3:4:5]
	fmt.Println(cap(s))
}

func TestMaps(T *testing.T) {
	banGame := make(map[int64]int64)

	banGame = nil
	if _, ok := banGame[123]; ok {
		fmt.Println("ok")
	} else {
		fmt.Println("not!!!")

	}

}

func TestPanic(T *testing.T) {
	var v int
	defer panic(333)
	if v == 0 {
		panic("123")
	}
}

func TestName(t *testing.T) {
	var s string
	s = "Aabcdefg"

	fmt.Printf("%d , %f", s[0], s[0])
}

func main() {
	TestName(nil)
	// TestPanic()
	// TestMaps()
	// TestLists()
	// Testerror()
	// Testmapappend()
	// Testintstring()
	// Testslice()
	// counts()
	// Test2(&Testing.T{})
	// Testint()
	// channel()
	// CompletePayOrder(234, 1896)
	// https()
	// Testarr()
	// chanTest()

	// gotoFun()
	// TestFor()
	// aaaa := boo()
	// fmt.Printf("%+v,%d ", aaaa, len(aaaa))

	// mapinit()
	// timess()
	// timeseco¡nd()
	// mapTest()
	// fmt.Print(1/2 + 1)
	// fmt.Print(2 / 2)
	// fmt.Print(1 % 2)
	// fmt.Print(2 % 2)
	// slices()
	// varsss()¡
	// arr := map[int]int32{}
	// arr[1] = 1
	// arr[2] = 1
	//
	// if _, ok := arr[12]; ok {
	// 	fmt.Printf("%3d", arr)
	// }
	//
	// for i := 0; i < 20; i++ {
	// 	fmt.Println(rand.Int31n(3) + 1)
	// }

	// tt := JSONTime{time.Now()}
	// arrTest()

	// Shuffle()
	// golang()

	// listcon()
	// t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	// timeStr := time.Now().Format("2006-01-02")
	// t2, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	// second := t2.AddDate(0, 0, 1).Unix() - time.Now().Unix()
	// fmt.Println(timeStr, t, t2)
	// fmt.Println(t.Unix() + 1)
	// fmt.Println(t2.AddDate(0, 0, 1).Unix() - time.Now().Unix())

	// return
	//
	// mapt()
	//
	// userinfo()
	// Cmd(frame.TODO())
	// return
	// // mm := new(gameserver.Room)
	// mm.Black = 2
	// mm.White = 4
	// if mm.InRoom == nil {
	// 	mm.InRoom = make(map[int64]int)
	// 	mm.InRoom[mm.Black] = 1
	//
	// }
	// if mm.InRoom[4] != 2 {
	// 	fmt.Println(3444)
	//
	// }
	//
	// fmt.Println(mm.InRoom, mm)
	// return
	//
	// var aa []*http.Server
	// if aa != nil && len(aa) >= 2 {
	// 	fmt.Println(11)
	// }
	// fmt.Println(aa)
	//
	// return
	// for i := 0; i < 2; i++ {
	// 	rand.Seed(int64(time.Now().UnixNano()))
	// 	fmt.Println(rand.Intn(2))
	// }
	// return

	// var i,j int32
	// arr := make([][]int32,0,15)

	// for i = 1; i <= 10; i++ {
	// 	for j = 1; j <= 10; j++ {
	// 		arr = append(arr,[]int32{i,j})
	// 	}
	// }
	// fmt.Println(arr,arrToRepeat(arr))

	// return
	// var m int
	// m = -1
	// fmt.Println(m)
	// // for m < 5{
	// // 	fmt.Print(m,&m)
	// // 	m++
	// // }
	//
	// return
	// for _, j := range Board {
	// 	fmt.Println(j)
	//
	// }
}

const a = 10 << iota

const (
	b = 10 << iota * iota
	c
	d
	e
)

// const (
// 	// c = 10
// 	// d
// 	// d = iota
// 	// e
// 	f = "hello"
// 	// nothing
// 	g
// 	h = iota
// 	i
// 	j = 0
// 	k
// 	l, m = iota, iota
// 	n, o
//
// 	p = iota + 1
// 	q
// 	_
// 	r = iota * iota
// 	s
// 	t = r
// 	u
// 	v = 1 << iota
// 	w
// 	x         = iota * 0.01
// 	y float32 = iota * 0.01
// 	z
// )

var Board [BOARD_SIZE + 2][BOARD_SIZE + 2]int

// func counts() {
// 	print(a, "\n", b, "\n", c, "\n", d, "\n", e, "\n", f, "\n", g, "\n", h, "\n", i, "\n", j, "\n", k, "\n", l, "\n", m, "\n", n, "\n", o, "\n", p, "\n", q, "\n", r, "\n", s, "\n", t, "\n", u, "\n",
// 		v, "\n", w, "\n", x, "\n", y, "\n", z)
// }

func timess() {

	aa := util.XTime(time.Now())

	fmt.Println(util.XTime(time.Now()), aa.String())

	formatTime, err := time.Parse("2006-01-02 15:04:05", aa.String())
	qq, _ := time.LoadLocation("Asia/Shanghai")
	url := fmt.Sprintf("%s%s", "123123", "order/interior/quickorder/disable-auto-grab")
	fmt.Println(formatTime.Unix(), err, time.Now().In(qq).Unix(), url)

}

func mapinit() {

	mmm := make(map[int64]string, 2)

	if mmm == nil {
		// mmm = make(map[int64]string, 2)
		mmm[1] = "123"
	}

	fmt.Println(mmm)

}
func boo() [][]Pieces {
	var arr []byte
	// var board [][]Pieces
	// 初始化随机数组
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for _, i := range r.Perm(16) {
		arr = append(arr, byte(i)+1)
	}
	board := make([][]Pieces, 4)
	for i := int(0); i < BOARD_SIZE_jun; i++ {
		board[i] = make([]Pieces, 4)
		icelog.Info(i)
		for j := int(0); j < BOARD_SIZE_jun; j++ {
			icelog.Info(board[i][j])
			board[i][j].SHOW = 1
			board[i][j].NUMBER = arr[j+i*4]
		}
	}
	return board
}
func TestFor(T *testing.T) {
	var data int
	for i := 0; i < 10; i++ {
		data++
		icelog.Info("***", data)
		go func(i int) {
			listen2(i)
		}(data)
	}
	<-time.After(time.Second)
}
func listen2(data int) {
	fmt.Print(data)
}

func gotoFun() {

	readChannl := make(chan int)

	go func(readerChannel chan int) {
		for i := 0; i < 100; i++ {

			icelog.Info(i)

			if i == 10 {
				select {
				// 判断管道是否关闭
				case _, ok := <-readerChannel:
					if !ok {
						icelog.Info("888")
						goto BB
					}
				}
			}

		}

	BB:
		icelog.Info("****")
	}(readChannl)

	<-time.After(time.Second * 2)

	go func(readChannl chan int) {
		// time.NewTimer(3 * time.Second)
		close(readChannl)
	}(readChannl)
}

func Testarr(T *testing.T) {
	var arr []int64

	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 121)
	arr = append(arr, 11)
	icelog.Info(arr)

	var days int64
	days = 100
	days = 0 - days
	icelog.Info(int(days))

}

func https() {
	client := http_api.NewClient()
	// resp, err := client.GETV1("https://latest-test-api.lygou.cc/order/plorder/is_show?god_id=13100179", nil)
	resp, err := client.POSTV2("https://quickorder-staging-api.lygou.cc/order/interior/quickorder/disable-auto-grab", map[string]interface{}{
		"god_ids": 1896,
		"reason":  1,
	})
	if err != nil {
		icelog.Error(err.Error())
	}
	qq, err := resp.ReadAll()
	var ress protocol.Message
	err = json.Unmarshal(qq, &ress)
	icelog.Info(ress, resp)

	// resp.Body.Close()
}

func CompletePayOrder(userId, godId int64) {
	client := http_api.NewClient()
	resp, err := client.POSTV2("https://latest-test-api.lygou.cc/order/interior/quickorder/complete", map[string]interface{}{
		"user_id": userId,
		"god_id":  godId,
	})
	if err != nil {
		icelog.Error(err.Error())
	}
	if resp.StatusCode == 200 {
		icelog.Info("支付成功,调用php服务成功")
		return
	}
	ms, _ := resp.ReadAll()
	var ress protocol.Message
	err = json.Unmarshal(ms, &ress)
	icelog.Info("支付成功,调用php失败：",
		ress.Errmsg)
}

func channel() {
	ch := make(chan int, 2)
	for i := 0; i <= 10; i++ {
		// i := 1
		// for {
		select {

		case ch <- i:
			fmt.Println(i, "%%%") // 0,2,4,6,8

		case x := <-ch:
			fmt.Println(x)        // 0,2,4,6,8
			fmt.Println(i, "&&&") // 0,2,4,6,8
		default:
			fmt.Println(i, "$$$") // 0,2,4,6,8

		}

		// i++
	}

}
func Testint(T *testing.T) {
	var eee int
	var ss = "1"
	fmt.Printf("%+v", eee)
	// fmt.Printf("%v", eee)
	fmt.Printf("%v", ss)
	fmt.Printf("%s", ss)

}

func Test2(T *testing.T) {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]int)

	for key, val := range slice {
		fmt.Printf("%+v", val)
		m[key] = val
	}

	if v, ok := m[8]; ok {
		fmt.Print(v, 111111111)
	}
	for k, v := range m {
		fmt.Printf("\n key: %d, value: %d \n", k, v)
	}

	const (
		mutexLocked = 1 << iota
		mutexWoken
		mutexStarving
		mutexWaiterShift = iota
	)
	sn2 := &struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	str := "123¡" + "abc"

	fmt.Print("&333333", mutexLocked, mutexWoken, mutexStarving, mutexWaiterShift)
	fmt.Print("&&&&&&&&&&", (*sn2).age, sn2.name, str)
}

// 初始化一个17*17的棋盘，把四周用9填充，棋盘点用0填充
// $$$$$$$$$$$$$$$$$
// $...............$
// $...............$
// $...............$
// $...............$
// $...............$
// $...............$
// $...............$
// $$$$$$$$$$$$$$$$$
func init() {
	var i, j int
	for i = 0; i < BOARD_SIZE+2; i++ {
		Board[0][i] = 9
		Board[BOARD_SIZE+1][i] = 9
		Board[i][0] = 9
		Board[i][BOARD_SIZE+1] = 9
	}

	for i = 1; i <= BOARD_SIZE; i++ {
		for j = 1; j <= BOARD_SIZE; j++ {
			Board[i][j] = 0
		}
	}

}

type Pieces struct {
	SHOW   byte
	NUMBER byte
}

func golang() {
	fmt.Println(111)
	arr := make([]int, 8, 111)
	arr[1] = 1231411231283611
	var aa = "qwe"[0:0]
	fmt.Println("%!", aa)

}

func arrtest() {

	aa := make(map[int]int)
	aa[1] = 1
	aa[2] = 1
	fmt.Println(aa)
	delete(aa, 1)
	fmt.Println(aa)

}

type JSONTime struct {
	time.Time
}

func (t *JSONTime) MarshalJSON() ([]byte, error) {
	// 这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
	return []byte(fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))), nil
}

func (t *JSONTime) UnmarshalJSON(data []byte) error {
	var err error

	t.Time, err = time.Parse(`"2006-01-02 15:04:05"`, string(data))
	if err != nil {
		return err
	}

	return nil
}

func varsss() bool {

	var isGodOpenCoupon bool
	fmt.Println(123123, isGodOpenCoupon)
	return isGodOpenCoupon
}

func slices() {

	tempBoard := make([]gs_gobang.History, 0)
	for i := 0; i < 10; i++ {
		tempBoard = append(tempBoard, gs_gobang.History{
			X: int32(i),
			Y: int32(i),
		})
	}
	Shuffle(tempBoard)

	icelog.Info(tempBoard, "打乱后的数组", tempBoard[2].X)

}

func Shuffle(slice []gs_gobang.History) {
	icelog.Info(slice, "打乱后的数组33")

	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
	icelog.Info(slice, "打乱后的数组12")

}
func maptest() {

	test := make(map[int64]int64, 0)
	test[123] = 222

	if n, ok := test[123]; ok {
		icelog.Info(n, ok)
	}

}

func timesecond() {
	icelog.Info("123")

	time.Sleep(time.Duration(5) * time.Second)
	icelog.Info("456")

}

func CheckWin() bool {
	var x, y, c int
	re := AddStone(x, y, c)
	if re > 0 {
		return true
	}
	return false
}

// 落子
func AddStone(x, y, cStone int) int {
	nResult := -1
	if cStone == BLACK_STONE {
		if IsFive(x, y, BLACK_STONE, 0) {
			nResult = BLACK_FIVE
		}
	} else if cStone == WHITE_STONE {
		if IsFive(x, y, WHITE_STONE, 0) {
			nResult = WHITE_FIVE
		}
	}
	return nResult
}

// 放置棋子
func SetStone(x, y, cStone int) {
	Board[x+1][y+1] = cStone
}

/*
判断是否连5，nDir参数：方向  1横的  2 竖的 3 斜线 4 反斜线
*/
func IsFive(x, y, cStone, nDir int) bool {
	if nDir > 0 {
		if Board[x+1][y+1] != EMPTY_STONE {
			return false
		}

		if count := CountDirection(x, y, cStone, nDir); count > 4 {
			return true
		}

	} else {
		for i := 1; i < 5; i++ {
			if IsFive(x, y, cStone, i) {
				return true
			}
		}
	}
	return false
}

func Cmd(c frame.Context) {
	msg := "有人评论了你的动态"
	imapipb.SendMessage(c, &imapipb.SendMessageReq{
		Thread:      imapipb.CreateNotificationMessageThread(40002).ThreadString(),
		FromId:      1992576,
		ToId:        1896,
		ContentType: imapipb.MESSAGE_CONTENT_TYPE_NEW_CMD,
		Subtype:     40002,
		Message:     msg,
		Pt:          imapipb.PLATFORM_TYPE_PLATFORM_TYPE_APP,
	})
}

func userinfo() {
	users := make([]interface{}, 0, 2)
	userA := gameserver.UserInfo{
		UserId:   123,
		Username: "12113",
	}
	userB := gameserver.UserInfo{
		UserId:   123,
		Username: "12223",
	}

	users = append(users, userA)
	users = append(users, userB)

	// da, _ := json.Marshal(users)
	fmt.Printf("%#v\n", users)
}

func chantest() {
	ff := "114.55:111:198:12345"
	arr := strings.Split(ff, ":")
	fmt.Println(arr[0], arr[1], len(arr))
}

func mapt() {

	var Event []map[string]interface{}

	arr := map[string]interface{}{
		"X": 2,
		"Y": 2,
		"C": 1,
	}
	Event = append(Event, arr)
	Event = append(Event, arr)
	Event = append(Event, arr)
	for k, v := range Event {
		fmt.Print(v["X"], k)
	}
	fmt.Print("********")

	fmt.Print(Event, Event[len(Event)-1], Event[0]["X"])
}

/*
横坐标，纵坐标，颜色，方向（- | / \）
*/
func CountDirection(x, y, cStone, nDir int) int {

	SetStone(x, y, cStone) // 放置作为判断

	nLine := 1

	switch nDir {

	case 1: // horizontal direction
		i := x
		for i > 0 {
			if Board[i-1][y+1] == cStone {
				nLine++
			} else {
				break
			}
		}

		i = x + 2
		for i < BOARD_SIZE+1 {
			if Board[i+1][y+1] == cStone {
				nLine++
			} else {
				break
			}
		}
		break
	case 2: // vertial direction
		i := y
		for i > 0 {
			if Board[x+1][i-1] == cStone {
				nLine++
			} else {
				break
			}
		}

		i = y + 2
		for i < BOARD_SIZE+1 {
			if Board[x+1][i+1] == cStone {
				nLine++
			} else {
				break
			}
		}
		break
	case 3: // diagonal direction - '/'
		i := x
		j := y
		for i > 0 && j > 0 {
			if Board[i-1][j-1] == cStone {
				nLine++
			} else {
				break
			}
		}

		i = x + 2
		j = y + 2
		for i < BOARD_SIZE+1 && j < BOARD_SIZE+1 {
			if Board[i+1][j+1] == cStone {
				nLine++
			} else {
				break
			}
		}
		break

	case 4: // diagonal direction - '\'

		i := x
		j := y + 2
		for i > 0 && j < BOARD_SIZE+1 {
			if Board[i-1][j+1] == cStone {
				nLine++
			} else {
				break
			}
		}

		i = x + 2
		j = y
		for i < BOARD_SIZE+1 && j > 0 {
			if Board[i+1][j-1] == cStone {
				nLine++
			} else {
				break
			}
		}
		break
	default:
		break
	}
	SetStone(x, y, EMPTY_STONE) // 还原置空
	return nLine
}
