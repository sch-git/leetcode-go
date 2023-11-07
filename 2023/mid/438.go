package main

func findAnagrams(s string, p string) []int {
	sLen, pLen := len(s), len(p)
	var ans []int
	if sLen < pLen {
		return ans
	}
	var sCount, pCount [26]int

	for i, ch := range p {
		pCount[ch-'a']++
		sCount[s[i]-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}

	return ans
}
