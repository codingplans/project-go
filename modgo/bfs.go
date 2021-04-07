package main

import (
	"container/list"
	"fmt"
	"sync"
)

type Node struct {
	v      string
	lchild *Node
	rchild *Node
}

// 栈信息
type Stack struct {
	list *list.List
	lock *sync.RWMutex
}

func main() {

	data := []string{"A", "B", "D", "#", "#", "E", "#", "#", "C", "F", "#", "#", "G", "#", "#"}

	t := new(Node)
	treeNodeCreate(t, data)
	depthSearch(t)
}
func NewStack() *Stack {
	list := list.New()
	l := &sync.RWMutex{}
	return &Stack{list, l}
}
func (stack *Stack) Push(value interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.list.PushBack(value)
}
func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}
func (stack *Stack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

var index, kk int

func treeNodeCreate(tree *Node, data []string) {
	e := data[index]
	index++
	// var root *Node
	if e == "#" {
		tree = nil
	} else {
		tree = new(Node)
		tree.v = e
		println(tree.v)
		treeNodeCreate(tree.lchild, data)
		treeNodeCreate(tree.rchild, data)
	}
}

func depthSearch(tree *Node) {
	// L := new(Stack)
	// L.Push(tree)
	// for !L.Empty() {
	// 	s := L.Pop()
	// 	if tree.rchild != nil {
	// 		L.Push(tree.rchild)
	// 	}
	// 	if tree.lchild != nil {
	// 		L.Push(tree.lchild)
	// 	}
	//
	// }
	// temp := tree
	if tree == nil {
		return
	}
	fmt.Printf("%+v3333", tree)

	if tree.lchild != nil {
		println(tree.v, 3333)

		depthSearch(tree.lchild)
	}
	if tree.rchild != nil {
		println(tree.v, 444)

		depthSearch(tree.rchild)
	}
	kk++
}

func breadSearch(tree *Node) {

	for tree != nil {

	}

}
