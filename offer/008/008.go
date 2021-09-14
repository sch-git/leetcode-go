package _008

import "math"

// middle
// 滑动窗口

func minSubArrayLen(target int, nums []int) int {
	windows := make([]int, 0)
	sum, res := 0, math.MaxInt64
	if len(nums) < 1 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		windows = append(windows, nums[i])
		for sum >= target {
			res = min(res, len(windows))
			sum -= nums[i-len(windows)]
			windows = windows[1:]
		}
	}

	if res == math.MaxInt64 {
		res = 0
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
