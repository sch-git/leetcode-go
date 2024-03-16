package main

func main() {
	maxMoves([][]int{
		{187, 167, 209, 251, 152, 236, 263, 128, 135},
		{267, 249, 251, 285, 73, 204, 70, 207, 74},
		{189, 159, 235, 66, 84, 89, 153, 111, 189},
		{120, 81, 210, 7, 2, 231, 92, 128, 218},
		{193, 131, 244, 293, 284, 175, 226, 205, 245},
	})
}

func maxMoves(grid [][]int) int {
	m, n := len(grid)-1, len(grid[0])-1
	ans := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		ans[i] = make([]int, n+1)
	}
	for j := n - 1; j >= 0; j-- {
		for i := 0; i <= m; i++ {
			curMax := 0
			if (i > 0) && grid[i][j] < grid[i-1][j+1] {
				curMax = max(curMax, ans[i-1][j+1]+1)
			}
			if grid[i][j] < grid[i][j+1] {
				curMax = max(curMax, ans[i][j+1]+1)
			}
			if (i < m) && grid[i][j] < grid[i+1][j+1] {
				curMax = max(curMax, ans[i+1][j+1]+1)
			}
			ans[i][j] = curMax
		}
	}

	res := 0
	for i := 0; i <= m; i++ {
		res = max(res, ans[i][0])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
