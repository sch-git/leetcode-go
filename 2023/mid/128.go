package main

func longestConsecutive(nums []int) int {
	maxLen := 0
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	for num, _ := range numSet {
		if !numSet[num-1] {
			curNum := num
			curLen := 1
			for numSet[curNum+1] {
				curLen++
				curNum++
			}
			maxLen = func() int {
				if curLen > maxLen {
					return curLen
				}
				return maxLen
			}()
		}
	}
	return maxLen
}
