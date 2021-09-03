package _31

// middle
// stack

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 && len(popped) == 0 {
		return true
	}

	stack := make([]int, 0)
	for len(stack) > 0 || len(pushed) > 0 {
		stack = append(stack, pushed[0])
		for len(stack) > 0 && len(popped) > 0 && stack[len(stack)-1] == popped[0] {
			stack = stack[:len(stack)-1]
			popped = popped[1:]
		}
		pushed = pushed[1:]
		if len(pushed) <= 0 {
			break
		}
	}

	return len(stack) <= 0
}
