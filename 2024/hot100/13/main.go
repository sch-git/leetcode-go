package main

func main() {

}

func maxSubArray(nums []int) int {
	curSum := nums[0]
	ans := curSum
	for i := 1; i < len(nums); i++ {
		if curSum >= 0 {
			curSum += nums[i]
		} else {

			curSum = maxNum(curSum, nums[i])
		}
		ans = maxNum(ans, curSum)
	}
	return ans
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
