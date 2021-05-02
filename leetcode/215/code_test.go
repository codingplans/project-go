package _32

import (
	"fmt"
	"testing"
)

//  215. 数组中的第K个最大元素
// 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 示例 1:
//
// 输入: [3,2,1,5,6,4] 和 k = 2
// 输出: 5
// 示例 2:
//
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
// 输出: 4
// 说明:
//
// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。

// 思考： 很简单的感觉 先快排，再找最后 4 个大小

type test struct {
	L1 []int
	K  int
}

var tests = []test{
	{[]int{5, 4, 3, 2, 1}, 1},
	{[]int{-2, 0, 1, 1, 2}, 1},
	{[]int{9, 7, 6, 4, 3}, 1},
	{[]int{4, 56, 235, 362, 3425, 1, 4, 5}, 4},
	{[]int{4, 56, 235, 36, 3425, 1, 4, 5}, 4},
	{[]int{4, 56, 235, 3, 3425, 1, 4, 5}, 4},
	{[]int{4, 56, 235, 36, 3425, 1, 4, 5}, 4},
	{[]int{4, 56, 235, 6, 3425, 1, 4, 5}, 4},
	{[]int{4, 56, 235, 36, 3425, 1, 4, 5}, 4},
	{[]int{4, 56, 235, 236, 3425, 1, 4, 5}, 4},
	{[]int{4, 56, 235, 36, 3425, 1, 4, 5}, 4},
	{[]int{4, 56, 235, 336, 3425, 1, 4, 5}, 4},
	{[]int{4, 56, 235, 361, 3425, 1, 4, 5}, 4},
	{[]int{4, 56, 235, 36, 43425, 1, 4, 5}, 4},
	{[]int{4, 516, 235, 36, 3425, 1, 4, 5}, 4},
	{[]int{4, 56, 2435, 36, 3425, 1, 4, 5}, 4},
	{[]int{-1, 0, 1, 2, -1, -4}, 4},
}

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		println("初始化")

		// rs := findKthLargest(tests[k1].L1, tests[k1].K)
		rs := 1
		heap := tests[k1].L1[:4:5]

		fmt.Println("结果：", rs, tests[k1].L1, heap, cap(heap))

	}

}

func BenchmarkBubb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k1 := range tests {

			findKthLargest(tests[k1].L1, tests[k1].K)

			// fmt.Println("结果：", rs)

		}

	}
}
func BenchmarkQuick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k1 := range tests {

			_ = findKthLargestV2(tests[k1].L1, tests[k1].K)

			// fmt.Println("结果：", rs)

		}
	}
}

func findKthLargest(nums []int, k int) int {

	// 	快排用了分治思想,左右下标换位置
	// quickSoft(nums, 0, len(nums)-1, k)
	bubblingSoft(nums, k)
	// fmt.Println(nums, k)
	return nums[len(nums)-k]

}
func bubblingSoft(nums []int, k int) {
	l := len(nums)
	x := 0
	for i := l - 1; i >= l-k; i-- {
		for j := i - 1; j >= 0; j-- {
			x++
			// println(i, j, x, k)
			if nums[j] > nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
}

func quickSoft(nums []int, m, n, max int) {
	if n <= m {
		return
	}
	o := m // 初始值
	r := n
	for m < n {
		for nums[o] <= nums[n] && m < n {
			n--
		}
		for nums[o] >= nums[m] && m < n {
			m++
		}
		nums[m], nums[n] = nums[n], nums[m]
	}
	nums[m], nums[o] = nums[o], nums[m]

	quickSoft(nums, o, m, max)
	quickSoft(nums, m+1, r, max)
}

// 堆排序
func findKthLargestV2(nums []int, k int) int {
	heap := nums[:k:k]
	for i := (k - 1) / 2; i >= 0; i-- {
		heapAdjust(heap, i, k-1)
	}
	for i := k; i < len(nums); i++ {
		if heap[0] < nums[i] {
			heap[0] = nums[i]
			heapAdjust(heap, 0, k-1)
		}
	}
	return heap[0]
}

func heapAdjust(nums []int, root, end int) {
	for i := root*2 + 1; i <= end; i = i*2 + 1 {
		if i < end && nums[i] > nums[i+1] {
			i++
		}
		if nums[root] <= nums[i] {
			break
		}
		nums[root], nums[i] = nums[i], nums[root]
		root = i
	}
}
