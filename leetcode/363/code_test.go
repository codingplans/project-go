package _32

import (
	"fmt"
	"testing"
)

// 20210422 未解决 推荐先去做53 题再回来

// 363. 矩形区域不超过 K 的最大数值和
// 给你一个 m x n 的矩阵 matrix 和一个整数 k ，找出并返回矩阵内部矩形区域的不超过 k 的最大数值和。
//
// 题目数据保证总会存在一个数值和不超过 k 的矩形区域。

// 给你一个 m x n 的矩阵 matrix 和一个整数 k ，找出并返回矩阵内部矩形区域的不超过 k 的最大数值和。
//
// 题目数据保证总会存在一个数值和不超过 k 的矩形区域。
//
//
//
// 示例 1：
//
//
// 输入：matrix = [[1,0,1],[0,-2,3]], k = 2
// 输出：2
// 解释：蓝色边框圈出来的矩形区域 [[0, 1], [-2, 3]] 的数值和是 2，且 2 是不超过 k 的最大数字（k = 2）。
// 示例 2：
//
// 输入：matrix = [[2,2,-1]], k = 3
// 输出：3

// 提示：
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -100 <= matrix[i][j] <= 100
// -105 <= k <= 105

type test struct {
	V [][]int
	K int
}

var tests = []test{
	{[][]int{
		{1, 2},
		{2, 1}}, 1},
	{[][]int{
		{1, 2, 3, 4},
		{4, 3, 2, 1},
		{1, 1, 1, 1}}, 5},
	{[][]int{
		{1, 0, 1},
		{0, -2, -3}}, 2},
	{[][]int{
		{1, 2, 3, 4},
		{-4, -3, -2, -1},
		{1, 1, 1, 1}}, 2},
}

func Test_upToDayUp(t *testing.T) {
	for k := range tests {
		rs := maxSumSubmatrix(tests[k].V, tests[k].K)
		println(rs)
	}
}

func maxSumSubmatrix(matrix [][]int, k int) int {

	for k0, v := range matrix {
		for k1, v1 := range matrix[k0] {
			fmt.Println(v[k1], v1)
			tmp := 0

			for tmp < k && matrix[k0][k1] < k {
				// aa := matrix[k0][k1]
				tmp++
			}
		}

	}
	return 0
}
