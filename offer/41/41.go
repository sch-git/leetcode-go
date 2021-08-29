package _41

import "sort"

// hard
// TODO 实现 heap 接口

type MedianFinder struct {
	min *minHeap
	max *maxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		min: &minHeap{
			slice: make(sort.IntSlice, 0),
		},
		max: &maxHeap{
			IntSlice: make(sort.IntSlice, 0),
		},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.min.slice) != len(this.max.IntSlice) {
		this.min.Push(num)
		this.max.Push(this.min.Pop())
	} else {
		this.max.Push(num)
		this.min.Push(this.max.Pop())
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.min.slice.Len() != this.max.IntSlice.Len() {
		return float64(this.min.Peek())
	} else {
		return (float64(this.min.Peek()) + float64(this.max.Peek())) / 2
	}
}

// sort 默认升序，用于构建小顶堆
type minHeap struct {
	slice sort.IntSlice
}

func (h *minHeap) Push(num int) {
	h.slice = append(h.slice, num)
}

func (h *minHeap) Pop() int {
	length := h.slice.Len()
	num := h.slice[length-1]
	h.slice = h.slice[:length-1]
	return num
}

func (h *minHeap) Peek() int {
	return h.slice[len(h.slice)-1]
}

// 构建大顶堆，重写 func Less
type maxHeap struct {
	sort.IntSlice
}

func (h *maxHeap) Push(num int) {
	h.IntSlice = append(h.IntSlice, num)
}

func (h *maxHeap) Pop() int {
	length := h.IntSlice.Len()
	num := h.IntSlice[length-1]
	h.IntSlice = h.IntSlice[:length-1]
	return num
}

func (h *maxHeap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *maxHeap) Peek() int {
	return h.IntSlice[len(h.IntSlice)-1]
}
