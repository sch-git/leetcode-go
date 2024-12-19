package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := `{"is_readcamp":true}`
	var obj OrderContext
	err := json.Unmarshal([]byte(str), &obj)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Printf("%v", *obj.IsReadcamp)

}

type OrderContext struct {
	IsReadcamp *bool `json:"is_readcamp"`
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
