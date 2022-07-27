package main

func main() {
	lengthOfLIS([]int{1,3,6,7,9,4,10,5,6})
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	res := 0
	for i:=0;i<len(nums);i++{
		dp[i] = 1
		for j:=0;j<i; j++{
			if nums[i]>nums[j]{
				dp[i] = maxNum(dp[j]+1,dp[i])
			}
		}
		res = maxNum(dp[i],res)
	}
	return res
}

func maxNum(a,b int) int {
	if a>b{
		return a
	}
	return b
}