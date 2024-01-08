package main

func main() {
	spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})
}

func spiralOrder(matrix [][]int) []int {
	ans := make([]int, 0)
	circle := 0
	sum := len(matrix) * len(matrix[0])
	for len(ans) < sum {
		i, j := circle, circle
		for ; len(ans) < sum && j < len(matrix[0])-circle; j++ {
			ans = append(ans, matrix[i][j])
		}

		i, j = circle+1, len(matrix[0])-circle-1
		for ; len(ans) < sum && i < len(matrix)-circle; i++ {
			ans = append(ans, matrix[i][j])
		}

		i, j = len(matrix)-circle-1, len(matrix[0])-circle-2
		for ; len(ans) < sum && j >= circle; j-- {
			ans = append(ans, matrix[i][j])
		}

		i, j = len(matrix)-circle-2, circle
		for ; len(ans) < sum && i > circle; i-- {
			ans = append(ans, matrix[i][j])
		}
		circle++
	}
	return ans
}
