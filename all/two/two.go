package two

type ListNode struct {
	Val  int       `json:"val"`
	Next *ListNode `json:"next"`
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	l := root
	var incr int
	for l1 != nil && l2 != nil {
		l.Next = &ListNode{
			Val: (l1.Val + l2.Val + incr) % 10,
		}
		incr = (l1.Val + l2.Val + incr) / 10
		l = l.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		l.Next = &ListNode{
			Val: (l1.Val + incr) % 10,
		}
		incr = (l1.Val + incr) / 10
		l = l.Next
		l1 = l1.Next
	}

	for l2 != nil {
		l.Next = &ListNode{
			Val: (l2.Val + incr) % 10,
		}
		incr = (l2.Val + incr) / 10
		l = l.Next
		l2 = l2.Next
	}

	if incr != 0 {
		l.Next = &ListNode{Val: incr}
	}

	return root.Next
}
