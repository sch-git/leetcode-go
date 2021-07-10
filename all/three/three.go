package three

import "math"

func LengthOfLongestSubstring(s string) int {
	m := make(map[string]int)
	maxLen := 0
	length := 0
	start := 1
	for index, val := range s {
		if m[string(val)] >= start {
			start = m[string(val)]
			length = index - start + 1
		} else {
			length++
		}
		m[string(val)] = index+1
		maxLen = int(math.Max(float64(maxLen), float64(length)))
	}
	maxLen = int(math.Max(float64(maxLen), float64(length)))
	return maxLen
}
