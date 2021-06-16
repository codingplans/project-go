package _00_init_code

import (
	"fmt"
	"testing"
)

type test struct {
	IntEs []int
	Lists []int
	K     int
	K2    int
	Str   string
	Str2  string
}

var tests = []test{}

//

func Test_upToDayUp(t *testing.T) {
	for _, v := range tests {
		fmt.Println("初始化")

		fmt.Println(v.IntEs)
		pre := 1
		fmt.Println("结果：", pre)
	}

}
