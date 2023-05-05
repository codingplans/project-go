package main

import (
	"unsafe"
)

//go:linkname overrideWrite runtime.overrideWrite
var overrideWrite func(fd uintptr, p unsafe.Pointer, n int32) int32

//go:linkname write1 runtime.write1
func write1(fd uintptr, p unsafe.Pointer, n int32) int32

//go:nosplit
func print2stdout(fd uintptr, p unsafe.Pointer, n int32) int32 {
	return write1(1, p, n)
}

func main() {
	// 默认stderr
	println("xx", "yy", "zz", false, '?', 0.75, 9)
	// 重定向到stdout
	overrideWrite = print2stdout
	println("xx", "yy", "zz", false, '?', 0.75, 9)
	// 恢复stderr
	overrideWrite = nil
	println("xx", "yy", "zz", false, '?', 0.75, 9)
}
