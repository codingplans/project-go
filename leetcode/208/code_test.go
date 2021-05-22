package _00_init_code

import (
	"fmt"
	"testing"
)

type test struct {
	IntEs []int
	Lists []int
	K     int
	Str   string
}

var tests = []test{}

// 208. 实现 Trie (前缀树)
// Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。
//
// 请你实现 Trie 类：
//
// Trie() 初始化前缀树对象。
// void insert(String word) 向前缀树中插入字符串 word 。
// boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
// boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。
//
//
// 示例：
//
// 输入
// ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
// [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
// 输出
// [null, null, true, false, true, null, true]
//
// 解释
// Trie trie = new Trie();
// trie.insert("apple");
// trie.search("apple");   // 返回 True
// trie.search("app");     // 返回 False
// trie.startsWith("app"); // 返回 True
// trie.insert("app");
// trie.search("app");     // 返回 True
//
//
// 提示：
//
// 1 <= word.length, prefix.length <= 2000
// word 和 prefix 仅由小写英文字母组成
// insert、search 和 startsWith 调用次数 总计 不超过 3 * 104 次
// 通过

func Test_upToDayUp(t *testing.T) {

	tt := Constructor()

	tt.Insert("apple")
	fmt.Println(tt.Search("apple"))
	fmt.Println(tt.Search("app"))
	fmt.Println(tt.StartsWith("app"))
	tt.Insert("app")
	fmt.Println(tt.Search("app"))

}

type Trie struct {
	arr   [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	arrs := this
	for k := range word {
		v := word[k] - 'a'
		if arrs.arr[v] == nil {
			arrs.arr[v] = &Trie{}
		}
		arrs = arrs.arr[v]
	}
	arrs.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	res := this
	for k := range word {
		v := word[k] - 'a'
		if res.arr[v] == nil {
			return false
		}
		res = res.arr[v]
	}
	return res != nil && res.isEnd == true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	res := this
	for k := range prefix {
		v := prefix[k] - 'a'
		if res.arr[v] == nil {
			return false
		}
		res = res.arr[v]
	}
	return res != nil

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
