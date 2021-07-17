package _87

import "math"

func firstUniqChar(s string) int {
	existTime := make(map[rune]int)
	idx := make(map[rune]int)
	for index, str := range s {
		existTime[str]++
		idx[str] = index
	}

	minIdx := math.MaxInt32
	for key, times := range existTime {
		if times == 1 && idx[key] < minIdx {
			minIdx = idx[key]
		}
	}

	if minIdx == math.MaxInt32{
		return -1
	}

	return minIdx
}
