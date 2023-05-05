package main

import (
	"fmt"
	"runtime"
	"time"
)

// 创建了一个极简的例子，A，B, C 实例都设置了 Finalizer 回调，故意让其中一个阻塞住，会影响到剩下的 Finalizer 的执行。

//Go 提供的 Finalizer 机制，让程序员创建的时候注册回调函数，能很好的帮助程序员解决资源安全释放的问题；
//Finalizer 的执行是全局单协程，且串行化执行的。所以可能会因为某一次的卡住导致全局的失效，切记；
//排查内存问题的时候，pprof 看现场很明确，但是根因可能是看似毫不相关的旮旯角落，有时候要把思维跳出来排查；

var (
	done chan struct{} = make(chan struct{})
)

type A struct {
	name string
}

type B struct {
	name string
}

type C struct {
	name string
}

// 由执行顺序决定  B阻塞了，所以A也无法释放
func newA() *A {
	v := &A{"n1"}
	runtime.SetFinalizer(v, func(p *A) {
		fmt.Println("gc Finalizer A")
	})
	return v
}

func newB() *B {
	v := &B{"n1"}
	runtime.SetFinalizer(v, func(p *B) {
		<-done
		fmt.Println("gc Finalizer B")
	})
	return v
}

func newC() *C {
	v := &C{"n1"}
	runtime.SetFinalizer(v, func(p *C) {
		fmt.Println("gc Finalizer C")
	})
	return v
}

func main() {
	a := newA()
	b := newB()
	c := newC()
	fmt.Println("== start ===")
	//_, _, _ = a, b, c
	_, _, _ = c, b, a
	fmt.Println("== ... ===")
	for i := 0; i < 10; i++ {
		//time.Sleep(time.Millisecond * 1)

		runtime.GC()
	}
	fmt.Println("== end ===")
	time.Sleep(time.Second * 1)
	//done = make(chan struct{})
	done <- struct{}{}
	fmt.Println("== real end ===")

	time.Sleep(time.Second * 10)
}

/*

输出：
== start ===
== ... ===
gc Finalizer C
== end ===
== real end ===
gc Finalizer B
gc Finalizer A

*/
