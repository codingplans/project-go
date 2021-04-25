package main

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value string
}

var DefaultValue string = "#"

func InsertNodeToTree(tree *TreeNode, node *TreeNode) {
	if tree == nil {
		return
	}
	if tree.Value == DefaultValue {
		tree.Value = node.Value
		return
	}
	if node.Value > tree.Value {
		if tree.Right == nil {
			tree.Right = &TreeNode{Value: DefaultValue}
		}
		InsertNodeToTree(tree.Right, node)
	}
	if node.Value < tree.Value {
		if tree.Left == nil {
			tree.Left = &TreeNode{Value: DefaultValue}
		}
		InsertNodeToTree(tree.Left, node)
	}
}

func InitTree(values ...string) (root *TreeNode) {
	rootNode := TreeNode{Value: DefaultValue, Right: nil, Left: nil}
	for _, value := range values {
		node := TreeNode{Value: value}
		InsertNodeToTree(&rootNode, &node)
	}
	return &rootNode
}

func main() {
	treeNode := InitTree("A", "B", "C", "D", "E", "F", "G")
	levelOrderBottoms(treeNode)

}

func levelOrderBottoms(root *TreeNode) [][]string {
	levelOrder := [][]string{}
	if root == nil {
		println(122)
		return levelOrder
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		level := []string{}
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			println(node.Value, 999)
			level = append(level, node.Value)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levelOrder = append(levelOrder, level)
	}
	for i := 0; i < len(levelOrder)/2; i++ {
		levelOrder[i], levelOrder[len(levelOrder)-1-i] = levelOrder[len(levelOrder)-1-i], levelOrder[i]
	}
	return levelOrder
}
