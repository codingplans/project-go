package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

var DefaultValue int = -1024

func InsertNodeToTree(tree *TreeNode, node *TreeNode) {
	if tree == nil {
		return
	}
	if tree.Value == DefaultValue {
		tree.Value = node.Value
		return
	}
	if node.Value > tree.Value {
		if tree.Right == nil {
			tree.Right = &TreeNode{Value: DefaultValue}
		}
		InsertNodeToTree(tree.Right, node)
	}
	if node.Value < tree.Value {
		if tree.Left == nil {
			tree.Left = &TreeNode{Value: DefaultValue}
		}
		InsertNodeToTree(tree.Left, node)
	}
}

func InitTree(values ...int) (root *TreeNode) {
	rootNode := TreeNode{Value: DefaultValue, Right: nil, Left: nil}
	for _, value := range values {
		node := TreeNode{Value: value}
		InsertNodeToTree(&rootNode, &node)
	}
	return &rootNode
}

func main() {
	treeNode := InitTree(5, 4, 6, 8, 9, 7, 1, 3, 2)
	tmp := treeNode
	for tmp != nil {
		fmt.Println(tmp.Value)
		if tmp.Left != nil {
			tmp = tmp.Left
		}
		tmp = tmp.Left
		// tmp = tmp.Right
	}

	// initBinary()
	// binary()
}

//
// type Node struct {
// 	left  *Node
// 	right *Node
// 	data  int
// }
//
// var Nd *Node
//
// // 二叉树
// func binary() {
// 	for Nd != nil {
// 		println(Nd.data, "\n")
// 		Nd = Nd.left
// 	}
//
// }
//
// func initBinary() {
// 	// Nd = &Node{data: 100}
// 	i := 1
// 	for i < 40 {
// 		tm := &Node{data: i + 1}
// 		i++
// 		InsertNodeToTree(Nd, tm)
// 	}
// }
//
// func InsertNodeToTree(tree, node *Node) {
// 	if tree == nil {
// 		tree = node
// 	}
// 	if tree.data > node.data {
// 		if tree.left == nil {
// 			tree.left = &Node{data: 1}
// 		}
// 		InsertNodeToTree(tree.left, node)
// 	}
// 	if tree.data < node.data {
// 		if tree.right == nil {
// 			tree.right = &Node{data: 1}
// 		}
// 		InsertNodeToTree(tree.right, node)
//
// 	}
// }
