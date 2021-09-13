package _51

func reversePairs(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] {
				dp[i]++
			}
		}
		if i > 0 {
			dp[i] += dp[i-1]
		}
	}
	return dp[len(nums)-1]
}
