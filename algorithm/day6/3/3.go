package _3

import "math"

func LengthOfLongestSubstring(s string) int {
	left, right := 1, 1
	m := make(map[int32]int)
	maxLen := 0
	for idx, val := range s {
		if m[val] == 0 {
			m[val] = idx + 1
		} else {
			if m[val] >= left {
				left = m[val] + 1
			}
			m[val] = idx + 1

		}
		right++
		maxLen = int(math.Max(float64(maxLen), float64(right-left)))
	}

	return maxLen
}
