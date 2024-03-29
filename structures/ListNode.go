package structures

import (
	"fmt"
)

// ListNode 是链接节点
// 这个不能复制到*_test.go文件中。会导致Travis失败
type ListNode struct {
	Val  int
	Next *ListNode
}

// List2Ints convert List to []int
func List2Ints(head *ListNode) []int {
	// 链条深度限制，链条深度超出此限制，会 panic
	limit := 100

	times := 0

	res := []int{}
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// Ints2List convert []int to List
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

// GetNodeWith returns the first node with val
func (l *ListNode) GetNodeWith(val int) *ListNode {
	res := l
	for res != nil {
		if res.Val == val {
			break
		}
		res = res.Next
	}
	return res
}

// Ints2ListWithCycle returns a list whose tail point to pos-indexed node
// head's index is 0
// if pos = -1, no cycle
func Ints2ListWithCycle(nums []int, pos int) *ListNode {
	head := Ints2List(nums)
	if pos == -1 {
		return head
	}
	c := head
	for pos > 0 {
		c = c.Next
		pos--
	}
	tail := c
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = c
	return head
}

func Travel(head *ListNode) {
	L := head
	for L != nil {
		fmt.Println(L.Val)
		L = L.Next
	}

}

func MergeList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	l := l1
	for l.Next != nil {
		l = l.Next
	}
	l.Next = l2
	return l1
}

// 翻转链表
func Reverse(pHead *ListNode) *ListNode {
	m := pHead
	// 方法一 ：最好理解
	var newList *ListNode
	for m != nil {
		// 先把下一个全部移送新变量
		nt := m.Next
		// 把反转链顺到主后面
		m.Next = newList
		// 主的是个反转完整链在给 反转链
		newList = m
		// 把临时变量放回主连
		m = nt
	}
	return newList
}
