package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkAddStr(b *testing.B) {
	str1 := "www"
	str := "eee"
	for i := 0; i < b.N; i++ {
		_ = str + str1
	}
	// _ = str
}

func BenchmarkSprintStr(b *testing.B) {
	str1 := "www"
	str := "eee"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s", str, str1)
	}
	// _ = str
}

func BenchmarkJoinStr(b *testing.B) {
	str1 := "www"
	str := "eee"
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{str, str1}, "")
	}
	// _ = str
}

func BenchmarkBufferStr(b *testing.B) {
	str1 := "www"
	str := "eee"
	for i := 0; i < b.N; i++ {
		var by bytes.Buffer
		by.WriteString(str1)
		by.WriteString(str)
		_ = by.String()
	}
	// _ = str

}

// BenchmarkAddStringWithOperator-8            50000000             30.3 ns/op
// BenchmarkAddStringWithSprintf-8             5000000              261  ns/op
// BenchmarkAddStringWithJoin-8                30000000             58.7 ns/op
// BenchmarkAddStringWithBuffer-8              2000000000           0.00 ns/op

// go test -bench='Str$' -run=none -cpu=8
// goos: darwin
// goarch: amd64
// pkg: testgo/modgo/test
// cpu: Intel(R) Core(TM) i7-6820HQ CPU @ 2.70GHz
// BenchmarkAddStr-8      	172992486	         7.022 ns/op
// BenchmarkSprintStr-8   	 8747796	       133.8 ns/op
// BenchmarkJoinStr-8     	33202748	        36.98 ns/op
// BenchmarkBufferStr-8   	23436049	        50.49 ns/op
