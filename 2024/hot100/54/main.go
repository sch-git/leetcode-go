package main

func main() {

}

func maxProduct(nums []int) int {
	ans := nums[0]
	fmin := make([]int, len(nums))
	fmin[0] = nums[0]
	for i, num := range nums {
		if i > 0 {
			nums[i] = maxNum(num, maxNum(num*nums[i-1], num*fmin[i-1]))
			fmin[i] = minNum(num, minNum(num*fmin[i-1], num*nums[i-1]))
			if ans < nums[i] {
				ans = nums[i]
			}
		}
	}
	return ans
}
func maxNum(a, b int) int {
	if a > b {
		return a

	}
	return b
}
func minNum(a, b int) int {
	if a > b {
		return b

	}
	return a
}
