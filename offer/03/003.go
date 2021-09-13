package _03

// simple
// 如果是奇数，则前一位数字的值 +1；如果是偶数，则值 =num/2 的值
// 2 (0010), 4 (0100), 8 (1000)： 1 的个数不变，位置变化

func countBits(n int) []int {
	if n<1{
		return []int{0}
	}
	res := make([]int, n+1)
	res[0] = 0
	res[1] = 1

	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i-1] + 1
		}
	}
	return res
}
