package _67

// simple

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for ; left < right; {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}

		if numbers[left]+numbers[right] < target {
			left++
		}

		if numbers[left]+numbers[right] > target {
			right--
		}
	}

	return []int{}
}
