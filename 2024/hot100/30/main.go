package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := &ListNode{Next: head}
	cur := pre.Next
	ans := pre
	for cur != nil && cur.Next != nil {
		next := cur.Next
		pre.Next = next
		cur.Next = next.Next
		next.Next = cur

		pre = cur
		cur = cur.Next
	}
	return ans.Next
}
