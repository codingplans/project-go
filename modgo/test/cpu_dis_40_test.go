package main

import (
	"math/rand"
	"testing"
	"time"
)

// 让cpu 利用率提升40%
type cpuDisBaz struct {
	bar int
	foo int
}

var cpuList []*cpuDisBaz

func init() {
	cpuList = make([]*cpuDisBaz, 0, 1000000)
	for i := 0; i < 1000000; i++ {
		cpuList = append(cpuList, &cpuDisBaz{bar: i, foo: i})
	}
}

func BenchmarkCpuDis(b *testing.B) {
	for i := 0; i < 100000; i++ {
		bar := foo1()
		_ = bar
	}
}
func foo1() *cpuDisBaz {
	rand.Seed(time.Now().UnixNano())
	t := rand.Intn(100000)
	rule := cpuList[t]
	return rule
}

// 然后我看到了——通过分配 rule := r[i]，我们将堆分配的 Rule 从 slice 复制到堆栈上，然后通过返回 &rule 我们创建了一个指向结构副本的（转义）指针。幸运的是，解决这个问题很容易。我们只需要将 & 符号向上移动一点，这样我们就可以引用切片中的结构而不是复制它：

func foo2() *cpuDisBaz {
	rand.Seed(time.Now().UnixNano())
	t := rand.Intn(100000)
	rule := &cpuList[t]
	return *rule
}

func BenchmarkCpuAdd(b *testing.B) {

	for i := 0; i < 100000; i++ {
		bar := foo2()
		_ = bar
	}
}

//  go test append_test.go -bench=. -run=none -benchtime=10s
// goos: darwin
// goarch: arm64
// BenchmarkAppendFor-8            1000000000               0.004315 ns/op
// BenchmarkAppendSlice-8          1000000000               0.001471 ns/op
// PASS
// ok      command-line-arguments  0.513s
