package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	quick := head.Next
	slow := head
	headpre := &ListNode{Next: head}
	cur := headpre
	change := false
	for quick != nil {
		for quick != nil && quick.Val == slow.Val {
			quick = quick.Next
			slow = slow.Next
			change = true
		}
		if change {
			cur.Next = quick
			if quick != nil {
				quick = quick.Next
			}
			slow = slow.Next
			change = false
		} else {
			cur = cur.Next
			if quick != nil {
				quick = quick.Next
			}
			slow = slow.Next
		}

	}
	return headpre.Next
}
