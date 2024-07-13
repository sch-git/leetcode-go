package main

import (
	"fmt"
)

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  2,
				Next: &ListNode{Val: 1},
			},
		},
	}
	fmt.Println(check(list))
	list = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{Val: 23, Next: &ListNode{Val: 1}},
			},
		},
	}
	fmt.Println(check(list))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 1
// 1 2 3 2 1
// 1 2 1
// 1 2 2 1
// 1 2 3
func check(h *ListNode) bool {
	hLen := 0
	curH := h
	for curH != nil {
		hLen++
		curH = curH.Next
	}

	mid := hLen/2 - 1
	curH = &ListNode{Next: h}
	secondH := h
	for ; mid > 0; mid-- {
		node := h.Next
		h.Next = node.Next
		node.Next = h
		curH.Next = node
		secondH = h.Next
	}
	h = curH.Next
	if hLen%2 != 0 {
		secondH = secondH.Next
	}
	for secondH != nil {
		if h.Val != secondH.Val {
			return false
		}
		secondH = secondH.Next
		h = h.Next
	}
	return true
}
