package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
)

func main() {
	a := [4][4]int{
		{0, 1, 2, 3},       /*  第一行索引为 0 */
		{4, 5, 6, 7},       /*  第二行索引为 1 */
		{8, 9, 10, 11},     /* 第三行索引为 2 */
		{81, 91, 101, 111}, /* 第三行索引为 2 */
	}

	machineName, err := os.Hostname()
	buf := md5.Sum([]byte(machineName))

	var staticIncrement int64
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		defer wg.Done()
		fmt.Sprint(atomic.AddInt64(&staticIncrement, 1))
	}()
	go func() {
		wg.Add(1)
		defer wg.Done()
		fmt.Sprint(atomic.AddInt64(&staticIncrement, 1))
	}()
	go func() {
		wg.Add(1)
		defer wg.Done()
		fmt.Sprint(atomic.AddInt64(&staticIncrement, 1))
	}()

	fmt.Println(machineName, err, buf, getMachineHash(), staticIncrement)
	fmt.Sprint(atomic.AddInt64(&staticIncrement, 1))

	println(a[1][2], 1, staticIncrement)
	wg.Wait()
	return

}
func getMachineHash() int32 {
	machineName, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	buf := md5.Sum([]byte(machineName))
	return (int32(buf[0])<<0x10 + int32(buf[1])<<8) + int32(buf[2])
}
