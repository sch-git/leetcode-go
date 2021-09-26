package _017

import "math"

// hard
// 滑动窗口

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	minLen := math.MaxInt64
	res := ""
	windows := make(map[rune]int)
	for _, ch := range t {
		windows[ch]++
	}

	for left, right := 0, 0; left <= right && right < len(s); right++ {
		ch := s[right]
		windows[rune(ch)]--
		windowsLen := right - left + 1
		if windowsLen >= len(t) {
			// 缩小窗口
			for isEmpty(windows) {
				if windowsLen < minLen {
					// 更新最小值
					minLen = windowsLen
					res = s[left : left+minLen]
				}
				windows[rune(s[left])]++
				left++
				windowsLen--
			}
		}
	}

	return res
}

// 判断是否所有字符存在
func isEmpty(m map[rune]int) bool {
	for _, v := range m {
		if v > 0 {
			return false
		}
	}

	return true
}
