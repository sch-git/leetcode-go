package _018

import "strings"

// simple
// 双指针

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for left, right := 0, len(s)-1; left <= right; {
		for left < right && !isNum(s[left]) && !isChar(s[left]) {
			left++
		}
		for left < right && !isNum(s[right]) && !isChar(s[right]) {
			right--
		}

		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}

	return true
}

func isNum(num uint8) bool {
	if num >= '0' && num <= '9' {
		return true
	}

	return false
}

func isChar(num uint8) bool {
	if num >= 'a' && num <= 'z' {
		return true
	}

	return false
}
