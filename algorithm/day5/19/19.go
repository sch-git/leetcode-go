package _9

// medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	root := &ListNode{
		Next: head,
	}
	first, second := root, root
	key := 0
	for first.Next != nil {
		if key >= n {
			second = second.Next
		}
		first = first.Next
		key++
	}

	if second.Next!=nil{
		second.Next = second.Next.Next
	}
	return root.Next
}
