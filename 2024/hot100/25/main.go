package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	q, s := head, head
	for q != nil && q.Next != nil {
		s = s.Next
		q = q.Next.Next
		if s == q {
			return true
		}
	}
	return false
}
