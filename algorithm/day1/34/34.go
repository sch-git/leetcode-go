package _4

// medium

func SearchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for ; left <= right; {
		middle := (left + right) / 2
		if target > nums[middle] {
			left = middle + 1
		} else if target < nums[middle] {
			right = middle - 1
		} else {
			start, end := middle, middle
			for start >= 0 && nums[start] == target {
				start--
			}
			for end < len(nums) && nums[end] == target {
				end++
			}

			return []int{start + 1, end - 1}
		}
	}

	return []int{-1, -1}
}
