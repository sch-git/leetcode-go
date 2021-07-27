package _67

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	m := make([]int, 26)
	for _, ch := range s1 {
		m[ch-'a']--
	}

	left := 0
	for right, ch := range s2 {
		x := ch - 'a'
		m[x]++
		for m[x] > 0 {
			m[s2[left]-'a']--
			left++
		}
		if right-left+1 == len(s1) {
			return true
		}
	}

	return false
}
