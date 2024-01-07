package main

func main() {
	rotate([]int{1, 2, 3}, 2)
	// 2341
	// 432
}
func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}

	right := len(nums) - 1
	knums := make([]int, 0)
	for left := right - k; left >= 0; {
		if len(knums) < k {
			knums = append(knums, nums[right])
		}
		nums[right] = nums[left]
		right--
		left = right - k
	}

	for len(knums) < k {
		knums = append(knums, nums[right])
		right--
	}
	for i := 0; i < k; i++ {
		nums[i] = knums[len(knums)-1-i]
	}
	return
}
