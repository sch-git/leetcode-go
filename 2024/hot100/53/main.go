package main

func main() {

}

func lengthOfLIS(nums []int) int {
	var res int = 1
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, 0)
	for _, _ = range nums {
		dp = append(dp, 1)
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = func() int {
					if dp[i] < dp[j]+1 {
						return dp[j] + 1
					}
					return dp[i]
				}()
				if dp[i] > res {
					res = dp[i]
				}
			}
		}
	}
	return res
}
