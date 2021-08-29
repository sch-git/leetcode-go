package _11

// middle
// 双指针

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for left < right {
		max = maxNum((right-left)*min(height[left], height[right]), max)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}

func maxNum(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
