package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}

	q, s, ans := head, head, &ListNode{Next: head}
	curHead := ans
	step := 1
	sstep := 1
	for ; q != nil; step++ {
		q = q.Next
		if step-sstep > k {
			cur := s.Next
			next := curHead.Next

			s.Next = cur.Next
			cur.Next = next
			curHead.Next = cur

			sstep++
			if sstep%k == 0 {
				curHead = s
				s = s.Next
				sstep++ // sstep 初始默认认为反转了 1，所以每次进入新的一段需要反转的链表需要+1
			}
		}
	}
	for sstep%k != 0 { // 需要反转的数量一定是 nk，不满足的需要再继续反转
		cur := s.Next
		next := curHead.Next

		s.Next = cur.Next
		cur.Next = next
		curHead.Next = cur

		sstep++
	}
	return ans.Next
}
