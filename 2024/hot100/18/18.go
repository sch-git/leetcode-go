package main

func main() {

}

func setZeroes(matrix [][]int) {
	row := make(map[int]bool)
	col := make(map[int]bool)
	for m := 0; m < len(matrix); m++ {
		for n := 0; n < len(matrix[m]); n++ {
			if matrix[m][n] == 0 {
				row[m] = true
				col[n] = true
			}
		}
	}
	for m := 0; m < len(matrix); m++ {
		for n := 0; n < len(matrix[m]); n++ {
			if row[m] || col[n] {
				matrix[m][n] = 0
			}
		}
	}
	return
}
