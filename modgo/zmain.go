package main

import (
	"fmt"

	"github.com/prometheus/common/log"
	"golang.org/x/text/feature/plural"
)

//go:generate gotext -srclang=en update -out=./catalog/catalog.go -lang=en,el

type entry struct {
	tag, key string
	msg      interface{}
}

var entries = [...]entry{
	{"en", "Hello World", "Hello World"},
	{"zh", "Hello World", "你好世界"},
	{"en", "%d task(s) remaining!", plural.Selectf(1, "%d",
		"=1", "One task remaining!",
		"=2", "Two tasks remaining!",
		"other", "[1]d tasks remaining!",
	)},
	{"zh", "%d task(s) remaining!", plural.Selectf(1, "%d",
		"=1", "剩余一项任务！",
		"=2", "剩余两项任务！",
		"other", "剩余 [1]d 项任务！",
	)},
}

func init() {
	// for _, e := range entries {
	// 	tag := language.MustParse(e.tag)
	// 	switch msg := e.msg.(type) {
	// 	case string:
	// 		message.SetString(tag, e.key, msg)
	// 	case catalog.Message:
	// 		message.Set(tag, e.key, msg)
	// 	case []catalog.Message:
	// 		message.Set(tag, e.key, msg...)
	// 	}
	// }
}
func main() {

	var v int
	var p *int
	var w interface{}
	fmt.Println(p == nil, w == nil)
	p = &v
	w = (*int)(nil)
	fmt.Println(p == nil, w == nil)
	p = nil
	w = p
	fmt.Println(p == nil, w == nil)

	// p := message.NewPrinter(language.Chinese)
	// p.Printf("Hello World")
	// p.Printf("Hello World")
	// p.Printf("Hello World!\n")
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Maxer(maxDepth(root.Left), maxDepth(root.Right))
}

func Maxer(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func minPathSum(grid [][]int) int {
// 	if len(grid) == 0 || len(grid[0]) == 0 {
// 		return 0
// 	}
// 	rows, columns := len(grid), len(grid[0])
// 	dp := make([][]int, rows)
// 	for i := 0; i < len(dp); i++ {
// 		dp[i] = make([]int, columns)
// 	}
// 	dp[0][0] = grid[0][0]
// 	for i := 1; i < rows; i++ {
// 		dp[i][0] = dp[i-1][0] + grid[i][0]
// 	}
// 	for j := 1; j < columns; j++ {
// 		dp[0][j] = dp[0][j-1] + grid[0][j]
// 	}
// 	for i := 1; i < rows; i++ {
// 		for j := 1; j < columns; j++ {
// 			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
// 		}
// 	}
// 	return dp[rows-1][columns-1]
// }

func minPathSum(grid [][]int) int {
	if len(grid[0]) == 0 || len(grid) == 0 {
		return 0
	}
	x, y := len(grid), len(grid[0])
	println(x, y)
	dp := make([][]int, x)
	dp[0] = make([]int, y)

	dp[0][0] = grid[0][0]

	for i := 1; i < x; i++ {
		dp[i] = make([]int, y)
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}

	for i := 1; i < y; i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}
	log.Info(grid, dp)
	for i := 1; i < x; i++ {
		for j := 1; j < y; j++ {
			m := dp[i-1][j]
			n := dp[i][j-1]
			dp[i][j] = min(n, m) + grid[i][j]
		}
	}
	return dp[x-1][y-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// aa := findMin([]int{3, 4, 5, 1, 2})
func findMin(nums []int) int {
	high := len(nums) - 1
	low := 0
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] {
			high = pivot
		} else if nums[pivot] > nums[high] {
			low = pivot + 1
		} else {
			high--
		}
	}
	return nums[low]
}

func minArray(numbers []int) int {
	temp := numbers[0]
	i := 1
	for i < len(numbers) {
		if temp > numbers[i] {
			temp = numbers[i]
		}
		i++
	}
	return temp
}
func generateTrees(n int) []*TreeNode {
	log.Info(123)
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

type TreeNode1 struct {
	i     int
	Left  *TreeNode
	Right *TreeNode
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTrees := []*TreeNode{}
	// 枚举可行根节点
	for i := start; i <= end; i++ {
		// 获得所有可行的左子树集合
		// leftTrees := helper(start, i-1)
		// // 获得所有可行的右子树集合
		// rightTrees := helper(i+1, end)
		// 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
		// for _, left := range leftTrees {
		// 	for _, right := range rightTrees {
		// 		// currTree := &TreeNode{i, nil, nil}
		// 		// currTree.Left, currTree.Right = left, right
		// 		// allTrees = append(allTrees, currTree)
		// 	}
		// }
	}
	return allTrees
}
