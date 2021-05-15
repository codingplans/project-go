package main

import (
	"fmt"
	"github.com/Darrenzzy/testgo/structures"
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

func BuildHeap(arr []int, lens int) {
	for i := lens / 2; i >= 0; i-- {
		Heapify(arr, i, lens)
	}

	lens--
	for i := lens; i >= 0; i-- {
		arr[0], arr[lens] = arr[lens], arr[0]
		lens--
		Heapify(arr, 0, lens)
	}
}

func Heapify(arr []int, n, lens int) {
	k := n
	i := k*2 + 1
	j := k*2 + 2

	if i >= lens {
		return
	}
	if i < lens && arr[i] > arr[k] {
		k = i
	}
	if j < lens && arr[j] > arr[k] {
		k = j
	}
	if n != k {
		arr[n], arr[k] = arr[k], arr[n]
		Heapify(arr, k, lens)
	}

}
