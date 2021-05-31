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
	{IntEs: []int{1, 2, structures.NULL, structures.NULL, 3, 4, 5, 6, structures.NULL, structures.NULL, structures.NULL, structures.NULL, 7, 8, 9}},
}

// 求直径

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		fmt.Println("初始化")

		fmt.Println(tests[k1].IntEs)
		t1 := structures.Ints2TreeNode(tests[k1].IntEs)

		fmt.Println(dp(t1))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// class Solution:
// def diameterOfBinaryTree(self, root: TreeNode) -> int:
// self.ans = 1
// def depth(node):
// # 访问到空节点了，返回0
// if not node:
// return 0
// # 左儿子为根的子树的深度
// L = depth(node.left)
// # 右儿子为根的子树的深度
// R = depth(node.right)
// # 计算d_node即L+R+1 并更新ans
// self.ans = max(self.ans, L + R + 1)
// # 返回该节点为根的子树的深度
// return max(L, R) + 1
//
// depth(root)
// return self.ans - 1

func dp(tree *structures.TreeNode) int {
	var l, r int
	if tree.Left != nil {
		l = dp(tree.Left)
	}
	if tree.Right != nil {
		r = dp(tree.Right)
	}
	println(l, r)
	return max(l, r) + 1
}
