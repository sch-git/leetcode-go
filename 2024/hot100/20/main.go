package main

func main() {
	rotate([][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}})
}

func rotate(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	circle := 0

	if m == 1 {
		return
	}

	for circle < m/2 {
		i, j := circle, circle
		for ; j < n-1-circle; j++ {
			offset := j - circle
			pre := matrix[i][j]
			tmp := matrix[circle+offset][n-1-circle]
			matrix[circle+offset][n-1-circle] = pre

			pre = tmp
			tmp = matrix[m-1-circle][n-1-circle-offset]
			matrix[m-1-circle][n-1-circle-offset] = pre

			pre = tmp
			tmp = matrix[m-1-circle-offset][circle]
			matrix[m-1-circle-offset][circle] = pre

			matrix[i][j] = tmp

		}
		circle++
	}
	return
}
