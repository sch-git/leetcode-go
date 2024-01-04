package main

func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	m := make(map[int]int)
	for idx, num := range nums {
		if _, ok := m[target-num]; ok {
			return []int{m[target-num], idx}
		}
		m[num] = idx
	}

	return res
}
