package main

import log "github.com/sirupsen/logrus"

// Definition for singly-linked list.

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func main() {

	levelOrderBottomDfs()
	// h := &ListNode2{
	// 	1, nil,
	// }
	// for i := 1; i < 10; i++ {
	// 	h.addCycle(&ListNode2{
	// 		i, nil,
	// 	})
	// }
	//
	// // aa := hasCycle()
	// aa := detectCycle(h)
	// log.Info(aa)
}

func levelOrderBottomDfs(root *TreeNode) [][]int {
	// 深度优先遍历(dfs)
	result := make([][]int, 0)
	level := 0
	if root == nil {
		return result
	}

	orderBottom(root, &result, level)

	// 数组翻转
	resultLength := len(result)
	left := 0
	right := resultLength - 1
	for left < right {
		temp := result[left]
		result[left] = result[right]
		result[right] = temp

		left++
		right--
	}

	return result
}

func orderBottom(root *TreeNode, result *[][]int, level int) {
	if root == nil {
		return
	}

	if len(*result) > level {
		(*result)[level] = append((*result)[level], root.Val)
	} else {
		*result = append(*result, []int{root.Val})
	}

	orderBottom(root.Left, result, level+1)
	orderBottom(root.Right, result, level+1)
}

func levelOrderBottom(root *TreeNode) [][]int {
	// 广度优先搜索(bfs)
	result := make([][]int, 0)
	level := 0
	if root == nil {
		return result
	}

	// 初始化队列
	list := []*TreeNode{root}
	length := 1 // 队列长度(即当前层节点数)
	for length > 0 {
		// 从队列中取出当前层
		for i := 0; i < length; i++ {
			// 出队
			node := list[0]
			list = list[1:]

			// 值放入result
			if len(result) > level {
				result[level] = append(result[level], node.Val)
			} else {
				result = append(result, []int{node.Val})
			}

			// 下一层入队
			if node.Left != nil {
				list = append(list, node.Left)
			}
			if node.Right != nil {
				list = append(list, node.Right)
			}
		}

		length = len(list)
		level++
	}

	// 数组翻转
	resultLength := len(result)
	left := 0
	right := resultLength - 1
	for left < right {
		temp := result[left]
		result[left] = result[right]
		result[right] = temp

		left++
		right--
	}

	return result
}

func (l *ListNode2) addCycle(head *ListNode2) {

	for l.Next != nil {
		l = l.Next
	}
	l.Next = head
}

func hasCycle(head *ListNode2) bool {
	if head == nil || head.Next == nil {
		return false
	}
	s := head
	k := head
	for k.Next != nil && k.Next.Next != nil {
		log.Info(k.Val)
		k = k.Next.Next
		s = s.Next
		if s == k {
			return true
		}
	}
	return false
}

func detectCycle(head *ListNode2) *ListNode2 {

	if head == nil || head.Next == nil {
		// return "no cycle"
	}

	return head
}
