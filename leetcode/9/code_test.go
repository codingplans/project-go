package _32

import (
	"fmt"
	"testing"
)

// 09. 用两个栈实现队列
// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
//
// 示例 1：
//
// 输入：
// ["CQueue","appendTail","deleteHead","deleteHead"]
// [[],[3],[],[]]
// 输出：[null,null,3,-1]
// 示例 2：
//
// 输入：
// ["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
// [[],[],[5],[2],[],[]]
// 输出：[null,-1,null,null,5,2]
// 提示：
//
// 1 <= values <= 10000
// 最多会对 appendTail、deleteHead 进行 10000 次调用

type test struct {
	V []int
}

var tests = []test{
	{[]int{3, 0, 0}},
	{[]int{0, 5, 2, 0, 0}},
	{[]int{1, 2, 0, 3, 4}},
	{[]int{0, 1, 0, 4, 0, 2, 1, 0, 4}},
	{[]int{0, 1}},
	{[]int{0}},
	{[]int{10, 0, 9}},
}

func Test_upToDayUp(t *testing.T) {
	for k := range tests {
		// println("初始化")
		Q := Constructor()
		fmt.Println(nil)
		for _, v := range tests[k].V {
			if v == 0 {
				a := Q.DeleteHead()
				fmt.Println(a)
				continue
			}
			Q.AppendTail(v)
			fmt.Println(nil)
		}

	}

}

type CQueue struct {
	st1, st2 []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.st1 = append(this.st1, value)
	for l := len(this.st1) - 1; l >= 0; l-- {
		this.st2 = append(this.st2, this.st1[l])
	}
	this.st1 = []int{}
}

func (this *CQueue) DeleteHead() int {
	// fmt.Println(this.st2)
	if len(this.st2) > 0 {
		defer func() { this.st2 = this.st2[1:] }()
		return this.st2[0]
	}
	return -1

}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
