package main

import (
	"testing"
)

func addInt(a, b int) int {
	return a + b
}

func addString(a, b string) string {
	return a + b
}

// go test -v any_test.go -bench=. -benchmem -benchtime=3s

func BenchmarkWithoutGenerics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = addInt(1, 2)
		_ = addString("foo", "bar")
	}
}

//
// type Addable interface {
// 	int | string
// }
//
// func add[T Addable](a, b T) T {
// 	return a + b
// }

func BenchmarkWithGenerics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// _ = add(1, 2)
		// _ = add("foo", "bar")
	}
}
