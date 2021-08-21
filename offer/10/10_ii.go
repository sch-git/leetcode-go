package _10

// simple
// 动态规划

func numWays(n int) int {
	nums := []int{1, 2}
	if n < 3 {
		if n == 0 {
			return 1
		}
		return nums[n-1]
	}

	for i := 3; i <= n; i++ {
		nums[0], nums[1] = nums[1], (nums[0]+nums[1])%(1e9+7)
	}
	return nums[1] % (1e9 + 7)
}
