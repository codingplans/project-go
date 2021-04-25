package _32

import (
	"testing"
)

// 53. 最大子序和
// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 示例 1：
//
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
// 示例 2：
//
// 输入：nums = [1]
// 输出：1
// 示例 3：
//
// 输入：nums = [0]
// 输出：0
// 示例 4：
//
// 输入：nums = [-1]
// 输出：-1
// 示例 5：
//
// 输入：nums = [-100000]
// 输出：-100000
//
//
// 提示：
//
// 1 <= nums.length <= 3 * 104
// -105 <= nums[i] <= 105
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/maximum-subarray
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type test struct {
	V []int
}

var tests = []test{
	{[]int{1, 2, -1, 3, 4}},
	{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
	{[]int{-1}},
	{[]int{-2, -1}},
	{[]int{-2, 1}},
	{[]int{0}},
	{[]int{10, -1, 9}},
	{[]int{10, -111, 9}},
}

func Test_upToDayUp(t *testing.T) {
	for k := range tests {
		// panic(121)
		rs := maxSubArray(tests[k].V)
		println(rs)

		rs = maxSubArrayV2(tests[k].V)
		println(rs)
	}
}

func BenchmarkUpToDayUp(t *testing.B) {
	for i := 0; i < t.N; i++ {
		for k := range tests {
			_ = maxSubArray(tests[k].V)
			// println(rs)
		}
	}
}
func BenchmarkMaxSubArrayV2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		for k := range tests {
			_ = maxSubArrayV2(tests[k].V)
			// println(rs)
		}
	}
}

func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	var x int = nums[0]

	for k := range nums {
		x = max(x, nums[k])
		sum := nums[k]
		for k2 := k + 1; k2 < l; k2++ {
			sum = sum + nums[k2]
			x = max(sum, x)
		}
	}

	return x
}

func maxSubArrayV2(nums []int) int {
	m := nums[0]
	for k := 1; k < len(nums); k++ {
		if nums[k] < nums[k-1]+nums[k] {
			nums[k] += nums[k-1]
		}
		if nums[k] > m {
			m = nums[k]
		}
	}

	return m
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y

}
