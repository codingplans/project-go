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

var tests = []test{
	{
		IntEs: []int{1, 2, 2, 3, 4, 4, 3},
	},
	{
		IntEs: []int{1, 2, 2, structures.NULL, 3, structures.NULL, 3},
	},
}

// 101. 对称二叉树
// 给定一个二叉树，检查它是否是镜像对称的。
//
//
//
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//    1
//   / \
//  2   2
// / \ / \
// 3  4 4  3
//
//
// 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//    1
//   / \
//  2   2
//   \   \
//   3    3
//

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		fmt.Println("初始化")
		list := structures.Ints2TreeNode(tests[k1].IntEs)
		pre := isSymmetric(list)
		fmt.Println("结果：", pre)
	}

}

func isSymmetric(root *structures.TreeNode) bool {

	return digui(root, root)

}

func digui(l, r *structures.TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}

	return l.Val == r.Val && digui(l.Left, r.Right) && digui(l.Right, r.Left)

}
