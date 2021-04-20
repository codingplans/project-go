package _32

import (
	"testing"
)

// 27. 移除元素
// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
//
// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
//
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
//
//
//
// 说明:
//
// 为什么返回数值是整数，但输出的答案是数组呢?
//
// 请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
//
// 你可以想象内部操作如下:
//
// // nums 是以“引用”方式传递的。也就是说，不对实参作任何拷贝
// int len = removeElement(nums, val);
//
// // 在函数里修改输入数组对于调用者是可见的。
// // 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
// for (int i = 0; i < len; i++) {
//     print(nums[i]);
// }
//
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/remove-element
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type test struct {
	V []int
}

var tests = []test{
	{[]int{1, 5, 9, 1, 5, 9}},
	{[]int{0}},
	{[]int{1, 2, 1, 1}},
	{[]int{0, 1, 2, 2, 3, 2, 2, 2, 2, 20, 4, 2}},
}

func Test_220(t *testing.T) {
	for k := range tests {
		rs := removeElement(tests[k].V, 3)
		println(rs)
	}
}

func removeElement(nums []int, val int) int {
	for k := 0; k < len(nums); {
		if k < len(nums) && nums[k] == val {
			nums = append(nums[:k], nums[k+1:]...)
			continue
		}
		k++
	}
	return len(nums)
}
