package _44

// simple

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for ; left < right; {
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		left++
		right--
	}

	return
}

// 官方题解
//func reverseString(s []byte) {
//	for left, right := 0, len(s)-1; left < right; left++ {
//		s[left], s[right] = s[right], s[left]
//		right--
//	}
//}

//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/reverse-string/solution/fan-zhuan-zi-fu-chuan-by-leetcode-solution/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
