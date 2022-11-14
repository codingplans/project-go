package main

import (
	"testing"
)

type appendBaz struct {
	bar int
	foo int
}

func BenchmarkAppendFor(b *testing.B) {

	m := make(map[int][]*appendBaz)
	for i := 0; i < 100000; i++ {
		m[1] = append(m[1], &appendBaz{bar: i, foo: i})
	}
	_ = len(m[1])
	// fmt.Println()
}

func BenchmarkAppendSlice(b *testing.B) {
	m := make(map[int][]*appendBaz)
	sls := make([]*appendBaz, 0, 100000)
	for i := 0; i < 100000; i++ {
		sls = append(sls, &appendBaz{bar: i, foo: i})
	}
	m[1] = sls
	_ = len(m[1])
	// fmt.Println(len(m[1]))
}

//  go test append_test.go -bench=. -run=none -benchtime=10s
// goos: darwin
// goarch: arm64
// BenchmarkAppendFor-8            1000000000               0.004315 ns/op
// BenchmarkAppendSlice-8          1000000000               0.001471 ns/op
// PASS
// ok      command-line-arguments  0.513s
