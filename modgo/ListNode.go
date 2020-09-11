package main

import log "github.com/sirupsen/logrus"

// Definition for singly-linked list.

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func main() {

	h := &ListNode2{
		1, nil,
	}
	for i := 1; i < 10; i++ {
		h.addCycle(&ListNode2{
			i, nil,
		})
	}

	// aa := hasCycle()
	aa := detectCycle(h)
	log.Info(aa)
}
func (l *ListNode2) addCycle(head *ListNode2) {

	for l.Next != nil {
		l = l.Next
	}
	l.Next = head
}

func hasCycle(head *ListNode2) bool {
	if head == nil || head.Next == nil {
		return false
	}
	s := head
	k := head
	for k.Next != nil && k.Next.Next != nil {
		log.Info(k.Val)
		k = k.Next.Next
		s = s.Next
		if s == k {
			return true
		}
	}
	return false
}

func detectCycle(head *ListNode2) *ListNode2 {

	if head == nil || head.Next == nil {
		// return "no cycle"
	}

	return head
}
