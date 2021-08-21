package _46

import "strconv"

// middle
// 动态规划
// dp[n] = dp[n]-1 + dp[n-2](if dp[n-2:] <= 26)
func translateNum(num int) int {
	s := strconv.Itoa(num)
	if len(s) == 1 {
		return 1
	}

	nums := []int{1, 1}
	for i := 1; i < len(s); i++ {
		pre := s[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			nums[0], nums[1] = nums[1], nums[0]+nums[1]
		} else {
			nums[0], nums[1] = nums[1], nums[1]
		}
	}
	return nums[1]
}
