package main

func main() {
	minWindow("ADOBECODEBANC", "ABC")
}

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	ans := ""
	left, right := 0, 0
	tchs := make(map[rune]int)
	for _, ch := range t {
		tchs[ch]++
	}

	isContainer := func(t, s map[rune]int) bool {
		for k, v := range t {
			if s[k] < v {
				return false
			}
		}
		return true
	}
	schs := make(map[rune]int)
	for ; right < len(s); right++ {
		schs[rune(s[right])]++
		for right-left >= len(t)-1 && tchs[rune(s[right])] > 0 && isContainer(tchs, schs) {
			if ans == "" || len(s[left:right+1]) < len(ans) {
				ans = s[left : right+1]
			}

			schs[rune(s[left])]--
			left++
		}
	}
	return ans
}
