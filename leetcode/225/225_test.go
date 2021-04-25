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

func Test_225(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	param_3 := obj.Top()
	println(param_3)
	param_2 := obj.Pop()
	println(param_2)

	param_1 := obj.Pop()
	println(param_1)

	param_4 := obj.Empty()
	println(param_4)

}

type MyStack struct {
	queue1, queue2 []int
}

/** Initialize your data structure here. */
func Constructor() (s MyStack) {
	return
}

/** Push element x onto stack. */
func (s *MyStack) Push(x int) {
	// 先添加到队列 2 里，遍历队1到队2里
	s.queue2 = append(s.queue2, x)
	for len(s.queue1) > 0 {
		s.queue2 = append(s.queue2, s.queue1[0])
		s.queue1 = s.queue1[1:]
	}
	// 然后置换队1
	s.queue1, s.queue2 = s.queue2, s.queue1
}

/** Removes the element on top of the stack and returns that element. */
func (s *MyStack) Pop() int {
	v := s.queue1[0]
	s.queue1 = s.queue1[1:]
	return v
}

/** Get the top element. */
func (s *MyStack) Top() int {
	return s.queue1[0]
}

/** Returns whether the stack is empty. */
func (s *MyStack) Empty() bool {
	return len(s.queue1) == 0
}
