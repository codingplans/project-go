package main

import (
	"bytes"
	"io"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func genRandString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	fn := func(length int, charset string) string {
		b := make([]byte, length)
		for i := range b {
			b[i] = charset[seededRand.Intn(len(charset))]
		}
		return string(b)
	}
	return fn(length, charset)
}

func TestIoCopy(t *testing.T) {
	s := genRandString(137)
	r := strings.NewReader(s)
	b := bytes.NewBufferString("")
	_, err := io.Copy(b, r)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(b.String())

}
func BenchmarkIoCopy(b *testing.B) {
	s := genRandString(1024 * 1024 * 10)
	for i := 0; i < b.N; i++ {
		r := strings.NewReader(s)
		bs := bytes.NewBufferString("")
		_, err := io.Copy(bs, r)
		if err != nil {
			b.Error(err.Error())
		}
		_ = bs.String()
		// b.Log()
	}
}

func BenchmarkIoReadall(b *testing.B) {
	s := genRandString(1024 * 1024 * 10)
	for i := 0; i < b.N; i++ {
		r := strings.NewReader(s)
		// bs := bytes.NewBufferString("")
		bs, err := io.ReadAll(r)
		_ = string(bs)
		if err != nil {
			b.Error(err.Error())
		}
	}
}

// go test -bench='BenchmarkIo' -run=none -cpu=8  -benchmem
// goos: darwin
// goarch: arm64
// pkg: testgo/modgo/test
// BenchmarkIoCopy-8      	     294	   3919850 ns/op	21042950 B/op	       4 allocs/op
// BenchmarkIoReadall-8   	     123	   8810197 ns/op	62919508 B/op	      39 allocs/op
// PASS
// ok  	testgo/modgo/test	4.872s
