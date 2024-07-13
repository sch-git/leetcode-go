package main

func main() {

}

func gameOfLife(board [][]int) {
	changeIdxI, changeIdxJ := make([]int, 0), make([]int, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			liveCount := countLive(i, j, &board)
			if board[i][j] == 1 {
				if liveCount < 2 || liveCount > 3 {
					changeIdxI = append(changeIdxI, i)
					changeIdxJ = append(changeIdxJ, j)
				}
			} else {
				if liveCount == 3 {
					changeIdxI = append(changeIdxI, i)
					changeIdxJ = append(changeIdxJ, j)
				}
			}
		}
	}
	for i := 0; i < len(changeIdxI); i++ {
		bi, bj := changeIdxI[i], changeIdxJ[i]
		if board[bi][bj] == 1 {
			board[bi][bj] = 0
		} else {
			board[bi][bj] = 1
		}
	}
	return
}

func countLive(i, j int, board *[][]int) int {
	return isLive(i-1, j-1, board) + isLive(i-1, j, board) + isLive(i-1, j+1, board) +
		isLive(i, j-1, board) + isLive(i, j+1, board) +
		isLive(i+1, j-1, board) + isLive(i+1, j, board) + isLive(i+1, j+1, board)
}

func isLive(i, j int, board *[][]int) int {
	if i < 0 || j < 0 || i >= len(*board) || j >= len((*board)[0]) {
		return 0
	}

	if (*board)[i][j] == 1 {
		return 1
	}
	return 0
}
