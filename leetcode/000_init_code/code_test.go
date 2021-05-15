package _00_init_code

import (
	"fmt"
	"testing"
	"time"
)

type test struct {
	IntEs []int
	Lists []int
	K     int
}

var tests = []test{}

//

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		fmt.Println("初始化")

		for k := range tests[k1].IntEs {
			fmt.Println(tests[k1].IntEs[k])

		}
		pre := 1
		fmt.Println("结果：", pre)

	}
}

func Test_Chan2(t *testing.T) {

	ch := make(chan struct{}, 10)

	fmt.Println(123)
	go func() {
		<-ch
		fmt.Println(666)
	}()

	fmt.Println(456)
	time.Sleep(10 * time.Second)

}
