package _56

// middle
// 位运算
// 分组异或

func singleNumbers(nums []int) []int {
	x := 0
	for _, num := range nums {
		x ^= num
	}

	div := 1
	for div&x == 0 {
		div <<= 1
	}

	a, b := 0, 0
	for _, num := range nums {
		if num&div == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}

	return []int{a, b}
}
