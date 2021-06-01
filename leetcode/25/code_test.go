package _32

import (
	"fmt"
	"testing"
)

// 25. K 个一组翻转链表
// 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。
//
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 进阶：
//
// 你可以设计一个只使用常数额外空间的算法来解决此问题吗？
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
//
// 示例 1：
//
//
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[2,1,4,3,5]
// 示例 2：
//
//
// 输入：head = [1,2,3,4,5], k = 3
// 输出：[3,2,1,4,5]
// 示例 3：
//
// 输入：head = [1,2,3,4,5], k = 1
// 输出：[1,2,3,4,5]
// 示例 4：
//
// 输入：head = [1], k = 1
// 输出：[1]
// 提示：
//
// 列表中节点的数量在范围 sz 内
// 1 <= sz <= 5000
// 0 <= Node.val <= 1000
// 1 <= k <= sz

// darren思考：
// K个一组翻转链表，考察链表翻转，链表合并，链表基础不能快速查找，只能遍历查找；
//
// 滑动窗口做法 用时 4ms，内存3.9M
// 从最开始的傻瓜做法实现，用 8ms， 后来改用滑动窗口法则，减少遍历次数，因为是链表（链式调用，每个节点都存下个节点地址）所以用一个指针变量来循环复用，遇到最后不是完整 k 个的情况，做链表翻转，其余情况直接 merge 到 list 中；其中巧妙用了栈思想，直接就是倒叙，不符合情况才翻转（正序）。

type test struct {
	V []int
	k int
}

var tests = []test{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3},
	{[]int{1, 2, 3}, 2},
	{[]int{1, 2, 3, 4, 5}, 3},
	{[]int{0, 1, 0, 4, 0, 2, 1, 0, 4}, 2},
	{[]int{0, 1}, 2},
	{[]int{0}, 1},
}

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		var list *ListNode
		for l := range tests[k1].V {
			list = NewNode(tests[k1].V[l], list)
		}
		println("初始化")
		travlist(list)
		// nl := reverseKGroup(list, tests[k1].k)
		nl := reverseKGroupV2(list, tests[k1].k)
		println("翻转后")

		travlist(nl)

	}

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 遍历 链表
func travlist(head *ListNode) {

	if head == nil {
		return
	}
	point := head
	fmt.Println(point.Val)

	for point.Next != nil {
		point = point.Next
		fmt.Println(point.Val)
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 优化第二版：
// 用栈思想 先每 k 个入栈 ，取出到新 list，

func reverseKGroupV2(head *ListNode, k int) *ListNode {
	Bhead := head
	var news *ListNode
	for Bhead != nil {
		i := k
		var temp *ListNode
		for i > 0 && Bhead != nil {
			temp = NewNode(Bhead.Val, temp)
			Bhead = Bhead.Next
			i--
		}

		// 说明满足k 个
		if i != 0 {
			// 读栈到链表
			temp = reverse(temp)
		}
		news = merge(news, temp)
	}
	return news
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	Bhead := head

	var arr []*ListNode
	for Bhead != nil {
		var news *ListNode
		var i int = 0
		for i < k && Bhead != nil {
			news = NewNode(Bhead.Val, news)
			i++
			Bhead = Bhead.Next
		}
		// 不是最后一个节点 都翻转
		if i != k {
			news = reverse(news)

		}
		arr = append(arr, news)
	}
	var news *ListNode
	for l := range arr {
		if news == nil {
			news = arr[l]
			continue
		}
		news = merge(news, arr[l])
	}
	return news
}

func merge(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	c := a
	for c.Next != nil {
		c = c.Next
	}
	c.Next = b
	return a
}

func NewNode(value int, next *ListNode) *ListNode {
	if next == nil {
		return &ListNode{Val: value}
	}
	n := &ListNode{Val: value}
	n.Next = next
	return n
}

// 翻转单链表
func reverse(list *ListNode) *ListNode {
	var rList *ListNode = list
	var news *ListNode
	for rList != nil {
		next := rList.Next
		rList.Next = news
		news = rList
		rList = next
	}
	return news
}
