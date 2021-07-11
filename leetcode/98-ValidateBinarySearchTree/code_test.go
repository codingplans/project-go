package _00_init_code

import (
	"fmt"
	"github.com/Darrenzzy/person-go/structures"
	"math"
	"testing"
)

type test struct {
	IntEs []int
	Lists []int
	K     int
	K2    int
	Str   string
	Str2  string
}

var tests = []test{
	{IntEs: []int{5, 1, 4, structures.NULL, structures.NULL, 3, 6}},
	{IntEs: []int{5, 1, 6, structures.NULL, structures.NULL, 5, 6}},
	{IntEs: []int{5, 1, 6, structures.NULL, structures.NULL, 5, 6}},
	{IntEs: []int{2, 1, 3}},
}

//

func Test_upToDayUp(t *testing.T) {
	for _, v := range tests {
		fmt.Println("初始化")
		tree := structures.Ints2TreeNode(v.IntEs)
		pre := isValidBST(tree)
		fmt.Println("结果：", pre)
	}

}

func isValidBST(root *structures.TreeNode) bool {
	if root == nil {
		return true
	}

	// 	大小边界测速方法：

	return ValidBfs(root, math.MinInt64, math.MaxInt64)
}

func ValidBfs(root *structures.TreeNode, l, r int) bool {
	if root.Val <= l || root.Val >= r {
		return false
	}

	return ValidBfs(root.Left, l, root.Val) && ValidBfs(root.Right, root.Val, r)
}
