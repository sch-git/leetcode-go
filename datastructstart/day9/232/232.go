package _32

type MyQueue struct {
	inStack, outStack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.outStack) > 0 {
		res := this.outStack[len(this.outStack)-1]
		this.outStack = this.outStack[:len(this.outStack)-1]
		return res
	}

	for idx := len(this.inStack) - 1; idx >= 0; idx-- {
		this.outStack = append(this.outStack, this.inStack[idx])
	}
	this.inStack = make([]int, 0)
	res := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.outStack) > 0 {
		return this.outStack[len(this.outStack)-1]
	}

	for idx := len(this.inStack) - 1; idx >= 0; idx-- {
		this.outStack = append(this.outStack, this.inStack[idx])
	}
	this.inStack = make([]int, 0)

	return this.outStack[len(this.outStack)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}
