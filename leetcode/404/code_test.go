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
	{IntEs: []int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
	{IntEs: []int{3, 4, 5, -7, -6, structures.NULL, structures.NULL, -7, structures.NULL, -5, structures.NULL, structures.NULL, structures.NULL, -4}},
}

// 404. 左叶子之和
// 计算给定二叉树的所有左叶子之和。
//
// 示例：
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
// 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		fmt.Println("初始化")

		tree := structures.Ints2TreeNode(tests[k1].IntEs)
		pre := sumOfLeftLeaves(tree)
		fmt.Println("结果：", pre)

	}

}

var sum int

func sumOfLeftLeaves(root *structures.TreeNode) int {
	if root == nil {
		return 0
	}
	sum = 0

	dep(root, nil)
	return sum

}

func dep(root, root2 *structures.TreeNode) {

	if root.Left != nil {
		dep(root.Left, root)
	}
	if root.Right != nil && (root.Right.Right != nil || root.Right.Left != nil) {
		dep(root.Right, root)

	}
	if root.Left == nil && root.Right == nil && root2 != nil {
		sum += root.Val
	}

}
