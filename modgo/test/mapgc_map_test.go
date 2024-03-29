package main

import (
	"fmt"
	"runtime"
	"testing"
)

func TestMapGC(t *testing.T) {
	v := "struct{} "
	a := make(map[int]string)
	for i := 0; i < 10000; i++ {
		a[i] = v
	}
	runtime.GC()
	fmt.Printf("%d\n", len(a))

	printMemStats("After Map Add 100000")

	for i := 0; i < 10000-1; i++ {
		delete(a, i)
	}
	runtime.GC()
	fmt.Printf("%d\n", len(a))
	printMemStats("After Map Delete 9999")

	for i := 0; i < 10000-1; i++ {
		a[i] = v
	}

	runtime.GC()
	fmt.Printf("%d\n", len(a))

	printMemStats("After Map Add 9999 again")

	a = nil
	runtime.GC()
	fmt.Printf("%d\n", len(a))

	printMemStats("After Map Set nil")
}

func printMemStats(mag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%v：memory = %vKB, GC Times = %v\n", mag, m.Alloc/1024, m.NumGC)
}
