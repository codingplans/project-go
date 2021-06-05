package _00_init_code

import (
	"fmt"
	"testing"
)

type test struct {
	IntEs []int
	Lists []int
	K     int
	K2    int
	Str   string
	Str2  string
}

var tests = []test{
	{Str: "()())())()()"},
	{Str: "((()(((()))()()())"},
	{Str: "()(((())()((((((((((((((((((()))))))))))))))))))())))))))))))))))((((((()))))"},
	{Str: "()(((())()((((((((()))))))))))))))))))((((((())))))))))))))))"},
}

// 32. 最长有效括号
// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//
//
// 示例 1：
//
// 输入：s = "(()"
// 输出：2
// 解释：最长有效括号子串是 "()"
// 示例 2：
//
// 输入：s = ")()())"
// 输出：4
// 解释：最长有效括号子串是 "()()"
// 示例 3：
//
// 输入：s = ""
// 输出：0
//
//
// 提示：
//
// 0 <= s.length <= 3 * 104

func Test_upToDayUp(t *testing.T) {
	for _, v := range tests {
		fmt.Println("初始化")

		pre := longestValidParentheses(v.Str)
		fmt.Println("结果：", pre)
	}

}
func longestValidParentheses(s string) int {
	// fmt.Println(s)
	maxl := 0
	i := 0
	arr := []int{-1, -1}
	fmt.Println(arr[:len(arr)-1])
	fmt.Println(arr[1:])
	for i < len(s) {
		if int(s[i]) == 40 {
			arr = append(arr, i)
		} else {
			arr = arr[:len(arr)-1]
			if len(arr) == 0 {
				arr = append(arr, i)
			} else {
				maxl = max(maxl, i-arr[len(arr)-1])
			}
		}

		i++
	}

	return maxl
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
