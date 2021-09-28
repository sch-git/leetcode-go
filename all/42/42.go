package _42

// hard
// 单调栈

func trap(height []int) (ans int) {
	if len(height) < 1 {
		return 0
	}

	stack := make([]int, 0)
	for idx, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			topIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			leftIdx := stack[len(stack)-1]
			curWidth := idx - leftIdx - 1
			curHeight := min(h, height[leftIdx]) - height[topIdx]
			ans += curWidth * curHeight
		}
		stack = append(stack, idx)
	}

	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
