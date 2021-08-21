package _10

// simple
// 动态规划

func Fib(n int) int {
	nums := []int{0, 1}
	if n < 2 {
		return nums[n]
	}

	for i := 2; i <= n; i++ {
		temp := nums[1]
		nums[1] = nums[0] + nums[1]
		nums[0] = temp
		// nums[0],nums[1] = nums[1],(nums[0] + nums[1])%(1e9+7)
	}

	return nums[1] % (1e9 + 7)
}
