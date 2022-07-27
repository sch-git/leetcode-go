package _800

func main() {

}

func maxAscendingSum(nums []int) int {
	res := 0
	windows := make([]int,0)
	sum := 0
	for idx := 0;idx <len(nums);idx++{
		if idx == 0{
			windows = append(windows,nums[idx])
			sum += nums[idx]
			res = maxNum(res,sum)
		}
		if nums[idx] > nums[idx-1]{
			windows=append(windows,nums[idx])
			sum += nums[idx]
			res = maxNum(res,sum)
		}else{
			windows = []int{nums[idx]}
			sum = nums[idx]
			res = maxNum(res,sum)
		}
	}
	return res
}

func maxNum(a,b int) int{
	if a>b{
		return a
	}
	return b
}