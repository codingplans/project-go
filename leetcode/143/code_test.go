package _32

import (
	"fmt"
	"testing"
)

//  2021。5.1 重新再做一遍，第一遍卡太久了； 思路： 首先 快慢找出后一半链表，再倒叙链表，再插入

// 143. 重排链表
// 给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
// 将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
// 示例 1:
//
// 给定链表 1->2->3->4, 重新排列为 1->4->2->3.
// 示例 2:
//
// 给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

type test struct {
	L1 []int
}

var tests = []test{
	{[]int{1, 2, 3, 4, 5}},
	{[]int{5, 4, 3, 2, 1}},
	{[]int{-2, 0, 1, 1, 2}},
	{[]int{0, 0, 0}},
	{[]int{0, 0, 0, 0}},
	{[]int{-1, 0, 1, 2, -1, -4}},
}

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		println("初始化")

		pre := &ListNode{Val: 1}
		l := pre
		for k := range tests[k1].L1 {
			l.Next = &ListNode{Val: tests[k1].L1[k]}
			l = l.Next
		}
		pre = pre.Next

		// reorderList(pre)
		reorderListV2(pre)
		Traverse(pre)

	}

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

	// fmt.Println("Traverse OK!")
}
func NewNode(value int, next *ListNode) *ListNode {
	var n ListNode
	n.Val = value
	n.Next = next
	return &n
}

// 定义一个 链表结构
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(head *ListNode, v int) *ListNode {
	if head == nil {
		return &ListNode{Val: v}
	}
	n := &ListNode{Val: v}
	n.Next = head
	return n
}
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func mergeList(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head
	i := 0
	for l1 != nil && l2 != nil {
		if i%2 == 0 {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}

		i++
		p = p.Next

		// l1Tmp = l1.Next
		// l2Tmp = l2.Next
		//
		// fmt.Println("#########")
		// Traverse(l1)
		// fmt.Println("#########")
		// Traverse(l2)
		//
		// l1.Next = l2
		// // 下一步
		// l1.Next = l1Tmp
		// l2 = l2Tmp
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return head.Next
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	// 寻找中间结点
	p1 := head
	p2 := head
	// 先找到中部节点
	p1 = middleNode(head)

	// 反转链表后半部分  1->2->3->4->5->6 to 1->2->3->6->5->4
	preMiddle := p1
	preCurrent := p1.Next

	Traverse(preCurrent)
	fmt.Println("bbbbbbb")
	for preCurrent.Next != nil {
		current := preCurrent.Next
		preCurrent.Next = current.Next
		current.Next = preMiddle.Next
		preMiddle.Next = current
	}

	// 重新拼接链表  1->2->3->6->5->4 to 1->6->2->5->3->4
	p1 = head
	p2 = preMiddle.Next
	for p1 != preMiddle {
		// 进一位
		preMiddle.Next = p2.Next
		p2.Next = p1.Next
		p1.Next = p2
		p1 = p2.Next
		p2 = preMiddle.Next
	}
}

func reorderListV2(head *ListNode) {

	if head == nil {
		return
	}
	// 寻找中间结点
	p1 := head
	p2 := head
	// 先找到中部节点
	p1 = middleNode(head)

	// 头结点
	preMiddle := p1
	preCurrent := p1.Next

	Traverse(preCurrent)
	fmt.Println("bbbbbbb")
	for preCurrent.Next != nil {
		current := preCurrent.Next
		preCurrent.Next = current.Next
		current.Next = preMiddle.Next
		preMiddle.Next = current
	}

	// 重新拼接链表  1->2->3->6->5->4 to 1->6->2->5->3->4
	p1 = head
	p2 = preMiddle.Next
	for p1 != preMiddle {
		// 进一位
		preMiddle.Next = p2.Next
		p2.Next = p1.Next
		p1.Next = p2
		p1 = p2.Next
		p2 = preMiddle.Next
	}

}
