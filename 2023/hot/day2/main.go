package main

import (
	"fmt"
)

// 使用 [26]int 数组记录字符串中的字母数量，字母异味词的对应字母数量一定相同，所以可以用[26]int 作为 map 的 key
func groupAnagrams(strs []string) [][]string {
	wordsNum2Key := make(map[[26]int][]string)
	for _, str := range strs {
		chs := [26]int{}
		for _, ch := range str {
			chs[ch-'a']++
		}
		wordsNum2Key[chs] = append(wordsNum2Key[chs], str)
	}

	res := make([][]string, 0)
	for _, val := range wordsNum2Key {
		res = append(res, val)
	}
	return res
}

func main() {
	trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
}

func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		fmt.Println(min(leftMax[i], rightMax[i]), "-", height[i], "-", min(leftMax[i], rightMax[i])-height[i])
		ans += min(leftMax[i], rightMax[i]) - height[i]

	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
