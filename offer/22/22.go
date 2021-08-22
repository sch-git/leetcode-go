package _22

// simple
// 双指针

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if k < 1 {
		return head
	}

	pre := &ListNode{
		Next: head,
	}

	for head != nil && head.Next != nil {
		if k <= 1 {
			head = head.Next
			pre = pre.Next
		} else {
			head = head.Next
			k--
		}
	}

	return pre.Next
}
