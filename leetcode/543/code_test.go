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

// 543. 二叉树的直径
// 给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
//
//
//
// 示例 :
// 给定二叉树
//
//          1
//         / \
//        2   3
//       / \
//      4   5
// 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
//
//
//
// 注意：两结点之间的路径长度是以它们之间边的数目表示。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode structures.TreeNode

var l, r, m int

func diameterOfBinaryTree(root *structures.TreeNode) int {
	dfsFind(root)
	return m

}

func dfsFind(root *structures.TreeNode) int {
	if root == nil {
		return 0
	}
	r = dfsFind(root.Right)
	l = dfsFind(root.Left)
	m = max(l+r+1, m)
	return max(l, r) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		fmt.Println("初始化")

		fmt.Println(tests[k1].IntEs)
	}

	diameterOfBinaryTree(structures.Ints2TreeNode([]int{1, 2, 3, 4, 5}))
	pre := m
	fmt.Println("结果：", pre)

}
