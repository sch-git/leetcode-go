package main

func main() {

}

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	lenFx := make(map[int]int)
	chIdx := make(map[int32]int)
	chExist := make(map[rune]bool)
	for idx, ch := range s {
		lenFx[idx] = 1
		if chExist[ch] {
			beforeIdx := chIdx[ch-'a']
			chIdx[ch-'a'] = idx
			lenFx[idx] = minNum(lenFx[idx-1]+1, idx-beforeIdx)
		} else {
			chExist[ch] = true
			chIdx[ch-'a'] = idx
			if idx > 0 {
				lenFx[idx] = lenFx[idx-1] + 1
			}
		}
		maxLen = maxNum(maxLen, lenFx[idx])
	}
	return maxLen
}

func maxNum(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func minNum(i, j int) int {
	if i > j {
		return i
	}
	return j
}
