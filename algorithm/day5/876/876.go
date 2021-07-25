package _76

// simple

type ListNode struct {
	Val   int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	sum := 0
	nodes := make([]*ListNode, 0)
	for head != nil {
		nodes = append(nodes, head)
		sum++
		head = head.Next
	}

	return nodes[sum/2]
}

// 快慢指针