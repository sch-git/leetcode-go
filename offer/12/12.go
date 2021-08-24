package _12

// middle
// 回溯

func Exist(board [][]byte, word string) bool {
	idx := 0
	for ridx, row := range board {
		for cidx, val := range row {
			if val == word[0] {
				res := findWord(board, word, ridx, cidx, idx)
				if res {
					return true
				}
			}
		}
	}

	return false
}

func findWord(board [][]byte, word string, ridx, cidx, idx int) bool {
	if ridx < 0 || cidx < 0 || ridx >= len(board) || cidx >= len(board[0]) || word[idx] != board[ridx][cidx] {
		return false
	}

	if board[ridx][cidx] == ' ' {
		return false
	}

	idx++
	board[ridx][cidx] = ' '
	if idx >= len(word) {
		return true
	}

	// 剪枝，一个成功便返回，不需要全部搜索
	if findWord(board, word, ridx-1, cidx, idx) || findWord(board, word, ridx+1, cidx, idx) || findWord(board, word, ridx, cidx-1, idx) || findWord(board, word, ridx, cidx+1, idx){
		return true
	}

	// 回溯，将当前位置重置
	board[ridx][cidx] = word[idx-1]
	return false
}
