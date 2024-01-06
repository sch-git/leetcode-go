package main

import (
	"sort"
)

func main() {
	merge([][]int{{2, 6}, {1, 5}})
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := make([][]int, 0)
	var preSli []int
	for _, sli := range intervals {
		if len(preSli) == 0 {
			preSli = append(preSli, sli...)
		} else {
			if preSli[0] <= sli[0] && sli[0] <= preSli[1] {
				preSli = []int{preSli[0], func() int {
					if sli[1] > preSli[1] {
						return sli[1]
					}
					return preSli[1]
				}()}
			} else {
				ans = append(ans, preSli)
				preSli = sli
			}
		}
	}
	ans = append(ans, preSli)
	return ans
}
