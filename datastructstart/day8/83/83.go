package _3

// simple

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	root := head
	if head == nil {
		return root
	}

	m := make(map[int]bool)
	m[head.Val] = true

	for head.Next != nil {
		if m[head.Next.Val] {
			head.Next = head.Next.Next
			continue
		}
		m[head.Next.Val] = true
		head = head.Next
	}

	return root
}
