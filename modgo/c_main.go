package main

// 最简单的一个C程序
// CGO_ENABLED=1 go run c_main.go

/*
#include <stdio.h>

void printint(int v) {
    printf("printint: %d\n", v);
}
*/
import "C"

func main() {
	v := 42
	C.printint(C.int(v))
}
