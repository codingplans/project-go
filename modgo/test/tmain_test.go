package main

import (
	"bytes"
	"fmt"
	"github.com/Darrenzzy/testgo/structures"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

// 随便练一下 二叉树排序 =》堆排序
func Test2TreeSoft(t *testing.T) {
	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	arr := []int{4, 3, 5, 1, 2, 6, 7}

	ts := arr[2 : len(arr)-1 : 7]
	fmt.Println(len(ts), cap(ts), ts)

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
	ww.Add(10)
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
	if tree == nil {
		return
	}
	var stack []*structures.TreeNode
	// 先压栈顶元素
	stack = []*structures.TreeNode{tree}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
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
	aa := 3
	// 0000100
	bb := 4
	t.Log(aa >> 1)
	t.Log(aa << 10)
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
	a := 1
	b := 2
	defer calc(a, calc(a, b, "0"), "1")
	a = 0
	defer calc(a, calc(a, b, "3"), "2")
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
}

// param: days 为多少天以后
// return: 今天+days 天之后的日期,所在月的最后一天, 按"2006年01月02日"格式化
func LastDateOfMonth(days int, ct time.Time) string {
	d := ct.AddDate(0, 0, days)              // time.Now()可以换成支持测试环境调时间的方法
	firstDate := d.AddDate(0, 0, -d.Day()+1) // 当月的第一天
	lastDate := firstDate.AddDate(0, 1, -1)  // 当月的最后一天
	return lastDate.Format("2006年01月02日")
}

type PayWay struct {
	//    支付id
	Id  int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Ids int64 `protobuf:"varint,2,opt,name=id,proto3" json:"ids,omitempty"`
	// 支付名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}
