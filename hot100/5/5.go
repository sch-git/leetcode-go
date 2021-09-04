package _5

// middle
// 最长回文子串

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	maxLen := 0
	res := s[0:0]

	for l := 2; l <= len(s); l++ { // 字符串长度
		for i := 0; i < len(s); i++ { // 左边界
			r := l + i - 1 // 右边界
			if r >= len(s) {
				continue
			}

			if s[i] == s[r] {
				if r-i < 3 {
					dp[i][r] = true
				} else {
					dp[i][r] = dp[i+1][r-1]
				}
				if dp[i][r] && r-i > maxLen {
					maxLen = r - i
					res = s[i : r+1]
				}
			}
		}
	}
	return res
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
