package main

import (
	"strconv"
	"testing"
)

func Benchmark_sliceRange(b *testing.B) {
	bb := []string{}
	for i := 0; i < 100000; i++ {
		bb = append(bb, strconv.Itoa(i)+"12")
	}
	for i := 0; i < b.N; i++ {
		a := []string{}
		sliceRange(a, bb)
	}
}

func Benchmark_sliceOnce(b *testing.B) {
	bb := []string{}
	for i := 0; i < 100000; i++ {
		bb = append(bb, strconv.Itoa(i)+"12")
	}
	for i := 0; i < b.N; i++ {
		a := []string{}
		sliceOnce(a, bb)
	}
}

func sliceRange(a, b []string) {
	for k := range b {
		a = append(a, b[k])
	}

}

func sliceOnce(a, b []string) {
	a = append(a, b...)
}

/*
从跑分结果来看： 推荐直接 arr... 来扩容 slice

// 结果一样 BenchmarkSliceOnce 推荐
go test  for_slice_test.go -bench=. -run=none -benchtime=10s
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-6820HQ CPU @ 2.70GHz
BenchmarkSliceRange-8   	   17137	    743108 ns/op
BenchmarkSliceOnce-8    	  125911	     96495 ns/op
PASS
ok  	command-line-arguments	33.049s


下面这个是把 slice 类型换成字符串 结果一样 BenchmarkSliceOnce 推荐
go test  for_slice_test.go -bench=. -run=none -benchtime=3s
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-6820HQ CPU @ 2.70GHz
BenchmarkSliceRange-8   	     412	   7389596 ns/op
BenchmarkSliceOnce-8    	    2719	   1333876 ns/op
PASS
ok  	command-line-arguments	11.030s

下面这个是把 slice 类型换成字符串 结果一样 BenchmarkSliceOnce 推荐
go test  for_slice_test.go -bench=. -run=none -benchtime=3s
goos: darwin
goarch: arm64
Benchmark_sliceRange-8              1081           3084000 ns/op
Benchmark_sliceOnce-8               5752            606444 ns/op
PASS
ok      command-line-arguments  7.612s

*/

