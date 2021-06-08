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
	obj := Constructor(3)
	obj.Put(2, 2)
	obj.travel()
	obj.Put(1, 1)
	obj.travel()

	obj.Get(2)
	obj.travel()

	obj.Get(1)
	obj.travel()
	obj.Get(2)
	obj.travel()
	obj.Put(3, 3)
	obj.travel()

	obj.Put(3, 4)
	obj.travel()

	obj.Put(3, 5)
	obj.travel()

	obj.Put(5, 5)
	obj.travel()

	obj.Put(4, 4)
	obj.travel()
	obj.Get(3)
	obj.travel()
	obj.Get(2)
	obj.travel()
	obj.Get(1)
	obj.travel()
	obj.Get(4)
	obj.travel()

}

type LFUCache struct {
	buf, min int
	md       map[int]*node // 存储 kv 结构， 存储 k 频率结构
	fks      map[int][]int // 记录每个频率的 key
}

type node struct {
	key int
	val int
	fa  int
}

func Constructor(capacity int) LFUCache {
	md1 := make(map[int]*node, capacity)
	fks := make(map[int][]int, capacity) // 存储同一个频次的 k 的顺序
	return LFUCache{buf: capacity, md: md1, fks: fks}
}

func (this *LFUCache) Get(key int) int {
	if v, ok := this.md[key]; ok {
		this.increaseMd(key)
		this.increaseFk(key)
		return v.val
	}
	return -1

}

func (this *LFUCache) Put(key int, value int) {
	if this.buf == 0 {
		return
	}
	n := new(node)
	if v, ok := this.md[key]; !ok {
		if len(this.md) == this.buf {
			delKey := this.delFks()
			println(delKey, 22222, this.min)
			if delKey >= 0 {
				delete(this.md, delKey)
			}
		}
		this.min = 1
		n = &node{
			val: value,
			key: key,
			fa:  1,
		}
	} else {
		v.fa++
		v.val = value
		n = v
	}
	this.md[key] = n
	this.increaseFk(key)
}

// 增加 kf
func (this *LFUCache) increaseFk(key int) {

	v, _ := this.md[key]

	this.delFksByfa(v.fa-1, key)
	if v.fa-1 == this.min {
		// 重置最小值
		this.findMin()
	}
	this.putFks(v.fa, key)
	return
}

func (this *LFUCache) increaseMd(key int) {
	v, ok := this.md[key]
	if ok {
		v.fa++
		this.md[key] = v
		// println(1111111111, v.fa-1, this.min)

		if v.fa-1 == this.min {
			// 重置最小值
			this.findMin()
		}
	}
}

func (this *LFUCache) findMin() {
	if v, ok := this.fks[this.min]; ok {
		if len(v) > 1 {
			return
		}
	}

	i := 0
	for _, v := range this.md {
		if i == 0 {
			this.min = v.fa
			i++
		}
		if this.min > v.fa {
			this.min = v.fa
		}
	}
}

func (this *LFUCache) putFks(fa, key int) {
	this.fks[fa] = append(this.fks[fa], key)
}

// 删除最小的最后一个 key
func (this *LFUCache) delFks() int {
	if len(this.fks[this.min]) == 0 {
		return -1
	}
	key := this.fks[this.min][0]
	this.fks[this.min] = this.fks[this.min][1:]
	return key
}

// 根据频率删除 key
func (this *LFUCache) delFksByfa(fa, key int) int {
	v := this.fks[fa]
	index := 0
	if len(v) == 1 {
		this.fks[fa] = []int{}
	} else {
		for k := range v {
			if v[k] == key {
				index = k
				this.fks[fa] = append(this.fks[fa][:index], this.fks[fa][index+1:]...)
				break
			}
		}
	}
	return key
}

func (this *LFUCache) travel() {
	for k := range this.md {
		fmt.Println("md：", this.md[k])
	}
	fmt.Println(this.fks, this.min)

}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
