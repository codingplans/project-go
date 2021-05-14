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

var tests = []test{
	{IntEs: []int{5, 4, 8, 11, structures.NULL, 13, 4, 7, 2, structures.NULL, structures.NULL, 5, 1}, K: 22},
	{IntEs: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, K: 15},
	{IntEs: []int{1, 2, 3}, K: 4},
	{IntEs: []int{1, 2}, K: 3},
}

// 113. 路径总和 II
// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//
// 叶子节点 是指没有子节点的节点。
//
//
//
// 示例 1：
//
//
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// 输出：[[5,4,11,2],[5,8,4,5]]
// 示例 2：
//
//
// 输入：root = [1,2,3], targetSum = 5
// 输出：[]
// 示例 3：
//
// 输入：root = [1,2], targetSum = 0
// 输出：[]
//
//
// 提示：
//
// 树中节点总数在范围 [0, 5000] 内
// -1000 <= Node.val <= 1000
// -1000 <= targetSum <= 1000

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		fmt.Println("初始化")
		root := structures.Ints2TreeNode(tests[k1].IntEs)
		pre := pathSum(root, tests[k1].K)
		fmt.Println("结果：", pre)

	}
}

var arrs [][]int

// 思路： 通过递归调用，每一层级 减掉目标值，到最根部判断是否为 0
// ，也用了分治思想， 分输到下一层去做判断，当前只减法

func pathSum(root *structures.TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	arrs = [][]int{}
	FindPath(root, []int{}, targetSum)

	// fmt.Println(arrs, 2222)

	return arrs
}

func FindPath(root *structures.TreeNode, arr []int, cur int) {
	if root == nil {
		return
		// return nil
	}
	cur -= root.Val
	arr = append(arr, root.Val)
	if cur == 0 && root.Left == nil && root.Right == nil {
		arrs = append(arrs, arr)
		arr = arr[:len(arr)-1]

	}
	if root.Left != nil {
		FindPath(root.Left, arr, cur)
	}
	if root.Right != nil {
		FindPath(root.Right, arr, cur)
	}

}

//
// func BfsFind(root *structures.TreeNode, level int) int {
// 	if root == nil {
// 		return 0
// 	}
//
// 	if len(arrs) <= level {
// 		arrs = append(arrs, []int{
// 			root.Val,
// 		})
//
// 	} else {
// 		arrs[level] = append(arrs[level], root.Val)
// 	}
//
// 	BfsFind(root.Left, level+1)
// 	BfsFind(root.Right, level+1)
//
// 	return 0
// }
//
// func DfsFind(root *structures.TreeNode, level, l int) int {
// 	if root == nil {
// 		println(11, level, l)
// 		return level + 1
// 	}
// 	for len(arrs) <= level {
// 		arrs = append(arrs, []int{})
// 	}
//
// 	// sum := root.Val
// 	arrs[level] = append(arrs[level], root.Val)
// 	arrs[l] = append(arrs[l], root.Val)
// 	level = DfsFind(root.Left, level, l+1)
//
// 	// arrs[level] = append(arrs[level], root.Val)
// 	level = DfsFind(root.Right, level, l+1)
// 	return level
// }
