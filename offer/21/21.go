package _21

// simple

func exchange(nums []int) []int {
	evenIdx := len(nums) - 1
	for idx, num := range nums {
		for num%2 == 0 && evenIdx > idx {
			if nums[evenIdx]%2 == 1 {
				nums[evenIdx], nums[idx] = nums[idx], nums[evenIdx]
				evenIdx--
				break
			}
			evenIdx--
		}
	}

	return nums
}
