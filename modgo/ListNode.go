package main

import (
	"fmt"
	"go.uber.org/zap"
	"testgo/modgo/xzap"
)

// 定义一个 链表结构
type ListNode struct {
	Val  int       `json:"val"`
	Next *ListNode `json:"next"`
	Nest *ListNode `json:"nest"`
	Data Element
}

func main() {
	node := Run()
	// aa := GetLength(&head)
	// aa := RevTraverse(&head)
	aa := DelNode(&node, 1)
	xzap.Info("aa", zap.Any("", aa))
}

func Run() ListNode {
	var head ListNode = ListNode{Data: 0, Nest: nil}
	head.Data = 0
	var nodeArray []Element
	for i := 0; i < 10; i++ {
		nodeArray = append(nodeArray, Element(i+1+i*100))
		Add(&head, nodeArray[i])
	}
	return head
}

func DelNode(head *ListNode, val int) *ListNode {
	cur := head
	pre := &ListNode{}
	for cur != nil {
		if cur.Val == val {
			pre.Nest = cur.Nest
			break
		}
		cur, pre = cur.Nest, cur
	}
	return cur
}

// 反转链表
func RevTraverse(head *ListNode) []int {
	count := 0
	temp := head
	for temp != nil {
		count++
		temp = temp.Nest
	}

	mk := make([]int, count)

	for head != nil {
		mk[count-1] = head.Val
		head = head.Nest
		count--
	}

	return mk

}

// 添加 头结点，数据
func Add(head *ListNode, data Element) {
	point := head // 临时指针
	for point.Nest != nil {
		point = point.Nest // 移位
	}
	var node ListNode  // 新节点
	point.Nest = &node // 赋值
	node.Data = data
	node.Val = int(data)

	head.Data = Element(GetLength(head)) // 打印全部的数据
	head.Val = GetLength(head)           // 打印全部的数据

	if GetLength(head) > 1 {
		Traverse(head)
	}

}

// 遍历 头结点
func Traverse(head *ListNode) {
	point := head
	// fmt.Println(head.Data, 999)

	for point.Nest != nil {
		fmt.Println(point.Data)
		point = point.Nest
	}
	fmt.Println("Traverse OK!")
}

// 获取长度 头结点
func GetLength(head *ListNode) int {
	point := head
	var length int
	for point.Nest != nil {
		length++
		point = point.Nest
	}
	return length
}

func newNode() *ListNode {
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
	return tail
}

// 定义元素类型
type Element int

// 函数接口
type LinkNoder interface {
	Add(head *ListNode, new *ListNode)              // 后面添加
	Delete(head *ListNode, index int)               // 删除指定index位置元素
	Insert(head *ListNode, index int, data Element) // 在指定index位置插入元素
	GetLength(head *ListNode) int                   // 获取长度
	Search(head *ListNode, data Element)            // 查询元素的位置
	GetData(head *ListNode, index int) Element      // 获取指定index位置的元素
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
