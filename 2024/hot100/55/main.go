package main

func main() {
	permute([]int{1, 2, 3})
}

func permute(nums []int) [][]int {
	var ans [][]int = make([][]int, 0)
	cal(make([]int, 0), nums, &ans)
	return ans
}

func cal(pre, cur []int, ans *[][]int) {
	if len(cur) == 0 {
		*ans = append(*ans, pre)
		return
	}

	for i := 0; i < len(cur); i++ {
		curPre := append(pre, cur[i])
		cal(curPre, append(append(make([]int, 0, len(cur)), cur[0:i]...), cur[i+1:]...), ans)
	}

	return
}
