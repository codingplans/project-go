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
	{
		IntEs: []int{1, 2, 3, 9},
		Lists: []int{1, 2, 3, 9},
		K:     111,
	},
	{
		IntEs: []int{9, 9, 9, 9, 9},
		Lists: []int{9, 9, 9, 9, 9},
		K:     11,
	},
	{
		IntEs: []int{7},
		Lists: []int{9, 9},
		K:     993,
	},
	{
		IntEs: []int{6},
		Lists: []int{7, 7},
		K:     809,
	},
	{
		IntEs: []int{0},
		Lists: []int{7, 7},
		K:     23,
	},
}

// 989. 数组形式的整数加法
// 对于非负整数 X 而言，X 的数组形式是每位数字按从左到右的顺序形成的数组。例如，如果 X = 1231，那么其数组形式为 [1,2,3,1]。
//
// 给定非负整数 X 的数组形式 A，返回整数 X+K 的数组形式。
//
//
//
// 示例 1：
//
// 输入：A = [1,2,0,0], K = 34
// 输出：[1,2,3,4]
// 解释：1200 + 34 = 1234
// 示例 2：
//
// 输入：A = [2,7,4], K = 181
// 输出：[4,5,5]
// 解释：274 + 181 = 455
// 示例 3：
//
// 输入：A = [2,1,5], K = 806
// 输出：[1,0,2,1]
// 解释：215 + 806 = 1021
// 示例 4：
//
// 输入：A = [9,9,9,9,9,9,9,9,9,9], K = 1
// 输出：[1,0,0,0,0,0,0,0,0,0,0]
// 解释：9999999999 + 1 = 10000000000
//
//
// 提示：
//
// 1 <= A.length <= 10000
// 0 <= A[i] <= 9
// 0 <= K <= 10000
// 如果 A.length > 1，那么 A[0] != 0

// 2021.6.25 思路 这道题比较简单，但是基于字节的要求，是两个大数数组求和，下面再做一遍

func Test_upToDayUp(t *testing.T) {
	for _, v := range tests {
		fmt.Println("初始化", v.IntEs)
		pre := addToArrayForm(v.IntEs, v.K)
		// pre := addToArrayFormV2(v.IntEs, v.Lists)
		fmt.Println("结果：", pre)
	}

}

func addToArrayFormV2(num, num2 []int) []int {
	if len(num) < len(num2) {
		num, num2 = num2, num
	}

	l1, l2 := len(num)-1, len(num2)-1
	x := 0
	for k := range num {
		if l2-k >= 0 {
			num[l1-k] += num2[l2-k]
		}
		num[l1-k] = num[l1-k] + x
		x = 0
		if num[l1-k] >= 10 {
			x = 1
			num[l1-k] %= 10
		}
	}
	if x > 0 {
		num = append([]int{1}, num...)
	}

	return num

}
func addToArrayForm(num []int, k int) []int {
	la := len(num) - 1
	x := 0
	for m := range num {
		num[la-m] = num[la-m] + k%10 + x
		x = 0
		if num[la-m] >= 10 {
			num[la-m] %= 10
			x = 1
		}
		k /= 10
	}

	for k > 0 {
		c := k%10 + x
		x = 0
		if c >= 10 {
			x = 1
			c %= 10
		}
		num = append([]int{c}, num...)
		k /= 10
	}
	if x > 0 {
		num = append([]int{1}, num...)
	}
	return num
}
