package main

func main() {

}

func climbStairs(n int) int {
	// f(n) 为爬到第 n 级台阶的方法数，则 f(n) = f(n-1) + f(n-2)
	fn := make(map[int]int)
	fn[0] = 1
	fn[1] = 1
	for i := 2; i <= n; i++ {
		fn[i] = fn[i-1] + fn[i-2]
	}
	return fn[n]
}
