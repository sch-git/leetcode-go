package _350

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, num := range nums1 {
		m[num]++
	}

	res := make([]int, 0)
	for _, num := range nums2 {
		if m[num] > 0 {
			m[num]--
			res = append(res, num)
		}
	}

	return res
}
