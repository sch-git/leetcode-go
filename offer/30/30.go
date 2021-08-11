package _30

import "math"

// simple

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if this.topMin() >= x {
		this.minStack = append(this.minStack, x)
	}
}

func (this *MinStack) Pop() {
	length := len(this.stack)
	res := this.stack[length-1]
	this.stack = this.stack[:length-1]

	if res == this.topMin() {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.topMin()
}

func (this *MinStack) topMin() int {
	length := len(this.minStack)
	if length > 0 {
		return this.minStack[length-1]
	}

	return math.MaxInt64
}
