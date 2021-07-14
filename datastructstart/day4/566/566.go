package _66

func matrixReshape(mat [][]int, r int, c int) [][]int {
	// 判断原矩阵的大小
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}

	rmat := make([][]int, r)
	for i := 0; i < r; i++ {
		rmat[i] = make([]int, c)
		for j := 0; j < c; j++ {
			idx := i*c + j
			si := idx / len(mat[0])
			sj := idx % len(mat[0])
			rmat[i][j] = mat[si][sj]
		}
	}

	return rmat
}
