package _14

// simple

func longestCommonPrefix(strs []string) string {
	res := ""
	idx := 0
	for {
		pre := ""
		for _, v := range strs {
			if pre == "" && len(v) > idx {
				pre = string(rune(v[idx]))
			}

			if idx >= len(v) || pre != string(rune(v[idx])) {
				return res
			}
		}
		res += pre
		idx++
	}

	return res
}
