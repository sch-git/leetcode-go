package _14

func cuttingRope(n int) int {
	if n < 4 {
		return n - 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 1
	dp[3] = 2
	for i := 4; i <= n; i++ {
		dp[i] = maxNum(dp[i-1], maxNum(dp[i-2]*2, dp[i-3]*3))
	}
	return dp[n]
}

func maxNum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
