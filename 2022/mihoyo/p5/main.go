package main

func main() {
	print(calculate(" 2-1 + 2 "))
}

func calculate(s string) int {
	ops := []int{1}
	sign ,ans:= 1,0
	n:=len(s)
	for i := 0;i<len(s);{
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = ops[len(ops)-1]
			i++
		case '-':
			sign = -ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops,sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num :=0
			for ;i<n && s[i]>='0' && s[i]<='9';i++{
				num = num*10+(int(s[i]-'0'))
			}
			ans += sign*num
		}
	}
	return ans
}
