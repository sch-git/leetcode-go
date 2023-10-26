package day2

// 使用 [26]int 数组记录字符串中的字母数量，字母异味词的对应字母数量一定相同，所以可以用[26]int 作为 map 的 key
func groupAnagrams(strs []string) [][]string {
	wordsNum2Key := make(map[[26]int][]string)
	for _, str := range strs {
		chs := [26]int{}
		for _, ch := range str {
			chs[ch-'a']++
		}
		wordsNum2Key[chs] = append(wordsNum2Key[chs], str)
	}

	res := make([][]string, 0)
	for _, val := range wordsNum2Key {
		res = append(res, val)
	}
	return res
}
