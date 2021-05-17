package _00_init_code

import (
	"fmt"
	"testing"
)

type test struct {
	IntEs []int
	Lists []int
	K     int
}

var tests = []test{}

// 146. LRU 缓存机制
// 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。
// 实现 LRUCache 类：
//
// LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//
//
// 进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？
//
//
//
// 示例：
//
// 输入
// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
// 输出
// [null, null, null, 1, null, -1, null, -1, 3, 4]
//
// 解释
// LRUCache lRUCache = new LRUCache(2);
// lRUCache.put(1, 1); // 缓存是 {1=1}
// lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
// lRUCache.get(1);    // 返回 1
// lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
// lRUCache.get(2);    // 返回 -1 (未找到)
// lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
// lRUCache.get(1);    // 返回 -1 (未找到)
// lRUCache.get(3);    // 返回 3
// lRUCache.get(4);    // 返回 4
//
//
// 提示：
//
// 1 <= capacity <= 3000
// 0 <= key <= 3000
// 0 <= value <= 104
// 最多调用 3 * 104 次 get 和 put

func Test_upToDayUp(t *testing.T) {
	obj := Constructor(3)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Put(3, 3)
	MList2Ints(&obj)

	aa := obj.Get(4)
	fmt.Println(aa)
	obj.Put(4, 4)
	aa = obj.Get(2)
	fmt.Println(aa)
	MList2Ints(&obj)

	// fmt.Printf("obj = %v\n", MList2Ints(&obj))
	// obj.Put(1, 1)
	// fmt.Printf("obj = %v\n", MList2Ints(&obj))
	// obj.Put(2, 2)
	// fmt.Printf("obj = %v\n", MList2Ints(&obj))
	// param1 := obj.Get(1)
	// fmt.Printf("param_1 = %v obj = %v\n", param1, MList2Ints(&obj))
	// obj.Put(3, 3)
	// fmt.Printf("obj = %v\n", MList2Ints(&obj))
	// param1 = obj.Get(2)
	// fmt.Printf("param_1 = %v obj = %v\n", param1, MList2Ints(&obj))
	// obj.Put(4, 4)
	// fmt.Printf("obj = %v\n", MList2Ints(&obj))
	// param1 = obj.Get(1)
	// fmt.Printf("param_1 = %v obj = %v\n", param1, MList2Ints(&obj))
	// param1 = obj.Get(3)
	// fmt.Printf("param_1 = %v obj = %v\n", param1, MList2Ints(&obj))
	// param1 = obj.Get(4)
	// fmt.Printf("param_1 = %v obj = %v\n", param1, MList2Ints(&obj))
}

type LRUCache struct {
	buf        int
	md         map[int]*ListNode
	head, tail *ListNode
}

type ListNode struct {
	prev, next *ListNode
	val, key   int
}

func MList2Ints(node *LRUCache) int {
	node.Travel()
	return 0
}

func Constructor(capacity int) LRUCache {
	md1 := make(map[int]*ListNode, capacity)
	return LRUCache{buf: capacity, md: md1}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.md[key]; ok {
		this.Remove(v)
		this.Add(v)
		return v.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.md[key]; ok {
		v.val = value
		this.Remove(v)
		this.Add(v)
		return
	} else {
		n := &ListNode{val: value, key: key}
		this.md[key] = n
		this.Add(n)

	}
	if len(this.md) > this.buf {
		delete(this.md, this.tail.key)
		this.Remove(this.tail)
	}

}

func (this *LRUCache) Travel() {
	if this.head == nil {
		return
	}
	a := this.head
	if a != nil {
		fmt.Println(a.val, 11)
		a = a.next
		// fmt.Println(a.val, 11)
	}

	for v := range this.md {
		fmt.Println(this.md[v].val)
	}

}

func (this *LRUCache) Add(node *ListNode) {
	node.prev = nil
	node.next = this.head
	if this.head != nil {
		this.head.prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
		this.tail.next = nil
	}
}

func (this *LRUCache) Remove(node *ListNode) {
	if this.head == node {
		this.head = node.next
		node.next = nil
		return
	}
	if this.tail == node {
		this.tail = node.prev
		node.prev.next = nil
		node.prev = nil
		return
	}

	node.prev.next = node.next
	node.next.prev = node.prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
