package _1109

// middle
// 滑动窗口

func corpFlightBookings(bookings [][]int, n int) []int {
	ans := make([]int, 0)
	for len(ans) < n {
		ans = append(ans, 0)
	}

	for _, booking := range bookings {

		for start := booking[0] - 1; start < booking[1]; start++ {
			ans[start] += booking[2]
		}
	}

	return ans
}
