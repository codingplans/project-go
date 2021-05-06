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
}

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		println("初始化", tests[k1].V)
		// fmt.Println(tests[k1].L1)
		// arr := threeSum(tests[k1].L1)
		rs := mySqrt(tests[k1].V)
		fmt.Println("结果：", rs)

	}

}
func mySqrt(x int) int {

	var a int
	shift := 2
	x, y := 1, 1
	for {
		println(a, 1<<shift, x, y, shift)

		if x == 0 && y == 0 {
			return shift
		}
		a = x >> shift
		if a < 1<<shift {
			shift--
			x = 0
			continue
		}
		if a > 1<<shift {
			shift++
			y = 0
			continue
		}
	}
	// println(a, 1<<shift)
	// return x

}
