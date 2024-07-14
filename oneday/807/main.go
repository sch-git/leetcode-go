package main

func main() {

}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	row, col := make([]int, 0), make([]int, 0)

	for i := 0; i < len(grid); i++ {

		rowMax := 0

		for j := 0; j < len(grid[0]); j++ {
			// 找到每行最大值
			rowMax = maxNum(rowMax, grid[i][j])
			// 找到每列最大值
			if len(col) < len(grid[0]) {
				col = append(col, 0)
			}
			col[j] = maxNum(col[j], grid[i][j])
		}

		row = append(row, rowMax)

	}

	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			ans += minNum(row[i], col[j]) - grid[i][j]
		}
	}

	return ans
}

func maxNum(a, b int) int {
	if a > b {
		return a

	}
	return b
}

func minNum(a, b int) int {
	if a < b {
		return a
	}
	return b
}
