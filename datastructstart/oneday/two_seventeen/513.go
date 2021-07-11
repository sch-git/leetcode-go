package two_seventeen

import "math"

func maxSubArray(nums []int) int {
	maxNum := math.MinInt32
	sum := 0
	for _, num := range nums {
		sum += num
		if num > sum {
			sum = num
		}

		if sum > maxNum {
			maxNum = sum
		}
	}
	return maxNum
}
