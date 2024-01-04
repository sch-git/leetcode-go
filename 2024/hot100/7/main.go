package main

func main() {
	trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
}

func trap(height []int) int {
	leftMax := make([]int, 0)
	rightMax := make([]int, len(height), len(height))
	area := 0
	curMax := 0
	for _, h := range height {
		curMax = max(curMax, h)
		leftMax = append(leftMax, curMax-h)
	}
	curMax = 0
	for j := len(height) - 1; j > 0; j-- {
		curMax = max(curMax, height[j])
		rightMax[j] = curMax - height[j]
	}

	for idx, _ := range height {
		area += min(leftMax[idx], rightMax[idx])
	}

	return area
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
