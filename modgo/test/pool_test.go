package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"sync"
	"testing"

	json "github.com/json-iterator/go"
)

type Student struct {
	Name   string
	Info   map[string]SSSS
	Remark [1024]byte
}

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

func randomBytes(n int) [1024]byte {
	b := make([]byte, n)
	rand.Read(b)
	var array [1024]byte
	copy(array[:], b)
	return array
}

type SSS struct {
	M map[string]any
}
type SSSS struct {
	MM map[string]SSS
}

func bufs() []byte {

	a := make(map[string]SSSS, 0)
	a["aa"] = SSSS{MM: map[string]SSS{"bb": {M: map[string]any{"bb": 1}}}}

	student := &Student{
		Name:   "randomString(10)",
		Info:   a,
		Remark: randomBytes(1024),
	}
	fmt.Println(student)
	buf, _ := json.Marshal(student)
	return buf
}

var buf = bufs()

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

var data = make([]byte, 10000)

func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
		buf.Reset()
		bufferPool.Put(buf)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
	}
}

/*

go test -v pool_test.go -bench . -benchmem -benchtime 5s  -cpu=1

goos: darwin
goarch: arm64
BenchmarkUnmarshal         	  278158	     18855 ns/op	    2152 B/op	      25 allocs/op
BenchmarkUnmarshalWithPool 	  326536	     18598 ns/op	     736 B/op	      21 allocs/op
BenchmarkBufferWithPool    	42681146	     140.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuffer            	 8205529	     712.6 ns/op	   10240 B/op	       1 allocs/op
PASS
ok  	command-line-arguments	24.790s

goos 和 goarch：这两行表示Go语言运行时的操作系统和架构。在这个例子中，操作系统（goos）是darwin（即macOS），架构（goarch）是arm64。
278158：表示在5秒内（由-benchtime 5s决定）这个测试被运行了278,158次。测试次数越多，结果越可靠。
18855 ns/op：表示每个操作（每次测试）的平均耗时是18,855纳秒。
2152 B/op：表示每次操作平均分配的内存是2152字节。这是由-benchmem标志启用的。
25 allocs/op：表示每次操作平均进行了25次内存分配。

Benchmark：基准函数的名称。
N：为基准执行的迭代次数。
ns/op：运行基准测试的一次迭代所花费的平均纳秒数。
B/op：在基准测试的每次迭代期间分配的平均字节数。
allocs/op：在基准测试的每次迭代期间分配内存的平均次数。



*/
