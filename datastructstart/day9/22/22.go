package _2

func isValid(s string) bool {
	stack := make([]rune, len(s))
	idx := 0
	for _, val := range s {
		if val == '(' || val == '{' || val == '[' {
			stack[idx] = val
			idx++
		} else {
			if idx-1 < 0 {
				return false
			}

			if val == ')' && stack[idx-1] == '(' {
				idx--
			} else if val == ']' && stack[idx-1] == '[' {
				idx--
			} else if val == '}' && stack[idx-1] == '{' {
				idx--
			} else {
				return false
			}
		}
	}

	return idx == 0
}
