package main

func main() {
	candy([]int{1, 6, 10, 8, 7, 3, 2})
	// 1 2 5 4 3 2 1
	// 1 2 3 1 2 4 5
}

func candy(ratings []int) int {
	if len(ratings) <= 1 {
		return len(ratings)
	}

	var sum, cur int = 1, 1
	listLen := 0
	curMax := 0
	//for idx, _ := range ratings {
	//	flag := false
	//	if idx == 0 {
	//		cur = 1
	//	}
	//	if idx > 0 {
	//		if ratings[idx] > ratings[idx-1] {
	//			listLen = 0
	//			cur++
	//			curMax = cur
	//		} else {
	//			if ratings[idx] == ratings[idx-1] {
	//				listLen = 0
	//				curMax = 0
	//			} else {
	//				listLen++
	//			}
	//			flag = cur == 1 && ratings[idx] < ratings[idx-1]
	//			cur = 1
	//		}
	//	}
	//	sum += cur
	//	if flag {
	//		sum += listLen
	//		if listLen < curMax {
	//			sum--
	//		}
	//	}
	//}

	// 代码优化
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			cur++
			curMax = cur
			listLen = 0
		} else if ratings[i] == ratings[i-1] {
			cur = 1
			curMax = 0
			listLen = 0
		} else {
			listLen++
			if cur == 1 {
				sum += listLen
				if listLen < curMax {
					sum--
				}
			}
			cur = 1
		}
		sum += cur
	}
	return sum
}
