package _6

import "math"

func isValidSudoku(board [][]byte) bool {
	m := make(map[byte][][]int)
	for row, datas := range board {
		for col, data := range datas {
			if string(data) == "." {
				continue
			}

			if m[data] != nil {
				if checkRowAndColAndSmallNine(m[data], row, col) {
					m[data] = append(m[data], []int{row, col})
				} else {
					return false
				}
			} else {
				m[data] = make([][]int, 0)
				m[data] = append(m[data], []int{row, col})
			}
		}
	}

	return true
}

func checkRowAndColAndSmallNine(board [][]int, row, col int) bool {
	for _, indexs := range board {
		// 判断在同一行或同一列
		if indexs[0] == row || col == indexs[1] {
			return false
		}

		// 判断在同一个小九宫格
		if abs(indexs[0], row) <= 2 && abs(indexs[1], col) <= 2 {
			for i := 2; i <= 8; i += 3 {
				for j := 2; j <= 8; j += 3 {
					if (indexs[0] > i-3 && indexs[0] <= i) && (row > i-3 && row <= i) && (indexs[1] > j-3 && indexs[1] <= j) && (col > j-3 && col <= j) {
						return false
					}
				}
			}
		}
	}
	return true
}

func abs(row, col int) int {
	return int(math.Abs(float64(row - col)))
}
