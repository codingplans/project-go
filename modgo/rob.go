package main

import "fmt"

func rob(nums []int) (int, []int) {
	n := len(nums)
	if n == 0 {
		return 0, []int{}
	}

	if n == 1 {
		return nums[0], []int{0}
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	// selectedHouses := make([]int, n)

	// if nums[0] > nums[1] {
	// 	selectedHouses[0] = 1
	// } else {
	// 	selectedHouses[1] = 1
	// }

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		// if dp[i-1] > dp[i-2]+nums[i] {
		// 	selectedHouses[i] = selectedHouses[i-1]
		// } else {
		// 	selectedHouses[i] = selectedHouses[i-2] + 1
		// }
	}

	// Traceback to find the houses to rob
	robbedHouses := make([]int, 0)
	i := n - 1
	for i >= 0 {
		if i == 0 {
			robbedHouses = append(robbedHouses, i+1)
			break
		} else if i == 1 {
			if dp[i] > dp[i-1] {
				robbedHouses = append(robbedHouses, i+1)
			}
			break
		} else if dp[i] > dp[i-1] {
			robbedHouses = append(robbedHouses, i+1)
			i -= 2
		} else {
			i--
		}
	}

	return dp[n-1], robbedHouses
	// maxWealth := dp[n-1]
	// maxWealthHouses := []int{}
	// i := n - 1
	//
	// for i >= 0 {
	// 	if selectedHouses[i] == 1 {
	// 		maxWealthHouses = append(maxWealthHouses, i)
	// 		break
	// 	}
	// 	if selectedHouses[i] == 2 {
	// 		maxWealthHouses = append(maxWealthHouses, i, i-1)
	// 		break
	// 	}
	// 	maxWealthHouses = append(maxWealthHouses, i)
	// 	i -= 2
	// }
	//
	// return maxWealth, reverse(maxWealthHouses)
}

func max(a, b int) int {
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
		return max(nums[1], nums[0])
	}
	fmt.Println(rob1(nums[1:]), rob1(nums[:l-1]), nums[1:], nums[:l-1])
	return max(rob1(nums[1:]), rob1(nums[:l-1]))

}

func rob1(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < l; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[l-1]
}

//
// func max(a,b int)int{
// 	if a>b{
// 		return a
// 	}
//
// 	return b
// }

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
