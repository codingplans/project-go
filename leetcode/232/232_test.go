package _32

import "testing"

// 232. 用栈实现队列
// 请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
//
// 实现 MyQueue 类：
//
// void push(int x) 将元素 x 推到队列的末尾
// int pop() 从队列的开头移除并返回元素
// int peek() 返回队列开头的元素
// boolean empty() 如果队列为空，返回 true ；否则，返回 false
//
//
// 说明：
//
// 你只能使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
// 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
//
//
// 进阶：
//
// 你能否实现每个操作均摊时间复杂度为 O(1) 的队列？换句话说，执行 n 个操作的总时间复杂度为 O(n) ，即使其中一个操作可能花费较长时间。
//
//
// 示例：
//
// 输入：
// ["MyQueue", "push", "push", "peek", "pop", "empty"]
// [[], [1], [2], [], [], []]
// 输出：
// [null, null, null, 1, 1, false]
//
// 解释：
// MyQueue myQueue = new MyQueue();
// myQueue.push(1); // queue is: [1]
// myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
// myQueue.peek(); // return 1
// myQueue.pop(); // return 1, queue is [2]
// myQueue.empty(); // return false
//
//
// 提示：
//
// 1 <= x <= 9
// 最多调用 100 次 push、pop、peek 和 empty
// 假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）
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

}

// 两个栈
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

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {

	// 当出栈全都空了再 轮训入栈
	if len(this.Q2) == 0 {
		for len(this.Q1) > 0 {
			q := this.Q1[len(this.Q1)-1]
			this.Q2 = append(this.Q2, q)
			this.Q1 = this.Q1[:len(this.Q1)-1]
		}
	}

	if len(this.Q2) == 0 {
		return 0
	}

	a := this.Q2[len(this.Q2)-1]
	this.Q2 = this.Q2[:len(this.Q2)-1]
	return a
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.Q2) == 0 {
		return 0
	}
	return this.Q2[len(this.Q2)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.Q2) == 0 {
		return true

	}
	return false
}
