package main

func main() {

}

// 递推的方式
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		dp[i] = make([]int, len(nums2))
	}
	//初始化第一行
	if nums1[0] == nums2[0] {
		dp[0][0] = 1
	}
	for j := 1; j < len(nums2); j++ {
		if nums1[0] == nums2[j] {
			dp[0][j] = 1
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}
	// 初始化第一列
	for i := 1; i < len(nums1); i++ {
		if nums2[0] == nums1[i] {
			dp[i][0] = 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}

	for i := 1; i < len(nums1); i++ {
		for j := 1; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(nums1)-1][len(nums2)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
