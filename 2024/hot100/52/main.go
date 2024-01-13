package main

func main() {
	print(wordBreak("leetcode", []string{"leet", "code"}))
}

// 设s[i:j] 为字符串 s[i:j] 是否能有字典中的单词组成
// 如果最后的单词 word 存在字典中，则s[i:j] = map[s[i:j]] || map[i+len(word):j] && map[word]
func wordBreak(s string, wordDict []string) bool {
	dp := make([][]bool, 0)
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	for i := 0; i < len(s); i++ {
		dp = append(dp, make([]bool, len(s)))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := len(s); j > i; j-- {
			dp[i][j-1] = func() bool {
				if wordMap[s[i:j]] {
					return true
				}
				for _, word := range wordDict {
					if j-i > len(word) {
						if wordMap[s[i:i+len(word)]] && dp[i+len(word)][j-1] {
							return true
						}
					}
				}
				return false
			}()
		}
	}
	return dp[0][len(s)-1]
}
