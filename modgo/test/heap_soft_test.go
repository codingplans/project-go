package main

import "testing"

func TestHeapSoft(t *testing.T) {

	arr1 := []int{1, 8, 0, 4, 3, 2, 7, 6, 5}
	t.Log(arr1)
	heapify(arr1)
	t.Log(arr1)
}

func heapSoft(arr []int, i, n int) {
	l := 2*i + 1
	r := 2*i + 2
	largest := i
	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapSoft(arr, largest, n)
	}
	return

}
func buildTopHeap(arr []int, n int) {
	// 倒数第二层开始 往顶点比较构造大顶堆
	for i := n/2 - 1; i >= 0; i-- {
		heapSoft(arr, i, n)
	}
	return

}

func heapify(arr []int) {
	n := len(arr)
	buildTopHeap(arr, n)
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapSoft(arr, 0, i)
	}
	return
}
