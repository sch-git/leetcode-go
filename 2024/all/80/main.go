package main

func main() {
	// 1 1 1 1 2 2 2 2 3 3
	println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	println(removeDuplicates([]int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3}))
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	curIdx, numTimes := 1, 1
	for i := 1; i < len(nums); i++ {
		for ; i < len(nums) && nums[i] == nums[i-1]; i++ {
			numTimes++
			nums[curIdx] = nums[i]
			if numTimes <= 2 {
				curIdx++
			}
		}
		numTimes = 1
		if i >= len(nums) {
			break
		}

		nums[curIdx] = nums[i]
		curIdx++
	}
	return curIdx
}
