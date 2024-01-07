package main

func productExceptSelf(nums []int) []int {
	lfx := make([]int, len(nums), len(nums))
	rfx := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		lfx[i] = nums[i]
		if i > 0 {
			lfx[i] *= lfx[i-1]
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		rfx[i] = nums[i]
		if i < len(nums)-1 {
			rfx[i] *= rfx[i+1]
		}
	}

	ans := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = 1
		if i-1 >= 0 {
			ans[i] *= lfx[i-1]
		}
		if i+1 < len(nums) {
			ans[i] *= rfx[i+1]
		}
	}

	return ans
}
