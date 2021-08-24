package _13

import "strconv"

// middle
// 递归回溯

func MovingCount(m int, n int, k int) int {
	board := make([][]int, m)
	for i := 0; i < m; i++ {
		board[i] = make([]int, n)
	}
	return run(board, m, n, 0, 0, k)
}

func run(board [][]int, m, n, r, c, k int) int {
	if r < 0 || c < 0 || r >= m || c >= n || board[r][c] > 0 {
		return 0
	}

	if count(strconv.Itoa(r)+strconv.Itoa(c)) > k {
		return 0
	}

	board[r][c] = 1
	return run(board, m, n, r-1, c, k) + run(board, m, n, r+1, c, k) + run(board, m, n, r, c-1, k) + run(board, m, n, r, c+1, k) + 1
}

func count(str string) int {
	sum := 0
	for _, val := range str {
		num, _ := strconv.Atoi(string(val))
		sum += num
	}
	return sum
}
