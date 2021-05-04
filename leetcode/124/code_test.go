package _00_init_code

import (
	"fmt"
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

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sums = 0

func maxPathSum(root *TreeNode) int {

	if root == nil {
		return 0
	}
	sum := recursive(root)
	return sum
}

// 递归方法 recursive
func recursive(node *TreeNode) int {
	if node == nil {
		return 0
	}
	fmt.Println(node.Val)
	if node.Left != nil || node.Right != nil {
		sums += max(recursive(node.Left), recursive(node.Right))
	}
	sums = node.Val + sums
	return sums
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		fmt.Println("初始化")

		for k := range tests[k1].Lists {
			fmt.Println(tests[k1].Lists[k])

		}
		pre := 1
		fmt.Println("结果：", pre)

	}
}

func CreateTree(node *TreeNode, v int) {
	if node == nil {
		node = &TreeNode{
			Val: v,
		}
	}

}
