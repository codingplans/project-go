package main

import (
	"fmt"
	"github.com/Darrenzzy/testgo/structures"
	"testing"
)

func TestReverse(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6, 7}
	list := structures.Ints2List(arr)
	list = Reverse(list)
	fmt.Println(structures.List2Ints(list))

}

func Reverse(pHead *structures.ListNode) *structures.ListNode {
	m := pHead
	newList := &structures.ListNode{}

	for m != nil {
		nt := m.Next
		m.Next = newList
		newList.Next = m
		m = nt
	}
	return newList
}
