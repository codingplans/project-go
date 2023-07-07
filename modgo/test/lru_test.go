package main

import "testing"

func TestLru(t *testing.T) {
	l := initlru(3)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, i2 := range arr {
		l.set(i2, i2)
	}
	l.rangeLru()
}

type Lru struct {
	m          map[int]*node
	max        int
	head, tail *node
}

type node struct {
	pre, next *node
	val, key  int
}

func initlru(k int) *Lru {
	return &Lru{
		m:    make(map[int]*node, k),
		max:  k,
		head: nil,
		tail: nil,
	}
}

func (l *Lru) remove(n *node) {
	if l.head == n {
		// l.head = l.head.next
		l.head = n.next
		l.head.pre = nil
		return
	}
	if l.tail == n {
		l.tail = n.pre
		n.pre.next = nil
		n.pre = nil
		return
	}
	n.pre.next = n.next
	n.next.pre = n.pre
}

func (l *Lru) get(k int) int {
	if v, ok := l.m[k]; ok {
		l.remove(v)
		l.add(v)
		return v.val
	}
	return -1
}

func (l *Lru) set(k, v int) {

	if s, ok := l.m[k]; ok {
		l.remove(s)
		l.add(s)
		return
	}
	if len(l.m) >= l.max {
		delete(l.m, l.tail.key)
		l.remove(l.tail)
	}
	n := &node{key: k, val: v}
	l.m[k] = n
	l.add(n)

}
func (l *Lru) add(n *node) {
	n.next = l.head
	if l.head != nil {
		l.head.pre = nil
	}
	l.head = n
	if l.tail == nil {
		l.tail = l.head
		l.tail.next = nil
	}

}
func (l *Lru) rangeLru() {
	n := l.head
	for n != nil {
		println(n.val, n.key)
		n = n.next
	}
}
