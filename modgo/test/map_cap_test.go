package main

import (
	"testing"
)

func MapCap2() {
	m := make(map[int]int, 100000)
	for i := 0; i < 100000; i++ {
		m[i] = i
	}
}
func MapCap1() {
	m := make(map[int]int)
	for i := 0; i < 100000; i++ {
		m[i] = i
	}
}

func MapCap3() {
	m := make(map[int]int, 100)
	for i := 0; i < 100; i++ {
		m[i] = i
	}
}
func MapCap4() {
	m := make(map[int]int)
	for i := 0; i < 100; i++ {
		m[i] = i
	}
}

func BenchmarkMapCap2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MapCap2()
	}
}

func BenchmarkMapCap1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MapCap1()
	}
}

func BenchmarkMapCap3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MapCap3()
	}
}

func BenchmarkMapCap4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MapCap4()
	}
}

/*
可以看出 给map初始化时候加上cap，可以提高性能
 go test map_cap_test.go  -bench=.
goos: darwin
goarch: arm64
BenchmarkMapCap2-8           433           3069372 ns/op
BenchmarkMapCap1-8           241           4419905 ns/op
BenchmarkMapCap3-8        528091              2248 ns/op
BenchmarkMapCap4-8        327957              3762 ns/op
PASS
ok      command-line-arguments  7.567s


*/
