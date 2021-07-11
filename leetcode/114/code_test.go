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
}

var tests = []test{
	{IntEs: []int{1, 2, 5, 3, 4, structures.NULL, 6}},
}

// 114. 二叉树展开为链表
// 给你二叉树的根结点 root ，请你将它展开为一个单链表：
//
// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。
//
//
// 示例 1：
//
//
// 输入：root = [1,2,5,3,4,null,6]
// 输出：[1,null,2,null,3,null,4,null,5,null,6]
// 示例 2：
//
// 输入：root = []
// 输出：[]
// 示例 3：
//
// 输入：root = [0]
// 输出：[0]
//
//
// 提示：
//
// 树中结点数在范围 [0, 2000] 内
// -100 <= Node.val <= 100

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		fmt.Println("初始化")

		fmt.Println(tests[k1].IntEs)
		t1 := structures.Ints2TreeNode(tests[k1].IntEs)
		Traval(t1)
		Bfs(t1)
		flatten(t1)
		Traval(t1)
		// Bfs(t1)

	}
}

var queue []*structures.TreeNode

// 思路： 先创建 2 叉树 在先序遍历， 存入 right 分支，然后删除节点
func flatten(root *structures.TreeNode) {
	Dfs(root)
	ttt(root)

}

func ttt(root *structures.TreeNode) {
	for i := 1; i < len(queue); i++ {
		pre, cur := queue[i-1], queue[i]
		pre.Left, pre.Right = nil, cur
	}

}

// 深度遍历存入队列
func Dfs(root *structures.TreeNode) {
	if root == nil {
		return
	}
	queue = append(queue, root)
	if root.Left != nil && root.Right != nil {

	}
	Dfs(root.Left)
	Dfs(root.Right)

}

// 先序遍历
func Traval(root *structures.TreeNode) {
	if root == nil {
		fmt.Println(nil)
		return
	}
	fmt.Println(root.Val)
	Traval(root.Left)
	Traval(root.Right)

}

// 先序遍历 v2
func PreTreeV2(root *structures.TreeNode) {
	if root == nil {
		return
	}
	// 用栈存储 树节点
	stack := []*structures.TreeNode{}

	// 保证节点不空 或者 栈不空
	for root != nil || len(stack) > 0 {
		// 深度遍历节点 到最低端 存入栈空间
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 最后一个节点时在取右边节点再次深度遍历
		if len(stack) > 0 {
			pop := stack[len(stack)-1]
			// 打印先序第一个节点
			println(pop.Val)
			// 最后一个节点的右节点赋值 ，用于下次深度遍历
			root = pop.Right
			// pop 出去该节点，
			stack = stack[:len(stack)-1]
		}
	}

}

// 保持手感
func Bfs(root *structures.TreeNode) {

	queue := []*structures.TreeNode{root}
	for len(queue) != 0 {
		tree := queue[0]
		queue = queue[1:]
		if tree.Left != nil {
			queue = append(queue, tree.Left)
		}
		if tree.Right != nil {
			queue = append(queue, tree.Right)
		}
		fmt.Println(tree.Val)

	}

}
