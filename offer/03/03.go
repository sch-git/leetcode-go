package _03

// simple

// 优化：原地置换，如果第 i 个位置的数不是 i 就将其换到相应的位置，如果是则判断是否已存在重复

func findRepeatNumber(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		if m[num] {
			return num
		}
		m[num] = true
	}
	return 0
}
