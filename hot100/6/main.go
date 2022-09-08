package main

func main() {
	convert("HA", 2)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	slice := make([][]rune, numRows)
	idx := 0
	flag := false
	for _, ch := range s {
		slice[idx] = append(slice[idx], ch)
		if idx >= numRows-1 || idx <= 0 {
			flag = !flag
		}

		if flag {
			idx++
		} else {
			idx--
		}
	}

	res := ""
	for _, chs := range slice {
		for _, ch := range chs {
			res += string(ch)
		}
	}

	return res
}
