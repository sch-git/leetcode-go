package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}}
	fmt.Println(closedIsland(grid))
}
func closedIsland(grid [][]int) int {
	var ans int
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == 1 || grid[x][y] == -1 {
				continue
			}
			if checkIslandClosed(grid, x, y) {
				ans++
			}
		}
	}
	return ans
}

func checkIslandClosed(grid [][]int, curx, cury int) (closed bool) {
	if curx < 0 || cury < 0 || curx > len(grid)-1 || cury > len(grid[0])-1 {
		return false
	}

	if grid[curx][cury] == 1 || grid[curx][cury] == -1 {
		return true
	}

	if grid[curx][cury] == 0 && (curx == 0 || cury == 0) {
		return false
	}

	grid[curx][cury] = -1
	up := checkIslandClosed(grid, curx-1, cury)
	down := checkIslandClosed(grid, curx+1, cury)
	left := checkIslandClosed(grid, curx, cury-1)
	right := checkIslandClosed(grid, curx, cury+1)

	return up && down && left && right
}
