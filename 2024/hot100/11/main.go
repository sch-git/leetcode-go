package main

func main() {

}

func maxSlidingWindow(nums []int, k int) []int {
	q := make([]int, 0)
	for i := 0; i < k; i++ {
		for len(q) > 0 && nums[i] > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	ans := make([]int, 0)
	ans = append(ans, nums[q[0]])
	for i := k; i < len(nums); i++ {
		for len(q) > 0 && nums[i] > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)

		for len(q) > 0 && q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}
