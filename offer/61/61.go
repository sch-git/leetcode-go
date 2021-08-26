package _61

import "sort"

// simple

func isStraight(nums []int) bool {
	sort.Ints(nums)
	count := 0

	m := make(map[int]bool)
	for _, num := range nums {
		if num == 0 {
			count++
			continue
		}
		if m[num] {
			return false
		}
		m[num] = true
	}

	if nums[len(nums)-1]-nums[count] < 5 {
		return true
	}
	return false
}
