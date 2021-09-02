package _39

import "sort"

func majorityElement(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	sort.Ints(nums)
	return nums[len(nums)/2]
}
