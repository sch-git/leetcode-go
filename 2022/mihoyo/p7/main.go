package main

func main() {

}

func balancedStringSplit(s string) int {
	res, count := 0, 0
	for _, ch := range s {
		if ch == 'L' {
			count--
		} else if ch == 'R' {
			count++
		}
		if count == 0 {
			res++
		}
	}

	return res
}
