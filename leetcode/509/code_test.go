package _00_init_code

import (
	"fmt"
	"testing"
)

type test struct {
	IntEs []int
	Lists []int
	K     int64
	Str   string
}

var tests = []test{
	{K: 1},
	{K: 0},
	{K: 10},
	{K: 32},
	{K: 100},
	{K: 200},
}

//
// ## 题目大意
//
// 斐波那契数，通常用 F(n) 表示，形成的序列称为斐波那契数列。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
//
// ```
// F(0) = 0,   F(1) = 1
// F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
// ```
//
// 给定 N，计算 F(N)。
//
// 提示：0 ≤ N ≤ 30

func Test_upToDayUp(t *testing.T) {
	for _, v := range tests {
		fmt.Println("初始化")
		// pre := fib(v.K)
		// fmt.Println("结果：", pre)

		pre := fibV2(v.K)
		fmt.Println("结果：", pre)

	}

}

// 递归方案
func fib(n int) int {
	if n <= 1 {
		return 1
	}
	x := fib(n-1) + fib(n-2)
	return x
}

// dp 动态规划， 每次记录前面 2 个值，为下一个做准备
func fibV2(n int64) int64 {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}
	if n == 2 {
		return 1
	}
	sum, pre1, pre2 := int64(0), int64(1), int64(1)
	for i := int64(3); i <= n; i++ {
		sum = pre2 + pre1
		pre2 = pre1
		pre1 = sum
	}
	return sum
}
