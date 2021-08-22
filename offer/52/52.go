package _52

// simple
// 双指针，遍历一个后再遍历另外一个，使两个链表遍历总数一致

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	rootA := headA
	rootB := headB
	for (headA != nil || headB != nil) || (headA != headB) {
		if headA == nil {
			headA = rootB
		} else if headB == nil {
			headB = rootA
		} else {
			headA = headA.Next
			headB = headB.Next
		}
	}

	return headA
}
