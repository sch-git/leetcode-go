package quicksort

import (
	"testing"
)

func TestName(t *testing.T) {
	nums := []int{1, 2, 3, 4, 6, 5, 8, 7}
	length := len(nums) - 1
	quicksort(nums, 0, length)
	//sort.Ints(nums)
	t.Log(nums)
}
