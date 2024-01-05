package main

func main() {
	findAnagrams("cbaebabacd", "abc")
}
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	ans := make([]int, 0)
	l, r := 0, 0
	windows := [26]int{}
	pw := [26]int{}
	for _, ch := range p {
		pw[ch-'a']++
	}

	for ; r < len(s); r++ {
		windows[s[r]-'a']++
		for r-l+1 == len(p) {
			//判断窗口是否满足条件
			if windows == pw {
				ans = append(ans, l)
			}
			windows[s[l]-'a']--
			l++
		}
	}
	return ans

}
