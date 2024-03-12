package main

import (
	"fmt"
	"testing"
)

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

type Node struct {
	key   int
	value interface{}
	prev  *Node
	next  *Node
}

type LRU struct {
	capacity int
	size     int
	head     *Node
	tail     *Node
	cache    map[int]*Node
}

func TestLRUN(t *testing.T) {
	lru := NewLRU(3)

	lru.Put(1, "a")
	lru.Put(2, "b")
	lru.Put(3, "c")

	fmt.Println(lru.Get(1)) // "a"
	fmt.Println(lru.Get(2)) // "b"

	lru.Put(4, "d")

	fmt.Println(lru.Get(3)) // nil

	lru.Put(5, "e")

	fmt.Println(lru.Get(1)) // nil

}

func NewLRU(capacity int) *LRU {
	lru := &LRU{
		capacity: capacity,
		size:     0,
		head:     nil,
		tail:     nil,
		cache:    make(map[int]*Node),
	}
	return lru
}

func (lru *LRU) Get(key int) interface{} {
	node, ok := lru.cache[key]
	if !ok {
		return nil
	}

	// 将节点移动到链表头部
	lru.moveToHead(node)

	return node.value
}

func (lru *LRU) Put(key int, value interface{}) {
	node, ok := lru.cache[key]
	if ok {
		// 更新节点值
		node.value = value

		// 将节点移动到链表头部
		lru.moveToHead(node)
	} else {
		// 创建新节点
		node = &Node{
			key:   key,
			value: value,
		}

		// 将新节点添加到链表头部
		lru.addNodeToHead(node)

		// 缓存容量已满，淘汰最久未使用的节点
		if lru.size == lru.capacity {
			lru.removeTail()
		}
		lru.cache[key] = node
	}
}

func (lru *LRU) moveToHead(node *Node) {
	if node == lru.head {
		return
	}

	// 将节点从原位置移除
	lru.removeNode(node)

	// 将节点添加到链表头部
	lru.addNodeToHead(node)
}

func (lru *LRU) addNodeToHead(node *Node) {
	if lru.head == nil {
		lru.head = node
		lru.tail = node
	} else {
		node.next = lru.head
		lru.head.prev = node
		lru.head = node
	}

	lru.size++
}

func (lru *LRU) removeNode(node *Node) {
	if node == lru.head {
		lru.head = node.next
		if lru.head != nil {
			lru.head.prev = nil
		} else {
			lru.tail = nil
		}
	} else if node == lru.tail {
		lru.tail = node.prev
		lru.tail.next = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}

	lru.size--
}

func (lru *LRU) removeTail() {
	node := lru.tail

	lru.removeNode(node)

	delete(lru.cache, node.key)
}
