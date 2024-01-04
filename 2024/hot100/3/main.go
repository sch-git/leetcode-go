package main

func main() {

}

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLen := 0
	for _, num := range nums {
		curLen := 1
		if numSet[num-1] {
			continue
		}

		tempNum := num
		for numSet[tempNum+1] {
			curLen++
			tempNum++
		}
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}
