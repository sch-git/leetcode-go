package _06

//simple

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var root, next *ListNode
	for head != nil {
		root = &ListNode{
			Val:  head.Val,
			Next: next,
		}
		next = root
		head = head.Next
	}

	return root
}
