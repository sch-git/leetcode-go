package _1034

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	dyeing(grid, row, col, grid[row][col], color)
	return grid
}

func dyeing(grid [][]int, row, col, source, color int) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) {
		return
	}

	if grid[row][col] == source {
		grid[row][col] = color
		dyeing(grid, row-1, col, source, color)
		dyeing(grid, row+1, col, source, color)
		dyeing(grid, row, col-1, source, color)
		dyeing(grid, row, col+1, source, color)
	}
}
