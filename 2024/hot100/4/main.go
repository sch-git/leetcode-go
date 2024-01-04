package main

func moveZeroes(nums []int) {
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[idx] = nums[i]
			idx++
		}
	}
	for i := idx; i < len(nums); i++ {
		nums[i] = 0
	}
}
