package _016

// middle
// 滑动窗口

func lengthOfLongestSubstring(s string) int {
	windows, max := 0, 0
	m := make(map[rune]int)
	exist := make(map[rune]bool)
	for idx, val := range s {
		windows++
		if exist[val]{
			if idx - m[val] <= windows{
				windows = idx-m[val]
			}
		}
		max = maxNum(windows, max)
		exist[val] = true
		m[val] = idx
	}

	return maxNum(windows, max)
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}

	return b
}

