package _63

import "math"

// middle
// 动态规划

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	min := math.MaxInt32
	res := 0
	for _, i := range prices {
		if min > i {
			min = i
		}

		if i-min > res {
			res = i - min
		}
	}

	return res
}
