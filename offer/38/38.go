package _38

var (
	res    []string
	filter map[string]bool
)

func permutation(s string) []string {
	if s == "" {
		return nil
	}

	visit := make(map[int]bool, 0)
	res = make([]string, 0)
	filter = make(map[string]bool)

	back(make([]rune, 0), visit, s)
	return res
}

func back(selected []rune, visited map[int]bool, s string) {
	if len(selected) == len(s) {
		if filter[string(selected)] {
			return
		}
		res = append(res, string(selected))
		filter[string(selected)] = true
		return
	}

	for idx, ch := range s {
		if visited[idx] {
			continue
		}

		visited[idx] = true
		selected = append(selected, ch)
		back(selected, visited, s)
		selected = selected[:len(selected)-1]
		visited[idx] = false
	}
}
