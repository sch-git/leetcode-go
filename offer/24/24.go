package _24

// simple

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	root := &ListNode{}
	for head != nil {
		temp := &ListNode{
			Val: head.Val,
		}
		temp.Next = root.Next
		root.Next = temp
		head = head.Next
	}

	return root.Next
}
