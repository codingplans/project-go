package main

import (
	"fmt"
	"golang.org/x/net/context"
	"runtime"
	"sync"
	"time"
)

func main() {
	// chann()

	a1 := []int{1, 3, 5, 7, 9}
	b1 := []int{2, 4, 6, 8, 10}
	c1 := []int{}

	// c1 := []int{ 1 ,1.5,2,3,4,5,6,7}

	n1 := 0
	m1 := len(a1)
	m2 := len(a1)
	k := len(a1) + len(b1)

	fmt.Println("排序前：", c1)

	for n1 < m2 && m1 < k && len(c1) <= k {
		if a1[n1] < b1[m1-m2] {
			c1 = append(c1, a1[n1])
			n1++
		} else {
			c1 = append(c1, b1[m1-m2])
			m1++
		}
	}

	for n1 < m2 {
		c1 = append(c1, a1[n1])
		n1++
	}
	for m1 < k {
		c1 = append(c1, b1[m1-m2])
		m1++
	}

	fmt.Println("排序后：", c1)

}

func chann() {
	fmt.Printf("%d", runtime.GOMAXPROCS)
	var wg sync.WaitGroup
	qq := make(chan bool)
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	testChan(ctx, qq, wg)
	if _, err := <-qq; err {
		println(" chan is open")
		qq <- true
	}
	wg.Wait()
}

func testChan(ctx context.Context, quit chan bool, wg sync.WaitGroup) {
	defer close(quit)
	defer wg.Done()
	// defer close(ctx)
	for {
		select {
		case <-ctx.Done():
			println("stop chan")
			return
		case OK, _ := <-quit:
			println(OK, 123)
			return
		default:
			println("sleep func")
			time.Sleep(time.Second)
		}
	}
}
