package main

import "testing"

func TestIsValid(t *testing.T) {
	// t.Log(isValid("[](){}"))
	// t.Log(isValid("[]({}[]{{{}}}){}"))
	// t.Log(isValid("{[}]"))
	// t.Log(isValid("]"))
	// t.Log(Fibonacci(4))
	// t.Log(FibonacciV2(4))
	// t.Log(Fibonacci(10))
	// t.Log(FibonacciV2(30))
	// t.Log(search([]int{1, 2, 2, 3, 3, 6, 8, 8, 8, 9, 9, 9}, 6))
	// t.Log(search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6))
	// t.Log(search([]int{1, 1, 1, 1, 6, 6, 6, 6, 6, 7, 8, 9, 10, 10, 101}, 101))
	// t.Log(search([]int{-2, 1, 2}, -2))
	t.Log(LRU([][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}, 3))
}

type Lru struct {
	md   map[int]*node
	max  int
	head *node
	tail *node
}

type node struct {
	pre, next *node
	val, key  int
}

func LRU(operators [][]int, k int) []int {
	// write code here
	lru := initLru(k)
	res := []int{}
	for i := range operators {
		if operators[i][0] == 1 {
			lru.set(operators[i][1], operators[i][2])

		} else {
			res = append(res, lru.get(operators[i][1]))
		}
	}
	return res
}
func initLru(k int) Lru {
	return Lru{
		md:  make(map[int]*node),
		max: k,
	}
}

func (this *Lru) get(k int) int {
	if v, ok := this.md[k]; ok {
		this.remove(v)
		this.add(v)
		return v.val
	}
	return -1
}
func (this *Lru) set(k, x int) {
	if v, ok := this.md[k]; ok {
		this.remove(v)
		this.add(v)
		return
	} else {
		n := &node{val: x, key: k}
		this.md[k] = n
		this.add(n)

	}

	if len(this.md) > this.max {
		delete(this.md, this.tail.key)
		this.remove(this.tail)
	}

}

func (this *Lru) remove(n *node) {
	if this.head == n {
		this.head = n.next
		this.head.pre = nil
		return
	}

	if this.tail == n {
		this.tail = n.pre
		n.pre.next = nil
		n.pre = nil
		return
	}

	n.pre.next = n.next
	n.next.pre = n.pre

	return

}

func (this *Lru) add(n *node) {
	n.next = this.head
	if this.head != nil {
		this.head.pre = n
	}
	this.head = n
	if this.tail == nil {
		this.tail = n
		this.tail.next = nil
	}

	return
}
