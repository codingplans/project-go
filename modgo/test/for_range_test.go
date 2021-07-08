package main

import "testing"

func BenchmarkForPointer(b *testing.B) {
	items := generateItems(1024)
	for i := 0; i < b.N; i++ {
		length := len(items)
		var tmp int
		for k := 0; k < length; k++ {
			tmp = items[k].id
		}
		_ = tmp
	}
}

func BenchmarkRangePointer(b *testing.B) {
	items := generateItems(1024)
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}
func BenchmarkRangeValPointer(b *testing.B) {
	items := generateItems(1024)
	for i := 0; i < b.N; i++ {
		var tmp int
		for k, item := range items {
			_ = k
			tmp = item.id
		}
		_ = tmp
	}
}

func BenchmarkRangeByKeyPointer(b *testing.B) {
	items := generateItems(1024)
	for i := 0; i < b.N; i++ {
		var tmp int
		for k := range items {
			tmp = items[k].id
		}
		_ = tmp
	}
}

// 对比 for 循环效率问题
// go test -bench=. -run=none
// 看到函数后面的-8了吗？这个表示运行时对应的GOMAXPROCS的值。
//
// 接着的20000000表示运行for循环的次数也就是调用被测试代码的次数
//
// 最后的117 ns/op表示每次需要话费117纳秒。(执行一次操作话费的时间)
func generateItems(n int) []Item {
	items := make([]Item, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, Item{id: i})
	}
	return items
}

type Item struct {
	id  int
	val [4096]byte
}
