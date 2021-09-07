package _38

var (
	res    []string
)

func permutation(s string) []string {
	if s == "" {
		return nil
	}

	visit := make(map[int]bool, 0)
	res = make([]string, 0)

	back(make([]rune, 0), visit, s)
	return res
}

func back(selected []rune, visited map[int]bool, s string) {
	if len(selected) == len(s) {
		res = append(res, string(selected))
		return
	}

	firstSelected := make(map[rune]int, 0)
	for idx, ch := range s {
		if visited[idx] {
			continue
		}

		if firstSelected[ch] != 0 && firstSelected[ch] == len(selected)+1 {
			continue
		}

		firstSelected[ch] = len(selected) + 1
		visited[idx] = true
		selected = append(selected, ch)
		back(selected, visited, s)
		selected = selected[:len(selected)-1]
		visited[idx] = false
	}
}
