package _007

import "sort"

// middle
// 思路：固定一个元素 i，变为求两数之和 在 i+1...nums.length-1 范围内查询和为 -nums[i] 的两个数
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); {
		target := -nums[i]
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[l] + nums[r]
			if sum == target {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum > target {
				r--
			} else if sum < target {
				l++
			}
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
