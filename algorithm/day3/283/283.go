package _83

func moveZeroes(nums []int) {
	move := false
	moveIdx := 0
	for idx, num := range nums {
		if num == 0 {
			move = true
			moveIdx++
		} else if move {
			nums[idx-moveIdx] = num
			nums[idx] = 0
		}
	}

	return
}
