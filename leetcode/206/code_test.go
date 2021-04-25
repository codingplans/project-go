package _32

import (
	"fmt"
	"testing"
)

// 206. 反转链表
// 反转一个单链表。

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
			NewNode(3,
				NewNode(4,
					NewNode(5,
						NewNode(6,
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
		rs := reverseList(tests[k].ListNode)
		Traverse(rs)
	}
}

func reverseList(head *ListNode) *ListNode {
	var news, next *ListNode
	now := head

	for now != nil {
		// 当前的给零时存起来，放置丢失
		next = now.Next
		// 每个节点的下一个指向新连
		now.Next = news
		// 当前节点复制新连
		news = now
		// 临时变量还原给当前节点 （当前节点向后移动）
		now = next

	}

	return news
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
