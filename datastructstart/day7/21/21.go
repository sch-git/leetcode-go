package _1

// simple

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	root := new(ListNode)
	cpRoot := root
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			cpRoot.Next = &ListNode{
				Val: l2.Val,
			}
			l2 = l2.Next
		} else {
			cpRoot.Next = &ListNode{
				Val: l1.Val,
			}
			l1 = l1.Next
		}

		cpRoot = cpRoot.Next
	}

	for l1 != nil {
		cpRoot.Next = &ListNode{
			Val: l1.Val,
		}
		cpRoot = cpRoot.Next
		l1 = l1.Next
	}

	for l2 != nil {
		cpRoot.Next = &ListNode{
			Val: l2.Val,
		}
		cpRoot = cpRoot.Next
		l2 = l2.Next
	}

	return root.Next
}
