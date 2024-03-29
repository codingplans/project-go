package main

import (
	"runtime"
	"testing"
)

func lastNumsBySlice(origin []int) []int {
	return origin[len(origin)-2:]
}

func lastNumsByCopy(origin []int) []int {
	result := make([]int, 2)
	copy(result, origin[len(origin)-2:])
	return result
}

func printMem(t *testing.T) {
	t.Helper()
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	t.Logf("%.2f MB", float64(rtm.Alloc)/1024./1024.)
}

func testLastChars(t *testing.T, f func([]int) []int) {
	t.Helper()
	ans := make([][]int, 0)
	for k := 0; k < 100; k++ {
		origin := generateWithCap(128 * 1024) // 1M
		ans = append(ans, f(origin))
	}
	printMem(t)
	_ = ans
}

// 大量内存得不到释放
// go test -v -run="TestLastChars"
func TestLastCharsBySlice(t *testing.T) { testLastChars(t, lastNumsBySlice) }
func TestLastCharsByCopy(t *testing.T)  { testLastChars(t, lastNumsByCopy) }

/*
=== RUN   TestLastCharsBySlice
    TestLastChars_test.go:46: 100.15 MB
--- PASS: TestLastCharsBySlice (0.28s)
=== RUN   TestLastCharsByCopy
    TestLastChars_test.go:47: 3.15 MB
--- PASS: TestLastCharsByCopy (0.26s)
PASS
ok  	testgo/modgo/test	0.620s
*/
