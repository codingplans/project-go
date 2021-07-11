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

	{
		Lists: []int{1, 9, 8, 6, 4, 3, 4, 5},
		IntEs: []int{5},
	},
}

//

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		fmt.Println("初始化")

		node1 := structures.Ints2List(tests[k1].Lists)
		node2 := structures.Ints2List(tests[k1].IntEs)
		node3 := getIntersectionNode(node1, node2)
		if node3 != nil {
			fmt.Println(node3.Val)
		}
		// Travel(node3)

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

// 思路：题目没有很明确，这里姑且理解为链表中遇到相同值的节点为公共节点
// 。 所以用了双指针 向后位移，做值比较

func getIntersectionNode(headA, headB *structures.ListNode) *structures.ListNode {
	// boundary check
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	for a != b {
		fmt.Printf("a = %v b = %v\n", a, b)
		if a != nil && b != nil {
			if a.Val == b.Val {
				return a
			}
			// fmt.Println(a.Val, b.Val)

		}
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

	}
	return a
}
