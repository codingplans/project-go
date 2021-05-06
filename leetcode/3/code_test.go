package _00_init_code

import (
	"fmt"
	"testing"
)

type test struct {
	IntEs []int
	Lists []int
	K     int
	S     string
}

var tests = []test{
	{S: "abcabcbb"},
	{S: "bbbbb"},
	{S: "pwwkew"},
	{S: ""},
	{S: "qwertyyyyyuiopasysdfg"},
}

// 3. 无重复字符的最长子串
// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
// 示例 1:
//
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:
//
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:
//
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
// 示例 4:
//
// 输入: s = ""
// 输出: 0
//
//
// 提示：
//
// 0 <= s.length <= 5 * 104
// s 由英文字母、数字、符号和空格组成

// 思考： 不重复最长 就用 map 实现，当大于 1，break出来
func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		fmt.Println("初始化")

		pre := lengthOfLongestSubstring(tests[k1].S)
		fmt.Println("结果：", pre)

	}
}

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	fmt.Println(s)

	m := [256]int{}
	max, l, r := 0, 0, -1

	for l < len(s) {
		if r+1 == len(s) {
			// 滑动完毕结束
			return max
		}
		// 初始滑动右边界
		if r+1 < len(s) && m[s[r+1]] == 0 {
			m[s[r+1]]++
			r++
		} else {
			// 若遇到重复 缩小边界 ，左边滑动，直到不重复
			m[s[l]]--
			l++

		}
		max = maxx(max, r-l+1)
	}

	return max
}

func maxx(x, y int) int {
	if x > y {
		return x
	}
	return y
}
