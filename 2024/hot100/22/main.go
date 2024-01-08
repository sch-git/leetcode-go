package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ma, mb := 0, 0
	cur := headA
	for cur != nil {
		ma++
		cur = cur.Next
	}

	cur = headB
	for cur != nil {
		mb++
		cur = cur.Next
	}

	if ma > mb {
		for i := 0; i < ma-mb; i++ {
			headA = headA.Next
		}
	} else if mb > ma {
		for i := 0; i < mb-ma; i++ {
			headB = headB.Next
		}
	}
	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}
