package _32

import (
	"testing"
)

// 28. 实现 strStr()
// 实现 strStr() 函数。
//
// 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。
//
//
//
// 说明：
//
// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//
// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。
//
//
//
// 示例 1：
//
// 输入：haystack = "hello", needle = "ll"
// 输出：2
// 示例 2：
//
// 输入：haystack = "aaaaa", needle = "bba"
// 输出：-1
// 示例 3：
//
// 输入：haystack = "", needle = ""
// 输出：0
//
//
// 提示：
//
// 0 <= haystack.length, needle.length <= 5 * 104
// haystack 和 needle 仅由小写英文字符组成
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/implement-strstr
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
type test struct {
	haystack string
	needle   string
}

var tests = []test{
	{"mississippi", "issip"},
	{"asdsad", "ds"},
	{"hello", "ll"},
	{"aaaaa", "bba"},
	{"", ""},
	{"", "a"},
	{"aa", "aaaaa"},
	{"aaa", "aaaa"},
	{"darrenzzy", "zaz"},
	{"mississippi", "issipi"},
}

func Test_upToDayUp(t *testing.T) {
	for k := range tests {
		rs := strStr(tests[k].haystack, tests[k].needle)
		println(rs)
	}
}

// 通过切片整体作比较
func strStr(haystack string, needle string) int {
	l := len(needle)
	if l < 0 {
		return 0
	}
	l2 := len(haystack)
	if l > len(haystack) {
		return -1
	}

	for k := 0; k < l2-l+1; k++ {
		if needle == haystack[k:k+l] {
			return k
		}
	}

	return -1

}

// 单个字符比较
func strStr2(haystack string, needle string) int {
	l := len(needle) - 1
	if l < 0 {
		return 0
	}
	l2 := len(haystack) - 1
	if l > len(haystack)-1 {
		return -1
	}

	for k := range haystack {
		i := 0
		k2 := k
		for k2 <= l2 && haystack[k2] == needle[i] {
			if i == l {
				return k2 - l
			}
			i++
			k2++
		}

	}
	return -1

}
