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
	param_2 := obj.Pop()
	println(param_2)

	param_1 := obj.Pop()
	println(param_1)

	param_4 := obj.Empty()
	println(param_4)

}

// 待完成
type MyQueue struct {
	// Q1, Q2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {

}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	return 1

}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return 1
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return true
}
