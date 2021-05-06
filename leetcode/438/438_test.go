package _32

import (
	"fmt"
	"testing"
)

// 给定一个字符串s和一个非空字符串p，找到s中所有是p的字母异位词的子串，返回这些子串的起始索引。
// 字符串只包含小写英文字母，并且字符串s和 p的长度都不超过 20100。
// 说明：
// 字母异位词指字母相同，但排列不同的字符串。
// 不考虑答案输出的顺序。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
// 示例 1:
//
// 输入:
// s: "cbaebabacd" p: "abc"
//
// 输出:
// [0, 6]
//
// 解释:
// 起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
// 起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
//  示例 2:
//
// 输入:
// s: "abab" p: "ab"
//
// 输出:
// [0, 1, 2]
//
// 解释:
// 起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
// 起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
// 起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
//

func Test_Problem438(t *testing.T) {

	flas := problem("abcbac", "cba")
	fmt.Printf("%v", flas)
}

func problem(s, p string) []int {

	mach := [26]int{}
	// mach := []byte{}
	// 先把 p 放到数组中
	for k := range p {
		mach[p[k]-'a']++
	}
	ls := len(s)
	l := 0
	// 存对应下标
	keys := []int{}
	// 一一比对
	for l < ls {
		m2 := mach
		i := 0
		for l+i < ls && m2[s[l+i]-'a'] > 0 {
			m2[s[l+i]-'a']--
			i++
		}
		if i == len(p) {
			keys = append(keys, l)
		}
		l++
	}
	return keys
}

func problemV2(s, p string) []int {
	m := [256]int{}
	count := 0
	flas := []int{}
	for k := range p {
		m[p[k]]++
		count++
	}

	for k := range s {
		count2 := count
		i := 0
		m2 := m
		for count2 >= 0 && k+i < len(s) && m2[s[k+i]] > 0 {
			// if m2[s[k+i]] > 0 {
			// 命中减一
			count2--
			m2[s[k+i]]--
			i++
			// } else {
			// 否则调回来
			// count2 += i
			// break
			// }
		}
		if count == i {
			flas = append(flas, k)
		}
		m2[s[k+i]]++

	}
	return flas
}
