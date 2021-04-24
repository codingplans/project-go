package _32

import "testing"

// 设计你的循环队列实现。 循环队列是一种线性数据结构，其操作表现基于 FIFO（先进先出）原则并且队尾被连接在队首之后以形成一个循环。它也被称为“环形缓冲器”。
//
// 循环队列的一个好处是我们可以利用这个队列之前用过的空间。在一个普通队列里，一旦一个队列满了，我们就不能插入下一个元素，即使在队列前面仍有空间。但是使用循环队列，我们能使用这些空间去存储新的值。
//
// 你的实现应该支持如下操作：
//
// MyCircularQueue(k): 构造器，设置队列长度为 k 。
// Front: 从队首获取元素。如果队列为空，返回 -1 。
// Rear: 获取队尾元素。如果队列为空，返回 -1 。
// enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
// deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
// isEmpty(): 检查循环队列是否为空。
// isFull(): 检查循环队列是否已满。
//
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/design-circular-queue
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func Test_622(t *testing.T) {
	k := 5
	obj := Constructor(k)
	param_1 := obj.EnQueue(2)
	obj.EnQueue(2)
	obj.EnQueue(7)
	obj.EnQueue(1)
	obj.EnQueue(3)
	obj.EnQueue(4)
	param_2 := obj.DeQueue()
	obj.DeQueue()
	param_3 := obj.Front()
	param_4 := obj.Rear()
	param_5 := obj.IsEmpty()
	param_6 := obj.IsFull()
	println(param_1, param_2, param_3, param_4, param_5, param_6)
}

type MyCircularQueue struct {
	Q1  []int
	len int
}

func Constructor(k int) MyCircularQueue {
	q := new(MyCircularQueue)
	q.len = k
	return *q

}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.len == len(this.Q1) {
		return false
	}
	this.Q1 = append(this.Q1, value)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if len(this.Q1) == 0 {
		return false
	}
	this.Q1 = this.Q1[1:]
	return true
}

func (this *MyCircularQueue) Front() int {
	if len(this.Q1) == 0 {
		return -1
	}
	return this.Q1[0]

}

func (this *MyCircularQueue) Rear() int {
	if len(this.Q1) == 0 {
		return -1
	}
	return this.Q1[len(this.Q1)-1]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return len(this.Q1) == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return len(this.Q1) == this.len
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
