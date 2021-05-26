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
	Str   string
}

var tests = []test{}

// 23. 合并K个升序链表
// 给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [
//  1->4->5,
//  1->3->4,
//  2->6
// ]
// 将它们合并到一个有序链表中得到。
// 1->1->2->3->4->4->5->6
// 示例 2：
//
// 输入：lists = []
// 输出：[]
// 示例 3：
//
// 输入：lists = [[]]
// 输出：[]
//
//
// 提示：
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		fmt.Println("初始化", k1)

		mergeKLists(nil)
	}
	mergeKLists(nil)

	pre := 1
	fmt.Println("结果：", pre)

}

func mergeKLists(lists []*structures.ListNode) *structures.ListNode {
	arr1 := []int{1, 2, 4, 6, 8, 99}
	arr2 := []int{2, 3, 5, 7, 9, 98}
	l1 := structures.Ints2List(arr1)
	l2 := structures.Ints2List(arr2)

	merge2(l1, l2)
	Travel(l1)

	return nil
}

func Travel(head *structures.ListNode) {
	L := head
	for L != nil {
		fmt.Println(L.Val, 22)
		L = L.Next
	}

}

func merge2(l1, l2 *structures.ListNode) *structures.ListNode {

	if l1 == nil {
		return l2

	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = merge2(l1.Next, l2)
		return l1
	}

	l2.Next = merge2(l1, l2.Next)
	return l2

}
