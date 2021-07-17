package _83

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, ch := range magazine {
		m[ch]++
	}

	for _, ch := range ransomNote {
		m[ch]--
		if m[ch] < 0 {
			return false
		}
	}

	return true
}
