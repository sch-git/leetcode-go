package _04

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for ; left <= right; {
		idx := (left + right) / 2
		if nums[idx] == target {
			return idx
		} else if nums[idx] > target {
			right = idx - 1
		} else if nums[idx] < target {
			left = idx + 1
		}
	}

	return -1
}
