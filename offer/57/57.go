package _57

func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	if len(nums) < 1 {
		return res
	}

	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left]+nums[right] == target {
			return []int{nums[left], nums[right]}
		} else if nums[left]+nums[right] > target {
			right--
		} else {
			left++
		}
	}

	return res
}
