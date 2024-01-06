package main

func main() {

}

func subarraySumV1(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		end := start
		sum := 0
		for ; end >= 0; end-- {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}

func subarraySum(nums []int, k int) int {
	count := 0
	mp := make(map[int]int)
	mp[0] = 1
	sum := 0
	for _, num := range nums {
		sum += num
		if _, ok := mp[sum-k]; ok {
			count += mp[sum-k]
		}
		mp[sum]++
	}
	return count
}
