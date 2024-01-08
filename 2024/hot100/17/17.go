package main

func main() {
	firstMissingPositive([]int{1, 2, 0})
}

func firstMissingPositive(nums []int) int {
	minNum := 1
	mp := make(map[int]int)
	for _, num := range nums {
		if num > 0 {
			mp[num] = num + 1
		}
	}

	_, ok := mp[minNum]
	for ok {
		minNum = mp[minNum]
		_, ok = mp[minNum]
	}
	return minNum
}
