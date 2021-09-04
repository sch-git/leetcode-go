package _23

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := new(minHeap)
	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}

	root := &ListNode{}
	head := root
	for h.Len() > 0 {
		tmp := heap.Pop(h).(*ListNode)
		head.Next = &ListNode{
			Val: tmp.Val,
		}
		head = head.Next

		tmp = tmp.Next
		if tmp != nil {
			heap.Push(h, tmp)
		}
	}

	return root.Next
}

type minHeap []*ListNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*ListNode))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
