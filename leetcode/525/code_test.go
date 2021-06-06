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
	{IntEs: []int{0, 0, 1, 0, 0, 0, 1, 1}},
	{IntEs: []int{1, 0, 1, 0, 1, 0, 1, 0}},
	{IntEs: []int{0, 1, 1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 0}},
	{IntEs: []int{0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1}},
	// {IntEs: []int{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1}},
}

// 525. 连续数组
// 给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
//
//
//
// 示例 1:
//
// 输入: nums = [0,1]
// 输出: 2
// 说明: [0, 1] 是具有相同数量0和1的最长连续子数组。
// 示例 2:
//
// 输入: nums = [0,1,0]
// 输出: 2
// 说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。
//
//
// 提示：
//
// 1 <= nums.length <= 105
// nums[i] 不是 0 就是 1

func Test_upToDayUp(t *testing.T) {
	for _, v := range tests {
		fmt.Println("初始化")

		fmt.Println(v.IntEs)
		pre := findMaxLength(v.IntEs)
		fmt.Println("结果：", pre)
	}

}
func findMaxLength(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
		sum -= 1 ^ v
		// println(v, 1^v, sum)
	}
	println(22222222, sum, len(nums))

	if (nums[0]+nums[len(nums)-1] == 0) || nums[0]+nums[len(nums)-1] == 2 {
		if len(nums) >= 2 && nums[0] != nums[1] {
			sum += 2
			// println(444444, sum, len(nums))
		}

	}
	if sum < 0 {
		sum *= -1

	}

	m := len(nums) - sum
	if m < 0 {
		m = 0
	}
	return m

}
