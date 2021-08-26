package _45

import (
	"sort"
	"strconv"
	"strings"
)

// middle
// 排序

type numSort []int

func minNumber(nums []int) string {
	numsSort := numSort(nums)
	sort.Sort(numsSort)
	sh := strings.Builder{}
	for _, v := range numsSort {
		sh.WriteString(strconv.Itoa(v))
	}
	return sh.String()
}

func (nums numSort) Len() int {
	return len(nums)
}

func (nums numSort) Less(i, j int) bool {
	tempi, tempj := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])

	ij, _ := strconv.Atoi(tempi + tempj)
	ji, _ := strconv.Atoi(tempj + tempi)

	return ij < ji
}

func (nums numSort) Swap(i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
	return
}
