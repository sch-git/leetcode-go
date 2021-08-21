package _47

// middle
// 动态规划
// dp[m][n] += max(dp[m-1][n],dp[m][n-1])

func maxValue(grid [][]int) int {
	m := len(grid) - 1
	n := len(grid[0]) - 1

	for rdx, nums := range grid {
		for cdx, _ := range nums {
			if rdx == 0 && cdx == 0{
				continue
			}
			if rdx == 0 && cdx != 0 {
				grid[rdx][cdx] += grid[rdx][cdx-1]
			} else if cdx == 0 && rdx != 0 {
				grid[rdx][cdx] += grid[rdx-1][cdx]
			} else {
				grid[rdx][cdx] += maxVal(grid[rdx-1][cdx], grid[rdx][cdx-1])
			}
		}
	}

	return grid[m][n]
}

func maxVal(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
