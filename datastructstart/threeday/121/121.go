package _121

func maxProfit(prices []int) int {
	min := prices[0]
	max := 0
	for left := 0; left < len(prices); left++ {
		if prices[left] < min {
			min = prices[left]
		}
		if prices[left]-min > max {
			max = prices[left] - min
		}
	}

	return max
}
