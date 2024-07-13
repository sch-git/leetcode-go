package main

import (
	"sort"
)

func main() {

}

// 贪心+双指针
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	l, r := 0, len(people)-1
	ans := 0
	for l <= r {
		if people[l]+people[r] > limit {
			r--
		} else {
			l++
			r--
		}
		ans++
	}
	return ans
}
