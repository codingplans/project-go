package _32

import "testing"

// 请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通队列的全部四种操作（push、top、pop 和 empty）。
//
// 实现 MyStack 类：
//
// void push(int x) 将元素 x 压入栈顶。
// int pop() 移除并返回栈顶元素。
// int top() 返回栈顶元素。
// boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。
//
//
// 注意：
//
// 你只能使用队列的基本操作 —— 也就是 push to back、peek/pop from front、size 和 is empty 这些操作。
// 你所使用的语言也许不支持队列。 你可以使用 list （列表）或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
//
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/implement-stack-using-queues
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func Test_232(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	obj.Push(5)
	param_2 := obj.Pop()
	println(param_2)

	param_1 := obj.Pop()
	println(param_1)
	param_3 := obj.Peek()
	println(param_3)

	param_4 := obj.Empty()
	println(param_4)
	println(obj.Pop())
	println(obj.Pop())
	println(obj.Pop())
	println(obj.Pop())
	println(obj.Pop())

}

// 待完成
type MyQueue struct {
	Q1, Q2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{Q1: []int{}, Q2: []int{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Q1 = append(this.Q1, x)
}

// 出栈 利用互搏，遍历给 q2， 直到 q1 最后一个元素，即是栈顶元素，然后在赋值回去
func (this *MyQueue) Pop() int {
	if len(this.Q1) == 0 {
		return -1
	}
	if len(this.Q1) > 1 {
		for len(this.Q1) != 1 {
			this.Q2 = append(this.Q2, this.Q1[0])
			this.Q1 = this.Q1[1:]
		}
	}
	a := this.Q1[0]
	this.Q1 = this.Q2
	this.Q2 = nil
	return a
}

// 同出栈原理， 当前要全部赋值回去
func (this *MyQueue) Peek() int {
	if len(this.Q1) == 0 {
		return 0
	}

	if len(this.Q1) > 1 {
		for len(this.Q1) != 1 {
			this.Q2 = append(this.Q2, this.Q1[0])
			this.Q1 = this.Q1[1:]
		}
	}
	a := this.Q1[0]
	this.Q2 = append(this.Q2, this.Q1[0])
	this.Q1 = this.Q2
	this.Q2 = nil

	return a
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.Q1) == 0 {
		return true

	}
	return false
}
