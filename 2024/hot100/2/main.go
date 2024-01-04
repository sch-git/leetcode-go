package main

func main() {
	groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
}
func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	m := make(map[[26]int][]string)
	for _, str := range strs {
		chs := [26]int{}
		for _, ch := range str {
			chs[ch-'a']++
		}
		if _, ok := m[chs]; !ok {
			m[chs] = make([]string, 0)
		}
		m[chs] = append(m[chs], str)
	}

	for _, ss := range m {
		res = append(res, ss)
	}
	return res
}
