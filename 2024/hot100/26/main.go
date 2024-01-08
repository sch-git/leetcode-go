package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	f, s := head, head
	if head == nil || head.Next == nil {
		return nil
	}
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			break
		}
	}

	if f == nil || f.Next == nil {
		return nil
	}
	f = head
	for f != nil {
		f = f.Next
		s = s.Next
		if f == s {
			return f
		}
	}
	return nil
}
