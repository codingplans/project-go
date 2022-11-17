package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type bazz struct {
	bar int
	foo int
}
type arrStructs []bazz

func main1() {
	arr := make(arrStructs, 0, 1000)
	for i := 1; i < 1000; i++ {
		arr = append(arr, bazz{bar: i, foo: i})
	}

	// fn := func(k int, v bazz, list arrStructs) {
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
