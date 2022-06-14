package main

import (
	"testing"
)

// 比较 len 的性能问题
// go test -v -run="*"
func BenchmarkStringLen(b *testing.B) {
	str := ""
	for i := 0; i < b.N; {
		if 0 == len(str) {
			i++
		}
	}
}
func BenchmarkStringEmpty(b *testing.B) {
	str := ""
	for i := 0; i < b.N; {
		if str == "" {
			i++
		}
	}
}
