package main

import (
	"testing"
	"time"

	"github.com/RoaringBitmap/roaring/roaring64"
)

func bitMapToSlice(bitmap *roaring64.Bitmap) []int64 {
	maxLen := 4096
	total := make([]int64, 0, maxLen)
	buf := make([]uint64, maxLen)
	bmIter := bitmap.ManyIterator()
	for n := bmIter.NextMany(buf); n != 0; n = bmIter.NextMany(buf) {
		for _, v := range buf[:n] {
			total = append(total, int64(v))
		}
	}
	return total
}
func BitMap1() {
	BitMap := roaring64.NewBitmap()
	ts := time.Now()
	l := 100000000
	for i := 0; i < l; i++ {
		BitMap.Add(uint64(i))
	}
	println(time.Since(ts).Microseconds(), "使用Bitmap方式")
	keys := []uint64{3333333, 63666666, 1233, 6224}

	for _, key := range keys {
		if BitMap.Contains(key) {
			println(time.Since(ts).Microseconds())
		}
	}

}

func BitMapV21() {
	l := 100000000
	arr := []uint64{}
	ts := time.Now()

	for i := 0; i < l; i++ {
		arr = append(arr, uint64(i))
	}
	println(time.Since(ts).Microseconds(), "使用切片方式")
	keys := []uint64{3333333, 63666666, 1233, 6224}

	for _, key := range keys {
		if IntInArray(key, keys) {
			println(time.Since(ts).Microseconds())
		}
	}
}

func IntInArray(val uint64, array []uint64) (exists bool) {
	exists = false
	if val < 6000000 {
		for _, v := range array {
			if val == v {
				exists = true
				return
			}
		}
	} else {

		l := len(array)
		for k := range array {
			if val == array[l-k-1] {
				exists = true
				return
			}
		}
	}

	return
}
func BenchmarkBitMapV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitMap1()
	}
}

func BenchmarkBitMapV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitMapV21()
	}

}

func TestBitMapV1(t *testing.T) {
	BitMap1()
}

func TestBitMapV2(t *testing.T) {
	BitMapV21()
}

// go test -v bitmap_test.go  -benchmem -benchtime=5s  -bench . -run=none
