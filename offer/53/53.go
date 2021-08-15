package _53

// simple

func search(nums []int, target int) int {
	times := 0
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			times += 1
			for i := mid - 1; i >= 0 && nums[i] == target; i-- {
				times++
			}
			for i := mid + 1; i < len(nums) && nums[i] == target; i++ {
				times++
			}
			break
		}
	}
	return times
}
