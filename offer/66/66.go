package _66

func constructArr(a []int) []int {
	if len(a)<1{
		return nil
	}
	length := len(a)
	res := make([]int, length)
	res[0] = 1
	for i := 1; i < length; i++ {
		res[i] = res[i-1] * a[i-1]
	}

	tmp := 1
	for i := length - 2; i >= 0; i-- {
		tmp *= a[i+1]
		res[i] *= tmp
	}

	return res
}
