package main

import (
	"math"
)

func main() {
	coinChange([]int{1, 2, 5}, 11)
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	res := []int{-1}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 1; i <= amount; i++ {
		var curmin = math.MaxInt32
		var tmpmin = -1
		for _, v := range coins {
			if i == v {
				tmpmin = 1
				break
			}
			if i-v >= 0 && res[i-v] != -1 {
				curmin = min(curmin, res[i-v]+1)
				tmpmin = curmin
			}
		}
		res = append(res, tmpmin)
	}

	return res[len(res)-1]
}
