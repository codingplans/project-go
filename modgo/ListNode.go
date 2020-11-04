package main

// 定义一个 链表结构
type ListNode struct {
	Val  int       `json:"val"`
	Next *ListNode `json:"next"`
}

func main() {
	var head = new(ListNode)
	head.Val = 0
	var tail *ListNode
	tail = head // tail用于记录头结点的地址，刚开始tail的的指针指向头结点
	for i := 1; i < 10; i++ {
		t := i
		if i == 8 {
			t = 5
		}
		var node = ListNode{Val: t}
		node.Next = tail // 将新插入的node的next指向头结点
		tail = &node     // 重新赋值头结点
	}

	// EachNodes(tail)
	aa := reverseList(tail)
	EachNodes(aa)

}

func IsClo(head *ListNode) (b bool) {

	f := head
	s := head

	for f != nil {
		if s.Val == f.Val {

		}

	}

	return b

}

// 遍历链表
func EachNodes(head *ListNode) {

	t := head
	for t != nil {
		println(t.Val)
		t = t.Next
	}
}

func reverseList(head *ListNode) *ListNode {

	var newp *ListNode
	cur := head

	for cur != nil {
		// t := cur.Next
		// cur.Next = newp
		// newp = cur
		// cur = t

		cur.Next, cur, newp = newp, cur.Next, cur
		EachNodes(newp)
		println("\n")
	}
	return newp
}
