package _32

import (
	"fmt"
	"testing"
)

// 234. 回文链表
//
// 请判断一个链表是否为回文链表。
//
// 示例 1:
//
// 输入: 1->2
// 输出: false
// 示例 2:
//
// 输入: 1->2->2->1
// 输出: true
// 进阶：
// 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

type test struct {
	*ListNode
}

// 定义一个 链表结构
type ListNode struct {
	Val  int
	Next *ListNode
}

var tests = []test{
	{NewNode(1,
		NewNode(2,
			NewNode(1,
				NewNode(1,
					NewNode(1,
						NewNode(1,
							nil))))))},

	{NewNode(1,
		NewNode(2,
			NewNode(2,
				NewNode(1,
					NewNode(2,
						NewNode(2,
							NewNode(1, nil)))))))},
	{NewNode(1,
		NewNode(2,
			NewNode(2,
				NewNode(3,
					NewNode(2,
						NewNode(1,
							nil))))))},
	{NewNode(1, nil)},
}

func Test_upToDayUp(t *testing.T) {
	for k := range tests {
		rs := isPalindrome(tests[k].ListNode)
		println(rs)
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {

}

func isPalindromeV2(head *ListNode) bool {
	temp := head

	// if temp.Next == nil {
	// 	return true
	// }
	a := []int{}
	for temp.Next != nil {
		a = append(a, temp.Val)
		temp = temp.Next
		if temp.Next == nil {
			a = append(a, temp.Val)
		}

	}
	// fmt.Println(a, 999)

	l := len(a)
	for k := 0; k < l/2; k++ {
		if l == 1 || a[k] != a[l-1-k] {
			// fmt.Println(a[k], a[l-1-k], k, l-1-k)
			return false
		}
	}
	return true
}

// 添加 头结点，数据
func Add(head *ListNode, data int) {
	point := head // 临时指针
	for point.Next != nil {
		point = point.Next // 移位
	}

	var node ListNode // 新节点

	point.Next = &node // 赋值
	node.Val = data
}

func NewNode(value int, next *ListNode) *ListNode {
	var n ListNode
	n.Val = value
	n.Next = next
	return &n
}

// 遍历 头结点
func Traverse(head *ListNode) {
	point := head
	// fmt.Println(head.Data, 999)

	for point.Next != nil {
		fmt.Println(point.Val)
		point = point.Next
	}
	fmt.Println(point.Val)

	fmt.Println("Traverse OK!")
}
