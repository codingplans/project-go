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
		K: 111,
	},
	{
		K: 11,
	},
	{
		K: 1,
	},
	{
		K: 5,
	},
	{
		K: 23,
	},
}

// 172. 阶乘后的零
// 给定一个整数 n ，返回 n! 结果中尾随零的数量。
//
// 提示n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1
//
//
//
// 示例 1：
//
// 输入：n = 3
// 输出：0
// 解释：3! = 6 ，不含尾随 0
// 示例 2：
//
// 输入：n = 5
// 输出：1
// 解释：5! = 120 ，有一个尾随 0
// 示例 3：
//
// 输入：n = 0
// 输出：0

// Darren思考： 这题来的太突然，没意识到 ，就过了，本想循环，但是直接递归了。也没有什么溢出问题。 下次再来练练巩固
func Test_upToDayUp(t *testing.T) {
	for _, v := range tests {
		fmt.Println("初始化", v.K)
		pre := trailingZeroes(v.K)
		fmt.Println("结果：", pre)
	}

}
func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}

	return n/5 + trailingZeroes(n/5)
}
