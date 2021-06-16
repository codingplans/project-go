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

var tests = []test{
	{IntEs: []int{1, 2, 43, 7, 3, 33, 1, 8, 12}},
}

//

func Test_upToDayUp(t *testing.T) {
	for _, v := range tests {
		fmt.Println("初始化")

		fmt.Println(v.IntEs)
		pre := getLeastNumbers(v.IntEs, 2)

		fmt.Println("结果：", pre)
	}

}
func getLeastNumbers(arr []int, k int) []int {
	l := len(arr)
	for i := (len(arr)) >> 1; i >= 0; i-- {
		buildHeap(arr, i, len(arr))
	}
	l--
	// println(2222)
	// 前 k 个 top
	for i := l; i >= 0 && l >= len(arr)-k; {
		// println(len(arr)-k, i, l)
		arr[0], arr[l] = arr[l], arr[0]
		buildHeap(arr, 0, l)
		l--
	}
	return arr[len(arr)-k:]
}

func buildHeap(arr []int, m, l int) {
	mid := m
	for m < l {
		i := m*2 + 1
		j := i + 1
		if i < l && arr[mid] > arr[i] {
			mid = i
		}
		if j < l && arr[mid] > arr[j] {
			mid = j
		}
		if mid != m {
			arr[m], arr[mid] = arr[mid], arr[m]
			fmt.Println(arr, mid, m)
			m = mid
		} else {
			break
		}

	}
}
