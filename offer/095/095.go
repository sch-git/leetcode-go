package _095

// middle
// 动态规划
// 最长公共子序列

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	dp[0] = make([]int, len(text2)+1)
	for i := 1; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
		for j := 1; j <= len(text2); j++ {
			pre := max(dp[i-1][j], dp[i][j-1])
			ch1, ch2 := text1[i-1], text2[j-1]

			if ch1 == ch2 {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = pre
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
