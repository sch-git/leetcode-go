package _64

// middle

func sumNums(n int) int {
	var sum func(int) bool
	ans := 0
	sum = func(i int) bool {
		ans += i
		return i > 0 && sum(i-1)
	}
	sum(n)
	return ans
}
