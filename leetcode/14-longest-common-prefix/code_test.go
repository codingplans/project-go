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
	Strs  []string
}

var tests = []test{
	{Strs: []string{"flower", "flow", "flight"}},
	{Strs: []string{"aass", "aass", "aass"}},
	{Strs: []string{"dog", "racecar", "car"}},
	{Strs: []string{"ab", "a"}},
}

// 14. 最长公共前缀
// 编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
//
//
// 示例 1：
//
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
// 示例 2：
//
// 输入：strs = ["dog","racecar","car"]
// 输出：""
// 解释：输入不存在公共前缀。

func Test_upToDayUp(t *testing.T) {
	for _, v := range tests {
		fmt.Println("初始化", v.Strs)
		pre := longestCommonPrefix(v.Strs)
		fmt.Println("结果：", pre)
	}

}

// 最简单的方式来遍历 时间复杂度是 mn
func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0 {
		return ""
	}
	res := 0
	for k := range strs[0] {
		for i := l - 1; i > 0; i-- {
			if len(strs[i]) <= k || strs[i][k] != strs[0][k] {
				return strs[0][:res]
			}
		}
		res++
	}
	return strs[0][:res]

}

// 二分法解决，很不错，值得学习 ，mn* Log n
func longestCommonPrefixV2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	isCommonPrefix := func(length int) bool {
		str0, count := strs[0][:length], len(strs)
		for i := 1; i < count; i++ {
			if strs[i][:length] != str0 {
				return false
			}
		}
		return true
	}
	minLength := len(strs[0])
	for _, s := range strs {
		if len(s) < minLength {
			minLength = len(s)
		}
	}
	low, high := 0, minLength
	for low < high {
		mid := (high-low+1)/2 + low
		if isCommonPrefix(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return strs[0][:low]
}
