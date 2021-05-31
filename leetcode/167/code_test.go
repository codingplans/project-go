package _00_init_code

import (
	"fmt"
	"testing"
)

type test struct {
	IntEs []int
	Lists []int
	K     int
	Str   string
}

var tests = []test{
	{IntEs: []int{1, 3, 4, 4, 4, 4, 4, 4, 4, 4, 5, 7, 9}, Lists: []int{0, 1, 2, 4, 4, 4, 4, 4, 4, 5, 6, 8, 88}, K: 9},
	{IntEs: []int{1, 3, 4, 4, 4, 4, 4, 4, 4, 4, 5, 7, 9}, Lists: []int{0, 1, 2, 4, 4, 4, 4, 4, 4, 5, 6, 8, 88}, K: 8},
}

// 求和
// ## 题目大意
//
// 找出两个数之和等于 target 的两个数字，要求输出它们的下标。注意一个数字不能使用 2 次。下标从小到大输出。假定题目一定有一个解。
//
// ## 解题思路
//
// 这一题比第 1 题 Two Sum 的问题还要简单，因为这里数组是有序的。可以直接用第一题的解法解决这道题。

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		res = [][]int{}
		fmt.Println("初始化")
		fmt.Println(tests[k1].IntEs)
		arrSumV2(tests[k1].IntEs, tests[k1].Lists, tests[k1].K)
		// arrSum(tests[k1].IntEs, tests[k1].Lists, tests[k1].K)
		fmt.Println("结果：", res)

	}
}

/*[1 3 4 4 4 4 4 4 4 4 5 7 9]
结果： [[1 8] [3 6] [4 5] [5 4] [7 2] [9 0]]
初始化
[1 3 4 4 4 4 4 4 4 4 5 7 9]
结果： [[3 5] [4 4] [4 4] [4 4] [4 4] [4 4] [4 4] [7 1]]
--- PASS: Test_upToDayUp (0.00s)*/

var res [][]int

func arrSum(arr1, arr2 []int, target int) {
	i, j := 0, len(arr2)-1
	for i < len(arr1) && j >= 0 {
		if arr1[i]+arr2[j] == target {
			res = append(res, []int{arr1[i], arr2[j]})
			i++
			j--
			continue
		}
		if arr1[i]+arr2[j] > target {
			j--
		} else if arr1[i]+arr2[j] < target {
			i++
		}
	}

}

func arrSumV2(arr1, arr2 []int, target int) {
	md := make(map[int]int, len(arr1))
	for k := range arr1 {
		md[arr1[k]] = k
	}
	for k := range arr2 {
		diff := target - arr2[k]
		if _, ok := md[diff]; ok {
			res = append(res, []int{diff, arr2[k]})
		}
	}

}
