package _32

import (
	"fmt"
	_ "net/http/pprof"
	"testing"
)

// 69. x 的平方根
// 实现 int sqrt(int x) 函数。
//
// 计算并返回 x 的平方根，其中 x 是非负整数。
//
// 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
//
// 示例 1:
//
// 输入: 4
// 输出: 2
// 示例 2:
//
// 输入: 8
// 输出: 2
// 说明: 8 的平方根是 2.82842...,
//     由于返回类型是整数，小数部分将被舍去。

type test struct {
	V int
}

var tests = []test{
	{10},
	{8},
	{9},
	{2},
	{88},
	{1},
	{0},
	{10000},
}

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		println("初始化", tests[k1].V)
		rs := mySqrt(tests[k1].V)
		fmt.Println("结果：", rs)

	}

}

// 思路 二分法来取值，用中间数来调和边界值,不要用 k*k来计算，可能导致整形溢出。

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	m, n := 1, x>>1
	for m <= n {
		mid := m + (n-m)/2
		// println(m, n, mid)
		if mid == (x / mid) {
			return mid
		} else {
			if mid < x/mid {
				m = mid + 1
			} else {
				n = mid - 1
			}
		}

	}
	return n
}
