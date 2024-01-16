package main

func main() {
	longestPalindrome("cbbd")
}

func longestPalindrome(s string) string {
	mp := make(map[string]bool)
	dp := make([][]int, len(s))
	ans := ""
	for i, ch := range s {
		mp[string(ch)] = true
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
		ans = string(ch)
	}
	for j := 1; j < len(s); j++ {
		for i := 0; i < len(s)-j; i++ {
			curj := i + j
			if s[i] == s[curj] && (((i+1 <= curj-1) && mp[s[i+1:curj-1+1]]) || curj == i+1) {
				cur := s[i : curj+1]
				dp[i][curj] = dp[i+1][curj-1] + 2
				if curj-i >= len(ans) {
					ans = cur
				}
				mp[cur] = true
			} else {
				dp[i][curj] = max(dp[i][curj-1], dp[i+1][curj])
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a

	}
	return b

}
