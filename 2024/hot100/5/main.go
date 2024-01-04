package main

func maxArea(height []int) int {
	maxarrea := 0
	i, j := 0, len(height)-1
	for i <= j {
		width := j - i
		curarea := width * min(height[i], height[j])
		maxarrea = max(maxarrea, curarea)

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxarrea
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
