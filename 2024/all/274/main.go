package main

import (
	"sort"
)

func main() {

}

func hIndex(citations []int) int {
	sort.Ints(citations)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	h := 0
	for idx, val := range citations {
		cur := min(len(citations)-idx, val)
		h = max(h, cur)
	}
	return h
}
