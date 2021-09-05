package _15

import "sort"

//  middle
// 双指针

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	res := make([][]int, 0)

	sort.Ints(nums)
	for idx, _ := range nums {
		if nums[idx] > 0 {
			return res
		}
		if idx > 0 && nums[idx] == nums[idx-1] {
			continue
		}

		l, r := idx+1, len(nums)-1
		for l < r {
			sum := nums[idx] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[idx], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
