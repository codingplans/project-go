package _00_init_code

import (
	"testing"
)

type test struct {
	IntEs []int
	Lists []int
	K     int
}

var tests = []test{}

//
// LFU的一个实现方法：
// 用一个主双向链表记录（访问次数，从链表头），从链表中按时间顺序记录着（key）
// 用一个哈希表记录（key，(value, 主链表ptr，从链表ptr)）ptr表示该key在链表中的地址
// 然后，get，put都在哈希表中操作，近似O(1)，哈希表中有个节点在链表中的地址，能O(1)找到，并把节点提搞访问频次，链表插入删除也都是O(1)。

// 运用你所掌握的数据结构，设计和实现一个  LFU (最近最少使用) 缓存机制 。
// 实现 LFUCache 类：
//
// LFUCache(int capacity) 以正整数作为容量 capacity 初始化 LFU 缓存
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
// ["LFUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
// 输出
// [null, null, null, 1, null, -1, null, -1, 3, 4]
//
// 解释
// LFUCache lRUCache = new LFUCache(2);
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
	// obj := Constructor(2)
	// obj.Put(1, 1)
	// obj.travel()
	//
	// obj.Put(2, 2)
	// obj.travel()
	//
	// aa := obj.Get(1)
	// obj.travel()
	//
	// fmt.Println(aa)
	// obj.Put(3, 3)
	// obj.travel()
	//
	// aa = obj.Get(2)
	// obj.travel()
	//
	// fmt.Println(aa)
	// aa = obj.Get(3)
	// obj.travel()
	//
	// fmt.Println(aa)
	//
	// obj.Put(4, 4)
	// obj.travel()
	//
	// aa = obj.Get(1)
	// obj.travel()
	//
	// fmt.Println(aa)
	// aa = obj.Get(3)
	// obj.travel()
	//
	// fmt.Println(aa)
	// aa = obj.Get(4)
	// obj.travel()
	//
	// fmt.Println(aa)

	// ["LFUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
	// 	[[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
	// [null,null,null,null,null,null,-1,null,19,17,null,-1,null,null,null,-1,null,-1,5,-1,12,null,null,3,5,5,null,null,1,null,-1,null,30,5,30,null,null,null,-1,null,-1,24,null,null,18,null,null,null,null,14,null,null,18,null,null,11,null,null,null,null,null,18,null,null,-1,null,4,29,30,null,12,11,null,null,null,null,29,null,null,null,null,17,-1,18,null,null,null,-1,null,null,null,20,null,null,null,29,18,18,null,null,null,null,20,null,null,null,null,null,null,null]
	// [null,null,null,null,null,null,-1,null,19,17,null,-1,null,null,null,-1,null,-1,5,-1,12,null,null,3,5,5,null,null,1,null,-1,null,30,5,30,null,null,null,-1,null,-1,24,null,null,18,null,null,null,null,14,null,null,18,null,null,11,null,null,null,null,null,18,null,null,24,null,4,29,30,null,12,11,null,null,null,null,29,null,null,null,null,17,22,18,null,null,null,24,null,null,null,20,null,null,null,29,18,18,null,null,null,null,20,null,null,null,null,null,null,null]
	// ["LFUCache","put","put","get","put","get","get","put","get","get","get"]
	// 	[[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]
	// obj := Constructor(3)
	// obj.Put(2, 2)
	// obj.travel()
	// obj.Put(1, 1)
	// obj.travel()
	//
	// obj.Get(2)
	// obj.travel()
	//
	// obj.Get(1)
	// obj.travel()
	// obj.Get(2)
	// obj.travel()
	// obj.Put(3, 3)
	// obj.travel()
	//
	// obj.Put(3, 4)
	// obj.travel()
	//
	// obj.Put(3, 5)
	// obj.travel()
	//
	// obj.Put(5, 5)
	// obj.travel()
	//
	// obj.Put(4, 4)
	// obj.travel()
	// obj.Get(3)
	// obj.travel()
	// obj.Get(2)
	// obj.travel()
	// obj.Get(1)
	// obj.travel()
	// obj.Get(4)
	// obj.travel()

}

type node struct {
	key   int
	value int
	freq  int
	pre   *node // Node所在频次的双向链表的前继Node
	tail  *node //
	head  *node //

	post *node // Node所在频次的双向链表的后继Node

	// doublyLinkedList DoublyLinkedList // Node所在频次的双向链表

}

type LFUCache struct {
	cache map[int]*node // 存储缓存的内容，Node中除了value值外，还有key、freq、所在doublyLinkedList、所在doublyLinkedList中的postNode、所在doublyLinkedList中的preNode，具体定义在下方。

	firstLinkedList *node // firstLinkedList.post 是频次最大的双向链表
	lastLinkedList  *node // lastLinkedList.pre 是频次最小的双向链表，满了之后删除 lastLinkedList.pre.tail.pre 这个Node即为频次最小且访问最早的Node

	size     int
	capacity int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cache:           make(map[int]*node, capacity),
		capacity:        capacity,
		firstLinkedList: &node{},
		lastLinkedList:  &node{},
	}
	// re
	// ​    cache = new
	// HashMap < > (capacity)
	// ​    firstLinkedList = new
	// DoublyLinkedList()
	// ​    lastLinkedList = new
	// DoublyLinkedList()
	// ​    firstLinkedList.post = lastLinkedList
	// ​    lastLinkedList.pre = firstLinkedList
	// ​    this.capacity = capacity
}

func (this *LFUCache) get(key int) int {

	if v, ok := this.cache[key]; ok {
		// 该key访问频次+1

		this.freqInc(v)
		return v.value
	}

	return -1

}

func (this *LFUCache) freqInc(nd *node) {

}

func (this *LFUCache) put(key, val int) {
	if this.capacity == 0 {
		return
	}
	nd, ok := this.cache[key]
	if ok {
		// 若key存在，则更新value，访问频次+1
		nd.value = val
		this.freqInc(nd)

	} else {
		// 若key不存在
		if this.size == this.capacity {
			// 如果缓存满了，删除lastLinkedList.pre这个链表（即表示最小频次的链表）中的tail.pre这个Node（即最小频次链表中最先访问的Node），如果该链表中的元素删空了，则删掉该链表。
			this.remove(this.lastLinkedList.pre.tail.pre.key)

			this.removeNode(this.lastLinkedList.pre.tail.pre)
			this.size--
			if this.lastLinkedList.pre.head.post == this.lastLinkedList.tail.pre {
				this.removeDoublyLinkedList(this.lastLinkedList.pre)
			}

		}
		// cache中put新Key-Node对儿，并将新node加入表示freq为1的DoublyLinkedList中，若不存在freq为1的DoublyLinkedList则新建。

		this.cache[key] = &node{key: key, value: val}
		if this.lastLinkedList.pre.freq != 1 {
			newDoublyLinedList := &node{freq: 1}
			this.addDoublyLinkedList(newDoublyLinedList, this.lastLinkedList.pre)

		} else {

		}
	}

}

//
//
// ​      Node newNode = new Node(key, value)
// ​      cache.put(key, newNode)
// ​      if (lastLinkedList.pre.freq != 1) {
// ​        DoublyLinkedList newDoublyLinedList = new DoublyLinkedList(1)
// ​        addDoublyLinkedList(newDoublyLinedList, lastLinkedList.pre)
// ​        newDoublyLinedList.addNode(newNode)
// ​
// } else {
// ​        lastLinkedList.pre.addNode(newNode)
// ​
// }
//
// ​      size++
// ​
// }
// }

func (this *LFUCache) remove(key int)                      {}
func (this *LFUCache) removeNode(nd *node)                 {}
func (this *LFUCache) removeDoublyLinkedList(nd *node)     {}
func (this *LFUCache) addDoublyLinkedList(nd, addNd *node) {}

//
// /**
//  * node的访问频次 + 1
//  */
// void freqInc(Node node) {
// ​ // 将node从原freq对应的双向链表里移除, 如果链表空了则删除链表。
//
// ​    DoublyLinkedList linkedList = node.doublyLinkedList
// ​    DoublyLinkedList preLinkedList = linkedList.pre
// ​    linkedList.removeNode(node)
// ​    if (linkedList.head.post == linkedList.tail) {
// ​      removeDoublyLinkedList(linkedList)
// ​
// }
//
//
// ​ // 将node加入新freq对应的双向链表，若该链表不存在，则先创建该链表。
//
// ​    node.freq++
// ​    if (preLinkedList.freq != node.freq) {
// ​      DoublyLinkedList newDoublyLinedList = new DoublyLinkedList(node.freq)
// ​      addDoublyLinkedList(newDoublyLinedList, preLinkedList)
// ​      newDoublyLinedList.addNode(node)
// ​
// } else {
// ​      preLinkedList.addNode(node)
// ​
// }
// }
//
// /**
//  * 增加代表某1频次的双向链表
//  */
// void addDoublyLinkedList(DoublyLinkedList newDoublyLinedList, DoublyLinkedList preLinkedList) {
// ​    newDoublyLinedList.post = preLinkedList.post
// ​    newDoublyLinedList.post.pre = newDoublyLinedList
// ​    newDoublyLinedList.pre = preLinkedList
// ​    preLinkedList.post = newDoublyLinedList
// }
//
// /**
//  * 删除代表某1频次的双向链表
//  */
// void removeDoublyLinkedList(DoublyLinkedList doublyLinkedList) {
// ​    doublyLinkedList.pre.post = doublyLinkedList.post
// ​    doublyLinkedList.post.pre = doublyLinkedList.pre
// }
// }
// class Node {
//
// int key;
//
// int value;
//
// int freq = 1;
//
// Node pre; // Node所在频次的双向链表的前继Node
//
// Node post; // Node所在频次的双向链表的后继Node
//
// DoublyLinkedList doublyLinkedList;  // Node所在频次的双向链表
//
// public Node() {}
//
//
//
// public Node(int key, int value) {
//
// ​    this.key = key;
//
// ​    this.value = value;
//
// }
//
// }
//
//
//
// class DoublyLinkedList {
//
// int freq; // 该双向链表表示的频次
//
// DoublyLinkedList pre; // 该双向链表的前继链表（pre.freq < this.freq）
//
// DoublyLinkedList post; // 该双向链表的后继链表 (post.freq > this.freq)
//
// Node head; // 该双向链表的头节点，新节点从头部加入，表示最近访问
//
// Node tail; // 该双向链表的尾节点，删除节点从尾部删除，表示最久访问
//
// public DoublyLinkedList() {
//
// ​    head = new Node();
//
// ​    tail = new Node();
//
// ​    head.post = tail;
//
// ​    tail.pre = head;
//
// }
//
//
//
// public DoublyLinkedList(int freq) {
//
// ​    head = new Node();
//
// ​    tail = new Node();
//
// ​    head.post = tail;
//
// ​    tail.pre = head;
//
// ​    this.freq = freq;
//
// }
//
//
//
// void removeNode(Node node) {
//
// ​    node.pre.post = node.post;
//
// ​    node.post.pre = node.pre;
//
// }
//
//
//
// void addNode(Node node) {
//
// ​    node.post = head.post;
//
// ​    head.post.pre = node;
//
// ​    head.post = node;
//
// ​    node.pre = head;
//
// ​    node.doublyLinkedList = this;
//
// }
//
//
//
// }
