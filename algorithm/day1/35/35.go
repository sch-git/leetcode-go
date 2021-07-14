package _5

func searchInsert(nums []int, target int) int {
	if nums[0] > target {
		return 0
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}

	left := 0
	right := len(nums) - 1
	idx := 0
	for ; left <= right; {
		idx = (left + right) / 2
		if nums[idx] == target {
			return idx
		} else if nums[idx] > target {
			right = idx - 1
		} else if nums[idx] < target {
			left = idx + 1
		}
	}

	if target > nums[idx] {
		return idx + 1
	} else {
		return idx
	}
}
