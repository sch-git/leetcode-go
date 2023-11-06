package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	windows := make([]rune, 0)
	kmap := make(map[rune]bool)
	for _, ch := range s {
		windows = append(windows, ch)
		if kmap[ch] {
			for windows[0] != ch {
				kmap[windows[0]] = false
				windows = windows[1:]
			}
			windows = windows[1:]
		}

		kmap[ch] = true
		maxLen = maxNum(maxLen, len(windows))
	}
	return maxLen
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
