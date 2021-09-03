package _62

// simple
// 约瑟夫环

func lastRemaining(n int, m int) int {
	x := 0
	for i := 2; i <= n; i++ {
		x = (x + m) % i
	}
	return x
}
