package day1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {
		if val, ok := m[target-num]; ok {
			return []int{idx, val}
		}
		m[num] = idx
	}

	return []int{}
}
