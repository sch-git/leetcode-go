package _18

import "sort"

// middle
// 双指针

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	sort.Ints(nums)
	res := make([][]int, 0)
	for idx, _ := range nums {
		if idx > 0 && nums[idx] == nums[idx-1] {
			continue
		}
		for i := idx + 1; i < len(nums)-1; i++ {
			if i > idx+1 && nums[i] == nums[i-1] {
				continue
			}

			l, r := i+1, len(nums)-1
			for l < r {
				sum := nums[idx] + nums[i] + nums[l] + nums[r]
				if sum == target {
					res = append(res, []int{nums[idx], nums[i], nums[l], nums[r]})
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
				} else {
					l++
				}
			}
		}
	}
	return res
}
