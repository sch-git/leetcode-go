package _42

// simple

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)
	for _, ch := range s {
		m[ch]++
	}

	for _, ch := range t {
		m[ch]--
	}

	for _, times := range m {
		if times != 0 {
			return false
		}
	}

	return true
}
