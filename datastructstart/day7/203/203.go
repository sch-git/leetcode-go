package _03

// simple

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	root := new(ListNode)
	cpRoot := root
	for head != nil {
		if head.Val != val {
			cpRoot.Next = &ListNode{
				Val: head.Val,
			}
			cpRoot = cpRoot.Next
		}
		head = head.Next
	}

	return root.Next
}
