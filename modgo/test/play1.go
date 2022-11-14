package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type baz struct {
	bar int
	foo int
}
type arrStruct []baz

func main() {
	arr := make(arrStruct, 0, 1000)
	for i := 1; i < 1000; i++ {
		arr = append(arr, baz{bar: i, foo: i})
	}

	// fn := func(k int, v baz, list arrStruct) {
	// 	for i, _ := range list {
	// 		if i == k {
	// 			continue
	// 		}
	// 		if list[i] != v {
	// 			fmt.Println("not match ", list[i], v, i)
	// 		}
	// 	}
	//
	// }

	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(len(arr) + 1)
	ar := []int{}
	for i := range arr {
		// go fn(i, arr[i], arr)
		s := i
		go func(k int) {
			defer wg.Done()
			fmt.Println(k, s)
			ar = append(ar, k)
		}(i)

	}

	ch := make(chan struct{})
	go func() {

		time.Sleep(2 * time.Second)
		ch <- struct{}{}
		wg.Done()
	}()
	<-ch
	wg.Wait()
	fmt.Println(len(ar), 444)

}
