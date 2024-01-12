package main

func main() {

}

func generate(numRows int) [][]int {
	// fmn 表示 m行 n列的数，则f(mn) = f(m-1n)+f(m-1n-1)
	ans := make([][]int, 0)
	for i := 0; i < numRows; i++ {
		row := make([]int, 0)
		for j := 0; j <= i; j++ {
			if j == 0 {
				row = append(row, 1)
			} else {
				sum := ans[i-1][j-1]
				if j < i {
					sum += ans[i-1][j]
				}
				row = append(row, sum)
			}
		}
		ans = append(ans, row)
	}
	return ans
}
