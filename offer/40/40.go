package _40

// simple
// 快排
// TODO 堆

func getLeastNumbers(arr []int, k int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr[:k]
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	i, j := l, r
	for i < j {		// 如果先从左向右，首个数就是最小的数的情况就会有问题
		for i < j && arr[j] >= arr[l] {	// 从右向左，找到首个比基准数大的数
			j--
		}
		for i < j && arr[i] <= arr[l] { // 从左向右，找到首个比基准数小的数
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[l], arr[i] = arr[i], arr[l]
	quickSort(arr, l, i-1)
	quickSort(arr, i+1, r)
	return
}
