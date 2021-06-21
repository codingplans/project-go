package _00_init_code

import (
	"fmt"
	"strconv"
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
	{IntEs: []int{1, 11, 19, 111, 123, 1, 11, 39, 111, 123, 9999, 9, 89, 349}},
	{IntEs: []int{0, 0}},
	{IntEs: []int{1, 2}},
	{IntEs: []int{2, 1}},
	{IntEs: []int{0, 0, 1}},
}

// 179. 最大数
// 给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
//
// 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
//
//
//
// 示例 1：
//
// 输入：nums = [10,2]
// 输出："210"
// 示例 2：
//
// 输入：nums = [3,30,34,5,9]
// 输出："9534330"
// 示例 3：
//
// 输入：nums = [1]
// 输出："1"
// 示例 4：
//
// 输入：nums = [10]
// 输出："10"

func Test_upToDayUp(t *testing.T) {
	for _, v := range tests {
		fmt.Println("初始化")

		fmt.Println(v.IntEs)
		pre := largestNumber(v.IntEs)
		fmt.Println("结果：", pre)
	}
}
func largestNumber(nums []int) string {

	i, j := 0, 0

	for i = 0; i < len(nums); i++ {
		for j = i + 1; j < len(nums); j++ {
			sum := fmt.Sprintf("%d%d", nums[i], nums[j])
			sum2 := fmt.Sprintf("%d%d", nums[j], nums[i])
			// sum := string(nums[i]) + string(nums[j])
			// sum2 := string(nums[j]) + string(nums[i])
			s1, _ := strconv.ParseInt(sum, 10, 64)
			s2, _ := strconv.ParseInt(sum2, 10, 64)
			if s1 < s2 {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	str := ""
	if nums[0] == 0 {
		return "0"
	}
	for k := range nums {

		str += strconv.FormatInt(int64(nums[k]), 10)

	}
	// fmt.Println(nums)
	return str

}
