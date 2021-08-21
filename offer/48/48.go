package _48

// middle
// 滑动窗口

func LengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}

	dis := make(map[rune]int)
	max, minIdx := 0, 1
	sub := ""
	for idx, val := range s {
		if dis[val] > 0 || (dis[val] == 0 && idx != 0 && val == int32(s[0])) {
			if minIdx < dis[val]+1 {
				minIdx = dis[val] + 1
			}
			sub = s[minIdx : idx+1]
		} else {
			sub += string(val)
		}
		dis[val] = idx
		if len(sub) > max {
			max = len(sub)
		}
	}
	return max
}
