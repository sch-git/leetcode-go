package main

func main() {
	reorderList(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}})
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head.Next == nil || head.Next.Next == nil {
		return
	}

	// 先生成逆序链表
	newHead := &ListNode{}
	cur := head
	nodeMap := make(map[*ListNode]*ListNode)
	listLen := 0
	for cur != nil {
		curNewHeadNext := newHead.Next
		newHead.Next = &ListNode{Val: cur.Val, Next: curNewHeadNext}
		nodeMap[newHead.Next] = cur
		cur = cur.Next
		listLen++
	}

	cur, newCur := head, newHead.Next
	newMap := make(map[*ListNode]*ListNode)
	for cur != nil {
		newMap[cur] = nodeMap[newCur]
		cur = cur.Next
		newCur = newCur.Next
	}

	listLen = listLen / 2
	cur = head
	for ; listLen > 0; listLen-- {
		next := cur.Next
		cur.Next = newMap[cur]
		cur.Next.Next = next
		cur = next
	}

	cur.Next = nil
	return
}
