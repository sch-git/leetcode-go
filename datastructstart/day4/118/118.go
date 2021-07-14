package _18

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for r := 0; r < numRows; r++ {
		res[r] = make([]int, r+1)
		for c := 0; c <= r; c++ {
			if c == 0 || c == r {
				res[r][c] = 1
			} else {
				res[r][c] = res[r-1][c] + res[r-1][c-1]
			}
		}
	}

	return res
}
