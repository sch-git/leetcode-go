package _58

// simple

func reverseLeftWords(s string, n int) string {
	n = n % len(s)
	pre := make([]rune, 0)
	res := make([]rune, 0)
	for idx, str := range s {
		if idx >= n {
			res = append(res, str)
		} else {
			pre = append(pre, str)
		}
	}

	return string(append(res, pre...))
	//return s[n:]+s[:n]
}
