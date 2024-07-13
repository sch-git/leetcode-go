package main

func main() {
	isMatch("aa", "a*")
}

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true
	for j := 2; j < len(p)+1; j += 2 {
		dp[0][j] = dp[0][j-2] && p[j-1] == '*'
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] != '*' {
				dp[i][j] = dp[i-1][j-1] && (p[j-1] == s[i-1] || p[j-1] == '.')
			} else {
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && p[j-2] == '.') || (dp[i-1][j] && s[i-1] == p[j-2])
			}
		}
	}
	return dp[len(s)][len(p)]
}
