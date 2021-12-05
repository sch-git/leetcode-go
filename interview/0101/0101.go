package _0101

// simple

func main() {

}


// 判定字符是否唯一
func isUnique(astr string) bool {
	m := make(map[rune]int)
	for _, v := range astr {
		m[v]++
		if m[v] > 1 {
			return false
		}
	}
	return true
}
