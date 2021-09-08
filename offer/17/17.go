package _17

import "math"

// simple

func printNumbers(n int) []int {
	length := int(math.Pow10(n))
	res := make([]int, length)
	for i := 1; i < length; i++ {
		res[i-1] = i
	}

	return res
}
