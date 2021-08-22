package _18

// simple

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	pre := &ListNode{
		Next: head,
	}
	root := pre
	for head != nil {
		if head.Val == val {
			pre.Next = head.Next
		}
		pre = pre.Next
		head = head.Next
	}

	return root.Next
}
