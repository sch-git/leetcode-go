package mid

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	minHeight := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	maxNum := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	for left < right {
		l := right - left
		h := minHeight(height[left], height[right])
		max = maxNum(max, l*h)

		if height[left] < height[right] {
			left++
		} else if height[right] < height[left] {
			right--
		} else {
			left++
			right--
		}
	}

	return max
}
