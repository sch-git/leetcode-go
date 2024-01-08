package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{Next: head}
	del := newHead
	pre := newHead
	o := 0
	for ; del != nil && del.Next != nil; o++ {
		del = del.Next
		if o >= n {
			pre = pre.Next
		}
	}
	if pre != nil && pre.Next != nil {
		pre.Next = pre.Next.Next
		return newHead.Next
	}
	return nil
}
