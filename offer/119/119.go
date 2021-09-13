package _119

// middle
// 最长连续序列

func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	max := 0
	for _, num := range nums {
		m[num] = 1
		max = 1
	}

	for _, num := range nums {
		if m[num] >= 2 {
			continue
		}

		temp := 0
		val, ok := m[num]
		for ok {
			if val >= 2 {
				temp += val
				m[num] = temp
				max = maxNum(temp, max)
				break
			}
			temp++
			m[num] = temp
			max = maxNum(temp, max)
			num++
			_, ok = m[num]
		}
	}

	return max
}

func maxNum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
