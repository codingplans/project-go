package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Darrenzzy/person-go/structures"
)

// 放一些 test 用到的 func  解耦开 test 文件

func Travel(node *structures.TreeNode) {
	// fmt.Println(node.Val, 222)

	if node.Left != nil {
		Travel(node.Left)
	}
	fmt.Println(node.Val, 222)

	if node.Right != nil {
		Travel(node.Right)
	}

}

func TravelList(node *structures.ListNode) {
	l := node
	if node == nil {
		return
	}
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

}

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

// 翻转链表
func Reverse(pHead *structures.ListNode) *structures.ListNode {
	m := pHead
	// 方法一 ：最好理解
	var newList *structures.ListNode
	for m != nil {
		nt := m.Next
		m.Next = newList
		newList = m
		m = nt
		TravelList(newList)
	}

	// 方法二：
	// var pnh *structures.ListNode
	// for m != nil {
	// 	newList = m
	// 	m = m.Next
	// 	TravelList(newList)
	//
	// 	newList.Next = pnh
	// 	pnh = newList
	//
	// 	TravelList(pnh)
	// }
	return newList
}

// 中序二叉树
func CreateTree(node *structures.TreeNode, v int) *structures.TreeNode {
	if node == nil {
		fmt.Println(v)
		return &structures.TreeNode{Val: v}
	}
	if node.Val <= v {
		// n := node.Right
		// cur := CreateTree(node.Right, v)
		// cur.Right = n
		// node.Right = cur
		node.Right = CreateTree(node.Right, v)
	} else {
		// n := node.Left
		// cur := CreateTree(node.Left, v)
		// cur.Left = n
		// node.Left = cur
		node.Left = CreateTree(node.Left, v)
	}
	return node
}

// 堆排序
func BuildHeap(arr []int, lens int) {
	for i := (lens - 1) / 2; i >= 0; i-- {
		Heapify(arr, i, lens)
	}

	lens--
	for i := lens; i >= 1; i-- {
		arr[0], arr[lens] = arr[lens], arr[0]
		lens--
		Heapify(arr, 0, lens)
	}
}

func Heapify(arr []int, n, lens int) {
	for {
		i := 2 * n
		if i > lens { // 保证该节点是非叶子节点
			break
		}
		if i < lens && arr[i+1] > arr[i] { // 选择较大的子节点
			i++
		}
		if arr[n] >= arr[i] { // 没下沉到底就构造好堆了
			break
		}
		arr[n], arr[i] = arr[i], arr[n]

		n = i
	}
}

func QuickSoft(arr []int, start, end int) {

	if start >= end {
		return
	}
	l := start
	r := end
	// fmt.Println(start, end, 999)

	for l < r {
		for arr[r] >= arr[start] && l < r {
			r--
		}
		for arr[l] <= arr[start] && l < r {
			l++
		}
		if arr[l] != arr[r] {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}

	arr[l], arr[start] = arr[start], arr[l]
	// fmt.Println(l, start, r, end)

	QuickSoft(arr, start, l-1)
	QuickSoft(arr, l+1, end)
}

func MergeList(l1, l2 *structures.ListNode) *structures.ListNode {
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
