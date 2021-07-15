package _3

func setZeroes(matrix [][]int) {
	cols := make(map[int]bool)
	rows := make(map[int]bool)
	for row, vals := range matrix {
		for col, val := range vals {
			if val == 0 {
				rows[row] = true
				cols[col] = true
			}
		}
	}

	for row, _ := range rows {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[row][i] = 0
		}
	}

	for col, _ := range cols {
		for i := 0; i < len(matrix); i++ {
			matrix[i][col] = 0
		}
	}

	return
}
