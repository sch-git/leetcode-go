package _58

import "strings"

func reverseWords(s string) string {
	if len(s) < 1 {
		return s
	}

	builder := strings.Builder{}
	res := make([]string, 0)
	for _, val := range s {
		if val == ' ' && builder.Len() > 0 {
			res = append([]string{builder.String()}, res...)
			builder = strings.Builder{}
		} else if val != ' ' {
			builder.WriteRune(val)
		}
	}
	if builder.Len() > 0 {
		res = append([]string{builder.String()}, res...)
	}

	return strings.Join(res, " ")
}
