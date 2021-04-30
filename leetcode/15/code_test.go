package _32

import (
	"fmt"
	_ "net/http/pprof"
	"reflect"
	"testing"
)

// 15. 三数之和
// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例 1：
//
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// 示例 2：
//
// 输入：nums = []
// 输出：[]
// 示例 3：
//
// 输入：nums = [0]
// 输出：[]
//
//
// 提示：
//
// 0 <= nums.length <= 3000
// -105 <= nums[i] <= 105

type test struct {
	L1 []int
}

var tests = []test{
	{[]int{-2, 0, 1, 1, 1, 2}},
	{[]int{10, -19, 90, 0, 9, 9, 9, 9, 9, -99, 9, -9, 9, 9}},   // [8,9,9,9,0,0,0,1]
	{[]int{10, -19, 90, 0, 9, 9, 9, 9, 9, -99, 9, -9, 9, 9}},   // [8,9,9,9,0,0,0,1]
	{[]int{10, -19, 90, 0, 9, 9, 9, 9, 9, -99, 9, -9, 9, 9}},   // [8,9,9,9,0,0,0,1]
	{[]int{10, -19, 90, 0, 9, 9, 659, 9, 9, -99, 9, -9, 9, 9}}, // [8,9,9,9,0,0,0,1]
	{[]int{10, -19, 90, 0, 9, 9, 19, 9, 9, -99, 9, -9, 9, 9}},  // [8,9,9,9,0,0,0,1]
	{[]int{10, -19, 90, 0, 9, 9, 29, 9, 9, -99, 9, -9, 9, 9}},  // [8,9,9,9,0,0,0,1]
	{[]int{10, -19, 90, 0, 9, 9, 39, 9, 9, -99, 9, -9, 9, 9}},  // [8,9,9,9,0,0,0,1]
	{[]int{10, -19, 90, 0, 9, 9, 9, 9, 9, -99, 9, -9, 9, 9}},   // [8,9,9,9,0,0,0,1]
	{[]int{10, -19, 90, 0, 9, 9, 9, 9, 9, -99, 9, -9, 9, 9}},   // [8,9,9,9,0,0,0,1]
	{[]int{10, -19, 90, 0, 9, 9, 49, 9, 9, -99, 9, -9, 9, 9}},  // [8,9,9,9,0,0,0,1]
	{[]int{10, -19, 90, 0, 9, 9, 9, 9, 9, -99, 9, -9, 9, 9}},   // [8,9,9,9,0,0,0,1]
	{[]int{10, -19, 90, 0, 9, 569, 9, 9, 9, -99, 9, -9, 9, 9}}, // [8,9,9,9,0,0,0,1]
	{[]int{10, -19, 90, 0, 9, 9, 9, 9, 9, -99, 9, -9, 9, 9}},   // [8,9,9,9,0,0,0,1]
	{[]int{10, -19, 90, 0, 9, 9, 9, 9, 9, -99, 9, -9, 9, 9}},   // [8,9,9,9,0,0,0,1]
	{[]int{10, -19, 90, 0, 9, 9, 9, 9, 9, -99, 9, -9, 9, 9}},   // [8,9,9,9,0,0,0,1]
	{[]int{5, 4, 3, 2, 1}},
	{[]int{9, 0, -9, -91}},    // [8,9,9,9,0,0,0,1]
	{[]int{-5, -10, 5, 6, 4}}, // 8 0 7
	{[]int{-1, -2, -3, 4, 5}},
	{[]int{-1, -2, -3, 4, 5}},
	{[]int{-1, -2, -3, 4, 5}},
	{[]int{-1, -2, -3, 4, 5}},
	{[]int{-1, -2, -3, 4, 5}},
	{[]int{1, 1, -2}},
	{[]int{-2, 0, 1, 1, 2}},
	{[]int{0, 0, 0}},
	{[]int{0, 0, 0, 0}},
	{[]int{-1, 0, 1, 2, -1, -4}},
}

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		println("初始化")
		// fmt.Println(tests[k1].L1)
		// arr := threeSum(tests[k1].L1)
		arr := threeSumV2(tests[k1].L1)
		fmt.Println("结果：", arr)

	}

}

func BenchmarkV2(b *testing.B) {
	benchmarks := []struct {
		name string
	}{
		// TODO: benchmarks
	}
	for _, bm := range benchmarks {
		b.Run(bm.name+"222", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for k1 := range tests {
					threeSumV2(tests[k1].L1)
				}
			}
		})
	}
}

func BenchmarkV3(b *testing.B) {
	benchmarks := []struct {
		name string
	}{
		// TODO: benchmarks
	}
	for _, bm := range benchmarks {
		b.Run(bm.name+"111", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for k1 := range tests {
					threeSumV3(tests[k1].L1)
				}
			}
		})
	}
}

// 方法二 排序，双指针
func threeSumV2(nums []int) [][]int {

	l := len(nums)
	if l < 3 {
		return nil
	}
	// 初始化容量，减少开辟内存开销
	arrs := make([][]int, 0, 5)
	// var arrs [][]int
	// fmt.Println(nums)
	QuickSort(nums, 0, l-1)
	// fmt.Println(nums)

	if nums[0] > 0 || nums[l-1] < 0 {
		return nil
	}
	var i, m, n int

	for i <= l-3 {
		m = i + 1
		n = l - 1
		for m <= l-2 && n >= i+1 && m < n && nums[m]+nums[i] <= 0 && nums[n] >= 0 {
			// 符合条件塞入数组
			if nums[i]+nums[m]+nums[n] == 0 {
				arr := []int{nums[i], nums[m], nums[n]}
				// 判断切片是否相等
				if len(arrs) > 0 && reflect.DeepEqual(arrs[len(arrs)-1], arr) {
					// break
				} else {
					arrs = append(arrs, arr)
				}
			}
			// 遇到重复数字跳过
			for n-1 > i+1 && nums[n] == nums[n-1] && nums[n-1]+nums[m]+nums[i] > 0 {
				n--
			}
			// 遇到重复数字跳过
			for m+1 <= l-2 && nums[m] == nums[m+1] && nums[n]+nums[m+1]+nums[i] < 0 {
				m++
			}
			if m == n {
				break
			}

			// 每轮范围缩减
			if nums[i]+nums[m]+nums[n] >= 0 {
				n--
			} else {
				m++
			}

		}
		i++
		// 遇到重复数字跳过
		for nums[i] == nums[i-1] && i <= l-3 {
			i++
		}
	}

	return arrs
}

// 方法二 排序，双指针
func threeSumV3(nums []int) [][]int {

	l := len(nums)
	if l < 3 {
		return nil
	}
	// 初始化容量，减少开辟内存开销
	arrs := make([][]int, 0, 5)
	// var arrs [][]int
	// fmt.Println(nums)
	QuickSort(nums, 0, l-1)
	// fmt.Println(nums)

	if nums[0] > 0 || nums[l-1] < 0 {
		return nil
	}
	var i, m, n int

	for i <= l-3 {
		m = i + 1
		n = l - 1
		for m <= l-2 && n >= i+1 && m < n && nums[m]+nums[i] <= 0 && nums[n] >= 0 {
			// 符合条件塞入数组
			if nums[i]+nums[m]+nums[n] == 0 {
				arr := []int{nums[i], nums[m], nums[n]}
				// 判断切片是否相等
				if len(arrs) > 0 && reflect.DeepEqual(arrs[len(arrs)-1], arr) {
					// break
				} else {
					arrs = append(arrs, arr)
				}
			}
			// 遇到重复数字跳过
			for n-1 > i+1 && nums[n] == nums[n-1] && nums[n-1]+nums[m]+nums[i] > 0 {
				n--
			}
			// 遇到重复数字跳过
			for m+1 <= l-2 && nums[m] == nums[m+1] && nums[n]+nums[m+1]+nums[i] < 0 {
				m++
			}
			if m == n {
				break
			}

			// 每轮范围缩减
			if nums[i]+nums[m]+nums[n] >= 0 {
				n--
			} else {
				m++
			}

		}
		i++
		// 遇到重复数字跳过
		for nums[i] == nums[i-1] && i <= l-3 {
			i++
		}
	}

	return arrs
}

func QuickSort(arr []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	i := startIndex
	j := endIndex
	for i != j {
		for arr[startIndex] <= arr[j] && i < j {
			j--
		}
		for arr[startIndex] >= arr[i] && i < j {
			i++
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}

	}
	arr[i], arr[startIndex] = arr[startIndex], arr[i]
	QuickSort(arr, startIndex, i-1)
	QuickSort(arr, i+1, endIndex)
}

func QuickSortByByte(arr []byte, startIndex, endIndex byte) {
	if startIndex >= endIndex {
		return
	}
	i := startIndex
	j := endIndex
	for i != j {
		for arr[startIndex] <= arr[j] && i < j {
			j--
		}
		for arr[startIndex] >= arr[i] && i < j {
			i++
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}

	}
	arr[i], arr[startIndex] = arr[startIndex], arr[i]
	QuickSortByByte(arr, startIndex, i-1)
	QuickSortByByte(arr, i+1, endIndex)
}

// 方案一 暴力解法
func threeSum(nums []int) [][]int {
	var arrs [][]int
	l := len(nums)
	if l < 3 {
		return nil
	}
	i, j := 1, 2
	for k, v := range nums {
		for i < l-k {
			for j < l-k {
				if v+nums[k+i]+nums[k+j] == 0 {
					sum, x1 := min(v, nums[k+i])
					sum2, x := min(nums[j+k], x1)
					sum3, x3 := min(sum2-x, sum-x1)
					arrs = append([][]int{{x, x3, sum3 - x3}}, arrs...)
				}
				j++
			}
			i++
			j = i + 1

		}
		i = k + 1
	}
	return arrs
}

func min(a, b int) (sum int, max int) {
	if a < b {
		return a + b, a
	}
	return a + b, b
}
