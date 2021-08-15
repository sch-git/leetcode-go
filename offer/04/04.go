package _04

// middle

// 从右上角开始查找

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) < 1 {
		return false
	}
	for n, m := 0, len(matrix[0])-1; n < len(matrix) && m >= 0; {
		if target == matrix[n][m] {
			return true
		}
		if target > matrix[n][m] {
			n++
		} else {
			m--
		}
	}
	return false
}
