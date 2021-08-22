package _25

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	head := root
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			root.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			root.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}

		root = root.Next
	}

	for l1 != nil {
		root.Next = &ListNode{Val: l1.Val}
		root = root.Next
		l1 = l1.Next
	}

	for l2 != nil {
		root.Next = &ListNode{Val: l2.Val}
		root = root.Next
		l2 = l2.Next
	}

	return head.Next
}
