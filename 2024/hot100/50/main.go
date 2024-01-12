package main

func main() {

}

func rob(nums []int) int {
	// 设 f(n) 为第n个位置能偷取的最大数，则f(n) = mac(n+f(n-2), f(n-1))
	if len(nums) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	if len(nums) <= 2 {
		return max(nums[0], nums[len(nums)-1])
	}
	nums[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		nums[i] = max(nums[i]+nums[i-2], nums[i-1])
	}
	return nums[len(nums)]
}
