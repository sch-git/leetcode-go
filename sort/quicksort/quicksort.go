package quicksort

// 快速排序

func quicksort(nums []int, l, r int) {
	if l > r {
		return
	}
	idx := findIndex(nums, l, r)
	nums[idx], nums[l] = nums[l], nums[idx]
	quicksort(nums, l, idx-1)
	quicksort(nums, idx+1, r)
}

func findIndex(nums []int, l, r int) int {
	idx := l
	for l < r {
		for l < r && nums[idx] >= nums[l] {
			l++
		}
		for l < r && nums[idx] <= nums[r] {
			r--
		}

		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

	return r // 返回得是 l 或 r 跟循环中是先移动 l 指针和 r 指针有关；如果返回先移动的 l，在最后 l 进行了++，存在越界风险；反之亦然。
}
