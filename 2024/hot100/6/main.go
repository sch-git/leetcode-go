package main

import (
	"sort"
)

func main() {
	threeSum([]int{-1, 0, 1, 2, -1, -4})
}

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)

	left, right := 0, len(nums)-1
	for ; left < right; left++ {
		cur := left + 1
		if left > 0 && nums[left] == nums[left-1] { // 和上次枚举的数不同
			continue
		}

		curRight := right
		for cur < curRight {
			count := nums[left] + nums[cur] + nums[curRight]
			if count == 0 {
				ans = append(ans, []int{nums[left], nums[cur], nums[curRight]})
				for cur++; cur < curRight && nums[cur] == nums[cur-1]; cur++ {
				} // 跳过重复数字
				for curRight--; cur < curRight && nums[curRight] == nums[curRight+1]; curRight-- {
				} // 跳过重复数字
			} else if count < 0 {
				cur++
			} else {
				curRight--
			}
		}
	}
	return ans
}
