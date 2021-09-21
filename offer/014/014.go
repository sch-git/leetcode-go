package _014

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	m := make(map[rune]int)
	windows := 0

	for i := 0; i < len(s1); i++ {
		m[rune(s1[i])]++
	}

	for idx, ch := range s2 {
		windows++
		m[ch]--
		if windows == len(s1) {
			count := 0
			for _, val := range m {
				count += val
				if val != 0 {
					count = -1
					break
				}
			}
			if count == 0 {
				return true
			}

			m[rune(s2[idx+1-windows])]++
			windows--
		}
	}

	return false
}
