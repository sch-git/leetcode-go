package _57

import "strings"

// simple

func reverseWords(s string) string {
	res := make([]string, 0)
	for _, str := range strings.Split(s, " ") {
		chs := []byte(str)	// 字符串 str 不可变
		reverseWord(chs)
		res = append(res, string(chs))
	}

	return strings.Join(res," ")
}

func reverseWord(s []byte) {
	for left, right := 0, len(s)-1; left < right; {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
