package _22

// middle
// 回溯

var res []string

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	seleted := make([]rune, 0)
	create(n*2, n, n, seleted)
	return res
}

func create(n, ln, rn int, selected []rune) {
	if n < 1 {
		res = append(res, string(selected))
		return
	}

	// 判断是否选择左括号
	if ln == rn || ln > 0 {
		selected = append(selected, '(')
		ln--
		create(n-1, ln, rn, selected)
		// 重置
		selected = selected[:len(selected)-1]
		ln++
	}

	// 判断是否选择右括号
	if rn > ln && rn > 0 {
		selected = append(selected, ')')
		rn--
		create(n-1, ln, rn, selected)
	}

	return
}
