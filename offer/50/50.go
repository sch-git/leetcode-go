package _50

import "strings"

// simple
// 可以考虑使用栈来实现

func firstUniqChar(s string) byte {
	if len(s) < 1 {
		return ' '
	}

	m := make(map[rune]bool)
	for idx, val := range s {
		if strings.Index(s[idx+1:], string(val)) == -1 && !m[val] {
			return byte(val)
		}
		m[val] = true
	}

	return ' '
}
