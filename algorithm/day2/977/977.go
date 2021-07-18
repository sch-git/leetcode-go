package _77

import "sort"

// simple

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	for idx, num := range nums {
		res[idx] = num * num
	}

	sort.Ints(res)
	return res
}
