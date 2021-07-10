package one

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		if v, ok := m[target-val]; ok {
			return []int{v, index}
		}

		m[val] = index
	}
	return []int{}
}
