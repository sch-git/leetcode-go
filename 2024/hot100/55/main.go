package main

func main() {
	longestValidParentheses("()(())")
}

func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}

	stack := make([]int, 0)
	maxLen := 0
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				if i-stack[len(stack)-1] > maxLen {
					maxLen = i - stack[len(stack)-1]
				}
			} else {
				stack = append(stack, i)
			}
		}
	}
	return maxLen
}
