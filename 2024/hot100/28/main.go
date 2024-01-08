package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := 0
	newList := &ListNode{}
	cur := newList
	for l1 != nil && l2 != nil {
		cur.Next = &ListNode{Val: (l1.Val + l2.Val + tmp) % 10}
		cur = cur.Next
		tmp = (l1.Val + l2.Val + tmp) / 10
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		cur.Next = &ListNode{Val: (l1.Val + tmp) % 10}
		cur = cur.Next
		tmp = (l1.Val + tmp) / 10
		l1 = l1.Next
	}
	for l2 != nil {
		cur.Next = &ListNode{Val: (l2.Val + tmp) % 10}
		cur = cur.Next
		tmp = (l2.Val + tmp) / 10
		l2 = l2.Next
	}

	if tmp != 0 {
		cur.Next = &ListNode{Val: tmp}
	}
	return newList.Next
}
