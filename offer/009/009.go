package _009

// middle
// 滑动窗口

func numSubarrayProductLessThanK(nums []int, k int) int {
	if len(nums) < 1 {
		return 0
	}
	res, product := 0, 1
	windows := make([]int, 0)
	for _, num := range nums {
		product *= num
		windows = append(windows, num)
		for product >= k && len(windows) > 0 {
			product /= windows[0]
			windows = windows[1:]
		}
		if product < k && len(windows) > 0 {
			res+=len(windows)
		}
	}

	return res
}
