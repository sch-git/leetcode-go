package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	newHead := &ListNode{}
	cur := newHead
	flag := true
	for flag {
		var minNode *ListNode
		var curMinNode *ListNode
		var curIdx int
		for idx, item := range lists {
			if item == nil {
				continue
			}

			if minNode == nil || minNode.Val > item.Val {
				minNode = &ListNode{Val: item.Val}
				curMinNode = item
				curIdx = idx
			}
		}
		if curMinNode == nil {
			break
		}
		cur.Next = minNode
		cur = cur.Next
		lists[curIdx] = lists[curIdx].Next
	}
	return newHead.Next
}
