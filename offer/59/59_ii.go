package _59

// middle
// 队列和栈

type MaxQueue struct {
	queue []int
	stack []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		queue: make([]int, 0),
		stack: make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.queue) < 1 {
		return -1
	}

	return this.stack[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)
	for len(this.stack) > 0 && this.stack[len(this.stack)-1] < value {
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) < 1 {
		return -1
	}
	val := this.queue[0]
	this.queue = this.queue[1:]
	if val == this.stack[0] {
		this.stack = this.stack[1:]
	}
	return val
}
