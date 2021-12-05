package _0105

import "math"

// middle
// 字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
func oneEditAway(first string, second string) bool {
	if math.Abs(float64(len(first)-len(second))) > 1 {
		return false
	}

	arr := make([][]int, len(first)+1)
	for k, _ := range arr {
		arr[k] = make([]int, len(second)+1)
	}

	// 初始化
	for i := 0; i <= len(first); i++ {
		arr[i][0] = i
	}

	for i := 0; i <= len(second); i++ {
		arr[0][i] = i
	}

	for m := 1; m <= len(first); m++ {
		for n := 1; n <= len(second); n++ {
			if first[m-1] == second[n-1] {
				arr[m][n] = arr[m-1][n-1]
			} else {
				arr[m][n] = int(math.Min(math.Min(float64(arr[m-1][n]), float64(arr[m][n-1])), float64(arr[m-1][n-1]))) + 1
			}
		}
	}

	return arr[len(first)][len(second)] <= 1
}
