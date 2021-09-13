package _006

// simple

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l, r}
		} else if sum > target {
			r--
		} else if sum < target {
			l++
		}
	}

	return nil
}
