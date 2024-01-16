package main

import (
	"strings"
)

func main() {

}

func minRemoveToMakeValid(s string) string {
	queue := make([]rune, 0)
	ans := strings.Builder{}
	for _, ch := range s {
		if ch != '(' && ch != ')' && len(queue) == 0 {
			ans.WriteRune(ch)
			continue
		}

		queue = append(queue, ch)
		if ch == ')' {
			if len(queue) > 0 && queue[0] == '(' {
				ans.WriteRune(queue[0])
				queue = queue[1:]
				for len(queue) > 0 && queue[0] != '(' {
					ans.WriteRune(queue[0])
					queue = queue[1:]
				}
			} else {
				queue = queue[1:]
			}
		}
	}

	for _, ch := range queue {
		if ch == '(' {
			continue
		}
		ans.WriteRune(ch)
	}

	return ans.String()
}
