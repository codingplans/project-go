package _00_init_code

import (
	"fmt"
	"github.com/Darrenzzy/person-go/structures"
	"testing"
)

type test struct {
	IntEs []int
	Lists []int
	K     int
	Str   string
}

var tests = []test{
	{IntEs: []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 4, 5, 5}},
	{IntEs: []int{1, 2, 3, 3, 4, 4, 5}},
}

// 83. 删除排序链表中的重复元素
// 存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。
//
// 返回同样按升序排列的结果链表。
//
//
//
// 示例 1：
//
//
// 输入：head = [1,2,3,3,4,4,5]
// 输出：[1,2,5]
// 示例 2：
//
//
// 输入：head = [1,1,1,2,3]
// 输出：[2,3]
//
//
// 提示：
//
// 链表中节点数目在范围 [0, 300] 内
// -100 <= Node.val <= 100
// 题目数据保证链表已经按升序排列

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		fmt.Println("初始化")
		list := structures.Ints2List(tests[k1].IntEs)
		Travel(list)
		// list = deleteDuplicates2(list)
		list = deleteDuplicates(list)
		fmt.Println(111111)
		Travel(list)
	}

}

func Travel(node *structures.ListNode) {
	l1 := node
	for l1 != nil {
		fmt.Println(l1.Val)
		l1 = l1.Next
	}
}

// 思考：很简单！
func deleteDuplicates2(head *structures.ListNode) *structures.ListNode {
	pre := &structures.ListNode{Next: head}
	head = pre

	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			head.Next = head.Next.Next
			continue
		}
		head = head.Next
	}

	return pre.Next

}

func deleteDuplicates(head *structures.ListNode) *structures.ListNode {
	pre := head

	for pre.Next != nil {
		if pre.Val == pre.Next.Val {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}
	return head
}
