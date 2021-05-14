package _00_init_code

import (
	"fmt"
	"github.com/Darrenzzy/testgo/structures"
	"testing"
)

type test struct {
	IntEs []int
	Lists []int
	K     int
}

var tests = []test{
	{
		Lists: []int{9, 4, 5},
		IntEs: []int{1, 4, 5},
	},
	{
		Lists: []int{9, 8, 6, 4, 3, 4, 5},
		IntEs: []int{1, 2, 3, 4, 5},
	},
}

//

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		fmt.Println("初始化")

		node1 := structures.Ints2List(tests[k1].Lists)
		node2 := structures.Ints2List(tests[k1].IntEs)
		node3 := getIntersectionNode(node1, node2)
		Travel(node3)

	}
}

func Travel(h *structures.ListNode) {
	if h == nil {
		return
	}
	c := h
	fmt.Println(c.Val)
	for c.Next != nil {
		c = c.Next
		fmt.Println(c.Val)
	}
}

func getIntersectionNode(headA, headB *structures.ListNode) *structures.ListNode {
	// boundary check
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	// if a & b have different len, then we will stop the loop after second iteration
	for a != b {
		// for the end of first iteration, we just reset the pointer to the head of another linkedlist
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
		fmt.Printf("a = %v b = %v\n", a, b)
		if a != nil && b != nil {
			fmt.Println(a.Val, b.Val)

		}
	}
	return a
}
