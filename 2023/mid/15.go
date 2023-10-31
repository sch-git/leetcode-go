package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
}

func threeSum(nums []int) [][]int {
	left, cur, right := 0, 1, len(nums)-1
	sort.Ints(nums)
	fmt.Println(nums)
	res := make([][]int, 0)
	m := make(map[string]bool)
	for left < right {
		for cur < right {
			count := nums[left] + nums[cur] + nums[right]
			if count > 0 {
				break
			} else if count < 0 {
				cur++
			} else if count == 0 {
				key := fmt.Sprintf("%d-%d-%d", nums[left], nums[cur], nums[right])
				if !m[key] {
					res = append(res, []int{nums[left], nums[cur], nums[right]})
					m[key] = true
				}
				cur++
				right--
			}
		}
		left++
		cur = left + 1
		right = len(nums) - 1
	}
	return res
}
