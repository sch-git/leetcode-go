package _0208

// middle
type ListNode struct {
	Val  int
	Next *ListNode
}

// 题解：设 环外部分长度为 a，slow 进入环后走了 b 距离与 fast 相遇，则 fast 走过的距离为 a+n（b+c）+b = a+（n+1）b+nc = 2（a+b）
// 当 slow 和 fast 相遇时，从 head 位置开始与 slow 同时向后移动一个位置，相遇时即是入环点
// 环路检测
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			ptr := head
			for slow != ptr {
				slow = slow.Next
				ptr = ptr.Next
			}
			return slow
		}
	}

	return nil
}
