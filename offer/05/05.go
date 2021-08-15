package _05

import "strings"

// simple

// 优化：原地算法，根据空格数量拓展原切片，从后往前遍历

func replaceSpace(s string) string {
	res := strings.Builder{}
	for _, str := range s {
		if rune(str) == ' '{
			res.WriteString("%20")
		}else{
			res.WriteString(string(str))
		}
	}
	return res.String()
}
