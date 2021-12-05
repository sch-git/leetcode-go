package _0102

// simple

func main() {

}

// 判定是否互为字符重排
func CheckPermutation(s1 string, s2 string) bool {
	m := make(map[rune]int)

	if len(s1) != len(s2) {
		return false
	}

	for _, v := range s1 {
		m[v]++
	}

	for _, v := range s2 {
		m[v]--
		if m[v] < 0 {
			return false
		}
	}

	return true
}
