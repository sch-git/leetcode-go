package main

func main() {

}

func maxProfit(prices []int) int {
	var res int
	pre := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > pre {
			res += prices[i] - pre
		}
		pre = prices[i]
	}
	return res
}
