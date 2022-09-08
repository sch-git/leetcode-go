package main

func main() {
	calculate("3+2*2")
}

func calculate(s string) int {
	op := '+'
	res := 0
	stack := make([]int, 0)
	num := 0
	for idx, ch := range s {
		isNum := ch >= '0' && ch <= '9'
		if isNum {
			num = 10*num + int(ch-'0')
		}
		if !isNum && ch != ' ' || idx == len(s)-1 {
			switch op {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '/':
				stack[len(stack)-1] /= num
			case '*':
				stack[len(stack)-1] *= num
			}
			op = ch
			num = 0
		}
	}
	for _, n := range stack {
		res += n
	}
	return res

}
