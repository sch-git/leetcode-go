package _62

// middle
// 动态规划，初始化 行、列

func uniquePaths(m int, n int) int {
	table := make([][]int, m)
	for i := 0; i < m; i++ {
		table[i] = make([]int, n)
		table[i][0] = 1
	}

	for i := 0; i < n; i++ {
		table[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			table[i][j] = table[i-1][j] + table[i][j-1]
		}
	}

	return table[m-1][n-1]
}
