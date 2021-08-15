package _53

// simple

// 将数组分为两部分，跳出循环时： left 指向右数组首位元素，right 指向左数组某位元素

func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if mid != nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
