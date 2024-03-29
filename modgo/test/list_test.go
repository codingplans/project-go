package main

import (
	"fmt"
	"github.com/Darrenzzy/person-go/structures"
	"testing"
)

func TestReverse(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6, 7}
	list := structures.Ints2List(arr)
	list = Reverse(list)
	//
	// fmt.Println(structures.List2Ints(list))

	t.Helper()
}

func TestSortList(t *testing.T) {
	arr := []int{1, 3, 2, 5, 4, 6, 8, 7, 9}
	list := structures.Ints2List(arr)

	// TravelList(list)
	aa := sortList(list)
	TravelList(aa)

}

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	if length <= 1 {
		return head
	}

	middleNode := MiddleNode(head)
	cur = middleNode.Next
	middleNode.Next = nil
	middleNode = cur
	TravelList(head)
	fmt.Println(222)
	TravelList(middleNode)
	fmt.Println(3333)

	left := sortList(head)
	right := sortList(middleNode)
	return mergeTwoLists(left, right)
}

func MiddleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1 := head
	p2 := head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return p1
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
