package _00_init_code

import (
	"fmt"
	"github.com/Darrenzzy/person-go/structures"
	"testing"
)

// 103. 二叉树的锯齿形层序遍历
// 给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
// 例如：
// 给定二叉树 [3,9,20,null,null,15,7],
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
// 返回锯齿形层序遍历如下：
//
// [
//  [3],
//  [20,9],
//  [15,7]
// ]

type test struct {
	IntEs []int
	Lists []int
	K     int
}

var tests = []test{
	{IntEs: []int{1, 2, 3, 4, 5}},
	{IntEs: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	{IntEs: []int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
}

//

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		fmt.Println("初始化")
		tree := structures.Ints2TreeNode(tests[k1].IntEs)
		pre := zigzagLevelOrder(tree)
		fmt.Println("结果：", pre)
	}
}

var arrsV2 [][]int

func zigzagLevelOrder(root *structures.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	// 	开始 Bfs
	// 初始化一个队列
	// var list []*structures.TreeNode
	// list = append(list, root)
	// return travelTreeBfs(list)
	arrsV2 = [][]int{}
	travelTreeDfs(root, 0)
	return arrsV2
}

func zigzagLevelOrderV2(root *structures.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	// 	开始 Bfs
	// 初始化一个队列
	var list []*structures.TreeNode
	list = append(list, root)
	return travelTreeBfs(list)
}

func travelTreeBfs(list []*structures.TreeNode) [][]int {
	var arrs [][]int
	isRevel := false

	for len(list) > 0 {
		i := 0
		l := len(list)
		// 初始化一个队列
		if isRevel {
			isRevel = false
		} else {
			isRevel = true
		}
		arr := []int{}
		for l > i {
			m := i
			if !isRevel {
				m = l - 1 - i
			}
			// println(m, i, isRevel)
			t := list[m]
			arr = append(arr, t.Val)

			node := list[i]
			if node.Left != nil {
				list = append(list, node.Left)
			}
			if node.Right != nil {
				list = append(list, node.Right)
			}
			i++
			// fmt.Println(i, node.Val)
		}
		arrs = append(arrs, arr)
		list = list[i:]
	}
	return arrs
}

// 递归 深度遍历  这个对性能影响大，不断扩容，是不划算的
func travelTreeDfs(root *structures.TreeNode, level int) {
	var arr []int
	if len(arrsV2) <= level {
		arrsV2 = append(arrsV2, []int{})
	}
	node := root
	arr = arrsV2[level]
	if level%2 == 0 {
		arr = append(arr, node.Val)

	} else {
		arr = append([]int{node.Val}, arr...)

	}

	// fmt.Println(node.Val, arr, arrsV2)
	// 归位
	arrsV2[level] = arr
	if node.Left != nil {
		travelTreeDfs(node.Left, level+1)
	}

	if node.Right != nil {
		travelTreeDfs(node.Right, level+1)
	}

}
