package _89

// simple

func Rotate(nums []int, k int) {
	if k == 0 || k == len(nums) {
		return
	}

	if k > len(nums) {
		k = k % len(nums)
	}

	temp := nums[len(nums)-k:]
	for idx := 0; idx < len(nums)-k; idx++ {
		temp = append(temp, nums[idx])
	}

	for idx, num := range temp {
		nums[idx] = num
	}

	return
}
