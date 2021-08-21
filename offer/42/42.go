package _42

import "math"

// simple
// 动态规划

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	max := math.MinInt32
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i]+nums[i-1] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
