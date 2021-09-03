package _57

// simple
// 滑动窗口

func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	for l, r := 1, 1; l <= target/2+1 && r <= target/2+1; {
		sum := (l + r) * (r - l + 1) / 2
		if sum == target {
			temp := make([]int, 0)
			for i := l; i <= r; i++ {
				temp = append(temp, i)
			}
			res = append(res, temp)
			l++
		} else if sum < target {
			r++
		} else {
			l++
		}
	}
	return res
}
