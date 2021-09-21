package _015

// middle
// 滑动窗口

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}

	m := make(map[rune]int)
	windows := 0
	res := make([]int, 0)

	for _, val := range p {
		m[val]++
	}

	for idx, val := range s {
		windows++
		m[val]--
		if windows == len(p) {
			flag := true
			for _, num := range m {
				if num != 0 {
					flag = false
					break
				}
			}
			if flag {
				res = append(res, idx+1-windows)
			}

			m[rune(s[idx+1-windows])]++
			windows--
		}
	}
	return res
}
