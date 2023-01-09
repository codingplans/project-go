package main

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func ConvertSliceUint32ToSet(dataList []uint32) map[uint32]bool {
	res := make(map[uint32]bool)
	for _, data := range dataList {
		res[data] = true
	}
	return res
}

func Uint32Intersect(l1 []uint32, l2 []uint32) []uint32 {
	l2Map := ConvertSliceUint32ToSet(l2)
	r := make([]uint32, 0)
	for _, l := range l1 {
		if _, ok := l2Map[l]; ok {
			r = append(r, l)
		}
	}
	return r
}

func Uint32IntersectV2(s1 []uint32, s2 []uint32) (res []uint32) {
	set := make(map[uint32]struct{}, len(s1))
	for _, v := range s1 {
		set[v] = struct{}{}
	}
	for _, v := range s2 {
		if _, ok := set[v]; ok {
			res = append(res, v)
		}
	}

	if res == nil {
		res = make([]uint32, 0)
	}

	return res
}

type Uint32Slice []uint32

func (p Uint32Slice) Len() int           { return len(p) }
func (p Uint32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Uint32IntersectV3(s1 []uint32, s2 []uint32) (res []uint32) {
	sort.Sort(Uint32Slice(s1))
	sort.Sort(Uint32Slice(s2))
	for i, j := 0, 0; i < len(s1) && j < len(s2); {
		x, y := s1[i], s2[j]
		if x == y {
			if res == nil || x > res[len(res)-1] {
				res = append(res, x)
			}
		} else if x < y {
			i++
		} else {
			j++
		}
	}

	if res == nil {
		res = make([]uint32, 0)
	}

	return res
}

var mockSlice = func(l int) []uint32 {
	r := make([]uint32, 0, l)
	for i := 0; i < l; i++ {
		r = append(r, rand.Uint32())
	}
	return r
}

var (
	aSlice10   = mockSlice(10)
	bSlice10   = mockSlice(10)
	aSlice100  = mockSlice(100)
	bSlice100  = mockSlice(100)
	aSlice1000 = mockSlice(1000)
	bSlice1000 = mockSlice(1000)
)

func TestUint32Intersect(t *testing.T) {
	expect1 := Uint32Intersect(aSlice10, bSlice10)
	got1 := Uint32IntersectV2(aSlice10, bSlice10)
	got2 := Uint32IntersectV3(aSlice10, bSlice10)

	expect2 := Uint32Intersect(aSlice100, bSlice100)
	got3 := Uint32IntersectV2(aSlice100, bSlice100)
	got4 := Uint32IntersectV3(aSlice100, bSlice100)

	expect3 := Uint32Intersect(aSlice1000, bSlice1000)
	got5 := Uint32IntersectV2(aSlice1000, bSlice1000)
	got6 := Uint32IntersectV3(aSlice1000, bSlice1000)

	require.EqualValues(t, expect1, got1)
	require.EqualValues(t, expect1, got2)
	require.EqualValues(t, expect2, got3)
	require.EqualValues(t, expect2, got4)
	require.EqualValues(t, expect3, got5)
	require.EqualValues(t, expect3, got6)
}

func BenchmarkUint32Intersect_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint32Intersect(aSlice10, bSlice10)
	}
}

func BenchmarkUint32Intersect_10_V2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint32IntersectV2(aSlice10, bSlice10)
	}
}

func BenchmarkUint32Intersect_10_V3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint32IntersectV3(aSlice10, bSlice10)
	}
}

func BenchmarkUint32Intersect_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint32Intersect(aSlice100, bSlice100)
	}
}

func BenchmarkUint32Intersect_100_V2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint32IntersectV2(aSlice100, bSlice100)
	}
}

func BenchmarkUint32Intersect_100_V3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint32IntersectV3(aSlice100, bSlice100)
	}
}

func BenchmarkUint32Intersect_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint32Intersect(aSlice1000, bSlice1000)
	}
}

func BenchmarkUint32Intersect_1000_V2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint32IntersectV2(aSlice1000, bSlice1000)
	}
}

func BenchmarkUint32Intersect_1000_V3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint32IntersectV3(aSlice1000, bSlice1000)
	}
}

// 结论， V3 最快， V2 次之， V1 最慢 ，先排序后，再做交集。

// 测试结果 go version go1.18.6 darwin/arm64
/*
 go test merge_slice_test.go  -bench=. -run=none -benchtime=10s  -benchmem  -cpu=1 -cpuprofile=cpu.profile -gcflags=-m

goos: darwin
goarch: arm64
BenchmarkUint32Intersect_10             38378931               396.1 ns/op           114 B/op          1 allocs/op
BenchmarkUint32Intersect_10_V2          40396842               293.9 ns/op            98 B/op          1 allocs/op
BenchmarkUint32Intersect_10_V3          89563322               135.6 ns/op            48 B/op          2 allocs/op
BenchmarkUint32Intersect_100             2537643              5003 ns/op            2193 B/op         16 allocs/op
BenchmarkUint32Intersect_100_V2          2922080              3508 ns/op            1030 B/op          6 allocs/op
BenchmarkUint32Intersect_100_V3          4460221              2706 ns/op              48 B/op          2 allocs/op
BenchmarkUint32Intersect_1000             213469             58191 ns/op           35041 B/op         59 allocs/op
BenchmarkUint32Intersect_1000_V2          402424             29420 ns/op           13705 B/op          6 allocs/op
BenchmarkUint32Intersect_1000_V3          223093             45618 ns/op              48 B/op          2 allocs/op
PASS
ok      command-line-arguments  123.139s


*/
