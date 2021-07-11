package main

import (
	"bytes"
	"fmt"
	"github.com/Darrenzzy/testgo/structures"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestReversListV2(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	tree := structures.Ints2TreeNode(arr)

	stack := []*structures.TreeNode{tree}
	// sss(stack, 1)
	fmt.Println(arr[1:2], stack)
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
				diff = (^diff + 1)
			}
			keys[diff]++
			// arr = mergeappend(arr, diff)

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
	println(a|b, x^a^b, (^b + 1))
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
		mid = int(uint((i + l)) >> 1)
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
	k := minPathSumV2([][]int{[]int{1, 3, 5, 9}, []int{8, 1, 3, 4}, []int{5, 0, 6, 1}, []int{8, 8, 4, 0}})
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
	arr := []int{4, 3, 5, 1, 2, 6, 33, 12, 1, 55, 3, 2, 111, 57, 7, 5}
	// arr := []int{4, 3, 5, 1, 2, 6}

	fmt.Println(arr)
	QuickSoft(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// 堆排序小练
func TestHeapSort(t *testing.T) {
	// arr := []int{4, 3, 5, 1, 2, 6, 7}
	arr := []int{1, 4, 3, 2, 6, 5, 8, 7, 9, 0}
	fmt.Println(arr)

	BuildHeap(arr, len(arr))
	fmt.Println(arr)
}

// 随便练一下 二叉树排序 =》堆排序
func Test2TreeSoft(t *testing.T) {
	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
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

type S struct {
	cl     chan struct{}
	num    int
	notity chan int
	wg     sync.WaitGroup
	sync.Mutex
}
