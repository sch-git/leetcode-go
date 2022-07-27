package main

import (
	"strconv"
)

func main() {
	println(getPermutation(3,3))
}

func getPermutation(n int, k int) string {
	nums := make([]int,0,n)
	for i:=1;i<=n;i++{
		nums = append(nums,i)
	}

	results := make([]string,0,k)
	selected := make([]int,0,n)
	backtrack(nums,selected,&results,k)
	return results[k-1]
}

func backtrack(nums,selected []int,results *[]string,k int)  {
	if len(*results) == k {
		return
	}
	if len(nums)<1{
		var result string
		for _,num := range selected{
			result += strconv.Itoa(num)
		}
		*results=append(*results,result)
		return
	}

	for idx,num := range nums{
		selected = append(selected,num)
		tmp := append(append([]int{},nums[:idx]...),nums[idx+1:]...)
		backtrack(tmp,selected,results,k)
		selected = selected[:len(selected)-1]
	}
}
