package _32

import (
	"testing"
)

// 2. 两数相加
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//
//
// 示例 1：
//
//
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
// 示例 2：
//
// 输入：l1 = [0], l2 = [0]
// 输出：[0]
// 示例 3：
//
// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// 输出：[8,9,9,9,0,0,0,1]
//
//
// 提示：
//
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零

// 思考：考察链表遍历 ，翻转，整形求和，
// 流程：遍历求和， 拆分数字入链表

type test struct {
	L1, L2 []int
}

var tests = []test{
	{[]int{2, 4, 3}, []int{5, 6, 4}}, // 8 0 7
	{[]int{0, 1}, []int{0, 1}},
	{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}}, // [8,9,9,9,0,0,0,1]
	{[]int{9, 9, 9, 9}, []int{9, 9, 9, 9, 9, 9, 9}}, // [8,9,9,9,0,0,0,1]
	{[]int{0}, []int{0, 1}},
	{[]int{0}, []int{0}},
}

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		var list, list2 *ListNode
		for l := range tests[k1].L1 {
			list = NewNode(tests[k1].L1[l], list)
		}
		for l := range tests[k1].L2 {
			list2 = NewNode(tests[k1].L2[l], list2)
		}
		println("初始化")
		travlist(list)
		println("初始化2")
		travlist(list2)

		// nl := addTwoNumbers(list, list2)
		nl := addTwoNumbersV2(list, list2)
		// nl := addTwoNumbersV3(list, list2)
		println("翻转后")
		travlist(nl)

	}

}
func NewNode(v int, list *ListNode) *ListNode {
	if list == nil {
		return &ListNode{Val: v}
	}
	n := &ListNode{Val: v}
	n.Next = list
	return n
}

// 正序添加节点
func NewNodeV2(v int, list *ListNode) *ListNode {
	if list == nil {
		return &ListNode{Val: v}
	}
	l := list
	for l.Next != nil {
		l = l.Next
	}
	l.Next = &ListNode{Val: v}
	return list
}

// 仅生成节点
func NewNodeV3(v int, list *ListNode) *ListNode {
	if list == nil {
		return &ListNode{Val: v}
	}
	l := list
	for l.Next != nil {
		l = l.Next
	}
	l.Next = &ListNode{Val: v}
	return list
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

//  方案 3
func addTwoNumbersV3(l1 *ListNode, l2 *ListNode) *ListNode {
	n := ListNode{}
	s := &n
	var count int
	for l1 != nil || l2 != nil {
		var a1, a2 int
		if l1 != nil {
			a1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			a2 = l2.Val
			l2 = l2.Next
		}

		s.Val = (a1 + a2 + count) % 10
		count = (a1 + a2 + count) / 10
		if l1 == nil && l2 == nil {
			if count > 0 {
				s.Next = &ListNode{Val: count}
			}
			break
		} else {
			s.Next = &ListNode{}
			s = s.Next

		}

	}

	return &n
}

// 方案 2
func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	// 定义一个新链表的表头和指针。新链表用来保存最终的计算结果
	head := &ListNode{Val: 0}
	pre := head
	sum := 0
	for {
		if l1 == nil && l2 == nil && sum == 0 {
			break
		}

		value1 := 0
		value2 := 0
		if l1 != nil {
			value1 = l1.Val
		}
		if l2 != nil {
			value2 = l2.Val
		}

		sum += value1 + value2

		pre.Next = &ListNode{Val: sum % 10}
		pre = pre.Next

		sum = sum / 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return head.Next
}

func revert(list *ListNode) *ListNode {
	if list == nil {
		return nil
	}
	l := list
	var t *ListNode
	for l != nil {
		next := l.Next
		l.Next = t
		t = l
		l = next
	}

	return t
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	list := l1
	list2 := l2

	count := 0
	sum := list.Val + list2.Val
	if sum >= 10 {
		count++
		sum = sum % 10
	}
	n := NewNodeV2(sum, nil)
	for list.Next != nil || list2.Next != nil {
		sum = 0
		if list.Next != nil {
			list = list.Next
			sum += list.Val
		}

		if list2.Next != nil {
			list2 = list2.Next
			sum += list2.Val
		}

		// 上一个大于 10 进一位
		for count > 0 {
			sum++
			count--
		}
		// 大于 10 给下一个 进一位
		if sum >= 10 {
			count++
			sum = sum % 10
		}
		// println(sum, 999)
		n = NewNodeV2(sum, n)
	}
	if count > 0 {
		n = NewNodeV2(1, n)
		count--
	}
	return n

}

func travlist(l *ListNode) {
	if l == nil {
		return
	}
	list := l
	println(list.Val)
	for list.Next != nil {
		list = list.Next
		println(list.Val)
	}
	// println("遍历结束")
}
