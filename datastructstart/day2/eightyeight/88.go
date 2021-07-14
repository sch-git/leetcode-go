package eightyeight

func Merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) < 1 {
		nums1 = nums2
		return
	}
	if len(nums2) < 1 {
		return
	}

	ns := make([]int, m)
	for i := 0; i < m; i++ {
		ns[i] = nums1[i]
	}

	var i, j int
	for idx := 0; idx < m+n; idx++ {
		if i >= m {
			for step, num := range nums2[j:] {
				nums1[idx+step] = num
			}
			break
		}

		if j >= n {
			for step, num := range ns[i:] {
				nums1[idx+step] = num
			}
			break
		}

		if ns[i] > nums2[j] {
			nums1[idx] = nums2[j]
			j++
		}else{
			nums1[idx] = ns[i]
			i++
		}
	}

	return
}
