package _00_init_code

import (
	"fmt"
	"github.com/Darrenzzy/testgo/xzap"
	"testing"
)

//
// LFU的一个实现方法：
// 用一个主双向链表记录（访问次数，从链表头），从链表中按时间顺序记录着（key）
// 用一个哈希表记录（key，(value, 主链表ptr，从链表ptr)）ptr表示该key在链表中的地址
// 然后，get，put都在哈希表中操作，近似O(1)，哈希表中有个节点在链表中的地址，能O(1)找到，并把节点提搞访问频次，链表插入删除也都是O(1)。

// 运用你所掌握的数据结构，设计和实现一个  LFU (最近最少使用) 缓存机制 。
// 实现 LFUCacheV2 类：
//
// LFUCacheV2(int capacity) 以正整数作为容量 capacity 初始化 LFU 缓存
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
// ["LFUCacheV2", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
// 输出
// [null, null, null, 1, null, -1, null, -1, 3, 4]
//
// 解释
// LFUCacheV2 lRUCache = new LFUCacheV2(2);
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

func Test_upToDayUpV2(t *testing.T) {
	// obj := Constructor(1)
	// obj.Put(1, 1)
	// obj.Put(2, 2)
	// obj.Put(3, 3)
	// obj.Put(2, 2)
	// obj.Put(4, 4)

	// obj := Constructor(2)
	// obj.Put(1, 1)
	// fmt.Println(obj.fks, obj.kf, 999)
	//
	// obj.Put(2, 2)
	// fmt.Println(obj.fks, obj.kf, 999)
	//
	// aa := obj.Get(1)
	// fmt.Println(obj.fks, obj.kf, 999)
	// fmt.Println(aa)
	// obj.Put(3, 3)
	// fmt.Println(obj.fks, obj.kf, 999)
	//
	// aa = obj.Get(2)
	// fmt.Println(obj.fks, obj.kf, 999)
	//
	// fmt.Println(aa)
	// aa = obj.Get(3)
	// fmt.Println(obj.fks, obj.kf, 999)
	//
	// fmt.Println(aa)
	//
	// obj.Put(4, 4)
	// fmt.Println(obj.fks, obj.kf, 999)
	//
	// aa = obj.Get(1)
	// fmt.Println(obj.fks, obj.kf, 999)
	//
	// fmt.Println(aa)
	// aa = obj.Get(3)
	// fmt.Println(obj.fks, obj.kf, 999)
	//
	// fmt.Println(aa)
	// aa = obj.Get(4)
	// fmt.Println(obj.fks, obj.kf, 999)
	// fmt.Println(aa)

	obj := ConstructorV2(3)
	obj.Put(2, 2)
	xzap.Info()
	fmt.Printf("%+v,%+v  \n", obj.fks, obj.md)
	obj.Put(1, 1)
	fmt.Printf("%+v,%+v  \n", obj.fks, obj.md)
	obj.Get(2)
	fmt.Printf("%+v,%+v  \n", obj.fks, obj.md)
	obj.Get(1)
	fmt.Printf("%+v,%+v  \n", obj.fks, obj.md)
	obj.Get(2)
	fmt.Printf("%+v,%+v  \n", obj.fks, obj.md)
	obj.Put(3, 3)
	fmt.Printf("%+v,%+v  \n", obj.fks, obj.md)
	obj.Put(4, 4)
	fmt.Printf("%+v,%+v  \n", obj.fks, obj.md)
	obj.Get(3)
	fmt.Printf("%+v,%+v  \n", obj.fks, obj.md)
	obj.Get(2)
	fmt.Printf("%+v,%+v  \n", obj.fks, obj.md)
	obj.Get(1)
	fmt.Printf("%+v,%+v  \n", obj.fks, obj.md)
	obj.Get(4)
	fmt.Printf("%+v,%+v  \n", obj.fks, obj.md)
}

type LFUCacheV2 struct {
	buf, min int
	md       map[int]*node // 存储 kv 结构， 存储 k 频率结构
	fks      map[int][]int // 记录每个频率的 key
}

type node struct {
	key int
	val int
	fa  int
}

func ConstructorV2(capacity int) LFUCacheV2 {
	md1 := make(map[int]*node, capacity)
	fks := make(map[int][]int, capacity) // 存储同一个频次的 k 的顺序
	return LFUCacheV2{buf: capacity, md: md1, fks: fks}
}

func (this *LFUCacheV2) Get(key int) int {

	if v, ok := this.md[key]; ok {
		this.increaseKf(key)
		return v.val
	}
	return -1

}

func (this *LFUCacheV2) Put(key int, value int) {
	if this.buf == 0 {
		return
	}
	n := new(node)
	if v, ok := this.md[key]; !ok {
		if len(this.md) == this.buf {
			delKey := this.delFks()
			// println(delKey, 444, key)
			if delKey > 0 {
				delete(this.md, delKey)
			}
			// fmt.Println(3333, delKey)
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
	// this.increaseKf(key)
}

// 增加 kf
func (this *LFUCacheV2) increaseKf(key int) {
	v, ok := this.md[key]
	// if len(this.md) == 0 {
	// 	this.min = 1
	// }
	if ok {
		fa := v.fa
		if fa == this.min {
			// 重置最小值
			this.min++
			// this.findMin()
		}
		v.fa++
		this.md[key] = v
		this.delFksByfa(fa, key)
		this.putFks(v.fa, key)
	}

	return
}

func (this *LFUCacheV2) findMin() {
	for _, v := range this.md {
		if this.min > v.fa {
			this.min = v.fa
		}
	}
}

func (this *LFUCacheV2) putFks(fa, key int) {
	this.fks[fa] = append(this.fks[fa], key)
	// fmt.Println(this.fks[fa], fa, key)
}

// 删除最小的最后一个 key
func (this *LFUCacheV2) delFks() int {
	if len(this.fks[this.min]) == 0 {
		return -1
	}
	key := this.fks[this.min][0]
	this.fks[this.min] = this.fks[this.min][1:]
	return key
}

// 根据频率删除 key
func (this *LFUCacheV2) delFksByfa(fa, key int) int {
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

/**
 * Your LFUCacheV2 object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
