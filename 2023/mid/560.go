package main

func main() {
	subarraySum([]int{1, -1, 0}, 0)
}

// [[0],[1,-1],[1,-1,0]]
func subarraySum(nums []int, k int) int {
	var ans int
	mp := make(map[int]int)
	mp[0] = 1 // 初始化前缀和为0的次数为1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := mp[sum-k]; ok {
			ans += mp[sum-k]
		}

		mp[sum]++
	}
	return ans
}
