package _59

// middle
// 双端队列

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < 1 {
		return nil
	}

	res := make([]int, 0)
	windows := make([]int, 0, k)
	queue := make([]int, 0)
	for _, num := range nums {
		windows = append(windows, num)
		if len(queue) < 1 {
			queue = append(queue, num)
		} else {
			for len(queue) > 0 && queue[len(queue)-1] < num { // 保证队列递减
				queue = queue[:len(queue)-1]
			}
			queue = append(queue, num)
		}

		if len(windows) == k { // 判断滑动窗口大小缩容
			res = append(res, queue[0])
			if queue[0] == windows[0] { // 如果队首是出窗口的数，队首出队列
				queue = queue[1:]
			}
			windows = windows[1:] // 窗口缩小
		}
	}
	return res
}
