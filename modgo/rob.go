package main

import "fmt"

// 打家劫舍

func maxx(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverse(arr []int) []int {
	n := len(arr)
	reversed := make([]int, n)
	for i, j := 0, n-1; i < n; i, j = i+1, j-1 {
		reversed[i] = arr[j]
	}
	return reversed
}

func rob2(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	if l == 0 {
		return 0
	}
	if l == 2 {
		return maxx(nums[1], nums[0])
	}
	fmt.Println(rob1(nums[1:]), rob1(nums[:l-1]), nums[1:], nums[:l-1])
	return maxx(rob1(nums[1:]), rob1(nums[:l-1]))

}

func rob1(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	dp[0] = nums[0]
	dp[1] = maxx(nums[0], nums[1])
	for i := 2; i < l; i++ {
		dp[i] = maxx(dp[i-2]+nums[i], dp[i-1])
	}
	// 返回数组最后一个值
	return dp[l-1]
}

// 解法二 DP 优化辅助空间，把迭代的值保存在 2 个变量中
func rob198_1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	curMax, preMax := 0, 0
	for i := 0; i < n; i++ {
		tmp := curMax
		curMax = maxx(curMax, nums[i]+preMax)
		preMax = tmp
	}
	return curMax
}

// 解法三 模拟
func rob11(nums []int) int {
	// a 对于偶数位上的最大值的记录
	// b 对于奇数位上的最大值的记录
	a, b := 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			a = maxx(a+nums[i], b)
		} else {
			b = maxx(a, b+nums[i])
		}
	}
	return maxx(a, b)
}

func main() {
	// Example usage:
	// wealth := []int{100, 27, 1, 1, 99}
	wealth := []int{1, 2, 3, 1}
	// maxWealth, houses := rob2(wealth)
	maxWealth := rob2(wealth)
	fmt.Println("Houses to rob:", wealth)
	fmt.Println("Maximum wealth:", maxWealth)
	// fmt.Println("Houses to rob:", houses)
}
