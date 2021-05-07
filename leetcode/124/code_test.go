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

// NULL 方便添加测试数据
var NULL = -1 << 63

var tests = []test{
	{Lists: []int{-10, 9, 20, NULL, NULL, 15, 7}},
	{Lists: []int{1, 2, 3}},
	{Lists: []int{1, -2, 3}},
}

// 124. 二叉树中的最大路径和
// 路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
//
// 路径和 是路径中各节点值的总和。
//
// 给你一个二叉树的根节点 root ，返回其 最大路径和 。
//
//
//
// 示例 1：
//
//
// 输入：root = [1,2,3]
// 输出：6
// 解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
// 示例 2：
//
//
// 输入：root = [-10,9,20,null,null,15,7]
// 输出：42
// 解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
//
//
// 提示：
//
// 树中节点数目范围是 [1, 3 * 104]
// -1000 <= Node.val <= 1000

// 思考：
// 方案一： 递归统计每个节点下最大橘子节点和，

var sums = 0

func maxPathSum(root *structures.TreeNode) int {
	if root == nil {
		return 0
	}
	sums = root.Val
	recursive(root)
	return sums
}

// 深度遍历 dfs 找最大路径和
// 递归方法 recursive
func recursive(node *structures.TreeNode) int {
	if node == nil {
		return 0
	}
	l := max(recursive(node.Left), 0)
	r := max(recursive(node.Right), 0)
	p := node.Val + l + r
	sums = max(p, sums)
	println(sums, p, node.Val)
	return node.Val + max(l, r)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		fmt.Println("初始化", tests[k1].Lists)
		// 快速生成二叉树
		tree := structures.Ints2TreeNode(tests[k1].Lists)
		pre := maxPathSum(tree)
		fmt.Println("结果：", pre)

	}
}
