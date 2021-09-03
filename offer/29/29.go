package _29

// simple

func spiralOrder(matrix [][]int) []int {
	if len(matrix) < 1 {
		return nil
	}
	res := make([]int, 0)
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	for {
		for j := l; j <= r; j++ { // left -> right
			res = append(res, matrix[t][j])
		}
		t++
		if t > b {
			break
		}

		for i := t; i <= b; i++ { // top -> bottom
			res = append(res, matrix[i][r])
		}
		r--
		if l > r {
			break
		}

		for i := r; i >= l; i-- { // right -> left
			res = append(res, matrix[b][i])
		}
		b--
		if t > b {
			break
		}

		for i := b; i >= t; i-- {
			res = append(res, matrix[i][l])
		}
		if l > r {
			break
		}
	}
	return res
}
