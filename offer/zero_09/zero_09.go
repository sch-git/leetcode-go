package zero_09

type CQueue struct {
	stackHead []int
	stackTail []int
}

func Constructor() CQueue {
	return CQueue{
		stackHead: make([]int, 0),
		stackTail: make([]int, 0),
	}
}

func (c *CQueue) AppendTail(value int) {
	c.stackHead = append(c.stackHead, value)
}

func (c *CQueue) DeleteHead() int {
	if len(c.stackTail) > 0 {
		res := c.stackTail[len(c.stackTail)-1]
		c.stackTail = c.stackTail[:len(c.stackTail)-1]
		return res
	}

	for i := len(c.stackHead) - 1; i >= 0; i-- {
		c.stackTail = append(c.stackTail, c.stackHead[i])
	}
	if len(c.stackTail)<1{
		return -1
	}
	res := c.stackTail[len(c.stackTail)-1]
	c.stackTail = c.stackTail[:len(c.stackTail)-1]
	c.stackHead = make([]int, 0)
	return res
}
