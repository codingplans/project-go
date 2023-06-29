package main

import (
	"math/rand"
	"testing"
	"time"
)

func randomStr() {
	for i := 0; i < 10; i++ {
		strs = append(strs, randomStrings(50))
	}
}

var strs []string

// 比较 len 的性能问题
func BenchmarkStringLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strs := randomStrings(50)
		if 0 == len(strs) {
			_ = i
		}
	}
}
func BenchmarkStringEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strs := randomStrings(50)
		if strs == "" {
			_ = i
		}
	}
}

func BenchmarkStringLenNull(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strs := ""
		if 0 == len(strs) {
			_ = i
		}
	}
}
func BenchmarkStringEmptyNull(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strs := ""
		if strs == "" {
			_ = i
		}
	}
}

/*

 go test -v stringlen_test.go  -bench=. -run=none -benchtime=5s
goos: darwin
goarch: arm64
BenchmarkStringLen
BenchmarkStringLen-8     	  555457	     10226 ns/op
BenchmarkStringEmpty
BenchmarkStringEmpty-8   	  582540	     10288 ns/op
PASS
ok  	command-line-arguments	17.082s



*/

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func stringWithCharset(length int, charset string) string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	if length > 50 {
		length = 50
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func randomStrings(length int) string {
	return stringWithCharset(length, charset)
}
